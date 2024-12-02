package entities

type Fundamental struct {
	Ticker       string `bson:"Ticker"`
	ReportMethod string `bson:"ReportMethod"`
	Report       string `bson:"Report"`
	Period       string `bson:"Period"`
	Date         string `bson:"Date" json:"Date"`
	Currency     string `bson:"Currency" json:"Currency"`
	SourceUrl    string `bson:"SourceUrl" json:"SourceUrl"`
	ReportUrl    string `bson:"ReportUrl" json:"ReportUrl"`
	Revenue      uint64 `bson:"Revenue" json:"Revenue"`
	NetIncome    uint64 `bson:"NetIncome" json:"NetIncome"` // Чистая прибыль
	BookValue    uint64 `bson:"BookValue" json:"BookValue"`
}

type FundamentalItem struct {
	Name    string `bson:"Name" json:"Name"`
	Value   string `bson:"Value" json:"Value"`
	Measure string `bson:"Measure" json:"Measure"`
}
