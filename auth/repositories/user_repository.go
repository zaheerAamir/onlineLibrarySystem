package repository

import (
	"fmt"
	"log"
	"searchRecommend/auth/schema"
	util "searchRecommend/auth/util"
)

type UserRepository struct {
	Db *util.Db
}

func (userquery *UserRepository) CreateUserQuery(user schema.UserSchema) bool {

	db, err := userquery.Db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	checkEmail := fmt.Sprintf("SELECT COUNT(*) FROM users WHERE email = '%s';", user.EMAIL)
	query, err1 := db.Query(checkEmail)
	if err1 != nil {
		panic(err1.Error())
	}

	var count int
	if query.Next() {
		data := query.Scan(&count)
		if data != nil {
			panic(data.Error())
		}
	}
	if count == 0 {
		insertUser := fmt.Sprintf(`INSERT INTO users 
		(userid, lastname, firstname, email, hashpass, salt)
        VALUES 
		('%d', '%s', '%s', '%s', '%s', '%s');`, user.ID, user.LAST_NAME, user.FIRST_NAME, user.EMAIL, user.HASH_PASSWORD, user.SALT)

		query, err1 := db.Query(insertUser)
		if err1 != nil {
			panic(err1.Error())
		}

		if query.Next() {
			data := query.Scan()
			if data != nil {
				panic(data.Error())
			}
		}

		return true
	}

	return false

}

func (userquery *UserRepository) LoginUserQuery(userLoginCred schema.UserLoginDto) (string, string) {

	db, err := userquery.Db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	checkEmail := fmt.Sprintf("SELECT COUNT(*) FROM users WHERE email = '%s';", userLoginCred.EMAIL)
	query, err1 := db.Query(checkEmail)
	if err1 != nil {
		panic(err1.Error())
	}

	var count int
	if query.Next() {
		data := query.Scan(&count)
		if data != nil {
			panic(data.Error())
		}
	}

	var hasshPass string
	var salt string
	if count != 0 {
		getUserPass := fmt.Sprintf("SELECT hashpass, salt FROM users WHERE email = '%s';", userLoginCred.EMAIL)
		query, err1 := db.Query(getUserPass)
		if err1 != nil {
			panic(err1.Error())
		}

		if query.Next() {
			data := query.Scan(
				&hasshPass,
				&salt,
			)

			if data != nil {
				panic(data.Error())
			}
		}
		return hasshPass, salt
	}

	return hasshPass, salt

}

func (userquery *UserRepository) CreateJWTQuery(email string) bool {

	db, err := userquery.Db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	getUserRole := fmt.Sprintf("SELECT admin FROM users WHERE email = '%s';", email)
	query, err1 := db.Query(getUserRole)
	if err1 != nil {
		panic(err1.Error())
	}

	var roleAdmin bool
	if query.Next() {
		data := query.Scan(
			&roleAdmin,
		)

		if data != nil {
			panic(data.Error())
		}
	}
	return roleAdmin
}

func (userquery *UserRepository) StoreRefreshTokenQuery(refreshToken, email string) {

	db, err := userquery.Db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	log.Println(refreshToken)
	log.Println(email)
	setRefreshToken := fmt.Sprintf("UPDATE users SET refresh_token = '%s' WHERE email = '%s';", refreshToken, email)
	log.Println(setRefreshToken)
	query, err1 := db.Query(setRefreshToken)
	if err1 != nil {
		panic(err1.Error())
	}

	if query.Next() {
		data := query.Scan()
		if data != nil {
			panic(data.Error())
		}
	}
}

func (userquery *UserRepository) RefreshTokenQuery(token string) bool {

	db, err := userquery.Db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	checkRefrehToken := fmt.Sprintf("SELECT COUNT(*) FROM users WHERE refresh_token = '%s';", token)
	query, err1 := db.Query(checkRefrehToken)
	if err1 != nil {
		panic(err1.Error())
	}

	var count int
	if query.Next() {
		data := query.Scan(&count)
		if data != nil {
			panic(data.Error())
		}
	}

	if count == 0 {
		return false
	}
	return true
}

func (userquery *UserRepository) LogoutQuery(email string) bool {

	db, err := userquery.Db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	checkEmail := fmt.Sprintf("SELECT COUNT(*) FROM users WHERE email = '%s';", email)
	query, err1 := db.Query(checkEmail)
	if err1 != nil {
		panic(err1.Error())
	}

	var count int
	if query.Next() {
		data := query.Scan(&count)
		if data != nil {
			panic(data.Error())
		}
	}

	if count != 0 {
		deleteToken := fmt.Sprintf("UPDATE users SET refresh_token = '' WHERE email = '%s';", email)
		query1, err2 := db.Query(deleteToken)
		if err2 != nil {
			panic(err2.Error())
		}

		if query1.Next() {
			data := query.Scan(&count)
			if data != nil {
				panic(data.Error())
			}
		}
		return true
	}
	return false
}
