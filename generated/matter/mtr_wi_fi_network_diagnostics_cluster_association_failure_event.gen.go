// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent] class.
var (
	MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass     _MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass
	MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClassOnce sync.Once
)

func getMTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass() _MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass {
	MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClassOnce.Do(func() {
		MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass = _MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass{objc.GetClass("MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent")}
	})
	return MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass
}

type _MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent] class.
type IMTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent interface {
	objectivec.IObject
	// properties:
	AssociationFailure() objc.IObject /* cross-framework: NSNumber */
	SetAssociationFailure(value objc.IObject /* cross-framework: NSNumber */)
	AssociationFailureCause() objc.IObject /* cross-framework: NSNumber */
	SetAssociationFailureCause(value objc.IObject /* cross-framework: NSNumber */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent
type MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent struct {
	objectivec.Object
}

// MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventFrom constructs a [MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent] from an unsafe.Pointer.
func MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventFrom(ptr unsafe.Pointer) MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent {
	return MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass) Alloc() MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass) New() MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent) Init() MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent) Autorelease() MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent creates a new MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent instance.
func NewMTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent() MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent {
	return getMTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwifinetworkdiagnosticsclusterassociationfailureevent/associationfailure
func (m_ MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent) AssociationFailure() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("associationFailure"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwifinetworkdiagnosticsclusterassociationfailureevent/associationfailure
func (m_ MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent) SetAssociationFailure(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAssociationFailure:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwifinetworkdiagnosticsclusterassociationfailureevent/associationfailurecause
func (m_ MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent) AssociationFailureCause() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("associationFailureCause"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwifinetworkdiagnosticsclusterassociationfailureevent/associationfailurecause
func (m_ MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent) SetAssociationFailureCause(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAssociationFailureCause:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwifinetworkdiagnosticsclusterassociationfailureevent/status
func (m_ MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwifinetworkdiagnosticsclusterassociationfailureevent/status
func (m_ MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}



