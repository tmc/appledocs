// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKPolygonView */


/* debug [class_header]: Header for MKPolygonView */
// The class instance for the [MKPolygonView] class.
var (
	MKPolygonViewClass     _MKPolygonViewClass
	MKPolygonViewClassOnce sync.Once
)

func getMKPolygonViewClass() _MKPolygonViewClass {
	MKPolygonViewClassOnce.Do(func() {
		MKPolygonViewClass = _MKPolygonViewClass{objc.GetClass("MKPolygonView")}
	})
	return MKPolygonViewClass
}

type _MKPolygonViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKPolygonView */
// An interface definition for the [MKPolygonView] class.
type IMKPolygonView interface {
	IMKOverlayPathView
	
/* debug [class_interface_properties]: Properties for MKPolygonView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKPolygonView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKPolygonView */
// Alloc allocates a new instance without initialization.
func (mc _MKPolygonViewClass) Alloc() MKPolygonView {
	rv := objc.Send[MKPolygonView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKPolygonViewClass) New() MKPolygonView {
	rv := objc.Send[MKPolygonView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKPolygonView) Init() MKPolygonView {
	rv := objc.Send[MKPolygonView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKPolygonView) Autorelease() MKPolygonView {
	rv := objc.Send[MKPolygonView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKPolygonView creates a new MKPolygonView instance.
func NewMKPolygonView() MKPolygonView {
	return getMKPolygonViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKPolygonView */
// Provides the visual representation for an annotation object.
//
// This view fills and strokes the area represented by the annotation. You can change the color and other drawing attributes of the polygon by modifying the properties inherited from the class. This class is typically used as is and not subclassed. In iOS 7 and later, use the class to display polygon overlays instead.


// Provides the visual representation for an annotation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygonView
type MKPolygonView struct {
	MKOverlayPathView
}

// MKPolygonViewFrom constructs a [MKPolygonView] from an unsafe.Pointer.
//
// Provides the visual representation for an annotation object.
func MKPolygonViewFrom(ptr unsafe.Pointer) MKPolygonView {
	return MKPolygonView{
		MKOverlayPathView: MKOverlayPathViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKPolygonView */

// Initializes and returns a new overlay view using the specified polygon overlay object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKPolygonView/initWithPolygon:
func NewMKPolygonViewWithPolygon(polygon IMKPolygon) MKPolygonView {
	instance := getMKPolygonViewClass().Alloc()
	rv := objc.Send[MKPolygonView](instance.ID, objc.Sel("initWithPolygon:"), polygon)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKPolygonViewWithPolygon */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKPolygonView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKPolygonView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKPolygonView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKPolygonView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKPolygonView */


