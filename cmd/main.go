package main

import (
	"fmt"

	"github.com/Th0masStorm/vakhtersha/docker"
)

func main() {
	fmt.Println("Testing...")
	_, err := docker.NewDockerClient()
	if err != nil {
		panic(err)
	}
}
