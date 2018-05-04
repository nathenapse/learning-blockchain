package pkg

import (
	"sort"

	bolt "github.com/coreos/bbolt"
)

// NewBlockchain initializes a new blockchain
func NewBlockchain(db *bolt.DB) error {
	var last []byte
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blockchain"))
		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte("blockchain"))
			if err != nil {
				return err
			}
			serialized, err := genesis.Serialize()
			if err != nil {
				return err
			}
			if err := b.Put([]byte(genesis.Hash), serialized); err != nil {
				return err
			}
			if err := b.Put([]byte("last"), []byte(genesis.Hash)); err != nil {
				return err
			}

			last = []byte(genesis.Hash)
		} else {
			last = b.Get([]byte("last"))
		}
		return nil
	})

	if err != nil {
		return err
	}

	return err
}

// AddBlock Addes a block in a blockchain
func AddBlock(money int, db *bolt.DB) (*Block, error) {
	var lastHash []byte
	var lastBlock *Block
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blockchain"))
		lastHash = b.Get([]byte("last"))
		block, err := DeserializeBlock(b.Get(lastHash))
		lastBlock = block
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	newBlock := NewBlock(money, lastBlock)
	if !IsBlockValid(newBlock, lastBlock) {
		return nil, NewInvalidBlockError("Invalid Block", newBlock)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blockchain"))

		serialized, err := newBlock.Serialize()
		if err != nil {
			return err
		}

		if err = b.Put([]byte(newBlock.Hash), serialized); err != nil {
			return err
		}

		if err := b.Put([]byte("last"), []byte(newBlock.Hash)); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return newBlock, nil
}

// GetBlock gets you the block with a specific block
func GetBlock(hash []byte, db *bolt.DB) (*Block, error) {
	var block *Block
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blockchain"))
		dblock, err := DeserializeBlock(b.Get(hash))
		block = dblock
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return block, nil
}

// AllBlocks get all blocks in the blockchain that are sorted
func AllBlocks(db *bolt.DB) ([]*Block, error) {
	var blocks []*Block
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blockchain"))
		b.ForEach(func(k, v []byte) error {
			block, err := DeserializeBlock(v)
			if err != nil {
				return err
			}
			blocks = append(blocks, block)
			return nil
		})

		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].Index < blocks[j].Index
	})
	return blocks, nil
}
