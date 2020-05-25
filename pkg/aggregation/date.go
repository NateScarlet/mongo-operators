package aggregation

// https://docs.mongodb.com/manual/reference/operator/aggregation/#date-expression-operators

// DateFromPartsCOperator returned from DateFromParts
type DateFromPartsCOperator M

// DateFromPartsC constructs and returns a Date object
// given the calendar date’s constituent properties.
// https://docs.mongodb.com/manual/reference/operator/aggregation/dateFromParts/
func DateFromPartsC(year interface{}) DateFromPartsCOperator {
	return DateFromPartsCOperator{"$dateFromParts": M{"year": year}}
}

// SetMonth option
func (op DateFromPartsCOperator) SetMonth(v interface{}) DateFromPartsCOperator {
	op["$dateFromParts"].(M)["month"] = v
	return op
}

// SetDay option
func (op DateFromPartsCOperator) SetDay(v interface{}) DateFromPartsCOperator {
	op["$dateFromParts"].(M)["day"] = v
	return op
}

// SetHour option
func (op DateFromPartsCOperator) SetHour(v interface{}) DateFromPartsCOperator {
	op["$dateFromParts"].(M)["hour"] = v
	return op
}

// SetMinute option
func (op DateFromPartsCOperator) SetMinute(v interface{}) DateFromPartsCOperator {
	op["$dateFromParts"].(M)["minute"] = v
	return op
}

// SetSecond option
func (op DateFromPartsCOperator) SetSecond(v interface{}) DateFromPartsCOperator {
	op["$dateFromParts"].(M)["second"] = v
	return op
}

// SetMillisecond option
func (op DateFromPartsCOperator) SetMillisecond(v interface{}) DateFromPartsCOperator {
	op["$dateFromParts"].(M)["millisecond"] = v
	return op
}

// SetTimezone option
func (op DateFromPartsCOperator) SetTimezone(v interface{}) DateFromPartsCOperator {
	op["$dateFromParts"].(M)["timezone"] = v
	return op
}

// DateFromPartsWOperator returned from DateFromWeekParts
type DateFromPartsWOperator M

// DateFromPartsW constructs and returns a Date object
// given the ISO week date’s constituent properties.
// https://docs.mongodb.com/manual/reference/operator/aggregation/dateFromParts/
func DateFromPartsW(year interface{}) DateFromPartsWOperator {
	return DateFromPartsWOperator{"$dateFromParts": M{"isoWeekYear": year}}
}

// SetWeek option
func (op DateFromPartsWOperator) SetWeek(v interface{}) DateFromPartsWOperator {
	op["$dateFromParts"].(M)["isoWeek"] = v
	return op
}

// SetDayOfWeek option
func (op DateFromPartsWOperator) SetDayOfWeek(v interface{}) DateFromPartsWOperator {
	op["$dateFromParts"].(M)["isoDayOfWeek"] = v
	return op
}

// SetHour option
func (op DateFromPartsWOperator) SetHour(v interface{}) DateFromPartsWOperator {
	op["$dateFromParts"].(M)["hour"] = v
	return op
}

// SetMinute option
func (op DateFromPartsWOperator) SetMinute(v interface{}) DateFromPartsWOperator {
	op["$dateFromParts"].(M)["minute"] = v
	return op
}

// SetSecond option
func (op DateFromPartsWOperator) SetSecond(v interface{}) DateFromPartsWOperator {
	op["$dateFromParts"].(M)["second"] = v
	return op
}

// SetMillisecond option
func (op DateFromPartsWOperator) SetMillisecond(v interface{}) DateFromPartsWOperator {
	op["$dateFromParts"].(M)["millisecond"] = v
	return op
}

// SetTimezone option
func (op DateFromPartsWOperator) SetTimezone(v interface{}) DateFromPartsWOperator {
	op["$dateFromParts"].(M)["timezone"] = v
	return op
}

// DateFromStringOperator returned from DateFromString
type DateFromStringOperator M

// DateFromString converts a date/time string to a date object.
// New in version 3.6.
// https://docs.mongodb.com/manual/reference/operator/aggregation/dateFromString/
func DateFromString(dateString interface{}) DateFromStringOperator {
	return DateFromStringOperator{"$dateFromString": M{"dateString": dateString}}
}

// SetFormat option
func (op DateFromStringOperator) SetFormat(v interface{}) DateFromStringOperator {
	op["$dateFromString"].(M)["format"] = v
	return op
}

// SetTimezone option
func (op DateFromStringOperator) SetTimezone(v interface{}) DateFromStringOperator {
	op["$dateFromString"].(M)["timezone"] = v
	return op
}

// SetOnError option
func (op DateFromStringOperator) SetOnError(v interface{}) DateFromStringOperator {
	op["$dateFromString"].(M)["onError"] = v
	return op
}

// SetOnNull option
func (op DateFromStringOperator) SetOnNull(v interface{}) DateFromStringOperator {
	op["$dateFromString"].(M)["onNull"] = v
	return op
}

// DateToPartsOperator returned from DateToParts
type DateToPartsOperator M

// DateToParts returns a document containing the constituent parts of a date.
// New in version 3.6.
// https://docs.mongodb.com/manual/reference/operator/aggregation/dateToParts/
func DateToParts(date interface{}) DateToPartsOperator {
	return DateToPartsOperator{"$dateToParts": M{"date": date}}
}

// SetTimezone option
func (op DateToPartsOperator) SetTimezone(v interface{}) DateToPartsOperator {
	op["$dateToParts"].(M)["timezone"] = v
	return op
}

// SetISO8601 option
func (op DateToPartsOperator) SetISO8601(v bool) DateToPartsOperator {
	op["$dateToParts"].(M)["iso8601"] = v
	return op
}

// DateToStringOperator returned from DateToString
type DateToStringOperator M

// DateToString converts a date object to a string according to a user-specified format.
// https://docs.mongodb.com/manual/reference/operator/aggregation/dateToString/
func DateToString(date interface{}) DateToStringOperator {
	return DateToStringOperator{"$dateToString": M{"date": date}}
}

// SetFormat option
// required before version 4.0
func (op DateToStringOperator) SetFormat(v string) DateToStringOperator {
	op["$dateToString"].(M)["format"] = v
	return op
}

// SetTimezone option
func (op DateToStringOperator) SetTimezone(v interface{}) DateToStringOperator {
	op["$dateToString"].(M)["timezone"] = v
	return op
}

// DayOfMonthOperator returned from DayOfMonth
type DayOfMonthOperator M

// DayOfMonth returns the day of the month for a date as a number between 1 and 31.
// https://docs.mongodb.com/manual/reference/operator/aggregation/dayOfMonth/
func DayOfMonth(date interface{}) DayOfMonthOperator {
	return DayOfMonthOperator{"$dayOfMonth": date}
}

// SetTimezone option
// New in version 3.6.
func (op DayOfMonthOperator) SetTimezone(v interface{}) DayOfMonthOperator {
	op["$dayOfMonth"] = wrap(op["$dayOfMonth"], "date")
	op["$dayOfMonth"].(M)["timezone"] = v
	return op
}

// DayOfWeekOperator returned from DayOfMonth
type DayOfWeekOperator M

// DayOfWeek returns the day of the week for a date as a number between 1 (Sunday) and 7 (Saturday).
// https://docs.mongodb.com/manual/reference/operator/aggregation/dayOfWeek/
func DayOfWeek(date interface{}) DayOfWeekOperator {
	return DayOfWeekOperator{"$dayOfWeek": date}
}

// SetTimezone option
// New in version 3.6.
func (op DayOfWeekOperator) SetTimezone(v interface{}) DayOfWeekOperator {
	op["$dayOfWeek"] = wrap(op["$dayOfWeek"], "date")
	op["$dayOfWeek"].(M)["timezone"] = v
	return op
}

// DayOfYearOperator returned from DayOfYear
type DayOfYearOperator M

// DayOfYear returns the day of the week for a date as a number between 1 (Sunday) and 7 (Saturday).
// https://docs.mongodb.com/manual/reference/operator/aggregation/dayOfYear/
func DayOfYear(date interface{}) DayOfYearOperator {
	return DayOfYearOperator{"$dayOfYear": date}
}

// SetTimezone option
// New in version 3.6.
func (op DayOfYearOperator) SetTimezone(v interface{}) DayOfYearOperator {
	op["$dayOfYear"] = wrap(op["$dayOfYear"], "date")
	op["$dayOfYear"].(M)["timezone"] = v
	return op
}

// HourOperator returned from Hour
type HourOperator M

// Hour returns the hour for a date as a number between 0 and 23.
// https://docs.mongodb.com/manual/reference/operator/aggregation/hour/
func Hour(date interface{}) HourOperator {
	return HourOperator{"$hour": date}
}

// SetTimezone option
// New in version 3.6.
func (op HourOperator) SetTimezone(v interface{}) HourOperator {
	op["$hour"] = wrap(op["$hour"], "date")
	op["$hour"].(M)["timezone"] = v
	return op
}

// ISODayOfWeekOperator returned from ISODayOfWeek
type ISODayOfWeekOperator M

// ISODayOfWeek returns the weekday number in ISO 8601 format, ranging from 1 (for Monday) to 7 (for Sunday).
// New in version 3.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/isoDayOfWeek/
func ISODayOfWeek(date interface{}) ISODayOfWeekOperator {
	return ISODayOfWeekOperator{"$isoDayOfWeek": date}
}

// SetTimezone option
// New in version 3.6.
func (op ISODayOfWeekOperator) SetTimezone(v interface{}) ISODayOfWeekOperator {
	op["$isoDayOfWeek"] = wrap(op["$isoDayOfWeek"], "date")
	op["$isoDayOfWeek"].(M)["timezone"] = v
	return op
}

// ISOWeekOperator returned from ISOWeek
type ISOWeekOperator M

// ISOWeek returns the week number in ISO 8601 format, ranging from 1 to 53.
// Week numbers start at 1 with the week (Monday through Sunday)
// that contains the year’s first Thursday.
// New in version 3.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/isoWeek/
func ISOWeek(date interface{}) ISOWeekOperator {
	return ISOWeekOperator{"$isoWeek": date}
}

// SetTimezone option
// New in version 3.6.
func (op ISOWeekOperator) SetTimezone(v interface{}) ISOWeekOperator {
	op["$isoWeek"] = wrap(op["$isoWeek"], "date")
	op["$isoWeek"].(M)["timezone"] = v
	return op
}

// ISOWeekYearOperator returned from ISOWeek
type ISOWeekYearOperator M

// ISOWeekYear returns the year number in ISO 8601 format.
// The year starts with the Monday of week 1
// and ends with the Sunday of the last week.
// New in version 3.4.
// https://docs.mongodb.com/manual/reference/operator/aggregation/isoWeekYear/
func ISOWeekYear(date interface{}) ISOWeekYearOperator {
	return ISOWeekYearOperator{"$isoWeekYear": date}
}

// SetTimezone option
// New in version 3.6.
func (op ISOWeekYearOperator) SetTimezone(v interface{}) ISOWeekYearOperator {
	op["$isoWeekYear"] = wrap(op["$isoWeekYear"], "date")
	op["$isoWeekYear"].(M)["timezone"] = v
	return op
}

// MillisecondOperator returned from Millisecond
type MillisecondOperator M

// Millisecond returns the milliseconds of a date as a number between 0 and 999.
// https://docs.mongodb.com/manual/reference/operator/aggregation/millisecond/
func Millisecond(date interface{}) MillisecondOperator {
	return MillisecondOperator{"$millisecond": date}
}

// SetTimezone option
// New in version 3.6.
func (op MillisecondOperator) SetTimezone(v interface{}) MillisecondOperator {
	op["$millisecond"] = wrap(op["$millisecond"], "date")
	op["$millisecond"].(M)["timezone"] = v
	return op
}

// MinuteOperator returned from Minute
type MinuteOperator M

// Minute returns the minute for a date as a number between 0 and 59.
// https://docs.mongodb.com/manual/reference/operator/aggregation/minute/
func Minute(date interface{}) MinuteOperator {
	return MinuteOperator{"$minute": date}
}

// SetTimezone option
// New in version 3.6.
func (op MinuteOperator) SetTimezone(v interface{}) MinuteOperator {
	op["$minute"] = wrap(op["$minute"], "date")
	op["$minute"].(M)["timezone"] = v
	return op
}

// MonthOperator returned from Month
type MonthOperator M

// Month returns the month for a date as a number between 1 (January) and 12 (December).
// https://docs.mongodb.com/manual/reference/operator/aggregation/month/
func Month(date interface{}) MonthOperator {
	return MonthOperator{"$month": date}
}

// SetTimezone option
// New in version 3.6.
func (op MonthOperator) SetTimezone(v interface{}) MonthOperator {
	op["$month"] = wrap(op["$month"], "date")
	op["$month"].(M)["timezone"] = v
	return op
}

// SecondOperator returned from Second
type SecondOperator M

// Second returns the seconds for a date as a number between 0 and 60 (leap seconds).
// https://docs.mongodb.com/manual/reference/operator/aggregation/second/
func Second(date interface{}) SecondOperator {
	return SecondOperator{"$second": date}
}

// SetTimezone option
// New in version 3.6.
func (op SecondOperator) SetTimezone(v interface{}) SecondOperator {
	op["$second"] = wrap(op["$second"], "date")
	op["$second"].(M)["timezone"] = v
	return op
}

// WeekOperator returned from Week
type WeekOperator M

// Week returns the week of the year for a date as a number between 0 and 53.
// https://docs.mongodb.com/manual/reference/operator/aggregation/week/
func Week(date interface{}) WeekOperator {
	return WeekOperator{"$week": date}
}

// SetTimezone option
// New in version 3.6.
func (op WeekOperator) SetTimezone(v interface{}) WeekOperator {
	op["$week"] = wrap(op["$week"], "date")
	op["$week"].(M)["timezone"] = v
	return op
}

// YearOperator returned from Year
type YearOperator M

// Year returns the year for a date as a number (e.g. 2014).
// https://docs.mongodb.com/manual/reference/operator/aggregation/year/
func Year(date interface{}) YearOperator {
	return YearOperator{"$year": date}
}

// SetTimezone option
// New in version 3.6.
func (op YearOperator) SetTimezone(v interface{}) YearOperator {
	op["$year"] = wrap(op["$year"], "date")
	op["$year"].(M)["timezone"] = v
	return op
}
