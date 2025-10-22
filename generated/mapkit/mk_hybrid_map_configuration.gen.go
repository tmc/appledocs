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
	PointOfInterestFilter() MKPointOfInterestFilter
	SetPointOfInterestFilter(value IMKPointOfInterestFilter)
	ShowsTraffic() bool
	SetShowsTraffic(value bool)
	PitchButtonVisibility() unsafe.Pointer
	SetPitchButtonVisibility(value unsafe.Pointer)
	PreferredConfiguration() MKMapConfiguration
	SetPreferredConfiguration(value IMKMapConfiguration)
	ShowsUserTrackingButton() bool
	SetShowsUserTrackingButton(value bool)
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
func (m_ MKHybridMapConfiguration) PointOfInterestFilter() MKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}


// SetPointOfInterestFilter sets the value of the pointOfInterestFilter property.
// The filter the framework uses to determine the points of interest to show on the map.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKHybridMapConfiguration/pointOfInterestFilter
func (m_ MKHybridMapConfiguration) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}

// A Boolean value that indicates whether the maps shows traffic conditions.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkhybridmapconfiguration/showstraffic
func (m_ MKHybridMapConfiguration) ShowsTraffic() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsTraffic"))
	return rv
}


// SetShowsTraffic sets the value of the showsTraffic property.
// A Boolean value that indicates whether the maps shows traffic conditions.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkhybridmapconfiguration/showstraffic
func (m_ MKHybridMapConfiguration) SetShowsTraffic(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsTraffic:"), value)
}

// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKHybridMapConfiguration) PitchButtonVisibility() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pitchButtonVisibility"))
	return rv
}


// SetPitchButtonVisibility sets the value of the pitchButtonVisibility property.
// A value that indicates whether the map’s pitch button is visible.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKHybridMapConfiguration) SetPitchButtonVisibility(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPitchButtonVisibility:"), value)
}

// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKHybridMapConfiguration) PreferredConfiguration() MKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](m_.ID, objc.Sel("preferredConfiguration"))
	return rv
}


// SetPreferredConfiguration sets the value of the preferredConfiguration property.
// The characteristics of the map view, including the map type and features the map displays.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKHybridMapConfiguration) SetPreferredConfiguration(value IMKMapConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredConfiguration:"), value)
}

// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKHybridMapConfiguration) ShowsUserTrackingButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsUserTrackingButton"))
	return rv
}


// SetShowsUserTrackingButton sets the value of the showsUserTrackingButton property.
// A Boolean value that indicates whether the map displays the user tracking button.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKHybridMapConfiguration) SetShowsUserTrackingButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsUserTrackingButton:"), value)
}


