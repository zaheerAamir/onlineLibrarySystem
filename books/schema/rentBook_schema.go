package schema

type RentBookDTO struct {
	RENTDURATION struct {
		MONTHS int `json:"months"`
		DAYS   int `json:"days"`
	} `json:"rentDuration"`
}

type RentBookSuccess struct {
	STATUS_CODE int    `json:"status_code"`
	STATUS_TEXT string `json:"status_text"`
	MESSAGE     string `json:"message"`
}

type RentBookSchema struct {
	USER_ID     int64
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
