// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Floor] class.
var floorClass = _FloorClass{objc.GetClass("CLFloor")}

type _FloorClass struct {
	class objc.Class
}

// The floor of a building on which the user’s device is located. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLFloor

type Floor struct {
	objectivec.Object
}

// FloorFrom constructs a [Floor] from an unsafe.Pointer.
//
// The floor of a building on which the user’s device is located.
func FloorFrom(ptr unsafe.Pointer) Floor {
	return Floor{objectivec.Object{objc.ID(ptr)}}
}



