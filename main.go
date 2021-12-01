package main

import (
	"net/http"
	"ranking/page"
)

func main() {
	http.HandleFunc("/getrankinglist", page.GetRankingList)
	http.ListenAndServe("localhost:8080", nil)
}
