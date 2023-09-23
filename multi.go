package main

import (
	"log"
	"time"
)

func tasks(id int) {

	log.Printf("Task %d Started\n", id)

	time.Sleep(time.Duration(id) * time.Second)

	log.Printf("Task %d Completed\n", id)
}

func Multi() {

	numTasks := 5

	tick := time.Now()
	log.Println("Task started")
	log.Println()

	for i := 1; i <= numTasks; i++ {
		go tasks(i)
	}

	// Sleep to allow goroutines to complete their tasks
	time.Sleep((time.Duration(numTasks) + 1) * time.Second)

	log.Println()
	log.Println("Task Completed")
	log.Printf("TIme took to complete the task: %s\n", time.Since(tick))

}
