// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKUserTrackingButton */


/* debug [class_header]: Header for MKUserTrackingButton */
// The class instance for the [MKUserTrackingButton] class.
var (
	MKUserTrackingButtonClass     _MKUserTrackingButtonClass
	MKUserTrackingButtonClassOnce sync.Once
)

func getMKUserTrackingButtonClass() _MKUserTrackingButtonClass {
	MKUserTrackingButtonClassOnce.Do(func() {
		MKUserTrackingButtonClass = _MKUserTrackingButtonClass{objc.GetClass("MKUserTrackingButton")}
	})
	return MKUserTrackingButtonClass
}

type _MKUserTrackingButtonClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKUserTrackingButton */
// An interface definition for the [MKUserTrackingButton] class.
type IMKUserTrackingButton interface {
	IView
	
/* debug [class_interface_properties]: Properties for MKUserTrackingButton */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKUserTrackingButton */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKUserTrackingButton */
// Alloc allocates a new instance without initialization.
func (mc _MKUserTrackingButtonClass) Alloc() MKUserTrackingButton {
	rv := objc.Send[MKUserTrackingButton](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKUserTrackingButtonClass) New() MKUserTrackingButton {
	rv := objc.Send[MKUserTrackingButton](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKUserTrackingButton) Init() MKUserTrackingButton {
	rv := objc.Send[MKUserTrackingButton](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKUserTrackingButton) Autorelease() MKUserTrackingButton {
	rv := objc.Send[MKUserTrackingButton](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKUserTrackingButton creates a new MKUserTrackingButton instance.
func NewMKUserTrackingButton() MKUserTrackingButton {
	return getMKUserTrackingButtonClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKUserTrackingButton */
// A specialized button that allows the user to toggle whether the map tracks to the heading the user is facing.
//
// Use this class when you need a standard button that you can incorporate into your view hierarchy. Tapping the button lets the user toggles between modes for displaying the map with and without the current heading applied. The button also reflects the current user tracking mode if set elsewhere.


// A specialized button that allows the user to toggle whether the map tracks to the heading the user is facing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserTrackingButton
type MKUserTrackingButton struct {
	View
}

// MKUserTrackingButtonFrom constructs a [MKUserTrackingButton] from an unsafe.Pointer.
//
// A specialized button that allows the user to toggle whether the map tracks to the heading the user is facing.
func MKUserTrackingButtonFrom(ptr unsafe.Pointer) MKUserTrackingButton {
	return MKUserTrackingButton{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKUserTrackingButton */

// Initializes the button with the map view that it should control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserTrackingButton/init(mapView:)
func NewMKUserTrackingButtonWithMapView(mapView IMKMapView) MKUserTrackingButton {
	rv := objc.Send[MKUserTrackingButton](objc.ID(getMKUserTrackingButtonClass().class), objc.Sel("userTrackingButtonWithMapView:"), mapView)
	return rv
}/* debug [class_init_methods/constructor]: NewMKUserTrackingButtonWithMapView */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKUserTrackingButton */

// Initializes the button with the map view that it should control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserTrackingButton/init(mapView:)
func (mc _MKUserTrackingButtonClass) UserTrackingButtonWithMapView(mapView IMKMapView) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("userTrackingButtonWithMapView:"), mapView)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UserTrackingButtonWithMapView) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKUserTrackingButton */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKUserTrackingButton */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKUserTrackingButton */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKUserTrackingButton */


