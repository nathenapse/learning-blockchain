package main

import (
	"reflect"
	"testing"

	bolt "github.com/coreos/bbolt"
	"github.com/nathenapse/learning-blockchain/pkg"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_runServer(t *testing.T) {
	type args struct {
		blockchain *pkg.Blockchain
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := runServer(tt.args.blockchain); (err != nil) != tt.wantErr {
				t.Errorf("runServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_setupDB(t *testing.T) {
	tests := []struct {
		name    string
		want    *bolt.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setupDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("setupDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
