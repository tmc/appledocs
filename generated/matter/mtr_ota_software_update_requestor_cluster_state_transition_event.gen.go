// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent] class.
var (
	MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass     _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass
	MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClassOnce sync.Once
)

func getMTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass() _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass {
	MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClassOnce.Do(func() {
		MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass = _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass{objc.GetClass("MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent")}
	})
	return MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass
}

type _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass struct {
	class objc.Class
}

// An interface definition for the [MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent] class.
type IMTROtaSoftwareUpdateRequestorClusterStateTransitionEvent interface {
	IMTROTASoftwareUpdateRequestorClusterStateTransitionEvent
	NewState() foundation.Number
	SetNewState(value foundation.INumber)
	PreviousState() foundation.Number
	SetPreviousState(value foundation.INumber)
	Reason() foundation.Number
	SetReason(value foundation.INumber)
	TargetSoftwareVersion() foundation.Number
	SetTargetSoftwareVersion(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent-1xzd5
type MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent struct {
	MTROTASoftwareUpdateRequestorClusterStateTransitionEvent
}

// MTROtaSoftwareUpdateRequestorClusterStateTransitionEventFrom constructs a [MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent] from an unsafe.Pointer.
func MTROtaSoftwareUpdateRequestorClusterStateTransitionEventFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	return MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent{
		MTROTASoftwareUpdateRequestorClusterStateTransitionEvent: MTROTASoftwareUpdateRequestorClusterStateTransitionEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass) Alloc() MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass) New() MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) Init() MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) Autorelease() MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateRequestorClusterStateTransitionEvent creates a new MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent instance.
func NewMTROtaSoftwareUpdateRequestorClusterStateTransitionEvent() MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	return getMTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-1xzd5/newstate
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) NewState() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("newState"))
	return rv
}


// SetNewState sets the value of the newState property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-1xzd5/newstate
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) SetNewState(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewState:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-1xzd5/previousstate
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) PreviousState() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("previousState"))
	return rv
}


// SetPreviousState sets the value of the previousState property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-1xzd5/previousstate
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) SetPreviousState(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreviousState:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-1xzd5/reason
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) Reason() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("reason"))
	return rv
}


// SetReason sets the value of the reason property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-1xzd5/reason
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) SetReason(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReason:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-1xzd5/targetsoftwareversion
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) TargetSoftwareVersion() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("targetSoftwareVersion"))
	return rv
}


// SetTargetSoftwareVersion sets the value of the targetSoftwareVersion property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-1xzd5/targetsoftwareversion
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) SetTargetSoftwareVersion(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetSoftwareVersion:"), value)
}



