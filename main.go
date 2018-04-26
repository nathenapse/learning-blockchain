package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"sync"
	"time"

	bolt "github.com/coreos/bbolt"
	"github.com/joho/godotenv"

	"github.com/nathenapse/learning-blockchain/api"
	"github.com/nathenapse/learning-blockchain/pkg"
)

// Message takes incoming JSON payload for writing Money
type Message struct {
	MONEY int
}

var mutex = &sync.Mutex{}
var usr, err = user.Current()

func main() {
	os.Mkdir(usr.HomeDir+"/.blockchain", os.ModePerm)
	godotenv.Load()
	db, err := setupDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	mutex.Lock()
	blockchain, err := pkg.NewBlockchain(db)
	mutex.Unlock()
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Fatal(runServer(blockchain))

}

// RunServer create a web listener for the webservice
func runServer(blockchain *pkg.Blockchain) error {
	mux := api.MakeMuxRouter(blockchain)
	httpPort, found := os.LookupEnv("PORT")

	if found == false {
		httpPort = "8080"
		fmt.Println("\"PORT\" Not Found in env using default 8080")
	}
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

func setupDB() (*bolt.DB, error) {
	database, found := os.LookupEnv("DB")

	if found == false {
		database = "blockchain.db"
	}

	fullpath := usr.HomeDir + "/.blockchain/" + database

	db, err := bolt.Open(fullpath, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	fmt.Println("DB Setup Done")
	return db, nil
}
