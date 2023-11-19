package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math/big"
	"net/http"
	"os"
	repository "searchRecommend/auth/repositories"
	"searchRecommend/auth/schema"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func (bookService *UserService) DbService() bool {

	count, err := bookService.UserRepo.QueryCount()
	if err != nil {
		panic(err)
	}
	//9938 is number of books present in booktwo table
	if count == 9938 {
		return true
	}

	return false
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
	max := big.NewInt(10)
	number := int64(0)

	for i := 0; i < 10; i++ {
		digit, err := rand.Int(rand.Reader, max)
		if err != nil {
			log.Fatalf(err.Error())
		}
		number = number*10 + digit.Int64()
	}
	user.ID = number
	// *****GENERATING RANDOM 10 DIGIT USERID END*****

	exists := userservice.UserRepo.CreateUserQuery(user)

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
	return false, error

}

func (userservice *UserService) CreateJWT(email string) (string, string, error) {

	envPath := os.Getenv("API_KEY")

	log.Println("ENV_PATH", envPath)
	if envPath == "" {
		if errr := godotenv.Load("/home/aamir/Desktop/My_code/searchRecommend/.env"); errr != nil {

			panic(errr.Error())
		}
	}

	accessSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	refreshSecret := os.Getenv("REFRESH_TOKEN_SECRET")

	admin, user_id := userservice.UserRepo.CreateJWTQuery(email)
	log.Println("User_ID:", user_id)

	accessTokenClaims := jwt.MapClaims{
		"exp":     time.Now().Add(5 * time.Minute).Unix(), // Token will expire in 1 minute
		"iat":     time.Now().Unix(),
		"role":    admin,
		"user_id": user_id,
		// Add other claims as needed
	}

	refreshTokenClaims := jwt.MapClaims{
		"iat":     time.Now().Unix(),
		"role":    admin,
		"user_id": user_id,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	accessTokeStr, err := accessToken.SignedString([]byte(accessSecret))
	if err != nil {
		panic(err.Error())
	}
	log.Println("ACCESS TOKEN:", accessTokeStr)

	refreshTokenStr, err1 := refreshToken.SignedString([]byte(refreshSecret))
	if err1 != nil {
		panic(err1.Error())
	}
	userservice.UserRepo.StoreRefreshTokenQuery(refreshTokenStr, email)

	return accessTokeStr, refreshTokenStr, nil
}

func (userservice *UserService) RefreshTokenService(token string) (bool, string) {

	check := userservice.UserRepo.RefreshTokenQuery(token)

	if !check {
		return check, ""
	}
	envPath := os.Getenv("API_KEY")

	log.Println("ENV_PATH", envPath)
	if envPath == "" {
		if errr := godotenv.Load("/home/aamir/Desktop/My_code/searchRecommend/.env"); errr != nil {

			panic(errr.Error())
		}
	}
	secret := os.Getenv("REFRESH_TOKEN_SECRET")

	checkToken, errr := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Println("[User Service] Incorrect Token!")
		}
		return []byte(secret), nil
	})
	if errr != nil {
		panic(errr.Error())
	}

	var accessTokeStr string
	var err1 error
	if checkToken.Valid {

		accessSecret := os.Getenv("ACCESS_TOKEN_SECRET")

		role := checkToken.Claims.(jwt.MapClaims)["role"].(bool)
		userID_FOLAT := checkToken.Claims.(jwt.MapClaims)["user_id"].(float64)
		accessTokenClaims := jwt.MapClaims{
			"exp":     time.Now().Add(5 * time.Minute).Unix(), // Token will expire in 1 minute
			"iat":     time.Now().Unix(),
			"role":    role,
			"user_id": int64(userID_FOLAT),
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

func (userservice *UserService) LogoutService(user_id int64) bool {

	check := userservice.UserRepo.LogoutQuery(user_id)
	return check
}
