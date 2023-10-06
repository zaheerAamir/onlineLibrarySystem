package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"searchRecommend/schema"
	service "searchRecommend/services"
)

type UserHandler struct {
	UserService *service.UserService
}

func (handler *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	req, errr := io.ReadAll(r.Body)
	if errr != nil {
		panic(errr.Error())
	}
	var user schema.UserDto
	if err := json.Unmarshal(req, &user); err != nil {
		panic(err.Error())
	}

	exists, err1 := handler.UserService.CreateUserService(user)
	if err1 != nil {
		panic(err1.Error())
	}

	var error schema.Error
	error.CODE = 400
	error.STATUSTEXT = http.StatusText(400)
	error.MESSAGE = "User already exists!"

	if exists {
		w.WriteHeader(201)
		w.Write([]byte("User Created"))
	} else {
		json, err := json.Marshal(error)
		if err != nil {
			panic(err.Error())
		}
		w.WriteHeader(error.CODE)
		w.Header().Set("content-type", "application/json")
		w.Write(json)
	}
}

func (handler *UserHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {

	req, errr := io.ReadAll(r.Body)
	if errr != nil {
		panic(errr.Error())
	}
	var user schema.UserLoginDto
	if err := json.Unmarshal(req, &user); err != nil {
		panic(err.Error())
	}

	loggedIN, error := handler.UserService.LoginUserService(user)
	if !loggedIN {
		json, err := json.Marshal(error)
		if err != nil {
			panic(err.Error())
		}
		w.WriteHeader(error.CODE)
		w.Header().Set("content-type", "application/json")
		w.Write(json)
	} else {
		w.WriteHeader(200)
		w.Write([]byte("Welcome to the worlds fastest BookSearch api!"))
	}

}
