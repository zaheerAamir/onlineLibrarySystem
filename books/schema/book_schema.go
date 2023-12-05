package schema

type Books struct {
	TITLE              string  `json:"title"`
	AUTHORS            string  `json:"authors"`
	AVERAGE_RATING     float64 `json:"avg_rating"`
	LANGUAGE_CODE      string  `json:"language_code"`
	NUM_PAGES          int     `json:"num_pages"`
	Text_REVIEWS_COUNT int     `json:"text_reviews_count"`
	PUBLICATION_DATE   string  `json:"publication_date"`
	PUBLISHER          string  `json:"publisher"`
}

type Error struct {
	CODE       int    `json:"statuscode"`
	STATUSTEXT string `json:"statustext"`
	MESSAGE    string `json:"message"`
}
