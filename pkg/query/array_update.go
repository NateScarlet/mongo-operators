package query

// https://docs.mongodb.com/manual/reference/operator/update-array/

// AddToSet adds elements to an array only if they do not already exist in the set.
// https://docs.mongodb.com/manual/reference/operator/update/addToSet/
func AddToSet(spec interface{}) M {
	return M{"$addToSet": spec}
}

// Pop removes the first or last element of an array.
// Pass a value of -1 to remove the first element of an array
// and 1 to remove the last element in an array.
// https://docs.mongodb.com/manual/reference/operator/update/pop/
func Pop(spec interface{}) M {
	return M{"$pop": spec}
}

// Pull removes from an existing array all instances of a value
// or values that match a specified condition.
// https://docs.mongodb.com/manual/reference/operator/update/pull/
func Pull(spec interface{}) M {
	return M{"$pull": spec}
}

// Push appends a specified value to an array.
// https://docs.mongodb.com/manual/reference/operator/update/push/
func Push(spec interface{}) M {
	return M{"$push": spec}
}

// PullAll removes all instances of the specified values from an existing array.
// Unlike the $pull operator that removes elements by specifying a query,
// $pullAll removes elements that match the listed values.
// https://docs.mongodb.com/manual/reference/operator/update/pullAll/
func PullAll(spec interface{}) M {
	return M{"$pullAll": spec}
}

// EachModifier returned from Each
type EachModifier M

// Each modifies the $push and $addToSet operators to
// append multiple items for array updates.
// https://docs.mongodb.com/manual/reference/operator/update/each/
func Each(value ...interface{}) EachModifier {
	return EachModifier{"$each": value}
}

// Position modifier specifies the location in the array at
// which the $push operator inserts elements.
// Without the $position modifier, the $push operator inserts elements
// to the end of the array.
// https://docs.mongodb.com/manual/reference/operator/update/position/
func (m EachModifier) Position(at int) EachModifier {
	m["$position"] = at
	return m
}

// Slice modifier limits the number of array elements during a $push operation.
// https://docs.mongodb.com/manual/reference/operator/update/slice/
func (m EachModifier) Slice(num int) EachModifier {
	m["$slice"] = num
	return m
}

// Sort modifier orders the elements of an array during a $push operation.
// https://docs.mongodb.com/manual/reference/operator/update/sort/
func (m EachModifier) Sort(sort interface{}) EachModifier {
	m["$sort"] = sort
	return m
}
