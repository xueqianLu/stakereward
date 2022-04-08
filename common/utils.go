package common

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func BigAmountFormatToFloat64OfDecimal(amount big.Int, decimal float64) float64 {
	amountFloat := new(big.Float).SetInt(&amount)
	dec := new(big.Float).SetFloat64(math.Pow(10, -decimal))
	mul, _ := new(big.Float).Mul(amountFloat, dec).Float64()
	return Decimal(mul, decimal)
}

func Decimal(value float64, decimal float64) float64 {
	format := "%." + Float64String(decimal) + "f"
	value, _ = strconv.ParseFloat(fmt.Sprintf(format, value), 64)
	return value
}

//float64转换string
func Float64String(data float64) string {
	var float64_strig string
	float64_strig = strconv.FormatFloat(data, 'f', -1, 64)
	return float64_strig
}
