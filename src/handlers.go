package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	data "./data"
	helpers "./helpers"

	"github.com/gorilla/mux"
)

// Start function starts the server
func Start(w http.ResponseWriter, r *http.Request) {
}

// UploadBlockchain function returns the entire blockchain
func UploadBlockchain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(blockchain); err != nil {
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
		if blocks, err := blockchain.GetLatestBlocks(int32(h)); err == nil {
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

	var blockdata data.BlockData

	// Pass read body
	var e error
	if body, e := ioutil.ReadAll(r.Body); e == nil {
		// Pass JSON decode
		if e := blockdata.DecodeFromJSON(string(body)); e == nil {
			// add it to queue
			queue.Put(blockdata)
			//fmt.Println(queue.queue, "Im here 3")
			w.WriteHeader(http.StatusOK)
		}
	}

	// If error exists show bad request
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// RegisterPeer func
func RegisterPeer(w http.ResponseWriter, r *http.Request) {
	// find ip adress and port
	if ip, port, er := helpers.ParseRemoteAddress(r.RemoteAddr); er == nil {
		// append new peer to peer list
		peers = append(peers, data.Peer{
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
}
