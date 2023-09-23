package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	service "searchRecommend/services"
)

type BookHandler struct {
	Bookservice *service.BookService
}

//Handlers
//healthCheck handler

func (handler *BookHandler) Query(w http.ResponseWriter, r *http.Request) {
	count := handler.Bookservice.DbService()
	dum := fmt.Sprintf("Helllo: %d", count)

	json, err := json.Marshal(dum)

	if err != nil {
		log.Fatal(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(json)
}

func (handler *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {

	books := handler.Bookservice.GetBooksService()

	json, err := json.Marshal(books)
	if err != nil {
		log.Fatal(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(json)
}
