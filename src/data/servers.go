package servers

import "time"

type Server struct {
	Name                string
	IP                  string
	URL                 string
	AvarageResponseTime time.Duration
	LastUpdate          time.Time
	LastCheck           time.Time
	LastStatus          int
	Monitor             bool
}


