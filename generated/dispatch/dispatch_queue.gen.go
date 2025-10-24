// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DispatchQueue */


/* debug [class_header]: Header for DispatchQueue */
// The class instance for the [DispatchQueue] class.
var (
	DispatchQueueClass     _DispatchQueueClass
	DispatchQueueClassOnce sync.Once
)

func getDispatchQueueClass() _DispatchQueueClass {
	DispatchQueueClassOnce.Do(func() {
		DispatchQueueClass = _DispatchQueueClass{objc.GetClass("DispatchQueue")}
	})
	return DispatchQueueClass
}

type _DispatchQueueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DispatchQueue */
// An interface definition for the [DispatchQueue] class.
type IDispatchQueue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DispatchQueue */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DispatchQueue */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DispatchQueue */
// Alloc allocates a new instance without initialization.
func (dc _DispatchQueueClass) Alloc() DispatchQueue {
	rv := objc.Send[DispatchQueue](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DispatchQueueClass) New() DispatchQueue {
	rv := objc.Send[DispatchQueue](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DispatchQueue) Init() DispatchQueue {
	rv := objc.Send[DispatchQueue](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DispatchQueue) Autorelease() DispatchQueue {
	rv := objc.Send[DispatchQueue](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDispatchQueue creates a new DispatchQueue instance.
func NewDispatchQueue() DispatchQueue {
	return getDispatchQueueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DispatchQueue */
// An object that manages the execution of tasks serially or concurrently on your app’s main thread or on a background thread.
//
// Dispatch queues are FIFO queues to which your application can submit tasks in the form of block objects. Dispatch queues execute tasks either serially or concurrently. Work submitted to dispatch queues executes on a pool of threads managed by the system. Except for the dispatch queue representing your app’s main thread, the system makes no guarantees about which thread it uses to execute a task. You schedule work items synchronously or asynchronously. When you schedule a work item synchronously, your code waits until that item finishes execution. When you schedule a work item asynchronously, your code continues executing while the work item runs elsewhere.


// An object that manages the execution of tasks serially or concurrently on your app’s main thread or on a background thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchQueue
type DispatchQueue struct {
	objectivec.Object
}

// DispatchQueueFrom constructs a [DispatchQueue] from an unsafe.Pointer.
//
// An object that manages the execution of tasks serially or concurrently on your app’s main thread or on a background thread.
func DispatchQueueFrom(ptr unsafe.Pointer) DispatchQueue {
	return DispatchQueue{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DispatchQueue *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DispatchQueue */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DispatchQueue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DispatchQueue */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DispatchQueue */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DispatchQueue */



