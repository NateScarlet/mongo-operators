package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#array-expression-operators

// ArrayElemAt returns the element at the specified array index.
// https://docs.mongodb.com/manual/reference/operator/aggregation/arrayElemAt
func ArrayElemAt(array, index interface{}) M {
	return M{"$arrayElemAt": A{array, index}}
}

// ArrayToObject converts an array into a single document.
// https://docs.mongodb.com/manual/reference/operator/aggregation/arrayToObject/
func ArrayToObject(expr interface{}) M {
	return M{"$arrayToObject": expr}
}

// ConcatArrays concatenates arrays to return the concatenated array.
// https://docs.mongodb.com/manual/reference/operator/aggregation/concatArrays/
func ConcatArrays(array ...interface{}) M {
	return M{"$concatArrays": array}
}

// FilterOperator returned from Filter
type FilterOperator M

// Filter selects a subset of the array to return an array
// with only the elements that match the filter condition.
// https://docs.mongodb.com/manual/reference/operator/aggregation/filter/
func Filter(input interface{}, cond interface{}) FilterOperator {
	return FilterOperator{"$filter": M{
		"input": input,
		"cond":  cond,
	}}
}

// SetAs option
func (op FilterOperator) SetAs(v string) FilterOperator {
	op["$filter"].(M)["as"] = v
	return op
}

// In returns a boolean indicating whether a specified value is in an array
// https://docs.mongodb.com/manual/reference/operator/aggregation/in
func In(elem, array interface{}) M {
	return M{"$in": A{elem, array}}
}

// IndexOfArrayOperator returned from IndexOfArray
type IndexOfArrayOperator M

// IndexOfArray searches an array for an occurence of a specified value
// and returns the array index of the first occurence.
// If the substring is not found, returns -1.
// https://docs.mongodb.com/manual/reference/operator/aggregation/indexOfArray/
func IndexOfArray(array, search interface{}) IndexOfArrayOperator {
	return IndexOfArrayOperator{"$indexOfArray": A{array, search}}
}

// SetStart expression
func (op IndexOfArrayOperator) SetStart(expr interface{}) IndexOfArrayOperator {
	var a = op["$indexOfArray"].(A)
	if len(a) <= 3 {
		op["$indexOfArray"] = append(a, expr)
	} else {
		a[2] = expr
	}
	return op
}

// SetEnd expression
func (op IndexOfArrayOperator) SetEnd(expr interface{}) IndexOfArrayOperator {
	var a = op["$indexOfArray"].(A)
	if len(a) <= 4 {
		op["$indexOfArray"] = append(a, expr)
	} else {
		a[3] = expr
	}
	return op
}

// IsArray determines if the operand is an array. Returns a boolean.
// https://docs.mongodb.com/manual/reference/operator/aggregation/isArray/
func IsArray(expr interface{}) M {
	return M{"$isArray": A{expr}}
}

// MapOperator returned from Map
type MapOperator M

// Map applies a subexpression to each element of an array
// and returns the array of resulting values in order.
// Accepts named parameters.
// https://docs.mongodb.com/manual/reference/operator/aggregation/map/
func Map(input, in interface{}) MapOperator {
	return MapOperator{"$map": M{
		"input": input,
		"in":    in,
	}}
}

// SetAs option
func (op MapOperator) SetAs(v string) MapOperator {
	op["$map"].(M)["as"] = v
	return op
}

// ObjectToArray converts a document to an array of documents representing key-value pairs.
// https://docs.mongodb.com/manual/reference/operator/aggregation/objectToArray/
func ObjectToArray(object interface{}) M {
	return M{"$objectToArray": object}
}

// RangeOperator returned from Range
type RangeOperator M

// Range outputs an array containing a sequence of integer
// according to user-defined inputs.
// https://docs.mongodb.com/manual/reference/operator/aggregation/range/
func Range(start, end interface{}) RangeOperator {
	return RangeOperator{"$range": A{start, end}}
}

// SetStep expression
func (op RangeOperator) SetStep(expr interface{}) RangeOperator {
	var a = op["$range"].(A)
	if len(a) <= 3 {
		op["$range"] = append(a, expr)
	} else {
		a[2] = expr
	}
	return op
}

// Reduce applies an expression to each element in an array
// and combines them into a single value.
// https://docs.mongodb.com/manual/reference/operator/aggregation/reduce/#exp._S_reduce
func Reduce(input, initialValue, in interface{}) M {
	return M{"$reduce": M{
		"input":        input,
		"initialValue": initialValue,
		"in":           in,
	}}
}

// ReverseArray returns an array with the elements in reverse order.
// https://docs.mongodb.com/manual/reference/operator/aggregation/reverseArray/
func ReverseArray(array interface{}) M {
	return M{"$reverseArray": array}
}

// Size returns the number of elements in the array.
// https://docs.mongodb.com/manual/reference/operator/aggregation/size/
func Size(array interface{}) M {
	return M{"$size": array}
}

// SliceOperator returned from slice
type SliceOperator M

// Slice returns a subset of an array.
// https://docs.mongodb.com/manual/reference/operator/aggregation/slice/
func Slice(array, n interface{}) SliceOperator {
	return SliceOperator{
		"$slice": A{array, n},
	}
}

// SetPos option
func (op SliceOperator) SetPos(v interface{}) SliceOperator {
	var a = op["$slice"].(A)
	if len(a) == 3 {
		a[1] = v
	} else {
		op["$slice"] = A{a[0], v, a[1]}
	}
	return op
}

// ZipOperator returned from Zip
type ZipOperator M

// Zip merge two arrays together.
// https://docs.mongodb.com/manual/reference/operator/aggregation/zip/
func Zip(inputs ...interface{}) ZipOperator {
	return ZipOperator{"$zip": M{"inputs": inputs}}
}

// SetUseLongestLength option
func (op ZipOperator) SetUseLongestLength(v bool) ZipOperator {
	op["$zip"].(M)["useLongestLength"] = v
	return op
}

// SetDefaults option, also set useLongestLength to true.
func (op ZipOperator) SetDefaults(array interface{}) ZipOperator {
	op.SetUseLongestLength(true)
	op["$zip"].(M)["defaults"] = array
	return op
}
