package query

// MergeOperators in a last win manner.
// only merge top level fields.
func MergeOperators(operators ...M) M {
	var ret = M{}
	for _, i := range operators {
		for k, v := range i {
			ret[k] = v
		}
	}
	return ret
}
