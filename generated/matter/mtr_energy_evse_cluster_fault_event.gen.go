// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTREnergyEVSEClusterFaultEvent] class.
var (
	MTREnergyEVSEClusterFaultEventClass     _MTREnergyEVSEClusterFaultEventClass
	MTREnergyEVSEClusterFaultEventClassOnce sync.Once
)

func getMTREnergyEVSEClusterFaultEventClass() _MTREnergyEVSEClusterFaultEventClass {
	MTREnergyEVSEClusterFaultEventClassOnce.Do(func() {
		MTREnergyEVSEClusterFaultEventClass = _MTREnergyEVSEClusterFaultEventClass{objc.GetClass("MTREnergyEVSEClusterFaultEvent")}
	})
	return MTREnergyEVSEClusterFaultEventClass
}

type _MTREnergyEVSEClusterFaultEventClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterFaultEvent] class.
type IMTREnergyEVSEClusterFaultEvent interface {
	objectivec.IObject
	// properties:
	FaultStateCurrentState() objc.IObject /* cross-framework: NSNumber */
	SetFaultStateCurrentState(value objc.IObject /* cross-framework: NSNumber */)
	FaultStatePreviousState() objc.IObject /* cross-framework: NSNumber */
	SetFaultStatePreviousState(value objc.IObject /* cross-framework: NSNumber */)
	SessionID() objc.IObject /* cross-framework: NSNumber */
	SetSessionID(value objc.IObject /* cross-framework: NSNumber */)
	State() objc.IObject /* cross-framework: NSNumber */
	SetState(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterFaultEvent
type MTREnergyEVSEClusterFaultEvent struct {
	objectivec.Object
}

// MTREnergyEVSEClusterFaultEventFrom constructs a [MTREnergyEVSEClusterFaultEvent] from an unsafe.Pointer.
func MTREnergyEVSEClusterFaultEventFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterFaultEvent {
	return MTREnergyEVSEClusterFaultEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterFaultEventClass) Alloc() MTREnergyEVSEClusterFaultEvent {
	rv := objc.Send[MTREnergyEVSEClusterFaultEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterFaultEventClass) New() MTREnergyEVSEClusterFaultEvent {
	rv := objc.Send[MTREnergyEVSEClusterFaultEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterFaultEvent) Init() MTREnergyEVSEClusterFaultEvent {
	rv := objc.Send[MTREnergyEVSEClusterFaultEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterFaultEvent) Autorelease() MTREnergyEVSEClusterFaultEvent {
	rv := objc.Send[MTREnergyEVSEClusterFaultEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterFaultEvent creates a new MTREnergyEVSEClusterFaultEvent instance.
func NewMTREnergyEVSEClusterFaultEvent() MTREnergyEVSEClusterFaultEvent {
	return getMTREnergyEVSEClusterFaultEventClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterFaultEvent/faultStateCurrentState
func (m_ MTREnergyEVSEClusterFaultEvent) FaultStateCurrentState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("faultStateCurrentState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterFaultEvent/faultStateCurrentState
func (m_ MTREnergyEVSEClusterFaultEvent) SetFaultStateCurrentState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFaultStateCurrentState:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterFaultEvent/faultStatePreviousState
func (m_ MTREnergyEVSEClusterFaultEvent) FaultStatePreviousState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("faultStatePreviousState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterFaultEvent/faultStatePreviousState
func (m_ MTREnergyEVSEClusterFaultEvent) SetFaultStatePreviousState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFaultStatePreviousState:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterFaultEvent/sessionID
func (m_ MTREnergyEVSEClusterFaultEvent) SessionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sessionID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterFaultEvent/sessionID
func (m_ MTREnergyEVSEClusterFaultEvent) SetSessionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterFaultEvent/state
func (m_ MTREnergyEVSEClusterFaultEvent) State() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("state"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterFaultEvent/state
func (m_ MTREnergyEVSEClusterFaultEvent) SetState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}



