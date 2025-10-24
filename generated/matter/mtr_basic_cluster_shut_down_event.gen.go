// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBasicClusterShutDownEvent] class.
var (
	MTRBasicClusterShutDownEventClass     _MTRBasicClusterShutDownEventClass
	MTRBasicClusterShutDownEventClassOnce sync.Once
)

func getMTRBasicClusterShutDownEventClass() _MTRBasicClusterShutDownEventClass {
	MTRBasicClusterShutDownEventClassOnce.Do(func() {
		MTRBasicClusterShutDownEventClass = _MTRBasicClusterShutDownEventClass{objc.GetClass("MTRBasicClusterShutDownEvent")}
	})
	return MTRBasicClusterShutDownEventClass
}

type _MTRBasicClusterShutDownEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBasicClusterShutDownEvent] class.
type IMTRBasicClusterShutDownEvent interface {
	IMTRBasicInformationClusterShutDownEvent
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBasicClusterShutDownEvent
type MTRBasicClusterShutDownEvent struct {
	MTRBasicInformationClusterShutDownEvent
}

// MTRBasicClusterShutDownEventFrom constructs a [MTRBasicClusterShutDownEvent] from an unsafe.Pointer.
func MTRBasicClusterShutDownEventFrom(ptr unsafe.Pointer) MTRBasicClusterShutDownEvent {
	return MTRBasicClusterShutDownEvent{
		MTRBasicInformationClusterShutDownEvent: MTRBasicInformationClusterShutDownEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBasicClusterShutDownEventClass) Alloc() MTRBasicClusterShutDownEvent {
	rv := objc.Send[MTRBasicClusterShutDownEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBasicClusterShutDownEventClass) New() MTRBasicClusterShutDownEvent {
	rv := objc.Send[MTRBasicClusterShutDownEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBasicClusterShutDownEvent) Init() MTRBasicClusterShutDownEvent {
	rv := objc.Send[MTRBasicClusterShutDownEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBasicClusterShutDownEvent) Autorelease() MTRBasicClusterShutDownEvent {
	rv := objc.Send[MTRBasicClusterShutDownEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBasicClusterShutDownEvent creates a new MTRBasicClusterShutDownEvent instance.
func NewMTRBasicClusterShutDownEvent() MTRBasicClusterShutDownEvent {
	return getMTRBasicClusterShutDownEventClass().New()
}




