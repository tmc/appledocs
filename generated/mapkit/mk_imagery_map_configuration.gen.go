// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MKImageryMapConfiguration] class.
var (
	MKImageryMapConfigurationClass     _MKImageryMapConfigurationClass
	MKImageryMapConfigurationClassOnce sync.Once
)

func getMKImageryMapConfigurationClass() _MKImageryMapConfigurationClass {
	MKImageryMapConfigurationClassOnce.Do(func() {
		MKImageryMapConfigurationClass = _MKImageryMapConfigurationClass{objc.GetClass("MKImageryMapConfiguration")}
	})
	return MKImageryMapConfigurationClass
}

type _MKImageryMapConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MKImageryMapConfiguration] class.
type IMKImageryMapConfiguration interface {
	IMKMapConfiguration
	PitchButtonVisibility() unsafe.Pointer
	SetPitchButtonVisibility(value unsafe.Pointer)
	PreferredConfiguration() MKMapConfiguration
	SetPreferredConfiguration(value IMKMapConfiguration)
	ShowsUserTrackingButton() bool
	SetShowsUserTrackingButton(value bool)
}

// The class that represents an imagery-based map presentation, such as one using satellite imagery.


// The class that represents an imagery-based map presentation, such as one using satellite imagery.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKImageryMapConfiguration
type MKImageryMapConfiguration struct {
	MKMapConfiguration
}

// MKImageryMapConfigurationFrom constructs a [MKImageryMapConfiguration] from an unsafe.Pointer.
//
// The class that represents an imagery-based map presentation, such as one using satellite imagery.
func MKImageryMapConfigurationFrom(ptr unsafe.Pointer) MKImageryMapConfiguration {
	return MKImageryMapConfiguration{
		MKMapConfiguration: MKMapConfigurationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MKImageryMapConfigurationClass) Alloc() MKImageryMapConfiguration {
	rv := objc.Send[MKImageryMapConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKImageryMapConfigurationClass) New() MKImageryMapConfiguration {
	rv := objc.Send[MKImageryMapConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKImageryMapConfiguration) Init() MKImageryMapConfiguration {
	rv := objc.Send[MKImageryMapConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKImageryMapConfiguration) Autorelease() MKImageryMapConfiguration {
	rv := objc.Send[MKImageryMapConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKImageryMapConfiguration creates a new MKImageryMapConfiguration instance.
func NewMKImageryMapConfiguration() MKImageryMapConfiguration {
	return getMKImageryMapConfigurationClass().New()
}



// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKImageryMapConfiguration) PitchButtonVisibility() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pitchButtonVisibility"))
	return rv
}


// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKImageryMapConfiguration) SetPitchButtonVisibility(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPitchButtonVisibility:"), value)
}


// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKImageryMapConfiguration) PreferredConfiguration() MKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](m_.ID, objc.Sel("preferredConfiguration"))
	return rv
}


// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKImageryMapConfiguration) SetPreferredConfiguration(value IMKMapConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredConfiguration:"), value)
}


// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKImageryMapConfiguration) ShowsUserTrackingButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsUserTrackingButton"))
	return rv
}


// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKImageryMapConfiguration) SetShowsUserTrackingButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsUserTrackingButton:"), value)
}



