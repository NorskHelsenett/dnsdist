package servers

import (
	"errors"

	"github.com/NorskHelsenett/dnsdist/transport"
)

type Client interface {
	// returns the raw showServers() output.
	List() (string, error)
	// adds a new backend. addr is "ip:port".
	Add(addr string) error
	// removes a backend by index, UUID or name.
	Remove(id string) error
	// administratively marks a backend as up.
	SetUp(id string) error
	// administratively marks a backend as down.
	SetDown(id string) error
}

type serversClient struct {
	transport transport.Transport
}

func New(transport transport.Transport) Client {
	return &serversClient{
		transport: transport,
	}
}

func (c *serversClient) List() (string, error) {
	return "", errors.New("un-implemented, accepting PRs on https://github.com/NorskHelsenett/dnsdist")
}
func (c *serversClient) Add(addr string) error {
	return errors.New("un-implemented, accepting PRs on https://github.com/NorskHelsenett/dnsdist")
}
func (c *serversClient) Remove(id string) error {
	return errors.New("un-implemented, accepting PRs on https://github.com/NorskHelsenett/dnsdist")
}
func (c *serversClient) SetUp(id string) error {
	return errors.New("un-implemented, accepting PRs on https://github.com/NorskHelsenett/dnsdist")
}
func (c *serversClient) SetDown(id string) error {
	return errors.New("un-implemented, accepting PRs on https://github.com/NorskHelsenett/dnsdist")
}
