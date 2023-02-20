package chain

import (
	"fmt"
	"time"
)

type Node struct {
	ID    string
	State map[string]interface{}
}
type Block struct {
	ID       int64
	NodeList []*Node
	Hash     string
}
type Blockchain struct {
	Blocks []*Block
}

func (b *Blockchain) AddBlock(block *Block) {
	if b.isValidBlock(block) {
		b.Blocks = append(b.Blocks, block)
	} else {
		fmt.Println("Invalid block")
	}
}
func (b *Blockchain) isValidBlock(block *Block) bool {
	// Implement your own block validation logic here
	return true
}

// PBFT
type PBFT struct {
	Blockchain *Blockchain
	Nodes      []*Node
	F          int
}

func (p *PBFT) ProposeBlock(block *Block) {
	// Broadcast the block to all nodes
	for _, node := range p.Nodes {
		go func(node *Node) {
			node.SendBlock(block)
		}(node)
	}

	// Wait for votes
	votes := make(chan bool, len(p.Nodes))
	for i := 0; i < len(p.Nodes)-p.F; i++ {
		select {
		case vote := <-votes:
			if !vote {
				fmt.Println("Block was not accepted by a node")
				return
			}
		case <-time.After(time.Second * 10):
			fmt.Println("Timeout reached, not enough votes received")
			return
		}
	}

	// Add block to the blockchain
	p.Blockchain.AddBlock(block)
	fmt.Println("Block added to the blockchain")
}

type Transaction struct {
	Data interface{}
}

func (p *PBFT) AddTransaction(transaction *Transaction) {
	p.TransactionBuffer = append(p.TransactionBuffer, transaction)
	if len(p.TransactionBuffer) >= p.TransactionThreshold {
		p.CreateAndProposeBlock()
	}
}

func (p *PBFT) CreateAndProposeBlock() {
	// Create new block
	block := &Block{
		ID:       p.NextBlockID,
		NodeList: p.Nodes,
		Hash:     "",
		Data:     p.TransactionBuffer,
	}
	// Hash the block
	block.Hash = HashBlock(block)

	// Clear the transaction buffer
	p.TransactionBuffer = nil

	// Propose the block
	p.ProposeBlock(block)

	// Increment the next block ID
	p.NextBlockID++
}
