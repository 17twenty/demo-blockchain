package main

import (
	"fmt"
	"testing"

	"github.com/satori/go.uuid"

	"github.com/shopspring/decimal"
)

func TestBlockchain(t *testing.T) {
	bc := NewBlockchain()

	nick := uuid.NewV4()
	fred := uuid.NewV4()
	bob := uuid.NewV4()

	bc.AddBlock(Transaction{
		Source:      nick,
		Destination: fred,
		Amount:      decimal.NewFromFloat(1000),
	})

	bc.AddBlock(Transaction{
		Source:      fred,
		Destination: bob,
		Amount:      decimal.NewFromFloat(200),
	})

	for _, block := range bc.blocks {
		fmt.Printf("Previous hash: %x\n", block.previousHash)
		fmt.Printf("Transaction: %s\n", block.transaction)
		fmt.Printf("Hash: %x\n", block.hash)
		fmt.Println()
	}
}
