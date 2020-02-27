package api

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
league - endpoint to get league table
*/
func getLeague(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
	board, err := db.GetLeaderBoard()
	if err != nil {
		log.Fatal(err)
	}

	js, err := json.Marshal(board)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Write(js)
}

/*
getResults - Endpoint for league game results
*/
func getResults(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	results, err := db.GetResults()
	if err != nil {
		log.Fatal(err)
	}

	js, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Write(js)
}

/*
getCalendar - Endpoint for future games
*/
func getCalendar(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	results, err := db.GetCalendar()
	if err != nil {
		log.Fatal(err)
	}

	js, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Write(js)
}
