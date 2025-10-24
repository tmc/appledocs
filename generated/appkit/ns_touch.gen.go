// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSTouch */


/* debug [class_header]: Header for NSTouch */
// The class instance for the [Touch] class.
var (
	TouchClass     _TouchClass
	TouchClassOnce sync.Once
)

func getTouchClass() _TouchClass {
	TouchClassOnce.Do(func() {
		TouchClass = _TouchClass{objc.GetClass("NSTouch")}
	})
	return TouchClass
}

type _TouchClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Touch */
// An interface definition for the [Touch] class.
type ITouch interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Touch */
	// properties:
	Identity() unsafe.Pointer
	Resting() bool
	NormalizedPosition() vision.Point
	Phase() TouchPhase
	Type() TouchType
	DeviceSize() Size /* not a class type */
	SetDeviceSize(value Size /* not a class type */)
	IsResting() bool
	SetIsResting(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Touch */
	// methods:
	PreviousLocationInView(view IView) vision.Point
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Touch */
// Alloc allocates a new instance without initialization.
func (tc _TouchClass) Alloc() Touch {
	rv := objc.Send[Touch](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TouchClass) New() Touch {
	rv := objc.Send[Touch](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Touch) Init() Touch {
	rv := objc.Send[Touch](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Touch) Autorelease() Touch {
	rv := objc.Send[Touch](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTouch creates a new Touch instance.
func NewTouch() Touch {
	return getTouchClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Touch */
// A snapshot of a particular touch at an instant in time.
//
// A touch event is not persistent throughout the touch. A touch creates new instances as it progresses. Use the identity property to follow a specific touch across its lifetime. Touches do not have a corresponding screen location. The first touch of a touch collection latches to the view underlying the cursor using the same hit detection as mouse events. Additional touches on the same device latch to the same view. Latches remain on views until the user ends a touch or an event cancels it.


// A snapshot of a particular touch at an instant in time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch
type Touch struct {
	objectivec.Object
}

// TouchFrom constructs a [Touch] from an unsafe.Pointer.
//
// A snapshot of a particular touch at an instant in time.
func TouchFrom(ptr unsafe.Pointer) Touch {
	return Touch{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Touch *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Touch */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Touch */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Touch */

// Indicates the previous location of the touch in the view’s coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/previousLocation(in:)
func (t_ Touch) PreviousLocationInView(view IView) vision.Point {
	rv := objc.Send[vision.Point](t_.ID, objc.Sel("previousLocationInView:"), view)
	return rv
}/* debug [instance_methods/method]: PreviousLocationInView */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Touch */

// The changes to a particular touch during its lifetime.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/identity
func (t_ Touch) Identity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("identity"))
	return rv
}/* debug [instance_properties/getter]: identity */


// The indicator for a resting touch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/isResting
func (t_ Touch) Resting() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("resting"))
	return rv
}/* debug [instance_properties/getter]: resting */


// The normalized position of the touch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/normalizedPosition
func (t_ Touch) NormalizedPosition() vision.Point {
	rv := objc.Send[vision.Point](t_.ID, objc.Sel("normalizedPosition"))
	return rv
}/* debug [instance_properties/getter]: normalizedPosition */


// The current phase of the touch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/phase-swift.property
func (t_ Touch) Phase() TouchPhase {
	rv := objc.Send[TouchPhase](t_.ID, objc.Sel("phase"))
	return rv
}/* debug [instance_properties/getter]: phase */


// A type of touch from a Touch Bar interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/type
func (t_ Touch) Type() TouchType {
	rv := objc.Send[TouchType](t_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The range of the touch device in points, such as 72 ppi.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/devicesize
func (t_ Touch) DeviceSize() Size /* not a class type */ {
	rv := objc.Send[Size](t_.ID, objc.Sel("deviceSize"))
	return rv
}/* debug [instance_properties/getter]: deviceSize */


// The range of the touch device in points, such as 72 ppi.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/devicesize
func (t_ Touch) SetDeviceSize(value Size /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDeviceSize:"), value)
}/* debug [instance_properties/setter]: deviceSize */


// The indicator for a resting touch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/isresting
func (t_ Touch) IsResting() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isResting"))
	return rv
}/* debug [instance_properties/getter]: isResting */


// The indicator for a resting touch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouch/isresting
func (t_ Touch) SetIsResting(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsResting:"), value)
}/* debug [instance_properties/setter]: isResting */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTouch */



