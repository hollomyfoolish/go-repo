package main

import (
	"context"
	"fmt"

	// "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	// "go.module.checker/pkg/mod/github.com/docker/docker@v1.13.1/api/types"
	// "go.module.checker/pkg/mod/github.com/docker/api/types"
	// "golang.org/x/net/context"
	// "github.com/docker/docker/api/types/filters"
	// "github.com/docker/docker/api/types/network"
)

const DOCKER_REGISTRY = "docker.wdf.sap.corp:51271"
const SAP_SME_NETWORK = "sap-sme"

var dockerClient *client.Client
var dockRegistry string

// func init() {
// 	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println("can get docker client")
// 		return
// 	}
// 	dockerClient = cli

// 	dockRegistry = os.Getenv("DOCKER_REGISTRY")
// 	if len(dockRegistry) == 0 {
// 		dockRegistry = DOCKER_REGISTRY
// 	}
// }

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}
