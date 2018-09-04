package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CallBack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	json.NewEncoder(w).Encode(id)

}
