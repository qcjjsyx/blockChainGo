package blc

type Blockchain struct {
	Blocks []*Block
}

func (blockchian *Blockchain) AddBlock(data string) {
	preHash := blockchian.Blocks[len(blockchian.Blocks)-1].Hash
	block := NewBlock(preHash, data)
	blockchian.Blocks = append(blockchian.Blocks, block)
}

func GenesisBlock() *Block {
	return NewBlock([]byte{}, "first Block")
}

func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}
