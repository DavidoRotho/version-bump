package main

import (
	"log"
	"os"

	"github.com/MichaelDao/version-bump/pkg/config"
	"github.com/urfave/cli"
)

func main() {
	cli := cli.NewApp()
	config.SetConfiguration(cli)
	err := cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}