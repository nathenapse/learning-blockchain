package app

import (
	"log"
	"net/http"
	"os"
	"os/user"
	"time"

	"github.com/gorilla/mux"

	bolt "github.com/coreos/bbolt"
)

// App is the structure that holds the router and database
type App struct {
	Router *mux.Router
	DB     *bolt.DB
}

var usr, err = user.Current()

// Initialize creates the folder, file and database
func (a *App) Initialize(database string) error {
	var directory = usr.HomeDir + "/.blockchain"
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if err := os.Mkdir(directory, os.ModePerm); err != nil {
			return err
		}
	}

	fullpath := directory + database

	db, err := bolt.Open(fullpath, 0600, nil)
	if err != nil {
		return err
	}
	a.DB = db

	log.Println("DB Setup Done")

	a.Router = mux.NewRouter()

	a.InitializeRoutes()

	return nil
}

// Run listens to the app requests in the port specified
func (a *App) Run(port string) {

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        a.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("HTTP Server Listening on port :", port)

	log.Fatal(s.ListenAndServe())
}
