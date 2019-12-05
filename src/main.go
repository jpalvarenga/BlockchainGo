package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	blockchain "./blockchain"
	data "./data"
)

// The current blockchain
var bc = blockchain.SyncBlockchain{
	BC: blockchain.Blockchain{
		Chain:  make(map[int32][]blockchain.Block),
		Length: 0,
	},
}

var peers = data.Peers{}

func main() {

	blockchainJSON, _ := ioutil.ReadFile("./json/blockchain.json")
	peersJSON, _ := ioutil.ReadFile("./json/peers.json")

	bc.DecodeFromJSON(string(blockchainJSON))
	error := json.Unmarshal(peersJSON, &peers)

	if error != nil {
		fmt.Println(error)
	}

	go func() {
		router := InitRouter()
		log.Fatal(http.ListenAndServe(":8080", router))
	}()

	router2 := InitRouter()
	log.Fatal(http.ListenAndServe(":3030", router2))
}
