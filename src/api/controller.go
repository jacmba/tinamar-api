package api

import (
	"encoding/json"
	"net/http"
)

func league(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
	board, _ := db.GetLeaderBoard()
	js, _ := json.Marshal(board)

	w.Header().Add("Content-Type", "application/json")
	w.Write(js)
}
