package validation

import (
	"custom-geth-exporter/structs"
	"log"
)

func ValidatePeer(peer structs.Peer) bool {
	if peer.ID == "" || peer.Name == "" || peer.Enode == "" ||
		peer.Network.LocalAddress == "" || peer.Network.RemoteAddress == "" ||
		peer.Protocols["eth"].Version == "" {
		log.Printf("Validation failed for peer: %+v", peer)
		return false
	}
	return true
}
