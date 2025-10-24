// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class leftEyePosition */


/* debug [class_header]: Header for leftEyePosition */
// The class instance for the [leftEyePosition] class.
var (
	LeftEyePositionClass     _leftEyePositionClass
	LeftEyePositionClassOnce sync.Once
)

func getleftEyePositionClass() _leftEyePositionClass {
	LeftEyePositionClassOnce.Do(func() {
		LeftEyePositionClass = _leftEyePositionClass{objc.GetClass("leftEyePosition")}
	})
	return LeftEyePositionClass
}

type _leftEyePositionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for leftEyePosition */
// An interface definition for the [leftEyePosition] class.
type IleftEyePosition interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for leftEyePosition */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for leftEyePosition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for leftEyePosition */
// Alloc allocates a new instance without initialization.
func (lc _leftEyePositionClass) Alloc() leftEyePosition {
	rv := objc.Send[leftEyePosition](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _leftEyePositionClass) New() leftEyePosition {
	rv := objc.Send[leftEyePosition](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ leftEyePosition) Init() leftEyePosition {
	rv := objc.Send[leftEyePosition](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ leftEyePosition) Autorelease() leftEyePosition {
	rv := objc.Send[leftEyePosition](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewleftEyePosition creates a new leftEyePosition instance.
func NewleftEyePosition() leftEyePosition {
	return getleftEyePositionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for leftEyePosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/leftEyePosition-c.ivar
type leftEyePosition struct {
	objectivec.Object
}

// leftEyePositionFrom constructs a [leftEyePosition] from an unsafe.Pointer.
func leftEyePositionFrom(ptr unsafe.Pointer) leftEyePosition {
	return leftEyePosition{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for leftEyePosition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for leftEyePosition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for leftEyePosition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for leftEyePosition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for leftEyePosition */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class leftEyePosition */



