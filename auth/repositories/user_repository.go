package repository

import (
	"fmt"
	"log"
	"reflect"
	"searchRecommend/auth/schema"
	util "searchRecommend/auth/util"
)

type UserRepository struct {
	Db *util.Db
}

func (bookquery *UserRepository) QueryCount() (int, error) {

	db, err := bookquery.Db.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()
	query, errr := db.Query("SELECT COUNT(*) FROM booktwo;")
	if errr != nil {
		panic(errr)
	}

	var count int
	if query.Next() {
		data := query.Scan(&count)
		if data != nil {
			panic(data.Error())
		}
	}

	return count, nil
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
		(userid, lastname, firstname, email, hashpass, salt, admin)
        VALUES 
		('%d', '%s', '%s', '%s', '%s', '%s', '%t');`, user.ID, user.LAST_NAME, user.FIRST_NAME, user.EMAIL, user.HASH_PASSWORD, user.SALT, false)

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

func (userquery *UserRepository) CreateJWTQuery(email string) (bool, int64) {

	db, err := userquery.Db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	getUserRole := fmt.Sprintf("SELECT admin, userid FROM users WHERE email = '%s';", email)
	query, err1 := db.Query(getUserRole)
	if err1 != nil {
		panic(err1.Error())
	}

	var roleAdmin bool
	var user_id int64 // Use an empty interface
	// var user_id int64
	if query.Next() {
		data := query.Scan(
			&roleAdmin,
			&user_id,
		)

		if data != nil {
			panic(data.Error())
		}
	}
	log.Println(reflect.TypeOf(user_id), user_id)
	return roleAdmin, user_id
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

func (userquery *UserRepository) LogoutQuery(user_id int64) bool {

	db, err := userquery.Db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	checkUserID := fmt.Sprintf("SELECT COUNT(*) FROM users WHERE userid = %d;", user_id)
	query, err1 := db.Query(checkUserID)
	if err1 != nil {
		panic(err1.Error())
	}

	var count1 int
	if query.Next() {
		data := query.Scan(&count1)
		if data != nil {
			panic(data.Error())
		}
	}

	if count1 == 0 {
		return false
	}

	deleteToken := fmt.Sprintf("UPDATE users SET refresh_token = '' WHERE userid = %d;", user_id)
	query1, err2 := db.Query(deleteToken)
	if err2 != nil {
		panic(err2.Error())
	}

	if query1.Next() {
		data := query.Scan()
		if data != nil {
			panic(data.Error())
		}
	}
	log.Println(deleteToken)
	return true
}
