// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent] class.
type IMTRWiFiNetworkDiagnosticsClusterDisconnectionEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent
type MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent struct {
	objectivec.Object
}

// MTRWiFiNetworkDiagnosticsClusterDisconnectionEventFrom constructs a [MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent] from an unsafe.Pointer.
func MTRWiFiNetworkDiagnosticsClusterDisconnectionEventFrom(ptr unsafe.Pointer) MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent {
	return MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWiFiNetworkDiagnosticsClusterDisconnectionEventClass) Alloc() MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwifinetworkdiagnosticsclusterdisconnectionevent/reasoncode
func (m_ MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent) ReasonCode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("reasonCode"))
	return rv
}


// SetReasonCode sets the value of the reasonCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwifinetworkdiagnosticsclusterdisconnectionevent/reasoncode
func (m_ MTRWiFiNetworkDiagnosticsClusterDisconnectionEvent) SetReasonCode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReasonCode:"), value)
}



