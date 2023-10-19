package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"searchRecommend/auth/schema"
)

func SignUp(next http.HandlerFunc) http.HandlerFunc {
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

			var data schema.UserDto
			if err := json.Unmarshal(body, &data); err != nil {
				panic(err.Error())

			}
			if data.FIRST_NAME == "" || data.LAST_NAME == "" || data.EMAIL == "" || data.PASSWORD == "" {
				error.MESSAGE = "missing values Request body must conatin this feilds: last_name, first_name, email, password"
				json, err := json.Marshal(error)
				if err != nil {
					panic(err.Error())
				}
				w.WriteHeader(error.CODE)
				w.Header().Set("content-type", "application/json")
				w.Write(json)
			} else {
				log.Println(data)
				r.Body = io.NopCloser(bytes.NewBuffer(body))
				next.ServeHTTP(w, r)
			}

		}

	})
}

func Login(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var error schema.Error
		error.CODE = 400
		error.STATUSTEXT = http.StatusText(400)
		error.MESSAGE = "Bad request it is a GET route"

		if r.Method != "GET" {
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

			var data schema.UserLoginDto
			if err := json.Unmarshal(body, &data); err != nil {
				panic(err.Error())

			}
			if data.EMAIL == "" || data.PASSWORD == "" {
				error.MESSAGE = "missing values Request body must conatin this feilds: email, password"
				json, err := json.Marshal(error)
				if err != nil {
					panic(err.Error())
				}
				w.WriteHeader(error.CODE)
				w.Header().Set("content-type", "application/json")
				w.Write(json)
			} else {
				log.Println(data)
				r.Body = io.NopCloser(bytes.NewBuffer(body))
				next.ServeHTTP(w, r)
			}
		}

	})

}

func Logout(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var error schema.Error
		error.CODE = 400
		error.STATUSTEXT = http.StatusText(400)
		error.MESSAGE = "Bad request it is a DELETE route"

		if r.Method != "DELETE" {
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
			var email schema.Logout
			if err1 := json.Unmarshal(body, &email); err1 != nil {
				panic(err1.Error())
			}
			log.Println(email.EMAIL)

			if email.EMAIL == "" {
				error.MESSAGE = "Please Enter the email Id"
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

	}
}
