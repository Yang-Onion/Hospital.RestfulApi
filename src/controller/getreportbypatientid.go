package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetReportByPatientId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bingRenId := vars["bingRenId"]
	json.NewEncoder(w).Encode(bingRenId)
}
