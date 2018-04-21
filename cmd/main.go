package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"

	"github.com/nathenapse/blockchain/api"
	"github.com/nathenapse/blockchain/pkg"
)

// Message takes incoming JSON payload for writing Money
type Message struct {
	MONEY int
}

var mutex = &sync.Mutex{}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	mutex.Lock()
	blockchain := pkg.NewBlockchain()
	mutex.Unlock()
	fmt.Print(blockchain)

	log.Fatal(runServer(blockchain))

}

// RunServer create a web listener for the webservice
func runServer(blockchain *pkg.Blockchain) error {
	mux := api.MakeMuxRouter(blockchain)
	httpPort := os.Getenv("PORT")
	log.Println("HTTP Server Listening on port :", httpPort)
	s := &http.Server{
		Addr:           ":" + httpPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()

	return err
}
