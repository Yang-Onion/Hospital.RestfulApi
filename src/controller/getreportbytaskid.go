package controller

import (
	"../github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetReportByTaskId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskid := vars["id"]
	fmt.Println(taskid)
	json.NewEncoder(w).Encode("abc")
}
