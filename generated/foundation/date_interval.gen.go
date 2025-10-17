// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DateInterval] class.
var dateIntervalClass = _DateIntervalClass{objc.GetClass("NSDateInterval")}

type _DateIntervalClass struct {
	class objc.Class
}

// An object representing the span of time between a specific start date and end date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateInterval

type DateInterval struct {
	objectivec.Object
}

// DateIntervalFrom constructs a [DateInterval] from an unsafe.Pointer.
//
// An object representing the span of time between a specific start date and end date.
func DateIntervalFrom(ptr unsafe.Pointer) DateInterval {
	return DateInterval{objectivec.Object{objc.ID(ptr)}}
}



