package main

import (
	"context"
	"fmt"

	"github.com/Th0masStorm/vakhtersha/docker"
)

func main() {
	fmt.Println("Starting...")
	cli, err := docker.NewDockerClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	docker.GetChannelFromSocket(cli, ctx)

}
