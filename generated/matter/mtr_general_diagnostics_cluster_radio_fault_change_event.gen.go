// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralDiagnosticsClusterRadioFaultChangeEvent] class.
var (
	MTRGeneralDiagnosticsClusterRadioFaultChangeEventClass     _MTRGeneralDiagnosticsClusterRadioFaultChangeEventClass
	MTRGeneralDiagnosticsClusterRadioFaultChangeEventClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterRadioFaultChangeEventClass() _MTRGeneralDiagnosticsClusterRadioFaultChangeEventClass {
	MTRGeneralDiagnosticsClusterRadioFaultChangeEventClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterRadioFaultChangeEventClass = _MTRGeneralDiagnosticsClusterRadioFaultChangeEventClass{objc.GetClass("MTRGeneralDiagnosticsClusterRadioFaultChangeEvent")}
	})
	return MTRGeneralDiagnosticsClusterRadioFaultChangeEventClass
}

type _MTRGeneralDiagnosticsClusterRadioFaultChangeEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralDiagnosticsClusterRadioFaultChangeEvent] class.
type IMTRGeneralDiagnosticsClusterRadioFaultChangeEvent interface {
	objectivec.IObject
	Current() unsafe.Pointer
	SetCurrent(value unsafe.Pointer)
	Previous() unsafe.Pointer
	SetPrevious(value unsafe.Pointer)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterRadioFaultChangeEvent
type MTRGeneralDiagnosticsClusterRadioFaultChangeEvent struct {
	objectivec.Object
}

// MTRGeneralDiagnosticsClusterRadioFaultChangeEventFrom constructs a [MTRGeneralDiagnosticsClusterRadioFaultChangeEvent] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterRadioFaultChangeEventFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterRadioFaultChangeEvent {
	return MTRGeneralDiagnosticsClusterRadioFaultChangeEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterRadioFaultChangeEventClass) Alloc() MTRGeneralDiagnosticsClusterRadioFaultChangeEvent {
	rv := objc.Send[MTRGeneralDiagnosticsClusterRadioFaultChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralDiagnosticsClusterRadioFaultChangeEventClass) New() MTRGeneralDiagnosticsClusterRadioFaultChangeEvent {
	rv := objc.Send[MTRGeneralDiagnosticsClusterRadioFaultChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterRadioFaultChangeEvent) Init() MTRGeneralDiagnosticsClusterRadioFaultChangeEvent {
	rv := objc.Send[MTRGeneralDiagnosticsClusterRadioFaultChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterRadioFaultChangeEvent) Autorelease() MTRGeneralDiagnosticsClusterRadioFaultChangeEvent {
	rv := objc.Send[MTRGeneralDiagnosticsClusterRadioFaultChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterRadioFaultChangeEvent creates a new MTRGeneralDiagnosticsClusterRadioFaultChangeEvent instance.
func NewMTRGeneralDiagnosticsClusterRadioFaultChangeEvent() MTRGeneralDiagnosticsClusterRadioFaultChangeEvent {
	return getMTRGeneralDiagnosticsClusterRadioFaultChangeEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterradiofaultchangeevent/current
func (m_ MTRGeneralDiagnosticsClusterRadioFaultChangeEvent) Current() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("current"))
	return rv
}


// SetCurrent sets the value of the current property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterradiofaultchangeevent/current
func (m_ MTRGeneralDiagnosticsClusterRadioFaultChangeEvent) SetCurrent(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterradiofaultchangeevent/previous
func (m_ MTRGeneralDiagnosticsClusterRadioFaultChangeEvent) Previous() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("previous"))
	return rv
}


// SetPrevious sets the value of the previous property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneraldiagnosticsclusterradiofaultchangeevent/previous
func (m_ MTRGeneralDiagnosticsClusterRadioFaultChangeEvent) SetPrevious(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrevious:"), value)
}



