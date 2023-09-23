package route

import (
	"net/http"
	handler "searchRecommend/src/handlers"
)

// routes
func Setuproutes(bookhandler *handler.BookHandler) {

	http.HandleFunc("/queryCount", bookhandler.Query)
	http.HandleFunc("/getBooks", bookhandler.GetBooks)
}
