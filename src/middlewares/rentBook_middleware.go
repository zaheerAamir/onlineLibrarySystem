package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"searchRecommend/schema"
)

func RentBook(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var error schema.Error
		error.CODE = 400
		error.STATUSTEXT = http.StatusText(400)
		error.MESSAGE = "Bad request it is a POST route"

		if r.Method != "POST" {
			json, err := json.Marshal(error)
			if err != nil {
				panic(err.Error())
			}
			w.WriteHeader(error.CODE)
			w.Header().Set("content-type", "application/json")
			w.Write(json)
		} else if r.Header.Get("Content-Type") != "application/json" {
			error.MESSAGE = "Content-Type should be application/json"
			json, err := json.Marshal(error)
			if err != nil {
				panic(err.Error())
			}
			w.WriteHeader(error.CODE)
			w.Header().Set("content-type", "application/json")
			w.Write(json)

		} else {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				panic(err.Error())
			}
			defer r.Body.Close()

			var duration schema.RentBookDTO
			if err1 := json.Unmarshal(body, &duration); err1 != nil {
				panic(err1.Error())
			}
			if duration.USER_ID == 0 {
				error.CODE = 404
				error.STATUSTEXT = http.StatusText(error.CODE)
				error.MESSAGE = "Unauthorized user!"
				json, err := json.Marshal(error)
				if err != nil {
					panic(err.Error())
				}
				w.WriteHeader(error.CODE)
				w.Header().Set("content-type", "application/json")
				w.Write(json)
			} else if duration.RENTDURATION.DAYS == 0 && duration.RENTDURATION.MONTHS == 0 {
				error.MESSAGE = "Please enter anyone field days or month"
				json, err := json.Marshal(error)
				if err != nil {
					panic(err.Error())
				}
				w.WriteHeader(error.CODE)
				w.Header().Set("content-type", "application/json")
				w.Write(json)
			} else if duration.RENTDURATION.DAYS > 30 {
				error.MESSAGE = "Please enter a valid filed Days, should be less than 31!"
				json, err := json.Marshal(error)
				if err != nil {
					panic(err.Error())
				}
				w.WriteHeader(error.CODE)
				w.Header().Set("content-type", "application/json")
				w.Write(json)
			} else if duration.RENTDURATION.MONTHS > 3 || duration.RENTDURATION.MONTHS == 3 && duration.RENTDURATION.DAYS != 0 {
				error.MESSAGE = "maximum period for rent is 3 months"
				json, err := json.Marshal(error)
				if err != nil {
					panic(err.Error())
				}
				w.WriteHeader(error.CODE)
				w.Header().Set("content-type", "application/json")
				w.Write(json)
			} else {
				r.Body = io.NopCloser(bytes.NewBuffer(body))
				next.ServeHTTP(w, r)
			}

		}
	})
}

func GiveBack(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var error schema.Error
		error.CODE = 400
		error.STATUSTEXT = http.StatusText(400)
		error.MESSAGE = "Bad request it is a PUT route"

		if r.Method != "PUT" {
			json, err := json.Marshal(error)
			if err != nil {
				panic(err.Error())
			}
			w.WriteHeader(error.CODE)
			w.Header().Set("content-type", "application/json")
			w.Write(json)
		} else if r.Header.Get("Content-Type") != "application/json" {
			error.MESSAGE = "Content-Type should be application/json"
			json, err := json.Marshal(error)
			if err != nil {
				panic(err.Error())
			}
			w.WriteHeader(error.CODE)
			w.Header().Set("content-type", "application/json")
			w.Write(json)

		} else {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				panic(err.Error())
			}
			defer r.Body.Close()

			var GiveBack schema.GiveBookBackDTO
			if err1 := json.Unmarshal(body, &GiveBack); err1 != nil {
				panic(err1.Error())
			}

			if GiveBack.BOOK_ID == 0 || GiveBack.EMAIL == "" {
				error.MESSAGE = "Please enter both the feilds book_id and email!"
				json, err := json.Marshal(error)
				if err != nil {
					panic(err.Error())
				}
				w.WriteHeader(error.CODE)
				w.Header().Set("content-type", "application/json")
				w.Write(json)
			} else {
				r.Body = io.NopCloser(bytes.NewBuffer(body))
				next.ServeHTTP(w, r)
			}

		}
	})
}
