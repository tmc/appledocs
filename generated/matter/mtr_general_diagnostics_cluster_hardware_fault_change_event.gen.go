// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent] class.
var (
	MTRGeneralDiagnosticsClusterHardwareFaultChangeEventClass     _MTRGeneralDiagnosticsClusterHardwareFaultChangeEventClass
	MTRGeneralDiagnosticsClusterHardwareFaultChangeEventClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterHardwareFaultChangeEventClass() _MTRGeneralDiagnosticsClusterHardwareFaultChangeEventClass {
	MTRGeneralDiagnosticsClusterHardwareFaultChangeEventClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterHardwareFaultChangeEventClass = _MTRGeneralDiagnosticsClusterHardwareFaultChangeEventClass{objc.GetClass("MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent")}
	})
	return MTRGeneralDiagnosticsClusterHardwareFaultChangeEventClass
}

type _MTRGeneralDiagnosticsClusterHardwareFaultChangeEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent] class.
type IMTRGeneralDiagnosticsClusterHardwareFaultChangeEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent
type MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent struct {
	objectivec.Object
}

// MTRGeneralDiagnosticsClusterHardwareFaultChangeEventFrom constructs a [MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterHardwareFaultChangeEventFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent {
	return MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterHardwareFaultChangeEventClass) Alloc() MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent {
	rv := objc.Send[MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralDiagnosticsClusterHardwareFaultChangeEventClass) New() MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent {
	rv := objc.Send[MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent) Init() MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent {
	rv := objc.Send[MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent) Autorelease() MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent {
	rv := objc.Send[MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterHardwareFaultChangeEvent creates a new MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent instance.
func NewMTRGeneralDiagnosticsClusterHardwareFaultChangeEvent() MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent {
	return getMTRGeneralDiagnosticsClusterHardwareFaultChangeEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterhardwarefaultchangeevent/previous
func (m_ MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent) Previous() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("previous"))
	return rv
}


// SetPrevious sets the value of the previous property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterhardwarefaultchangeevent/previous
func (m_ MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent) SetPrevious(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrevious:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterhardwarefaultchangeevent/current
func (m_ MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent) Current() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("current"))
	return rv
}


// SetCurrent sets the value of the current property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterhardwarefaultchangeevent/current
func (m_ MTRGeneralDiagnosticsClusterHardwareFaultChangeEvent) SetCurrent(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrent:"), value)
}



