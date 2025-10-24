// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class leftEyeClosed */


/* debug [class_header]: Header for leftEyeClosed */
// The class instance for the [leftEyeClosed] class.
var (
	LeftEyeClosedClass     _leftEyeClosedClass
	LeftEyeClosedClassOnce sync.Once
)

func getleftEyeClosedClass() _leftEyeClosedClass {
	LeftEyeClosedClassOnce.Do(func() {
		LeftEyeClosedClass = _leftEyeClosedClass{objc.GetClass("leftEyeClosed")}
	})
	return LeftEyeClosedClass
}

type _leftEyeClosedClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for leftEyeClosed */
// An interface definition for the [leftEyeClosed] class.
type IleftEyeClosed interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for leftEyeClosed */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for leftEyeClosed */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for leftEyeClosed */
// Alloc allocates a new instance without initialization.
func (lc _leftEyeClosedClass) Alloc() leftEyeClosed {
	rv := objc.Send[leftEyeClosed](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _leftEyeClosedClass) New() leftEyeClosed {
	rv := objc.Send[leftEyeClosed](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ leftEyeClosed) Init() leftEyeClosed {
	rv := objc.Send[leftEyeClosed](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ leftEyeClosed) Autorelease() leftEyeClosed {
	rv := objc.Send[leftEyeClosed](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewleftEyeClosed creates a new leftEyeClosed instance.
func NewleftEyeClosed() leftEyeClosed {
	return getleftEyeClosedClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for leftEyeClosed */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/leftEyeClosed-c.ivar
type leftEyeClosed struct {
	objectivec.Object
}

// leftEyeClosedFrom constructs a [leftEyeClosed] from an unsafe.Pointer.
func leftEyeClosedFrom(ptr unsafe.Pointer) leftEyeClosed {
	return leftEyeClosed{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for leftEyeClosed *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for leftEyeClosed */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for leftEyeClosed */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for leftEyeClosed */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for leftEyeClosed */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class leftEyeClosed */



