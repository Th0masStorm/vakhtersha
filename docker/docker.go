// Package vakhtersha/docker implements reading Docker API
// Package docker implements interaction with Docker API
package docker

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
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

func GetChannelFromSocket(client *client.Client, ctx context.Context) (<-chan events.Message, <-chan error) {
	now := time.Now()
	options := types.EventsOptions{
		Since: now.Format(time.RFC3339),
		Until: now.AddDate(1, 0, 0).Format(time.RFC3339),
		Filters: filters.NewArgs(filters.KeyValuePair{
			Key:   "type",
			Value: "container",
		}),
	}
	events, errors := client.Events(ctx, options)
	return events, errors

}

func ReadEventsFromChannel(events <-chan events.Message) {
	for event := range events {
		fmt.Println(event)
	}
}

func ReadErrorsFromChannel(errors <-chan error) {
	for err := range errors {
		fmt.Println(err)
	}
}
