// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MKHybridMapConfiguration] class.
var (
	MKHybridMapConfigurationClass     _MKHybridMapConfigurationClass
	MKHybridMapConfigurationClassOnce sync.Once
)

func getMKHybridMapConfigurationClass() _MKHybridMapConfigurationClass {
	MKHybridMapConfigurationClassOnce.Do(func() {
		MKHybridMapConfigurationClass = _MKHybridMapConfigurationClass{objc.GetClass("MKHybridMapConfiguration")}
	})
	return MKHybridMapConfigurationClass
}

type _MKHybridMapConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MKHybridMapConfiguration] class.
type IMKHybridMapConfiguration interface {
	IMKMapConfiguration
}

// The class that represents a satellite image of the area with road and road name information layers on top.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKHybridMapConfiguration
type MKHybridMapConfiguration struct {
	MKMapConfiguration
}

// MKHybridMapConfigurationFrom constructs a [MKHybridMapConfiguration] from an unsafe.Pointer.
//
// The class that represents a satellite image of the area with road and road name information layers on top.
func MKHybridMapConfigurationFrom(ptr unsafe.Pointer) MKHybridMapConfiguration {
	return MKHybridMapConfiguration{
		MKMapConfiguration: MKMapConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKHybridMapConfigurationClass) Alloc() MKHybridMapConfiguration {
	rv := objc.Send[MKHybridMapConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKHybridMapConfigurationClass) New() MKHybridMapConfiguration {
	rv := objc.Send[MKHybridMapConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKHybridMapConfiguration) Init() MKHybridMapConfiguration {
	rv := objc.Send[MKHybridMapConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKHybridMapConfiguration) Autorelease() MKHybridMapConfiguration {
	rv := objc.Send[MKHybridMapConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKHybridMapConfiguration creates a new MKHybridMapConfiguration instance.
func NewMKHybridMapConfiguration() MKHybridMapConfiguration {
	return getMKHybridMapConfigurationClass().New()
}



// The filter the framework uses to determine the points of interest to show on the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKHybridMapConfiguration/pointOfInterestFilter
func (m_ MKHybridMapConfiguration) PointOfInterestFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}


// SetPointOfInterestFilter sets the value of the pointOfInterestFilter property.
// The filter the framework uses to determine the points of interest to show on the map.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKHybridMapConfiguration/pointOfInterestFilter
func (m_ MKHybridMapConfiguration) SetPointOfInterestFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}

