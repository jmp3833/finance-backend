package models

type FcU struct{ Bank }
type Simple struct{ Bank }
type Chase struct{ Bank }
type BofA struct{ Bank }

type Bank struct {
	Ref         string
	Date        string
	Description string
	Amount      float64
	Transtype   string
}
