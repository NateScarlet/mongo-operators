package aggregation

import "github.com/NateScarlet/mongo-operators/pkg/query"

// noop if v is a M and has key, else wrap v with M using key
func wrap(v interface{}, key string) M {
	if m, ok := v.(M); ok {
		_, ok = m[key]
		if ok {
			return m
		}
	}
	return M{key: v}
}

// MatchExpr is a shortcut for `{ $match: { $expr: expr } }``
func MatchExpr(expr interface{}) M {
	return Match(query.Expr(expr))
}

// Unless returns `expr if v else { $not : expr }`
func Unless(expr M, v bool) M {
	if v {
		return expr
	}
	return M{"$not": expr}
}

// DateTrunc trunc date to lower precision.
//
// For date=ISODate("2001-02-03T15:04:05.6Z), timezone="UTC":
//   to="year"   -> ISODate("2001-01-01T00:00:00Z")
//   to="month"  -> ISODate("2001-02-01T00:00:00Z")
//   to="day"    -> ISODate("2001-02-03T00:00:00Z")
//   to="hour"   -> ISODate("2001-02-03T15:00:00Z")
//   to="minute" -> ISODate("2001-02-03T15:04:00Z")
//   to="second" -> ISODate("2001-02-03T15:04:05Z")
func DateTrunc(to string, date interface{}, timezone string) M {
	var ret = DateFromPartsC(Year(date).SetTimezone(timezone)).
		SetTimezone(timezone)
	switch to {
	case "second":
		ret.SetSecond(Second(date).SetTimezone(timezone))
		fallthrough
	case "minute":
		ret.SetMinute(Minute(date).SetTimezone(timezone))
		fallthrough
	case "hour":
		ret.SetHour(Hour(date).SetTimezone(timezone))
		fallthrough
	case "day":
		ret.SetDay(DayOfMonth(date).SetTimezone(timezone))
		fallthrough
	case "month":
		ret.SetMonth(Month(date).SetTimezone(timezone))
	}
	return M(ret)
}
