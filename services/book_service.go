package service

import (
	"log"
	"math"
	repository "searchRecommend/repositories"
	"searchRecommend/schema"
	"strconv"
	"time"
)

type BookService struct {
	BookRepo *repository.BookQuery
}

// main logic
func (bookService *BookService) DbService() int {

	count, err := bookService.BookRepo.QueryCount()
	if err != nil {
		panic(err)
	}

	return count
}

func (bookservice *BookService) GetBooksService(limit, page_no string) ([]schema.Books, float64, int) {

	tick := time.Now()

	limitInt, err1 := strconv.Atoi(limit)
	if err1 != nil {
		panic(err1.Error())
	}

	page_noInt, err2 := strconv.Atoi(page_no)
	if err2 != nil {
		panic(err2.Error())
	}

	offset := (page_noInt - 1) * limitInt
	log.Println("offset:", offset, "pageNumber:", page_noInt, "limit:", limitInt)

	books, dbLength, err := bookservice.BookRepo.GetBooksQuery(offset, limitInt, page_noInt)
	if err != nil {
		panic(err)
	}

	num_pages := math.Ceil(float64(dbLength) / float64(limitInt))
	log.Println("max num_pages:", num_pages)
	log.Println(time.Since(tick))

	return books, num_pages, page_noInt
}

func (bookservice *BookService) FilterBooksService(query []string) []schema.Books {

	author := query[0]
	publisher := query[1]
	avg_rating := query[2]
	num_pages := query[3]

	books, err := bookservice.BookRepo.FilterBooksQuery(author, publisher, avg_rating, num_pages)
	if err != nil {
		panic(err.Error())
	}

	return books

}
