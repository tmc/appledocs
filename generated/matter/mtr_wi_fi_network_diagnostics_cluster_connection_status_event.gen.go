// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent] class.
type IMTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent interface {
	objectivec.IObject
	// properties:
	ConnectionStatus() objc.IObject /* cross-framework: NSNumber */
	SetConnectionStatus(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent
type MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent struct {
	objectivec.Object
}

// MTRWiFiNetworkDiagnosticsClusterConnectionStatusEventFrom constructs a [MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent] from an unsafe.Pointer.
func MTRWiFiNetworkDiagnosticsClusterConnectionStatusEventFrom(ptr unsafe.Pointer) MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent {
	return MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWiFiNetworkDiagnosticsClusterConnectionStatusEventClass) Alloc() MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwifinetworkdiagnosticsclusterconnectionstatusevent/connectionstatus
func (m_ MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent) ConnectionStatus() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("connectionStatus"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwifinetworkdiagnosticsclusterconnectionstatusevent/connectionstatus
func (m_ MTRWiFiNetworkDiagnosticsClusterConnectionStatusEvent) SetConnectionStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConnectionStatus:"), value)
}



