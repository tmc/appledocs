// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTREnergyEVSEClusterEnergyTransferStoppedEvent] class.
var (
	MTREnergyEVSEClusterEnergyTransferStoppedEventClass     _MTREnergyEVSEClusterEnergyTransferStoppedEventClass
	MTREnergyEVSEClusterEnergyTransferStoppedEventClassOnce sync.Once
)

func getMTREnergyEVSEClusterEnergyTransferStoppedEventClass() _MTREnergyEVSEClusterEnergyTransferStoppedEventClass {
	MTREnergyEVSEClusterEnergyTransferStoppedEventClassOnce.Do(func() {
		MTREnergyEVSEClusterEnergyTransferStoppedEventClass = _MTREnergyEVSEClusterEnergyTransferStoppedEventClass{objc.GetClass("MTREnergyEVSEClusterEnergyTransferStoppedEvent")}
	})
	return MTREnergyEVSEClusterEnergyTransferStoppedEventClass
}

type _MTREnergyEVSEClusterEnergyTransferStoppedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterEnergyTransferStoppedEvent] class.
type IMTREnergyEVSEClusterEnergyTransferStoppedEvent interface {
	objectivec.IObject
	// properties:
	EnergyTransferred() objc.IObject /* cross-framework: NSNumber */
	SetEnergyTransferred(value objc.IObject /* cross-framework: NSNumber */)
	Reason() objc.IObject /* cross-framework: NSNumber */
	SetReason(value objc.IObject /* cross-framework: NSNumber */)
	SessionID() objc.IObject /* cross-framework: NSNumber */
	SetSessionID(value objc.IObject /* cross-framework: NSNumber */)
	State() objc.IObject /* cross-framework: NSNumber */
	SetState(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent
type MTREnergyEVSEClusterEnergyTransferStoppedEvent struct {
	objectivec.Object
}

// MTREnergyEVSEClusterEnergyTransferStoppedEventFrom constructs a [MTREnergyEVSEClusterEnergyTransferStoppedEvent] from an unsafe.Pointer.
func MTREnergyEVSEClusterEnergyTransferStoppedEventFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterEnergyTransferStoppedEvent {
	return MTREnergyEVSEClusterEnergyTransferStoppedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterEnergyTransferStoppedEventClass) Alloc() MTREnergyEVSEClusterEnergyTransferStoppedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEnergyTransferStoppedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterEnergyTransferStoppedEventClass) New() MTREnergyEVSEClusterEnergyTransferStoppedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEnergyTransferStoppedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) Init() MTREnergyEVSEClusterEnergyTransferStoppedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEnergyTransferStoppedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) Autorelease() MTREnergyEVSEClusterEnergyTransferStoppedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEnergyTransferStoppedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterEnergyTransferStoppedEvent creates a new MTREnergyEVSEClusterEnergyTransferStoppedEvent instance.
func NewMTREnergyEVSEClusterEnergyTransferStoppedEvent() MTREnergyEVSEClusterEnergyTransferStoppedEvent {
	return getMTREnergyEVSEClusterEnergyTransferStoppedEventClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/energyTransferred
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) EnergyTransferred() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("energyTransferred"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/energyTransferred
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) SetEnergyTransferred(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnergyTransferred:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/reason
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) Reason() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("reason"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/reason
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) SetReason(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReason:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/sessionID
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) SessionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sessionID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/sessionID
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) SetSessionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/state
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) State() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("state"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/state
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) SetState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}



