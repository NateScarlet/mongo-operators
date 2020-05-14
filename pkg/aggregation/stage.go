package aggregation

// Match https://docs.mongodb.com/manual/reference/operator/aggregation/match
func Match(query interface{}) M {
	return M{"$match": query}
}

// MatchExpr create `{ $match: { $expr: query } }``
func MatchExpr(query interface{}) M {
	return M{"$match": M{"$expr": query}}
}

// LookupF https://docs.mongodb.com/manual/reference/operator/aggregation/lookup/#equality-match
func LookupF(from, localField, foreignField, as string) M {
	return M{"$lookup": M{
		"from":         from,
		"localField":   localField,
		"foreignField": foreignField,
		"as":           as,
	}}
}

// LookupP https://docs.mongodb.com/manual/reference/operator/aggregation/lookup/#join-conditions-and-uncorrelated-sub-queries
func LookupP(from string, let M, pipeline A, as string) M {
	var m = M{
		"from":     from,
		"pipeline": pipeline,
		"as":       as,
	}
	if let != nil {
		m["let"] = let
	}
	return M{"$lookup": m}
}

// UnwindOptions https://docs.mongodb.com/manual/reference/operator/aggregation/unwind/#document-operand-with-options
type UnwindOptions M

// Unwind https://docs.mongodb.com/manual/reference/operator/aggregation/unwind
func Unwind(path string) UnwindOptions {
	if path[0] != '$' {
		path = "$" + path
	}
	return UnwindOptions{"$unwind": M{"path": path}}
}

// SetIncludeArrayIndex option
func (o UnwindOptions) SetIncludeArrayIndex(v string) UnwindOptions {
	o["$unwind"].(M)["includeArrayIndex"] = v
	return o
}

// SetPreserveNullAndEmptyArrays option
// https://docs.mongodb.com/manual/reference/operator/aggregation/unwind/#document-operand-with-options
func (o UnwindOptions) SetPreserveNullAndEmptyArrays(v bool) UnwindOptions {
	o["$unwind"].(M)["preserveNullAndEmptyArrays"] = v
	return o
}

// AddFields https://docs.mongodb.com/manual/reference/operator/aggregation/addFields
func AddFields(fields M) M {
	return M{"$addFields": fields}
}

// Unset https://docs.mongodb.com/manual/reference/operator/aggregation/unset
func Unset(field ...string) M {
	return M{"$unset": field}
}
