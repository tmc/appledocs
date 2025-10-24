// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CalendarDate] class.
var (
	CalendarDateClass     _CalendarDateClass
	CalendarDateClassOnce sync.Once
)

func getCalendarDateClass() _CalendarDateClass {
	CalendarDateClassOnce.Do(func() {
		CalendarDateClass = _CalendarDateClass{objc.GetClass("NSCalendarDate")}
	})
	return CalendarDateClass
}

type _CalendarDateClass struct {
	class objc.Class
}

// An interface definition for the [CalendarDate] class.
type ICalendarDate interface {
	IDate
	// properties:
	// methods:
}

// A specialized date object with embedded calendar information.


// A specialized date object with embedded calendar information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate
type CalendarDate struct {
	Date
}

// CalendarDateFrom constructs a [CalendarDate] from an unsafe.Pointer.
//
// A specialized date object with embedded calendar information.
func CalendarDateFrom(ptr unsafe.Pointer) CalendarDate {
	return CalendarDate{
		Date: DateFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CalendarDateClass) Alloc() CalendarDate {
	rv := objc.Send[CalendarDate](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CalendarDateClass) New() CalendarDate {
	rv := objc.Send[CalendarDate](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CalendarDate) Init() CalendarDate {
	rv := objc.Send[CalendarDate](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CalendarDate) Autorelease() CalendarDate {
	rv := objc.Send[CalendarDate](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCalendarDate creates a new CalendarDate instance.
func NewCalendarDate() CalendarDate {
	return getCalendarDateClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/initWithString:
func NewCalendarDateWithString(description IString) CalendarDate {
	instance := getCalendarDateClass().Alloc()
	rv := objc.Send[CalendarDate](instance.ID, objc.Sel("initWithString:"), description)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/initWithString:calendarFormat:
func NewCalendarDateWithStringCalendarFormat(description IString, format IString) CalendarDate {
	instance := getCalendarDateClass().Alloc()
	rv := objc.Send[CalendarDate](instance.ID, objc.Sel("initWithString:calendarFormat:"), description, format)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/initWithString:calendarFormat:locale:
func NewCalendarDateWithStringCalendarFormatLocale(description IString, format IString, locale objectivec.IObject) CalendarDate {
	instance := getCalendarDateClass().Alloc()
	rv := objc.Send[CalendarDate](instance.ID, objc.Sel("initWithString:calendarFormat:locale:"), description, format, locale)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/initWithYear:month:day:hour:minute:second:timeZone:
func NewCalendarDateWithYearMonthDayHourMinuteSecondTimeZone(year int /* primitive/slice/pointer. */, month uint /* primitive/slice/pointer. */, day uint /* primitive/slice/pointer. */, hour uint /* primitive/slice/pointer. */, minute uint /* primitive/slice/pointer. */, second uint /* primitive/slice/pointer. */, aTimeZone ITimeZone) CalendarDate {
	instance := getCalendarDateClass().Alloc()
	rv := objc.Send[CalendarDate](instance.ID, objc.Sel("initWithYear:month:day:hour:minute:second:timeZone:"), year, month, day, hour, minute, second, aTimeZone)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/calendarDate
func (cc _CalendarDateClass) CalendarDate() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("calendarDate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/dateWithString:calendarFormat:
func (cc _CalendarDateClass) DateWithStringCalendarFormat(description IString, format IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("dateWithString:calendarFormat:"), description, format)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/dateWithString:calendarFormat:locale:
func (cc _CalendarDateClass) DateWithStringCalendarFormatLocale(description IString, format IString, locale objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("dateWithString:calendarFormat:locale:"), description, format, locale)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/dateWithYear:month:day:hour:minute:second:timeZone:
func (cc _CalendarDateClass) DateWithYearMonthDayHourMinuteSecondTimeZone(year int /* primitive/slice/pointer. */, month uint /* primitive/slice/pointer. */, day uint /* primitive/slice/pointer. */, hour uint /* primitive/slice/pointer. */, minute uint /* primitive/slice/pointer. */, second uint /* primitive/slice/pointer. */, aTimeZone ITimeZone) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("dateWithYear:month:day:hour:minute:second:timeZone:"), year, month, day, hour, minute, second, aTimeZone)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/distantFuture
func (cc _CalendarDateClass) DistantFuture() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("distantFuture"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/distantPast
func (cc _CalendarDateClass) DistantPast() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("distantPast"))
	return rv
}


