package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"searchRecommend/books/schema"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func AuthorizeAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		envPath := os.Getenv("API_KEY")

		log.Println("ENV_PATH", envPath)
		if envPath == "" {
			if errr := godotenv.Load("/home/aamir/Desktop/My_code/searchRecommend/.env"); errr != nil {

				panic(errr.Error())
			}
		}
		secret := os.Getenv("ACCESS_TOKEN_SECRET")
		log.Println(r.Header["Bearer"])

		var errorr schema.Error
		errorr.CODE = 401
		errorr.STATUSTEXT = http.StatusText(errorr.CODE)
		errorr.MESSAGE = "Unauthorized User!"

		if r.Header["Authorization"] != nil {
			tokenJwt := strings.Split(r.Header["Authorization"][0], " ")[1]
			log.Println(tokenJwt)
			token, err := jwt.Parse(tokenJwt, func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					json, err := json.Marshal(errorr)
					if err != nil {
						panic(err.Error())
					}
					w.WriteHeader(errorr.CODE)
					w.Header().Set("content-type", "application/json")
					w.Write(json)
				}
				return []byte(secret), nil
			})
			if err != nil {
				json, err := json.Marshal(errorr)
				if err != nil {
					panic(err.Error())
				}
				w.WriteHeader(errorr.CODE)
				w.Header().Set("content-type", "application/json")
				w.Write(json)
			}

			role := token.Claims.(jwt.MapClaims)["role"].(bool)
			if token.Valid {
				if role {
					log.Println("admin:", role)
					next.ServeHTTP(w, r)
				} else {
					log.Println("admin:", role)
					errorr.MESSAGE = "Unauthorized user. This a admin protected route!"
					json, err := json.Marshal(errorr)
					if err != nil {
						panic(err.Error())
					}
					w.WriteHeader(errorr.CODE)
					w.Header().Set("content-type", "application/json")
					w.Write(json)
				}

			}
			// Check the expiration time
			expirationTime, ok := token.Claims.(jwt.MapClaims)["exp"].(float64)
			if !ok {
				log.Println("unable to extract expiration time")
			}

			// Convert the expiration time to a Go time.Time object
			expiration := time.Unix(int64(expirationTime), 0)

			// Compare with the current time
			log.Println("Token Expired:", time.Now().After(expiration))

		} else {
			json, err := json.Marshal(errorr)
			if err != nil {
				panic(err.Error())
			}
			w.WriteHeader(errorr.CODE)
			w.Header().Set("content-type", "application/json")
			w.Write(json)
		}
	}
}

func AuthorizeUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		envPath := os.Getenv("API_KEY")

		log.Println("ENV_PATH", envPath)
		if envPath == "" {
			if errr := godotenv.Load("/home/aamir/Desktop/My_code/searchRecommend/.env"); errr != nil {

				panic(errr.Error())
			}
		}
		secret := os.Getenv("ACCESS_TOKEN_SECRET")
		log.Println(r.Header["Bearer"])

		var errorr schema.Error
		errorr.CODE = 401
		errorr.STATUSTEXT = http.StatusText(errorr.CODE)
		errorr.MESSAGE = "Unauthorized User!"

		if r.Header["Authorization"] != nil {
			tokenJwt := strings.Split(r.Header["Authorization"][0], " ")[1]
			log.Println(tokenJwt)
			token, err := jwt.Parse(tokenJwt, func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					json, err := json.Marshal(errorr)
					if err != nil {
						panic(err.Error())
					}
					w.WriteHeader(errorr.CODE)
					w.Header().Set("content-type", "application/json")
					w.Write(json)
				}
				return []byte(secret), nil
			})
			if err != nil {
				json, err := json.Marshal(errorr)
				if err != nil {
					panic(err.Error())
				}
				w.WriteHeader(errorr.CODE)
				w.Header().Set("content-type", "application/json")
				w.Write(json)
			}

			role := token.Claims.(jwt.MapClaims)["role"].(bool)
			if token.Valid {
				log.Println("admin:", role)
				next.ServeHTTP(w, r)
			}
			// Check the expiration time
			expirationTime, ok := token.Claims.(jwt.MapClaims)["exp"].(float64)
			if !ok {
				log.Println("unable to extract expiration time")
			}

			// Convert the expiration time to a Go time.Time object
			expiration := time.Unix(int64(expirationTime), 0)

			// Compare with the current time
			log.Println("Token Expired:", time.Now().After(expiration))

		} else {
			json, err := json.Marshal(errorr)
			if err != nil {
				panic(err.Error())
			}
			w.WriteHeader(errorr.CODE)
			w.Header().Set("content-type", "application/json")
			w.Write(json)
		}

	}
}
