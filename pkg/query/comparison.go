package query

// https://docs.mongodb.com/manual/reference/operator/query-comparison/

// Eq matches values that are equal to a specified value.
// https://docs.mongodb.com/manual/reference/operator/query/eq
func Eq(value interface{}) M {
	return M{"$eq": value}
}

// Gt matches values that are greater than a specified value.
// https://docs.mongodb.com/manual/reference/operator/query/gt
func Gt(value interface{}) M {
	return M{"$gt": value}
}

// Gte matches values that are greater than or equal to a specified value.
// https://docs.mongodb.com/manual/reference/operator/query/gte/
func Gte(value interface{}) M {
	return M{"$gte": value}
}

// In matches any of the values specified in an array.
// https://docs.mongodb.com/manual/reference/operator/query/in/
func In(value interface{}) M {
	return M{"$in": value}
}

// Lt matches values that are less than a specified value.
// https://docs.mongodb.com/manual/reference/operator/query/lt/
func Lt(value interface{}) M {
	return M{"$lt": value}
}

// Lte matches values that are less than or equal to a specified value.
// https://docs.mongodb.com/manual/reference/operator/query/lte/
func Lte(value interface{}) M {
	return M{"$lte": value}
}

// Ne matches all values that are not equal to a specified value.
// https://docs.mongodb.com/manual/reference/operator/query/ne/
func Ne(value interface{}) M {
	return M{"$ne": value}
}

// Nin matches none of the values specified in an array.
// https://docs.mongodb.com/manual/reference/operator/query/nin/
func Nin(value interface{}) M {
	return M{"$nin": value}
}
