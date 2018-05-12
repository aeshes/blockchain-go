package main

import "fmt"
import "strconv"

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC")
	bc.AddBlock("Send 2 more BTC")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := NewProofOfWork(block)
		fmt.Printf("Proof of Work: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}