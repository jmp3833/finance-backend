package models

const FcuName = "Empower Federal Credit Union"
const ChaseName = "Chase Bank"
const BofAName = "Bank of America"
const SimpleName = "Simple inc. Bank"

type fcU struct{ Bank }
type simple struct{ Bank }
type chase struct{ Bank }
type bofA struct{ Bank }

type Bank struct {
	Transtype   string
	Description string
	Amount      float64
	Date        string
	Ref         string
	BankName    string
	DbName      string
}

func GetEmpowerModel(bank Bank) Bank {
	bank.BankName = FcuName
	bank.DbName = "fcu"
	return bank
}

func GetChaseModel(bank Bank) Bank {
	bank.BankName = ChaseName
	bank.DbName = "chase"
	return bank
}
