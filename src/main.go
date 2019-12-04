package main

import (
	"fmt"
	"log"
	"net/http"

	miner "./blockchain"
	helper "./helper"
)

// The current blockchain
var blockchain = miner.Blockchain{Chain: make(map[int32][]miner.Block), Length: 0}

// A queue holding the blocks needed to be processed
var queue = SyncBlockDataQueue{queue: []helper.BlockData{}}

func main() {

	go func() {
		i := 0
		for {
			if len(queue.queue) > 0 {
				fmt.Println(queue.Get())
			}
			if i%1400234240 == 0 {
				fmt.Println("Im alive")
				fmt.Println(queue.queue)
			}
			i++
		}
	}()

	// Create queue that will hold all blocks yet to process

	// json, _ := ioutil.ReadFile("./json/blockchain.json")

	// blockchain.DecodeFromJSON(string(json))

	// value, _ := blockchain.EncodeToJSON()
	// fmt.Println(value)
	// fmt.Println(blockchain.Length)

	// if err == nil {
	// 	blockchain.DecodeFromJSON(string(json))
	// } else {
	// 	// request blockchain from users
	// }

	// f, _ := os.Create("./json/blockchain.json")

	// blockchain2 := new(miner.Blockchain)
	// blockchain2.Initial()

	// // Create genesis block
	// genesis := new(miner.Block)
	// genesis.Initial(0, "genesis", "genesis")

	// // Insert genesis block to the blockchain
	// blockchain2.Insert(*genesis)

	// // Create and insert first block
	// block1 := new(miner.Block)
	// block1.Initial(1, blockchain.Chain[blockchain.Length-1][0].Header.Hash, "1BTC")

	// blockchain.Insert(*block1)

	// // Create and insert second block
	// block2 := new(miner.Block)
	// block2.Initial(2, blockchain.Chain[blockchain.Length-1][0].Header.Hash, "2BTC")

	// blockchain.Insert(*block2)

	// a, _ := blockchain.EncodeToJSON()
	// f.WriteString(a)

	// _ = f.Close()

	router := InitRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
