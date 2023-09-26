package repository

import (
	"fmt"
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

func (bookquery *BookQuery) GetBooksQuery(offset, limit, page_no int) ([]schema.Books, int, error) {

	db, err := bookquery.BookDb.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()
	sqlStatement := fmt.Sprintf(`
	SELECT bookone.title, 
	booktwo.authors, 
	bookone.textreviewscount, 
	booktwo.langcode, 
	booktwo.numpages, 
	bookthree.avg_rating, 
	bookthree.publisher, 
	bookthree.publishingdate FROM ((bookone INNER JOIN booktwo on bookone.bookid = booktwo.bookid) 
	INNER JOIN bookthree on bookone.bookid = bookthree.bookid)
	OFFSET %d
	LIMIT %d;`, offset, limit)

	query, errr := db.Query(sqlStatement)
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

	queryCount, err1 := db.Query("SELECT COUNT(*) FROM bookone;")
	if err1 != nil {
		panic(err1.Error())
	}
	var dbLength int

	if queryCount.Next() {
		data := queryCount.Scan(&dbLength)
		if data != nil {
			panic(data.Error())
		}
	}

	return books, dbLength, nil
}
