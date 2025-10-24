// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DispatchSerialQueue */


/* debug [class_header]: Header for DispatchSerialQueue */
// The class instance for the [DispatchSerialQueue] class.
var (
	DispatchSerialQueueClass     _DispatchSerialQueueClass
	DispatchSerialQueueClassOnce sync.Once
)

func getDispatchSerialQueueClass() _DispatchSerialQueueClass {
	DispatchSerialQueueClassOnce.Do(func() {
		DispatchSerialQueueClass = _DispatchSerialQueueClass{objc.GetClass("DispatchSerialQueue")}
	})
	return DispatchSerialQueueClass
}

type _DispatchSerialQueueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DispatchSerialQueue */
// An interface definition for the [DispatchSerialQueue] class.
type IDispatchSerialQueue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DispatchSerialQueue */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DispatchSerialQueue */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DispatchSerialQueue */
// Alloc allocates a new instance without initialization.
func (dc _DispatchSerialQueueClass) Alloc() DispatchSerialQueue {
	rv := objc.Send[DispatchSerialQueue](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DispatchSerialQueueClass) New() DispatchSerialQueue {
	rv := objc.Send[DispatchSerialQueue](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DispatchSerialQueue) Init() DispatchSerialQueue {
	rv := objc.Send[DispatchSerialQueue](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DispatchSerialQueue) Autorelease() DispatchSerialQueue {
	rv := objc.Send[DispatchSerialQueue](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDispatchSerialQueue creates a new DispatchSerialQueue instance.
func NewDispatchSerialQueue() DispatchSerialQueue {
	return getDispatchSerialQueueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DispatchSerialQueue */
// A custom dispatch queue that schedules tasks for serial execution on an arbitrary thread.
//
// You do not create objects of this type directly. You receive a queue of the appropriate type when you create a new object.


// A custom dispatch queue that schedules tasks for serial execution on an arbitrary thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchSerialQueue
type DispatchSerialQueue struct {
	objectivec.Object
}

// DispatchSerialQueueFrom constructs a [DispatchSerialQueue] from an unsafe.Pointer.
//
// A custom dispatch queue that schedules tasks for serial execution on an arbitrary thread.
func DispatchSerialQueueFrom(ptr unsafe.Pointer) DispatchSerialQueue {
	return DispatchSerialQueue{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DispatchSerialQueue *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DispatchSerialQueue */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DispatchSerialQueue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DispatchSerialQueue */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DispatchSerialQueue */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DispatchSerialQueue */



