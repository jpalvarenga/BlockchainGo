package main

import "fmt"

func main() {

	// Create blockchain
	blockchain := new(Blockchain)
	blockchain.Initial()

	// Create genesis block
	genesis := new(Block)
	genesis.Initial(0, "genesis", "genesis")

	// Insert genesis block to the blockchain
	blockchain.Insert(*genesis)

	// Create and insert first block
	block1 := new(Block)
	block1.Initial(1, blockchain.Chain[blockchain.Length-1][0].Header.Hash, "1BTC")

	blockchain.Insert(*block1)

	// Create and insert second block
	block2 := new(Block)
	block2.Initial(2, blockchain.Chain[blockchain.Length-1][0].Header.Hash, "2BTC")

	blockchain.Insert(*block2)

	// Get Json Representation of blockchain
	json, _ := blockchain.EncodeToJSON()

	// Print Blockchain 1
	fmt.Println()
	fmt.Println("BLOCKCHAIN 1 JSON")
	fmt.Println()
	fmt.Println(json)

	// Create second blockchain
	blockchain2 := new(Blockchain)
	blockchain2.Initial()

	// Insert all blocks from JSON to blockchain
	blockchain2.DecodeFromJSON(json)

	// Get Json Representation of blockchain 2
	json2, _ := blockchain.EncodeToJSON()

	// Print Blockchain 2
	fmt.Println()
	fmt.Println("BLOCKCHAIN 2 JSON")
	fmt.Println()
	fmt.Println(json2)
}
