package pools

import (
	"errors"

	"github.com/NorskHelsenett/dnsdist/transport"
)

type Client interface {
	// returns the raw showPools() output.
	List() (string, error)
}

type poolsClient struct {
	transport transport.Transport
}

func New(transport transport.Transport) Client {
	return &poolsClient{
		transport: transport,
	}
}

func (c *poolsClient) List() (string, error) {
	return "", errors.New("un-implemented, accepting PRs on https://github.com/NorskHelsenett/dnsdist")
}
