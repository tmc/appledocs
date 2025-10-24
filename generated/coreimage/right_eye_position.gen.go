// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class rightEyePosition */


/* debug [class_header]: Header for rightEyePosition */
// The class instance for the [rightEyePosition] class.
var (
	RightEyePositionClass     _rightEyePositionClass
	RightEyePositionClassOnce sync.Once
)

func getrightEyePositionClass() _rightEyePositionClass {
	RightEyePositionClassOnce.Do(func() {
		RightEyePositionClass = _rightEyePositionClass{objc.GetClass("rightEyePosition")}
	})
	return RightEyePositionClass
}

type _rightEyePositionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for rightEyePosition */
// An interface definition for the [rightEyePosition] class.
type IrightEyePosition interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for rightEyePosition */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for rightEyePosition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for rightEyePosition */
// Alloc allocates a new instance without initialization.
func (rc _rightEyePositionClass) Alloc() rightEyePosition {
	rv := objc.Send[rightEyePosition](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _rightEyePositionClass) New() rightEyePosition {
	rv := objc.Send[rightEyePosition](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ rightEyePosition) Init() rightEyePosition {
	rv := objc.Send[rightEyePosition](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ rightEyePosition) Autorelease() rightEyePosition {
	rv := objc.Send[rightEyePosition](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewrightEyePosition creates a new rightEyePosition instance.
func NewrightEyePosition() rightEyePosition {
	return getrightEyePositionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for rightEyePosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/rightEyePosition-c.ivar
type rightEyePosition struct {
	objectivec.Object
}

// rightEyePositionFrom constructs a [rightEyePosition] from an unsafe.Pointer.
func rightEyePositionFrom(ptr unsafe.Pointer) rightEyePosition {
	return rightEyePosition{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for rightEyePosition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for rightEyePosition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for rightEyePosition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for rightEyePosition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for rightEyePosition */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class rightEyePosition */



