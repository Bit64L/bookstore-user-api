package user_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysql_user_username = "mysql_user_username"
	mysql_user_password = "mysql_user_password"
	mysql_user_host     = "mysql_user_host"
	mysql_user_schema   = "mysql_user_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_user_username)
	password = os.Getenv(mysql_user_password)
	host     = os.Getenv(mysql_user_host)
	schema   = os.Getenv(mysql_user_schema)
)

func init() {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema)

	//log.Println(fmt.Sprintf("connect to %s", dataSourceName))

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("connect to mysql successfully!")



}
