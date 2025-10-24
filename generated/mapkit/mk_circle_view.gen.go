// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MKCircleView */


/* debug [class_header]: Header for MKCircleView */
// The class instance for the [MKCircleView] class.
var (
	MKCircleViewClass     _MKCircleViewClass
	MKCircleViewClassOnce sync.Once
)

func getMKCircleViewClass() _MKCircleViewClass {
	MKCircleViewClassOnce.Do(func() {
		MKCircleViewClass = _MKCircleViewClass{objc.GetClass("MKCircleView")}
	})
	return MKCircleViewClass
}

type _MKCircleViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MKCircleView */
// An interface definition for the [MKCircleView] class.
type IMKCircleView interface {
	IMKOverlayPathView
	
/* debug [class_interface_properties]: Properties for MKCircleView */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MKCircleView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MKCircleView */
// Alloc allocates a new instance without initialization.
func (mc _MKCircleViewClass) Alloc() MKCircleView {
	rv := objc.Send[MKCircleView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MKCircleViewClass) New() MKCircleView {
	rv := objc.Send[MKCircleView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MKCircleView) Init() MKCircleView {
	rv := objc.Send[MKCircleView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MKCircleView) Autorelease() MKCircleView {
	rv := objc.Send[MKCircleView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKCircleView creates a new MKCircleView instance.
func NewMKCircleView() MKCircleView {
	return getMKCircleViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MKCircleView */
// Provides the visual representation for an annotation object.
//
// This view fills and strokes the circle represented by the annotation. You can change the color and other drawing attributes of the circle by modifying the properties inherited from the class. This class is typically used as is and not subclassed. Use of this class is discouraged in iOS 7 and later. Use the class instead.


// Provides the visual representation for an annotation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleView
type MKCircleView struct {
	MKOverlayPathView
}

// MKCircleViewFrom constructs a [MKCircleView] from an unsafe.Pointer.
//
// Provides the visual representation for an annotation object.
func MKCircleViewFrom(ptr unsafe.Pointer) MKCircleView {
	return MKCircleView{
		MKOverlayPathView: MKOverlayPathViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MKCircleView */

// Initializes and returns a new overlay view using the specified circle overlay object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MapKit/MKCircleView/initWithCircle:
func NewMKCircleViewWithCircle(circle IMKCircle) MKCircleView {
	instance := getMKCircleViewClass().Alloc()
	rv := objc.Send[MKCircleView](instance.ID, objc.Sel("initWithCircle:"), circle)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMKCircleViewWithCircle */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MKCircleView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MKCircleView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MKCircleView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MKCircleView */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MKCircleView */


