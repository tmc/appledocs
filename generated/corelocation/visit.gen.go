// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Visit] class.
var visitClass = _VisitClass{objc.GetClass("CLVisit")}

type _VisitClass struct {
	class objc.Class
}

// Information about the user’s location during a specific period of time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLVisit

type Visit struct {
	objectivec.Object
}

// VisitFrom constructs a [Visit] from an unsafe.Pointer.
//
// Information about the user’s location during a specific period of time.
func VisitFrom(ptr unsafe.Pointer) Visit {
	return Visit{objectivec.Object{objc.ID(ptr)}}
}



