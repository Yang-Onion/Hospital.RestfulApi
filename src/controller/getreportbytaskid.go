package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetReportByTaskId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskid := vars["id"]
	json.NewEncoder(w).Encode(taskid)
}
