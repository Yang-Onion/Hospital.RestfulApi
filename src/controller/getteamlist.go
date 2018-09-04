package controller

import (
	"encoding/json"
	"net/http"
)

func GetTeamList(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("teamlist")
}
