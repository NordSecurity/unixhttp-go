package unixhttp

import (
	"context"
	"github.com/pkg/errors"
	"net"
	"net/http"
	"os"
	"time"
)

func NewListener(socket string) (net.Listener, error) {
	if fileExists(socket) {
		err := os.Remove(socket)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	listener, err := net.Listen("unix", socket)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = os.Chmod(socket, 0777)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return listener, nil
}

func NewClient(socket string) *http.Client {
	transport := &http.Transport{
		DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
			return net.Dial("unix", socket)
		},
	}

	return &http.Client{
		Transport: transport,
		Timeout:   time.Minute,
	}
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}
