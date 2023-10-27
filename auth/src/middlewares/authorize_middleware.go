package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"searchRecommend/auth/schema"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type ContextKey string

const UserIDKey ContextKey = "user_id"

func AuthorizeAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
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
			w.Write(json)
		}
	}
}

func AuthorizeUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("content-type", "application/json")
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
				w.Write(json)
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				log.Println("Unable to extract claims")
				// Handle the error
				return
			}
			log.Println("CLAIMS2", claims)
			// OUTPUT
			// 2023/10/24 23:44:42 map[exp:1.698171316e+09 iat:1.698171256e+09 role:false user_id:5.462434098e+09]

			user_ID_Float := token.Claims.(jwt.MapClaims)["user_id"].(float64)
			user_ID := int64(user_ID_Float)
			log.Println(user_ID)
			ctx := context.WithValue(r.Context(), UserIDKey, user_ID)
			r = r.WithContext(ctx)

			role := token.Claims.(jwt.MapClaims)["role"].(bool)
			if token.Valid && !role {
				log.Println("admin:", role)
				next.ServeHTTP(w, r)
			}
			// Check the expiration time
			expirationTime := token.Claims.(jwt.MapClaims)["exp"].(float64)

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
			w.Write(json)
		}

	}
}
