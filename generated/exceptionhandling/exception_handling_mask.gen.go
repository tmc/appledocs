// Code generated from Apple documentation for ExceptionHandling. DO NOT EDIT.

package exceptionhandling

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class exceptionHandlingMask */


/* debug [class_header]: Header for exceptionHandlingMask */
// The class instance for the [exceptionHandlingMask] class.
var (
	ExceptionHandlingMaskClass     _exceptionHandlingMaskClass
	ExceptionHandlingMaskClassOnce sync.Once
)

func getexceptionHandlingMaskClass() _exceptionHandlingMaskClass {
	ExceptionHandlingMaskClassOnce.Do(func() {
		ExceptionHandlingMaskClass = _exceptionHandlingMaskClass{objc.GetClass("exceptionHandlingMask")}
	})
	return ExceptionHandlingMaskClass
}

type _exceptionHandlingMaskClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for exceptionHandlingMask */
// An interface definition for the [exceptionHandlingMask] class.
type IexceptionHandlingMask interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for exceptionHandlingMask */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for exceptionHandlingMask */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for exceptionHandlingMask */
// Alloc allocates a new instance without initialization.
func (ec _exceptionHandlingMaskClass) Alloc() exceptionHandlingMask {
	rv := objc.Send[exceptionHandlingMask](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _exceptionHandlingMaskClass) New() exceptionHandlingMask {
	rv := objc.Send[exceptionHandlingMask](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ exceptionHandlingMask) Init() exceptionHandlingMask {
	rv := objc.Send[exceptionHandlingMask](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ exceptionHandlingMask) Autorelease() exceptionHandlingMask {
	rv := objc.Send[exceptionHandlingMask](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewexceptionHandlingMask creates a new exceptionHandlingMask instance.
func NewexceptionHandlingMask() exceptionHandlingMask {
	return getexceptionHandlingMaskClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for exceptionHandlingMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/struct_(unnamed)/exceptionHandlingMask
type exceptionHandlingMask struct {
	objectivec.Object
}

// exceptionHandlingMaskFrom constructs a [exceptionHandlingMask] from an unsafe.Pointer.
func exceptionHandlingMaskFrom(ptr unsafe.Pointer) exceptionHandlingMask {
	return exceptionHandlingMask{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for exceptionHandlingMask *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for exceptionHandlingMask */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for exceptionHandlingMask */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for exceptionHandlingMask */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for exceptionHandlingMask */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class exceptionHandlingMask */



