package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	/*
		"github.com/gorilla/mux"
		"log"
		"net/http"
		"database/sql"*/)

var (
	debug         = flag.Bool("debug", false, "false")
	password      = flag.String("password", "", "P@ssw0rd")
	port     *int = flag.Int("port", 1433, "1433")
	server        = flag.String("server", "", "10.4.69.1")
	user          = flag.String("user", "", "sa")
)

func main() {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	stmt, err := conn.Prepare("select 1, 'abc'")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var somenumber int64
	var somechars string
	err = row.Scan(&somenumber, &somechars)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("somenumber:%d\n", somenumber)
	fmt.Printf("somechars:%s\n", somechars)

	fmt.Printf("bye\n")

	/*db := dbcommon.Mssql{
		DataSource: "10.4.69.1",
		Database:   "Src.Basics",
		Windows:    false,
		Sa: SA{
			User:   "sa",
			Passwd: "P@ssw0rd",
		},

		User:   "sa",
		Passwd: "P@ssw0rd",
	}*/

	/*router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{id}", TodoDetail)

	//体检单位
	router.HandleFunc("/api/physical/team/list", TeamList)

	log.Fatal(http.ListenAndServe(":8080", router))*/

}

/*func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TodoIndex")
}

func TodoDetail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Detail")
}*/
