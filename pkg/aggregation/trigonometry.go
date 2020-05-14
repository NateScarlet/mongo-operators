package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#trigonometry-expression-operators

// Sin returns the sine of a value that is measured in radians.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/sin/
func Sin(number interface{}) M {
	return M{"$sin": number}
}

// Cos returns the cosine of a value that is measured in radians.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/cos/
func Cos(number interface{}) M {
	return M{"$cos": number}
}

// Tan returns the tangent of a value that is measured in radians.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/tan/
func Tan(number interface{}) M {
	return M{"$tan": number}
}

// ASin returns the inverse sin (arc sine) of a value in radians.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/asin/
func ASin(number interface{}) M {
	return M{"$asin": number}
}

// ACos returns the inverse cosine (arc cosine) of a value in radians.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/acos/
func ACos(number interface{}) M {
	return M{"$acos": number}
}

// ATan returns the inverse tangent (arc tangent) of a value in radians.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/atan/
func ATan(number interface{}) M {
	return M{"$atan": number}
}

// ATan2 returns the inverse tangent (arc tangent) of y / x in radians,
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/atan2/
func ATan2(y, x interface{}) M {
	return M{"$atan2": A{y, x}}
}

// ASinH returns the inverse hyperbolic sine (hyperbolic arc sine) of a value in radians.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/asin/
func ASinH(number interface{}) M {
	return M{"$asinh": number}
}

// ACosH returns the inverse hyperbolic cosine (hyperbolic arc cosine) of a value in radians.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/acos/
func ACosH(number interface{}) M {
	return M{"$acosh": number}
}

// ATanH returns the inverse hyperbolic tangent (hyperbolic arc tangent) of a value in radians.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/atan/
func ATanH(number interface{}) M {
	return M{"$atanh": number}
}

// DegreesToRadians converts a value from degrees to radians.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/degreesToRadians/
func DegreesToRadians(number interface{}) M {
	return M{"$degreesToRadians": number}
}

// RadiansToDegrees converts a value from radians to degrees.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/radiansToDegrees/
func RadiansToDegrees(number interface{}) M {
	return M{"$radiansToDegrees": number}
}
