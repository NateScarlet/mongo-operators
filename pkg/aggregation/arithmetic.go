package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#arithmetic-expression-operators

// Abs returns the absolute value of a number.
// https://docs.mongodb.com/manual/reference/operator/aggregation/abs/
func Abs(number interface{}) M {
	return M{"$abs": number}
}

// Add numbers to return the sum, or adds numbers and a date to return a new date.
// If adding numbers and a date, treats the numbers as milliseconds.
// Accepts any number of argument expressions, but at most,
// one expression can resolve to a date.
// https://docs.mongodb.com/manual/reference/operator/aggregation/add/
func Add(numberOrDate ...interface{}) M {
	return M{"$add": numberOrDate}
}

// Ceil returns the smallest integer greater than or equal to the specified number.
// https://docs.mongodb.com/manual/reference/operator/aggregation/ceil/
func Ceil(number interface{}) M {
	return M{"$ceil": number}
}

// Divide one number by another and returns the result.
// https://docs.mongodb.com/manual/reference/operator/aggregation/divide/
func Divide(dividend, divisor interface{}) M {
	return M{"$divide": A{dividend, divisor}}
}

// Exp raises Euler’s number (i.e. e ) to the specified exponent and returns the result.
// https://docs.mongodb.com/manual/reference/operator/aggregation/exp/
func Exp(exponent interface{}) M {
	return M{"$exp": exponent}
}

// Floor returns the largest integer less than or equal to the specified number.
// https://docs.mongodb.com/manual/reference/operator/aggregation/floor/
func Floor(number interface{}) M {
	return M{"$floor": number}
}

// Ln calculates the natural logarithm ln (i.e loge) of a number
// and returns the result as a double.
// https://docs.mongodb.com/manual/reference/operator/aggregation/ln/
func Ln(number interface{}) M {
	return M{"$ln": number}
}

// Log calculates the log of a number in the specified base
// and returns the result as a double.
// https://docs.mongodb.com/manual/reference/operator/aggregation/log/
func Log(number, base interface{}) M {
	return M{"$log": A{number, base}}
}

// Log10 calculates the log base 10 of a number and returns the result as a double.
// https://docs.mongodb.com/manual/reference/operator/aggregation/log10/
func Log10(number interface{}) M {
	return M{"$log10": number}
}

// Mod divides one number by another and returns the remainder.
// https://docs.mongodb.com/manual/reference/operator/aggregation/mod/
func Mod(dividend, divisor interface{}) M {
	return M{"$mod": A{dividend, divisor}}
}

// Multiply numbers together and returns the result.
// https://docs.mongodb.com/manual/reference/operator/aggregation/multiply/
func Multiply(number1, number2 interface{}) M {
	return M{"$multiply": A{number1, number2}}
}

// Pow raises a number to the specified exponent and returns the result.
// https://docs.mongodb.com/manual/reference/operator/aggregation/pow/
func Pow(number, exponent interface{}) M {
	return M{"$pow": A{number, exponent}}
}

// Round a number to to a whole integer or to a specified decimal place
// https://docs.mongodb.com/manual/reference/operator/aggregation/round/
func Round(number, place interface{}) M {
	return M{"$round": A{number, place}}
}

// Sqrt calculates the square root of a positive number
// and returns the result as a double.
// https://docs.mongodb.com/manual/reference/operator/aggregation/sqrt/
func Sqrt(number interface{}) M {
	return M{"$sqrt": number}
}

// Subtract two numbers to return the difference,
// or two dates to return the difference in milliseconds,
// or a date and a number in milliseconds to return the resulting date.
// https://docs.mongodb.com/manual/reference/operator/aggregation/subtract/
func Subtract(left, right interface{}) M {
	return M{"$subtract": A{left, right}}
}

// Trunc truncates a number to a whole integer or to a specified decimal place.
// https://docs.mongodb.com/manual/reference/operator/aggregation/trunc/
func Trunc(number, place interface{}) M {
	if place == nil {
		return M{"$trunc": number}
	}
	return M{"$trunc": A{number, place}}
}
