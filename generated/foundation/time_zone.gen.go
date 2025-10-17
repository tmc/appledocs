// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TimeZone] class.
var timeZoneClass = _TimeZoneClass{objc.GetClass("NSTimeZone")}

type _TimeZoneClass struct {
	class objc.Class
}

// Information about standard time conventions associated with a specific geopolitical region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTimeZone

type TimeZone struct {
	objectivec.Object
}

// TimeZoneFrom constructs a [TimeZone] from an unsafe.Pointer.
//
// Information about standard time conventions associated with a specific geopolitical region.
func TimeZoneFrom(ptr unsafe.Pointer) TimeZone {
	return TimeZone{objectivec.Object{objc.ID(ptr)}}
}



