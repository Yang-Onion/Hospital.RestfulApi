package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
import (
	_ "github.com/mattn/go-adodb"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{id}", TodoDetail)

	//体检单位
	router.HandleFunc("/api/physical/team/list", TeamList)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TodoIndex")
}

func TodoDetail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Detail")
}
