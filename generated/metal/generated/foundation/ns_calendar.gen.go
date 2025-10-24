// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Calendar] class.
var (
	CalendarClass     _CalendarClass
	CalendarClassOnce sync.Once
)

func getCalendarClass() _CalendarClass {
	CalendarClassOnce.Do(func() {
		CalendarClass = _CalendarClass{objc.GetClass("NSCalendar")}
	})
	return CalendarClass
}

type _CalendarClass struct {
	class objc.Class
}

// An interface definition for the [Calendar] class.
type ICalendar interface {
	objectivec.IObject
	// properties:
	AmSymbol() IString
	SetAmSymbol(value IString)
	CalendarIdentifier() unsafe.Pointer
	SetCalendarIdentifier(value unsafe.Pointer)
	EraSymbols() IString
	SetEraSymbols(value IString)
	FirstWeekday() int
	SetFirstWeekday(value int)
	Locale() ILocale
	SetLocale(value ILocale)
	LongEraSymbols() IString
	SetLongEraSymbols(value IString)
	MinimumDaysInFirstWeek() int
	SetMinimumDaysInFirstWeek(value int)
	MonthSymbols() IString
	SetMonthSymbols(value IString)
	PmSymbol() IString
	SetPmSymbol(value IString)
	QuarterSymbols() IString
	SetQuarterSymbols(value IString)
	ShortMonthSymbols() IString
	SetShortMonthSymbols(value IString)
	ShortQuarterSymbols() IString
	SetShortQuarterSymbols(value IString)
	ShortStandaloneMonthSymbols() IString
	SetShortStandaloneMonthSymbols(value IString)
	ShortStandaloneQuarterSymbols() IString
	SetShortStandaloneQuarterSymbols(value IString)
	ShortStandaloneWeekdaySymbols() IString
	SetShortStandaloneWeekdaySymbols(value IString)
	ShortWeekdaySymbols() IString
	SetShortWeekdaySymbols(value IString)
	StandaloneMonthSymbols() IString
	SetStandaloneMonthSymbols(value IString)
	StandaloneQuarterSymbols() IString
	SetStandaloneQuarterSymbols(value IString)
	StandaloneWeekdaySymbols() IString
	SetStandaloneWeekdaySymbols(value IString)
	TimeZone() ITimeZone
	SetTimeZone(value ITimeZone)
	VeryShortMonthSymbols() IString
	SetVeryShortMonthSymbols(value IString)
	VeryShortStandaloneMonthSymbols() IString
	SetVeryShortStandaloneMonthSymbols(value IString)
	VeryShortStandaloneWeekdaySymbols() IString
	SetVeryShortStandaloneWeekdaySymbols(value IString)
	VeryShortWeekdaySymbols() IString
	SetVeryShortWeekdaySymbols(value IString)
	WeekdaySymbols() IString
	SetWeekdaySymbols(value IString)
	// methods:
}

// A definition of the relationships between calendar units and absolute points in time, providing features for calculation and comparison of dates.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. objects encapsulate information about systems of reckoning time in which the beginning, length, and divisions of a year are defined. They provide information about the calendar and support for calendrical computations such as determining the range of a given calendrical unit and adding units to a given absolute time. is with its Core Foundation counterpart, . See for more information on toll-free bridging.


// A definition of the relationships between calendar units and absolute points in time, providing features for calculation and comparison of dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar
type Calendar struct {
	objectivec.Object
}

// CalendarFrom constructs a [Calendar] from an unsafe.Pointer.
//
// A definition of the relationships between calendar units and absolute points in time, providing features for calculation and comparison of dates.
func CalendarFrom(ptr unsafe.Pointer) Calendar {
	return Calendar{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CalendarClass) Alloc() Calendar {
	rv := objc.Send[Calendar](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CalendarClass) New() Calendar {
	rv := objc.Send[Calendar](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Calendar) Init() Calendar {
	rv := objc.Send[Calendar](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Calendar) Autorelease() Calendar {
	rv := objc.Send[Calendar](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCalendar creates a new Calendar instance.
func NewCalendar() Calendar {
	return getCalendarClass().New()
}



// The symbol used to represent “AM” for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/amsymbol
func (c_ Calendar) AmSymbol() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("amSymbol"))
	return rv
}


// The symbol used to represent “AM” for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/amsymbol
func (c_ Calendar) SetAmSymbol(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAmSymbol:"), value)
}


// An identifier for the calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/calendaridentifier
func (c_ Calendar) CalendarIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("calendarIdentifier"))
	return rv
}


// An identifier for the calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/calendaridentifier
func (c_ Calendar) SetCalendarIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCalendarIdentifier:"), value)
}


// A list of era symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/erasymbols
func (c_ Calendar) EraSymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("eraSymbols"))
	return rv
}


// A list of era symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/erasymbols
func (c_ Calendar) SetEraSymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEraSymbols:"), value)
}


// The index of the first weekday of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/firstweekday
func (c_ Calendar) FirstWeekday() int {
	rv := objc.Send[int](c_.ID, objc.Sel("firstWeekday"))
	return rv
}


// The index of the first weekday of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/firstweekday
func (c_ Calendar) SetFirstWeekday(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFirstWeekday:"), value)
}


// The locale of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/locale
func (c_ Calendar) Locale() ILocale {
	rv := objc.Send[Locale](c_.ID, objc.Sel("locale"))
	return rv
}


// The locale of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/locale
func (c_ Calendar) SetLocale(value ILocale) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocale:"), value)
}


// A list of long era symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/longerasymbols
func (c_ Calendar) LongEraSymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("longEraSymbols"))
	return rv
}


// A list of long era symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/longerasymbols
func (c_ Calendar) SetLongEraSymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLongEraSymbols:"), value)
}


// The minimum number of days in the first week of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/minimumdaysinfirstweek
func (c_ Calendar) MinimumDaysInFirstWeek() int {
	rv := objc.Send[int](c_.ID, objc.Sel("minimumDaysInFirstWeek"))
	return rv
}


// The minimum number of days in the first week of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/minimumdaysinfirstweek
func (c_ Calendar) SetMinimumDaysInFirstWeek(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinimumDaysInFirstWeek:"), value)
}


// A list of month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/monthsymbols
func (c_ Calendar) MonthSymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("monthSymbols"))
	return rv
}


// A list of month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/monthsymbols
func (c_ Calendar) SetMonthSymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMonthSymbols:"), value)
}


// The symbol used to represent “PM” for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/pmsymbol
func (c_ Calendar) PmSymbol() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("pmSymbol"))
	return rv
}


// The symbol used to represent “PM” for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/pmsymbol
func (c_ Calendar) SetPmSymbol(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPmSymbol:"), value)
}


// A list of quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/quartersymbols
func (c_ Calendar) QuarterSymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("quarterSymbols"))
	return rv
}


// A list of quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/quartersymbols
func (c_ Calendar) SetQuarterSymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQuarterSymbols:"), value)
}


// A list of short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortmonthsymbols
func (c_ Calendar) ShortMonthSymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("shortMonthSymbols"))
	return rv
}


// A list of short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortmonthsymbols
func (c_ Calendar) SetShortMonthSymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShortMonthSymbols:"), value)
}


// A list of short quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortquartersymbols
func (c_ Calendar) ShortQuarterSymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("shortQuarterSymbols"))
	return rv
}


// A list of short quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortquartersymbols
func (c_ Calendar) SetShortQuarterSymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShortQuarterSymbols:"), value)
}


// A list of short standalone month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortstandalonemonthsymbols
func (c_ Calendar) ShortStandaloneMonthSymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("shortStandaloneMonthSymbols"))
	return rv
}


// A list of short standalone month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortstandalonemonthsymbols
func (c_ Calendar) SetShortStandaloneMonthSymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShortStandaloneMonthSymbols:"), value)
}


// A list of short standalone quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortstandalonequartersymbols
func (c_ Calendar) ShortStandaloneQuarterSymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("shortStandaloneQuarterSymbols"))
	return rv
}


// A list of short standalone quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortstandalonequartersymbols
func (c_ Calendar) SetShortStandaloneQuarterSymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShortStandaloneQuarterSymbols:"), value)
}


// A list of short standalone weekday symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortstandaloneweekdaysymbols
func (c_ Calendar) ShortStandaloneWeekdaySymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("shortStandaloneWeekdaySymbols"))
	return rv
}


// A list of short standalone weekday symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortstandaloneweekdaysymbols
func (c_ Calendar) SetShortStandaloneWeekdaySymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShortStandaloneWeekdaySymbols:"), value)
}


// A list of shorter-named weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortweekdaysymbols
func (c_ Calendar) ShortWeekdaySymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("shortWeekdaySymbols"))
	return rv
}


// A list of shorter-named weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/shortweekdaysymbols
func (c_ Calendar) SetShortWeekdaySymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShortWeekdaySymbols:"), value)
}


// A list of standalone month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standalonemonthsymbols
func (c_ Calendar) StandaloneMonthSymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("standaloneMonthSymbols"))
	return rv
}


// A list of standalone month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standalonemonthsymbols
func (c_ Calendar) SetStandaloneMonthSymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStandaloneMonthSymbols:"), value)
}


// A list of standalone quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standalonequartersymbols
func (c_ Calendar) StandaloneQuarterSymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("standaloneQuarterSymbols"))
	return rv
}


// A list of standalone quarter symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standalonequartersymbols
func (c_ Calendar) SetStandaloneQuarterSymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStandaloneQuarterSymbols:"), value)
}


// A list of standalone weekday symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standaloneweekdaysymbols
func (c_ Calendar) StandaloneWeekdaySymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("standaloneWeekdaySymbols"))
	return rv
}


// A list of standalone weekday symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/standaloneweekdaysymbols
func (c_ Calendar) SetStandaloneWeekdaySymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStandaloneWeekdaySymbols:"), value)
}


// The time zone for the calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/timezone
func (c_ Calendar) TimeZone() ITimeZone {
	rv := objc.Send[TimeZone](c_.ID, objc.Sel("timeZone"))
	return rv
}


// The time zone for the calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/timezone
func (c_ Calendar) SetTimeZone(value ITimeZone) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeZone:"), value)
}


// A list of very short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortmonthsymbols
func (c_ Calendar) VeryShortMonthSymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("veryShortMonthSymbols"))
	return rv
}


// A list of very short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortmonthsymbols
func (c_ Calendar) SetVeryShortMonthSymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVeryShortMonthSymbols:"), value)
}


// A list of very short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortstandalonemonthsymbols
func (c_ Calendar) VeryShortStandaloneMonthSymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("veryShortStandaloneMonthSymbols"))
	return rv
}


// A list of very short month symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortstandalonemonthsymbols
func (c_ Calendar) SetVeryShortStandaloneMonthSymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVeryShortStandaloneMonthSymbols:"), value)
}


// A list of very short standalone weekday symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortstandaloneweekdaysymbols
func (c_ Calendar) VeryShortStandaloneWeekdaySymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("veryShortStandaloneWeekdaySymbols"))
	return rv
}


// A list of very short standalone weekday symbols for this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortstandaloneweekdaysymbols
func (c_ Calendar) SetVeryShortStandaloneWeekdaySymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVeryShortStandaloneWeekdaySymbols:"), value)
}


// A list of very-shortly-named weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortweekdaysymbols
func (c_ Calendar) VeryShortWeekdaySymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("veryShortWeekdaySymbols"))
	return rv
}


// A list of very-shortly-named weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/veryshortweekdaysymbols
func (c_ Calendar) SetVeryShortWeekdaySymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVeryShortWeekdaySymbols:"), value)
}


// A list of weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/weekdaysymbols
func (c_ Calendar) WeekdaySymbols() IString {
	rv := objc.Send[String](c_.ID, objc.Sel("weekdaySymbols"))
	return rv
}


// A list of weekdays in this calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscalendar/weekdaysymbols
func (c_ Calendar) SetWeekdaySymbols(value IString) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeekdaySymbols:"), value)
}



