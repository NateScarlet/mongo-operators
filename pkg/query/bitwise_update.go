package query

// https://docs.mongodb.com/manual/reference/operator/update-bitwise/

// Bit performs a bitwise update of a field.
// The operator supports bitwise and, bitwise or,
// and bitwise xor (i.e. exclusive or) operations.
// https://docs.mongodb.com/manual/reference/operator/update/bit/
func Bit(spec interface{}) M {
	return M{"$bit": spec}
}
