package service

import (
	"fin_data_processing/internal/entities"
	"github.com/shopspring/decimal"
)

const TARGET_PBV = "pbv"

func TargetsAchievementCheck(target entities.TargetUser, fundamental entities.Fundamental, quote entities.Quote) (isAchieved bool, currResult float64) {
	capitalization := quote.CalcCapitalization()
	targetVal := decimal.NewFromFloat(target.Target.Value)

	switch target.Target.ValuationRatio {
	case TARGET_PBV:
		return CheckPBv(targetVal, capitalization, fundamental)
	default:
		return false, 0
	}
}

// CheckPBv капитализация к балансовой стоимости
func CheckPBv(target decimal.Decimal, capitalization decimal.Decimal, fundamental entities.Fundamental) (achieved bool, current float64) {
	bookValue := decimal.NewFromUint64(fundamental.BookValue)
	res := capitalization.Div(bookValue)
	result, _ := res.Float64()

	// Чем меньше - тем лучше
	if res.GreaterThanOrEqual(target) {
		return false, result
	}
	return true, result
}
