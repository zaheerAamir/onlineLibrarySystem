package repository

import (
	"database/sql"
	"fmt"
	"log"
	"searchRecommend/schema"
	util "searchRecommend/utils"
)

type BookQuery struct {
	Db *util.Db
}

func (bookquery *BookQuery) QueryCount() (int, error) {

	db, err := bookquery.Db.ConnectDB()
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

	db, err := bookquery.Db.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()
	sqlStatement := fmt.Sprintf(`
	SELECT bookone.title, 
	booktwo.authors, 
	booktwo.lang_code, 
	booktwo.num_pages, 
	bookthree.avg_rating, 
	bookthree.publisher, 
	bookthree.publishing_date FROM ((bookone INNER JOIN booktwo on bookone.bookid = booktwo.bookid) 
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

func makeQueryHelper(queryStr string, db *sql.DB, books []schema.Books) ([]schema.Books, error) {

	log.Println("executing")
	query, err1 := db.Query(queryStr)
	if err1 != nil {
		panic(err1.Error())
	}

	for query.Next() {
		var book schema.Books
		data := query.Scan(
			&book.TITLE,
			&book.AUTHORS,
			&book.LANGUAGE_CODE,
			&book.NUM_PAGES,
			&book.AVERAGE_RATING,
			&book.PUBLISHER,
			&book.PUBLICATION_DATE,
		)
		if data != nil {
			panic(data.Error())
		}

		books = append(books, book)

	}

	if err := query.Err(); err != nil {
		return nil, err
	}

	return books, nil

}

func (bookquery *BookQuery) FilterBooksQuery(author, publisher, avg_rating, num_pages string) ([]schema.Books, error) {

	db, err := bookquery.Db.ConnectDB()
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	sqlStatement := `
	SELECT 
    bookone.title, 
    booktwo.authors, 
    booktwo.lang_code, 
    booktwo.num_pages, 
    bookthree.avg_rating, 
    bookthree.publisher, 
    bookthree.publishing_date 
    FROM 
        bookone 
    INNER JOIN 
        booktwo ON bookone.bookid = booktwo.bookid
    INNER JOIN 
        bookthree ON bookone.bookid = bookthree.bookid
    `
	var books []schema.Books

	if author != "" && publisher != "" && avg_rating != "" {

		joinAuthorANDPublisherAvg := fmt.Sprintf(`
		%s 
		WHERE 
		    booktwo.authors LIKE '%%%s%%' 
		AND 
		    bookthree.publisher = '%s' ORDER BY bookthree.avg_rating %s`,
			sqlStatement, author, publisher, avg_rating)

		log.Println(joinAuthorANDPublisherAvg)

		books, err = makeQueryHelper(joinAuthorANDPublisherAvg, db, books)
		if err != nil {
			panic(err.Error())
		}

	} else if author != "" && publisher != "" && num_pages != "" {

		joinAuthorANDPublisherNum := fmt.Sprintf(`
		%s 
		WHERE 
		    booktwo.authors LIKE '%%%s%%' 
		AND 
		    bookthree.publisher = '%s' ORDER BY booktwo.numpages %s`,
			sqlStatement, author, publisher, num_pages)

		log.Println(joinAuthorANDPublisherNum)

		books, err = makeQueryHelper(joinAuthorANDPublisherNum, db, books)
		if err != nil {
			panic(err.Error())
		}

	} else if author != "" && publisher != "" {
		joinAuthorANDPublisher := fmt.Sprintf(`%s WHERE booktwo.authors LIKE '%%%s%%' AND bookthree.publisher = '%s'`, sqlStatement, author, publisher)
		log.Println(joinAuthorANDPublisher)

		books, err = makeQueryHelper(joinAuthorANDPublisher, db, books)
		if err != nil {
			panic(err.Error())
		}

	} else if author != "" && publisher == "" && avg_rating != "" {

		joinAuthorAvg := fmt.Sprintf(`%s WHERE booktwo.authors LIKE '%%%s%%' ORDER BY bookthree.avg_rating %s`, sqlStatement, author, avg_rating)

		log.Println(joinAuthorAvg)

		books, err = makeQueryHelper(joinAuthorAvg, db, books)
		if err != nil {
			panic(err.Error())
		}

	} else if author != "" && publisher == "" && num_pages != "" {

		joinAuthorNum := fmt.Sprintf(`%s WHERE booktwo.authors LIKE '%%%s%%' ORDER BY booktwo.numpages %s`, sqlStatement, author, num_pages)

		log.Println(joinAuthorNum)

		books, err = makeQueryHelper(joinAuthorNum, db, books)
		if err != nil {
			panic(err.Error())
		}

	} else if author != "" && publisher == "" {

		joinAuthor := fmt.Sprintf(`%s WHERE booktwo.authors LIKE '%%%s%%' `, sqlStatement, author)

		log.Println(joinAuthor)

		books, err = makeQueryHelper(joinAuthor, db, books)
		if err != nil {
			panic(err.Error())
		}

	} else if publisher != "" && author == "" && avg_rating != "" {

		joinPublisherAvg := fmt.Sprintf(`%s WHERE bookthree.publisher = '%s' ORDER BY bookthree.avg_rating %s`, sqlStatement, publisher, avg_rating)

		log.Println(joinPublisherAvg)

		books, err = makeQueryHelper(joinPublisherAvg, db, books)
		if err != nil {
			panic(err.Error())
		}

	} else if publisher != "" && author == "" && num_pages != "" {

		joinPublisherNum := fmt.Sprintf(`%s WHERE bookthree.publisher = '%s' ORDER BY booktwo.numpages %s`, sqlStatement, publisher, num_pages)

		log.Println(joinPublisherNum)

		books, err = makeQueryHelper(joinPublisherNum, db, books)
		if err != nil {
			panic(err.Error())
		}

	} else if publisher != "" && author == "" {

		joinPublisher := fmt.Sprintf(`%s WHERE bookthree.publisher = '%s'`, sqlStatement, publisher)

		log.Println(joinPublisher)

		books, err = makeQueryHelper(joinPublisher, db, books)
		if err != nil {
			panic(err.Error())
		}

	} else {
		log.Fatal("Atleast one parameter must be set!")
	}

	return books, nil

}
