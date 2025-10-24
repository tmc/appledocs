// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKStandardMapConfiguration */


/* debug [class_header]: Header for MKStandardMapConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKStandardMapConfiguration */
// An interface definition for the [MKStandardMapConfiguration] class.
type IMKStandardMapConfiguration interface {
	IMKMapConfiguration
	
/* debug [class_interface_properties]: Properties for MKStandardMapConfiguration */
	// properties:
	EmphasisStyle() MKStandardMapEmphasisStyle
	SetEmphasisStyle(value MKStandardMapEmphasisStyle)
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

	
/* debug [class_interface_methods]: Methods for MKStandardMapConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKStandardMapConfiguration */
// Alloc allocates a new instance without initialization.
func (mc _MKStandardMapConfigurationClass) Alloc() MKStandardMapConfiguration {
	rv := objc.Send[MKStandardMapConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKStandardMapConfiguration */
// The class that represents the default map presentation, which is a street map that shows the position of all roads and some road names.


// The class that represents the default map presentation, which is a street map that shows the position of all roads and some road names.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKStandardMapConfiguration */

// Creates a new standard map configuration with the specified elevation style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/init(elevationStyle:)
func NewMKStandardMapConfigurationWithElevationStyle(elevationStyle MKMapElevationStyle) MKStandardMapConfiguration {
	instance := getMKStandardMapConfigurationClass().Alloc()
	rv := objc.Send[MKStandardMapConfiguration](instance.ID, objc.Sel("initWithElevationStyle:"), elevationStyle)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKStandardMapConfigurationWithElevationStyle */


// Creates a standard map configuration with the specified elevation and emphasis styles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/init(elevationStyle:emphasisStyle:)
func NewMKStandardMapConfigurationWithElevationStyleEmphasisStyle(elevationStyle MKMapElevationStyle, emphasisStyle MKStandardMapEmphasisStyle) MKStandardMapConfiguration {
	instance := getMKStandardMapConfigurationClass().Alloc()
	rv := objc.Send[MKStandardMapConfiguration](instance.ID, objc.Sel("initWithElevationStyle:emphasisStyle:"), elevationStyle, emphasisStyle)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKStandardMapConfigurationWithElevationStyleEmphasisStyle */


// Creates a standard map configuration with the specified emphasis style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/init(emphasisStyle:)
func NewMKStandardMapConfigurationWithEmphasisStyle(emphasisStyle MKStandardMapEmphasisStyle) MKStandardMapConfiguration {
	instance := getMKStandardMapConfigurationClass().Alloc()
	rv := objc.Send[MKStandardMapConfiguration](instance.ID, objc.Sel("initWithEmphasisStyle:"), emphasisStyle)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKStandardMapConfigurationWithEmphasisStyle */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKStandardMapConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKStandardMapConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKStandardMapConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKStandardMapConfiguration */

// The value that indicates how the framework emphasizes map features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/emphasisStyle-swift.property
func (m_ MKStandardMapConfiguration) EmphasisStyle() MKStandardMapEmphasisStyle {
	rv := objc.Send[MKStandardMapEmphasisStyle](m_.ID, objc.Sel("emphasisStyle"))
	return rv
}/* debug [instance_properties/getter]: emphasisStyle */


// The value that indicates how the framework emphasizes map features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/emphasisStyle-swift.property
func (m_ MKStandardMapConfiguration) SetEmphasisStyle(value MKStandardMapEmphasisStyle) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEmphasisStyle:"), value)
}/* debug [instance_properties/setter]: emphasisStyle */


// The filter used to determine the points of interest shown on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/pointOfInterestFilter
func (m_ MKStandardMapConfiguration) PointOfInterestFilter() IMKPointOfInterestFilter {
	rv := objc.Send[MKPointOfInterestFilter](m_.ID, objc.Sel("pointOfInterestFilter"))
	return rv
}/* debug [instance_properties/getter]: pointOfInterestFilter */


// The filter used to determine the points of interest shown on the map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/pointOfInterestFilter
func (m_ MKStandardMapConfiguration) SetPointOfInterestFilter(value IMKPointOfInterestFilter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPointOfInterestFilter:"), value)
}/* debug [instance_properties/setter]: pointOfInterestFilter */


// A Boolean value that controls whether the map displays traffic conditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/showsTraffic
func (m_ MKStandardMapConfiguration) ShowsTraffic() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsTraffic"))
	return rv
}/* debug [instance_properties/getter]: showsTraffic */


// A Boolean value that controls whether the map displays traffic conditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKStandardMapConfiguration/showsTraffic
func (m_ MKStandardMapConfiguration) SetShowsTraffic(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsTraffic:"), value)
}/* debug [instance_properties/setter]: showsTraffic */


// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKStandardMapConfiguration) PitchButtonVisibility() MKFeatureVisibility {
	rv := objc.Send[MKFeatureVisibility](m_.ID, objc.Sel("pitchButtonVisibility"))
	return rv
}/* debug [instance_properties/getter]: pitchButtonVisibility */


// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKStandardMapConfiguration) SetPitchButtonVisibility(value MKFeatureVisibility) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPitchButtonVisibility:"), value)
}/* debug [instance_properties/setter]: pitchButtonVisibility */


// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKStandardMapConfiguration) PreferredConfiguration() IMKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](m_.ID, objc.Sel("preferredConfiguration"))
	return rv
}/* debug [instance_properties/getter]: preferredConfiguration */


// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKStandardMapConfiguration) SetPreferredConfiguration(value IMKMapConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredConfiguration:"), value)
}/* debug [instance_properties/setter]: preferredConfiguration */


// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKStandardMapConfiguration) ShowsUserTrackingButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsUserTrackingButton"))
	return rv
}/* debug [instance_properties/getter]: showsUserTrackingButton */


// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKStandardMapConfiguration) SetShowsUserTrackingButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsUserTrackingButton:"), value)
}/* debug [instance_properties/setter]: showsUserTrackingButton */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKStandardMapConfiguration */


