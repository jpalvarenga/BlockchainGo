package data

// Peer represents a node in the network
type Peer struct {
	ID   string `json:"id"`
	IP   string `json:"ip"`
	Port string `json:"port"`
}

// Peers represents a Peer list
type Peers []Peer
