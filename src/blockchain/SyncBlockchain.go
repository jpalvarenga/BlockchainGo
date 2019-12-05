package blockchain

import "sync"

// SyncBlockchain is a thread safe blockchain
type SyncBlockchain struct {
	BC  Blockchain
	mux sync.Mutex
}

// Insert func
func (bc *SyncBlockchain) Insert(block Block) error {
	bc.mux.Lock()
	err := bc.BC.Insert(block)
	bc.mux.Unlock()
	return err
}

// GetLatestBlocks func
func (bc *SyncBlockchain) GetLatestBlocks(height int32) ([]Block, error) {
	bc.mux.Lock()
	blocks, err := bc.BC.GetLatestBlocks(height)
	bc.mux.Unlock()
	return blocks, err
}

// GetParentBlock func
func (bc *SyncBlockchain) GetParentBlock(block Block) *Block {
	bc.mux.Lock()
	b := bc.BC.GetParentBlock(block)
	bc.mux.Unlock()
	return b
}

// EncodeToJSON func
func (bc *SyncBlockchain) EncodeToJSON() (string, error) {
	bc.mux.Lock()
	json, err := bc.BC.EncodeToJSON()
	bc.mux.Unlock()
	return json, err
}

// DecodeFromJSON func
func (bc *SyncBlockchain) DecodeFromJSON(data string) error {
	bc.mux.Lock()
	err := bc.BC.DecodeFromJSON(data)
	bc.mux.Unlock()
	return err
}
