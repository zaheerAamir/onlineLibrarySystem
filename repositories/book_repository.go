package repository

import (
	"searchRecommend/schema"
	util "searchRecommend/utils"
)

type BookQuery struct {
	BookDb *util.BookDb
}

func (bookquery *BookQuery) QueryCount() (int, error) {

	db, err := bookquery.BookDb.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()
	query, errr := db.Query("SELECT COUNT(*) FROM bookone;")
	if errr != nil {
		panic(errr)
	}

	var count int
	if query.Next() {
		data := query.Scan(&count)
		if data != nil {
			panic(data.Error())
		}
	}

	return count, nil
}

func (bookquery *BookQuery) GetBooksQuery() ([]schema.Books, error) {

	db, err := bookquery.BookDb.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()
	query, errr := db.Query("SELECT bookone.title, booktwo.authors, bookone.textreviewscount, booktwo.langcode, booktwo.numpages, bookthree.avg_rating, bookthree.publisher, bookthree.publishingdate FROM ((bookone INNER JOIN booktwo on bookone.bookid = booktwo.bookid) INNER JOIN bookthree on bookone.bookid = bookthree.bookid);")
	if errr != nil {
		panic(errr)
	}

	var books []schema.Books

	for query.Next() {

		var book schema.Books
		data := query.Scan(
			&book.TITLE,
			&book.AUTHORS,
			&book.Text_REVIEWS_COUNT,
			&book.LANGUAGE_CODE,
			&book.NUM_PAGES,
			&book.AVERAGE_RATING,
			&book.PUBLISHER,
			&book.PUBLICATION_DATE,
		)

		if data != nil {
			panic(data)
		}

		books = append(books, book)

	}

	return books, nil
}
