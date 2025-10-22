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
	LocationSourceInformationClass     _LocationSourceInformationClass
	LocationSourceInformationClassOnce sync.Once
)

func getLocationSourceInformationClass() _LocationSourceInformationClass {
	LocationSourceInformationClassOnce.Do(func() {
		LocationSourceInformationClass = _LocationSourceInformationClass{objc.GetClass("CLLocationSourceInformation")}
	})
	return LocationSourceInformationClass
}

type _LocationSourceInformationClass struct {
	class objc.Class
}

// An interface definition for the [LocationSourceInformation] class.
type ILocationSourceInformation interface {
	objectivec.IObject
	IsProducedByAccessory() bool
	SetIsProducedByAccessory(value bool)
	IsSimulatedBySoftware() bool
	SetIsSimulatedBySoftware(value bool)
}

// Information about the source that provides a location.
//
// contains information about the source that provides a instance, such as instances that delivers. For example, an app may choose to check the source information and reject locations if the property is when the developer isn’t debugging or testing the app.
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


// A Boolean value that indicates whether the system receives the location from an external accessory.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationSourceInformation/isProducedByAccessory
func (l_ LocationSourceInformation) IsProducedByAccessory() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isProducedByAccessory"))
	return rv
}


// SetIsProducedByAccessory sets the value of the isProducedByAccessory property.
// A Boolean value that indicates whether the system receives the location from an external accessory.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationSourceInformation/isProducedByAccessory
func (l_ LocationSourceInformation) SetIsProducedByAccessory(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsProducedByAccessory:"), value)
}

// A Boolean value that indicates whether the system generates the location using on-device software simulation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationSourceInformation/isSimulatedBySoftware
func (l_ LocationSourceInformation) IsSimulatedBySoftware() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isSimulatedBySoftware"))
	return rv
}


// SetIsSimulatedBySoftware sets the value of the isSimulatedBySoftware property.
// A Boolean value that indicates whether the system generates the location using on-device software simulation.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationSourceInformation/isSimulatedBySoftware
func (l_ LocationSourceInformation) SetIsSimulatedBySoftware(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsSimulatedBySoftware:"), value)
}



