package service

import (
	"errors"
	"fin_data_processing/internal/entities"
	"github.com/shopspring/decimal"
)

const (
	TargetPbv   = "pbv"
	TargetPe    = "pe"
	TargetPs    = "ps"
	TargetPrice = "price"
)

func TargetsAchievementCheck(target entities.TargetUser, fundamental entities.Fundamental, quote entities.Quote) (isAchieved bool, currResult float64, err error) {
	capitalization := quote.CalcTotalCapitalization()
	targetDec := decimal.NewFromFloat(target.Target.Value)

	switch target.Target.ValuationRatio {
	case TargetPbv:
		return checkRatio(targetDec, capitalization, fundamental.BookValue, "book value")
	case TargetPe:
		return checkRatio(targetDec, capitalization, fundamental.NetIncome, "net income")
	case TargetPs:
		return checkRatio(targetDec, capitalization, fundamental.Revenue, "revenue")
	case TargetPrice:
		return target.Target.Value >= float64(quote.Price), float64(quote.Price), nil
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

	return !result.GreaterThan(target), current, nil
}
