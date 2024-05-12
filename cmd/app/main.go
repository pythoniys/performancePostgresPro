package main

import (
	"fmt"
	"log"
	"performanceTest/cmd/helpfunc"
	"performanceTest/internal/storage/postgres"
	"time"
)

func main() {
	storage, err := postgres.New()
	if err != nil {
		log.Fatal(err)
	}
	_ = storage

	var num int
	var dur time.Duration
	fmt.Println("Количество параллельных потоков:")
	fmt.Scanf("%d\n", &num)
	fmt.Println(num)

	fmt.Println("Длительность вставки:")
	fmt.Scanf("%d\n", &dur)
	fmt.Println(dur)

	for i := 0; i < num; i++ {
		go GetResult(dur, num, storage)
	}

	go GetResult(dur, num, storage)
}

func GetResult(duration time.Duration, numOfThreads int, storage *postgres.Storage) {
	_ = duration
	_ = numOfThreads
	fmt.Println("зашел")

	err := storage.InsertIntoStorage(helpfunc.RandStringRunes(8), helpfunc.RandStringRunes(32))
	if err != nil {
		log.Fatal(err)
		//return fmt.Errorf("Get result %v", err)
	}

}
