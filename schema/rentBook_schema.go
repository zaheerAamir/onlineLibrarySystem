package schema

type RentBookDTO struct {
	USER_ID      int `josn:"user_id"`
	RENTDURATION struct {
		MONTHS int `json:"months"`
		DAYS   int `json:"days"`
	} `json:"rentDuration"`
}

type RentBookSchema struct {
	USER_ID     int
	RENTED_DATE struct {
		DAY   int
		MONTH int
		YEAR  int
	}
	RENT_DAY struct {
		DAYS   int
		MONTHS int
		YEARS  int
	}
}

type GiveBookBackDTO struct {
	BOOK_ID int    `json:"book_id"`
	EMAIL   string `json:"email"`
}

type GiveBookBackSchema struct {
	BOOK_ID int
	EMAIL   string
}
