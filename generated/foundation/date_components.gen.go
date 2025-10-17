// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DateComponents] class.
var dateComponentsClass = _DateComponentsClass{objc.GetClass("NSDateComponents")}

type _DateComponentsClass struct {
	class objc.Class
}

// An object that specifies a date or time in terms of units (such as year, month, day, hour, and minute) to be evaluated in a calendar system and time zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDateComponents

type DateComponents struct {
	objectivec.Object
}

// DateComponentsFrom constructs a [DateComponents] from an unsafe.Pointer.
//
// An object that specifies a date or time in terms of units (such as year, month, day, hour, and minute) to be evaluated in a calendar system and time zone.
func DateComponentsFrom(ptr unsafe.Pointer) DateComponents {
	return DateComponents{objectivec.Object{objc.ID(ptr)}}
}



