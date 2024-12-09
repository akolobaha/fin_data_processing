package entities

import (
	"testing"
)

func TestFundamentalCache(t *testing.T) {
	cache := NewFundamentalCache()

	// Создаем пример данных
	fundamental1 := Fundamental{
		Ticker:       "AAPL",
		ReportMethod: "RSBU",
		Report:       "Q4 2023",
		Period:       "2023-12-31",
		Date:         "2023-12-31",
		Currency:     "USD",
		SourceURL:    "http://example.com/source1",
		ReportURL:    "http://example.com/report1",
		Revenue:      123456789,
		NetIncome:    98765432,
		BookValue:    100000000,
	}

	fundamental2 := Fundamental{
		Ticker:       "AAPL",
		ReportMethod: "MSFO",
		Report:       "Q4 2023",
		Period:       "2023-12-31",
		Date:         "2023-12-31",
		Currency:     "USD",
		SourceURL:    "http://example.com/source2",
		ReportURL:    "http://example.com/report2",
		Revenue:      223456789,
		NetIncome:    198765432,
		BookValue:    200000000,
	}

	// Тест: Установка значений в кэш
	cache.Set(fundamental1.Ticker, fundamental1.ReportMethod, fundamental1)
	cache.Set(fundamental2.Ticker, fundamental2.ReportMethod, fundamental2)

	// Тест: Получение значений из кэша
	retrieved1, ok1 := cache.Get(fundamental1.Ticker, fundamental1.ReportMethod)
	if !ok1 {
		t.Fatalf("expected to find ticker %s with report method %s in cache", fundamental1.Ticker, fundamental1.ReportMethod)
	}
	if retrieved1 != fundamental1 {
		t.Errorf("expected %v, got %v", fundamental1, retrieved1)
	}

	retrieved2, ok2 := cache.Get(fundamental2.Ticker, fundamental2.ReportMethod)
	if !ok2 {
		t.Fatalf("expected to find ticker %s with report method %s in cache", fundamental2.Ticker, fundamental2.ReportMethod)
	}
	if retrieved2 != fundamental2 {
		t.Errorf("expected %v, got %v", fundamental2, retrieved2)
	}

	// Тест: Попытка получения несуществующего тикера
	_, ok3 := cache.Get("GOOGL", "RSBU")
	if ok3 {
		t.Fatal("expected not to find ticker GOOGL in cache")
	}

	// Тест: Попытка получения несуществующего метода отчета
	_, ok4 := cache.Get(fundamental1.Ticker, "NonexistentMethod")
	if ok4 {
		t.Fatal("expected not to find report method NonexistentMethod in cache")
	}

	// Тест: Удаление значения из кэша
	cache.Delete(fundamental1.Ticker, fundamental1.ReportMethod)

	// Тест: Проверка, что тикер и метод отчета были удалены
	_, ok5 := cache.Get(fundamental1.Ticker, fundamental1.ReportMethod)
	if ok5 {
		t.Fatalf("expected ticker %s with report method %s to be deleted from cache", fundamental1.Ticker, fundamental1.ReportMethod)
	}

	// Тест: Удаление метода отчета, когда тикер все еще существует
	cache.Delete(fundamental2.Ticker, fundamental2.ReportMethod)

	// Тест: Проверка, что тикер был удален, если больше нет методов отчета
	_, ok6 := cache.Get(fundamental2.Ticker, fundamental2.ReportMethod)
	if ok6 {
		t.Fatalf("expected ticker %s with report method %s to be deleted from cache", fundamental2.Ticker, fundamental2.ReportMethod)
	}
	// Проверяем, что тикер полностью удален из кэша
	_, ok7 := cache.Get(fundamental2.Ticker, "RSBU")
	if ok7 {
		t.Fatalf("expected ticker %s to be deleted from cache", fundamental2.Ticker)
	}
}
