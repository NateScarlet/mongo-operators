package aggregation

// SetField adds, updates, or removes a specified field in a document.
// New in version 5.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/setField/
func SetField(field string, inputExpr interface{}, valueExpr interface{}) M {
	return M{
		"$setField": M{
			"field": field,
			"input": inputExpr,
			"value": valueExpr,
		},
	}
}

// UnsetField removes a specified field in a document.
// New in version 5.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/unsetField/
func UnsetField(field string, inputExpr interface{}) M {
	return M{
		"$unsetField": M{
			"field": field,
			"input": inputExpr,
		},
	}
}
