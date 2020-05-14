package query

// https://docs.mongodb.com/manual/reference/operator/query-logical/

// And joins query clauses with a logical AND returns all documents
// that match the conditions of both clauses.
// https://docs.mongodb.com/manual/reference/operator/query/and/
func And(expr ...interface{}) M {
	return M{"$and": expr}
}

// Not inverts the effect of a query expression and returns documents
// that do not match the query expression.
// https://docs.mongodb.com/manual/reference/operator/query/not/
func Not(expr ...interface{}) M {
	return M{"$not": expr}
}

// Nor inverts the effect of a query expression and returns documents
// that do not match the query expression.
// https://docs.mongodb.com/manual/reference/operator/query/nor/
func Nor(expr ...interface{}) M {
	return M{"$nor": expr}
}

// Or joins query clauses with a logical OR returns all documents
// that match the conditions of either clause.
// https://docs.mongodb.com/manual/reference/operator/query/or/
func Or(expr ...interface{}) M {
	return M{"$or": expr}
}
