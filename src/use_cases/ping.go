package usecases

import (
	"log"
	"net/http"

	data "github.com/moohbr/WebMonitor/src/data"
	infrastructure "github.com/moohbr/WebMonitor/src/infrastructure/database"
)

func PingServer(server data.Server) data.Server {
	resp, err := http.Get(server.URL)
	if err != nil {
		log.Fatal(err)
	}
	server.LastStatus = resp.StatusCode
	log.Println(server)
	return server
}

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
	log.Println("PingTest - Starting ping test")
	req, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Google:", req.StatusCode)
}
