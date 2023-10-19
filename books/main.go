package main

import (
	"context"
	"log"
	"net/http"
	"os"
	repository "searchRecommend/books/repositories"
	service "searchRecommend/books/services"
	route "searchRecommend/books/src"
	handler "searchRecommend/books/src/handlers"
	"searchRecommend/books/src/job"
	util "searchRecommend/books/utils"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

func startPostgresContainer() error {

	envPath := os.Getenv("API_KEY")

	log.Println("ENV_PATH", envPath)
	if envPath == "" {
		if errr := godotenv.Load("/home/aamir/Desktop/My_code/searchRecommend/.env"); errr != nil {

			panic(errr.Error())
		}

		ctx := context.Background()
		cli, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			return err
		}

		// Replace "your-postgres-container-name" with the actual container name
		containerID := os.Getenv("POSTGRES_CONTAINER_ID")

		// Start the PostgreSQL container
		err1 := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
		if err1 != nil {
			return err1
		}

		log.Println("PostgreSQL container started successfully......")
	}

	return nil
}

func main() {

	// Start the PostgreSQL container
	if err := startPostgresContainer(); err != nil {
		log.Fatalf("Failed to start PostgreSQL container: %v", err)
	}

	// Create a new cron scheduler
	c := cron.New()
	DB := &util.Db{}

	bookRepo := &repository.BookQuery{Db: DB}
	bookService := &service.BookService{BookRepo: bookRepo}
	bookHandler := &handler.BookHandler{Bookservice: bookService}

	rentBookRepo := &repository.RentBookRepo{Db: DB}
	rentBookService := &service.RentBookService{RentBookRepo: rentBookRepo}
	rentBookHandler := &handler.RentBookHandler{Bookservice: rentBookService}

	sendemail := &job.SendEmail{Db: DB}
	c.AddFunc("00 00 11 * *", sendemail.CheckUsers)

	// Start the cron scheduler
	c.Start()
	route.Setuproutes(bookHandler, rentBookHandler)

	log.Println("Server Running on Port :8080")
	http.ListenAndServe(":8080", nil)

}
