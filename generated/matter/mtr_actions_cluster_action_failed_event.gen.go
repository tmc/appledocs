// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRActionsClusterActionFailedEvent */


/* debug [class_header]: Header for MTRActionsClusterActionFailedEvent */
// The class instance for the [MTRActionsClusterActionFailedEvent] class.
var (
	MTRActionsClusterActionFailedEventClass     _MTRActionsClusterActionFailedEventClass
	MTRActionsClusterActionFailedEventClassOnce sync.Once
)

func getMTRActionsClusterActionFailedEventClass() _MTRActionsClusterActionFailedEventClass {
	MTRActionsClusterActionFailedEventClassOnce.Do(func() {
		MTRActionsClusterActionFailedEventClass = _MTRActionsClusterActionFailedEventClass{objc.GetClass("MTRActionsClusterActionFailedEvent")}
	})
	return MTRActionsClusterActionFailedEventClass
}

type _MTRActionsClusterActionFailedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRActionsClusterActionFailedEvent */
// An interface definition for the [MTRActionsClusterActionFailedEvent] class.
type IMTRActionsClusterActionFailedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRActionsClusterActionFailedEvent */
	// properties:
	ActionID() objc.IObject /* cross-framework: NSNumber */
	SetActionID(value objc.IObject /* cross-framework: NSNumber */)
	Error() objc.IObject /* cross-framework: NSNumber */
	SetError(value objc.IObject /* cross-framework: NSNumber */)
	InvokeID() objc.IObject /* cross-framework: NSNumber */
	SetInvokeID(value objc.IObject /* cross-framework: NSNumber */)
	NewState() objc.IObject /* cross-framework: NSNumber */
	SetNewState(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRActionsClusterActionFailedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRActionsClusterActionFailedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterActionFailedEventClass) Alloc() MTRActionsClusterActionFailedEvent {
	rv := objc.Send[MTRActionsClusterActionFailedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRActionsClusterActionFailedEventClass) New() MTRActionsClusterActionFailedEvent {
	rv := objc.Send[MTRActionsClusterActionFailedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterActionFailedEvent) Init() MTRActionsClusterActionFailedEvent {
	rv := objc.Send[MTRActionsClusterActionFailedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterActionFailedEvent) Autorelease() MTRActionsClusterActionFailedEvent {
	rv := objc.Send[MTRActionsClusterActionFailedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterActionFailedEvent creates a new MTRActionsClusterActionFailedEvent instance.
func NewMTRActionsClusterActionFailedEvent() MTRActionsClusterActionFailedEvent {
	return getMTRActionsClusterActionFailedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRActionsClusterActionFailedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionFailedEvent
type MTRActionsClusterActionFailedEvent struct {
	objectivec.Object
}

// MTRActionsClusterActionFailedEventFrom constructs a [MTRActionsClusterActionFailedEvent] from an unsafe.Pointer.
func MTRActionsClusterActionFailedEventFrom(ptr unsafe.Pointer) MTRActionsClusterActionFailedEvent {
	return MTRActionsClusterActionFailedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRActionsClusterActionFailedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRActionsClusterActionFailedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRActionsClusterActionFailedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRActionsClusterActionFailedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRActionsClusterActionFailedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionFailedEvent/actionID
func (m_ MTRActionsClusterActionFailedEvent) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}/* debug [instance_properties/getter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionFailedEvent/actionID
func (m_ MTRActionsClusterActionFailedEvent) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}/* debug [instance_properties/setter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionFailedEvent/error
func (m_ MTRActionsClusterActionFailedEvent) Error() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionFailedEvent/error
func (m_ MTRActionsClusterActionFailedEvent) SetError(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setError:"), value)
}/* debug [instance_properties/setter]: error */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionFailedEvent/invokeID
func (m_ MTRActionsClusterActionFailedEvent) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}/* debug [instance_properties/getter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionFailedEvent/invokeID
func (m_ MTRActionsClusterActionFailedEvent) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}/* debug [instance_properties/setter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionFailedEvent/newState
func (m_ MTRActionsClusterActionFailedEvent) NewState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newState"))
	return rv
}/* debug [instance_properties/getter]: newState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterActionFailedEvent/newState
func (m_ MTRActionsClusterActionFailedEvent) SetNewState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewState:"), value)
}/* debug [instance_properties/setter]: newState */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRActionsClusterActionFailedEvent */



