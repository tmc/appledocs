// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTREnergyEVSEClusterEVConnectedEvent] class.
var (
	MTREnergyEVSEClusterEVConnectedEventClass     _MTREnergyEVSEClusterEVConnectedEventClass
	MTREnergyEVSEClusterEVConnectedEventClassOnce sync.Once
)

func getMTREnergyEVSEClusterEVConnectedEventClass() _MTREnergyEVSEClusterEVConnectedEventClass {
	MTREnergyEVSEClusterEVConnectedEventClassOnce.Do(func() {
		MTREnergyEVSEClusterEVConnectedEventClass = _MTREnergyEVSEClusterEVConnectedEventClass{objc.GetClass("MTREnergyEVSEClusterEVConnectedEvent")}
	})
	return MTREnergyEVSEClusterEVConnectedEventClass
}

type _MTREnergyEVSEClusterEVConnectedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterEVConnectedEvent] class.
type IMTREnergyEVSEClusterEVConnectedEvent interface {
	objectivec.IObject
	// properties:
	SessionID() objc.IObject /* cross-framework: NSNumber */
	SetSessionID(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVConnectedEvent
type MTREnergyEVSEClusterEVConnectedEvent struct {
	objectivec.Object
}

// MTREnergyEVSEClusterEVConnectedEventFrom constructs a [MTREnergyEVSEClusterEVConnectedEvent] from an unsafe.Pointer.
func MTREnergyEVSEClusterEVConnectedEventFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterEVConnectedEvent {
	return MTREnergyEVSEClusterEVConnectedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterEVConnectedEventClass) Alloc() MTREnergyEVSEClusterEVConnectedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEVConnectedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterEVConnectedEventClass) New() MTREnergyEVSEClusterEVConnectedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEVConnectedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterEVConnectedEvent) Init() MTREnergyEVSEClusterEVConnectedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEVConnectedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterEVConnectedEvent) Autorelease() MTREnergyEVSEClusterEVConnectedEvent {
	rv := objc.Send[MTREnergyEVSEClusterEVConnectedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterEVConnectedEvent creates a new MTREnergyEVSEClusterEVConnectedEvent instance.
func NewMTREnergyEVSEClusterEVConnectedEvent() MTREnergyEVSEClusterEVConnectedEvent {
	return getMTREnergyEVSEClusterEVConnectedEventClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVConnectedEvent/sessionID
func (m_ MTREnergyEVSEClusterEVConnectedEvent) SessionID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("sessionID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterEVConnectedEvent/sessionID
func (m_ MTREnergyEVSEClusterEVConnectedEvent) SetSessionID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSessionID:"), value)
}



