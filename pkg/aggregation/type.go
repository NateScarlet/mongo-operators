package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#type-expression-operators

// ConvertOperator returned from Convert
type ConvertOperator M

// Convert a value to a specified type.
// New in version 4.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/convert/
func Convert(input, to interface{}) ConvertOperator {
	return ConvertOperator{"$convert": M{
		"input": input,
		"to":    to,
	}}
}

// SetOnError option
func (op ConvertOperator) SetOnError(expr interface{}) ConvertOperator {
	op["$convert"].(M)["onError"] = expr
	return op
}

// SetOnNull option
func (op ConvertOperator) SetOnNull(expr interface{}) ConvertOperator {
	op["$convert"].(M)["onNull"] = expr
	return op
}

// ToBool converts value to a boolean.
// New in version 4.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/toBool/
func ToBool(expr interface{}) M {
	return M{"$toBool": expr}
}

// ToDate converts a value to a date.
// If the value cannot be converted to a date, $toDate errors.
// If the value is null or missing, $toDate returns null.
// New in version 4.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/toDate/
func ToDate(expr interface{}) M {
	return M{"$toDate": expr}
}

// ToDecimal converts value to a Decimal128.
// New in version 4.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/toDecimal/
func ToDecimal(expr interface{}) M {
	return M{"$toDecimal": expr}
}

// ToDouble converts value to a double.
// New in version 4.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/toDouble/
func ToDouble(expr interface{}) M {
	return M{"$toDouble": expr}
}

// ToInt converts value to a integer.
// New in version 4.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/toInt/
func ToInt(expr interface{}) M {
	return M{"$toInt": expr}
}

// ToLong converts value to a long.
// New in version 4.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/toLong/
func ToLong(expr interface{}) M {
	return M{"$toLong": expr}
}

// ToObjectID converts value to a ObjectId.
// New in version 4.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/toObjectId/
func ToObjectID(expr interface{}) M {
	return M{"$toObjectId": expr}
}

// ToString converts value to a string.
// New in version 4.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/toString/
func ToString(expr interface{}) M {
	return M{"$toString": expr}
}

// Type return the BSON data type of the field.
// New in version 3.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/type
func Type(expr interface{}) M {
	return M{"$type": expr}
}
