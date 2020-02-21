package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func league(w http.ResponseWriter, req *http.Request) {
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
	w.Write(js)
}
