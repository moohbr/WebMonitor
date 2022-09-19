package main

import (
	"fmt"

	data "github.com/moohbr/WebMonitor/src/data"
)

func main() {
	s := data.Server{
		Name: "Test",
		IP:   "someIP",
		URL:  "someURL",
	}
	fmt.Println(s)
}
