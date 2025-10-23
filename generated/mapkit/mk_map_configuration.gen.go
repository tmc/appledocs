// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MKMapConfiguration] class.
var (
	MKMapConfigurationClass     _MKMapConfigurationClass
	MKMapConfigurationClassOnce sync.Once
)

func getMKMapConfigurationClass() _MKMapConfigurationClass {
	MKMapConfigurationClassOnce.Do(func() {
		MKMapConfigurationClass = _MKMapConfigurationClass{objc.GetClass("MKMapConfiguration")}
	})
	return MKMapConfigurationClass
}

type _MKMapConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MKMapConfiguration] class.
type IMKMapConfiguration interface {
	objectivec.IObject
	ElevationStyle() unsafe.Pointer
	SetElevationStyle(value unsafe.Pointer)
	PitchButtonVisibility() unsafe.Pointer
	SetPitchButtonVisibility(value unsafe.Pointer)
	PreferredConfiguration() MKMapConfiguration
	SetPreferredConfiguration(value IMKMapConfiguration)
	ShowsUserTrackingButton() bool
	SetShowsUserTrackingButton(value bool)
}

// An abstract class that represents the shared elements of map configurations.


// An abstract class that represents the shared elements of map configurations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapConfiguration
type MKMapConfiguration struct {
	objectivec.Object
}

// MKMapConfigurationFrom constructs a [MKMapConfiguration] from an unsafe.Pointer.
//
// An abstract class that represents the shared elements of map configurations.
func MKMapConfigurationFrom(ptr unsafe.Pointer) MKMapConfiguration {
	return MKMapConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MKMapConfigurationClass) Alloc() MKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MKMapConfigurationClass) New() MKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKMapConfiguration) Init() MKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKMapConfiguration) Autorelease() MKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKMapConfiguration creates a new MKMapConfiguration instance.
func NewMKMapConfiguration() MKMapConfiguration {
	return getMKMapConfigurationClass().New()
}



// The value that indicates the map’s elevation style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapconfiguration/elevationstyle-swift.property
func (m_ MKMapConfiguration) ElevationStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("elevationStyle"))
	return rv
}


// The value that indicates the map’s elevation style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapconfiguration/elevationstyle-swift.property
func (m_ MKMapConfiguration) SetElevationStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setElevationStyle:"), value)
}


// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKMapConfiguration) PitchButtonVisibility() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pitchButtonVisibility"))
	return rv
}


// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKMapConfiguration) SetPitchButtonVisibility(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPitchButtonVisibility:"), value)
}


// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKMapConfiguration) PreferredConfiguration() MKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](m_.ID, objc.Sel("preferredConfiguration"))
	return rv
}


// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKMapConfiguration) SetPreferredConfiguration(value IMKMapConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredConfiguration:"), value)
}


// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKMapConfiguration) ShowsUserTrackingButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsUserTrackingButton"))
	return rv
}


// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKMapConfiguration) SetShowsUserTrackingButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsUserTrackingButton:"), value)
}



