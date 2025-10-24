// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MKUserLocationView */


/* debug [class_header]: Header for MKUserLocationView */
// The class instance for the [MKUserLocationView] class.
var (
	MKUserLocationViewClass     _MKUserLocationViewClass
	MKUserLocationViewClassOnce sync.Once
)

func getMKUserLocationViewClass() _MKUserLocationViewClass {
	MKUserLocationViewClassOnce.Do(func() {
		MKUserLocationViewClass = _MKUserLocationViewClass{objc.GetClass("MKUserLocationView")}
	})
	return MKUserLocationViewClass
}

type _MKUserLocationViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKUserLocationView */
// An interface definition for the [MKUserLocationView] class.
type IMKUserLocationView interface {
	IMKAnnotationView
	
/* debug [class_interface_properties]: Properties for MKUserLocationView */
	// properties:
	ZPriority() MKAnnotationViewZPriority /* typedef */
	SetZPriority(value MKAnnotationViewZPriority /* typedef */)
	ShowsUserLocation() bool
	SetShowsUserLocation(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKUserLocationView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKUserLocationView */
// Alloc allocates a new instance without initialization.
func (mc _MKUserLocationViewClass) Alloc() MKUserLocationView {
	rv := objc.Send[MKUserLocationView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKUserLocationViewClass) New() MKUserLocationView {
	rv := objc.Send[MKUserLocationView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKUserLocationView) Init() MKUserLocationView {
	rv := objc.Send[MKUserLocationView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKUserLocationView) Autorelease() MKUserLocationView {
	rv := objc.Send[MKUserLocationView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKUserLocationView creates a new MKUserLocationView instance.
func NewMKUserLocationView() MKUserLocationView {
	return getMKUserLocationViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKUserLocationView */
// A configurable annotation that shows the user’s location using the default MapKit style.
//
// If you don’t need additional configuration, you can show an annotation with the user’s location by setting on the map to . If you want to specify additional configuration, such as , create this annotation view directly. To display the annotation view, return the instance from . The user location view provides the MapKit default style and behavior. The visual display varies with the level of authorization the user grants your app.


// A configurable annotation that shows the user’s location using the default MapKit style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKUserLocationView
type MKUserLocationView struct {
	MKAnnotationView
}

// MKUserLocationViewFrom constructs a [MKUserLocationView] from an unsafe.Pointer.
//
// A configurable annotation that shows the user’s location using the default MapKit style.
func MKUserLocationViewFrom(ptr unsafe.Pointer) MKUserLocationView {
	return MKUserLocationView{
		MKAnnotationView: MKAnnotationViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKUserLocationView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKUserLocationView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKUserLocationView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKUserLocationView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKUserLocationView */

// The relative importance of the annotation view when in an unselected state with respect to its ordering along the z-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/zpriority
func (m_ MKUserLocationView) ZPriority() MKAnnotationViewZPriority /* typedef */ {
	rv := objc.Send[float32](m_.ID, objc.Sel("zPriority"))
	return rv
}/* debug [instance_properties/getter]: zPriority */


// The relative importance of the annotation view when in an unselected state with respect to its ordering along the z-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkannotationview/zpriority
func (m_ MKUserLocationView) SetZPriority(value MKAnnotationViewZPriority /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setZPriority:"), value)
}/* debug [instance_properties/setter]: zPriority */


// A Boolean value that indicates whether the map tries to display the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsuserlocation
func (m_ MKUserLocationView) ShowsUserLocation() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsUserLocation"))
	return rv
}/* debug [instance_properties/getter]: showsUserLocation */


// A Boolean value that indicates whether the map tries to display the user’s location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mapkit/mkmapview/showsuserlocation
func (m_ MKUserLocationView) SetShowsUserLocation(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsUserLocation:"), value)
}/* debug [instance_properties/setter]: showsUserLocation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKUserLocationView */



