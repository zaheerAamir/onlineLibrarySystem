package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"searchRecommend/books/schema"
	service "searchRecommend/books/services"
)

type BookHandler struct {
	Bookservice *service.BookService
}

//Handlers
//healthCheck handler

func (handler *BookHandler) Query(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	healthCheck := handler.Bookservice.DbService()
	var dum string
	if healthCheck {

		dum = fmt.Sprintf("Helllo: %s", "Bookservice is healthy")

	}

	json, err := json.Marshal(dum)

	if err != nil {
		log.Fatal(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

// @Summary GetBooks user route
// @Description User can get the list of books with pagination
// @Tags books
// @Produce json
// @Security bearerToken
// @Param page_no query int true "Page number (default 1)"
// @Param limit query int true "Number of itmes per page (default 10)"
// @Success 200 {object} schema.Books
// @Failure 400 {object}  schema.Error
// @Router /getBooks  [get]
func (handler *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
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
		w.WriteHeader(error.CODE)
		w.Write(json)
	} else if page_no == "0" {
		error.MESSAGE = "Invalid Page number! Page number starts from 1."
		json, err := json.Marshal(error)
		if err != nil {
			log.Fatal(err.Error())
		}
		w.WriteHeader(error.CODE)
		w.Write(json)

	} else {

		books, num_pages, page_no := handler.Bookservice.GetBooksService(limit, page_no)

		error.CODE = 400
		error.STATUSTEXT = http.StatusText(error.CODE)
		error.MESSAGE = "Page number does exist"

		if page_no > int(num_pages) {
			errMessg := fmt.Sprintf("Incorrect page number value! Page nuber should not be greater that %d for limit %s", int(num_pages), limit)
			error.MESSAGE = errMessg
			json, err := json.Marshal(error)
			if err != nil {
				log.Fatal(err.Error())
			}
			w.WriteHeader(error.CODE)
			w.Write(json)

		} else {
			json, err := json.Marshal(books)
			if err != nil {
				log.Fatal(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(json)
		}

	}

}

func (handler *BookHandler) convJsonHelper(query []string, w http.ResponseWriter) {

	books := handler.Bookservice.FilterBooksService(query)

	json, err := json.Marshal(books)
	if err != nil {
		log.Fatal(err.Error())
	}

	w.Write(json)
}

// @Summary FilterBooks user route
// @Description User can filter books author or publisher name also sort books by number of pages or average rating
// @Tags books
// @Produce json
// @Security bearerToken
// @Param author query string false "Author name to filter by"
// @Param publisher query string false "Publisher name to filter by"
// @Param avg_rating query string false "Sort by average rating ASC or DESC"
// @Param num_pages query string false "Sort by number of pages ASC or DESC"
// @Success 200 {object} schema.Books
// @Failure 400 {object}  schema.Error
// @Router /filterBooks  [get]
func (handler *BookHandler) FilterBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	param := r.URL.Query()

	checkAvg_rating := param.Has("avg_rating")
	checkNum_pages := param.Has("num_pages")

	author := param.Get("author")
	publisher := param.Get("publisher")

	var query []string

	var error schema.Error
	error.CODE = 400
	error.STATUSTEXT = http.StatusText(error.CODE)
	error.MESSAGE = "Parameters author and publisher name should not be empty"

	if checkAvg_rating && checkNum_pages {
		error.MESSAGE = "Cannot set both filters avg_rating and num_pages on a single param"
		json, err := json.Marshal(error)
		if err != nil {
			log.Fatal(err.Error())
		}
		w.WriteHeader(error.CODE)
		w.Write(json)

	} else if author == "" && publisher == "" {
		json, err := json.Marshal(error)
		if err != nil {
			log.Fatal(err.Error())
		}
		w.WriteHeader(error.CODE)
		w.Write(json)

	} else if publisher == "" && checkAvg_rating {

		avg_rating := param.Get("avg_rating")

		for i := 0; i < 4; i++ {

			if i == 0 {
				query = append(query, author)
			} else if i == 1 {
				query = append(query, "")
			} else if i == 2 {
				query = append(query, avg_rating)
			} else {
				query = append(query, "")
			}

		}

		handler.convJsonHelper(query, w)

	} else if publisher == "" && checkNum_pages {

		num_pages := param.Get("num_pages")

		for i := 0; i < 4; i++ {

			if i == 0 {
				query = append(query, author)
			} else if i == 1 {
				query = append(query, "")
			} else if i == 2 {
				query = append(query, "")
			} else {
				query = append(query, num_pages)
			}

		}

		handler.convJsonHelper(query, w)

	} else if publisher == "" {

		log.Println(author)

		for i := 0; i < 4; i++ {
			if i == 0 {
				query = append(query, author)
			} else if i == 1 {
				query = append(query, "")
			} else if i == 2 {
				query = append(query, "")
			} else {
				query = append(query, "")
			}
		}

		handler.convJsonHelper(query, w)

	} else if author == "" && checkAvg_rating {

		avg_rating := param.Get("avg_rating")

		log.Println(publisher)

		for i := 0; i < 4; i++ {
			if i == 0 {
				query = append(query, "")
			} else if i == 1 {
				query = append(query, publisher)
			} else if i == 2 {
				query = append(query, avg_rating)
			} else {
				query = append(query, "")
			}
		}

		handler.convJsonHelper(query, w)

	} else if author == "" && checkNum_pages {

		num_pages := param.Get("num_pages")

		for i := 0; i < 4; i++ {
			if i == 0 {
				query = append(query, "")
			} else if i == 1 {
				query = append(query, publisher)
			} else if i == 2 {
				query = append(query, "")
			} else {
				query = append(query, num_pages)
			}
		}

		handler.convJsonHelper(query, w)

	} else if author == "" {

		log.Println(publisher)

		for i := 0; i < 4; i++ {
			if i == 0 {
				query = append(query, "")
			} else if i == 1 {
				query = append(query, publisher)
			} else if i == 2 {
				query = append(query, "")
			} else {
				query = append(query, "")
			}
		}

		handler.convJsonHelper(query, w)

	} else if author != "" && publisher == "" && checkNum_pages {

		num_pages := param.Get("num_pages")

		for i := 0; i < 4; i++ {
			if i == 0 {
				query = append(query, author)
			} else if i == 1 {
				query = append(query, publisher)
			} else if i == 2 {
				query = append(query, "")
			} else {
				query = append(query, num_pages)
			}
		}

		handler.convJsonHelper(query, w)

	} else if author != "" && publisher != "" && checkAvg_rating {

		avg_rating := param.Get("avg_rating")
		log.Println(author, publisher)

		for i := 0; i < 4; i++ {
			if i == 0 {
				query = append(query, author)
			} else if i == 1 {
				query = append(query, publisher)
			} else if i == 2 {
				query = append(query, avg_rating)
			} else {
				query = append(query, "")
			}
		}

		handler.convJsonHelper(query, w)

	} else if author != "" && publisher != "" && checkNum_pages {

		num_pages := param.Get("num_pages")
		log.Println(author, publisher)

		for i := 0; i < 4; i++ {
			if i == 0 {
				query = append(query, author)
			} else if i == 1 {
				query = append(query, publisher)
			} else if i == 2 {
				query = append(query, "")
			} else {
				query = append(query, num_pages)
			}
		}

		handler.convJsonHelper(query, w)
	} else {
		log.Println(author, publisher)

		for i := 0; i < 4; i++ {
			if i == 0 {
				query = append(query, author)
			} else if i == 1 {
				query = append(query, publisher)
			} else if i == 2 {
				query = append(query, "")
			} else {
				query = append(query, "")
			}
		}

		handler.convJsonHelper(query, w)
	}

}
