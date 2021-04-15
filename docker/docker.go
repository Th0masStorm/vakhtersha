// Package vakhtersha/docker implements reading Docker API
// Package docker implements interaction with Docker API
package docker

import (
	"github.com/docker/docker/client"
)

// NewDockerClient return a Docker API client or an error
// Docker configuration is taken from the environment
// By default, the client will use docker.sock
func NewDockerClient() (*client.Client, error) {
	ClientPtr, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return ClientPtr, nil
}
