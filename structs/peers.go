package structs

type Network struct {
	LocalAddress  string `json:"localAddress"`
	RemoteAddress string `json:"remoteAddress"`
}

type EthProtocol struct {
	Difficulty uint64 `json:"difficulty"`
	Head       string `json:"head"`
	Version    string `json:"version"`
}

type Peer struct {
	Caps      []string               `json:"caps"`
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Enode     string                 `json:"enode"`
	Network   Network                `json:"network"`
	Protocols map[string]EthProtocol `json:"protocols"`
	Inbound   bool                   `json:"inbound"`
	Static    bool                   `json:"static"`
	Trusted   bool                   `json:"trusted"`
}
