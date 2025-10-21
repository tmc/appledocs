// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/energyTransferred
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) EnergyTransferred() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("energyTransferred"))
	return rv
}


// SetEnergyTransferred sets the value of the energyTransferred property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/energyTransferred
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) SetEnergyTransferred(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnergyTransferred:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/reason
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) Reason() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("reason"))
	return rv
}


// SetReason sets the value of the reason property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/reason
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) SetReason(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReason:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/sessionID
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) SessionID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("sessionID"))
	return rv
}


// SetSessionID sets the value of the sessionID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/sessionID
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) SetSessionID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionID:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/state
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStoppedEvent/state
func (m_ MTREnergyEVSEClusterEnergyTransferStoppedEvent) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}


