package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	service "searchRecommend/services"
	"strconv"
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

	param := r.URL.Query()

	page_no := param.Get("page_no")
	limit := param.Get("limit")

	limitInt, err1 := strconv.Atoi(limit)
	if err1 != nil {
		panic(err1.Error())
	}

	page_noInt, err2 := strconv.Atoi(page_no)
	if err2 != nil {
		panic(err2.Error())
	}

	books, num_pages := handler.Bookservice.GetBooksService(limitInt, page_noInt)

	if page_noInt > int(num_pages) || page_noInt < 1 {
		// log.Fatal(http.StatusBadRequest, ",", http.StatusText(400))
		http.Error(w, "Page number does not exist!", 400.0)
		//log.Fatal(http.StatusBadRequest, ",", http.StatusText(400))

		return
	}

	json, err := json.Marshal(books)
	if err != nil {
		log.Fatal(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(json)
}
