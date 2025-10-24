// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DispatchGroup */


/* debug [class_header]: Header for DispatchGroup */
// The class instance for the [DispatchGroup] class.
var (
	DispatchGroupClass     _DispatchGroupClass
	DispatchGroupClassOnce sync.Once
)

func getDispatchGroupClass() _DispatchGroupClass {
	DispatchGroupClassOnce.Do(func() {
		DispatchGroupClass = _DispatchGroupClass{objc.GetClass("DispatchGroup")}
	})
	return DispatchGroupClass
}

type _DispatchGroupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DispatchGroup */
// An interface definition for the [DispatchGroup] class.
type IDispatchGroup interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DispatchGroup */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DispatchGroup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DispatchGroup */
// Alloc allocates a new instance without initialization.
func (dc _DispatchGroupClass) Alloc() DispatchGroup {
	rv := objc.Send[DispatchGroup](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DispatchGroupClass) New() DispatchGroup {
	rv := objc.Send[DispatchGroup](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DispatchGroup) Init() DispatchGroup {
	rv := objc.Send[DispatchGroup](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DispatchGroup) Autorelease() DispatchGroup {
	rv := objc.Send[DispatchGroup](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDispatchGroup creates a new DispatchGroup instance.
func NewDispatchGroup() DispatchGroup {
	return getDispatchGroupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DispatchGroup */
// A group of tasks that you monitor as a single unit.
//
// Groups allow you to aggregate a set of tasks and synchronize behaviors on the group. You attach multiple work items to a group and schedule them for asynchronous execution on the same queue or different queues. When all work items finish executing, the group executes its completion handler. You can also wait synchronously for all tasks in the group to finish executing.


// A group of tasks that you monitor as a single unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Dispatch/DispatchGroup
type DispatchGroup struct {
	objectivec.Object
}

// DispatchGroupFrom constructs a [DispatchGroup] from an unsafe.Pointer.
//
// A group of tasks that you monitor as a single unit.
func DispatchGroupFrom(ptr unsafe.Pointer) DispatchGroup {
	return DispatchGroup{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DispatchGroup *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DispatchGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DispatchGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DispatchGroup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DispatchGroup */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DispatchGroup */



