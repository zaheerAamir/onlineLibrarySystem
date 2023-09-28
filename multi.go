package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	util "searchRecommend/utils"
	"time"
)

type BookQuery struct {
	BookDb *util.BookDb
}

type GoogleBooksResponse struct {
	Items []struct {
		VolumeInfo struct {
			Publisher   string   `json:"publisher"`
			Description string   `json:"description"`
			Categories  []string `json:"categories"`
		} `json:"volumeInfo"`
	} `json:"items"`
}

func tasks(id int) {

	log.Printf("Task %d Started\n", id)

	time.Sleep(time.Duration(id) * time.Second)

	log.Printf("Task %d Completed\n", id)
}

func Multi() {

	// numTasks := 5

	// tick := time.Now()
	// log.Println("Task started")
	// log.Println()

	// for i := 1; i <= numTasks; i++ {
	// 	go tasks(i)
	// }

	// // Sleep to allow goroutines to complete their tasks
	// time.Sleep((time.Duration(numTasks) + 1) * time.Second)

	// log.Println()
	// log.Println("Task Completed")
	// log.Printf("TIme took to complete the task: %s\n", time.Since(tick))

	Category()

}

func Category() {

	book_name := "The Hitchhiker's Guide to the Galaxy (Hitchhiker's Guide to the Galaxy  #1)"
	publisher := "Del Rey Books"
	api_key := "AIzaSyAHiZDcVX77s1bM9fLakwPTOb_35QFRtdo"

	URL := "https://www.googleapis.com/books/v1/volumes"
	encoded := url.QueryEscape(book_name)

	query := fmt.Sprintf("%s?key=%s&q=%s", URL, api_key, encoded)

	res, err := http.Get(query)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()

	var bookResponse GoogleBooksResponse
	if err := json.NewDecoder(res.Body).Decode(&bookResponse); err != nil {
		log.Fatal(err.Error())
	}

	if len(bookResponse.Items) > 0 {
		for _, v := range bookResponse.Items {
			if v.VolumeInfo.Publisher == publisher {
				log.Println(v.VolumeInfo.Publisher)
				log.Println(v.VolumeInfo.Description)
				log.Println(v.VolumeInfo.Categories)
			}
		}
	}

	// if len(bookResponse.Items) > 0 && bookResponse.Items[0].VolumeInfo.Publisher == publisher {
	// 	log.Println(bookResponse.Items[0].VolumeInfo.Categories, nil)
	// 	log.Println(bookResponse.Items[0].VolumeInfo.Description)
	// }

	log.Println(query)
}
