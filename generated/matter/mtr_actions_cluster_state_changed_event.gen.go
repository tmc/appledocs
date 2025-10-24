// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRActionsClusterStateChangedEvent */


/* debug [class_header]: Header for MTRActionsClusterStateChangedEvent */
// The class instance for the [MTRActionsClusterStateChangedEvent] class.
var (
	MTRActionsClusterStateChangedEventClass     _MTRActionsClusterStateChangedEventClass
	MTRActionsClusterStateChangedEventClassOnce sync.Once
)

func getMTRActionsClusterStateChangedEventClass() _MTRActionsClusterStateChangedEventClass {
	MTRActionsClusterStateChangedEventClassOnce.Do(func() {
		MTRActionsClusterStateChangedEventClass = _MTRActionsClusterStateChangedEventClass{objc.GetClass("MTRActionsClusterStateChangedEvent")}
	})
	return MTRActionsClusterStateChangedEventClass
}

type _MTRActionsClusterStateChangedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRActionsClusterStateChangedEvent */
// An interface definition for the [MTRActionsClusterStateChangedEvent] class.
type IMTRActionsClusterStateChangedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRActionsClusterStateChangedEvent */
	// properties:
	ActionID() objc.IObject /* cross-framework: NSNumber */
	SetActionID(value objc.IObject /* cross-framework: NSNumber */)
	InvokeID() objc.IObject /* cross-framework: NSNumber */
	SetInvokeID(value objc.IObject /* cross-framework: NSNumber */)
	NewState() objc.IObject /* cross-framework: NSNumber */
	SetNewState(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRActionsClusterStateChangedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRActionsClusterStateChangedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRActionsClusterStateChangedEventClass) Alloc() MTRActionsClusterStateChangedEvent {
	rv := objc.Send[MTRActionsClusterStateChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRActionsClusterStateChangedEventClass) New() MTRActionsClusterStateChangedEvent {
	rv := objc.Send[MTRActionsClusterStateChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRActionsClusterStateChangedEvent) Init() MTRActionsClusterStateChangedEvent {
	rv := objc.Send[MTRActionsClusterStateChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRActionsClusterStateChangedEvent) Autorelease() MTRActionsClusterStateChangedEvent {
	rv := objc.Send[MTRActionsClusterStateChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRActionsClusterStateChangedEvent creates a new MTRActionsClusterStateChangedEvent instance.
func NewMTRActionsClusterStateChangedEvent() MTRActionsClusterStateChangedEvent {
	return getMTRActionsClusterStateChangedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRActionsClusterStateChangedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStateChangedEvent
type MTRActionsClusterStateChangedEvent struct {
	objectivec.Object
}

// MTRActionsClusterStateChangedEventFrom constructs a [MTRActionsClusterStateChangedEvent] from an unsafe.Pointer.
func MTRActionsClusterStateChangedEventFrom(ptr unsafe.Pointer) MTRActionsClusterStateChangedEvent {
	return MTRActionsClusterStateChangedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRActionsClusterStateChangedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRActionsClusterStateChangedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRActionsClusterStateChangedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRActionsClusterStateChangedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRActionsClusterStateChangedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStateChangedEvent/actionID
func (m_ MTRActionsClusterStateChangedEvent) ActionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("actionID"))
	return rv
}/* debug [instance_properties/getter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStateChangedEvent/actionID
func (m_ MTRActionsClusterStateChangedEvent) SetActionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActionID:"), value)
}/* debug [instance_properties/setter]: actionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStateChangedEvent/invokeID
func (m_ MTRActionsClusterStateChangedEvent) InvokeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("invokeID"))
	return rv
}/* debug [instance_properties/getter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStateChangedEvent/invokeID
func (m_ MTRActionsClusterStateChangedEvent) SetInvokeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInvokeID:"), value)
}/* debug [instance_properties/setter]: invokeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStateChangedEvent/newState
func (m_ MTRActionsClusterStateChangedEvent) NewState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newState"))
	return rv
}/* debug [instance_properties/getter]: newState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRActionsClusterStateChangedEvent/newState
func (m_ MTRActionsClusterStateChangedEvent) SetNewState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewState:"), value)
}/* debug [instance_properties/setter]: newState */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRActionsClusterStateChangedEvent */



