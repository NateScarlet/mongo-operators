package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#literal-expression-operator

// Literal return a value without parsing.
// Use for values that the aggregation pipeline may interpret as an expression.
// For example, use a $literal expression to a string
// that starts with a $ to avoid parsing as a field path.
// https://docs.mongodb.com/manual/reference/operator/aggregation/literal/
func Literal(value interface{}) M {
	return M{"$literal": value}
}
