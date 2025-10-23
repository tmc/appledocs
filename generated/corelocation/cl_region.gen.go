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
	RegionClass     _RegionClass
	RegionClassOnce sync.Once
)

func getRegionClass() _RegionClass {
	RegionClassOnce.Do(func() {
		RegionClass = _RegionClass{objc.GetClass("CLRegion")}
	})
	return RegionClass
}

type _RegionClass struct {
	class objc.Class
}

// An interface definition for the [Region] class.
type IRegion interface {
	objectivec.IObject
	// properties:
	Center() LocationCoordinate2D /* not a class type */
	Identifier() string /* primitive/slice/pointer. */
	NotifyOnEntry() bool /* primitive/slice/pointer. */
	SetNotifyOnEntry(value bool /* primitive/slice/pointer. */)
	NotifyOnExit() bool /* primitive/slice/pointer. */
	SetNotifyOnExit(value bool /* primitive/slice/pointer. */)
	Radius() LocationDistance /* not a class type */
	// methods:
}

// A base class representing an area that can be monitored.
//
// This is an abstract base class. Instantiate one of the provided subclasses that define specific types of regions. After you create a region, register it with a object with the method. The location manager generates appropriate events whenever the user crosses the boundaries of the region.


// A base class representing an area that can be monitored.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion/init(circularRegionWithCenter:radius:identifier:)
func NewRegionCircularRegionWithCenterRadiusIdentifier(center LocationCoordinate2D /* not a class type */, radius LocationDistance /* not a class type */, identifier string /* primitive/slice/pointer. */) Region {
	instance := getRegionClass().Alloc()
	rv := objc.Send[Region](instance.ID, objc.Sel("initCircularRegionWithCenter:radius:identifier:"), center, radius, objc.String(identifier))
	rv.Autorelease()
	return rv
}



// The center point of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion/center
func (r_ Region) Center() LocationCoordinate2D /* not a class type */ {
	rv := objc.Send[LocationCoordinate2D](r_.ID, objc.Sel("center"))
	return rv
}


// The identifier for the region object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion/identifier
func (r_ Region) Identifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](r_.ID, objc.Sel("identifier"))
	return rv
}


// A Boolean indicating that notifications are generated upon entry into the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion/notifyOnEntry
func (r_ Region) NotifyOnEntry() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("notifyOnEntry"))
	return rv
}


// A Boolean indicating that notifications are generated upon entry into the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion/notifyOnEntry
func (r_ Region) SetNotifyOnEntry(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNotifyOnEntry:"), value)
}


// A Boolean indicating that notifications are generated upon exit from the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion/notifyOnExit
func (r_ Region) NotifyOnExit() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("notifyOnExit"))
	return rv
}


// A Boolean indicating that notifications are generated upon exit from the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion/notifyOnExit
func (r_ Region) SetNotifyOnExit(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setNotifyOnExit:"), value)
}


// The radius (measured in meters) that defines the region’s outer boundary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion/radius
func (r_ Region) Radius() LocationDistance /* not a class type */ {
	rv := objc.Send[LocationDistance](r_.ID, objc.Sel("radius"))
	return rv
}


