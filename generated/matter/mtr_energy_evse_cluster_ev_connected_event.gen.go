// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTREnergyEVSEClusterEVConnectedEvent */


/* debug [class_header]: Header for MTREnergyEVSEClusterEVConnectedEvent */
// The class instance for the [MTREnergyEVSEClusterEVConnectedEvent] class.
var (
	MTREnergyEVSEClusterEVConnectedEventClass     _MTREnergyEVSEClusterEVConnectedEventClass
	MTREnergyEVSEClusterEVConnectedEventClassOnce sync.Once
)

func getMTREnergyEVSEClusterEVConnectedEventClass() _MTREnergyEVSEClusterEVConnectedEventClass {
	MTREnergyEVSEClusterEVConnectedEventClassOnce.Do(func() {
		MTREnergyEVSEClusterEVConnectedEventClass = _MTREnergyEVSEClusterEVConnectedEventClass{objc.GetClass("MTREnergyEVSEClusterEVConnectedEvent")}
	})
	return MTREnergyEVSEClusterEVConnectedEventClass
}

type _MTREnergyEVSEClusterEVConnectedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTREnergyEVSEClusterEVConnectedEvent */
// An interface definition for the [MTREnergyEVSEClusterEVConnectedEvent] class.
type IMTREnergyEVSEClusterEVConnectedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTREnergyEVSEClusterEVConnectedEvent */
	// properties:
	SessionID() objc.IObject /* cross-framework: NSNumber */
	SetSessionID(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTREnergyEVSEClusterEVConnectedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTREnergyEVSEClusterEVConnectedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterEVConnectedEventClass) Alloc() MTREnergyEVSEClusterEVConnectedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEVConnectedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTREnergyEVSEClusterEVConnectedEventClass) New() MTREnergyEVSEClusterEVConnectedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEVConnectedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterEVConnectedEvent) Init() MTREnergyEVSEClusterEVConnectedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEVConnectedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterEVConnectedEvent) Autorelease() MTREnergyEVSEClusterEVConnectedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEVConnectedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterEVConnectedEvent creates a new MTREnergyEVSEClusterEVConnectedEvent instance.
func NewMTREnergyEVSEClusterEVConnectedEvent() MTREnergyEVSEClusterEVConnectedEvent {
	return getMTREnergyEVSEClusterEVConnectedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTREnergyEVSEClusterEVConnectedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVConnectedEvent
type MTREnergyEVSEClusterEVConnectedEvent struct {
	objectivec.Object
}

// MTREnergyEVSEClusterEVConnectedEventFrom constructs a [MTREnergyEVSEClusterEVConnectedEvent] from an unsafe.Pointer.
func MTREnergyEVSEClusterEVConnectedEventFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterEVConnectedEvent {
	return MTREnergyEVSEClusterEVConnectedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTREnergyEVSEClusterEVConnectedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTREnergyEVSEClusterEVConnectedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTREnergyEVSEClusterEVConnectedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTREnergyEVSEClusterEVConnectedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTREnergyEVSEClusterEVConnectedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVConnectedEvent/sessionID
func (m_ MTREnergyEVSEClusterEVConnectedEvent) SessionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sessionID"))
	return rv
}/* debug [instance_properties/getter]: sessionID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVConnectedEvent/sessionID
func (m_ MTREnergyEVSEClusterEVConnectedEvent) SetSessionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionID:"), value)
}/* debug [instance_properties/setter]: sessionID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTREnergyEVSEClusterEVConnectedEvent */



