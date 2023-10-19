package main

import (
	"log"
	"net/http"
	repository "searchRecommend/auth/repositories"
	service "searchRecommend/auth/services"
	route "searchRecommend/auth/src"
	handler "searchRecommend/auth/src/handlers"
	util "searchRecommend/auth/util"
)

func main() {

	DB := &util.Db{}

	userRepo := &repository.UserRepository{Db: DB}
	userService := &service.UserService{UserRepo: userRepo}
	userHandler := &handler.UserHandler{UserService: userService}

	route.SetupRoutes(userHandler)

	log.Println("Server Running on Port :8081")
	http.ListenAndServe(":8081", nil)

}
