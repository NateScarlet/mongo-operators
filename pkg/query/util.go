package query

// Update merges multiple update operators in a last win manner.
func Update(operators ...M) M {
	var ret = M{}
	for _, i := range operators {
		for k, v := range i {
			ret[k] = v
		}
	}
	return ret
}
