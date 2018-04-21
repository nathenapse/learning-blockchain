package pkg

// Blockchain is the list of linked blocks
type Blockchain struct {
	blocks []*Block
}

// NewBlockchain initializes a new blockchain
func NewBlockchain() *Blockchain {
	newBlock := NewGenesisBlock()
	return &Blockchain{[]*Block{newBlock}}
}

// AddBlock Addes a block in a blockchain
func (bc *Blockchain) AddBlock(money int) (*Block, error) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(money, prevBlock)
	if !IsBlockValid(newBlock, prevBlock) {
		return nil, NewInvalidBlockError("Invalid Block", newBlock)
	}
	bc.blocks = append(bc.blocks, newBlock)
	return newBlock, nil
}

// GetBlock gets you the block with a specific block
func (bc *Blockchain) GetBlock(index int) (*Block, error) {
	for i := 0; i < len(bc.blocks); i++ {
		if bc.blocks[i].Index == index {
			return bc.blocks[i], nil
		}
	}
	return nil, NewBlockNotFoundError("Block Not found", index)
}

// AllBlocks get all blocks in the blockchain
func (bc *Blockchain) AllBlocks() []Block {
	blocks := make([]Block, len(bc.blocks))
	for i := 0; i < len(bc.blocks); i++ {
		blocks[i] = *bc.blocks[i]
	}
	return blocks
}
