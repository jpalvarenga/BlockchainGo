package miner

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"io/ioutil"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	blockchain "../blockchain"
	data "../data"
)

// RequestParentBlock func
// Requests the Parent block to a peer
func RequestParentBlock(block blockchain.Block, peer data.Peer) (blockchain.Block, error) {
	var url string
	var parent blockchain.Block
	var err error
	// Create url
	url = strings.Join([]string{"http://", peer.ID, "/block/", strconv.Itoa(int(block.Header.Height - 1)), "/", block.Header.ParentHash}, url)
	// http get from peer
	if response, err := http.Get(url); err == nil {
		if body, err := ioutil.ReadAll(response.Body); err == nil {
			if err := parent.DecodeFromJSON(string(body)); err == nil {
				return parent, err
			}
		}
	}
	return parent, err
}

// Broadcast func
// Broadcasts the block to all the peers in the peer list
func Broadcast(blockdata data.BlockData, peers []data.Peer) error {
	var err error
	if json, err := blockdata.EncodeToJSON(); err == nil {
		for _, peer := range peers {
			var url string
			strings.Join([]string{"http://", peer.ID, "/block/receive"}, url)
			http.Post(url, "application/json", bytes.NewBuffer([]byte(json)))
		}
	} else {
		return err
	}
	return err
}

// Pow Function
// Description: Stands for Proof of Work. A block and a difficulty is passed as a parameter and a nonce value is returned
// along with the hash value of sha256(block.Header.parentHash || nonce || block.Value)
// Arguments: block to find nonce and the difficulty
// Return: The nonce as a [8]byte and the hash as a [32] byte
func Pow(block blockchain.Block) ([8]byte, [32]byte) {

	// creates random generador and seed block timestamp
	random := rand.New(rand.NewSource(block.Header.Timestamp))

	// create random nonce integer
	var nonceInteger uint64 = random.Uint64()
	var nonceSlice [8]byte

	// find the max hash value according to the difficulty
	maxInteger := MaxInt(block.Header.Difficulty)

	// loop through all possible Int64 values
	for nonceInteger <= math.MaxUint64 {
		// convert nonce integer to a byte slice for hashing
		binary.BigEndian.PutUint64(nonceSlice[:], nonceInteger)

		// get hash as byte array and integer
		hashSlice := HashConcat(hex.EncodeToString(nonceSlice[:]), block)
		hashInteger := new(big.Int).SetBytes(hashSlice[:])

		// check if the hash integer is smaller than max integer allowed
		if hashInteger.Cmp(maxInteger) == -1 {
			return nonceSlice, hashSlice
		}
		// increase nonce
		nonceInteger = (nonceInteger + 1) % math.MaxInt64
	}

	return [8]byte{}, [32]byte{}
}

// CheckNonce Function
// Description: Function takes a nonce a block and a difficulty and checks if the nonce is indeed valid.
// Arguments: The nonce found, the block and it's difficulty
// Return: If nonce is valid ( boolean )
func CheckNonce(nonce string, block blockchain.Block) bool {
	// gets hash value as a byte slice
	hashSlice := HashConcat(nonce, block)
	// compute integer values
	hashInteger := new(big.Int).SetBytes(hashSlice[:])
	maxInteger := MaxInt(block.Header.Difficulty)
	// compare both values
	return hashInteger.Cmp(maxInteger) == -1
}

// MaxInt Function
// Description: MaxInt will the least integer that converted to hexadecimal will not be a valid hash value
// E.g For difficulty of 4 it will return the Hexadecimal 0x00010000...(x56) as an integer
// Return: Maximum integer as a big.Int pointer
func MaxInt(difficulty int32) *big.Int {
	max := [32]byte{}
	// find correct bracket according to difficulty
	bracket := (difficulty / 2) - 1
	remainder := difficulty % 2
	// change byte slice values manually
	if remainder == 0 {
		max[bracket] = 1
	} else {
		max[bracket+1] = 16
	}
	// return value as an integer
	return new(big.Int).SetBytes(max[:])
}

// HashConcat Function
// Description: Takes a Nonce and a Block and returns sha256(block.Header.parentHash || nonce || block.Value)
// Arguments: The nonce and the block
// Return: The hash of sha256(block.Header.parentHash || nonce || block.Value) as [32] byte
func HashConcat(nonce string, block blockchain.Block) [32]byte {
	return sha256.Sum256(bytes.Join([][]byte{
		[]byte(block.Header.ParentHash),
		[]byte(nonce),
		[]byte(block.Value)},
		[]byte{}))
}
