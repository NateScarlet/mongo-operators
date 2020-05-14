package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#set-expression-operators

// SetIntersection https://docs.mongodb.com/manual/reference/operator/aggregation/setIntersection
func SetIntersection(arrayExpr ...interface{}) M {
	return M{"$setIntersection": arrayExpr}
}
