package db

import (
  "github.com/jmp3833/finance-backend/models"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "time"
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

func GetAllTransactions(db sql.DB, bankName string) ([]models.Transaction, error) {
  query := "select * FROM transactions WHERE bank_name = '" +
    bankName +
    "' LIMIT 100;"
  return parseTransactionsFromDb(db, query)
}

func parseTransactionsFromDb(db sql.DB, query string) ([]models.Transaction, error) {
  var (
    transaction models.Transaction
    transactions []models.Transaction
  )
  rows, err := db.Query(query)
  if err != nil { return nil, err }
  for rows.Next() {
    err = rows.Scan(
      &transaction.Id,
      &transaction.UserId,
      &transaction.AccountId,
      &transaction.BankName,
      &transaction.BankType,
      &transaction.TransactionType,
      &transaction.Amount,
      &transaction.Date,
      &transaction.Description,
    )
    transactions = append(transactions, transaction)
    if err != nil { return nil, err }
  }
  defer rows.Close()
  return transactions, nil
}

func GetTransactionsAfterDate(db sql.DB, bankName string, date time.Time) ([]models.Transaction, error) {
  query := `select * from transactions where date > ` + date.Format(models.DbTimeFormat)
  return parseTransactionsFromDb(db, query)
}

func InsertTransaction(db *sql.DB, t models.Transaction) error {
  preparedStmt := `insert into transactions` +
  `(
    user_id,
    account_id,
    bank_name,
    bank_type,
    transaction_type,
    amount,
    date,
    description
  )` +  `values (?, ?, ?, ?, ?, ?, ?, ?)`

  stmt, err := db.Prepare(preparedStmt)
  if err != nil {
    return err
  }
  _, err = stmt.Exec(
    t.UserId,
    t.AccountId,
    t.BankName,
    t.BankType,
    t.TransactionType,
    t.Amount,
    t.Date,
    t.Description,
  )
  if err != nil { return err }
  return nil
}
