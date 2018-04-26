package block

import (
	"net/http"

	"github.com/nathenapse/learning-blockchain/pkg"
)

// Get All the blocks in the blockchain
func Get(blockchain *pkg.Blockchain) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		blocks, err := blockchain.AllBlocks()
		if err != nil {
			respondWithJSON(w, r, http.StatusInternalServerError, nil)
			return
		}
		respondWithJSON(w, r, http.StatusOK, blocks)
	}
}
