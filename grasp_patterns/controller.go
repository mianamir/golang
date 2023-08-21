package grasp_patterns

import (
	"fmt"
)

/**
In this example, we have a simplified representation of a blockchain-based fintech application.
The Transaction struct represents a financial transaction, and the Block struct represents a
block in the blockchain containing transactions.

The BlockchainController class handles blockchain-related operations, such as adding transactions
to the blockchain. It follows the GRASP Controller pattern, as it acts as an intermediary that
coordinates the interactions between the client code and the blockchain data.

The main function demonstrates the GRASP Controller pattern by creating a BlockchainController
instance and using it to add transactions to the blockchain. The controller manages the blockchain
operations while keeping the client code separate and loosely coupled.

Please note that this example is a simplified illustration of a blockchain application and does
not include the actual hashing and validation mechanisms used in real blockchain systems.

*/

// Transaction represents a financial transaction on the blockchain
type Transaction struct {
	From   string
	To     string
	Amount float64
}

// Block represents a block in the blockchain containing transactions
type Block struct {
	Transactions []Transaction
	PreviousHash string
}

// BlockchainController handles blockchain-related operations
type BlockchainController struct {
	blockchain []Block
}

func NewBlockchainController() *BlockchainController {
	genesisBlock := Block{
		Transactions: []Transaction{},
		PreviousHash: "",
	}
	return &BlockchainController{blockchain: []Block{genesisBlock}}
}

func (bc *BlockchainController) AddTransaction(transaction Transaction) {
	lastBlock := bc.blockchain[len(bc.blockchain)-1]
	newBlock := Block{
		Transactions: append(lastBlock.Transactions, transaction),
		PreviousHash: "hash of last block", // Placeholder for actual hash calculation
	}
	bc.blockchain = append(bc.blockchain, newBlock)
	fmt.Println("Transaction added to the blockchain")
}

func main() {
	blockchainController := NewBlockchainController()

	transaction1 := Transaction{From: "Alice", To: "Bob", Amount: 10.0}
	transaction2 := Transaction{From: "Bob", To: "Charlie", Amount: 5.0}

	blockchainController.AddTransaction(transaction1)
	blockchainController.AddTransaction(transaction2)

	fmt.Println("Blockchain:")
	for _, block := range blockchainController.blockchain {
		fmt.Printf("Block - Previous Hash: %s\n", block.PreviousHash)
		for _, transaction := range block.Transactions {
			fmt.Printf("  - Transaction: From %s to %s, Amount: %.2f\n", transaction.From, transaction.To, transaction.Amount)
		}
	}
}
