package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nathenapse/blockchain/api/block"
	"github.com/nathenapse/blockchain/pkg"
)

// MakeMuxRouter returns the routes
func MakeMuxRouter(blockchain *pkg.Blockchain) http.Handler {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/", block.Get(blockchain)).Methods("GET")
	muxRouter.HandleFunc("/", block.Post(blockchain)).Methods("POST")
	return muxRouter
}
