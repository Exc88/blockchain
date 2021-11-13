package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64 //block creation time
	Data []byte //very important information contained in the block
	PrevBlockHash []byte //stores the previous block hash
	Hash []byte //contains hash of the block
}
//todo *Data* should be contained in a different structure, but we created the structure for simplicity

func (b *Block) SetHash(){
	timestamp := []byte(strconv.FormatInt(b.Timestamp,10))
	headers := bytes.Join([][]byte{b.PrevBlockHash,b.Data,timestamp},[]byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string,PrevBlockHash []byte) *Block{
	block := &Block{time.Now().Unix(),[]byte(data),PrevBlockHash,[]byte{}}
	block.SetHash()
	return block
}