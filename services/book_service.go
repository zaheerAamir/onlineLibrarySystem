package service

import (
	"log"
	"math"
	repository "searchRecommend/repositories"
	"searchRecommend/schema"
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

func (bookservice *BookService) GetBooksService(limit, page_no int) ([]schema.Books, float64) {

	tick := time.Now()

	offset := (page_no - 1) * limit
	log.Println(offset, page_no, limit)

	books, dbLength, err := bookservice.BookRepo.GetBooksQuery(offset, limit, page_no)
	if err != nil {
		panic(err)
	}

	num_pages := math.Ceil(float64(dbLength) / float64(limit))
	log.Println(num_pages)
	log.Println(time.Since(tick))

	return books, num_pages
}
