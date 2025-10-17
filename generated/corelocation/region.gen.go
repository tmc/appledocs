// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Region] class.
var regionClass = _RegionClass{objc.GetClass("CLRegion")}

type _RegionClass struct {
	class objc.Class
}

// A base class representing an area that can be monitored. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion

type Region struct {
	objectivec.Object
}

// RegionFrom constructs a [Region] from an unsafe.Pointer.
//
// A base class representing an area that can be monitored.
func RegionFrom(ptr unsafe.Pointer) Region {
	return Region{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (rc _RegionClass) Alloc() Region {
	rv := objc.Send[Region](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (rc _RegionClass) New() Region {
	rv := objc.Send[Region](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ Region) Init() Region {
	rv := objc.Send[Region](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ Region) Autorelease() Region {
	rv := objc.Send[Region](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRegion creates a new Region instance.
func NewRegion() Region {
	return regionClass.New()
}
// Initializes and returns a region object defining a circular area. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion/init(circularRegionWithCenter:radius:identifier:)
func NewRegionCircularRegionWithCenterRadiusIdentifier(center unsafe.Pointer, radius unsafe.Pointer, identifier string) Region {
	instance := regionClass.Alloc()
	rv := objc.Send[Region](instance.ID, objc.Sel("initCircularRegionWithCenter:radius:identifier:"), center, radius, identifier)
	rv.Autorelease()
	return rv
}


// Returns a Boolean value indicating whether the region contains the specified coordinate. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion/contains(_:)
func (r_ Region) ContainsCoordinate(coordinate unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("containsCoordinate:"), coordinate)
	return rv
}

