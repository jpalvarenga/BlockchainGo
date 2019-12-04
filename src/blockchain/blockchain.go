/*
Package miner created by Joao Alvarenga

Title: Blockchain Data Structure

Description: Simple implementation of a blockchain data structure in Golang

*/
package miner

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Blockchain Struct
// Description: Represents a blockchain
type Blockchain struct {
	Chain  map[int32][]Block
	Length int32
}

// Initial Function
// Description: This function forms the blockchain data structure.
func (blockchain *Blockchain) Initial() {
	blockchain.Chain = make(map[int32][]Block)
	blockchain.Length = 0
}

// Insert Function
// Description: This function takes a block as the argument, use its height to find the corresponding list in blockchain's Chain map. If the list has already contained that block's hash, ignore it because we don't store duplicate blocks; if not, insert the block into the list.
// Argument: block
func (blockchain *Blockchain) Insert(block Block) error {

	height := block.Header.Height
	forks := blockchain.Chain[height]

	// Check if block already exists in this level
	for _, fork := range forks {
		if fork.Header.Hash == block.Header.Hash {
			return errors.New("block already exists")
		}
	}

	// Add block to Blockchain
	blockchain.Chain[height] = append(blockchain.Chain[height], block)

	// Increase blockchain height
	blockchain.Length = int32(len(blockchain.Chain))

	// No errors, Yay!
	return nil
}

// GetLatestBlocks Function
// Description: This function takes a height as the argument, returns the list of blocks stored in that height or None if the height doesn't exist.
// Argument: int32
// Return type: []Block
func (blockchain Blockchain) GetLatestBlocks(height int32) ([]Block, error) {
	if blockchain.Length >= height {
		return blockchain.Chain[height], nil
	}
	err := errors.New("empty height level in blockchain")
	return nil, err
}

// GetParentBlock Function
// Description: This function takes a height as the argument, returns the list of blocks stored in that height or None if the height doesn't exist.
// Argument: int32
// Return type: []Block
func (blockchain Blockchain) GetParentBlock(block Block) (*Block, error) {
	parentHeight := block.Header.Height - 1
	for _, b := range blockchain.Chain[parentHeight] {
		if block.Header.ParentHash == b.Header.Hash {
			return &b, nil
		}
	}
	return nil, errors.New("parent not found")
}

// EncodeToJSON Function
// Description: This function iterates over all the blocks, generate blocks' JSON String by the function you implemented for a single Block, and return the list of those JSON Strings.
// Return type: string.
func (blockchain Blockchain) EncodeToJSON() (string, error) {

	// Create list
	chain := make([]Block, 0)

	// Add all blocks to list
	for index := 0; index < int(blockchain.Length); index++ {
		chain = append(chain, blockchain.Chain[int32(index)][0])
		fmt.Println("hello")
	}

	//Encode list to JSON
	var data []byte
	data, err := json.MarshalIndent(chain, "", "  ")
	return string(data), err
}

// DecodeFromJSON Function
// Description: This function is called upon a blockchain instance. It takes a blockchain JSON string as input, decodes the JSON string back to a list of block JSON strings, decodes each block JSON string back to a block instance, and inserts every block into the blockchain.
// Argument: self, string
func (blockchain *Blockchain) DecodeFromJSON(data string) (*Blockchain, error) {

	// Turn JSON into list
	chain := new([]Block)
	err := json.Unmarshal([]byte(data), chain)

	// Go over every block in the list and insert it to the blockchain
	for _, block := range *chain {
		blockchain.Insert(block)
	}

	return blockchain, err
}
