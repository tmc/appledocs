// Code generated from Apple documentation for ExceptionHandling. DO NOT EDIT.

package exceptionhandling

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class exceptionHangingMask */


/* debug [class_header]: Header for exceptionHangingMask */
// The class instance for the [exceptionHangingMask] class.
var (
	ExceptionHangingMaskClass     _exceptionHangingMaskClass
	ExceptionHangingMaskClassOnce sync.Once
)

func getexceptionHangingMaskClass() _exceptionHangingMaskClass {
	ExceptionHangingMaskClassOnce.Do(func() {
		ExceptionHangingMaskClass = _exceptionHangingMaskClass{objc.GetClass("exceptionHangingMask")}
	})
	return ExceptionHangingMaskClass
}

type _exceptionHangingMaskClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for exceptionHangingMask */
// An interface definition for the [exceptionHangingMask] class.
type IexceptionHangingMask interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for exceptionHangingMask */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for exceptionHangingMask */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for exceptionHangingMask */
// Alloc allocates a new instance without initialization.
func (ec _exceptionHangingMaskClass) Alloc() exceptionHangingMask {
	rv := objc.Send[exceptionHangingMask](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _exceptionHangingMaskClass) New() exceptionHangingMask {
	rv := objc.Send[exceptionHangingMask](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ exceptionHangingMask) Init() exceptionHangingMask {
	rv := objc.Send[exceptionHangingMask](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ exceptionHangingMask) Autorelease() exceptionHangingMask {
	rv := objc.Send[exceptionHangingMask](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewexceptionHangingMask creates a new exceptionHangingMask instance.
func NewexceptionHangingMask() exceptionHangingMask {
	return getexceptionHangingMaskClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for exceptionHangingMask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/struct_(unnamed)/exceptionHangingMask
type exceptionHangingMask struct {
	objectivec.Object
}

// exceptionHangingMaskFrom constructs a [exceptionHangingMask] from an unsafe.Pointer.
func exceptionHangingMaskFrom(ptr unsafe.Pointer) exceptionHangingMask {
	return exceptionHangingMask{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for exceptionHangingMask *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for exceptionHangingMask */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for exceptionHangingMask */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for exceptionHangingMask */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for exceptionHangingMask */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class exceptionHangingMask */



