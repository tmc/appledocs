// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralDiagnosticsClusterBootReasonEvent] class.
var (
	MTRGeneralDiagnosticsClusterBootReasonEventClass     _MTRGeneralDiagnosticsClusterBootReasonEventClass
	MTRGeneralDiagnosticsClusterBootReasonEventClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterBootReasonEventClass() _MTRGeneralDiagnosticsClusterBootReasonEventClass {
	MTRGeneralDiagnosticsClusterBootReasonEventClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterBootReasonEventClass = _MTRGeneralDiagnosticsClusterBootReasonEventClass{objc.GetClass("MTRGeneralDiagnosticsClusterBootReasonEvent")}
	})
	return MTRGeneralDiagnosticsClusterBootReasonEventClass
}

type _MTRGeneralDiagnosticsClusterBootReasonEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralDiagnosticsClusterBootReasonEvent] class.
type IMTRGeneralDiagnosticsClusterBootReasonEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterBootReasonEvent
type MTRGeneralDiagnosticsClusterBootReasonEvent struct {
	objectivec.Object
}

// MTRGeneralDiagnosticsClusterBootReasonEventFrom constructs a [MTRGeneralDiagnosticsClusterBootReasonEvent] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterBootReasonEventFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterBootReasonEvent {
	return MTRGeneralDiagnosticsClusterBootReasonEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterBootReasonEventClass) Alloc() MTRGeneralDiagnosticsClusterBootReasonEvent {
	rv := objc.Send[MTRGeneralDiagnosticsClusterBootReasonEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralDiagnosticsClusterBootReasonEventClass) New() MTRGeneralDiagnosticsClusterBootReasonEvent {
	rv := objc.Send[MTRGeneralDiagnosticsClusterBootReasonEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterBootReasonEvent) Init() MTRGeneralDiagnosticsClusterBootReasonEvent {
	rv := objc.Send[MTRGeneralDiagnosticsClusterBootReasonEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterBootReasonEvent) Autorelease() MTRGeneralDiagnosticsClusterBootReasonEvent {
	rv := objc.Send[MTRGeneralDiagnosticsClusterBootReasonEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterBootReasonEvent creates a new MTRGeneralDiagnosticsClusterBootReasonEvent instance.
func NewMTRGeneralDiagnosticsClusterBootReasonEvent() MTRGeneralDiagnosticsClusterBootReasonEvent {
	return getMTRGeneralDiagnosticsClusterBootReasonEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterbootreasonevent/bootreason
func (m_ MTRGeneralDiagnosticsClusterBootReasonEvent) BootReason() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("bootReason"))
	return rv
}


// SetBootReason sets the value of the bootReason property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterbootreasonevent/bootreason
func (m_ MTRGeneralDiagnosticsClusterBootReasonEvent) SetBootReason(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBootReason:"), value)
}



