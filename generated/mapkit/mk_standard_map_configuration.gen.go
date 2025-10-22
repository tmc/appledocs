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
	PointOfInterestFilter() MKPointOfInterestFilter
	SetPointOfInterestFilter(value IMKPointOfInterestFilter)
	PitchButtonVisibility() unsafe.Pointer
	SetPitchButtonVisibility(value unsafe.Pointer)
	PreferredConfiguration() MKMapConfiguration
	SetPreferredConfiguration(value IMKMapConfiguration)
	ShowsUserTrackingButton() bool
	SetShowsUserTrackingButton(value bool)
	EmphasisStyle() unsafe.Pointer
	SetEmphasisStyle(value unsafe.Pointer)
	ShowsTraffic() bool
	SetShowsTraffic(value bool)
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




// Creates a standard map configuration with the specified emphasis style.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/init(emphasisStyle:)
func NewMKStandardMapConfigurationWithEmphasisStyle(emphasisStyle unsafe.Pointer) MKStandardMapConfiguration {
	instance := getMKStandardMapConfigurationClass().Alloc()
	rv := objc.Send[MKStandardMapConfiguration](instance.ID, objc.Sel("initWithEmphasisStyle:"), emphasisStyle)
	rv.Autorelease()
	return rv
}


// The filter used to determine the points of interest shown on the map.
//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/pointOfInterestFilter
func (m_ MKStandardMapConfiguration) PointOfInterestFilter() MKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}


// SetPointOfInterestFilter sets the value of the pointOfInterestFilter property.
// The filter used to determine the points of interest shown on the map.

//
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/pointOfInterestFilter
func (m_ MKStandardMapConfiguration) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}

// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKStandardMapConfiguration) PitchButtonVisibility() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pitchButtonVisibility"))
	return rv
}


// SetPitchButtonVisibility sets the value of the pitchButtonVisibility property.
// A value that indicates whether the map’s pitch button is visible.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKStandardMapConfiguration) SetPitchButtonVisibility(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPitchButtonVisibility:"), value)
}

// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKStandardMapConfiguration) PreferredConfiguration() MKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](m_.ID, objc.Sel("preferredConfiguration"))
	return rv
}


// SetPreferredConfiguration sets the value of the preferredConfiguration property.
// The characteristics of the map view, including the map type and features the map displays.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKStandardMapConfiguration) SetPreferredConfiguration(value IMKMapConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredConfiguration:"), value)
}

// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKStandardMapConfiguration) ShowsUserTrackingButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsUserTrackingButton"))
	return rv
}


// SetShowsUserTrackingButton sets the value of the showsUserTrackingButton property.
// A Boolean value that indicates whether the map displays the user tracking button.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKStandardMapConfiguration) SetShowsUserTrackingButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsUserTrackingButton:"), value)
}

// The value that indicates how the framework emphasizes map features.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkstandardmapconfiguration/emphasisstyle-swift.property
func (m_ MKStandardMapConfiguration) EmphasisStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("emphasisStyle"))
	return rv
}


// SetEmphasisStyle sets the value of the emphasisStyle property.
// The value that indicates how the framework emphasizes map features.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkstandardmapconfiguration/emphasisstyle-swift.property
func (m_ MKStandardMapConfiguration) SetEmphasisStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEmphasisStyle:"), value)
}

// A Boolean value that controls whether the map displays traffic conditions.
//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkstandardmapconfiguration/showstraffic
func (m_ MKStandardMapConfiguration) ShowsTraffic() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsTraffic"))
	return rv
}


// SetShowsTraffic sets the value of the showsTraffic property.
// A Boolean value that controls whether the map displays traffic conditions.

//
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkstandardmapconfiguration/showstraffic
func (m_ MKStandardMapConfiguration) SetShowsTraffic(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsTraffic:"), value)
}


