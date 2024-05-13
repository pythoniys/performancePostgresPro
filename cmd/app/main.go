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
	var dur int
	fmt.Println("Количество параллельных потоков:")
	fmt.Scanf("%d\n", &num)

	fmt.Println("Длительность вставки в минутах:")
	fmt.Scanf("%d\n", &dur)

	for j := 0; j < num; j++ {
		go CallGet(dur, storage)
	}
	time.Sleep(time.Duration(dur) * time.Minute)
}

func CallGet(t int, storage *postgres.Storage) {
	now := time.Now()
	for time.Since(now) < time.Minute*time.Duration(t) {
		GetResult(storage)
		time.Sleep(time.Millisecond * 10)
	}
}

func GetResult(storage *postgres.Storage) {

	fmt.Println("ready")
	err := storage.InsertIntoStorage(helpfunc.RandStringRunes(8), helpfunc.RandStringRunes(32))
	if err != nil {
		log.Fatal(err)
	}

}
