package service

import (
	"errors"
	"fin_data_processing/internal/entities"
	"github.com/shopspring/decimal"
)

const (
	TARGET_PBV   = "pbv"
	TARGET_PE    = "pe"
	TARGET_PS    = "ps"
	TARGET_PRICE = "price"
)

func TargetsAchievementCheck(target entities.TargetUser, fundamental entities.Fundamental, quote entities.Quote) (isAchieved bool, currResult float64, err error) {
	capitalization := quote.CalcCapitalization()
	targetDec := decimal.NewFromFloat(target.Target.Value)

	switch target.Target.ValuationRatio {
	case TARGET_PBV:
		return checkRatio(targetDec, capitalization, fundamental.BookValue, "book value")
	case TARGET_PE:
		return checkRatio(targetDec, capitalization, fundamental.NetIncome, "net income")
	case TARGET_PS:
		return checkRatio(targetDec, capitalization, fundamental.Revenue, "revenue")
	case TARGET_PRICE:
		return target.Target.Value <= float64(quote.Price), float64(quote.Price), nil
	default:
		return false, 0, nil
	}
}

func checkRatio(target decimal.Decimal, capitalization decimal.Decimal, value uint64, valueName string) (achieved bool, current float64, err error) {
	if value == 0 {
		return false, 0, errors.New(valueName + " is zero")
	}

	result := capitalization.Div(decimal.NewFromUint64(value))
	current, _ = result.Float64()

	return !result.GreaterThanOrEqual(target), current, nil
}
