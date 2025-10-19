// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Region] class.
var (
	regionClass     _RegionClass
	regionClassOnce sync.Once
)

func getRegionClass() _RegionClass {
	regionClassOnce.Do(func() {
		regionClass = _RegionClass{objc.GetClass("CLRegion")}
	})
	return regionClass
}

type _RegionClass struct {
	class objc.Class
}

// An interface definition for the [Region] class.
type IRegion interface {
	objectivec.IObject
	ContainsCoordinate(coordinate unsafe.Pointer) bool
}

// A base class representing an area that can be monitored.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getRegionClass().New()
}


// Initializes and returns a region object defining a circular area.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion/init(circularRegionWithCenter:radius:identifier:)
func NewRegionCircularRegionWithCenterRadiusIdentifier(center unsafe.Pointer, radius unsafe.Pointer, identifier string) Region {
	instance := getRegionClass().Alloc()
	rv := objc.Send[Region](instance.ID, objc.Sel("initCircularRegionWithCenter:radius:identifier:"), center, radius, objc.String(identifier))
	rv.Autorelease()
	return rv
}


// Returns a Boolean value indicating whether the region contains the specified coordinate.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion/contains(_:)
func (r_ Region) ContainsCoordinate(coordinate unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("containsCoordinate:"), coordinate)
	return rv
}

