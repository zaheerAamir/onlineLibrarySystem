package route

import (
	"net/http"
	handler "searchRecommend/books/src/handlers"
	"searchRecommend/books/src/middlewares"
)

// routes
func Setuproutes(bookhandler *handler.BookHandler,
	rentbookhandler *handler.RentBookHandler) {

	bookRoutes(bookhandler)
	rentBookRoutes(rentbookhandler)
}

func bookRoutes(bookhandler *handler.BookHandler) {
	http.HandleFunc("/queryCount", middlewares.LoggerMiddleware(middlewares.AuthorizeUser(bookhandler.Query)))
	http.HandleFunc("/getBooks", middlewares.LoggerMiddleware(middlewares.AuthorizeUser(bookhandler.GetBooks)))
	http.HandleFunc("/filterBooks", middlewares.LoggerMiddleware(middlewares.AuthorizeUser(bookhandler.FilterBooks)))

}

func rentBookRoutes(rentbookhander *handler.RentBookHandler) {

	http.HandleFunc("/rentBook/", middlewares.LoggerMiddleware(middlewares.SetContentType(middlewares.AuthorizeUser(middlewares.RentBook(rentbookhander.RentbookHandler)))))
	http.HandleFunc("/giveBookBack", middlewares.LoggerMiddleware(middlewares.SetContentType(middlewares.AuthorizeAdmin(middlewares.GiveBack(rentbookhander.GiveBookBackHandler)))))
}
