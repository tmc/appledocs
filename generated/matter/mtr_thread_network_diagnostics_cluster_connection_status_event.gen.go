// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent] class.
var (
	MTRThreadNetworkDiagnosticsClusterConnectionStatusEventClass     _MTRThreadNetworkDiagnosticsClusterConnectionStatusEventClass
	MTRThreadNetworkDiagnosticsClusterConnectionStatusEventClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterConnectionStatusEventClass() _MTRThreadNetworkDiagnosticsClusterConnectionStatusEventClass {
	MTRThreadNetworkDiagnosticsClusterConnectionStatusEventClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterConnectionStatusEventClass = _MTRThreadNetworkDiagnosticsClusterConnectionStatusEventClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent")}
	})
	return MTRThreadNetworkDiagnosticsClusterConnectionStatusEventClass
}

type _MTRThreadNetworkDiagnosticsClusterConnectionStatusEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent] class.
type IMTRThreadNetworkDiagnosticsClusterConnectionStatusEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent
type MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterConnectionStatusEventFrom constructs a [MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterConnectionStatusEventFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent {
	return MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterConnectionStatusEventClass) Alloc() MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDiagnosticsClusterConnectionStatusEventClass) New() MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent) Init() MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent) Autorelease() MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterConnectionStatusEvent creates a new MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent instance.
func NewMTRThreadNetworkDiagnosticsClusterConnectionStatusEvent() MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent {
	return getMTRThreadNetworkDiagnosticsClusterConnectionStatusEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterconnectionstatusevent/connectionstatus
func (m_ MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent) ConnectionStatus() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("connectionStatus"))
	return rv
}


// SetConnectionStatus sets the value of the connectionStatus property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterconnectionstatusevent/connectionstatus
func (m_ MTRThreadNetworkDiagnosticsClusterConnectionStatusEvent) SetConnectionStatus(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConnectionStatus:"), value)
}



