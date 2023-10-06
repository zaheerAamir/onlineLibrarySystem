package repository

import (
	"fmt"
	"searchRecommend/schema"
	util "searchRecommend/utils"
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
