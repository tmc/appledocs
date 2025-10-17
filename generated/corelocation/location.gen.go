// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Location] class.
var locationClass = _LocationClass{objc.GetClass("CLLocation")}

type _LocationClass struct {
	class objc.Class
}

// The latitude, longitude, and course information reported by the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation

type Location struct {
	objectivec.Object
}

// LocationFrom constructs a [Location] from an unsafe.Pointer.
//
// The latitude, longitude, and course information reported by the system.
func LocationFrom(ptr unsafe.Pointer) Location {
	return Location{objectivec.Object{objc.ID(ptr)}}
}

// Returns the distance (measured in meters) from the current object’s location to the specified location. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/distance(from:)
func (l_ Location) Distance() {
	objc.Send[objc.ID](l_.ID, objc.Sel("distance"))
}
// Returns the distance (measured in meters) from the current object’s location to the specified location. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/getDistanceFrom(_:)
func (l_ Location) GetDistanceFrom() {
	objc.Send[objc.ID](l_.ID, objc.Sel("getDistanceFrom"))
}


