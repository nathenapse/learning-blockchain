package main

import (
	"os"
	"os/user"
	"testing"

	"github.com/joho/godotenv"

	"github.com/nathenapse/learning-blockchain/app"
)

var a app.App
var usr, err = user.Current()

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"noenv"},
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name != "noenv" {
				godotenv.Load()
			}
			a = app.App{}
			testDb, found := os.LookupEnv("TEST_DB")
			if found == false && tt.name != "noenv" {
				t.Fatal("no env found")
			} else {

				a.Initialize(testDb)

				if _, err := os.Stat(usr.HomeDir + "/.blockchain"); os.IsNotExist(err) {
					t.Errorf("Initialize not creating directory at %v", err)
				}
				if _, err := os.Stat(usr.HomeDir + "/.blockchain/" + testDb); os.IsNotExist(err) {
					t.Errorf("Initialize not creating db file at %v", err)
				}
			}
		})
	}
}
