package data

import (
	"errors"
	"sync"
)

// SyncBlockDataQueue struct
type SyncBlockDataQueue struct {
	queue []BlockData
	mux   sync.Mutex
}

// Put func
func (q *SyncBlockDataQueue) Put(data BlockData) {
	q.mux.Lock()
	q.queue = append(q.queue, data)
	q.mux.Unlock()
}

// Get func
func (q *SyncBlockDataQueue) Get() (BlockData, error) {
	q.mux.Lock()
	length := len(q.queue)
	if length == 0 {
		return BlockData{}, errors.New("empty queue")
	}
	data := q.queue[length-1]
	// Delete
	q.queue = append(q.queue[:length-1], q.queue[length:]...)
	q.mux.Unlock()
	return data, nil
}
