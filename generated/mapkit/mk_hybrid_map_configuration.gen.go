// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKHybridMapConfiguration */


/* debug [class_header]: Header for MKHybridMapConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKHybridMapConfiguration */
// An interface definition for the [MKHybridMapConfiguration] class.
type IMKHybridMapConfiguration interface {
	IMKMapConfiguration
	
/* debug [class_interface_properties]: Properties for MKHybridMapConfiguration */
	// properties:
	PointOfInterestFilter() IMKPointOfInterestFilter
	SetPointOfInterestFilter(value IMKPointOfInterestFilter)
	ShowsTraffic() bool
	SetShowsTraffic(value bool)
	PitchButtonVisibility() MKFeatureVisibility
	SetPitchButtonVisibility(value MKFeatureVisibility)
	PreferredConfiguration() IMKMapConfiguration
	SetPreferredConfiguration(value IMKMapConfiguration)
	ShowsUserTrackingButton() bool
	SetShowsUserTrackingButton(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKHybridMapConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKHybridMapConfiguration */
// Alloc allocates a new instance without initialization.
func (mc _MKHybridMapConfigurationClass) Alloc() MKHybridMapConfiguration {
	rv := objc.Send[MKHybridMapConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKHybridMapConfiguration */
// The class that represents a satellite image of the area with road and road name information layers on top.


// The class that represents a satellite image of the area with road and road name information layers on top.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKHybridMapConfiguration */

// Creates a new hybrid map configuration with the specified elevation style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKHybridMapConfiguration/init(elevationStyle:)
func NewMKHybridMapConfigurationWithElevationStyle(elevationStyle MKMapElevationStyle) MKHybridMapConfiguration {
	instance := getMKHybridMapConfigurationClass().Alloc()
	rv := objc.Send[MKHybridMapConfiguration](instance.ID, objc.Sel("initWithElevationStyle:"), elevationStyle)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKHybridMapConfigurationWithElevationStyle */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKHybridMapConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKHybridMapConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKHybridMapConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKHybridMapConfiguration */

// The filter the framework uses to determine the points of interest to show on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKHybridMapConfiguration/pointOfInterestFilter
func (m_ MKHybridMapConfiguration) PointOfInterestFilter() IMKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}/* debug [instance_properties/getter]: pointOfInterestFilter */


// The filter the framework uses to determine the points of interest to show on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKHybridMapConfiguration/pointOfInterestFilter
func (m_ MKHybridMapConfiguration) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}/* debug [instance_properties/setter]: pointOfInterestFilter */


// A Boolean value that indicates whether the maps shows traffic conditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKHybridMapConfiguration/showsTraffic
func (m_ MKHybridMapConfiguration) ShowsTraffic() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsTraffic"))
	return rv
}/* debug [instance_properties/getter]: showsTraffic */


// A Boolean value that indicates whether the maps shows traffic conditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKHybridMapConfiguration/showsTraffic
func (m_ MKHybridMapConfiguration) SetShowsTraffic(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsTraffic:"), value)
}/* debug [instance_properties/setter]: showsTraffic */


// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKHybridMapConfiguration) PitchButtonVisibility() MKFeatureVisibility {
	rv := objc.Send[MKFeatureVisibility](m_.ID, objc.Sel("pitchButtonVisibility"))
	return rv
}/* debug [instance_properties/getter]: pitchButtonVisibility */


// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKHybridMapConfiguration) SetPitchButtonVisibility(value MKFeatureVisibility) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPitchButtonVisibility:"), value)
}/* debug [instance_properties/setter]: pitchButtonVisibility */


// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKHybridMapConfiguration) PreferredConfiguration() IMKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](m_.ID, objc.Sel("preferredConfiguration"))
	return rv
}/* debug [instance_properties/getter]: preferredConfiguration */


// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKHybridMapConfiguration) SetPreferredConfiguration(value IMKMapConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredConfiguration:"), value)
}/* debug [instance_properties/setter]: preferredConfiguration */


// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKHybridMapConfiguration) ShowsUserTrackingButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsUserTrackingButton"))
	return rv
}/* debug [instance_properties/getter]: showsUserTrackingButton */


// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKHybridMapConfiguration) SetShowsUserTrackingButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsUserTrackingButton:"), value)
}/* debug [instance_properties/setter]: showsUserTrackingButton */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKHybridMapConfiguration */


