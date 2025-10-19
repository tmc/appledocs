// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LocationSourceInformation] class.
var (
	locationSourceInformationClass     _LocationSourceInformationClass
	locationSourceInformationClassOnce sync.Once
)

func getLocationSourceInformationClass() _LocationSourceInformationClass {
	locationSourceInformationClassOnce.Do(func() {
		locationSourceInformationClass = _LocationSourceInformationClass{objc.GetClass("CLLocationSourceInformation")}
	})
	return locationSourceInformationClass
}

type _LocationSourceInformationClass struct {
	class objc.Class
}

// An interface definition for the [LocationSourceInformation] class.
type ILocationSourceInformation interface {
	objectivec.IObject
}

// Information about the source that provides a location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationSourceInformation
type LocationSourceInformation struct {
	objectivec.Object
}

// LocationSourceInformationFrom constructs a [LocationSourceInformation] from an unsafe.Pointer.
//
// Information about the source that provides a location.
func LocationSourceInformationFrom(ptr unsafe.Pointer) LocationSourceInformation {
	return LocationSourceInformation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LocationSourceInformationClass) Alloc() LocationSourceInformation {
	rv := objc.Send[LocationSourceInformation](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LocationSourceInformationClass) New() LocationSourceInformation {
	rv := objc.Send[LocationSourceInformation](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LocationSourceInformation) Init() LocationSourceInformation {
	rv := objc.Send[LocationSourceInformation](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LocationSourceInformation) Autorelease() LocationSourceInformation {
	rv := objc.Send[LocationSourceInformation](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLocationSourceInformation creates a new LocationSourceInformation instance.
func NewLocationSourceInformation() LocationSourceInformation {
	return getLocationSourceInformationClass().New()
}




