// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRSwitchClusterInitialPressEvent] class.
var (
	MTRSwitchClusterInitialPressEventClass     _MTRSwitchClusterInitialPressEventClass
	MTRSwitchClusterInitialPressEventClassOnce sync.Once
)

func getMTRSwitchClusterInitialPressEventClass() _MTRSwitchClusterInitialPressEventClass {
	MTRSwitchClusterInitialPressEventClassOnce.Do(func() {
		MTRSwitchClusterInitialPressEventClass = _MTRSwitchClusterInitialPressEventClass{objc.GetClass("MTRSwitchClusterInitialPressEvent")}
	})
	return MTRSwitchClusterInitialPressEventClass
}

type _MTRSwitchClusterInitialPressEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSwitchClusterInitialPressEvent] class.
type IMTRSwitchClusterInitialPressEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterInitialPressEvent
type MTRSwitchClusterInitialPressEvent struct {
	objectivec.Object
}

// MTRSwitchClusterInitialPressEventFrom constructs a [MTRSwitchClusterInitialPressEvent] from an unsafe.Pointer.
func MTRSwitchClusterInitialPressEventFrom(ptr unsafe.Pointer) MTRSwitchClusterInitialPressEvent {
	return MTRSwitchClusterInitialPressEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSwitchClusterInitialPressEventClass) Alloc() MTRSwitchClusterInitialPressEvent {
	rv := objc.Send[MTRSwitchClusterInitialPressEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSwitchClusterInitialPressEventClass) New() MTRSwitchClusterInitialPressEvent {
	rv := objc.Send[MTRSwitchClusterInitialPressEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSwitchClusterInitialPressEvent) Init() MTRSwitchClusterInitialPressEvent {
	rv := objc.Send[MTRSwitchClusterInitialPressEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSwitchClusterInitialPressEvent) Autorelease() MTRSwitchClusterInitialPressEvent {
	rv := objc.Send[MTRSwitchClusterInitialPressEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSwitchClusterInitialPressEvent creates a new MTRSwitchClusterInitialPressEvent instance.
func NewMTRSwitchClusterInitialPressEvent() MTRSwitchClusterInitialPressEvent {
	return getMTRSwitchClusterInitialPressEventClass().New()
}




