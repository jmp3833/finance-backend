package db

import (
	"fmt"
	"github.com/jmp3833/finance-backend/models"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
	"log"
)

//TODO: Abstract away from local instance
func GetDbInstance() mysql.Conn {
	user := "justin"
	dbname := "finance"
  pass := "test123"

	dbinstance := mysql.New("tcp", "", "127.0.0.1:3306", user, pass, dbname)

	err := dbinstance.Connect()
	if err != nil {
		log.Fatal(err)
	}

	return dbinstance
}

func InsertTransaction(db mysql.Conn, t models.Transaction) {
	preparedStmt := `insert into ` + t.DbName +
		`(ref, transtype, description, amount, date) values (?, ?, ?, ?, ?)`

	stmt, err := db.Prepare(preparedStmt)

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Run(genid(t),
		t.Transtype,
		t.Description,
		t.Amount,
		t.Date)

	if err != nil {
		log.Fatal(err)
	}
}

func genid(t models.Transaction) string {
	fmt.Printf("Generating Id for Transaction:")
	fmt.Printf("%+v\n", t)
	return t.Date + t.Description
}
