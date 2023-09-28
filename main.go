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

	bookdata := &util.BookDb{}
	bookrepo := &repository.BookQuery{BookDb: bookdata}
	bookservice := &service.BookService{BookRepo: bookrepo}
	bookhandler := &handler.BookHandler{Bookservice: bookservice}

	foo := "Hello"
	point := &foo
	val := *point
	log.Println(foo, point, val)

	route.Setuproutes(bookhandler)
	log.Println("Server Running on Port :8080")
	http.ListenAndServe(":8080", nil)

	// Multi()
}
