// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CalendarDate] class.
var calendarDateClass = _CalendarDateClass{objc.GetClass("NSCalendarDate")}

type _CalendarDateClass struct {
	class objc.Class
}

// An interface definition for the [CalendarDate] class.
type ICalendarDate interface {
	IDate
}

// A specialized date object with embedded calendar information. [Full Topic]
//
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



