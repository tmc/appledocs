// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent */


/* debug [class_header]: Header for MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent */
// The class instance for the [MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent] class.
var (
	MTRWiFiNetworkDiagnosticsClusterDisconnectionEventClass     _MTRWiFiNetworkDiagnosticsClusterDisconnectionEventClass
	MTRWiFiNetworkDiagnosticsClusterDisconnectionEventClassOnce sync.Once
)

func getMTRWiFiNetworkDiagnosticsClusterDisconnectionEventClass() _MTRWiFiNetworkDiagnosticsClusterDisconnectionEventClass {
	MTRWiFiNetworkDiagnosticsClusterDisconnectionEventClassOnce.Do(func() {
		MTRWiFiNetworkDiagnosticsClusterDisconnectionEventClass = _MTRWiFiNetworkDiagnosticsClusterDisconnectionEventClass{objc.GetClass("MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent")}
	})
	return MTRWiFiNetworkDiagnosticsClusterDisconnectionEventClass
}

type _MTRWiFiNetworkDiagnosticsClusterDisconnectionEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent */
// An interface definition for the [MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent] class.
type IMTRWiFiNetworkDiagnosticsClusterDisconnectionEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent */
	// properties:
	ReasonCode() objc.IObject /* cross-framework: NSNumber */
	SetReasonCode(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRWiFiNetworkDiagnosticsClusterDisconnectionEventClass) Alloc() MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWiFiNetworkDiagnosticsClusterDisconnectionEventClass) New() MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent) Init() MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent) Autorelease() MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWiFiNetworkDiagnosticsClusterDisconnectionEvent creates a new MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent instance.
func NewMTRWiFiNetworkDiagnosticsClusterDisconnectionEvent() MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent {
	return getMTRWiFiNetworkDiagnosticsClusterDisconnectionEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent
type MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent struct {
	objectivec.Object
}

// MTRWiFiNetworkDiagnosticsClusterDisconnectionEventFrom constructs a [MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent] from an unsafe.Pointer.
func MTRWiFiNetworkDiagnosticsClusterDisconnectionEventFrom(ptr unsafe.Pointer) MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent {
	return MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent/reasonCode
func (m_ MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent) ReasonCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("reasonCode"))
	return rv
}/* debug [instance_properties/getter]: reasonCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent/reasonCode
func (m_ MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent) SetReasonCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReasonCode:"), value)
}/* debug [instance_properties/setter]: reasonCode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent */



