package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)
/*
cannot find package "cloud.google.com/go/civil"
*/
func test() {
	query := url.Values{}
	query.Add("app name", "MyAppName")

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(sa, P@ssw0rd),
		Host:   fmt.Sprintf("%s:%d", 10.4.69.1, 1433),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	conn, err := sql.Open("sqlserver", u.String())
	if err !=nil {
		 log.Fatal("Open connection failed:",err.Error())
	}
	defer conn.Close()

	stmt,err:=conn.Prepare("select name,id from basics.location")
	if err !=nil {
		log.Fatal("Prepare failed:",err.Error())
	}
	defer stmt.Close()

	row :=stmt.QueryRow()
	var name string
	var id int64
	err =row.Scan(&name,&id)
	if err !=nil {
		log.Fatal("Scan failed:",err.Error())
	}
	fmt.Printf("id:%d\n", id)
	fmt.Printf("name:%s\n", name)


}
