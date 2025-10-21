// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTREnergyEVSEClusterEnergyTransferStartedEvent] class.
var (
	MTREnergyEVSEClusterEnergyTransferStartedEventClass     _MTREnergyEVSEClusterEnergyTransferStartedEventClass
	MTREnergyEVSEClusterEnergyTransferStartedEventClassOnce sync.Once
)

func getMTREnergyEVSEClusterEnergyTransferStartedEventClass() _MTREnergyEVSEClusterEnergyTransferStartedEventClass {
	MTREnergyEVSEClusterEnergyTransferStartedEventClassOnce.Do(func() {
		MTREnergyEVSEClusterEnergyTransferStartedEventClass = _MTREnergyEVSEClusterEnergyTransferStartedEventClass{objc.GetClass("MTREnergyEVSEClusterEnergyTransferStartedEvent")}
	})
	return MTREnergyEVSEClusterEnergyTransferStartedEventClass
}

type _MTREnergyEVSEClusterEnergyTransferStartedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterEnergyTransferStartedEvent] class.
type IMTREnergyEVSEClusterEnergyTransferStartedEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStartedEvent
type MTREnergyEVSEClusterEnergyTransferStartedEvent struct {
	objectivec.Object
}

// MTREnergyEVSEClusterEnergyTransferStartedEventFrom constructs a [MTREnergyEVSEClusterEnergyTransferStartedEvent] from an unsafe.Pointer.
func MTREnergyEVSEClusterEnergyTransferStartedEventFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterEnergyTransferStartedEvent {
	return MTREnergyEVSEClusterEnergyTransferStartedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterEnergyTransferStartedEventClass) Alloc() MTREnergyEVSEClusterEnergyTransferStartedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEnergyTransferStartedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterEnergyTransferStartedEventClass) New() MTREnergyEVSEClusterEnergyTransferStartedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEnergyTransferStartedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterEnergyTransferStartedEvent) Init() MTREnergyEVSEClusterEnergyTransferStartedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEnergyTransferStartedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterEnergyTransferStartedEvent) Autorelease() MTREnergyEVSEClusterEnergyTransferStartedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEnergyTransferStartedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterEnergyTransferStartedEvent creates a new MTREnergyEVSEClusterEnergyTransferStartedEvent instance.
func NewMTREnergyEVSEClusterEnergyTransferStartedEvent() MTREnergyEVSEClusterEnergyTransferStartedEvent {
	return getMTREnergyEVSEClusterEnergyTransferStartedEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStartedEvent/maximumCurrent
func (m_ MTREnergyEVSEClusterEnergyTransferStartedEvent) MaximumCurrent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("maximumCurrent"))
	return rv
}


// SetMaximumCurrent sets the value of the maximumCurrent property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStartedEvent/maximumCurrent
func (m_ MTREnergyEVSEClusterEnergyTransferStartedEvent) SetMaximumCurrent(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumCurrent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStartedEvent/sessionID
func (m_ MTREnergyEVSEClusterEnergyTransferStartedEvent) SessionID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("sessionID"))
	return rv
}


// SetSessionID sets the value of the sessionID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStartedEvent/sessionID
func (m_ MTREnergyEVSEClusterEnergyTransferStartedEvent) SetSessionID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStartedEvent/state
func (m_ MTREnergyEVSEClusterEnergyTransferStartedEvent) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEnergyTransferStartedEvent/state
func (m_ MTREnergyEVSEClusterEnergyTransferStartedEvent) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}



