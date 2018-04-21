package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nathenapse/learning-blockchain/api/block"
	"github.com/nathenapse/learning-blockchain/pkg"
)

// MakeMuxRouter returns the routes
func MakeMuxRouter(blockchain *pkg.Blockchain) http.Handler {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/", block.Get(blockchain)).Methods("GET")
	muxRouter.HandleFunc("/", block.Post(blockchain)).Methods("POST")
	return muxRouter
}
