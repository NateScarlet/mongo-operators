package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#type-expression-operators

// Type https://docs.mongodb.com/manual/reference/operator/aggregation/type
func Type(expr interface{}) M {
	return M{"$type": expr}
}
