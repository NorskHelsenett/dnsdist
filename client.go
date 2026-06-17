package dnsdist

import (
	"fmt"
	"strings"

	"github.com/NorskHelsenett/dnsdist/pools"
	"github.com/NorskHelsenett/dnsdist/rules"
	"github.com/NorskHelsenett/dnsdist/servers"
	"github.com/NorskHelsenett/dnsdist/stats"
	"github.com/NorskHelsenett/dnsdist/transport"
)

type Client interface {
	Pools() pools.Client
	Rules() rules.Client
	Servers() servers.Client
	Stats() stats.Client
	Ping() error
	Close() error
}

type dnsDistClient struct {
	pools     pools.Client
	rules     rules.Client
	servers   servers.Client
	stats     stats.Client
	transport transport.Transport
}

func NewClient(transport transport.Transport) Client {
	return &dnsDistClient{
		rules:     rules.New(transport),
		transport: transport,
	}
}

func (c *dnsDistClient) Pools() pools.Client {
	return c.pools
}

func (c *dnsDistClient) Rules() rules.Client {
	return c.rules
}

func (c *dnsDistClient) Servers() servers.Client {
	return c.servers
}

func (c *dnsDistClient) Stats() stats.Client {
	return c.stats
}

func (c *dnsDistClient) Ping() error {
	resp, err := c.transport.Execute("")
	if err != nil {
		return fmt.Errorf("dnsdist ping failed: %w", err)
	}

	if strings.Contains(resp, "Error") {
		return fmt.Errorf("ping failed: got response: %s", resp)
	}

	return nil
}

func (c *dnsDistClient) Close() error {
	return c.transport.Close()
}
