// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEClusterFaultEvent */


/* debug [class_header]: Header for MTREnergyEVSEClusterFaultEvent */
// The class instance for the [MTREnergyEVSEClusterFaultEvent] class.
var (
	MTREnergyEVSEClusterFaultEventClass     _MTREnergyEVSEClusterFaultEventClass
	MTREnergyEVSEClusterFaultEventClassOnce sync.Once
)

func getMTREnergyEVSEClusterFaultEventClass() _MTREnergyEVSEClusterFaultEventClass {
	MTREnergyEVSEClusterFaultEventClassOnce.Do(func() {
		MTREnergyEVSEClusterFaultEventClass = _MTREnergyEVSEClusterFaultEventClass{objc.GetClass("MTREnergyEVSEClusterFaultEvent")}
	})
	return MTREnergyEVSEClusterFaultEventClass
}

type _MTREnergyEVSEClusterFaultEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEClusterFaultEvent */
// An interface definition for the [MTREnergyEVSEClusterFaultEvent] class.
type IMTREnergyEVSEClusterFaultEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEClusterFaultEvent */
	// properties:
	FaultStateCurrentState() objc.IObject /* cross-framework: NSNumber */
	SetFaultStateCurrentState(value objc.IObject /* cross-framework: NSNumber */)
	FaultStatePreviousState() objc.IObject /* cross-framework: NSNumber */
	SetFaultStatePreviousState(value objc.IObject /* cross-framework: NSNumber */)
	SessionID() objc.IObject /* cross-framework: NSNumber */
	SetSessionID(value objc.IObject /* cross-framework: NSNumber */)
	State() objc.IObject /* cross-framework: NSNumber */
	SetState(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEClusterFaultEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEClusterFaultEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterFaultEventClass) Alloc() MTREnergyEVSEClusterFaultEvent {
	rv := objc.Send[MTREnergyEVSEClusterFaultEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEClusterFaultEventClass) New() MTREnergyEVSEClusterFaultEvent {
	rv := objc.Send[MTREnergyEVSEClusterFaultEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterFaultEvent) Init() MTREnergyEVSEClusterFaultEvent {
	rv := objc.Send[MTREnergyEVSEClusterFaultEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterFaultEvent) Autorelease() MTREnergyEVSEClusterFaultEvent {
	rv := objc.Send[MTREnergyEVSEClusterFaultEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterFaultEvent creates a new MTREnergyEVSEClusterFaultEvent instance.
func NewMTREnergyEVSEClusterFaultEvent() MTREnergyEVSEClusterFaultEvent {
	return getMTREnergyEVSEClusterFaultEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEClusterFaultEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterFaultEvent
type MTREnergyEVSEClusterFaultEvent struct {
	objectivec.Object
}

// MTREnergyEVSEClusterFaultEventFrom constructs a [MTREnergyEVSEClusterFaultEvent] from an unsafe.Pointer.
func MTREnergyEVSEClusterFaultEventFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterFaultEvent {
	return MTREnergyEVSEClusterFaultEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEClusterFaultEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEClusterFaultEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEClusterFaultEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEClusterFaultEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEClusterFaultEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterFaultEvent/faultStateCurrentState
func (m_ MTREnergyEVSEClusterFaultEvent) FaultStateCurrentState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("faultStateCurrentState"))
	return rv
}/* debug [instance_properties/getter]: faultStateCurrentState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterFaultEvent/faultStateCurrentState
func (m_ MTREnergyEVSEClusterFaultEvent) SetFaultStateCurrentState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFaultStateCurrentState:"), value)
}/* debug [instance_properties/setter]: faultStateCurrentState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterfaultevent/faultstatepreviousstate
func (m_ MTREnergyEVSEClusterFaultEvent) FaultStatePreviousState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("faultStatePreviousState"))
	return rv
}/* debug [instance_properties/getter]: faultStatePreviousState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterfaultevent/faultstatepreviousstate
func (m_ MTREnergyEVSEClusterFaultEvent) SetFaultStatePreviousState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFaultStatePreviousState:"), value)
}/* debug [instance_properties/setter]: faultStatePreviousState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterfaultevent/sessionid
func (m_ MTREnergyEVSEClusterFaultEvent) SessionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sessionID"))
	return rv
}/* debug [instance_properties/getter]: sessionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterfaultevent/sessionid
func (m_ MTREnergyEVSEClusterFaultEvent) SetSessionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionID:"), value)
}/* debug [instance_properties/setter]: sessionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterfaultevent/state
func (m_ MTREnergyEVSEClusterFaultEvent) State() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrenergyevseclusterfaultevent/state
func (m_ MTREnergyEVSEClusterFaultEvent) SetState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}/* debug [instance_properties/setter]: state */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEClusterFaultEvent */



