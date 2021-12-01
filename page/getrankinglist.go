package page

import (
	"html/template"
	"net/http"
)

func GetRankingList(w http.ResponseWriter, r *http.Request) {
	var rankingId = r.FormValue("ranking_id")
	parsedTemplate, _ := template.ParseFiles("template/getrankinglist.html")
	parsedTemplate.Execute(w, nil)

	w.Write([]byte(rankingId))
}
