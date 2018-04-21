package block

import (
	"net/http"

	"github.com/nathenapse/blockchain/pkg"
)

// Get All the blocks in the blockchain
func Get(blockchain *pkg.Blockchain) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		blocks := blockchain.AllBlocks()
		respondWithJSON(w, r, http.StatusOK, blocks)
	}
}
