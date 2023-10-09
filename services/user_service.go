package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math/big"
	"net/http"
	"os"
	repository "searchRecommend/repositories"
	"searchRecommend/schema"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func (userservice *UserService) CreateUserService(userDto schema.UserDto) (bool, error) {

	var user schema.UserSchema
	user.EMAIL = userDto.EMAIL
	user.FIRST_NAME = userDto.FIRST_NAME
	user.LAST_NAME = userDto.LAST_NAME

	// *****HASHING THE PASSWORD START*****
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err.Error())
	}
	saltString := hex.EncodeToString(salt)
	user.SALT = saltString

	saltPass := saltString + userDto.PASSWORD
	hash := sha256.New()
	hash.Write([]byte(saltPass))

	hashedPass := hex.EncodeToString(hash.Sum(nil))
	user.HASH_PASSWORD = hashedPass
	// *****HASHING THE PASSWORD END*****

	// *****GENERATING RANDOM 10 DIGIT USERID START*****
	const charset = "0123456789"
	var result strings.Builder
	for i := 0; i < 10; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result.WriteByte(charset[index.Int64()])
	}
	id, errr := strconv.Atoi(result.String())
	if errr != nil {
		panic(errr.Error())
	}
	user.ID = id
	// *****GENERATING RANDOM 10 DIGIT USERID END*****

	log.Println(user)
	exists := userservice.UserRepo.CreateUserQuery(user)

	if exists {
		log.Println("User Created")
	} else {
		log.Println("Email alredy exists")
	}
	return exists, nil

}

func (userservice *UserService) LoginUserService(user schema.UserLoginDto) (bool, schema.Error) {

	hashPassDB, salt := userservice.UserRepo.LoginUserQuery(user)

	var error schema.Error

	if hashPassDB != "" && salt != "" {
		password := salt + user.PASSWORD
		hash := sha256.New()
		hash.Write([]byte(password))
		hashedPassUser := hex.EncodeToString(hash.Sum(nil))

		if hashedPassUser == hashPassDB {
			return true, error
		} else {
			error.CODE = 401
			error.STATUSTEXT = http.StatusText(error.CODE)
			error.MESSAGE = "Incorrect Password! ⚠️"

			return false, error
		}

	}

	error.CODE = 404
	error.STATUSTEXT = http.StatusText(error.CODE)
	error.MESSAGE = "User does not exist!"
	log.Println("User does not exists")
	return false, error

}

func (userservice *UserService) CreateJWT(email string) (string, string, error) {
	if err := godotenv.Load(); err != nil {

		panic(err.Error())
	}

	accessSecret := os.Getenv("ACCES_TOKEN_SECRET")
	refreshSecret := os.Getenv("REFRESH_TOKEN_SECRET")

	admin := userservice.UserRepo.CreateJWTQuery(email)

	accessTokenClaims := jwt.MapClaims{
		"exp":  time.Now().Add(time.Minute).Unix(), // Token will expire in 1 minute
		"iat":  time.Now().Unix(),
		"role": admin,
		// Add other claims as needed
	}

	refreshTokenClaims := jwt.MapClaims{
		"iat":  time.Now().Unix(),
		"role": admin,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	accessTokeStr, err := accessToken.SignedString([]byte(accessSecret))
	if err != nil {
		panic(err.Error())
	}

	refreshTokenStr, err1 := refreshToken.SignedString([]byte(refreshSecret))
	if err1 != nil {
		panic(err1.Error())
	}
	userservice.UserRepo.StoreRefreshTokenQuery(refreshTokenStr, email)

	return accessTokeStr, refreshTokenStr, nil
}

func (userservice *UserService) RefreshTokenService(token string) (bool, string) {

	log.Println("token:", token)
	check := userservice.UserRepo.RefreshTokenQuery(token)

	log.Println("Service:", check)
	if !check {
		return check, ""
	}
	if err := godotenv.Load(); err != nil {

		panic(err.Error())
	}

	secret := os.Getenv("REFRESH_TOKEN_SECRET")

	checkToken, errr := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Println("Incorrect Token!")
		}
		return []byte(secret), nil
	})
	if errr != nil {
		panic(errr.Error())
	}

	var accessTokeStr string
	var err1 error
	if checkToken.Valid {

		accessSecret := os.Getenv("ACCES_TOKEN_SECRET")

		role := checkToken.Claims.(jwt.MapClaims)["role"].(bool)
		accessTokenClaims := jwt.MapClaims{
			"exp":  time.Now().Add(time.Minute).Unix(), // Token will expire in 1 minute
			"iat":  time.Now().Unix(),
			"role": role,
			// Add other claims as needed
		}

		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

		accessTokeStr, err1 = accessToken.SignedString([]byte(accessSecret))
		if err1 != nil {
			panic(err1.Error())
		}

	}
	return true, accessTokeStr

}

func (userservice *UserService) LogoutService(email string) bool {

	check := userservice.UserRepo.LogoutQuery(email)
	return check
}
