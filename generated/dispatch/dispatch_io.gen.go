// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DispatchIO */


/* debug [class_header]: Header for DispatchIO */
// The class instance for the [DispatchIO] class.
var (
	DispatchIOClass     _DispatchIOClass
	DispatchIOClassOnce sync.Once
)

func getDispatchIOClass() _DispatchIOClass {
	DispatchIOClassOnce.Do(func() {
		DispatchIOClass = _DispatchIOClass{objc.GetClass("DispatchIO")}
	})
	return DispatchIOClass
}

type _DispatchIOClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DispatchIO */
// An interface definition for the [DispatchIO] class.
type IDispatchIO interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DispatchIO */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DispatchIO */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DispatchIO */
// Alloc allocates a new instance without initialization.
func (dc _DispatchIOClass) Alloc() DispatchIO {
	rv := objc.Send[DispatchIO](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DispatchIOClass) New() DispatchIO {
	rv := objc.Send[DispatchIO](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DispatchIO) Init() DispatchIO {
	rv := objc.Send[DispatchIO](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DispatchIO) Autorelease() DispatchIO {
	rv := objc.Send[DispatchIO](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDispatchIO creates a new DispatchIO instance.
func NewDispatchIO() DispatchIO {
	return getDispatchIOClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DispatchIO */
// An object that manages operations on a file descriptor using either stream-based or random-access semantics.


// An object that manages operations on a file descriptor using either stream-based or random-access semantics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchIO
type DispatchIO struct {
	objectivec.Object
}

// DispatchIOFrom constructs a [DispatchIO] from an unsafe.Pointer.
//
// An object that manages operations on a file descriptor using either stream-based or random-access semantics.
func DispatchIOFrom(ptr unsafe.Pointer) DispatchIO {
	return DispatchIO{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DispatchIO *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DispatchIO */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DispatchIO */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DispatchIO */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DispatchIO */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DispatchIO */



