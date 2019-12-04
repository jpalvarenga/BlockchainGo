package helper

// Peer represents a node in the network
type Peer struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

// Peers represents a Peer list
type Peers []Peer
