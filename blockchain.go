package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"

	"github.com/satori/go.uuid"

	"github.com/shopspring/decimal"
)

// blockchain: the world's worst database
type Transaction struct {
	ID          uuid.UUID
	Source      uuid.UUID
	Destination uuid.UUID
	Reference   string
	Amount      decimal.Decimal
	Defined     time.Time
}

type Ledgerer interface {
	getBalance() decimal.Decimal
	commitTransaction(*Transaction)
}

type Block struct {
	timestamp    int64
	transaction  Transaction
	previousHash []byte
	hash         []byte
}

func (t Transaction) String() string {
	return fmt.Sprintf("{ from: %v, to: %v, description: %v, amount: %s", t.Source, t.Destination, t.Reference, t.Amount.StringFixed(4))
}

func (b *Block) CalculateHash() {
	timestamp := []byte(strconv.FormatInt(b.timestamp, 10))
	headers := bytes.Join([][]byte{b.previousHash, []byte(fmt.Sprintf("%v", b.transaction)), timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.hash = hash[:]
}

func NewBlock(tx Transaction, previousHash []byte) *Block {
	block := &Block{
		timestamp:    time.Now().Unix(),
		transaction:  tx,
		previousHash: previousHash,
	}
	block.CalculateHash()
	return block
}

type Blockchain struct {
	blocks []*Block
}

func NewGenesisBlock() *Block {
	return NewBlock(Transaction{
		Reference: "genesis",
	}, []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func (bc *Blockchain) AddBlock(tx Transaction) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(tx, prevBlock.hash)
	bc.blocks = append(bc.blocks, newBlock)
}
