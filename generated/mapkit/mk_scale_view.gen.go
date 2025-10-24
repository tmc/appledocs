// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKScaleView */


/* debug [class_header]: Header for MKScaleView */
// The class instance for the [MKScaleView] class.
var (
	MKScaleViewClass     _MKScaleViewClass
	MKScaleViewClassOnce sync.Once
)

func getMKScaleViewClass() _MKScaleViewClass {
	MKScaleViewClassOnce.Do(func() {
		MKScaleViewClass = _MKScaleViewClass{objc.GetClass("MKScaleView")}
	})
	return MKScaleViewClass
}

type _MKScaleViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKScaleView */
// An interface definition for the [MKScaleView] class.
type IMKScaleView interface {
	IView
	
/* debug [class_interface_properties]: Properties for MKScaleView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKScaleView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKScaleView */
// Alloc allocates a new instance without initialization.
func (mc _MKScaleViewClass) Alloc() MKScaleView {
	rv := objc.Send[MKScaleView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKScaleViewClass) New() MKScaleView {
	rv := objc.Send[MKScaleView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKScaleView) Init() MKScaleView {
	rv := objc.Send[MKScaleView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKScaleView) Autorelease() MKScaleView {
	rv := objc.Send[MKScaleView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKScaleView creates a new MKScaleView instance.
func NewMKScaleView() MKScaleView {
	return getMKScaleViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKScaleView */
// A specialized view that displays the scale information for its associated map.
//
// Use this class when you want to incorporate a standard scale view into your own view hierarchy. A scale view displays a legend with distance information for its associated map view. As the map region changes, the scale view updates automatically to reflect any changes in scale.


// A specialized view that displays the scale information for its associated map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKScaleView
type MKScaleView struct {
	View
}

// MKScaleViewFrom constructs a [MKScaleView] from an unsafe.Pointer.
//
// A specialized view that displays the scale information for its associated map.
func MKScaleViewFrom(ptr unsafe.Pointer) MKScaleView {
	return MKScaleView{
		View: ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKScaleView */

// Creates a scale view and associates it with the specified map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKScaleView/init(mapView:)
func NewMKScaleViewWithMapView(mapView IMKMapView) MKScaleView {
	rv := objc.Send[MKScaleView](objc.ID(getMKScaleViewClass().class), objc.Sel("scaleViewWithMapView:"), mapView)
	return rv
}/* debug [class_init_methods/constructor]: NewMKScaleViewWithMapView */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKScaleView */

// Creates a scale view and associates it with the specified map view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKScaleView/init(mapView:)
func (mc _MKScaleViewClass) ScaleViewWithMapView(mapView IMKMapView) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("scaleViewWithMapView:"), mapView)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ScaleViewWithMapView) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKScaleView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKScaleView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKScaleView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKScaleView */


