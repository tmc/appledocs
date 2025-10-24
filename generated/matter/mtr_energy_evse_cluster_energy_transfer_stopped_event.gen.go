// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEClusterEnergyTransferStoppedEvent */


/* debug [class_header]: Header for MTREnergyEVSEClusterEnergyTransferStoppedEvent */
// The class instance for the [MTREnergyEVSEClusterEnergyTransferStoppedEvent] class.
var (
	MTREnergyEVSEClusterEnergyTransferStoppedEventClass     _MTREnergyEVSEClusterEnergyTransferStoppedEventClass
	MTREnergyEVSEClusterEnergyTransferStoppedEventClassOnce sync.Once
)

func getMTREnergyEVSEClusterEnergyTransferStoppedEventClass() _MTREnergyEVSEClusterEnergyTransferStoppedEventClass {
	MTREnergyEVSEClusterEnergyTransferStoppedEventClassOnce.Do(func() {
		MTREnergyEVSEClusterEnergyTransferStoppedEventClass = _MTREnergyEVSEClusterEnergyTransferStoppedEventClass{objc.GetClass("MTREnergyEVSEClusterEnergyTransferStoppedEvent")}
	})
	return MTREnergyEVSEClusterEnergyTransferStoppedEventClass
}

type _MTREnergyEVSEClusterEnergyTransferStoppedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEClusterEnergyTransferStoppedEvent */
// An interface definition for the [MTREnergyEVSEClusterEnergyTransferStoppedEvent] class.
type IMTREnergyEVSEClusterEnergyTransferStoppedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEClusterEnergyTransferStoppedEvent */
	// properties:
	EnergyTransferred() objc.IObject /* cross-framework: NSNumber */
	SetEnergyTransferred(value objc.IObject /* cross-framework: NSNumber */)
	Reason() objc.IObject /* cross-framework: NSNumber */
	SetReason(value objc.IObject /* cross-framework: NSNumber */)
	SessionID() objc.IObject /* cross-framework: NSNumber */
	SetSessionID(value objc.IObject /* cross-framework: NSNumber */)
	State() objc.IObject /* cross-framework: NSNumber */
	SetState(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEClusterEnergyTransferStoppedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEClusterEnergyTransferStoppedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterEnergyTransferStoppedEventClass) Alloc() MTREnergyEVSEClusterEnergyTransferStoppedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEnergyTransferStoppedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEClusterEnergyTransferStoppedEventClass) New() MTREnergyEVSEClusterEnergyTransferStoppedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEnergyTransferStoppedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) Init() MTREnergyEVSEClusterEnergyTransferStoppedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEnergyTransferStoppedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) Autorelease() MTREnergyEVSEClusterEnergyTransferStoppedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEnergyTransferStoppedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterEnergyTransferStoppedEvent creates a new MTREnergyEVSEClusterEnergyTransferStoppedEvent instance.
func NewMTREnergyEVSEClusterEnergyTransferStoppedEvent() MTREnergyEVSEClusterEnergyTransferStoppedEvent {
	return getMTREnergyEVSEClusterEnergyTransferStoppedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEClusterEnergyTransferStoppedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent
type MTREnergyEVSEClusterEnergyTransferStoppedEvent struct {
	objectivec.Object
}

// MTREnergyEVSEClusterEnergyTransferStoppedEventFrom constructs a [MTREnergyEVSEClusterEnergyTransferStoppedEvent] from an unsafe.Pointer.
func MTREnergyEVSEClusterEnergyTransferStoppedEventFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterEnergyTransferStoppedEvent {
	return MTREnergyEVSEClusterEnergyTransferStoppedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEClusterEnergyTransferStoppedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEClusterEnergyTransferStoppedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEClusterEnergyTransferStoppedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEClusterEnergyTransferStoppedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEClusterEnergyTransferStoppedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/energyTransferred
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) EnergyTransferred() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("energyTransferred"))
	return rv
}/* debug [instance_properties/getter]: energyTransferred */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/energyTransferred
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) SetEnergyTransferred(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnergyTransferred:"), value)
}/* debug [instance_properties/setter]: energyTransferred */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterenergytransferstoppedevent/reason
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) Reason() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("reason"))
	return rv
}/* debug [instance_properties/getter]: reason */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterenergytransferstoppedevent/reason
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) SetReason(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReason:"), value)
}/* debug [instance_properties/setter]: reason */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterenergytransferstoppedevent/sessionid
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) SessionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sessionID"))
	return rv
}/* debug [instance_properties/getter]: sessionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterenergytransferstoppedevent/sessionid
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) SetSessionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionID:"), value)
}/* debug [instance_properties/setter]: sessionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterenergytransferstoppedevent/state
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) State() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterenergytransferstoppedevent/state
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) SetState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEClusterEnergyTransferStoppedEvent */



