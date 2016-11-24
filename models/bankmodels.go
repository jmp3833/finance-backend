package models

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
	Name        string
	DbName      string
}

type Record interface {
	GetTransaction(csvLine []string) Transaction
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

func (t ChaseTransaction) GetTransaction(csvLine []string) Transaction {
	//TODO: Finish me
	return Transaction{}
}
