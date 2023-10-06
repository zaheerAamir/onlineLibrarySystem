package repository

import (
	"fmt"
	"log"
	"net/http"
	"searchRecommend/schema"
	util "searchRecommend/utils"
)

type RentBookRepo struct {
	Db *util.Db
}

func (rentbookquery *RentBookRepo) RentbookRepo(rent schema.RentBookSchema, bookId int) schema.Error {

	db, err := rentbookquery.Db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	checkUser := fmt.Sprintf("SELECT COUNT(*) FROM users WHERE userid = '%d';", rent.USER_ID)
	query, err1 := db.Query(checkUser)
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

	var error schema.Error
	if count != 0 {

		checkBook_id := fmt.Sprintf("SELECT COUNT(*) FROM bookone WHERE bookid = '%d';", bookId)
		query, err1 := db.Query(checkBook_id)
		if err1 != nil {
			panic(err1.Error())
		}
		log.Println(checkBook_id)

		var count1 int
		if query.Next() {
			data := query.Scan(&count1)
			if data != nil {
				panic(data.Error())
			}
		}

		if count1 != 0 {
			insertRent := fmt.Sprintf(`UPDATE bookone 
            SET rented = true, 
            userid = %d, 
            rented_date = CURRENT_TIMESTAMP, 
            rentdate = TO_DATE('%d-%d-%d', 'DD-MM-YYYY') 
            WHERE bookid = %d;
            `, rent.USER_ID, rent.RENT_DAY.DAYS, rent.RENT_DAY.MONTHS, rent.RENT_DAY.YEARS, bookId)
			log.Println(insertRent)

			query, err1 := db.Query(insertRent)
			if err1 != nil {
				panic(err1.Error())
			}

			if query.Next() {
				data := query.Scan()
				if data != nil {
					panic(data.Error())
				}
			}

			return error
		} else {
			error.CODE = 404
			error.STATUSTEXT = http.StatusText(error.CODE)
			error.MESSAGE = "BookId does not exist!"

			fmt.Println("BookId does not exists")
			return error
		}

	}
	error.CODE = 401
	error.STATUSTEXT = http.StatusText(error.CODE)
	error.MESSAGE = "Unauthorized user does not exist!"

	log.Println("Error User does not exist")
	return error

}

func (rentbookquery *RentBookRepo) GiveBookBackRepo(req schema.GiveBookBackDTO) schema.Error {

	db, err := rentbookquery.Db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	checkEmail := fmt.Sprintf("SELECT userid FROM users WHERE email = '%s'", req.EMAIL)
	query, err1 := db.Query(checkEmail)
	if err1 != nil {
		panic(err1.Error())
	}

	var user_id int64
	if query.Next() {
		data := query.Scan(&user_id)
		if data != nil {
			panic(data.Error())
		}
	}
	log.Println(user_id)
	var error schema.Error
	if user_id != 0 {

		checkBook_id := fmt.Sprintf("SELECT bookid FROM bookone WHERE userid = %d;", user_id)
		query3, errr := db.Query(checkBook_id)
		if errr != nil {
			panic(errr.Error())
		}

		var countArr []int
		for query3.Next() {
			var count int
			data := query3.Scan(&count)
			if data != nil {
				panic(data.Error())
			}
			countArr = append(countArr, count)
		}
		log.Println(countArr)

		for _, v := range countArr {
			if req.BOOK_ID != v {
				error.CODE = 404
				error.STATUSTEXT = http.StatusText(error.CODE)
				error.MESSAGE = "Book not rented!"
				return error

			} else {
				getUserRentedBooks := fmt.Sprintf(`UPDATE bookone
		        SET
		            rented = false,
		            userid = NULL,
		            rented_date = NULL,
		            rentdate = NULL
		        WHERE
		            userid = %d
		        AND
		            bookid = %d;
		        `, user_id, req.BOOK_ID)

				query1, err2 := db.Query(getUserRentedBooks)
				if err2 != nil {
					panic(err2.Error())
				}

				if query1.Next() {
					data := query1.Scan()
					if data != nil {
						panic(data.Error())
					}
				}
				break
			}

		}

		return error

	}
	error.CODE = 404
	error.STATUSTEXT = http.StatusText(error.CODE)
	error.MESSAGE = "UserID not found!"

	log.Println("UserId not found")
	return error

}
