package main

import (
	"fmt"

	"github.com/NabeelAhmed1721/GoBlockChain/blockchain"
)

func main() {
	b := blockchain.InitBlockChain()

	b.AddBlock("Hello!")
	b.AddBlock("Bye!")

	for _, block := range b.Blocks {
		fmt.Printf("Block Hash: %x\n", block.Hash)
		fmt.Printf("Block Data: %s\n", block.Data)
	}
}
