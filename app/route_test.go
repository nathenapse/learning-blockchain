package app

import (
	"net/http"
	"testing"
)

func TestApp_InitializeRoutes(t *testing.T) {
	tests := []struct {
		name string
		app  *App
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.app.InitializeRoutes()
		})
	}
}

func Test_respondWithJSON(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		r       *http.Request
		code    int
		payload interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respondWithJSON(tt.args.w, tt.args.r, tt.args.code, tt.args.payload)
		})
	}
}
