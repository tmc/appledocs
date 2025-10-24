// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROTASoftwareUpdateRequestorClusterStateTransitionEvent */


/* debug [class_header]: Header for MTROTASoftwareUpdateRequestorClusterStateTransitionEvent */
// The class instance for the [MTROTASoftwareUpdateRequestorClusterStateTransitionEvent] class.
var (
	MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass     _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass
	MTROTASoftwareUpdateRequestorClusterStateTransitionEventClassOnce sync.Once
)

func getMTROTASoftwareUpdateRequestorClusterStateTransitionEventClass() _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass {
	MTROTASoftwareUpdateRequestorClusterStateTransitionEventClassOnce.Do(func() {
		MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass = _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass{objc.GetClass("MTROTASoftwareUpdateRequestorClusterStateTransitionEvent")}
	})
	return MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass
}

type _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROTASoftwareUpdateRequestorClusterStateTransitionEvent */
// An interface definition for the [MTROTASoftwareUpdateRequestorClusterStateTransitionEvent] class.
type IMTROTASoftwareUpdateRequestorClusterStateTransitionEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROTASoftwareUpdateRequestorClusterStateTransitionEvent */
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

	
/* debug [class_interface_methods]: Methods for MTROTASoftwareUpdateRequestorClusterStateTransitionEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROTASoftwareUpdateRequestorClusterStateTransitionEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass) Alloc() MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterStateTransitionEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass) New() MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterStateTransitionEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) Init() MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterStateTransitionEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) Autorelease() MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterStateTransitionEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateRequestorClusterStateTransitionEvent creates a new MTROTASoftwareUpdateRequestorClusterStateTransitionEvent instance.
func NewMTROTASoftwareUpdateRequestorClusterStateTransitionEvent() MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	return getMTROTASoftwareUpdateRequestorClusterStateTransitionEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROTASoftwareUpdateRequestorClusterStateTransitionEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterStateTransitionEvent-3xhxb
type MTROTASoftwareUpdateRequestorClusterStateTransitionEvent struct {
	objectivec.Object
}

// MTROTASoftwareUpdateRequestorClusterStateTransitionEventFrom constructs a [MTROTASoftwareUpdateRequestorClusterStateTransitionEvent] from an unsafe.Pointer.
func MTROTASoftwareUpdateRequestorClusterStateTransitionEventFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	return MTROTASoftwareUpdateRequestorClusterStateTransitionEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROTASoftwareUpdateRequestorClusterStateTransitionEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROTASoftwareUpdateRequestorClusterStateTransitionEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROTASoftwareUpdateRequestorClusterStateTransitionEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROTASoftwareUpdateRequestorClusterStateTransitionEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROTASoftwareUpdateRequestorClusterStateTransitionEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterStateTransitionEvent-3xhxb/newState
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) NewState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newState"))
	return rv
}/* debug [instance_properties/getter]: newState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterStateTransitionEvent-3xhxb/newState
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) SetNewState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewState:"), value)
}/* debug [instance_properties/setter]: newState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterStateTransitionEvent-3xhxb/previousState
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) PreviousState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("previousState"))
	return rv
}/* debug [instance_properties/getter]: previousState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterStateTransitionEvent-3xhxb/previousState
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) SetPreviousState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreviousState:"), value)
}/* debug [instance_properties/setter]: previousState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterStateTransitionEvent-3xhxb/reason
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) Reason() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("reason"))
	return rv
}/* debug [instance_properties/getter]: reason */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterStateTransitionEvent-3xhxb/reason
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) SetReason(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReason:"), value)
}/* debug [instance_properties/setter]: reason */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterStateTransitionEvent-3xhxb/targetSoftwareVersion
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) TargetSoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("targetSoftwareVersion"))
	return rv
}/* debug [instance_properties/getter]: targetSoftwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterStateTransitionEvent-3xhxb/targetSoftwareVersion
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) SetTargetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetSoftwareVersion:"), value)
}/* debug [instance_properties/setter]: targetSoftwareVersion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROTASoftwareUpdateRequestorClusterStateTransitionEvent */



