// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKImageryMapConfiguration */


/* debug [class_header]: Header for MKImageryMapConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKImageryMapConfiguration */
// An interface definition for the [MKImageryMapConfiguration] class.
type IMKImageryMapConfiguration interface {
	IMKMapConfiguration
	
/* debug [class_interface_properties]: Properties for MKImageryMapConfiguration */
	// properties:
	PitchButtonVisibility() MKFeatureVisibility
	SetPitchButtonVisibility(value MKFeatureVisibility)
	PreferredConfiguration() IMKMapConfiguration
	SetPreferredConfiguration(value IMKMapConfiguration)
	ShowsUserTrackingButton() bool
	SetShowsUserTrackingButton(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKImageryMapConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKImageryMapConfiguration */
// Alloc allocates a new instance without initialization.
func (mc _MKImageryMapConfigurationClass) Alloc() MKImageryMapConfiguration {
	rv := objc.Send[MKImageryMapConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKImageryMapConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKImageryMapConfiguration */

// Creates a new imagery based map configuration with the specified elevation style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKImageryMapConfiguration/init(elevationStyle:)
func NewMKImageryMapConfigurationWithElevationStyle(elevationStyle MKMapElevationStyle) MKImageryMapConfiguration {
	instance := getMKImageryMapConfigurationClass().Alloc()
	rv := objc.Send[MKImageryMapConfiguration](instance.ID, objc.Sel("initWithElevationStyle:"), elevationStyle)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKImageryMapConfigurationWithElevationStyle */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKImageryMapConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKImageryMapConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKImageryMapConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKImageryMapConfiguration */

// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKImageryMapConfiguration) PitchButtonVisibility() MKFeatureVisibility {
	rv := objc.Send[MKFeatureVisibility](m_.ID, objc.Sel("pitchButtonVisibility"))
	return rv
}/* debug [instance_properties/getter]: pitchButtonVisibility */


// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKImageryMapConfiguration) SetPitchButtonVisibility(value MKFeatureVisibility) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPitchButtonVisibility:"), value)
}/* debug [instance_properties/setter]: pitchButtonVisibility */


// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKImageryMapConfiguration) PreferredConfiguration() IMKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](m_.ID, objc.Sel("preferredConfiguration"))
	return rv
}/* debug [instance_properties/getter]: preferredConfiguration */


// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKImageryMapConfiguration) SetPreferredConfiguration(value IMKMapConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredConfiguration:"), value)
}/* debug [instance_properties/setter]: preferredConfiguration */


// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKImageryMapConfiguration) ShowsUserTrackingButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsUserTrackingButton"))
	return rv
}/* debug [instance_properties/getter]: showsUserTrackingButton */


// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKImageryMapConfiguration) SetShowsUserTrackingButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsUserTrackingButton:"), value)
}/* debug [instance_properties/setter]: showsUserTrackingButton */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKImageryMapConfiguration */


