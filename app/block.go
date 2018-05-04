package app

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/nathenapse/learning-blockchain/pkg"
)

func (app *App) getBlocks() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		blocks, err := pkg.AllBlocks(app.DB)
		if err != nil {
			respondWithJSON(w, r, http.StatusInternalServerError, nil)
			return
		}
		respondWithJSON(w, r, http.StatusOK, blocks)
	}
}

// Message is the request the server is expecting
type Message struct {
	Money int `json:"money"`
}

var mutex = &sync.Mutex{}

// CreateBlock Create a new Block to a blockchain
func (app *App) createBlock() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var m Message

		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			respondWithJSON(w, r, http.StatusBadRequest, m)
			return
		}

		defer r.Body.Close()

		if m.Money <= 0 {
			respondWithJSON(w, r, http.StatusBadRequest, m)
			return
		}

		mutex.Lock()
		newBlock, err := pkg.AddBlock(m.Money, app.DB)
		mutex.Unlock()

		if err != nil {
			respondWithJSON(w, r, http.StatusBadRequest, *newBlock)
		} else {
			respondWithJSON(w, r, http.StatusCreated, *newBlock)
		}
	}
}
