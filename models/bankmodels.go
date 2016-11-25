package models

import (
	"math"
	"strconv"
)

const (
	EmpowerFCU    = "Empower Federal Credit Union"
	BankOfAmerica = "Bank of America"
	Simple        = "Simple inc. Bank"
	Chase         = "Chase Bank"
)

type Transaction struct {
	Transtype   string
	Description string
	Amount      float64
	Date        string
	BankName    string
	DbName      string
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
	transamount, err := strconv.ParseFloat(csvLine[4], 2)
	if err != nil {
		panic(err)
	}
	return Transaction{
		Transtype:   csvLine[0],
		Description: csvLine[3],
		Amount:      math.Abs(transamount),
		Date:        csvLine[1],
		DbName:      "chase",
		BankName:    Chase}
}

func (t EmpowerFcuTransaction) ParseTransactionFromCsvLine(csvLine []string) Transaction {
	transamount, err := strconv.ParseFloat(csvLine[4], 2)
	if err != nil {
		panic(err)
	}
	return Transaction{
		Transtype:   csvLine[3],
		Description: csvLine[8],
		Amount:      math.Abs(transamount),
		Date:        csvLine[1],
		DbName:      "fcu",
		BankName:    EmpowerFCU}
}
