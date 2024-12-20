package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddress := "localhost:4000"

	transport := NewTCPTransport(listenAddress)

	assert.Equal(t, transport.listenAddress, listenAddress)
}
