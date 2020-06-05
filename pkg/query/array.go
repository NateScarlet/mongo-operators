package query

// https://docs.mongodb.com/manual/reference/operator/query-array/

// All selects the documents where the value of a field
// is an array that contains all the specified elements.
// https://docs.mongodb.com/manual/reference/operator/query/all/
func All(value interface{}) M {
	return M{"$all": value}
}

// ElemMatch matches documents that contain an array field with
// at least one element that matches all the specified query criteria.
// https://docs.mongodb.com/manual/reference/operator/query/elemMatch/
func ElemMatch(query interface{}) M {
	return M{"$elemMatch": query}
}

// Size matches any array with the number of elements specified by the argument.
// https://docs.mongodb.com/manual/reference/operator/query/size/
func Size(n int) M {
	return M{"$size": n}
}
