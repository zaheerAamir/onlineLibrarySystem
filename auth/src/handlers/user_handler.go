package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"searchRecommend/auth/schema"
	service "searchRecommend/auth/services"
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
		var resp schema.UserSucces
		resp.CODE = 200
		resp.TEXT = "SuccesFully logged In. Please use the access token to execute any bookAPI"

		accessToken, refreshToken, err := handler.UserService.CreateJWT(user.EMAIL)
		if err != nil {
			panic(err.Error())
		}
		resp.ACCES_TOKEN = accessToken
		resp.REFRESH_TOKEN = refreshToken

		w.WriteHeader(200)
		json, err := json.Marshal(resp)
		if err != nil {
			panic(err.Error())
		}
		w.Header().Set("content-type", "application/json")
		w.Write(json)
	}

}

func (handler *UserHandler) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	var refresh_token schema.RefreshTokenDTO

	if err := json.Unmarshal(body, &refresh_token); err != nil {
		panic(err.Error())
	}

	check, token := handler.UserService.RefreshTokenService(refresh_token.REFRESH_TOKEN)

	var error schema.Error

	if !check && token == "" {
		log.Println(check)
		error.CODE = 403
		error.STATUSTEXT = http.StatusText(error.CODE)
		error.MESSAGE = "Access Denied! Refresh Token invalid 🛑"

		json, err := json.Marshal(error)
		if err != nil {
			panic(err.Error())
		}
		w.WriteHeader(error.CODE)
		w.Header().Set("content-type", "application/json")
		w.Write(json)
	} else {
		var access_token schema.AccessTokenSchema
		access_token.AccessToken = token
		w.WriteHeader(200)
		json, err := json.Marshal(access_token)
		if err != nil {
			panic(err.Error())
		}
		w.Header().Set("content-type", "application/json")
		w.Write(json)
	}
}

func (handler *UserHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var email schema.Logout
	if err1 := json.Unmarshal(body, &email); err1 != nil {
		panic(err1.Error())
	}

	check := handler.UserService.LogoutService(email.EMAIL)
	var error schema.Error
	if !check {
		error.CODE = 404
		error.STATUSTEXT = http.StatusText(error.CODE)
		error.MESSAGE = "User Not Found! Email does not exist"

		w.WriteHeader(error.CODE)
		json, err := json.Marshal(error)
		if err != nil {
			panic(err.Error())
		}
		w.Header().Set("content-type", "application/json")
		w.Write(json)
	} else {
		w.WriteHeader(204)
		w.Write([]byte("User Successfully Logged Out!"))
	}
}
