/*
Package api has the modules to handle HTTP RESTful API
*/
package api

import (
	"log"
	"net/http"

	"../persistence"
)

/*
Server main data structure to handle HTTP server
*/
type Server struct {
	Port string
	DB   *persistence.Persistence
}

var router = map[string]func(http.ResponseWriter, *http.Request){
	"/league":          getLeague,
	"/league/results":  getResults,
	"/league/calendar": getCalendar,
}
var db *persistence.Persistence

/*
Init initializes HTTP server
*/
func (s *Server) Init() {
	// Initialize routes
	for k, v := range router {
		log.Println("Mounting route", k)
		http.HandleFunc(k, v)
	}

	db = s.DB

	log.Fatal(http.ListenAndServe("0.0.0.0:"+s.Port, nil))
}
