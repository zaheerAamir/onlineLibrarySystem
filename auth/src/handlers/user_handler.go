package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"searchRecommend/auth/schema"
	service "searchRecommend/auth/services"
	"searchRecommend/auth/src/middlewares"
)

type UserHandler struct {
	UserService *service.UserService
}

// @Summary SignUp user route
// @Description User can create Account
// @Tags users
// @Accept json
// @Produce json
// @Param body body schema.UserDto true "Request body in JSON format"
// @Success 201 {object} schema.UserSuccess
// @Failure 409 {object}  schema.Error
// @Router /createUser [post]
func (handler *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
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
	error.CODE = 409
	error.STATUSTEXT = http.StatusText(error.CODE)
	error.MESSAGE = "User already exists!"

	if exists {
		var success schema.UserSuccess
		success.CODE = 201
		success.TEXT = http.StatusText(success.CODE)
		success.MESSAGE = "User created Successfully!"

		json, err := json.Marshal(success)
		if err != nil {
			panic(err.Error())
		}

		w.WriteHeader(success.CODE)
		w.Write(json)
	} else {
		json, err := json.Marshal(error)
		if err != nil {
			panic(err.Error())
		}
		w.WriteHeader(error.CODE)
		w.Write(json)
	}
}

// @Summary Login user route
// @Description User can Login to their Account
// @Tags users
// @Accept json
// @Produce json
// @Param body body schema.UserLoginDto true "Request body in JSON format"
// @Success 200 {object} schema.UserLoginSuccess
// @Failure 401 {object}  schema.Error
// @Router /login [post]
func (handler *UserHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
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
		w.Write(json)
	} else {
		var resp schema.UserLoginSuccess
		resp.CODE = 200
		resp.TEXT = "SuccesFully logged In. Please use the access token to execute any bookAPI"

		accessToken, refreshToken, err := handler.UserService.CreateJWT(user.EMAIL)
		if err != nil {
			panic(err.Error())
		}
		resp.ACCES_TOKEN = accessToken
		resp.REFRESH_TOKEN = refreshToken

		json, err := json.Marshal(resp)
		if err != nil {
			panic(err.Error())
		}
		w.WriteHeader(200)
		w.Write(json)
	}

}

// @Summary Refresh Access token route
// @Description User can Refresh the expired Access token from the Refresh token
// @Tags users
// @Accept json
// @Produce json
// @Param body body schema.RefreshTokenDTO true "Request body in JSON format"
// @Success 200 {object} schema.AccessTokenSchema
// @Failure 403 {object} schema.Error
// @Router /token [post]
func (handler *UserHandler) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
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
		log.Println("[User Handler] [Refresh Token Handler] Valid Refresh token:", check)
		error.CODE = 403
		error.STATUSTEXT = http.StatusText(error.CODE)
		error.MESSAGE = "Access Denied! Refresh Token invalid 🛑"

		json, err := json.Marshal(error)
		if err != nil {
			panic(err.Error())
		}
		w.WriteHeader(error.CODE)
		w.Write(json)
	} else {
		var access_token schema.AccessTokenSchema
		access_token.AccessToken = token
		json, err := json.Marshal(access_token)
		if err != nil {
			panic(err.Error())
		}
		w.WriteHeader(200)
		w.Write(json)
	}
}

// @Summary Logout user route
// @Description User can Logout of their Account
// @Tags users
// @Produce json
// @Security bearerToken
// @Success 204
// @Failure 404 {object}  schema.Error
// @Router /logout   [delete]
func (handler *UserHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	var error schema.Error
	error.CODE = 400
	error.STATUSTEXT = http.StatusText(400)
	error.MESSAGE = "Bad request it is a DELETE route"

	if r.Method != http.MethodDelete {
		json, err := json.Marshal(error)
		if err != nil {
			panic(err.Error())
		}
		w.WriteHeader(error.CODE)
		w.Write(json)

	}

	userID, ok := r.Context().Value(middlewares.UserIDKey).(int64)
	log.Println(userID)

	if !ok {
		panic("UserID not found or has unexpected type")
	}

	check := handler.UserService.LogoutService(userID)
	if !check {
		error.CODE = 404
		error.STATUSTEXT = http.StatusText(error.CODE)
		error.MESSAGE = "User Not Found! User does not exist"

		json, err := json.Marshal(error)
		if err != nil {
			panic(err.Error())
		}
		w.WriteHeader(error.CODE)
		w.Write(json)
	} else {

		w.WriteHeader(204)
	}
}
