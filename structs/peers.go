package structs

import (
	"encoding/json"
)

type Version string

func (v *Version) UnmarshalJSON(data []byte) error {
	var temp json.Number
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	*v = Version(temp.String())
	return nil
}

type Network struct {
	LocalAddress  string `json:"localAddress"`
	RemoteAddress string `json:"remoteAddress"`
}

type EthProtocol struct {
	Difficulty uint64  `json:"difficulty"`
	Head       string  `json:"head"`
	Version    Version `json:"version"`
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
