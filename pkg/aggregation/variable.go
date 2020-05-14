package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#variable-expression-operators

// Let defines variables for use within the scope of a subexpression and returns the result of the subexpression. Accepts named parameters.
// Accepts any number of argument expressions.
// https://docs.mongodb.com/manual/reference/operator/aggregation/let/
func Let(vars, in interface{}) M {
	return M{"$let": M{
		"vars": vars,
		"in":   in,
	}}
}
