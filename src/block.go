/*

Title: Block Data Structure
Author: Joao Alvarenga

Description: Simple implementation of a block structure in blockchain

*/
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

// Block Struct
// Description: Represents a block in the blockchain
type Block struct {
	Header Header
	Value  string
}

// Header Struct
// Description: Represents header in a Block object
type Header struct {
	Height     int32
	Timestamp  int64
	Hash       string
	ParentHash string
	Size       int32
}

// Initial Function
// Description: This function takes arguments(such as height, parentHash, and value) and forms a block. This is a method of the block struct.
func (block *Block) Initial(height int32, parentHash string, value string) {
	block.Header = Header{
		Height:     height,
		Timestamp:  time.Now().Unix(),
		ParentHash: parentHash,
		Size:       32,
	}
	block.Value = value
	// Getting hash value
	h := sha256.New()
	h.Write(
		[]byte(string(block.Header.Height) +
			string(block.Header.Timestamp) +
			string(block.Header.ParentHash) +
			string(block.Header.Size) +
			block.Value))
	block.Header.Hash = hex.EncodeToString(h.Sum(nil))
}

// EncodeToJSON Function
// Description: This function encodes a block instance into a JSON format string.
// Argument: a block or you may define this as a method of the block struct
// Return value: a string of JSON format.
func (block Block) EncodeToJSON() (string, error) {
	var data []byte
	data, err := json.MarshalIndent(block, "", "  ")
	return string(data), err
}

// DecodeFromJSON Function
// Description: This function takes a string that represents the JSON value of a block as an input, and decodes the input string back to a block instance.
// Argument: a string of JSON format
// Return value: a block instance
func (block Block) DecodeFromJSON(data string) (Block, error) {
	var b *Block
	err := json.Unmarshal([]byte(data), b)
	return *b, err
}
