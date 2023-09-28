package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"searchRecommend/schema"
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

	param := r.URL.Query()

	page_no := param.Get("page_no")
	limit := param.Get("limit")

	var error schema.Error
	error.CODE = 400
	error.STATUSTEXT = http.StatusText(error.CODE)
	error.MESSAGE = "Parameters page number and limit should not be empty"

	if page_no == "" && limit == "" {
		json, err := json.Marshal(error)
		if err != nil {
			log.Fatal(err.Error())
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(error.CODE)
		w.Write(json)
	} else if page_no == "" || limit == "" {
		json, err := json.Marshal(error)
		if err != nil {
			log.Fatal(err.Error())
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(error.CODE)
		w.Write(json)
	} else {

		books, num_pages, page_no := handler.Bookservice.GetBooksService(limit, page_no)

		error.CODE = 400
		error.STATUSTEXT = http.StatusText(error.CODE)
		error.MESSAGE = "Page number does exist"

		if page_no > int(num_pages) || page_no < 1 {
			json, err := json.Marshal(error)
			if err != nil {
				log.Fatal(err.Error())
			}
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(error.CODE)
			w.Write(json)

		} else {
			json, err := json.Marshal(books)
			if err != nil {
				log.Fatal(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("content-type", "application/json")
			w.Write(json)
		}

	}

}

func (handler *BookHandler) FilterBooks(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()

	author := param.Get("author")
	publisher := param.Get("publisher")

	var query []string

	var error schema.Error
	error.CODE = 400
	error.STATUSTEXT = http.StatusText(error.CODE)
	error.MESSAGE = "Parameters author and publisher name should not be empty"

	if author == "" && publisher == "" {
		json, err := json.Marshal(error)
		if err != nil {
			log.Fatal(err.Error())
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(error.CODE)
		w.Write(json)
	} else if publisher == "" {
		log.Println(author)

		for i := 0; i < 2; i++ {
			if i == 0 {
				query = append(query, author)
			} else {
				query = append(query, "")
			}
		}
		books := handler.Bookservice.FilterBooksService(query)

		json, err := json.Marshal(books)
		if err != nil {
			log.Fatal(err.Error())
		}

		w.Header().Set("content-type", "application/json")
		w.Write(json)

	} else if author == "" {
		log.Println(publisher)

		for i := 0; i < 2; i++ {
			if i == 0 {
				query = append(query, "")
			} else {
				query = append(query, publisher)
			}
		}
		books := handler.Bookservice.FilterBooksService(query)

		json, err := json.Marshal(books)
		if err != nil {
			log.Fatal(err.Error())
		}

		w.Header().Set("content-type", "application/json")
		w.Write(json)

	} else {
		log.Println(author, publisher)

		for i := 0; i < 2; i++ {
			if i == 0 {
				query = append(query, author)
			} else {
				query = append(query, publisher)
			}
		}
		books := handler.Bookservice.FilterBooksService(query)

		json, err := json.Marshal(books)
		if err != nil {
			log.Fatal(err.Error())
		}

		w.Header().Set("content-type", "application/json")
		w.Write(json)
	}

}
