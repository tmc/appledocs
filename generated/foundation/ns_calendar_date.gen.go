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
	DayOfCommonEra() int
	DescriptionWithCalendarFormat(format string) String
	DescriptionWithCalendarFormatLocale(format string, locale objectivec.IObject) String
	HourOfDay() int
}

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



//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/dateWithYear:month:day:hour:minute:second:timeZone:

func (cc _CalendarDateClass) DateWithYearMonthDayHourMinuteSecondTimeZone(year int, month uint, day uint, hour uint, minute uint, second uint, aTimeZone ITimeZone) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("dateWithYear:month:day:hour:minute:second:timeZone:"), year, month, day, hour, minute, second, aTimeZone)
	return rv
}


//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/distantPast

func (cc _CalendarDateClass) DistantPast() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("distantPast"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/dayOfCommonEra
func (c_ CalendarDate) DayOfCommonEra() int {
	rv := objc.Send[int](c_.ID, objc.Sel("dayOfCommonEra"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/descriptionWithCalendarFormat:
func (c_ CalendarDate) DescriptionWithCalendarFormat(format string) String {
	rv := objc.Send[String](c_.ID, objc.Sel("descriptionWithCalendarFormat:"), objc.String(format))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/descriptionWithCalendarFormat:locale:
func (c_ CalendarDate) DescriptionWithCalendarFormatLocale(format string, locale objectivec.IObject) String {
	rv := objc.Send[String](c_.ID, objc.Sel("descriptionWithCalendarFormat:locale:"), objc.String(format), locale)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/hourOfDay
func (c_ CalendarDate) HourOfDay() int {
	rv := objc.Send[int](c_.ID, objc.Sel("hourOfDay"))
	return rv
}



