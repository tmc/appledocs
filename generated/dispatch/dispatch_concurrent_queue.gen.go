// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DispatchConcurrentQueue */


/* debug [class_header]: Header for DispatchConcurrentQueue */
// The class instance for the [DispatchConcurrentQueue] class.
var (
	DispatchConcurrentQueueClass     _DispatchConcurrentQueueClass
	DispatchConcurrentQueueClassOnce sync.Once
)

func getDispatchConcurrentQueueClass() _DispatchConcurrentQueueClass {
	DispatchConcurrentQueueClassOnce.Do(func() {
		DispatchConcurrentQueueClass = _DispatchConcurrentQueueClass{objc.GetClass("DispatchConcurrentQueue")}
	})
	return DispatchConcurrentQueueClass
}

type _DispatchConcurrentQueueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DispatchConcurrentQueue */
// An interface definition for the [DispatchConcurrentQueue] class.
type IDispatchConcurrentQueue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DispatchConcurrentQueue */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DispatchConcurrentQueue */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DispatchConcurrentQueue */
// Alloc allocates a new instance without initialization.
func (dc _DispatchConcurrentQueueClass) Alloc() DispatchConcurrentQueue {
	rv := objc.Send[DispatchConcurrentQueue](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DispatchConcurrentQueueClass) New() DispatchConcurrentQueue {
	rv := objc.Send[DispatchConcurrentQueue](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DispatchConcurrentQueue) Init() DispatchConcurrentQueue {
	rv := objc.Send[DispatchConcurrentQueue](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DispatchConcurrentQueue) Autorelease() DispatchConcurrentQueue {
	rv := objc.Send[DispatchConcurrentQueue](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDispatchConcurrentQueue creates a new DispatchConcurrentQueue instance.
func NewDispatchConcurrentQueue() DispatchConcurrentQueue {
	return getDispatchConcurrentQueueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DispatchConcurrentQueue */
// A custom dispatch queue that schedules tasks for concurrent execution.
//
// You do not create objects of this type directly. You receive a queue of the appropriate type when you create a new object.


// A custom dispatch queue that schedules tasks for concurrent execution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchConcurrentQueue
type DispatchConcurrentQueue struct {
	objectivec.Object
}

// DispatchConcurrentQueueFrom constructs a [DispatchConcurrentQueue] from an unsafe.Pointer.
//
// A custom dispatch queue that schedules tasks for concurrent execution.
func DispatchConcurrentQueueFrom(ptr unsafe.Pointer) DispatchConcurrentQueue {
	return DispatchConcurrentQueue{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DispatchConcurrentQueue *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DispatchConcurrentQueue */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DispatchConcurrentQueue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DispatchConcurrentQueue */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DispatchConcurrentQueue */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DispatchConcurrentQueue */



