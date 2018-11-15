package main

import (
	"fmt"
	"omise/challenges/commandline"
	"omise/challenges/executor"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Flags = commandline.Flags
	app.Action = executor.Execute
	app.Usage = "Omise Challenge"
	app.HideVersion = true

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)

	}
}
