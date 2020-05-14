package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#text-expression-operator

// Meta access text search metadata.
// https://docs.mongodb.com/manual/reference/operator/aggregation/meta/
func Meta(metaDataKeyword string) M {
	return M{"$meta": metaDataKeyword}
}
