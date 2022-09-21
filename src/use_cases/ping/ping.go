package ping

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	data "github.com/moohbr/WebMonitor/src/data"
	database "github.com/moohbr/WebMonitor/src/infrastructure/database"
	templates "github.com/moohbr/WebMonitor/src/providers/mail/templates"
	mail "github.com/moohbr/WebMonitor/src/use_cases/sendMail"
)

var (
	wg sync.WaitGroup
)

func PingServer(server data.Server) {
	defer wg.Wait()

	if server.Monitor {
		start := time.Now()
		response, err := http.Get("https://" + server.URL)
		if err != nil {
			log.Println("[ERROR] Error pinging server: " + server.Name)
			log.Println("[ERROR] Error: " + err.Error())
			server.LastStatus = 0
		} else {
			server.LastStatus = response.StatusCode
		}

		server.LastCheck = time.Now().UTC().Format("2006-01-02 15:04:05")
		server.AvarageResponseTime = time.Now().Sub(start)

		log.Println("[SYSTEM] " + server.Name + " | " + server.IP + " | " + server.URL + " | " +
			strconv.FormatInt(server.AvarageResponseTime.Milliseconds(), 10) + "ms | " + server.LastUpdate + " | " + server.LastCheck +
			" | " + strconv.Itoa(server.LastStatus) + " | " + strconv.FormatBool(server.Monitor))
		db := database.OpenDatabase()
		defer db.Close()

		db.UpdateServer(server)
		users := db.GetUsers()

		if len(users) > 0 {
			if server.LastStatus != 200 {
				wg.Add(len(users))
				for _, user := range users {
					go mail.SendMail([]string{user.Email}, templates.ServerDown(server), &wg)
				}
				return
			}
		}
	}
}

func PingAllServers() {
	db := database.OpenDatabase()

	defer db.Close()

	servers := db.GetServers()
	for _, server := range servers {
		PingServer(server)
	}
}
