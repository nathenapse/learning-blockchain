package pkg

import (
	"reflect"
	"testing"
)

func TestNewBlockchain(t *testing.T) {
	tests := []struct {
		name string
		want *Blockchain
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBlockchain(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBlockchain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockchain_AddBlock(t *testing.T) {
	type fields struct {
		blocks []*Block
	}
	type args struct {
		money int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Block
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &Blockchain{
				blocks: tt.fields.blocks,
			}
			got, err := bc.AddBlock(tt.args.money)
			if (err != nil) != tt.wantErr {
				t.Errorf("Blockchain.AddBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blockchain.AddBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockchain_GetBlock(t *testing.T) {
	type fields struct {
		blocks []*Block
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Block
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &Blockchain{
				blocks: tt.fields.blocks,
			}
			got, err := bc.GetBlock(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Blockchain.GetBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blockchain.GetBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockchain_AllBlocks(t *testing.T) {
	type fields struct {
		blocks []*Block
	}
	tests := []struct {
		name   string
		fields fields
		want   []Block
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &Blockchain{
				blocks: tt.fields.blocks,
			}
			if got := bc.AllBlocks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Blockchain.AllBlocks() = %v, want %v", got, tt.want)
			}
		})
	}
}
