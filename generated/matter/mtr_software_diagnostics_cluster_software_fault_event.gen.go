// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclustersoftwarefaultevent/name
func (m_ MTRSoftwareDiagnosticsClusterSoftwareFaultEvent) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclustersoftwarefaultevent/name
func (m_ MTRSoftwareDiagnosticsClusterSoftwareFaultEvent) SetName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclustersoftwarefaultevent/id
func (m_ MTRSoftwareDiagnosticsClusterSoftwareFaultEvent) Id() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("id"))
	return rv
}


// SetId sets the value of the id property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclustersoftwarefaultevent/id
func (m_ MTRSoftwareDiagnosticsClusterSoftwareFaultEvent) SetId(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclustersoftwarefaultevent/faultrecording
func (m_ MTRSoftwareDiagnosticsClusterSoftwareFaultEvent) FaultRecording() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("faultRecording"))
	return rv
}


// SetFaultRecording sets the value of the faultRecording property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsoftwarediagnosticsclustersoftwarefaultevent/faultrecording
func (m_ MTRSoftwareDiagnosticsClusterSoftwareFaultEvent) SetFaultRecording(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFaultRecording:"), value)
}



