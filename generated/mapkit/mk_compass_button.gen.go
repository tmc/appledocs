// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKCompassButton */


/* debug [class_header]: Header for MKCompassButton */
// The class instance for the [MKCompassButton] class.
var (
	MKCompassButtonClass     _MKCompassButtonClass
	MKCompassButtonClassOnce sync.Once
)

func getMKCompassButtonClass() _MKCompassButtonClass {
	MKCompassButtonClassOnce.Do(func() {
		MKCompassButtonClass = _MKCompassButtonClass{objc.GetClass("MKCompassButton")}
	})
	return MKCompassButtonClass
}

type _MKCompassButtonClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKCompassButton */
// An interface definition for the [MKCompassButton] class.
type IMKCompassButton interface {
	IView
	
/* debug [class_interface_properties]: Properties for MKCompassButton */
	// properties:
	CompassVisibility() MKFeatureVisibility
	SetCompassVisibility(value MKFeatureVisibility)
	MapView() IMKMapView
	SetMapView(value IMKMapView)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKCompassButton */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKCompassButton */
// Alloc allocates a new instance without initialization.
func (mc _MKCompassButtonClass) Alloc() MKCompassButton {
	rv := objc.Send[MKCompassButton](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKCompassButtonClass) New() MKCompassButton {
	rv := objc.Send[MKCompassButton](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKCompassButton) Init() MKCompassButton {
	rv := objc.Send[MKCompassButton](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKCompassButton) Autorelease() MKCompassButton {
	rv := objc.Send[MKCompassButton](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKCompassButton creates a new MKCompassButton instance.
func NewMKCompassButton() MKCompassButton {
	return getMKCompassButtonClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKCompassButton */
// A specialized view that displays the compass heading for its associated map.
//
// Use this class when you want to incorporate a standard compass button into your own view hierarchy. A compass button reflects the current orientation of its associated map view. Tapping the compass button reorients the map so that due north is at the top of the map view.


// A specialized view that displays the compass heading for its associated map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCompassButton
type MKCompassButton struct {
	View
}

// MKCompassButtonFrom constructs a [MKCompassButton] from an unsafe.Pointer.
//
// A specialized view that displays the compass heading for its associated map.
func MKCompassButtonFrom(ptr unsafe.Pointer) MKCompassButton {
	return MKCompassButton{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKCompassButton */

// Creates a compass button and associates it with the specified map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCompassButton/init(mapView:)
func NewMKCompassButtonWithMapView(mapView IMKMapView) MKCompassButton {
	rv := objc.Send[MKCompassButton](objc.ID(getMKCompassButtonClass().class), objc.Sel("compassButtonWithMapView:"), mapView)
	return rv
}/* debug [class_init_methods/constructor]: NewMKCompassButtonWithMapView */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKCompassButton */

// Creates a compass button and associates it with the specified map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCompassButton/init(mapView:)
func (mc _MKCompassButtonClass) CompassButtonWithMapView(mapView IMKMapView) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("compassButtonWithMapView:"), mapView)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CompassButtonWithMapView) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKCompassButton */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKCompassButton */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKCompassButton */

// The visibility of the compass button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCompassButton/compassVisibility
func (m_ MKCompassButton) CompassVisibility() MKFeatureVisibility {
	rv := objc.Send[MKFeatureVisibility](m_.ID, objc.Sel("compassVisibility"))
	return rv
}/* debug [instance_properties/getter]: compassVisibility */


// The visibility of the compass button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCompassButton/compassVisibility
func (m_ MKCompassButton) SetCompassVisibility(value MKFeatureVisibility) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCompassVisibility:"), value)
}/* debug [instance_properties/setter]: compassVisibility */


// The map view that provides the heading information for the compass button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCompassButton/mapView
func (m_ MKCompassButton) MapView() IMKMapView {
	rv := objc.Send[MKMapView](m_.ID, objc.Sel("mapView"))
	return rv
}/* debug [instance_properties/getter]: mapView */


// The map view that provides the heading information for the compass button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCompassButton/mapView
func (m_ MKCompassButton) SetMapView(value IMKMapView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapView:"), value)
}/* debug [instance_properties/setter]: mapView */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKCompassButton */


