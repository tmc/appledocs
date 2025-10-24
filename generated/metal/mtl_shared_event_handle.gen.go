// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLSharedEventHandle */


/* debug [class_header]: Header for MTLSharedEventHandle */
// The class instance for the [SharedEventHandle] class.
var (
	SharedEventHandleClass     _SharedEventHandleClass
	SharedEventHandleClassOnce sync.Once
)

func getSharedEventHandleClass() _SharedEventHandleClass {
	SharedEventHandleClassOnce.Do(func() {
		SharedEventHandleClass = _SharedEventHandleClass{objc.GetClass("MTLSharedEventHandle")}
	})
	return SharedEventHandleClass
}

type _SharedEventHandleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SharedEventHandle */
// An interface definition for the [SharedEventHandle] class.
type ISharedEventHandle interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SharedEventHandle */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SharedEventHandle */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SharedEventHandle */
// Alloc allocates a new instance without initialization.
func (sc _SharedEventHandleClass) Alloc() SharedEventHandle {
	rv := objc.Send[SharedEventHandle](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SharedEventHandleClass) New() SharedEventHandle {
	rv := objc.Send[SharedEventHandle](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SharedEventHandle) Init() SharedEventHandle {
	rv := objc.Send[SharedEventHandle](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SharedEventHandle) Autorelease() SharedEventHandle {
	rv := objc.Send[SharedEventHandle](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSharedEventHandle creates a new SharedEventHandle instance.
func NewSharedEventHandle() SharedEventHandle {
	return getSharedEventHandleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SharedEventHandle */
// An instance you use to recreate a shareable event.
//
// To create a instance, call the method on an instance. Use an XPC conection to pass a instance to another process. To recreate the event, call the on an instance.


// An instance you use to recreate a shareable event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSharedEventHandle
type SharedEventHandle struct {
	objectivec.Object
}

// SharedEventHandleFrom constructs a [SharedEventHandle] from an unsafe.Pointer.
//
// An instance you use to recreate a shareable event.
func SharedEventHandleFrom(ptr unsafe.Pointer) SharedEventHandle {
	return SharedEventHandle{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SharedEventHandle *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SharedEventHandle */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SharedEventHandle */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SharedEventHandle */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SharedEventHandle */

// A string that identifies the shareable event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSharedEventHandle/label
func (s_ SharedEventHandle) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLSharedEventHandle */



