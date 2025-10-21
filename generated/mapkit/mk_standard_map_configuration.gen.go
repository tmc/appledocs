// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MKStandardMapConfiguration] class.
var (
	MKStandardMapConfigurationClass     _MKStandardMapConfigurationClass
	MKStandardMapConfigurationClassOnce sync.Once
)

func getMKStandardMapConfigurationClass() _MKStandardMapConfigurationClass {
	MKStandardMapConfigurationClassOnce.Do(func() {
		MKStandardMapConfigurationClass = _MKStandardMapConfigurationClass{objc.GetClass("MKStandardMapConfiguration")}
	})
	return MKStandardMapConfigurationClass
}

type _MKStandardMapConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MKStandardMapConfiguration] class.
type IMKStandardMapConfiguration interface {
	IMKMapConfiguration
}

// The class that represents the default map presentation, which is a street map that shows the position of all roads and some road names.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration
type MKStandardMapConfiguration struct {
	MKMapConfiguration
}

// MKStandardMapConfigurationFrom constructs a [MKStandardMapConfiguration] from an unsafe.Pointer.
//
// The class that represents the default map presentation, which is a street map that shows the position of all roads and some road names.
func MKStandardMapConfigurationFrom(ptr unsafe.Pointer) MKStandardMapConfiguration {
	return MKStandardMapConfiguration{
		MKMapConfiguration: MKMapConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKStandardMapConfigurationClass) Alloc() MKStandardMapConfiguration {
	rv := objc.Send[MKStandardMapConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKStandardMapConfigurationClass) New() MKStandardMapConfiguration {
	rv := objc.Send[MKStandardMapConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKStandardMapConfiguration) Init() MKStandardMapConfiguration {
	rv := objc.Send[MKStandardMapConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKStandardMapConfiguration) Autorelease() MKStandardMapConfiguration {
	rv := objc.Send[MKStandardMapConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKStandardMapConfiguration creates a new MKStandardMapConfiguration instance.
func NewMKStandardMapConfiguration() MKStandardMapConfiguration {
	return getMKStandardMapConfigurationClass().New()
}


// The filter used to determine the points of interest shown on the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/pointOfInterestFilter
func (m_ MKStandardMapConfiguration) PointOfInterestFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}


// SetPointOfInterestFilter sets the value of the pointOfInterestFilter property.
// The filter used to determine the points of interest shown on the map.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/pointOfInterestFilter
func (m_ MKStandardMapConfiguration) SetPointOfInterestFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}


