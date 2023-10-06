package main

import (
	"log"
	"net/http"
	repository "searchRecommend/repositories"
	service "searchRecommend/services"
	route "searchRecommend/src"
	handler "searchRecommend/src/handlers"
	util "searchRecommend/utils"
)

func main() {

	DB := &util.Db{}

	bookRepo := &repository.BookQuery{Db: DB}
	bookService := &service.BookService{BookRepo: bookRepo}
	bookHandler := &handler.BookHandler{Bookservice: bookService}

	userRepo := &repository.UserRepository{Db: DB}
	userService := &service.UserService{UserRepo: userRepo}
	userHandler := &handler.UserHandler{UserService: userService}

	rentBookRepo := &repository.RentBookRepo{Db: DB}
	rentBookService := &service.RentBookService{RentBookRepo: rentBookRepo}
	rentBookHandler := &handler.RentBookHandler{Bookservice: rentBookService}

	foo := "Hello"
	point := &foo
	val := *point
	log.Println(foo, point, val)

	route.Setuproutes(bookHandler, userHandler, rentBookHandler)

	//Creating The DB
	// sql, err := DB.ConnectDB()
	// if err != nil {
	// 	panic(err)
	// }
	// util.Task(sql)

	log.Println("Server Running on Port :8080")
	http.ListenAndServe(":8080", nil)

	// Multi()
}
