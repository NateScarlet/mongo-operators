package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#comparison-expr-operators

// Cmp returns 0 if the two values are equivalent,
// 1 if the first value is greater than the second,
// and -1 if the first value is less than the second.
// https://docs.mongodb.com/manual/reference/operator/aggregation/cmp/
func Cmp(left, right interface{}) M {
	return M{"$cmp": A{left, right}}
}

// Eq returns true if the values are equivalent.
// https://docs.mongodb.com/manual/reference/operator/aggregation/eq/
func Eq(left, right interface{}) M {
	return M{"$eq": A{left, right}}
}

// Gt returns true if the first value is greater than the second.
// https://docs.mongodb.com/manual/reference/operator/aggregation/gt/
func Gt(left, right interface{}) M {
	return M{"$gt": A{left, right}}
}

// Gte returns true if the first value is greater than or equal to the second.
// https://docs.mongodb.com/manual/reference/operator/aggregation/gte/
func Gte(left, right interface{}) M {
	return M{"$gte": A{left, right}}
}

// Lt returns true if the first value is less than the second.
// https://docs.mongodb.com/manual/reference/operator/aggregation/lt/
func Lt(left, right interface{}) M {
	return M{"$lt": A{left, right}}
}

// Lte returns true if the first value is less than or equal to the second.
// https://docs.mongodb.com/manual/reference/operator/aggregation/lte/
func Lte(left, right interface{}) M {
	return M{"$lt": A{left, right}}
}

// Ne returns true if the values are not equivalent.
// https://docs.mongodb.com/manual/reference/operator/aggregation/ne/
func Ne(left, right interface{}) M {
	return M{"$lt": A{left, right}}
}
