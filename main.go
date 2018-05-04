package main

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"

	"github.com/nathenapse/learning-blockchain/app"
	"github.com/nathenapse/learning-blockchain/pkg"
)

// Message takes incoming JSON payload for writing Money
type Message struct {
	MONEY int
}

var mutex = &sync.Mutex{}

func main() {
	godotenv.Load()
	a := app.App{}

	database, found := os.LookupEnv("DB")
	if found == false {
		database = "blockchain.db"
	}

	err := a.Initialize(database)

	if err != nil {
		log.Fatal(err)
	}

	mutex.Lock()
	err = pkg.NewBlockchain(a.DB)
	if err != nil {
		log.Fatal(err)
	}
	mutex.Unlock()

	port, found := os.LookupEnv("PORT")
	if found == false {
		port = "8080"
	}
	a.Run(port)

}
