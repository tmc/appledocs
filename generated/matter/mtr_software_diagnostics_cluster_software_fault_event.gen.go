// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRSoftwareDiagnosticsClusterSoftwareFaultEvent] class.
var (
	MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass     _MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass
	MTRSoftwareDiagnosticsClusterSoftwareFaultEventClassOnce sync.Once
)

func getMTRSoftwareDiagnosticsClusterSoftwareFaultEventClass() _MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass {
	MTRSoftwareDiagnosticsClusterSoftwareFaultEventClassOnce.Do(func() {
		MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass = _MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass{objc.GetClass("MTRSoftwareDiagnosticsClusterSoftwareFaultEvent")}
	})
	return MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass
}

type _MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSoftwareDiagnosticsClusterSoftwareFaultEvent] class.
type IMTRSoftwareDiagnosticsClusterSoftwareFaultEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterSoftwareFaultEvent
type MTRSoftwareDiagnosticsClusterSoftwareFaultEvent struct {
	objectivec.Object
}

// MTRSoftwareDiagnosticsClusterSoftwareFaultEventFrom constructs a [MTRSoftwareDiagnosticsClusterSoftwareFaultEvent] from an unsafe.Pointer.
func MTRSoftwareDiagnosticsClusterSoftwareFaultEventFrom(ptr unsafe.Pointer) MTRSoftwareDiagnosticsClusterSoftwareFaultEvent {
	return MTRSoftwareDiagnosticsClusterSoftwareFaultEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass) Alloc() MTRSoftwareDiagnosticsClusterSoftwareFaultEvent {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterSoftwareFaultEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass) New() MTRSoftwareDiagnosticsClusterSoftwareFaultEvent {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterSoftwareFaultEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSoftwareDiagnosticsClusterSoftwareFaultEvent) Init() MTRSoftwareDiagnosticsClusterSoftwareFaultEvent {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterSoftwareFaultEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSoftwareDiagnosticsClusterSoftwareFaultEvent) Autorelease() MTRSoftwareDiagnosticsClusterSoftwareFaultEvent {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterSoftwareFaultEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSoftwareDiagnosticsClusterSoftwareFaultEvent creates a new MTRSoftwareDiagnosticsClusterSoftwareFaultEvent instance.
func NewMTRSoftwareDiagnosticsClusterSoftwareFaultEvent() MTRSoftwareDiagnosticsClusterSoftwareFaultEvent {
	return getMTRSoftwareDiagnosticsClusterSoftwareFaultEventClass().New()
}




