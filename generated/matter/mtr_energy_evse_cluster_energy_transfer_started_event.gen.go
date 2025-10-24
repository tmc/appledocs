// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEClusterEnergyTransferStartedEvent */


/* debug [class_header]: Header for MTREnergyEVSEClusterEnergyTransferStartedEvent */
// The class instance for the [MTREnergyEVSEClusterEnergyTransferStartedEvent] class.
var (
	MTREnergyEVSEClusterEnergyTransferStartedEventClass     _MTREnergyEVSEClusterEnergyTransferStartedEventClass
	MTREnergyEVSEClusterEnergyTransferStartedEventClassOnce sync.Once
)

func getMTREnergyEVSEClusterEnergyTransferStartedEventClass() _MTREnergyEVSEClusterEnergyTransferStartedEventClass {
	MTREnergyEVSEClusterEnergyTransferStartedEventClassOnce.Do(func() {
		MTREnergyEVSEClusterEnergyTransferStartedEventClass = _MTREnergyEVSEClusterEnergyTransferStartedEventClass{objc.GetClass("MTREnergyEVSEClusterEnergyTransferStartedEvent")}
	})
	return MTREnergyEVSEClusterEnergyTransferStartedEventClass
}

type _MTREnergyEVSEClusterEnergyTransferStartedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEClusterEnergyTransferStartedEvent */
// An interface definition for the [MTREnergyEVSEClusterEnergyTransferStartedEvent] class.
type IMTREnergyEVSEClusterEnergyTransferStartedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEClusterEnergyTransferStartedEvent */
	// properties:
	MaximumCurrent() objc.IObject /* cross-framework: NSNumber */
	SetMaximumCurrent(value objc.IObject /* cross-framework: NSNumber */)
	SessionID() objc.IObject /* cross-framework: NSNumber */
	SetSessionID(value objc.IObject /* cross-framework: NSNumber */)
	State() objc.IObject /* cross-framework: NSNumber */
	SetState(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEClusterEnergyTransferStartedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEClusterEnergyTransferStartedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterEnergyTransferStartedEventClass) Alloc() MTREnergyEVSEClusterEnergyTransferStartedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEnergyTransferStartedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEClusterEnergyTransferStartedEventClass) New() MTREnergyEVSEClusterEnergyTransferStartedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEnergyTransferStartedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterEnergyTransferStartedEvent) Init() MTREnergyEVSEClusterEnergyTransferStartedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEnergyTransferStartedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterEnergyTransferStartedEvent) Autorelease() MTREnergyEVSEClusterEnergyTransferStartedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEnergyTransferStartedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterEnergyTransferStartedEvent creates a new MTREnergyEVSEClusterEnergyTransferStartedEvent instance.
func NewMTREnergyEVSEClusterEnergyTransferStartedEvent() MTREnergyEVSEClusterEnergyTransferStartedEvent {
	return getMTREnergyEVSEClusterEnergyTransferStartedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEClusterEnergyTransferStartedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStartedEvent
type MTREnergyEVSEClusterEnergyTransferStartedEvent struct {
	objectivec.Object
}

// MTREnergyEVSEClusterEnergyTransferStartedEventFrom constructs a [MTREnergyEVSEClusterEnergyTransferStartedEvent] from an unsafe.Pointer.
func MTREnergyEVSEClusterEnergyTransferStartedEventFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterEnergyTransferStartedEvent {
	return MTREnergyEVSEClusterEnergyTransferStartedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEClusterEnergyTransferStartedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEClusterEnergyTransferStartedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEClusterEnergyTransferStartedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEClusterEnergyTransferStartedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEClusterEnergyTransferStartedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStartedEvent/maximumCurrent
func (m_ MTREnergyEVSEClusterEnergyTransferStartedEvent) MaximumCurrent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maximumCurrent"))
	return rv
}/* debug [instance_properties/getter]: maximumCurrent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStartedEvent/maximumCurrent
func (m_ MTREnergyEVSEClusterEnergyTransferStartedEvent) SetMaximumCurrent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumCurrent:"), value)
}/* debug [instance_properties/setter]: maximumCurrent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStartedEvent/sessionID
func (m_ MTREnergyEVSEClusterEnergyTransferStartedEvent) SessionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sessionID"))
	return rv
}/* debug [instance_properties/getter]: sessionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStartedEvent/sessionID
func (m_ MTREnergyEVSEClusterEnergyTransferStartedEvent) SetSessionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionID:"), value)
}/* debug [instance_properties/setter]: sessionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterenergytransferstartedevent/state
func (m_ MTREnergyEVSEClusterEnergyTransferStartedEvent) State() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterenergytransferstartedevent/state
func (m_ MTREnergyEVSEClusterEnergyTransferStartedEvent) SetState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEClusterEnergyTransferStartedEvent */



