// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBasicClusterStartUpEvent] class.
var (
	MTRBasicClusterStartUpEventClass     _MTRBasicClusterStartUpEventClass
	MTRBasicClusterStartUpEventClassOnce sync.Once
)

func getMTRBasicClusterStartUpEventClass() _MTRBasicClusterStartUpEventClass {
	MTRBasicClusterStartUpEventClassOnce.Do(func() {
		MTRBasicClusterStartUpEventClass = _MTRBasicClusterStartUpEventClass{objc.GetClass("MTRBasicClusterStartUpEvent")}
	})
	return MTRBasicClusterStartUpEventClass
}

type _MTRBasicClusterStartUpEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBasicClusterStartUpEvent] class.
type IMTRBasicClusterStartUpEvent interface {
	IMTRBasicInformationClusterStartUpEvent
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBasicClusterStartUpEvent
type MTRBasicClusterStartUpEvent struct {
	MTRBasicInformationClusterStartUpEvent
}

// MTRBasicClusterStartUpEventFrom constructs a [MTRBasicClusterStartUpEvent] from an unsafe.Pointer.
func MTRBasicClusterStartUpEventFrom(ptr unsafe.Pointer) MTRBasicClusterStartUpEvent {
	return MTRBasicClusterStartUpEvent{
		MTRBasicInformationClusterStartUpEvent: MTRBasicInformationClusterStartUpEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBasicClusterStartUpEventClass) Alloc() MTRBasicClusterStartUpEvent {
	rv := objc.Send[MTRBasicClusterStartUpEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBasicClusterStartUpEventClass) New() MTRBasicClusterStartUpEvent {
	rv := objc.Send[MTRBasicClusterStartUpEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBasicClusterStartUpEvent) Init() MTRBasicClusterStartUpEvent {
	rv := objc.Send[MTRBasicClusterStartUpEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBasicClusterStartUpEvent) Autorelease() MTRBasicClusterStartUpEvent {
	rv := objc.Send[MTRBasicClusterStartUpEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBasicClusterStartUpEvent creates a new MTRBasicClusterStartUpEvent instance.
func NewMTRBasicClusterStartUpEvent() MTRBasicClusterStartUpEvent {
	return getMTRBasicClusterStartUpEventClass().New()
}




