package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/Th0masStorm/vakhtersha/docker"
)

func main() {
	fmt.Println("Starting...")
	cli, err := docker.NewDockerClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	events, errors := docker.GetChannelFromSocket(cli, ctx)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go docker.ReadEventsFromChannel(events)
	go docker.ReadErrorsFromChannel(errors)
	wg.Wait()

}
