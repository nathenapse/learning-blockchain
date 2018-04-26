package api

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/nathenapse/learning-blockchain/pkg"
)

func TestMakeMuxRouter(t *testing.T) {
	type args struct {
		blockchain *pkg.Blockchain
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeMuxRouter(tt.args.blockchain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeMuxRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
