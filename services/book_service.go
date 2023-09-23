package service

import (
	repository "searchRecommend/repositories"
	"searchRecommend/schema"
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

func (bookservice *BookService) GetBooksService() []schema.Books {

	books, err := bookservice.BookRepo.GetBooksQuery()
	if err != nil {
		panic(err)
	}

	return books
}
