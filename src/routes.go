package route

import (
	"net/http"
	handler "searchRecommend/src/handlers"
	"searchRecommend/src/middlewares"
)

// routes
func Setuproutes(bookhandler *handler.BookHandler,
	userhandler *handler.UserHandler,
	rentbookhandler *handler.RentBookHandler) {

	bookRoutes(bookhandler)
	userRoutes(userhandler)
	rentBookRoutes(rentbookhandler)
}

func bookRoutes(bookhandler *handler.BookHandler) {
	http.HandleFunc("/queryCount", middlewares.LoggerMiddleware(bookhandler.Query))
	http.HandleFunc("/getBooks", middlewares.LoggerMiddleware(bookhandler.GetBooks))
	http.HandleFunc("/filterBooks", middlewares.LoggerMiddleware(bookhandler.FilterBooks))

}

func userRoutes(userhandler *handler.UserHandler) {

	http.HandleFunc("/createUser", middlewares.LoggerMiddleware(middlewares.SetContentType(middlewares.SignUp(userhandler.CreateUserHandler))))
	http.HandleFunc("/login", middlewares.LoggerMiddleware(middlewares.SetContentType(middlewares.Login(userhandler.LoginUserHandler))))
}

func rentBookRoutes(rentbookhander *handler.RentBookHandler) {

	http.HandleFunc("/rentBook/", middlewares.LoggerMiddleware(middlewares.SetContentType(middlewares.RentBook(rentbookhander.RentbookHandler))))
	http.HandleFunc("/giveBookBack", middlewares.LoggerMiddleware(middlewares.SetContentType(middlewares.GiveBack(rentbookhander.GiveBookBackHandler))))
}
