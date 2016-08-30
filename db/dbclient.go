package db

import (
	"fmt"
	"github.com/jmp3833/finance/models"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
	"log"
)

//TODO: Abstract away from local instance
func GetDBInstance() mysql.Conn {
	user := "justin"
	pass := "test123"
	dbname := "test"

	dbinstance := mysql.New("tcp", "", "127.0.0.1:3306", user, pass, dbname)

	err := dbinstance.Connect()
	if err != nil {
		log.Fatal(err)
	}

	return dbinstance
}

func AddRecord(record models.Chase, db mysql.Conn) {
	stmt, err := db.Prepare(
		`insert into chase 
	  (ref, transtype, description, amount, date)
	  values (?, ?, ?, ?, ?)`)

	if err != nil {
		log.Fatal(err)
	}

	//TODO: remove me
	fmt.Printf("%+v\n", record)

	_, err = stmt.Run(
		genid(record),
		record.Transtype,
		record.Description,
		record.Amount,
		record.Date)

	if err != nil {
		log.Fatal(err)
	}
}

func GetRecordById(id int, dbinstance mysql.Conn) []mysql.Row {
	rows, _, err := dbinstance.Query(
		"select * from chase where id = '%s'", id)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func genid(record models.Chase) string {
	return record.Date + record.Description
}
