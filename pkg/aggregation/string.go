package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#string-expression-operators

// Concat any number of strings.
// https://docs.mongodb.com/manual/reference/operator/aggregation/concat/
func Concat(stringExpr ...interface{}) M {
	return M{"$concat": stringExpr}
}

// IndexOfBytesOperator returned from IndexOfBytes
type IndexOfBytesOperator M

// IndexOfBytes searches a string for an occurence of a substring
// and returns the UTF-8 byte index of the first occurence.
// If the substring is not found, returns -1.
// New in version 3.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/indexOfBytes/
func IndexOfBytes(stringExpr, substringExpr interface{}) IndexOfBytesOperator {
	return IndexOfBytesOperator{"$indexOfBytes": A{stringExpr, substringExpr}}
}

// IndexOfCPOperator returned from IndexOfCP
type IndexOfCPOperator M

// IndexOfCP searches a string for an occurence of a substring
// and returns the UTF-8 code point index of the first occurence.
// If the substring is not found, returns -1
// https://docs.mongodb.com/manual/reference/operator/aggregation/indexOfCP/
func IndexOfCP(stringExpr, subStringExpr interface{}) IndexOfCPOperator {
	return IndexOfCPOperator{"$indexOfCP": A{stringExpr, subStringExpr}}
}

// SetStart expression
func (op IndexOfCPOperator) SetStart(expr interface{}) IndexOfCPOperator {
	var a = op["$indexOfCP"].(A)
	if len(a) <= 3 {
		op["$indexOfCP"] = append(a, expr)
	} else {
		a[2] = expr
	}
	return op
}

// SetEnd expression
func (op IndexOfCPOperator) SetEnd(expr interface{}) IndexOfCPOperator {
	var a = op["$indexOfCP"].(A)
	if len(a) <= 4 {
		op["$indexOfCP"] = append(a, expr)
	} else {
		a[3] = expr
	}
	return op
}

// LTrimOperator returned from LTrim
type LTrimOperator M

// LTrim removes whitespace characters, including null,
// or the specified characters from the beginning of a string.
// New in version 4.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/ltrim/
func LTrim(input interface{}) LTrimOperator {
	return LTrimOperator{"$ltrim": M{"input": input}}
}

// SetChars option
func (op LTrimOperator) SetChars(expr interface{}) LTrimOperator {
	op["$ltrim"].(M)["chars"] = expr
	return op
}

// RegexFindOperator returned from RegexFind
type RegexFindOperator M

// RegexFind applies a regular expression (regex) to a string
// and returns information on the first matched substring.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/regexFind/
func RegexFind(input, regex interface{}) RegexFindOperator {
	return RegexFindOperator{"$regexFind": M{
		"input": input,
		"regex": regex,
	}}
}

// SetOptions option
func (op RegexFindOperator) SetOptions(expr interface{}) RegexFindOperator {
	op["$regexFind"].(M)["options"] = expr
	return op
}

// RegexFindAllOperator returned from RegexFindAll
type RegexFindAllOperator M

// RegexFindAll Applies a regular expression (regex) to a string
// and returns information on the all matched substrings.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/regexFindAll/
func RegexFindAll(input, regex interface{}) RegexFindAllOperator {
	return RegexFindAllOperator{"$regexFindAll": M{
		"input": input,
		"regex": regex,
	}}
}

// SetOptions option
func (op RegexFindAllOperator) SetOptions(expr interface{}) RegexFindAllOperator {
	op["$regexFindAll"].(M)["options"] = expr
	return op
}

// RegexMatchOperator returned from RegexMatch
type RegexMatchOperator M

// RegexMatch applies a regular expression (regex) to a string
// and returns a boolean that indicates if a match is found or not.
// New in version 4.2.
// https://docs.mongodb.com/manual/reference/operator/aggregation/regexMatch/
func RegexMatch(input, regex interface{}) RegexMatchOperator {
	return RegexMatchOperator{"$regexMatch": M{
		"input": input,
		"regex": regex,
	}}
}

// SetOptions option
func (op RegexMatchOperator) SetOptions(expr interface{}) RegexMatchOperator {
	op["$regexMatch"].(M)["options"] = expr
	return op
}

// RTrimOperator returned from RTrim
type RTrimOperator M

// RTrim removes whitespace or the specified characters from the end of a string.
// New in version 4.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/rtrim/
func RTrim(input interface{}) RTrimOperator {
	return RTrimOperator{"$rtrim": M{"input": input}}
}

// SetChars option
func (op RTrimOperator) SetChars(expr interface{}) RTrimOperator {
	op["$rtrim"].(M)["chars"] = expr
	return op
}

// Split a string into substrings based on a delimiter.
// Returns an array of substrings. If the delimiter is not found within the string,
// returns an array containing the original string.
// https://docs.mongodb.com/manual/reference/operator/aggregation/split/
func Split(stringExpr, delimiter interface{}) M {
	return M{"$split": A{stringExpr, delimiter}}
}

// StrLenBytes returns the number of UTF-8 encoded bytes in a string.
// New in version 3.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/strLenBytes/
func StrLenBytes(stringExpr interface{}) M {
	return M{"$strLenBytes": stringExpr}
}

// StrLenCP returns the number of UTF-8 code points in a string.
// New in version 3.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/strLenCP/
func StrLenCP(stringExpr interface{}) M {
	return M{"$strLenCP": stringExpr}
}

// StrCaseCmp performs case-insensitive string comparison
// and returns: 0 if two strings are equivalent,
// 1 if the first string is greater than the second,
// and -1 if the first string is less than the second.
// https://docs.mongodb.com/manual/reference/operator/aggregation/strcasecmp/
func StrCaseCmp(string1, string2 interface{}) M {
	return M{"$strcasecmp": A{string1, string2}}
}

// SubstrBytes returns the substring of a string.
// Starts with the character at the specified UTF-8 byte index (zero-based)
// in the string and continues for the specified number of bytes.
// New in version 3.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/substrBytes/
func SubstrBytes(stringExpr, startIndex, count interface{}) M {
	return M{"$substrBytes": A{stringExpr, startIndex, count}}
}

// SubstrCP returns the substring of a string.
// Starts with the character at the specified UTF-8 code point (CP) index (zero-based)
// in the string and continues for the number of code points specified.
// https://docs.mongodb.com/manual/reference/operator/aggregation/substrCP/
func SubstrCP(stringExpr, startIndex, count interface{}) M {
	return M{"$substrCP": A{stringExpr, startIndex, count}}
}

// ToLower converts a string to lowercase.
// https://docs.mongodb.com/manual/reference/operator/aggregation/toLower/
func ToLower(stringExpr interface{}) M {
	return M{"$toLower": stringExpr}
}

// TrimOperator returned from Trim
type TrimOperator M

// Trim removes whitespace or the specified characters
// from the beginning and end of a string.
// New in version 4.0.
// https://docs.mongodb.com/manual/reference/operator/aggregation/trim/
func Trim(input interface{}) TrimOperator {
	return TrimOperator{"$trim": M{"input": input}}
}

// SetChars option
func (op TrimOperator) SetChars(expr interface{}) TrimOperator {
	op["$trim"].(M)["chars"] = expr
	return op
}

// ToUpper converts a string to uppercase.
// https://docs.mongodb.com/manual/reference/operator/aggregation/toUpper/
func ToUpper(stringExpr interface{}) M {
	return M{"$toUpper": stringExpr}
}
