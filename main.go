package main

import (
	cmd "github.com/moohbr/WebMonitor/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()

	if err != nil {
		panic(err)
	}

}
