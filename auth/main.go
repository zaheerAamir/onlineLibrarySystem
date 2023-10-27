package main

import (
	"io"
	"log"
	"net/http"
	"os"
	repository "searchRecommend/auth/repositories"
	service "searchRecommend/auth/services"
	route "searchRecommend/auth/src"
	handler "searchRecommend/auth/src/handlers"
	util "searchRecommend/auth/util"
)

// @securityDefinitions.apikey bearerToken
// @in header
// @name Authorization
// @description Enter your access_token in the form of <b>Bearer &lt;access_token&gt;</b>
func main() {

	// Serve Swagger JSON
	http.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		swaggerFile, err := os.Open("docs/swagger.json")
		if err != nil {
			http.Error(w, "Unable to serve Swagger JSON", http.StatusInternalServerError)
			return
		}
		defer swaggerFile.Close()
		io.Copy(w, swaggerFile)
	})

	http.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("docs/swagger-ui"))))
	DB := &util.Db{}

	userRepo := &repository.UserRepository{Db: DB}
	userService := &service.UserService{UserRepo: userRepo}
	userHandler := &handler.UserHandler{UserService: userService}

	route.SetupRoutes(userHandler)

	log.Println("Server Running on Port :8081")
	http.ListenAndServe(":8081", nil)

}
