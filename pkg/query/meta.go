package query

// Explain provides information on the query plan.
// It returns a document that describes the process
// and indexes used to return the query.
// This may provide useful insight when attempting to optimize a query.
// https://docs.mongodb.com/manual/reference/operator/meta/explain/
func Explain() M {
	return M{"$explain": 1}
}

// Hint operator forces the query optimizer
// to use a specific index to fulfill the query.
// Specify the index either by the index name or by document.
// https://docs.mongodb.com/manual/reference/operator/meta/hint/
func Hint(hint interface{}) M {
	return M{"$hint": hint}
}

// MaxIndex specify the exclusive upper bound for
// a specific index in order to constrain the results of find().
// The $max specifies the upper bound for all keys of a specific index in order.
// https://docs.mongodb.com/manual/reference/operator/meta/max/
func MaxIndex(bound interface{}) M {
	return M{"$max": bound}
}

// MaxTimeMS specifies a cumulative time limit in milliseconds
// for processing operations on the cursor.
// MongoDB interrupts the operation at the earliest following interrupt point.
// https://docs.mongodb.com/manual/reference/operator/meta/maxTimeMS/
func MaxTimeMS(ms int64) M {
	return M{"$maxTimeMS": ms}
}

// MinIndex specify a $min value to specify the inclusive lower bound
// for a specific index in order to constrain the results of find().
// The $min specifies the lower bound for all keys of a specific index in order.
// https://docs.mongodb.com/manual/reference/operator/meta/min/
func MinIndex(bound interface{}) M {
	return M{"$min": bound}
}

// OrderBy sorts the results of a query in ascending or descending order.
// https://docs.mongodb.com/manual/reference/operator/meta/orderby/
func OrderBy(sort interface{}) M {
	return M{"$orderBy": sort}
}

// Query operator forces MongoDB to interpret an expression as a query.
// https://docs.mongodb.com/manual/reference/operator/meta/query/
func Query(expr interface{}) M {
	return M{"$query": expr}
}

// ReturnKey only return the index field or fields for the results of the query.
// If $returnKey is set to true and the query does not use an index
// to perform the read operation, the returned documents will not contain any fields
// https://docs.mongodb.com/manual/reference/operator/meta/returnKey/
func ReturnKey() M {
	return M{"$returnKey": true}
}

// ShowDiskLoc adds a field $diskLoc to the returned documents.
// The value of the added $diskLoc field is a document that contains the disk location information:
//
// "$diskLoc": {
//   "file": <int>,
//   "offset": <int>
// }
//
// https://docs.mongodb.com/manual/reference/operator/meta/showDiskLoc/
func ShowDiskLoc() M {
	return M{"$showDiskLoc": true}
}

// Natural use in conjunction with cursor.hint() to perform a collection scan
// to return documents in natural order.
// https://docs.mongodb.com/manual/reference/operator/meta/natural/
func Natural(direction int) M {
	return M{"$natural": direction}
}
