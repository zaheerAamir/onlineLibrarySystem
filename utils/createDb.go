package util

func task( /*Db *sql.DB, BooksStore []books*/ ) {

	// tick := time.Now()

	// for _, v := range BooksStore {

	// 	// insert values into bookOne table:
	// 	// sqlStatement := "INSERT INTO BookOne (bookid, title, rented, textreviewscount) VALUES ($1, $2, $3, $4)"
	// 	// query , err := Db.Query(sqlStatement, v.BOOK_ID, v.TITLE, false, v.Text_REVIEWS_COUNT)

	// 	// insert values into bookTwo table:

	// 	// var Exists bool
	// 	// err := Db.QueryRow("SELECT EXISTS(SELECT 1 FROM bookone WHERE bookid = $1)", v.BOOK_ID).Scan(&Exists)
	// 	// if err != nil {
	// 	// 	panic(err)
	// 	// }

	// 	// if !Exists {
	// 	// 	log.Printf("BookId %d not exists in bookOne table\n", v.BOOK_ID)
	// 	// } else {
	// 	// 	sqlStatement := "INSERT INTO BookTwo (authors, langcode, numpages, bookid) VALUES ($1, $2, $3, $4)"

	// 	// 	query, err := Db.Query(sqlStatement, v.AUTHORS, v.LANGUAGE_CODE, v.NUM_PAGES, v.BOOK_ID)

	// 	// 	if err != nil {
	// 	// 		panic(err.Error())
	// 	// 	}

	// 	// 	if query.Next() {
	// 	// 		data := query.Scan()

	// 	// 		if data != nil {
	// 	// 			panic(data)
	// 	// 		}
	// 	// 	}

	// 	// }

	// 	//insert values into bookThree table:

	// 	var Exists bool
	// 	err := Db.QueryRow("SELECT EXISTS(SELECT 1 FROM bookone WHERE bookid = $1)", v.BOOK_ID).Scan(&Exists)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if !Exists {
	// 		log.Printf("BookId %d not exists in bookOne table\n", v.BOOK_ID)
	// 	} else {
	// 		sqlStatement := "INSERT INTO BookThree (publisher, publishingdate, avg_rating, bookid) VALUES ($1, $2, $3, $4)"

	// 		query, err := Db.Query(sqlStatement, v.PUBLISHER, v.PUBLICATION_DATE, v.AVERAGE_RATING, v.BOOK_ID)

	// 		if err != nil {
	// 			panic(err.Error())
	// 		}

	// 		if query.Next() {
	// 			data := query.Scan()

	// 			if data != nil {
	// 				panic(data)
	// 			}
	// 		}

	// 	}

	//}

	//log.Printf("Time took: %s\n", time.Since(tick))

}

// func (bookService *BookService) Service() ([]books, error) {

// 	file, err := os.Open("books.tsv")
// 	if err != nil {
// 		log.Panic("Error cannot open books.tsv: ", err)
// 	}
// 	defer file.Close()

// 	//Scanner to read the tsv file line by line
// 	scanner := bufio.NewScanner(file)

// 	var records [][]string
// 	//Read each line
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		//split the lines into fileds based on \t character
// 		fields := strings.Split(line, "\t")

// 		records = append(records, fields)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Println(err.Error())
// 	}

// 	var BooksStore []books

// 	var newRecords [][]string
// 	for i := range records {
// 		if i == 0 {
// 			newRecords = records[1:]
// 		}
// 	}

// 	for i := 0; i < len(newRecords); i++ {

// 		arr := newRecords[i]
// 		id, err1 := strconv.Atoi(arr[0])

// 		if err1 != nil {
// 			log.Println(err1.Error())
// 			panic(err1)
// 		}

// 		avg_rate, err2 := strconv.ParseFloat(arr[3], 64)

// 		if err2 != nil {
// 			//log.Println(err2.Error())
// 			log.Println(arr[3], arr[0])
// 			panic(err2)
// 		}

// 		num_pages, err3 := strconv.Atoi(arr[7])

// 		if err3 != nil {
// 			log.Println(err3.Error())
// 			panic(err3)
// 		}

// 		text_reviews_count, err4 := strconv.Atoi(arr[9])

// 		if err4 != nil {
// 			log.Println(err4.Error())
// 			panic(err4)
// 		}

// 		BooksStore = append(BooksStore, books{
// 			BOOK_ID:            id,
// 			TITLE:              arr[1],
// 			AUTHORS:            arr[2],
// 			AVERAGE_RATING:     avg_rate,
// 			LANGUAGE_CODE:      arr[6],
// 			NUM_PAGES:          num_pages,
// 			Text_REVIEWS_COUNT: text_reviews_count,
// 			PUBLICATION_DATE:   arr[10],
// 			PUBLISHER:          arr[11],
// 		})
// 	}

// 	log.Println(len(newRecords) == len(BooksStore))

// 	return BooksStore, nil
// }
