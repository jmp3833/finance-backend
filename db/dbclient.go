package db

import (
	"fmt"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
	"log"
)

//TODO: Abstract away from local instance
func GetDB() {
	user := "justin"
	pass := "test123"
	dbname := "test"

	dbinstance := mysql.New("tcp", "", "127.0.0.1:3306", user, pass, dbname)

	err := dbinstance.Connect()
	if err != nil {
		log.Fatal(err)
	}

	rows, res, err := dbinstance.Query("select * from chase where description = ''")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rows)
	fmt.Println(res)
}
