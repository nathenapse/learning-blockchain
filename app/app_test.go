package app

import "testing"

func TestApp_Initialize(t *testing.T) {
	type args struct {
		database string
	}
	tests := []struct {
		name    string
		a       *App
		args    args
		wantErr bool
	}{
		{"Success", &App{}, args{"test.db"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.Initialize(tt.args.database); (err != nil) != tt.wantErr {
				t.Errorf("App.Initialize() error = %v, wantErr %v", err, tt.wantErr)
				if tt.a.DB == nil {
					t.Errorf("App.Initialize() DID not initialize the database")
				}
			}
		})
	}
}

func TestApp_Run(t *testing.T) {
	type args struct {
		port string
	}
	tests := []struct {
		name string
		a    *App
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Run(tt.args.port)
		})
	}
}
