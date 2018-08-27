package dbcommon

import (
	"database/sql"
	"flag"
	_ "github.com/denisenkom/go-mssqldb"
	//"strings"
)

type Mssql struct {
	Server   string
	Database string
	User     string
	Passwd   string
	Port     int
}

func (m *Mssql) Open() (err error) {
	var conf []string
	conf = append(conf, "server="+m.Server)
	conf = append(conf, "Initial Catalog="+m.Database)
	conf = append(conf, "user id="+m.User)
	conf = append(conf, "password="+m.Passwd)
	conf = append(conf, "port="+m.Port)

	//m.DB, err = sql.Open("mssql", strings.Join(conf, ";"))
	//m.DB, err = sql.Open("mssql", "Data Source=10.4.69.1;User ID=sa;Password=P@ssw0rd;Initial Catalog=Src.Basics;")
	connString := "Data Source=10.4.69.1;User ID=sa;Password=P@ssw0rd;Initial Catalog=Src.Basics;"
	conn, err := sql.Open("mssql", connString)

	if err != nil {
		return err
	}
	return nil

}

/*
func (m *Mssql) Close() (err error) {


}
*/
