package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#miscellaneous-operators

// GetField returns the value of a specified field from a document.
// You can use $getField to retrieve the value of fields with names
// that contain periods (.) or start with dollar signs ($).
// New in version 5.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/getField/
func GetField(field string, inputExpr interface{}) M {
	return M{
		"$getField": M{
			"field": field,
			"input": inputExpr,
		},
	}
}

// Rand returns a random float between 0 and 1 each time it is called.
// New in version 4.4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/rand/
func Rand() M {
	return M{"$rand": M{}}
}

// SampleRate matches a random selection of input documents.
// The number of documents selected approximates
// the sample rate expressed as a percentage of the total number of documents.
// New in version 4.4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/sampleRate/
func SampleRate(rate float64) M {
	return M{"$sampleRate": rate}
}
