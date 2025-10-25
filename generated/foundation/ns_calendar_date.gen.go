// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCalendarDate */


/* debug [class_header]: Header for NSCalendarDate */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CalendarDate */
// An interface definition for the [CalendarDate] class.
type ICalendarDate interface {
	IDate
	
/* debug [class_interface_properties]: Properties for CalendarDate */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CalendarDate */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CalendarDate */
// Alloc allocates a new instance without initialization.
func (cc _CalendarDateClass) Alloc() CalendarDate {
	rv := objc.Send[CalendarDate](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CalendarDate */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CalendarDate */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/initWithString:
func NewCalendarDateWithString(description IString) CalendarDate {
	instance := getCalendarDateClass().Alloc()
	rv := objc.Send[CalendarDate](instance.ID, objc.Sel("initWithString:"), description)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCalendarDateWithString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/initWithString:calendarFormat:
func NewCalendarDateWithStringCalendarFormat(description IString, format IString) CalendarDate {
	instance := getCalendarDateClass().Alloc()
	rv := objc.Send[CalendarDate](instance.ID, objc.Sel("initWithString:calendarFormat:"), description, format)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCalendarDateWithStringCalendarFormat */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/initWithString:calendarFormat:locale:
func NewCalendarDateWithStringCalendarFormatLocale(description IString, format IString, locale objc.IObject) CalendarDate {
	instance := getCalendarDateClass().Alloc()
	rv := objc.Send[CalendarDate](instance.ID, objc.Sel("initWithString:calendarFormat:locale:"), description, format, locale)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCalendarDateWithStringCalendarFormatLocale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/initWithYear:month:day:hour:minute:second:timeZone:
func NewCalendarDateWithYearMonthDayHourMinuteSecondTimeZone(year int, month uint, day uint, hour uint, minute uint, second uint, aTimeZone ITimeZone) CalendarDate {
	instance := getCalendarDateClass().Alloc()
	rv := objc.Send[CalendarDate](instance.ID, objc.Sel("initWithYear:month:day:hour:minute:second:timeZone:"), year, month, day, hour, minute, second, aTimeZone)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCalendarDateWithYearMonthDayHourMinuteSecondTimeZone */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CalendarDate */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/calendarDate
func (cc _CalendarDateClass) CalendarDate() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("calendarDate"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CalendarDate) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/dateWithString:calendarFormat:
func (cc _CalendarDateClass) DateWithStringCalendarFormat(description IString, format IString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("dateWithString:calendarFormat:"), description, format)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DateWithStringCalendarFormat) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/dateWithString:calendarFormat:locale:
func (cc _CalendarDateClass) DateWithStringCalendarFormatLocale(description IString, format IString, locale objc.IObject) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("dateWithString:calendarFormat:locale:"), description, format, locale)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DateWithStringCalendarFormatLocale) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/dateWithYear:month:day:hour:minute:second:timeZone:
func (cc _CalendarDateClass) DateWithYearMonthDayHourMinuteSecondTimeZone(year int, month uint, day uint, hour uint, minute uint, second uint, aTimeZone ITimeZone) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("dateWithYear:month:day:hour:minute:second:timeZone:"), year, month, day, hour, minute, second, aTimeZone)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DateWithYearMonthDayHourMinuteSecondTimeZone) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/distantFuture
func (cc _CalendarDateClass) DistantFuture() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("distantFuture"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DistantFuture) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendarDate/distantPast
func (cc _CalendarDateClass) DistantPast() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("distantPast"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DistantPast) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CalendarDate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CalendarDate */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CalendarDate */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCalendarDate */


