package query

// https://docs.mongodb.com/manual/reference/operator/update-field/

// CurrentDate sets the value of a field to current date, either as a Date or a Timestamp.
// https://docs.mongodb.com/manual/reference/operator/update/currentDate/
func CurrentDate(spec interface{}) M {
	return M{"$currentDate": spec}
}

// Inc increments the value of the field by the specified amount.
// https://docs.mongodb.com/manual/reference/operator/update/inc/
func Inc(spec interface{}) M {
	return M{"$inc": spec}
}

// Min only updates the field if the specified value is less than the existing field value.
// https://docs.mongodb.com/manual/reference/operator/update/min/
func Min(spec interface{}) M {
	return M{"$min": spec}
}

// Max only updates the field if the specified value is greater than the existing field value.
// https://docs.mongodb.com/manual/reference/operator/update/max/
func Max(spec interface{}) M {
	return M{"$max": spec}
}

// Mul multiplies the value of the field by the specified amount.
// https://docs.mongodb.com/manual/reference/operator/update/mul/
func Mul(spec interface{}) M {
	return M{"$mul": spec}
}

// Rename updates the name of a field.
// https://docs.mongodb.com/manual/reference/operator/update/rename/
func Rename(spec interface{}) M {
	return M{"$rename": spec}
}

// Set the value of a field in a document.
// https://docs.mongodb.com/manual/reference/operator/update/set/
func Set(spec interface{}) M {
	return M{"$set": spec}
}

// SetOnInsert sets the value of a field if an update results
// in an insert of a document.
// Has no effect on update operations that modify existing documents.
// https://docs.mongodb.com/manual/reference/operator/update/setOnInsert/
func SetOnInsert(spec interface{}) M {
	return M{"$setOnInsert": spec}
}

// Unset removes the specified field from a document.
// https://docs.mongodb.com/manual/reference/operator/update/unset/
func Unset(spec interface{}) M {
	return M{"$unset": spec}
}
