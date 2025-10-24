// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKPolylineView */


/* debug [class_header]: Header for MKPolylineView */
// The class instance for the [MKPolylineView] class.
var (
	MKPolylineViewClass     _MKPolylineViewClass
	MKPolylineViewClassOnce sync.Once
)

func getMKPolylineViewClass() _MKPolylineViewClass {
	MKPolylineViewClassOnce.Do(func() {
		MKPolylineViewClass = _MKPolylineViewClass{objc.GetClass("MKPolylineView")}
	})
	return MKPolylineViewClass
}

type _MKPolylineViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKPolylineView */
// An interface definition for the [MKPolylineView] class.
type IMKPolylineView interface {
	IMKOverlayPathView
	
/* debug [class_interface_properties]: Properties for MKPolylineView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKPolylineView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKPolylineView */
// Alloc allocates a new instance without initialization.
func (mc _MKPolylineViewClass) Alloc() MKPolylineView {
	rv := objc.Send[MKPolylineView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKPolylineViewClass) New() MKPolylineView {
	rv := objc.Send[MKPolylineView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKPolylineView) Init() MKPolylineView {
	rv := objc.Send[MKPolylineView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKPolylineView) Autorelease() MKPolylineView {
	rv := objc.Send[MKPolylineView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKPolylineView creates a new MKPolylineView instance.
func NewMKPolylineView() MKPolylineView {
	return getMKPolylineViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKPolylineView */
// Provides the visual representation for an annotation object.
//
// This view strokes the path represented by the annotation. (This class does not fill the area enclosed by the path.) You can change the color and other drawing attributes of the path by modifying the properties inherited from the class. This class is typically used as is and not subclassed. In iOS 7 and later, use the class to display polyline overlays instead.


// Provides the visual representation for an annotation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolylineView
type MKPolylineView struct {
	MKOverlayPathView
}

// MKPolylineViewFrom constructs a [MKPolylineView] from an unsafe.Pointer.
//
// Provides the visual representation for an annotation object.
func MKPolylineViewFrom(ptr unsafe.Pointer) MKPolylineView {
	return MKPolylineView{
		MKOverlayPathView: MKOverlayPathViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKPolylineView */

// Initializes and returns a new overlay view using the specified polyline overlay object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolylineView/initWithPolyline:
func NewMKPolylineViewWithPolyline(polyline IMKPolyline) MKPolylineView {
	instance := getMKPolylineViewClass().Alloc()
	rv := objc.Send[MKPolylineView](instance.ID, objc.Sel("initWithPolyline:"), polyline)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKPolylineViewWithPolyline */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKPolylineView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKPolylineView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKPolylineView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKPolylineView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKPolylineView */


