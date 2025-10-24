// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MKOverlayPathView */


/* debug [class_header]: Header for MKOverlayPathView */
// The class instance for the [MKOverlayPathView] class.
var (
	MKOverlayPathViewClass     _MKOverlayPathViewClass
	MKOverlayPathViewClassOnce sync.Once
)

func getMKOverlayPathViewClass() _MKOverlayPathViewClass {
	MKOverlayPathViewClassOnce.Do(func() {
		MKOverlayPathViewClass = _MKOverlayPathViewClass{objc.GetClass("MKOverlayPathView")}
	})
	return MKOverlayPathViewClass
}

type _MKOverlayPathViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKOverlayPathView */
// An interface definition for the [MKOverlayPathView] class.
type IMKOverlayPathView interface {
	IMKOverlayView
	
/* debug [class_interface_properties]: Properties for MKOverlayPathView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKOverlayPathView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKOverlayPathView */
// Alloc allocates a new instance without initialization.
func (mc _MKOverlayPathViewClass) Alloc() MKOverlayPathView {
	rv := objc.Send[MKOverlayPathView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKOverlayPathViewClass) New() MKOverlayPathView {
	rv := objc.Send[MKOverlayPathView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKOverlayPathView) Init() MKOverlayPathView {
	rv := objc.Send[MKOverlayPathView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKOverlayPathView) Autorelease() MKOverlayPathView {
	rv := objc.Send[MKOverlayPathView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKOverlayPathView creates a new MKOverlayPathView instance.
func NewMKOverlayPathView() MKOverlayPathView {
	return getMKOverlayPathViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKOverlayPathView */
// Represents a generic overlay that draws its contents using a Core Graphics path data type.
//
// You can use this class to implement simple path-based overlay views or subclass it to define additional drawing behaviors. The default drawing behavior of this class is to apply the object’s current fill attributes, fill the path, apply the current stroke attributes, and then stroke the path. If you subclass, you should override the method and use that method to build the appropriate path for the overlay. You can invalidate this path as needed and force the path to be recreated using whatever new data your subclass has obtained. In iOS 7 and later, use the class to display path-based overlays instead.


// Represents a generic overlay that draws its contents using a Core Graphics path data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKOverlayPathView
type MKOverlayPathView struct {
	MKOverlayView
}

// MKOverlayPathViewFrom constructs a [MKOverlayPathView] from an unsafe.Pointer.
//
// Represents a generic overlay that draws its contents using a Core Graphics path data type.
func MKOverlayPathViewFrom(ptr unsafe.Pointer) MKOverlayPathView {
	return MKOverlayPathView{
		MKOverlayView: MKOverlayViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKOverlayPathView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKOverlayPathView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKOverlayPathView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKOverlayPathView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKOverlayPathView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKOverlayPathView */


