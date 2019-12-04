/*
Package miner created by Joao Alvarenga

Title: Block Data Structure

Description: Simple implementation of a block structure in blockchain

*/
package miner

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"
)

// DIFFICULTY represents difficulty to harvest block
const DIFFICULTY int32 = 6

// Block Struct
// Description: Represents a block in the blockchain
type Block struct {
	Header Header `json:"header"`
	Value  string `json:"value"`
}

// Header Struct
// Description: Represents header in a Block object
type Header struct {
	Height     int32  `json:"height"`
	Timestamp  int64  `json:"timestamp"`
	Hash       string `json:"hash"`
	ParentHash string `json:"parenthash"`
	Size       int32  `json:"size"`
	Difficulty int32  `json:"difficulty"`
	Nonce      string `json:"nonce"`
}

// Initial Function
// Description: This function takes arguments(such as height, parentHash, and value) and forms a block. This is a method of the block struct.
// Argument: Height of the block, it's parent hash and the value ( transactions )
func (block *Block) Initial(height int32, parentHash string, value string) {
	block.Header = Header{
		Height:     height,
		Timestamp:  time.Now().Unix(),
		ParentHash: parentHash,
		Size:       32,
		Difficulty: DIFFICULTY,
	}
	block.Value = value
	// Proof of work ...
	nonce, _ := Pow(*block)
	block.Header.Nonce = hex.EncodeToString(nonce[:])
	// Getting hash value
	hash := sha256.Sum256(bytes.Join([][]byte{
		[]byte(strconv.Itoa(int(block.Header.Height))),
		[]byte(strconv.Itoa(int(block.Header.Timestamp))),
		[]byte(block.Header.ParentHash),
		[]byte(strconv.Itoa(int(block.Header.Size))),
		[]byte(block.Value)},
		[]byte{}))
	block.Header.Hash = hex.EncodeToString(hash[:])
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
