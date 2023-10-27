package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"searchRecommend/auth/schema"
)

func isGmailAddress(email string) bool {
	// Define a regular expression pattern for Gmail addresses
	pattern := `^[a-zA-Z0-9._%+-]+@gmail\.com$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

func SignUp(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("content-type", "application/json")
		var error schema.Error
		error.CODE = 400
		error.STATUSTEXT = http.StatusText(400)
		error.MESSAGE = "Bad request it is a POST route"

		if r.Method != http.MethodPost {
			json, err := json.Marshal(error)
			if err != nil {
				panic(err.Error())
			}
			w.WriteHeader(error.CODE)
			w.Write(json)
		} else if r.Header.Get("Content-Type") != "application/json" {
			error.MESSAGE = "Content-Type should be application/json"
			json, err := json.Marshal(error)
			if err != nil {
				panic(err.Error())
			}
			w.WriteHeader(error.CODE)
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
				w.Write(json)
			} else if !isGmailAddress(data.EMAIL) {
				error.MESSAGE = "Given Gmail ID is not a Gmail address!"
				json, err := json.Marshal(error)
				if err != nil {
					panic(err.Error())
				}
				w.WriteHeader(error.CODE)
				w.Write(json)

			} else {
				log.Println("[User Middleware] [SignUp route]", data)
				r.Body = io.NopCloser(bytes.NewBuffer(body))
				next.ServeHTTP(w, r)
			}

		}

	})
}

func Login(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("content-type", "application/json")
		var error schema.Error
		error.CODE = 400
		error.STATUSTEXT = http.StatusText(400)
		error.MESSAGE = "Bad request it is a POST route"

		if r.Method != http.MethodPost {
			json, err := json.Marshal(error)
			if err != nil {
				panic(err.Error())
			}
			w.WriteHeader(error.CODE)
			w.Write(json)
		} else if r.Header.Get("Content-Type") != "application/json" {
			error.MESSAGE = "Content-Type should be application/json"
			json, err := json.Marshal(error)
			if err != nil {
				panic(err.Error())
			}
			w.WriteHeader(error.CODE)
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
				w.Write(json)
			} else {
				log.Println("[User Middleware] [Login route]", data)
				r.Body = io.NopCloser(bytes.NewBuffer(body))
				next.ServeHTTP(w, r)
			}
		}

	})

}
