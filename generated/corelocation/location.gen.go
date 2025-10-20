// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Location] class.
var (
	locationClass     _LocationClass
	locationClassOnce sync.Once
)

func getLocationClass() _LocationClass {
	locationClassOnce.Do(func() {
		locationClass = _LocationClass{objc.GetClass("CLLocation")}
	})
	return locationClass
}

type _LocationClass struct {
	class objc.Class
}

// An interface definition for the [Location] class.
type ILocation interface {
	objectivec.IObject
	Distance()
	GetDistanceFrom()
}

// The latitude, longitude, and course information reported by the system.
//
// A object contains the geographical location and altitude of a device, along with values indicating the accuracy of those measurements and when they were collected. In iOS, a location object also contains course information — that is, the speed and heading in which the device was moving. Typically, you don’t create location objects yourself. After you request location updates from your object, the system uses onboard sensors to gather location data and report that data to your app. Some services also return previously collected location data, which you can use as context to improve your services. You can always retrieve the most recently collected location from the property of your object. You may create location objects yourself when you want to cache custom location data or calculate the distance between two geographical coordinates. Use objects as-is, and don’t subclass them.
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

// Alloc allocates a new instance without initialization.
func (lc _LocationClass) Alloc() Location {
	rv := objc.Send[Location](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LocationClass) New() Location {
	rv := objc.Send[Location](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ Location) Init() Location {
	rv := objc.Send[Location](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ Location) Autorelease() Location {
	rv := objc.Send[Location](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLocation creates a new Location instance.
func NewLocation() Location {
	return getLocationClass().New()
}


// Returns the distance (measured in meters) from the current object’s location to the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/distance(from:)
func (l_ Location) Distance() {
	objc.Send[objc.ID](l_.ID, objc.Sel("distance"))
}

// Returns the distance (measured in meters) from the current object’s location to the specified location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocation/getDistanceFrom(_:)
func (l_ Location) GetDistanceFrom() {
	objc.Send[objc.ID](l_.ID, objc.Sel("getDistanceFrom"))
}



