package models

type Chase struct {
	Transtype   string
	Description string
	Amount      float64
	Date        string
}

type Simple struct {
	Transtype   string
	Amount      float64
	Pending     bool
	Desctiption string
	City        string
	Memo        string
}

type BofA struct {
	Date        string
	Ref         float64
	Description string
	Address     string
	Amount      float64
}

type FcU struct {
	Ref         string
	Date        string
	Description string
	Amount      float64
}
