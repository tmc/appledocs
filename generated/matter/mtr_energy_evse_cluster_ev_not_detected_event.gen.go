// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTREnergyEVSEClusterEVNotDetectedEvent] class.
var (
	MTREnergyEVSEClusterEVNotDetectedEventClass     _MTREnergyEVSEClusterEVNotDetectedEventClass
	MTREnergyEVSEClusterEVNotDetectedEventClassOnce sync.Once
)

func getMTREnergyEVSEClusterEVNotDetectedEventClass() _MTREnergyEVSEClusterEVNotDetectedEventClass {
	MTREnergyEVSEClusterEVNotDetectedEventClassOnce.Do(func() {
		MTREnergyEVSEClusterEVNotDetectedEventClass = _MTREnergyEVSEClusterEVNotDetectedEventClass{objc.GetClass("MTREnergyEVSEClusterEVNotDetectedEvent")}
	})
	return MTREnergyEVSEClusterEVNotDetectedEventClass
}

type _MTREnergyEVSEClusterEVNotDetectedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterEVNotDetectedEvent] class.
type IMTREnergyEVSEClusterEVNotDetectedEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent
type MTREnergyEVSEClusterEVNotDetectedEvent struct {
	objectivec.Object
}

// MTREnergyEVSEClusterEVNotDetectedEventFrom constructs a [MTREnergyEVSEClusterEVNotDetectedEvent] from an unsafe.Pointer.
func MTREnergyEVSEClusterEVNotDetectedEventFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterEVNotDetectedEvent {
	return MTREnergyEVSEClusterEVNotDetectedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterEVNotDetectedEventClass) Alloc() MTREnergyEVSEClusterEVNotDetectedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEVNotDetectedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterEVNotDetectedEventClass) New() MTREnergyEVSEClusterEVNotDetectedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEVNotDetectedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) Init() MTREnergyEVSEClusterEVNotDetectedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEVNotDetectedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) Autorelease() MTREnergyEVSEClusterEVNotDetectedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEVNotDetectedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterEVNotDetectedEvent creates a new MTREnergyEVSEClusterEVNotDetectedEvent instance.
func NewMTREnergyEVSEClusterEVNotDetectedEvent() MTREnergyEVSEClusterEVNotDetectedEvent {
	return getMTREnergyEVSEClusterEVNotDetectedEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/sessionDuration
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SessionDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("sessionDuration"))
	return rv
}


// SetSessionDuration sets the value of the sessionDuration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/sessionDuration
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SetSessionDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionDuration:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/sessionEnergyCharged
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SessionEnergyCharged() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("sessionEnergyCharged"))
	return rv
}


// SetSessionEnergyCharged sets the value of the sessionEnergyCharged property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/sessionEnergyCharged
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SetSessionEnergyCharged(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionEnergyCharged:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/sessionID
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SessionID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("sessionID"))
	return rv
}


// SetSessionID sets the value of the sessionID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/sessionID
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SetSessionID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionID:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/state
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/state
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}


