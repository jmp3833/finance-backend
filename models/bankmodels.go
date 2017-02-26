package models

import (
	"math"
	"strconv"
  "time"
)

const csvDateFormat = "01/02/2006"

const (
	BankOfAmerica = "Bank of America"
	Simple        = "Simple inc. Bank"
	Chase         = "Chase Bank"
	NotImplemented = "NA"
)

const (
  Sale = "Sale"
  Payment = "Payment"
  Transfer = "Transfer"
)

const (
  Credit = "Credit"
  Debit = "Debit"
)

//TODO: Distinguish between multiple cards
//TODO: Bubble up errors rather than panicking
//TODO: Concept of users
type Transaction struct {
	Id				  int
  UserId      int
  AccountId   int
	BankName    string
  BankType    string
	TransactionType string
	Amount      float64
	Date        string
	Description string
}

type Record interface {
	ParseTransactionFromCsvLine(csvLine []string) Transaction
	GetTransaction() Transaction
}

type EmpowerFcuTransaction struct {
	Transaction
}

type SimpleTransaction struct {
	Transaction
}

type ChaseTransaction struct {
	Transaction
}

type BankOfAmericaTransaction struct {
	Transaction
}

func (t Transaction) GetTransaction() Transaction {
	return t
}

func (t Transaction) ParseTransactionFromCsvLine(csvLine []string) Transaction {
	panic(t)
}

func (t ChaseTransaction) ParseTransactionFromCsvLine(csvLine []string) Transaction {
  var transactionType string
  transactionAmount, err := strconv.ParseFloat(csvLine[4], 2)
	if err != nil {
    panic(err)
  }
  if csvLine[0] == "Sale" {
    transactionType = Sale
  } else {
    transactionType = Payment
  }
  date, err := time.Parse(csvDateFormat, csvLine[1])
	if err != nil {
    panic(err)
  }
	return Transaction{
    TransactionType: transactionType,
    BankType: Credit,
		Description: csvLine[3],
		Amount:      math.Abs(transactionAmount),
		Date:        date.Format("2006-01-02"),
		BankName:    Chase,
  }
}

func (t SimpleTransaction) ParseTransactionFromCsvLine(csvLine []string) Transaction {
	transactionAmount, err := strconv.ParseFloat(csvLine[3], 2)
	if err != nil { panic(err) }

  var transactionType string
  if csvLine[0] == "purchase" {
    transactionType = Sale
  } else {
    transactionType = Transfer
  }

	return Transaction{
		TransactionType: transactionType,
    BankType: Debit,
		Description: csvLine[7],
		Amount:      math.Abs(transactionAmount),
		Date:        csvLine[0],
		BankName:    Simple,
  }
}

func (t BankOfAmericaTransaction) ParseTransactionFromCsvLine(csvLine []string) Transaction {
	transactionAmount, err := strconv.ParseFloat(csvLine[4], 2)
	if err != nil { panic (err) }

  var transactionType string
  if transactionAmount < 0 {
    transactionType = Payment
  } else {
    transactionType = Sale
  }
	return Transaction{
		TransactionType: transactionType,
    BankType: Credit,
		Description: csvLine[2],
		Amount:      math.Abs(transactionAmount),
		Date:        csvLine[0],
		BankName:    BankOfAmerica}
}
