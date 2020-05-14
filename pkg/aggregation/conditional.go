package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#conditional-expression-operators

// Cond A ternary operator that evaluates one expression,
// and depending on the result, returns the value of one of the other two expressions.
// https://docs.mongodb.com/manual/reference/operator/aggregation/cond/
func Cond(ifExpr, thenExpr, elseExpr interface{}) M {
	return M{"$cond": A{ifExpr, thenExpr, elseExpr}}
}

// IfNull returns either the non-null result of the first expression
// or the result of the second expression if the first expression results in a null result.
// Null result encompasses instances of undefined values or missing fields.
// https://docs.mongodb.com/manual/reference/operator/aggregation/ifNull
func IfNull(expr, replaceExprIfNull interface{}) M {
	return M{"$ifNull": A{expr, replaceExprIfNull}}
}

// Switch evaluates a series of case expressions.
// When it finds an expression which evaluates to true,
// $switch executes a specified expression and breaks out of the control flow.
// Accept args as `case, then, [case, then, ..., default]`.
// https://docs.mongodb.com/manual/reference/operator/aggregation/switch/
func Switch(args ...interface{}) M {
	var b = A{}
	var s = M{}
	for i := 0; i < len(args); i += 2 {
		if i+1 > len(args)-1 {
			s["default"] = args[i]
		} else {
			b = append(b, M{
				"case": args[i],
				"then": args[i+1],
			})
		}
	}
	s["branches"] = b
	return M{"$switch": s}
}
