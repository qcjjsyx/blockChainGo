package blc

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	TimeStamp    int64
	Data         []byte
	PreBlockHash []byte
	Hash         []byte
}

func (block *Block) SetHsah() {
	timeStamp := []byte(strconv.FormatInt(block.TimeStamp, 10))

	header := bytes.Join([][]byte{timeStamp, block.Data, block.PreBlockHash}, []byte{})
	hash := sha256.Sum256((header))

	block.Hash = hash[:]
}

func NewBlock(preHash []byte, data string) *Block {
	block := &Block{time.Now().Unix(), []byte(data), preHash, []byte{}}
	block.SetHsah()
	return block
}
