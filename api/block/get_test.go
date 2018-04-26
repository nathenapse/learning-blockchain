package block

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/nathenapse/learning-blockchain/pkg"
)

func TestGet(t *testing.T) {
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
			if got := Get(tt.args.blockchain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
