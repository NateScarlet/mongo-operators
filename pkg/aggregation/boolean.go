package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#boolean-expression-operators

// And returns true only when all its expressions evaluate to true.
// https://docs.mongodb.com/manual/reference/operator/aggregation/and
func And(expr ...interface{}) M {
	return M{"$and": expr}
}

// Not returns the boolean value that is the opposite of its argument expression.
// https://docs.mongodb.com/manual/reference/operator/aggregation/not
func Not(expr interface{}) M {
	return M{"$not": A{expr}}
}

// Or returns true when any of its expressions evaluates to true.
// https://docs.mongodb.com/manual/reference/operator/aggregation/or
func Or(expr ...interface{}) M {
	return M{"$or": expr}
}
