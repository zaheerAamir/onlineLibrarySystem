package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"searchRecommend/auth/schema"
)

func CreateToken(next http.HandlerFunc) http.HandlerFunc {
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
			var token schema.RefreshTokenDTO
			if err := json.Unmarshal(body, &token); err != nil {
				panic(err.Error())

			}
			log.Println(token)
			if token.REFRESH_TOKEN == "" {
				error.MESSAGE = "Refresh token is empty!"
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
