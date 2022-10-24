package data

import "time"

type Server struct {
	Name                string
	IP                  string
	URL                 string
	AvarageResponseTime time.Duration
	LastUpdate          string
	LastCheck           string
	LastStatus          int
	Monitor             bool
}
