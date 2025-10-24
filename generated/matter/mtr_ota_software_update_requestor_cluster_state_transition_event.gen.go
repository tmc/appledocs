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
	// properties:
	NewState() objc.IObject /* cross-framework: NSNumber */
	SetNewState(value objc.IObject /* cross-framework: NSNumber */)
	PreviousState() objc.IObject /* cross-framework: NSNumber */
	SetPreviousState(value objc.IObject /* cross-framework: NSNumber */)
	Reason() objc.IObject /* cross-framework: NSNumber */
	SetReason(value objc.IObject /* cross-framework: NSNumber */)
	TargetSoftwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetTargetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-1xzd5/newstate
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) NewState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-1xzd5/newstate
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) SetNewState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewState:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-1xzd5/previousstate
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) PreviousState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("previousState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-1xzd5/previousstate
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) SetPreviousState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreviousState:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-1xzd5/reason
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) Reason() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("reason"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-1xzd5/reason
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) SetReason(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReason:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-1xzd5/targetsoftwareversion
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) TargetSoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("targetSoftwareVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdaterequestorclusterstatetransitionevent-1xzd5/targetsoftwareversion
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) SetTargetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetSoftwareVersion:"), value)
}



