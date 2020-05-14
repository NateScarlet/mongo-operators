package query

// https://docs.mongodb.com/manual/reference/operator/query-bitwise/

// BitsAllClear matches documents where all of the bit positions
// given by the query are clear (i.e. 0) in field.
// New in version 3.2.
// https://docs.mongodb.com/manual/reference/operator/query/bitsAllClear/
func BitsAllClear(mask interface{}) M {
	return M{"$bitsAllClear": mask}
}

// BitsAllSet matches documents where all of the bit positions
// given by the query are set (i.e. 1) in field.
// New in version 3.2.
// https://docs.mongodb.com/manual/reference/operator/query/bitsAllSet/
func BitsAllSet(mask interface{}) M {
	return M{"$bitsAllSet": mask}
}

// BitsAnyClear matches documents where all of the bit positions
// given by the query are clear (i.e. 0) in field.
// New in version 3.2.
// https://docs.mongodb.com/manual/reference/operator/query/bitsAllClear/
func BitsAnyClear(mask interface{}) M {
	return M{"$bitsAnyClear": mask}
}

// BitsAnySet matches documents where any of the bit positions
// given by the query are set (i.e. 1) in field.
// New in version 3.2.
// https://docs.mongodb.com/manual/reference/operator/query/bitsAllSet/
func BitsAnySet(mask interface{}) M {
	return M{"$bitsAnySet": mask}
}
