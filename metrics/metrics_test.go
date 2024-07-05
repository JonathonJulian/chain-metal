package metrics

import (
	"custom-geth-exporter/structs"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestInit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRPCClient := NewMockRPCClientInterface(ctrl)
	rpcDial = func(rawurl string) (RPCClientInterface, error) {
		return mockRPCClient, nil
	}
	fileExists = func(path string) bool {
		return true
	}

	err := Init("/tmp/geth.ipc", "http://localhost:8545", true)
	if err != nil {
		t.Fatalf("Init() error = %v", err)
	}
}

func TestUpdatePeerMetrics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRPCClient := NewMockRPCClientInterface(ctrl)
	RPCClient = mockRPCClient

	peers := []structs.Peer{
		{ID: "peer1", Name: "peer1", Enode: "enode1", Network: structs.Network{LocalAddress: "local1", RemoteAddress: "remote1"}, Protocols: map[string]structs.EthProtocol{"eth": {Version: "1"}}},
		{ID: "peer2", Name: "peer2", Enode: "enode2", Network: structs.Network{LocalAddress: "local2", RemoteAddress: "remote2"}, Protocols: map[string]structs.EthProtocol{"eth": {Version: "1"}}},
	}

	mockRPCClient.EXPECT().Call(gomock.Any(), "admin_peers").SetArg(0, peers).Return(nil)

	err := UpdatePeerMetrics()
	if err != nil {
		t.Fatalf("UpdatePeerMetrics() error = %v", err)
	}

	collected := testutil.CollectAndCount(peerCountGauge)
	if collected != len(peers) {
		t.Errorf("collected %v metrics, want %v", collected, len(peers))
	}
}
