// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestEventEvent] class.
var (
	MTRTestClusterClusterTestEventEventClass     _MTRTestClusterClusterTestEventEventClass
	MTRTestClusterClusterTestEventEventClassOnce sync.Once
)

func getMTRTestClusterClusterTestEventEventClass() _MTRTestClusterClusterTestEventEventClass {
	MTRTestClusterClusterTestEventEventClassOnce.Do(func() {
		MTRTestClusterClusterTestEventEventClass = _MTRTestClusterClusterTestEventEventClass{objc.GetClass("MTRTestClusterClusterTestEventEvent")}
	})
	return MTRTestClusterClusterTestEventEventClass
}

type _MTRTestClusterClusterTestEventEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestEventEvent] class.
type IMTRTestClusterClusterTestEventEvent interface {
	IMTRUnitTestingClusterTestEventEvent
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestEventEvent
type MTRTestClusterClusterTestEventEvent struct {
	MTRUnitTestingClusterTestEventEvent
}

// MTRTestClusterClusterTestEventEventFrom constructs a [MTRTestClusterClusterTestEventEvent] from an unsafe.Pointer.
func MTRTestClusterClusterTestEventEventFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestEventEvent {
	return MTRTestClusterClusterTestEventEvent{
		MTRUnitTestingClusterTestEventEvent: MTRUnitTestingClusterTestEventEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestEventEventClass) Alloc() MTRTestClusterClusterTestEventEvent {
	rv := objc.Send[MTRTestClusterClusterTestEventEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestEventEventClass) New() MTRTestClusterClusterTestEventEvent {
	rv := objc.Send[MTRTestClusterClusterTestEventEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestEventEvent) Init() MTRTestClusterClusterTestEventEvent {
	rv := objc.Send[MTRTestClusterClusterTestEventEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestEventEvent) Autorelease() MTRTestClusterClusterTestEventEvent {
	rv := objc.Send[MTRTestClusterClusterTestEventEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestEventEvent creates a new MTRTestClusterClusterTestEventEvent instance.
func NewMTRTestClusterClusterTestEventEvent() MTRTestClusterClusterTestEventEvent {
	return getMTRTestClusterClusterTestEventEventClass().New()
}




