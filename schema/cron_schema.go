package schema

type RentDetails struct {
	USER_ID     int64
	EMAIL       string
	FIRST_NAME  string
	LAST_NAME   string
	BOOK_ID     int
	TITLE       string
	RENTED_DATE string
	RENTDATE    struct {
		DAY   int
		MONTH int
		YEAR  int
	}
}
