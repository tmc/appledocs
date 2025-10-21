// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent] class.
var (
	MTRGeneralDiagnosticsClusterNetworkFaultChangeEventClass     _MTRGeneralDiagnosticsClusterNetworkFaultChangeEventClass
	MTRGeneralDiagnosticsClusterNetworkFaultChangeEventClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterNetworkFaultChangeEventClass() _MTRGeneralDiagnosticsClusterNetworkFaultChangeEventClass {
	MTRGeneralDiagnosticsClusterNetworkFaultChangeEventClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterNetworkFaultChangeEventClass = _MTRGeneralDiagnosticsClusterNetworkFaultChangeEventClass{objc.GetClass("MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent")}
	})
	return MTRGeneralDiagnosticsClusterNetworkFaultChangeEventClass
}

type _MTRGeneralDiagnosticsClusterNetworkFaultChangeEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent] class.
type IMTRGeneralDiagnosticsClusterNetworkFaultChangeEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent
type MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent struct {
	objectivec.Object
}

// MTRGeneralDiagnosticsClusterNetworkFaultChangeEventFrom constructs a [MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterNetworkFaultChangeEventFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent {
	return MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterNetworkFaultChangeEventClass) Alloc() MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralDiagnosticsClusterNetworkFaultChangeEventClass) New() MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent) Init() MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent) Autorelease() MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterNetworkFaultChangeEvent creates a new MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent instance.
func NewMTRGeneralDiagnosticsClusterNetworkFaultChangeEvent() MTRGeneralDiagnosticsClusterNetworkFaultChangeEvent {
	return getMTRGeneralDiagnosticsClusterNetworkFaultChangeEventClass().New()
}




