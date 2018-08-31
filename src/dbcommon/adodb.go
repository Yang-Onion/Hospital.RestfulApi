package main

import (
	"database/sql"
	"fmt"
	"strings"
)

import (
	_ "github.com/mattn/go-adodb"
)

type Mssql struct {
	*sql.DB
	DataSource string
	Database   string
	Windows    bool
	Sa         SA
}

type SA struct {
	User   string
	Passwd string
}

func (m *Mssql) Open() (err error) {
	var conf []string
	conf = append(conf, "Provider=SQLOLEDB")
	conf = append(conf, "Data Source="+m.DataSource)
	if m.Windows {
		// Integrated Security=SSPI 这个表示以当前WINDOWS系统用户身去登录SQL SERVER服务器(需要在安装sqlserver时候设置)，
		// 如果SQL SERVER服务器不支持这种方式登录时，就会出错。
		conf = append(conf, "integrated security=SSPI")
	}
	conf = append(conf, "Initial Catalog="+m.Database)
	conf = append(conf, "user id="+m.Sa.User)
	conf = append(conf, "password="+m.Sa.Passwd)

	m.DB, err = sql.Open("adodb", strings.Join(conf, ";"))
	if err != nil {
		return err
	}
	return nil
}

func main() {

	db := Mssql{
		DataSource: "10.4.69.1",
		Database:   "Src.Basics",
		// windwos: true 为windows身份验证，false 必须设置sa账号和密码
		Windows: false,
		Sa: SA{
			User:   "sa",
			Passwd: "P@ssw0rd",
		},
	}
	// 连接数据库
	err := db.Open()
	if err != nil {
		fmt.Println("sql open:", err)
		return
	}
	defer db.Close()

	// 执行SQL语句
	rows, err := db.Query("select name,id from basics.location")
	if err != nil {
		fmt.Println("query: ", err)
		return
	}
	news := make(map[int]string)
	for rows.Next() {
		var name string
		var number int
		rows.Scan(&name, &number)
		news[number] = name
		//fmt.Printf("Name: %s \t Number: %d\n", name, number)
	}
	rows.Close()
	for newsid, title := range news {
		fmt.Printf("newsid: %d \t title: %s\n", newsid, title)
	}

}
