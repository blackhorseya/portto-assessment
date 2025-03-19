package netx

import (
	"net"
	"strconv"
	"testing"
)

func TestGetAvailablePort(t *testing.T) {
	// Test that the port is not 0
	port := GetAvailablePort()
	if port == 0 {
		t.Error("GetAvailablePort() returned 0")
	}

	// Test that the port is within the range
	if port < _startPort || port > _endPort {
		t.Errorf("GetAvailablePort() returned %d, which is out of range", port)
	}

	// Test that the port is available
	address := ":" + strconv.Itoa(port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		t.Errorf("GetAvailablePort() returned %d, which is not available", port)
	} else {
		listener.Close()
	}
}
