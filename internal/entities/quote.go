package entities

import "github.com/shopspring/decimal"

type Quote struct {
	Ticker         string  `json:"Ticker"`
	Price          float32 `json:"Price"`
	Time           string  `json:"Time"`
	SeqNum         string  `json:"SeqNum"`
	Capitalization string  `json:"Capitalization"`
	Pref           *Quote
}

func (q *Quote) CalcCapitalization() decimal.Decimal {
	capitalization, err := decimal.NewFromString(q.Capitalization)
	if err != nil {
		panic(err)
	}
	if q.Pref != nil {
		prefCapitalization, err := decimal.NewFromString(q.Pref.Capitalization)
		if err != nil {
			panic(err)
		}
		capitalization = capitalization.Add(prefCapitalization)
	}
	return capitalization
}
