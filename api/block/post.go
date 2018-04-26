package block

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/nathenapse/learning-blockchain/pkg"
)

// Message takes incoming JSON payload for writing Money
type Message struct {
	Money int `json:"money"`
}

var mutex = &sync.Mutex{}

// Post Create a new Block to a blockchain
func Post(blockchain *pkg.Blockchain) func(w http.ResponseWriter, r *http.Request) {
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
		newBlock, err := blockchain.AddBlock(m.Money)
		mutex.Unlock()

		if err != nil {
			respondWithJSON(w, r, http.StatusBadRequest, *newBlock)
		} else {
			respondWithJSON(w, r, http.StatusCreated, *newBlock)
		}
	}
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
