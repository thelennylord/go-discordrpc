package discordrpc

import (
	"encoding/json"
	"fmt"

	"github.com/thelennylord/go-discordrpc/ipc"
)

// Client wrapper for the Discord RPC client
type Client struct {
	ClientID string
	Socket   *ipc.Socket
}

// New sends a handshake in the socket and returns an error or nil and an instance of Client
func New(clientid string) (*Client, error) {
	if clientid == "" {
		return nil, fmt.Errorf("no clientid set")
	}

	payload, err := json.Marshal(handshake{"1", clientid})
	if err != nil {
		return nil, err
	}

	sock, err := ipc.NewConnection()
	if err != nil {
		return nil, err
	}

	c := &Client{Socket: sock, ClientID: clientid}

	r, err := c.Socket.Send(0, string(payload))
	if err != nil {
		return nil, err
	}

	var responseBody Data
	if err := json.Unmarshal([]byte(r), &responseBody); err != nil {
		return nil, err
	}

	if responseBody.Code > 1000 {
		return nil, fmt.Errorf(responseBody.Message)
	}

	return c, nil
}
