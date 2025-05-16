package main

	// MoonBucks is a simple blockchain implementation in Go.
	// It demonstrates the basic concepts of a blockchain, including blocks, hashing, and proof of work.
	// It is not intended for production use and should not be used as such.
	// The code is for educational purposes only.

	// Part 1: Block Structure, Hashing, and Creation

type Block struct {				// Block represents a single block in the blockchain
	Timestamp int64				// This is a simplified version of a Block in a Blockchain
	Data      []byte
	PrevBlockHash []byte
	Hash      []byte
}

func (b.Block) BlockHash() {																// BlockHash takes Block fields, concatenates them, and hashes the result
	Timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, Timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string, PrevBlockHash []byte) *Block {									// NewBlock function creates a new block with the given data and previous block hash
	block := &Block{time.Now().Unix(), []byte(data), PrevBlockHash, []byte{}}
	block.BlockHash()
	return block
}

	// Part 2: Implementing the Blockchain

type Blockchain struct {		// My first Blockchain!
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {					// AddBlock adds a new block to the blockchain
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock) 
}

func GenesisBlock() *Block {									// GenesisBlock creates the first block in the blockchain
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {								// NewBlockchain creates a new blockchain with the genesis block
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func main() {													// Main function to run the blockchain
	bc := newBlockchain()

	bc.AddBlock("Send 1 MBX to Edward")
	bc.AddBlock("Send 5 MBX to Awn")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}

