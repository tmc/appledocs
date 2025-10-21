// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSwitchClusterLongReleaseEvent] class.
var (
	MTRSwitchClusterLongReleaseEventClass     _MTRSwitchClusterLongReleaseEventClass
	MTRSwitchClusterLongReleaseEventClassOnce sync.Once
)

func getMTRSwitchClusterLongReleaseEventClass() _MTRSwitchClusterLongReleaseEventClass {
	MTRSwitchClusterLongReleaseEventClassOnce.Do(func() {
		MTRSwitchClusterLongReleaseEventClass = _MTRSwitchClusterLongReleaseEventClass{objc.GetClass("MTRSwitchClusterLongReleaseEvent")}
	})
	return MTRSwitchClusterLongReleaseEventClass
}

type _MTRSwitchClusterLongReleaseEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSwitchClusterLongReleaseEvent] class.
type IMTRSwitchClusterLongReleaseEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterLongReleaseEvent
type MTRSwitchClusterLongReleaseEvent struct {
	objectivec.Object
}

// MTRSwitchClusterLongReleaseEventFrom constructs a [MTRSwitchClusterLongReleaseEvent] from an unsafe.Pointer.
func MTRSwitchClusterLongReleaseEventFrom(ptr unsafe.Pointer) MTRSwitchClusterLongReleaseEvent {
	return MTRSwitchClusterLongReleaseEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSwitchClusterLongReleaseEventClass) Alloc() MTRSwitchClusterLongReleaseEvent {
	rv := objc.Send[MTRSwitchClusterLongReleaseEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSwitchClusterLongReleaseEventClass) New() MTRSwitchClusterLongReleaseEvent {
	rv := objc.Send[MTRSwitchClusterLongReleaseEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSwitchClusterLongReleaseEvent) Init() MTRSwitchClusterLongReleaseEvent {
	rv := objc.Send[MTRSwitchClusterLongReleaseEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSwitchClusterLongReleaseEvent) Autorelease() MTRSwitchClusterLongReleaseEvent {
	rv := objc.Send[MTRSwitchClusterLongReleaseEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSwitchClusterLongReleaseEvent creates a new MTRSwitchClusterLongReleaseEvent instance.
func NewMTRSwitchClusterLongReleaseEvent() MTRSwitchClusterLongReleaseEvent {
	return getMTRSwitchClusterLongReleaseEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrswitchclusterlongreleaseevent/previousposition
func (m_ MTRSwitchClusterLongReleaseEvent) PreviousPosition() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("previousPosition"))
	return rv
}


// SetPreviousPosition sets the value of the previousPosition property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrswitchclusterlongreleaseevent/previousposition
func (m_ MTRSwitchClusterLongReleaseEvent) SetPreviousPosition(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreviousPosition:"), value)
}



