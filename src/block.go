/*

Title: Block Data Structure
Author: Joao Alvarenga

Description: Simple implementation of a block structure in blockchain

*/
package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"strconv"
	"time"
)

// DIFFICULTY represents difficulty to harvest block
const DIFFICULTY int32 = 6

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
	Hash       [32]byte
	ParentHash [32]byte
	Size       int32
	Difficulty int32
	Nonce      [8]byte
}

// Initial Function
// Description: This function takes arguments(such as height, parentHash, and value) and forms a block. This is a method of the block struct.
// Argument: Height of the block, it's parent hash and the value ( transactions )
func (block *Block) Initial(height int32, parentHash [32]byte, value string) {
	block.Header = Header{
		Height:     height,
		Timestamp:  time.Now().Unix(),
		ParentHash: parentHash,
		Size:       32,
		Difficulty: DIFFICULTY,
	}
	block.Value = value
	// Proof of work ...
	block.Header.Nonce, _ = Pow(*block)
	// Getting hash value
	block.Header.Hash = sha256.Sum256(bytes.Join([][]byte{
		[]byte(strconv.Itoa(int(block.Header.Height))),
		[]byte(strconv.Itoa(int(block.Header.Timestamp))),
		block.Header.ParentHash[:],
		[]byte(strconv.Itoa(int(block.Header.Size))),
		[]byte(block.Value)},
		[]byte{}))
}

// EncodeToJSON Function
// Description: This function encodes a block instance into a JSON format string.
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
