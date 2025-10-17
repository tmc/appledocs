// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Calendar] class.
var calendarClass = _CalendarClass{objc.GetClass("NSCalendar")}

type _CalendarClass struct {
	class objc.Class
}

// A definition of the relationships between calendar units and absolute points in time, providing features for calculation and comparison of dates. [Full Topic]
//
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



