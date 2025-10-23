// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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




