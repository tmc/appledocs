// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRAsyncCallbackQueueWorkItem */


/* debug [class_header]: Header for MTRAsyncCallbackQueueWorkItem */
// The class instance for the [MTRAsyncCallbackQueueWorkItem] class.
var (
	MTRAsyncCallbackQueueWorkItemClass     _MTRAsyncCallbackQueueWorkItemClass
	MTRAsyncCallbackQueueWorkItemClassOnce sync.Once
)

func getMTRAsyncCallbackQueueWorkItemClass() _MTRAsyncCallbackQueueWorkItemClass {
	MTRAsyncCallbackQueueWorkItemClassOnce.Do(func() {
		MTRAsyncCallbackQueueWorkItemClass = _MTRAsyncCallbackQueueWorkItemClass{objc.GetClass("MTRAsyncCallbackQueueWorkItem")}
	})
	return MTRAsyncCallbackQueueWorkItemClass
}

type _MTRAsyncCallbackQueueWorkItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRAsyncCallbackQueueWorkItem */
// An interface definition for the [MTRAsyncCallbackQueueWorkItem] class.
type IMTRAsyncCallbackQueueWorkItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRAsyncCallbackQueueWorkItem */
	// properties:
	CancelHandler() unsafe.Pointer
	SetCancelHandler(value unsafe.Pointer)
	ReadyHandler() unsafe.Pointer
	SetReadyHandler(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRAsyncCallbackQueueWorkItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRAsyncCallbackQueueWorkItem */
// Alloc allocates a new instance without initialization.
func (mc _MTRAsyncCallbackQueueWorkItemClass) Alloc() MTRAsyncCallbackQueueWorkItem {
	rv := objc.Send[MTRAsyncCallbackQueueWorkItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRAsyncCallbackQueueWorkItemClass) New() MTRAsyncCallbackQueueWorkItem {
	rv := objc.Send[MTRAsyncCallbackQueueWorkItem](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAsyncCallbackQueueWorkItem) Init() MTRAsyncCallbackQueueWorkItem {
	rv := objc.Send[MTRAsyncCallbackQueueWorkItem](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAsyncCallbackQueueWorkItem) Autorelease() MTRAsyncCallbackQueueWorkItem {
	rv := objc.Send[MTRAsyncCallbackQueueWorkItem](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAsyncCallbackQueueWorkItem creates a new MTRAsyncCallbackQueueWorkItem instance.
func NewMTRAsyncCallbackQueueWorkItem() MTRAsyncCallbackQueueWorkItem {
	return getMTRAsyncCallbackQueueWorkItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRAsyncCallbackQueueWorkItem */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAsyncCallbackQueueWorkItem
type MTRAsyncCallbackQueueWorkItem struct {
	objectivec.Object
}

// MTRAsyncCallbackQueueWorkItemFrom constructs a [MTRAsyncCallbackQueueWorkItem] from an unsafe.Pointer.
func MTRAsyncCallbackQueueWorkItemFrom(ptr unsafe.Pointer) MTRAsyncCallbackQueueWorkItem {
	return MTRAsyncCallbackQueueWorkItem{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRAsyncCallbackQueueWorkItem */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAsyncCallbackQueueWorkItem/init(queue:)
func NewMTRAsyncCallbackQueueWorkItemWithQueue(queue unsafe.Pointer) MTRAsyncCallbackQueueWorkItem {
	instance := getMTRAsyncCallbackQueueWorkItemClass().Alloc()
	rv := objc.Send[MTRAsyncCallbackQueueWorkItem](instance.ID, objc.Sel("initWithQueue:"), queue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRAsyncCallbackQueueWorkItemWithQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRAsyncCallbackQueueWorkItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRAsyncCallbackQueueWorkItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRAsyncCallbackQueueWorkItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRAsyncCallbackQueueWorkItem */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAsyncCallbackQueueWorkItem/cancelHandler
func (m_ MTRAsyncCallbackQueueWorkItem) CancelHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cancelHandler"))
	return rv
}/* debug [instance_properties/getter]: cancelHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAsyncCallbackQueueWorkItem/cancelHandler
func (m_ MTRAsyncCallbackQueueWorkItem) SetCancelHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCancelHandler:"), value)
}/* debug [instance_properties/setter]: cancelHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAsyncCallbackQueueWorkItem/readyHandler
func (m_ MTRAsyncCallbackQueueWorkItem) ReadyHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("readyHandler"))
	return rv
}/* debug [instance_properties/getter]: readyHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAsyncCallbackQueueWorkItem/readyHandler
func (m_ MTRAsyncCallbackQueueWorkItem) SetReadyHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReadyHandler:"), value)
}/* debug [instance_properties/setter]: readyHandler */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRAsyncCallbackQueueWorkItem */


