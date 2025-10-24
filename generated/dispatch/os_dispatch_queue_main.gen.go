// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class OS_dispatch_queue_main */


/* debug [class_header]: Header for OS_dispatch_queue_main */
// The class instance for the [OS_dispatch_queue_main] class.
var (
	OS_dispatch_queue_mainClass     _OS_dispatch_queue_mainClass
	OS_dispatch_queue_mainClassOnce sync.Once
)

func getOS_dispatch_queue_mainClass() _OS_dispatch_queue_mainClass {
	OS_dispatch_queue_mainClassOnce.Do(func() {
		OS_dispatch_queue_mainClass = _OS_dispatch_queue_mainClass{objc.GetClass("OS_dispatch_queue_main")}
	})
	return OS_dispatch_queue_mainClass
}

type _OS_dispatch_queue_mainClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OS_dispatch_queue_main */
// An interface definition for the [OS_dispatch_queue_main] class.
type IOS_dispatch_queue_main interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OS_dispatch_queue_main */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OS_dispatch_queue_main */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OS_dispatch_queue_main */
// Alloc allocates a new instance without initialization.
func (oc _OS_dispatch_queue_mainClass) Alloc() OS_dispatch_queue_main {
	rv := objc.Send[OS_dispatch_queue_main](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OS_dispatch_queue_mainClass) New() OS_dispatch_queue_main {
	rv := objc.Send[OS_dispatch_queue_main](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OS_dispatch_queue_main) Init() OS_dispatch_queue_main {
	rv := objc.Send[OS_dispatch_queue_main](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OS_dispatch_queue_main) Autorelease() OS_dispatch_queue_main {
	rv := objc.Send[OS_dispatch_queue_main](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOS_dispatch_queue_main creates a new OS_dispatch_queue_main instance.
func NewOS_dispatch_queue_main() OS_dispatch_queue_main {
	return getOS_dispatch_queue_mainClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OS_dispatch_queue_main */
// A system-provided dispatch queue that schedules tasks for serial execution on the app’s main thread.
//
// You do not create objects of this type directly. You receive a queue of the appropriate type when you create a new object.


// A system-provided dispatch queue that schedules tasks for serial execution on the app’s main thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/OS_dispatch_queue_main-swift.class
type OS_dispatch_queue_main struct {
	objectivec.Object
}

// OS_dispatch_queue_mainFrom constructs a [OS_dispatch_queue_main] from an unsafe.Pointer.
//
// A system-provided dispatch queue that schedules tasks for serial execution on the app’s main thread.
func OS_dispatch_queue_mainFrom(ptr unsafe.Pointer) OS_dispatch_queue_main {
	return OS_dispatch_queue_main{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OS_dispatch_queue_main *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OS_dispatch_queue_main */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OS_dispatch_queue_main */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OS_dispatch_queue_main */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OS_dispatch_queue_main */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class OS_dispatch_queue_main */





