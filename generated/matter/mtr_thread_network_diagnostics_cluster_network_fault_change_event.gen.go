// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent] class.
var (
	MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass     _MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass
	MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass() _MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass {
	MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass = _MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent")}
	})
	return MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass
}

type _MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent] class.
type IMTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent
type MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventFrom constructs a [MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent {
	return MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass) Alloc() MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass) New() MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent) Init() MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent) Autorelease() MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent creates a new MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent instance.
func NewMTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent() MTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEvent {
	return getMTRThreadNetworkDiagnosticsClusterNetworkFaultChangeEventClass().New()
}




