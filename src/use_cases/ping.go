package usecases

import (
	"fmt"
	"log"
	"net/http"

	data "github.com/moohbr/WebMonitor/src/data"
	infrastructure "github.com/moohbr/WebMonitor/src/infrastructure/database"
)

// PingServer pings a server
func PingServer(server data.Server) data.Server {
	resp, err := http.Get(server.URL)
	if err != nil {
		log.Fatal(err)
	}
	server.LastStatus = resp.StatusCode
	fmt.Println(server)
	return server
}

// PingAllServers pings all servers
func PingAllServers() {
	database := infrastructure.Database{}
	defer database.Close()
	database.InitDatabase()

	servers := database.GetServers()
	for _, server := range servers {
		PingServer(server)
	}
}

func PingTest() {
	log.Println("PingTest")
	req, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Google:", req.StatusCode)
}
