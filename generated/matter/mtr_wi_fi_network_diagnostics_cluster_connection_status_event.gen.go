// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent */


/* debug [class_header]: Header for MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent */
// The class instance for the [MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent] class.
var (
	MTRWiFiNetworkDiagnosticsClusterConnectionStatusEventClass     _MTRWiFiNetworkDiagnosticsClusterConnectionStatusEventClass
	MTRWiFiNetworkDiagnosticsClusterConnectionStatusEventClassOnce sync.Once
)

func getMTRWiFiNetworkDiagnosticsClusterConnectionStatusEventClass() _MTRWiFiNetworkDiagnosticsClusterConnectionStatusEventClass {
	MTRWiFiNetworkDiagnosticsClusterConnectionStatusEventClassOnce.Do(func() {
		MTRWiFiNetworkDiagnosticsClusterConnectionStatusEventClass = _MTRWiFiNetworkDiagnosticsClusterConnectionStatusEventClass{objc.GetClass("MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent")}
	})
	return MTRWiFiNetworkDiagnosticsClusterConnectionStatusEventClass
}

type _MTRWiFiNetworkDiagnosticsClusterConnectionStatusEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent */
// An interface definition for the [MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent] class.
type IMTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent */
	// properties:
	ConnectionStatus() objc.IObject /* cross-framework: NSNumber */
	SetConnectionStatus(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRWiFiNetworkDiagnosticsClusterConnectionStatusEventClass) Alloc() MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWiFiNetworkDiagnosticsClusterConnectionStatusEventClass) New() MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent) Init() MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent) Autorelease() MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent creates a new MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent instance.
func NewMTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent() MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent {
	return getMTRWiFiNetworkDiagnosticsClusterConnectionStatusEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent
type MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent struct {
	objectivec.Object
}

// MTRWiFiNetworkDiagnosticsClusterConnectionStatusEventFrom constructs a [MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent] from an unsafe.Pointer.
func MTRWiFiNetworkDiagnosticsClusterConnectionStatusEventFrom(ptr unsafe.Pointer) MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent {
	return MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent/connectionStatus
func (m_ MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent) ConnectionStatus() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("connectionStatus"))
	return rv
}/* debug [instance_properties/getter]: connectionStatus */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent/connectionStatus
func (m_ MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent) SetConnectionStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConnectionStatus:"), value)
}/* debug [instance_properties/setter]: connectionStatus */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent */



