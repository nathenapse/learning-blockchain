package pkg

import (
	"reflect"
	"testing"
)

func TestNewInvalidBlockError(t *testing.T) {
	type args struct {
		message string
		block   *Block
	}
	tests := []struct {
		name string
		args args
		want *InvalidBlockError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInvalidBlockError(tt.args.message, tt.args.block); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInvalidBlockError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvalidBlockError_Error(t *testing.T) {
	type fields struct {
		message string
		block   Block
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &InvalidBlockError{
				message: tt.fields.message,
				block:   tt.fields.block,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("InvalidBlockError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBlockNotFoundError(t *testing.T) {
	type args struct {
		message string
		index   int
	}
	tests := []struct {
		name string
		args args
		want *BlockNotFoundError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBlockNotFoundError(tt.args.message, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBlockNotFoundError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockNotFoundError_Error(t *testing.T) {
	type fields struct {
		message string
		block   Block
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &BlockNotFoundError{
				message: tt.fields.message,
				block:   tt.fields.block,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("BlockNotFoundError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEnvNotFoundError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *EnvNotFoundError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEnvNotFoundError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEnvNotFoundError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvNotFoundError_Error(t *testing.T) {
	type fields struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EnvNotFoundError{
				message: tt.fields.message,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("EnvNotFoundError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
