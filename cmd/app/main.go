package main

import (
	"log"
	"performanceTest/internal/storage/postgres"
)

func main() {
	storage, err := postgres.New()
	if err != nil {
		log.Fatal(err)
	}
	_ = storage

}
