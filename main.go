package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
)

// MoonBucks is a simple blockchain implementation in Go.
// It demonstrates the basic concepts of a blockchain, including blocks, hashing, and proof of work.
// It is not intended for production use and should not be used as such.
// The code is for educational purposes only.

// Part 1: Block Structure, Hashing, and Creation

type Block struct { // Block represents a single block in the blockchain
	Timestamp     int64 // This is a simplified version of a Block in a Blockchain
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func (b *Block) BlockHash() { // BlockHash takes Block fields, concatenates them, and hashes the result
	Timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, Timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Part 2: Implementing the Blockchain

type Blockchain struct { // My first Blockchain!
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) { // AddBlock adds a new block to the blockchain
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func GenesisBlock() *Block { // GenesisBlock creates the first block in the blockchain
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain { // NewBlockchain creates a new blockchain with the genesis block
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func main() { // Main function to run the blockchain
	bc := NewBlockchain()

	bc.AddBlock("Send 1 MBX to Edward")
	bc.AddBlock("Send 3 MBX to Lucia")
	bc.AddBlock("Send 5 MBX to Awn")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()

		pow := NewProofOfWork(block)
		fmt.Printf("PoW Valid: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

// Part 3: Proof of Work

const (
	targetBits = 24
	maxNonce   = math.MaxInt64
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}
	return pow
}

func IntToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
