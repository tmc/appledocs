// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKPitchControl */


/* debug [class_header]: Header for MKPitchControl */
// The class instance for the [MKPitchControl] class.
var (
	MKPitchControlClass     _MKPitchControlClass
	MKPitchControlClassOnce sync.Once
)

func getMKPitchControlClass() _MKPitchControlClass {
	MKPitchControlClassOnce.Do(func() {
		MKPitchControlClass = _MKPitchControlClass{objc.GetClass("MKPitchControl")}
	})
	return MKPitchControlClass
}

type _MKPitchControlClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKPitchControl */
// An interface definition for the [MKPitchControl] class.
type IMKPitchControl interface {
	IView
	
/* debug [class_interface_properties]: Properties for MKPitchControl */
	// properties:
	MapView() IMKMapView
	SetMapView(value IMKMapView)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKPitchControl */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKPitchControl */
// Alloc allocates a new instance without initialization.
func (mc _MKPitchControlClass) Alloc() MKPitchControl {
	rv := objc.Send[MKPitchControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKPitchControlClass) New() MKPitchControl {
	rv := objc.Send[MKPitchControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKPitchControl) Init() MKPitchControl {
	rv := objc.Send[MKPitchControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKPitchControl) Autorelease() MKPitchControl {
	rv := objc.Send[MKPitchControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKPitchControl creates a new MKPitchControl instance.
func NewMKPitchControl() MKPitchControl {
	return getMKPitchControlClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKPitchControl */
// A specialized view that displays and controls the pitch angle of the map view.
//
// Use this class when you want to incorporate a standard, fixed-size pitch control into your own view hierarchy. A pitch control allows the user to change the pitch angle of its associated map view.


// A specialized view that displays and controls the pitch angle of the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPitchControl
type MKPitchControl struct {
	View
}

// MKPitchControlFrom constructs a [MKPitchControl] from an unsafe.Pointer.
//
// A specialized view that displays and controls the pitch angle of the map view.
func MKPitchControlFrom(ptr unsafe.Pointer) MKPitchControl {
	return MKPitchControl{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKPitchControl */

// Creates a pitch control and associates it with the specified map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPitchControl/init(mapView:)
func NewMKPitchControlWithMapView(mapView IMKMapView) MKPitchControl {
	rv := objc.Send[MKPitchControl](objc.ID(getMKPitchControlClass().class), objc.Sel("pitchControlWithMapView:"), mapView)
	return rv
}/* debug [class_init_methods/constructor]: NewMKPitchControlWithMapView */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKPitchControl */

// Creates a pitch control and associates it with the specified map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPitchControl/init(mapView:)
func (mc _MKPitchControlClass) PitchControlWithMapView(mapView IMKMapView) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("pitchControlWithMapView:"), mapView)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PitchControlWithMapView) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKPitchControl */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKPitchControl */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKPitchControl */

// The map view associated with this control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPitchControl/mapView
func (m_ MKPitchControl) MapView() IMKMapView {
	rv := objc.Send[MKMapView](m_.ID, objc.Sel("mapView"))
	return rv
}/* debug [instance_properties/getter]: mapView */


// The map view associated with this control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPitchControl/mapView
func (m_ MKPitchControl) SetMapView(value IMKMapView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapView:"), value)
}/* debug [instance_properties/setter]: mapView */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKPitchControl */


