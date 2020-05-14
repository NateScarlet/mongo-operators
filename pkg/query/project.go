package query

// https://docs.mongodb.com/manual/reference/operator/projection/

// FirstElemMatch the first element in an array
// that matches the specified $elemMatch condition
// https://docs.mongodb.com/manual/reference/operator/projection/elemMatch/
func FirstElemMatch(spec interface{}) M {
	return M{"$elemMatch": spec}
}

// Meta projects the document’s score assigned during $text operation..
// https://docs.mongodb.com/manual/reference/operator/projection/meta/
func Meta(metaDataKeyword string) M {
	return M{"$meta": metaDataKeyword}
}

// SliceN limits the number of elements projected from an array
// https://docs.mongodb.com/manual/reference/operator/projection/slice/
func SliceN(num int) M {
	return M{"$slice": num}
}

// Slice limits the number of elements projected from an array.
// use skip and limit syntax.
// https://docs.mongodb.com/manual/reference/operator/projection/slice/
func SliceSkip(skip, limit int) M {
	return M{"$slice": A{skip, limit}}
}
