package main

import (
	"fmt"

	cmd "github.com/moohbr/WebMonitor/src/cmd"
)

func main() {
	fmt.Println("WebMonitor")

	cmd.RootCmd.Execute()
}
