package block

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/nathenapse/learning-blockchain/pkg"
)

func TestPost(t *testing.T) {
	type args struct {
		blockchain *pkg.Blockchain
	}
	tests := []struct {
		name string
		args args
		want func(w http.ResponseWriter, r *http.Request)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Post(tt.args.blockchain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Post() = %v, want %v", got, tt.want)
			}
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
