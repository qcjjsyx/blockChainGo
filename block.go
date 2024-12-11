package blockchain

type Block struct {
	timeStamp    int64
	data         []byte
	preBlockHash []byte
	Hash         []byte
}
