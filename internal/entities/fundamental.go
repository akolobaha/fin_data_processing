package entities

type Fundamental struct {
	Ticker             string          `bson:"Ticker"`
	ReportMethod       string          `bson:"ReportMethod"`
	PeriodType         string          `bson:"PeriodType"`
	Period             string          `bson:"Period"`
	Date               string          `bson:"Date" json:"Date"`
	Currency           string          `bson:"Currency" json:"Currency"`
	ReportUrl          string          `bson:"ReportUrl"`
	Revenue            FundamentalItem `bson:"Revenue" json:"Revenue"`
	OperatingIncome    FundamentalItem `bson:"OperatingIncome" json:"OperatingIncome"`
	EBITDA             FundamentalItem `bson:"EBITDA" json:"EBITDA"`
	NetIncome          FundamentalItem `bson:"NetIncome" json:"NetIncome"`
	Ocf                FundamentalItem `bson:"Ocf" json:"Ocf"`
	Capex              FundamentalItem `bson:"Capex" json:"Capex"`
	Fcf                FundamentalItem `bson:"Fcf" json:"Fcf"`
	DividendPayout     FundamentalItem `bson:"DividendPayout" json:"DividendPayout"`
	Dividend           FundamentalItem `bson:"Dividend" json:"Dividend"`
	DivYield           FundamentalItem `bson:"DivYield" json:"DivYield"`
	DivPayoutRatio     FundamentalItem `bson:"DivPayoutRatio" json:"DivPayoutRatio"`
	Opex               FundamentalItem `bson:"Opex" json:"Opex"`
	Amortization       FundamentalItem `bson:"Amortization" json:"Amortization"`
	EmploymentExpenses FundamentalItem `bson:"EmploymentExpenses" json:"EmploymentExpenses"`
	InterestExpenses   FundamentalItem `bson:"InterestExpenses" json:"InterestExpenses"`
	Assets             FundamentalItem `bson:"Assets" json:"Assets"`
	NetAssets          FundamentalItem `bson:"NetAssets" json:"NetAssets"`
	Debt               FundamentalItem `bson:"Debt" json:"Debt"`
	Cash               FundamentalItem `bson:"Cash" json:"Cash"`
	NetDebt            FundamentalItem `bson:"NetDebt" json:"NetDebt"`
	CommonShare        FundamentalItem `bson:"CommonShare" json:"CommonShare"`
	NumberOfShares     FundamentalItem `bson:"NumberOfShares" json:"NumberOfShares"`
	MarketCap          FundamentalItem `bson:"MarketCap" json:"MarketCap"`
	Ev                 FundamentalItem `bson:"Ev" json:"Ev"`
	BookValue          FundamentalItem `bson:"BookValue" json:"BookValue"`
	Eps                FundamentalItem `bson:"Eps" json:"Eps"`
	FcfShare           FundamentalItem `bson:"FcfShare" json:"FcfShare"`
	BvShare            FundamentalItem `bson:"BvShare" json:"BvShare"`
	EbitdaMargin       FundamentalItem `bson:"EbitdaMargin" json:"EbitdaMargin"`
	NetMargin          FundamentalItem `bson:"NetMargin" json:"NetMargin"`
	FcfYield           FundamentalItem `bson:"FcfYield" json:"FcfYield"`
	Roe                FundamentalItem `bson:"Roe" json:"Roe"`
	Roa                FundamentalItem `bson:"Roa" json:"Roa"`
	PE                 string          `bson:"PE" json:"PE"`
	PFcf               FundamentalItem `bson:"PFcf" json:"PFcf"`
	PS                 string          `bson:"PS" json:"ps"`
	PBv                string          `bson:"PBv" json:"pbv"`
	EvEbitda           FundamentalItem `bson:"EvEbitda" json:"EvEbitda"`
	DebtEbitda         FundamentalItem `bson:"DebtEbitda" json:"DebtEbitda"`
	RAndDCapex         FundamentalItem `bson:"RAndDCapex" json:"RAndDCapex"`
	CapexRevenue       FundamentalItem `bson:"CapexRevenue" json:"CapexRevenue"`
}

type FundamentalItem struct {
	Name    string `bson:"Name" json:"Name"`
	Value   string `bson:"Value" json:"Value"`
	Measure string `bson:"Measure" json:"Measure"`
}
