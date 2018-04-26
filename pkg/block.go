package pkg

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"strconv"
	"time"
)

// Block is basic construct of the Block
type Block struct {
	Index     int
	Timestamp int64
	Money     int
	Hash      string
	PrevHash  string
}

// NewGenesisBlock creates the GenesisBlock
func NewGenesisBlock() *Block {
	block := &Block{0, time.Now().Unix(), 0, "", ""}
	block.Hash = block.calculateHash()
	return block
}

func (block *Block) calculateHash() string {
	record := strconv.Itoa(block.Index) + strconv.FormatInt(block.Timestamp, 10) + strconv.Itoa(block.Money) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// NewBlock Creates new Block
func NewBlock(money int, prevBlock *Block) *Block {

	block := &Block{prevBlock.Index + 1, time.Now().Unix(), money, "", prevBlock.Hash}
	block.Hash = block.calculateHash()
	return block
}

// IsBlockValid make sure block is valid by checking index, and comparing the hash of the previous block
func IsBlockValid(newBlock *Block, oldBlock *Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if newBlock.calculateHash() != newBlock.Hash {
		return false
	}

	return true
}

// Serialize serializes the block to []byte
func (block *Block) Serialize() ([]byte, error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	if err := encoder.Encode(block); err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}

// DeserializeBlock unserializes block
func DeserializeBlock(d []byte) (*Block, error) {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	if err := decoder.Decode(&block); err != nil {
		return nil, err
	}

	return &block, nil
}
