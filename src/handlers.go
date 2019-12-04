package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	helper "./helper"

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

	var blockdata helper.BlockData

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

func RegisterPeer(w http.ResponseWriter, r *http.Request) {
}
