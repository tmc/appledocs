// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROTASoftwareUpdateRequestorClusterStateTransitionEvent] class.
var (
	MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass     _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass
	MTROTASoftwareUpdateRequestorClusterStateTransitionEventClassOnce sync.Once
)

func getMTROTASoftwareUpdateRequestorClusterStateTransitionEventClass() _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass {
	MTROTASoftwareUpdateRequestorClusterStateTransitionEventClassOnce.Do(func() {
		MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass = _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass{objc.GetClass("MTROTASoftwareUpdateRequestorClusterStateTransitionEvent")}
	})
	return MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass
}

type _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass struct {
	class objc.Class
}

// An interface definition for the [MTROTASoftwareUpdateRequestorClusterStateTransitionEvent] class.
type IMTROTASoftwareUpdateRequestorClusterStateTransitionEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterStateTransitionEvent-3xhxb
type MTROTASoftwareUpdateRequestorClusterStateTransitionEvent struct {
	objectivec.Object
}

// MTROTASoftwareUpdateRequestorClusterStateTransitionEventFrom constructs a [MTROTASoftwareUpdateRequestorClusterStateTransitionEvent] from an unsafe.Pointer.
func MTROTASoftwareUpdateRequestorClusterStateTransitionEventFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	return MTROTASoftwareUpdateRequestorClusterStateTransitionEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass) Alloc() MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterStateTransitionEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass) New() MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterStateTransitionEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) Init() MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterStateTransitionEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) Autorelease() MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterStateTransitionEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateRequestorClusterStateTransitionEvent creates a new MTROTASoftwareUpdateRequestorClusterStateTransitionEvent instance.
func NewMTROTASoftwareUpdateRequestorClusterStateTransitionEvent() MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	return getMTROTASoftwareUpdateRequestorClusterStateTransitionEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-3xhxb/newstate
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) NewState() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("newState"))
	return rv
}


// SetNewState sets the value of the newState property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-3xhxb/newstate
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) SetNewState(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewState:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-3xhxb/previousstate
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) PreviousState() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("previousState"))
	return rv
}


// SetPreviousState sets the value of the previousState property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-3xhxb/previousstate
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) SetPreviousState(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreviousState:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-3xhxb/reason
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) Reason() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("reason"))
	return rv
}


// SetReason sets the value of the reason property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-3xhxb/reason
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) SetReason(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReason:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-3xhxb/targetsoftwareversion
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) TargetSoftwareVersion() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("targetSoftwareVersion"))
	return rv
}


// SetTargetSoftwareVersion sets the value of the targetSoftwareVersion property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-3xhxb/targetsoftwareversion
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) SetTargetSoftwareVersion(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetSoftwareVersion:"), value)
}



