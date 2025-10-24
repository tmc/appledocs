// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKMapConfiguration */


/* debug [class_header]: Header for MKMapConfiguration */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKMapConfiguration */
// An interface definition for the [MKMapConfiguration] class.
type IMKMapConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MKMapConfiguration */
	// properties:
	ElevationStyle() MKMapElevationStyle
	SetElevationStyle(value MKMapElevationStyle)
	PitchButtonVisibility() MKFeatureVisibility
	SetPitchButtonVisibility(value MKFeatureVisibility)
	PreferredConfiguration() IMKMapConfiguration
	SetPreferredConfiguration(value IMKMapConfiguration)
	ShowsUserTrackingButton() bool
	SetShowsUserTrackingButton(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKMapConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKMapConfiguration */
// Alloc allocates a new instance without initialization.
func (mc _MKMapConfigurationClass) Alloc() MKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKMapConfiguration */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKMapConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKMapConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKMapConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKMapConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKMapConfiguration */

// The value that indicates the map’s elevation style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapConfiguration/elevationStyle-swift.property
func (m_ MKMapConfiguration) ElevationStyle() MKMapElevationStyle {
	rv := objc.Send[MKMapElevationStyle](m_.ID, objc.Sel("elevationStyle"))
	return rv
}/* debug [instance_properties/getter]: elevationStyle */


// The value that indicates the map’s elevation style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKMapConfiguration/elevationStyle-swift.property
func (m_ MKMapConfiguration) SetElevationStyle(value MKMapElevationStyle) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setElevationStyle:"), value)
}/* debug [instance_properties/setter]: elevationStyle */


// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKMapConfiguration) PitchButtonVisibility() MKFeatureVisibility {
	rv := objc.Send[MKFeatureVisibility](m_.ID, objc.Sel("pitchButtonVisibility"))
	return rv
}/* debug [instance_properties/getter]: pitchButtonVisibility */


// A value that indicates whether the map’s pitch button is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/pitchbuttonvisibility
func (m_ MKMapConfiguration) SetPitchButtonVisibility(value MKFeatureVisibility) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPitchButtonVisibility:"), value)
}/* debug [instance_properties/setter]: pitchButtonVisibility */


// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKMapConfiguration) PreferredConfiguration() IMKMapConfiguration {
	rv := objc.Send[MKMapConfiguration](m_.ID, objc.Sel("preferredConfiguration"))
	return rv
}/* debug [instance_properties/getter]: preferredConfiguration */


// The characteristics of the map view, including the map type and features the map displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/preferredconfiguration
func (m_ MKMapConfiguration) SetPreferredConfiguration(value IMKMapConfiguration) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredConfiguration:"), value)
}/* debug [instance_properties/setter]: preferredConfiguration */


// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKMapConfiguration) ShowsUserTrackingButton() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsUserTrackingButton"))
	return rv
}/* debug [instance_properties/getter]: showsUserTrackingButton */


// A Boolean value that indicates whether the map displays the user tracking button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsusertrackingbutton
func (m_ MKMapConfiguration) SetShowsUserTrackingButton(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsUserTrackingButton:"), value)
}/* debug [instance_properties/setter]: showsUserTrackingButton */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKMapConfiguration */



