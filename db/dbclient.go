package db

import (
	"fmt"
	"github.com/jmp3833/finance-backend/models"
	"database/sql"
  _ "github.com/go-sql-driver/mysql"
	"log"
)


//TODO: Abstract away from local instance
func GetDbInstance() *sql.DB {
	db, err := sql.Open("mysql", "justin:test123@/finance")
	if err != nil {
		log.Fatal(err)
	}

	//Incorrect username or password, or some other connection err
	if db.Ping()  != nil {
		log.Fatal(err)
	}
	return db
}

func GetAllTransactions(db sql.DB, bankName string) []models.Transaction {
	var (
		transaction models.Transaction
		transactions []models.Transaction
	)
  //TODO bind to bank type
	preparedStmt := "select * from " + bankName + " LIMIT 100" + `;`
  fmt.Print(preparedStmt)
	rows, err := db.Query(preparedStmt)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err = rows.Scan(
			&transaction.Id,
			&transaction.RefString,
			&transaction.Transtype,
			&transaction.Description,
			&transaction.Amount,
			&transaction.Date,
		)
		transactions = append(transactions, transaction)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	defer rows.Close()
	return transactions
}

func InsertTransaction(db *sql.DB, t models.Transaction) {
	preparedStmt := `insert into ` + t.DbName +
	`(ref, transtype, description, amount, date) values (?, ?, ?, ?, ?)`

	stmt, err := db.Prepare(preparedStmt)

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(genid(t),
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
