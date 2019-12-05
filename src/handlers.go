package main

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	blockchain "./blockchain"
	data "./data"
	helpers "./helpers"
	miner "./miner"

	"github.com/gorilla/mux"
)

// Start function starts the server
func Start(w http.ResponseWriter, r *http.Request) {
	block := blockchain.Block{
		Header: blockchain.Header{
			Height:     1,
			Timestamp:  1575171860,
			Hash:       "7d1922155e6724f00c10fb3e5aa836480a1593136e90c635424b4527e862b082",
			ParentHash: "201176e2b58a005e3925315011e40a3cdf49afff5a0459d2f9db92ab79e084f0",
			Size:       32,
			Difficulty: 6,
			Nonce:      "72879827ab97977d",
		},
		Value: "1BTC",
	}
	miner.RequestParentBlock(block, data.Peer{
		ID:   "localhost:8080",
		IP:   "localhost",
		Port: "8080",
	})

	http.Get("http://localhost:8080/upload")

}

// UploadBlockchain function returns the entire blockchain
func UploadBlockchain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(bc); err != nil {
		panic(err)
	}
}

// UploadBlock function returns a block
func UploadBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	height := mux.Vars(r)["height"]
	hash := mux.Vars(r)["hash"]

	// Tries to convert height to a number
	if h, err := strconv.Atoi(height); err == nil {
		if blocks, err := bc.GetLatestBlocks(int32(h)); err == nil {
			for _, block := range blocks {
				if block.Header.Hash == hash {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(block)
				}
			}
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// ReceiveBlock func
func ReceiveBlock(w http.ResponseWriter, r *http.Request) {

	blockdata := new(data.BlockData)

	if body, er := ioutil.ReadAll(r.Body); er == nil {
		if er = json.Unmarshal(body, blockdata); er == nil {

			var block = blockdata.Block
			//var parent *blockchain.Block = nil

			// loop until find parent in the blockchain
			if bc.GetParentBlock(block) == nil {
				miner.RequestParentBlock(block, blockdata.Peer)
			}

			if block.Header.Nonce == "" {
				nonce, _ := miner.Pow(block)
				block.Header.Nonce = hex.EncodeToString(nonce[:])
			}
			if miner.CheckNonce(block.Header.Nonce, block) {
				miner.Broadcast(block, peers)
			}
		}
	}

	// Pass read body
	var e error
	if body, e := ioutil.ReadAll(r.Body); e == nil {
		// Pass JSON decode
		if e := blockdata.DecodeFromJSON(string(body)); e == nil {
			//fmt.Println(queue.queue, "Im here 3")
			w.WriteHeader(http.StatusOK)
		}
	}

	// If error exists show bad request
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// RegisterPeer function
// Endpoint: /peer
// Purpose: A node can register itself as a validator by calling the /peer get
// request on any node in the network
func RegisterPeer(w http.ResponseWriter, r *http.Request) {
	// find ip adress and port
	if ip, port, er := helpers.ParseRemoteAddress(r.RemoteAddr); er == nil {
		// append new peer to peer list
		peers = append(peers, data.Peer{
			ID:   string(ip + ":" + port),
			IP:   ip,
			Port: port,
		})
		// good request
		w.WriteHeader(http.StatusOK)
		// return id
		w.Write([]byte(string(ip + ":" + port)))
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	// TODO: Broadcast Peer
}
