package metrics

import (
	"log"
	"os"
	"time"

	"custom-geth-exporter/internal/validation"
	"custom-geth-exporter/structs"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

type RPCClientInterface interface {
	Call(result interface{}, method string, args ...interface{}) error
}

var (
	Client     *ethclient.Client
	RPCClient  RPCClientInterface
	rpcDial    = func(rawurl string) (RPCClientInterface, error) { return rpc.Dial(rawurl) }
	fileExists = func(path string) bool { _, err := os.Stat(path); return err == nil }

	peerCountGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ethereum_peer_count",
		Help: "Ethereum protocol version used by peers",
	}, []string{"peer_id", "version", "peer_name", "local_address", "remote_address", "enode"})
)

func Init(ipcPath, httpURL string, httpFallback bool) error {
	var err error
	if ipcPath != "" {
		for {
			if fileExists(ipcPath) {
				log.Printf("IPC path %s is now available", ipcPath)
				break
			}
			log.Printf("Waiting for IPC path %s to become available...", ipcPath)
			time.Sleep(1 * time.Second)
		}

		var rpcClient RPCClientInterface
		rpcClient, err = rpcDial(ipcPath)
		if err != nil {
			log.Printf("error: Failed to connect to IPC at %s: %v", ipcPath, err)
			if !httpFallback {
				return err
			}
		} else {
			RPCClient = rpcClient
			if client, ok := rpcClient.(*rpc.Client); ok {
				Client = ethclient.NewClient(client)
			}
			log.Printf("Successfully connected to IPC at %s", ipcPath)
			return nil
		}
	}

	if httpFallback || ipcPath == "" {
		var rpcClient RPCClientInterface
		rpcClient, err = rpcDial(httpURL)
		if err != nil {
			log.Printf("error: Failed to connect to HTTP at %s: %v", httpURL, err)
			return err
		} else {
			RPCClient = rpcClient
			if client, ok := rpcClient.(*rpc.Client); ok {
				Client = ethclient.NewClient(client)
			}
		}
	}

	return nil
}

func UpdatePeerMetrics() error {
	var peers []structs.Peer
	err := RPCClient.Call(&peers, "admin_peers")
	if err != nil {
		log.Printf("RPC Call Error: %v", err)
		return err
	}

	for _, peer := range peers {
		if !validation.ValidatePeer(peer) {
			continue
		}

		id := peer.ID
		name := peer.Name
		localAddress := peer.Network.LocalAddress
		remoteAddress := peer.Network.RemoteAddress
		version := string(peer.Protocols["eth"].Version)
		enode := peer.Enode

		peerCountGauge.WithLabelValues(id, version, name, localAddress, remoteAddress, enode).Set(float64(len(peers)))

	}
	return nil
}

func init() {
	prometheus.MustRegister(peerCountGauge)
}
