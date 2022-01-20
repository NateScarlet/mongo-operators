package aggregation

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// https://docs.mongodb.com/manual/reference/operator/aggregation/#custom-aggregation-expression-operators

type AccumulatorOperator M

// Accumulator defines a custom accumulator function.
// New in version 4.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/accumulator/
func Accumulator(
	init primitive.JavaScript,
	initArgs interface{},
	accumulate primitive.JavaScript,
	accumulateArgs interface{},
	merge primitive.JavaScript,
	finalize primitive.JavaScript,
) M {
	return M{
		"$accumulator": M{
			"init":           init,
			"initArgs":       initArgs,
			"accumulate":     accumulate,
			"accumulateArgs": accumulateArgs,
			"merge":          merge,
			"finalize":       finalize,
			"lang":           "js",
		},
	}
}

// Function defines a custom function.
// New in version 4.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/function/
func Function(
	body primitive.JavaScript,
	args interface{},

) M {
	return M{
		"$function": M{
			"body": body,
			"args": args,
			"lang": "js",
		},
	}
}
