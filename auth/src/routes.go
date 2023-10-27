package route

import (
	"net/http"
	handler "searchRecommend/auth/src/handlers"
	"searchRecommend/auth/src/middlewares"
)

func SetupRoutes(userhandler *handler.UserHandler) {

	http.HandleFunc("/createUser", middlewares.LoggerMiddleware(middlewares.SetContentType(middlewares.SignUp(userhandler.CreateUserHandler))))
	http.HandleFunc("/login", middlewares.LoggerMiddleware(middlewares.SetContentType(middlewares.Login(userhandler.LoginUserHandler))))

	http.HandleFunc("/token", middlewares.LoggerMiddleware(middlewares.SetContentType(middlewares.CreateToken(userhandler.RefreshTokenHandler))))
	http.HandleFunc("/logout", middlewares.LoggerMiddleware(middlewares.SetContentType(middlewares.AuthorizeUser(userhandler.LogoutHandler))))
}
