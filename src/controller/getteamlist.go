package controller

import (
	db "dbcommon"
	"encoding/json"
	"net/http"
)

func GetTeamList(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("teamlist")
}
