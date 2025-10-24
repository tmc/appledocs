// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEClusterEVNotDetectedEvent */


/* debug [class_header]: Header for MTREnergyEVSEClusterEVNotDetectedEvent */
// The class instance for the [MTREnergyEVSEClusterEVNotDetectedEvent] class.
var (
	MTREnergyEVSEClusterEVNotDetectedEventClass     _MTREnergyEVSEClusterEVNotDetectedEventClass
	MTREnergyEVSEClusterEVNotDetectedEventClassOnce sync.Once
)

func getMTREnergyEVSEClusterEVNotDetectedEventClass() _MTREnergyEVSEClusterEVNotDetectedEventClass {
	MTREnergyEVSEClusterEVNotDetectedEventClassOnce.Do(func() {
		MTREnergyEVSEClusterEVNotDetectedEventClass = _MTREnergyEVSEClusterEVNotDetectedEventClass{objc.GetClass("MTREnergyEVSEClusterEVNotDetectedEvent")}
	})
	return MTREnergyEVSEClusterEVNotDetectedEventClass
}

type _MTREnergyEVSEClusterEVNotDetectedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEClusterEVNotDetectedEvent */
// An interface definition for the [MTREnergyEVSEClusterEVNotDetectedEvent] class.
type IMTREnergyEVSEClusterEVNotDetectedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEClusterEVNotDetectedEvent */
	// properties:
	SessionDuration() objc.IObject /* cross-framework: NSNumber */
	SetSessionDuration(value objc.IObject /* cross-framework: NSNumber */)
	SessionEnergyCharged() objc.IObject /* cross-framework: NSNumber */
	SetSessionEnergyCharged(value objc.IObject /* cross-framework: NSNumber */)
	SessionID() objc.IObject /* cross-framework: NSNumber */
	SetSessionID(value objc.IObject /* cross-framework: NSNumber */)
	State() objc.IObject /* cross-framework: NSNumber */
	SetState(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEClusterEVNotDetectedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEClusterEVNotDetectedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterEVNotDetectedEventClass) Alloc() MTREnergyEVSEClusterEVNotDetectedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEVNotDetectedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEClusterEVNotDetectedEventClass) New() MTREnergyEVSEClusterEVNotDetectedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEVNotDetectedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) Init() MTREnergyEVSEClusterEVNotDetectedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEVNotDetectedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) Autorelease() MTREnergyEVSEClusterEVNotDetectedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEVNotDetectedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterEVNotDetectedEvent creates a new MTREnergyEVSEClusterEVNotDetectedEvent instance.
func NewMTREnergyEVSEClusterEVNotDetectedEvent() MTREnergyEVSEClusterEVNotDetectedEvent {
	return getMTREnergyEVSEClusterEVNotDetectedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEClusterEVNotDetectedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent
type MTREnergyEVSEClusterEVNotDetectedEvent struct {
	objectivec.Object
}

// MTREnergyEVSEClusterEVNotDetectedEventFrom constructs a [MTREnergyEVSEClusterEVNotDetectedEvent] from an unsafe.Pointer.
func MTREnergyEVSEClusterEVNotDetectedEventFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterEVNotDetectedEvent {
	return MTREnergyEVSEClusterEVNotDetectedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEClusterEVNotDetectedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEClusterEVNotDetectedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEClusterEVNotDetectedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEClusterEVNotDetectedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEClusterEVNotDetectedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/sessionDuration
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SessionDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sessionDuration"))
	return rv
}/* debug [instance_properties/getter]: sessionDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/sessionDuration
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SetSessionDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionDuration:"), value)
}/* debug [instance_properties/setter]: sessionDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterevnotdetectedevent/sessionenergycharged
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SessionEnergyCharged() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sessionEnergyCharged"))
	return rv
}/* debug [instance_properties/getter]: sessionEnergyCharged */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterevnotdetectedevent/sessionenergycharged
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SetSessionEnergyCharged(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionEnergyCharged:"), value)
}/* debug [instance_properties/setter]: sessionEnergyCharged */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterevnotdetectedevent/sessionid
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SessionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sessionID"))
	return rv
}/* debug [instance_properties/getter]: sessionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterevnotdetectedevent/sessionid
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SetSessionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionID:"), value)
}/* debug [instance_properties/setter]: sessionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterevnotdetectedevent/state
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) State() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterevnotdetectedevent/state
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SetState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEClusterEVNotDetectedEvent */



