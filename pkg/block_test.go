package pkg

import (
	"reflect"
	"testing"
)

func TestNewGenesisBlock(t *testing.T) {
	tests := []struct {
		name string
		want *Block
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGenesisBlock(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGenesisBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlock_calculateHash(t *testing.T) {
	type fields struct {
		Index     int
		Timestamp int64
		Money     int
		Hash      string
		PrevHash  string
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
			block := &Block{
				Index:     tt.fields.Index,
				Timestamp: tt.fields.Timestamp,
				Money:     tt.fields.Money,
				Hash:      tt.fields.Hash,
				PrevHash:  tt.fields.PrevHash,
			}
			if got := block.calculateHash(); got != tt.want {
				t.Errorf("Block.calculateHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBlock(t *testing.T) {
	type args struct {
		money     int
		prevBlock *Block
	}
	tests := []struct {
		name string
		args args
		want *Block
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBlock(tt.args.money, tt.args.prevBlock); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBlockValid(t *testing.T) {
	type args struct {
		newBlock *Block
		oldBlock *Block
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlockValid(tt.args.newBlock, tt.args.oldBlock); got != tt.want {
				t.Errorf("IsBlockValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlock_Serialize(t *testing.T) {
	tests := []struct {
		name    string
		block   *Block
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.block.Serialize()
			if (err != nil) != tt.wantErr {
				t.Errorf("Block.Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Block.Serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeserializeBlock(t *testing.T) {
	type args struct {
		d []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Block
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeserializeBlock(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeserializeBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeserializeBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}
