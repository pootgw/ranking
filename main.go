package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/getrankinglist", GetRankingList)
	http.ListenAndServe("localhost:8080", nil)
}

func GetRankingList(httpWriter http.ResponseWriter, httpRequest *http.Request) {
	var rankingId = httpRequest.FormValue("ranking_id")
	httpWriter.Write([]byte(rankingId))
}
