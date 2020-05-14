package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#set-expression-operators

// AllElementsTrue returns true if no element of a set evaluates to false,
// otherwise, returns false.
// https://docs.mongodb.com/manual/reference/operator/aggregation/allElementsTrue/
func AllElementsTrue(set interface{}) M {
	return M{"$allElementsTrue": A{set}}
}

// AnyElementsTrue returns true if any elements of a set evaluate to true;
// otherwise, returns false.
// https://docs.mongodb.com/manual/reference/operator/aggregation/anyElementTrue/
func AnyElementsTrue(set interface{}) M {
	return M{"$anyElementTrue": A{set}}
}

// SetDifference returns a set with elements that appear in the first set
// but not in the second set; i.e. performs a relative complement of
// the second set relative to the first.
// https://docs.mongodb.com/manual/reference/operator/aggregation/setDifference/
func SetDifference(set1, set2 interface{}) M {
	return M{"$setDifference": A{set1, set2}}
}

// SetEquals returns true if the input sets have the same distinct elements.
// Accepts two or more argument expressions.
// https://docs.mongodb.com/manual/reference/operator/aggregation/setEquals/
func SetEquals(set ...interface{}) M {
	return M{"$setEquals": set}
}

// SetIntersection returns a set with elements that appear in all
// of the input sets. Accepts any number of argument expressions.
// https://docs.mongodb.com/manual/reference/operator/aggregation/setIntersection
func SetIntersection(set ...interface{}) M {
	return M{"$setIntersection": set}
}

// SetIsSubset returns true if all elements of the first set
// appear in the second set, including when the first set equals the second set;
// i.e. not a strict subset.
// https://docs.mongodb.com/manual/reference/operator/aggregation/setIsSubset/
func SetIsSubset(set1, set2 interface{}) M {
	return M{"$setIsSubset": A{set1, set2}}
}

// SetUnion returns a set with elements that appear in any of the input sets.
// https://docs.mongodb.com/manual/reference/operator/aggregation/setUnion/
func SetUnion(set ...interface{}) M {
	return M{"$setUnion": set}
}
