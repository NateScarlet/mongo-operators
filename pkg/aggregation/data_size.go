package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#data-size-operators

// BinarySize returns the size of
// a given string or binary data value's content in bytes.
// New in version 4.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/binarySize/
func BinarySize(expr interface{}) M {
	return M{"$binarySize": expr}
}

// BSONSize returns the size in bytes of a given document
// (i.e. bsontype Object) when encoded as BSON.
// https://docs.mongodb.com/manual/reference/operator/aggregation/bsonSize/
func BSONSize(expr interface{}) M {
	return M{"$bsonSize": expr}
}
