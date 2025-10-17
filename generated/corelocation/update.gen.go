// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Update] class.
var updateClass = _UpdateClass{objc.GetClass("CLUpdate")}

type _UpdateClass struct {
	class objc.Class
}

// An object that represents a location update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLUpdate

type Update struct {
	objectivec.Object
}

// UpdateFrom constructs a [Update] from an unsafe.Pointer.
//
// An object that represents a location update.
func UpdateFrom(ptr unsafe.Pointer) Update {
	return Update{objectivec.Object{objc.ID(ptr)}}
}



