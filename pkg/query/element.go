package query

// https://docs.mongodb.com/manual/reference/operator/query-element/

// Exists matches documents that have the specified field.
// https://docs.mongodb.com/manual/reference/operator/query/exists/
func Exists(v bool) M {
	return M{"$exists": v}
}

// Type selects documents if a field is of the specified type.
// https://docs.mongodb.com/manual/reference/operator/query/type/
func Type(v bool) M {
	return M{"$type": v}
}
