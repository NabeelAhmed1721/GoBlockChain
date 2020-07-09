package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// BlockChain Storage
type BlockChain struct {
	Blocks []*Block
}

// Block Individual
type Block struct {
	Hash		[]byte
	Data		[]byte
	PrevHash 	[]byte
}

// computeHash unique for each block in chain
func (block *Block) computeHash() {
	blockInfo := bytes.Join([][]byte{block.Data, block.PrevHash}, []byte{})
	blockHash := sha256.Sum256(blockInfo)
	block.Hash = blockHash[:]
}

// createBlock at a given stage in blockchain
func createBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.computeHash()
	return block
}

// AddBlock to blockchain end
func (blockchain *BlockChain) AddBlock(data string) {
	prevBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := createBlock(data, prevBlock.Hash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

// Genesis block stores empty hash. Origin of blockchain
func Genesis() *Block {
	return createBlock("GenesisBlock", []byte{})
}

// InitBlockChain must be declared. Creates empty blockchain with genesis block
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}