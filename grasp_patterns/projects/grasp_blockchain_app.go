package projects

import (
	"crypto/sha256"
	"fmt"
	"time"
)

/**

In this example, we've created a simple blockchain project in Go that incorporates all the GRASP patterns:

Information Expert: Each Block instance encapsulates information about itself, including its index,
timestamp, data, hash, and previous hash.

Creator: The NewBlockchainController function acts as the creator for the BlockchainController
instances, initializing the blockchain with the genesis block.

Controller: The BlockchainController class acts as the controller, coordinating interactions
between the client code and the blockchain data.

Low Coupling: The BlockchainController handles the interactions between blocks and their data,
promoting loose coupling.

High Cohesion: The Block struct's attributes and methods are closely related to representing and
managing blocks in the blockchain.

Polymorphism: While not explicitly shown in this example, blockchain systems can utilize polymorphism
to handle different types of transactions.

Pure Fabrication: In this simplified example, there's no need for introducing artificial classes.

Indirection: The BlockchainController class indirectly handles the creation and addition of blocks
through its methods.

Protected Variations: The code structure allows for changes to be made to individual blocks without
affecting the overall system.

Responsibility Layer: The separation between the Block struct and the BlockchainController class
creates a clear responsibility layer.

*/

// Block represents a block in the blockchain
type Block struct {
	Index     int
	Timestamp time.Time
	Data      string
	Hash      string
	PrevHash  string
}

// BlockchainController handles blockchain operations
type BlockchainController struct {
	chain []Block
}

// NewBlockchainController creates a new blockchain controller
func NewBlockchainController() *BlockchainController {
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now(),
		Data:      "Genesis Block",
		Hash:      "",
		PrevHash:  "",
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
	return &BlockchainController{chain: []Block{genesisBlock}}
}

// AddBlock adds a new block to the blockchain
func (bc *BlockchainController) AddBlock(data string) {
	prevBlock := bc.chain[len(bc.chain)-1]
	newBlock := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now(),
		Data:      data,
		Hash:      "",
		PrevHash:  prevBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	bc.chain = append(bc.chain, newBlock)
}

func calculateHash(block Block) string {
	data := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp.String(), block.Data, block.PrevHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}

func main() {
	// Information Expert: Each Block instance encapsulates information about itself
	// Creator: Create a new blockchain controller
	blockchainController := NewBlockchainController()

	// Controller & Low Coupling: Use the controller to add blocks to the blockchain
	blockchainController.AddBlock("Transaction 1")
	blockchainController.AddBlock("Transaction 2")
	blockchainController.AddBlock("Transaction 3")

	// Responsibility Layer: The BlockchainController class is responsible for managing the blockchain

	// Display blockchain information
	for _, block := range blockchainController.chain {
		fmt.Printf("Block %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Previous Hash: %s\n", block.PrevHash)
		fmt.Println()
	}
}
