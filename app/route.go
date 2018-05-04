package app

import (
	"encoding/json"
	"net/http"
)

// InitializeRoutes returns the routes
func (app *App) InitializeRoutes() {
	app.Router.HandleFunc("/", app.getBlocks()).Methods("GET")
	app.Router.HandleFunc("/", app.createBlock()).Methods("POST")
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
