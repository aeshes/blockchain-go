package main

import "time"
import "strconv"
import "bytes"
import "crypto/sha256"

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{ block.PrevBlockHash, block.Data, timestamp }, []byte{})
	hash := sha256.Sum256(headers)

	block.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{ time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash();
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}