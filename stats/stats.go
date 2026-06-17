package stats

import (
	"errors"

	"github.com/NorskHelsenett/dnsdist/transport"
)

type Client interface {
	// returns the dnsdist version string.
	Version() (string, error)
	// returns the full dumpStats() output.
	Dump() (string, error)
	// returns the raw getStatisticsCounters() output (Lua table as string).
	Counters() (string, error)
}

type statsClient struct {
	transport transport.Transport
}

func New(transport transport.Transport) Client {
	return &statsClient{
		transport: transport,
	}
}

func (c *statsClient) Version() (string, error) {
	return "", errors.New("un-implemented, accepting PRs on https://github.com/NorskHelsenett/dnsdist")
}
func (c *statsClient) Dump() (string, error) {
	return "", errors.New("un-implemented, accepting PRs on https://github.com/NorskHelsenett/dnsdist")
}
func (c *statsClient) Counters() (string, error) {
	return "", errors.New("un-implemented, accepting PRs on https://github.com/NorskHelsenett/dnsdist")
}
