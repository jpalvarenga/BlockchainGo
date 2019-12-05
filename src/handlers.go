package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
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
	// Todo
}

// UploadBlockchain function
// Endpoint: /upload
// Purpose: Get the blockchain from peer
func UploadBlockchain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if json, err := bc.EncodeToJSON(); err == nil {
		w.Write([]byte(json))
	}
}

// UploadBlock function
// Endpoint: /block/{height}/{hash}
// Purpose: Get a specific block from peer
func UploadBlock(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// variables
	height := mux.Vars(r)["height"]
	hash := mux.Vars(r)["hash"]

	// get blocks from certain height with specific hash
	if h, err := strconv.Atoi(height); err == nil {
		if blocks, err := bc.GetLatestBlocks(int32(h)); err == nil {
			for _, block := range blocks {
				if block.Header.Hash == hash {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(block)
				}
			}
		}
	}
}

// ReceiveBlock func
func ReceiveBlock(w http.ResponseWriter, r *http.Request) {
	// blockdata
	blockdata := new(data.BlockData)

	var err error
	if err = ParseRequest(r.Body, blockdata); err == nil {
		// if parent block is not in the blockchain request all the parent blocks to the miner
		if bc.GetParentBlock(blockdata.Block) == nil {
			if err = GetAllParentBlocks(blockdata.Block, blockdata.Peer); err == nil {
				// if all the parent blocks have been added
				if blockdata.Block.Header.Nonce == "" {
					nonce, _ := miner.Pow(blockdata.Block)
					blockdata.Block.Header.Nonce = hex.EncodeToString(nonce[:])
					bc.Insert(blockdata.Block)
					miner.Broadcast(*blockdata, peers)
				} else {
					if miner.CheckNonce(blockdata.Block.Header.Nonce, blockdata.Block) {
						miner.Broadcast(*blockdata, peers)
					}
				}
			}
		}
	}
}

// GetAllParentBlocks func
func GetAllParentBlocks(block blockchain.Block, peer data.Peer) error {
	// parents to be added to current blockchain
	var parents = []blockchain.Block{}
	var b blockchain.Block = block
	var err error
	for bc.GetParentBlock(b) == nil {
		if b, err = miner.RequestParentBlock(block, peer); err == nil {
			parents = append(parents, b)
		} else {
			return err
		}
	}
	for index := len(parents) - 1; index > 0; index-- {
		if miner.CheckNonce(parents[index].Header.Nonce, parents[index]) {
			bc.Insert(parents[index])
		} else {
			return errors.New("invalid nonce")
		}
	}
	return err
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

// ParseRequest func
func ParseRequest(r io.Reader, v interface{}) error {
	var er error
	if body, er := ioutil.ReadAll(r); er == nil {
		er = json.Unmarshal(body, v)
	}
	return er
}
