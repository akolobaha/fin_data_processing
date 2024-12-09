package entities

import "sync"

type Fundamental struct {
	Ticker       string
	ReportMethod string
	Report       string
	Period       string
	Date         string `bson:"Date" json:"Date"`
	Currency     string `bson:"Currency" json:"Currency"`
	SourceURL    string
	ReportURL    string
	Revenue      uint64 `bson:"Revenue" json:"Revenue"`
	NetIncome    uint64 `bson:"NetIncome" json:"NetIncome"` // Чистая прибыль
	BookValue    uint64 `bson:"BookValue" json:"BookValue"`
}

// FundamentalCache map[тикер]map[rsbu/msfo]Fundamental
type FundamentalCache struct {
	mu    sync.RWMutex
	cache map[string]map[string]Fundamental
}

func NewFundamentalCache() *FundamentalCache {
	return &FundamentalCache{
		cache: make(map[string]map[string]Fundamental, 220),
	}
}

func (f *FundamentalCache) Set(ticker string, reportMethod string, fundamental Fundamental) {
	f.mu.Lock()
	defer f.mu.Unlock()

	// Инициализация вложенной карты, если она не существует
	if f.cache[ticker] == nil {
		f.cache[ticker] = make(map[string]Fundamental)
	}
	f.cache[ticker][reportMethod] = fundamental
}

func (f *FundamentalCache) Get(ticker string, reportMethod string) (Fundamental, bool) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	fundamental, ok := f.cache[ticker][reportMethod]
	return fundamental, ok
}

func (f *FundamentalCache) Delete(ticker string, reportMethod string) {
	f.mu.Lock()
	defer f.mu.Unlock()

	// Проверка наличия тикера и удаление метода отчета
	if _, exists := f.cache[ticker]; exists {
		delete(f.cache[ticker], reportMethod)
		// Если вложенная карта пуста, можно удалить и сам тикер
		if len(f.cache[ticker]) == 0 {
			delete(f.cache, ticker)
		}
	}
}
