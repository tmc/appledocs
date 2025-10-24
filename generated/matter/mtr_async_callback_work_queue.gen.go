// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRAsyncCallbackWorkQueue */


/* debug [class_header]: Header for MTRAsyncCallbackWorkQueue */
// The class instance for the [MTRAsyncCallbackWorkQueue] class.
var (
	MTRAsyncCallbackWorkQueueClass     _MTRAsyncCallbackWorkQueueClass
	MTRAsyncCallbackWorkQueueClassOnce sync.Once
)

func getMTRAsyncCallbackWorkQueueClass() _MTRAsyncCallbackWorkQueueClass {
	MTRAsyncCallbackWorkQueueClassOnce.Do(func() {
		MTRAsyncCallbackWorkQueueClass = _MTRAsyncCallbackWorkQueueClass{objc.GetClass("MTRAsyncCallbackWorkQueue")}
	})
	return MTRAsyncCallbackWorkQueueClass
}

type _MTRAsyncCallbackWorkQueueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRAsyncCallbackWorkQueue */
// An interface definition for the [MTRAsyncCallbackWorkQueue] class.
type IMTRAsyncCallbackWorkQueue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRAsyncCallbackWorkQueue */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRAsyncCallbackWorkQueue */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRAsyncCallbackWorkQueue */
// Alloc allocates a new instance without initialization.
func (mc _MTRAsyncCallbackWorkQueueClass) Alloc() MTRAsyncCallbackWorkQueue {
	rv := objc.Send[MTRAsyncCallbackWorkQueue](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRAsyncCallbackWorkQueueClass) New() MTRAsyncCallbackWorkQueue {
	rv := objc.Send[MTRAsyncCallbackWorkQueue](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAsyncCallbackWorkQueue) Init() MTRAsyncCallbackWorkQueue {
	rv := objc.Send[MTRAsyncCallbackWorkQueue](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAsyncCallbackWorkQueue) Autorelease() MTRAsyncCallbackWorkQueue {
	rv := objc.Send[MTRAsyncCallbackWorkQueue](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAsyncCallbackWorkQueue creates a new MTRAsyncCallbackWorkQueue instance.
func NewMTRAsyncCallbackWorkQueue() MTRAsyncCallbackWorkQueue {
	return getMTRAsyncCallbackWorkQueueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRAsyncCallbackWorkQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAsyncCallbackWorkQueue
type MTRAsyncCallbackWorkQueue struct {
	objectivec.Object
}

// MTRAsyncCallbackWorkQueueFrom constructs a [MTRAsyncCallbackWorkQueue] from an unsafe.Pointer.
func MTRAsyncCallbackWorkQueueFrom(ptr unsafe.Pointer) MTRAsyncCallbackWorkQueue {
	return MTRAsyncCallbackWorkQueue{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRAsyncCallbackWorkQueue */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAsyncCallbackWorkQueue/init(context:queue:)
func NewMTRAsyncCallbackWorkQueueWithContextQueue(context objc.IObject, queue unsafe.Pointer) MTRAsyncCallbackWorkQueue {
	instance := getMTRAsyncCallbackWorkQueueClass().Alloc()
	rv := objc.Send[MTRAsyncCallbackWorkQueue](instance.ID, objc.Sel("initWithContext:queue:"), context, queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRAsyncCallbackWorkQueueWithContextQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRAsyncCallbackWorkQueue */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRAsyncCallbackWorkQueue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRAsyncCallbackWorkQueue */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRAsyncCallbackWorkQueue */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRAsyncCallbackWorkQueue */


