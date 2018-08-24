package dbcommon

import (
	"database/sql"
	_ "github.com/mattn/go-adodb"
	"strings"
)

type Mssql struct {
	*sql.DB
	dataSource string
	database   string
	windows    bool
	sa         SA
}

type SA struct {
	user   string
	passwd string
}

func (m *Mssql) Open() (err error) {
	var conf []string
	conf = append(conf, "Provider=SQLOLEDB")
	conf = append(conf, "Data Source="+m.dataSource)
	conf = append(conf, "Initial Catalog="+m.database)
	conf = append(conf, "user id="+m.sa.user)
	conf = append(conf, "password="+m.sa.passwd)

	m.DB, err = sql.Open("driverName", strings.Join(conf, ";"))

	if err != nil {
		return err
	}
	return nil

}
