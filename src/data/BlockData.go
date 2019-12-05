package data

import (
	"encoding/json"

	blockchain "../blockchain"
)

// BlockData ...
type BlockData struct {
	Peer  Peer             `json:"peer"`
	Block blockchain.Block `json:"block"`
}

// DecodeFromJSON ...
func (blockdata *BlockData) DecodeFromJSON(data string) error {
	return json.Unmarshal([]byte(data), blockdata)
}

// EncodeToJSON ...
func (blockdata BlockData) EncodeToJSON() (string, error) {
	data, err := json.MarshalIndent(blockdata, "", "  ")
	return string(data), err
}
