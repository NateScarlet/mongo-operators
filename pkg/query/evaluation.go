package query

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// https://docs.mongodb.com/manual/reference/operator/query-evaluation/

// Expr allows use of aggregation expressions within the query language.
// https://docs.mongodb.com/manual/reference/operator/query/expr/
func Expr(expr interface{}) M {
	return M{"$expr": expr}
}

// JSONSchema validate documents against the given JSON Schema.
// https://docs.mongodb.com/manual/reference/operator/query/jsonSchema/
func JSONSchema(schema interface{}) M {
	return M{"$jsonSchema": schema}
}

// Mod performs a modulo operation on the value of a field
// and selects documents with a specified result.
// https://docs.mongodb.com/manual/reference/operator/query/mod/
func Mod(divisor, remainder int64) M {
	return M{"$mod": A{divisor, remainder}}
}

// Regex selects documents where values match a specified regular expression.
// https://docs.mongodb.com/manual/reference/operator/query/regex/
func Regex(re primitive.Regex) M {
	return M{"$regex": re}
}

// TextOperator returned by Text
type TextOperator M

// Text performs text search.
// https://docs.mongodb.com/manual/reference/operator/query/text/
func Text(search string) TextOperator {
	return TextOperator{"$text": M{
		"$search": search,
	}}
}

// SetLanguage option
func (op TextOperator) SetLanguage(v string) TextOperator {
	op["$search"].(M)["$language"] = v
	return op
}

// SetCaseSensitive option
func (op TextOperator) SetCaseSensitive(v bool) TextOperator {
	op["$search"].(M)["$caseSensitive"] = v
	return op
}

// SetDiacriticSensitive option
func (op TextOperator) SetDiacriticSensitive(v bool) TextOperator {
	op["$search"].(M)["$diacriticSensitive"] = v
	return op
}

// Where matches documents that satisfy a JavaScript expression.
func Where(js primitive.JavaScript) M {
	return M{"$where": js}
}
