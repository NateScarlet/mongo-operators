package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#window-operators

// CountAccumulator Returns the number of documents in a group.
//
// $count is available in these stages:
//
// - $bucket
// - $bucketAuto
// - $group
// - $setWindowFields (Available starting in MongoDB 5.0)
//
// https://docs.mongodb.com/manual/reference/operator/aggregation/count-accumulator/
func CountAccumulator() M {
	return M{"$count": M{}}
}

// CovariancePop returns the population covariance of two numeric expressions.
// New in version 5.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/covariancePop/
func CovariancePop(expr1, expr2 interface{}) M {
	return M{
		"$covariancePop": A{
			expr1,
			expr2,
		},
	}
}

// CovarianceSamp returns the sample covariance of two numeric expressions.
// New in version 5.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/covarianceSamp/
func CovarianceSamp(expr1, expr2 interface{}) M {
	return M{
		"$covarianceSamp": A{
			expr1,
			expr2,
		},
	}
}

// DenseRank returns the document position (known as the rank)
// relative to other documents in the $setWindowFields stage partition.
// There are no gaps in the ranks. Ties receive the same rank.
// New in version 5.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/denseRank/
func DenseRank() M {
	return M{"$denseRank": M{}}
}

type DerivativeOperator M

type DerivativeUnit string

const (
	DUWeek        DerivativeUnit = "week"
	DUDay         DerivativeUnit = "day"
	DUHour        DerivativeUnit = "hour"
	DUMinute      DerivativeUnit = "minute"
	DUSecond      DerivativeUnit = "second"
	DUMillisecond DerivativeUnit = "millisecond"
)

func (op DerivativeOperator) SetUnit(unit DerivativeUnit) DerivativeOperator {
	op["$derivative"].(M)["unit"] = unit
	return op
}

// Derivative returns the average rate of change
// within the specified window.
// New in version 5.0.
func Derivative(inputExpr interface{}) DerivativeOperator {
	return DerivativeOperator{
		"$derivative": M{
			"input": inputExpr,
		},
	}
}

// DocumentNumber returns the position of a document (known as the document number)
// in the $setWindowFields stage partition.
// Ties result in different adjacent document numbers.
// New in version 5.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/documentNumber/
func DocumentNumber() M {
	return M{"$documentNumber": M{}}
}

// ExpMovingAvgN returns the exponential moving average for the numeric expression.
// New in version 5.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/expMovingAvg/
func ExpMovingAvg(inputExpr interface{}, n int64) M {
	return M{"$expMovingAvg": M{
		"input": inputExpr,
		"N":     n,
	}}
}

// ExpMovingAvgAlpha returns the exponential moving average for the numeric expression.
// New in version 5.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/expMovingAvg/
func ExpMovingAvgAlpha(inputExpr interface{}, alpha float64) M {
	return M{"$expMovingAvg": M{
		"input": inputExpr,
		"alpha": alpha,
	}}
}

type IntegralOperator M

type IntegralUnit string

const (
	IUWeek        IntegralUnit = "week"
	IUDay         IntegralUnit = "day"
	IUHour        IntegralUnit = "hour"
	IUMinute      IntegralUnit = "minute"
	IUSecond      IntegralUnit = "second"
	IUMillisecond IntegralUnit = "millisecond"
)

func (op IntegralOperator) SetUnit(unit IntegralUnit) IntegralOperator {
	op["$integral"].(M)["unit"] = unit
	return op
}

// Integral returns the approximation of the area
// under a curve.
// New in version 5.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/integral/
func Integral(inputExpr interface{}) IntegralOperator {
	return IntegralOperator{
		"$integral": M{
			"input": inputExpr,
		},
	}
}

// Rank returns the document position (known as the rank)
// relative to other documents in the $setWindowFields stage partition.
// New in version 5.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/rank/
func Rank() M {
	return M{"$rank": M{}}
}

type ShiftOperator M

func (op ShiftOperator) SetDefault(expr interface{}) ShiftOperator {
	op["$shift"].(M)["default"] = expr
	return op
}

// Shift returns the value from an expression applied to a document
// in a specified position relative to the current document
// in the $setWindowFields stage partition.
// New in version 5.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/shift/
func Shift(outputExpr interface{}, by int64) ShiftOperator {
	return ShiftOperator{
		"$shift": M{
			"output": outputExpr,
			"by":     by,
		},
	}
}
