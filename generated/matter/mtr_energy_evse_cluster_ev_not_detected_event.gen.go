// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	SessionDuration() objc.IObject /* cross-framework: NSNumber */
	SetSessionDuration(value objc.IObject /* cross-framework: NSNumber */)
	SessionEnergyCharged() objc.IObject /* cross-framework: NSNumber */
	SetSessionEnergyCharged(value objc.IObject /* cross-framework: NSNumber */)
	SessionID() objc.IObject /* cross-framework: NSNumber */
	SetSessionID(value objc.IObject /* cross-framework: NSNumber */)
	State() objc.IObject /* cross-framework: NSNumber */
	SetState(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/sessionDuration
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SessionDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sessionDuration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/sessionDuration
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SetSessionDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionDuration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/sessionEnergyCharged
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SessionEnergyCharged() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sessionEnergyCharged"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/sessionEnergyCharged
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SetSessionEnergyCharged(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionEnergyCharged:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/sessionID
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SessionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sessionID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/sessionID
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SetSessionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/state
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) State() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("state"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVNotDetectedEvent/state
func (m_ MTREnergyEVSEClusterEVNotDetectedEvent) SetState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setState:"), value)
}



