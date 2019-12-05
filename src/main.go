package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	miner "./blockchain"
	data "./data"
)

// The current blockchain
var blockchain = miner.Blockchain{Chain: make(map[int32][]miner.Block), Length: 0}

// A queue holding the blocks needed to be processed
var queue = data.SyncBlockDataQueue{}

var peers = data.Peers{}

func main() {

	blockchainJSON, _ := ioutil.ReadFile("./json/blockchain.json")
	peersJSON, _ := ioutil.ReadFile("./json/peers.json")

	blockchain.DecodeFromJSON(string(blockchainJSON))
	error := json.Unmarshal(peersJSON, &peers)

	if error != nil {
		fmt.Println(error)
	}

	router := InitRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
