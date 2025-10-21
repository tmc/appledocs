// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTREnergyEVSEClusterRFIDEvent] class.
var (
	MTREnergyEVSEClusterRFIDEventClass     _MTREnergyEVSEClusterRFIDEventClass
	MTREnergyEVSEClusterRFIDEventClassOnce sync.Once
)

func getMTREnergyEVSEClusterRFIDEventClass() _MTREnergyEVSEClusterRFIDEventClass {
	MTREnergyEVSEClusterRFIDEventClassOnce.Do(func() {
		MTREnergyEVSEClusterRFIDEventClass = _MTREnergyEVSEClusterRFIDEventClass{objc.GetClass("MTREnergyEVSEClusterRFIDEvent")}
	})
	return MTREnergyEVSEClusterRFIDEventClass
}

type _MTREnergyEVSEClusterRFIDEventClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterRFIDEvent] class.
type IMTREnergyEVSEClusterRFIDEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterRFIDEvent
type MTREnergyEVSEClusterRFIDEvent struct {
	objectivec.Object
}

// MTREnergyEVSEClusterRFIDEventFrom constructs a [MTREnergyEVSEClusterRFIDEvent] from an unsafe.Pointer.
func MTREnergyEVSEClusterRFIDEventFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterRFIDEvent {
	return MTREnergyEVSEClusterRFIDEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterRFIDEventClass) Alloc() MTREnergyEVSEClusterRFIDEvent {
	rv := objc.Send[MTREnergyEVSEClusterRFIDEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterRFIDEventClass) New() MTREnergyEVSEClusterRFIDEvent {
	rv := objc.Send[MTREnergyEVSEClusterRFIDEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterRFIDEvent) Init() MTREnergyEVSEClusterRFIDEvent {
	rv := objc.Send[MTREnergyEVSEClusterRFIDEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterRFIDEvent) Autorelease() MTREnergyEVSEClusterRFIDEvent {
	rv := objc.Send[MTREnergyEVSEClusterRFIDEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterRFIDEvent creates a new MTREnergyEVSEClusterRFIDEvent instance.
func NewMTREnergyEVSEClusterRFIDEvent() MTREnergyEVSEClusterRFIDEvent {
	return getMTREnergyEVSEClusterRFIDEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterRFIDEvent/uid
func (m_ MTREnergyEVSEClusterRFIDEvent) Uid() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("uid"))
	return rv
}


// SetUid sets the value of the uid property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterRFIDEvent/uid
func (m_ MTREnergyEVSEClusterRFIDEvent) SetUid(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUid:"), value)
}



