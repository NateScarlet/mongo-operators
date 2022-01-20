package query

// Comment Adds a comment to a query predicate.
// https://docs.mongodb.com/manual/reference/operator/query/comment/
func Comment(comment string) M {
	return M{"$comment": comment}
}

// Rand generates a random float between 0 and 1.
// https://docs.mongodb.com/manual/reference/operator/query/rand/
func Rand() M {
	return M{"$rand": M{}}
}
