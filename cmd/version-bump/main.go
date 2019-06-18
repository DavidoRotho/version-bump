package main

import (
	"log"
	"os"

	"github.com/MichaelDao/version-bump/pkg/config"
	"github.com/urfave/cli"
)

func main() {
	cli := cli.NewApp()
		
	config.ParseCLi(cli)
	err := cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// TODO LOCK  

	// TODO getHistory()
	// TODO Calculate()
	// TODO BuildMeta()
	// TODO AddBuild()

	// TODO UNLOCK  

	// TODO final PR stuff
}

