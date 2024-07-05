package ui

import (
	"log"
	"net/http"
	"time"

	"custom-geth-exporter/internal/validation"
	"custom-geth-exporter/metrics"
	"custom-geth-exporter/structs"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeUI(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/public/index.html")
}

func ServeRPCData(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		peers := getPeerData()
		log.Printf("Sending peer data: %v", peers)

		validPeers := validatePeerData(peers)
		if err := conn.WriteJSON(validPeers); err != nil {
			log.Println(err)
			return
		}
		time.Sleep(5 * time.Second)
	}
}

func getPeerData() []structs.Peer {
	var peers []structs.Peer
	err := metrics.RPCClient.Call(&peers, "admin_peers")
	if err != nil {
		log.Println(err)
		return nil
	}
	return peers
}

func validatePeerData(peers []structs.Peer) []structs.Peer {
	var validPeers []structs.Peer
	for _, peer := range peers {
		if validation.ValidatePeer(peer) {
			validPeers = append(validPeers, peer)
		} else {
			log.Printf("Validation failed for peer: %v", peer)
		}
	}
	return validPeers
}
