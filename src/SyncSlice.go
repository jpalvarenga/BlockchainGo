package main

import (
	"errors"
	"sync"

	helper "./helper"
)

// SyncBlockDataQueue struct
type SyncBlockDataQueue struct {
	queue []helper.BlockData
	mux   sync.Mutex
}

// Put func
func (q *SyncBlockDataQueue) Put(data helper.BlockData) {
	q.mux.Lock()
	q.queue = append(q.queue, data)
	q.mux.Unlock()
}

// Get func
func (q *SyncBlockDataQueue) Get() (helper.BlockData, error) {
	q.mux.Lock()
	length := len(q.queue)
	if length == 0 {
		return helper.BlockData{}, errors.New("empty queue")
	}
	data := q.queue[length-1]
	// Delete
	q.queue = append(q.queue[:length-1], q.queue[length:]...)
	q.mux.Unlock()
	return data, nil
}

// type SyncIntQueue struct {
// 	queue []int
// 	mux   sync.Mutex
// }

// // Put func
// func (q *SyncIntQueue) Put(data int) {
// 	q.mux.Lock()
// 	q.queue = append(q.queue, data)
// 	q.mux.Unlock()
// }

// // Get func
// func (q *SyncIntQueue) Get() (int, error) {
// 	q.mux.Lock()
// 	length := len(q.queue)
// 	if length == 0 {
// 		return -1, errors.New("empty queue")
// 	}
// 	data := q.queue[length-1]
// 	// Delete
// 	q.queue = append(q.queue[:length-1], q.queue[length:]...)
// 	q.mux.Unlock()
// 	return data, nil
// }
