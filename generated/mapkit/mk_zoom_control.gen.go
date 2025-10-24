// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKZoomControl */


/* debug [class_header]: Header for MKZoomControl */
// The class instance for the [MKZoomControl] class.
var (
	MKZoomControlClass     _MKZoomControlClass
	MKZoomControlClassOnce sync.Once
)

func getMKZoomControlClass() _MKZoomControlClass {
	MKZoomControlClassOnce.Do(func() {
		MKZoomControlClass = _MKZoomControlClass{objc.GetClass("MKZoomControl")}
	})
	return MKZoomControlClass
}

type _MKZoomControlClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKZoomControl */
// An interface definition for the [MKZoomControl] class.
type IMKZoomControl interface {
	IView
	
/* debug [class_interface_properties]: Properties for MKZoomControl */
	// properties:
	MapView() IMKMapView
	SetMapView(value IMKMapView)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKZoomControl */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKZoomControl */
// Alloc allocates a new instance without initialization.
func (mc _MKZoomControlClass) Alloc() MKZoomControl {
	rv := objc.Send[MKZoomControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKZoomControlClass) New() MKZoomControl {
	rv := objc.Send[MKZoomControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKZoomControl) Init() MKZoomControl {
	rv := objc.Send[MKZoomControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKZoomControl) Autorelease() MKZoomControl {
	rv := objc.Send[MKZoomControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKZoomControl creates a new MKZoomControl instance.
func NewMKZoomControl() MKZoomControl {
	return getMKZoomControlClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKZoomControl */
// A specialized view that displays and controls the zoom level of the map view.
//
// Use this class when you want to incorporate a standard, fixed-size zoom control into your own view hierarchy. A zoom control enables the user to change the zoom level of its associated map view.


// A specialized view that displays and controls the zoom level of the map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKZoomControl
type MKZoomControl struct {
	View
}

// MKZoomControlFrom constructs a [MKZoomControl] from an unsafe.Pointer.
//
// A specialized view that displays and controls the zoom level of the map view.
func MKZoomControlFrom(ptr unsafe.Pointer) MKZoomControl {
	return MKZoomControl{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKZoomControl */

// Creates a zoom control and associates it with the specified map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKZoomControl/init(mapView:)
func NewMKZoomControlWithMapView(mapView IMKMapView) MKZoomControl {
	rv := objc.Send[MKZoomControl](objc.ID(getMKZoomControlClass().class), objc.Sel("zoomControlWithMapView:"), mapView)
	return rv
}/* debug [class_init_methods/constructor]: NewMKZoomControlWithMapView */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKZoomControl */

// Creates a zoom control and associates it with the specified map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKZoomControl/init(mapView:)
func (mc _MKZoomControlClass) ZoomControlWithMapView(mapView IMKMapView) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("zoomControlWithMapView:"), mapView)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ZoomControlWithMapView) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKZoomControl */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKZoomControl */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKZoomControl */

// The map view associated with this control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKZoomControl/mapView
func (m_ MKZoomControl) MapView() IMKMapView {
	rv := objc.Send[MKMapView](m_.ID, objc.Sel("mapView"))
	return rv
}/* debug [instance_properties/getter]: mapView */


// The map view associated with this control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKZoomControl/mapView
func (m_ MKZoomControl) SetMapView(value IMKMapView) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapView:"), value)
}/* debug [instance_properties/setter]: mapView */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKZoomControl */


