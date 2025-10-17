// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CircularRegion] class.
var circularRegionClass = _CircularRegionClass{objc.GetClass("CLCircularRegion")}

type _CircularRegionClass struct {
	class objc.Class
}

// A circular geographic region that a center point and radius deine. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLCircularRegion

type CircularRegion struct {
	Region
}

// CircularRegionFrom constructs a [CircularRegion] from an unsafe.Pointer.
//
// A circular geographic region that a center point and radius deine.
func CircularRegionFrom(ptr unsafe.Pointer) CircularRegion {
	return CircularRegion{
		Region: RegionFrom(ptr),
	}
}



