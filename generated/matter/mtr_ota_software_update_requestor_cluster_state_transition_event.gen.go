// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent */


/* debug [class_header]: Header for MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent */
// The class instance for the [MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent] class.
var (
	MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass     _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass
	MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClassOnce sync.Once
)

func getMTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass() _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass {
	MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClassOnce.Do(func() {
		MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass = _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass{objc.GetClass("MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent")}
	})
	return MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass
}

type _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent */
// An interface definition for the [MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent] class.
type IMTROtaSoftwareUpdateRequestorClusterStateTransitionEvent interface {
	IMTROTASoftwareUpdateRequestorClusterStateTransitionEvent
	
/* debug [class_interface_properties]: Properties for MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent */
	// properties:
	NewState() objc.IObject /* cross-framework: NSNumber */
	SetNewState(value objc.IObject /* cross-framework: NSNumber */)
	PreviousState() objc.IObject /* cross-framework: NSNumber */
	SetPreviousState(value objc.IObject /* cross-framework: NSNumber */)
	Reason() objc.IObject /* cross-framework: NSNumber */
	SetReason(value objc.IObject /* cross-framework: NSNumber */)
	TargetSoftwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetTargetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass) Alloc() MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass) New() MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) Init() MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) Autorelease() MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateRequestorClusterStateTransitionEvent creates a new MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent instance.
func NewMTROtaSoftwareUpdateRequestorClusterStateTransitionEvent() MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	return getMTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent-1xzd5
type MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent struct {
	MTROTASoftwareUpdateRequestorClusterStateTransitionEvent
}

// MTROtaSoftwareUpdateRequestorClusterStateTransitionEventFrom constructs a [MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent] from an unsafe.Pointer.
func MTROtaSoftwareUpdateRequestorClusterStateTransitionEventFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	return MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent{
		MTROTASoftwareUpdateRequestorClusterStateTransitionEvent: MTROTASoftwareUpdateRequestorClusterStateTransitionEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent-1xzd5/newState
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) NewState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newState"))
	return rv
}/* debug [instance_properties/getter]: newState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent-1xzd5/newState
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) SetNewState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewState:"), value)
}/* debug [instance_properties/setter]: newState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent-1xzd5/previousState
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) PreviousState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("previousState"))
	return rv
}/* debug [instance_properties/getter]: previousState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent-1xzd5/previousState
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) SetPreviousState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreviousState:"), value)
}/* debug [instance_properties/setter]: previousState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent-1xzd5/reason
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) Reason() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("reason"))
	return rv
}/* debug [instance_properties/getter]: reason */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent-1xzd5/reason
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) SetReason(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReason:"), value)
}/* debug [instance_properties/setter]: reason */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent-1xzd5/targetSoftwareVersion
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) TargetSoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("targetSoftwareVersion"))
	return rv
}/* debug [instance_properties/getter]: targetSoftwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent-1xzd5/targetSoftwareVersion
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) SetTargetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetSoftwareVersion:"), value)
}/* debug [instance_properties/setter]: targetSoftwareVersion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent */



