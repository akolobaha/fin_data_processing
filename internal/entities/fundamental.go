package entities

type Fundamental struct {
	Ticker       string
	ReportMethod string
	Report       string
	Period       string
	Date         string `bson:"Date" json:"Date"`
	Currency     string `bson:"Currency" json:"Currency"`
	SourceUrl    string
	ReportUrl    string
	Revenue      uint64 `bson:"Revenue" json:"Revenue"`
	NetIncome    uint64 `bson:"NetIncome" json:"NetIncome"` // Чистая прибыль
	BookValue    uint64 `bson:"BookValue" json:"BookValue"`
}
