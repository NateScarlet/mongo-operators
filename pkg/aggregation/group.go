package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#accumulators---group---bucket---bucketauto---setwindowfields-

// AddToSet returns an array of all unique values that results
// from applying an expression to each document in a group of documents
// that share the same group by key.
// Order of the elements in the output array is unspecified.
// $addToSet is only available in the $group stage.
// https://docs.mongodb.com/manual/reference/operator/aggregation/addToSet/
func AddToSet(expr interface{}) M {
	return M{"$addToSet": expr}
}

// Avg returns the average value of the numeric values.
// In MongoDB 3.2 and earlier, $avg is available in the $group stage only.
// https://docs.mongodb.com/manual/reference/operator/aggregation/avg/
func Avg(expr ...interface{}) M {
	if len(expr) == 1 {
		return M{"$avg": expr[0]}
	}
	return M{"$avg": expr}
}

// First returns the value that results from applying an expression
// to the first document in a group of documents that share the same group by key.
// Only meaningful when documents are in a defined order.
// $first is only available in the $group stage.
// https://docs.mongodb.com/manual/reference/operator/aggregation/first/
func First(expr interface{}) M {
	return M{"$first": expr}
}

// Last returns the value that results from applying an expression
// to the last document in a group of documents that share the same group by key.
// Only meaningful when documents are in a defined order.
// $last is only available in the $group stage.
// https://docs.mongodb.com/manual/reference/operator/aggregation/last/
func Last(expr interface{}) M {
	return M{"$last": expr}
}

// Max returns the maximum value. $max compares both value and type,
// using the specified BSON comparison order for values of different types.
// In MongoDB 3.2 and earlier, $max is available in the $group stage only.
// https://docs.mongodb.com/manual/reference/operator/aggregation/max/
func Max(expr ...interface{}) M {
	if len(expr) == 1 {
		return M{"$max": expr[0]}
	}
	return M{"$max": expr}
}

// MergeObjects combines multiple documents into a single document.
// New in version 3.6.
// https://docs.mongodb.com/manual/reference/operator/aggregation/mergeObjects/
func MergeObjects(expr ...interface{}) M {
	if len(expr) == 1 {
		return M{"$mergeObjects": expr[0]}
	}
	return M{"$mergeObjects": expr}
}

// Min Returns the minimum value. $min compares both value and type,
// using the specified BSON comparison order for values of different types.
// In MongoDB 3.2 and earlier, $min is available in the $group stage only.
// https://docs.mongodb.com/manual/reference/operator/aggregation/min/
func Min(expr ...interface{}) M {
	if len(expr) == 1 {
		return M{"$min": expr[0]}
	}
	return M{"$min": expr}
}

// Push Returns an array of all values that result from applying an expression
// to each document in a group of documents that share the same group by key.
// $push is only available in the $group stage.
// https://docs.mongodb.com/manual/reference/operator/aggregation/push/
func Push(expr interface{}) M {
	return M{"$push": expr}
}

// StdDevPop returns the population standard deviation of the input values.
// New in version 3.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/stdDevPop/
func StdDevPop(expr ...interface{}) M {
	if len(expr) == 1 {
		return M{"$stdDevPop": expr[0]}
	}
	return M{"$stdDevPop": expr}
}

// StdDevSamp returns the sample standard deviation of the input values.
// New in version 3.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/stdDevSamp/
func StdDevSamp(expr ...interface{}) M {
	if len(expr) == 1 {
		return M{"$stdDevSamp": expr[0]}
	}
	return M{"$stdDevSamp": expr}
}

// Sum calculates and returns the sum of numeric values.
// $sum ignores non-numeric values.
// In MongoDB 3.2 and earlier, $sum is available in the $group stage only.
// https://docs.mongodb.com/manual/reference/operator/aggregation/sum/
func Sum(expr ...interface{}) M {
	if len(expr) == 1 {
		return M{"$sum": expr[0]}
	}
	return M{"$sum": expr}
}
