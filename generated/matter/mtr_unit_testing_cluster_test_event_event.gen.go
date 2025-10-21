// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRUnitTestingClusterTestEventEvent] class.
var (
	MTRUnitTestingClusterTestEventEventClass     _MTRUnitTestingClusterTestEventEventClass
	MTRUnitTestingClusterTestEventEventClassOnce sync.Once
)

func getMTRUnitTestingClusterTestEventEventClass() _MTRUnitTestingClusterTestEventEventClass {
	MTRUnitTestingClusterTestEventEventClassOnce.Do(func() {
		MTRUnitTestingClusterTestEventEventClass = _MTRUnitTestingClusterTestEventEventClass{objc.GetClass("MTRUnitTestingClusterTestEventEvent")}
	})
	return MTRUnitTestingClusterTestEventEventClass
}

type _MTRUnitTestingClusterTestEventEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestEventEvent] class.
type IMTRUnitTestingClusterTestEventEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestEventEvent
type MTRUnitTestingClusterTestEventEvent struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestEventEventFrom constructs a [MTRUnitTestingClusterTestEventEvent] from an unsafe.Pointer.
func MTRUnitTestingClusterTestEventEventFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestEventEvent {
	return MTRUnitTestingClusterTestEventEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestEventEventClass) Alloc() MTRUnitTestingClusterTestEventEvent {
	rv := objc.Send[MTRUnitTestingClusterTestEventEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestEventEventClass) New() MTRUnitTestingClusterTestEventEvent {
	rv := objc.Send[MTRUnitTestingClusterTestEventEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestEventEvent) Init() MTRUnitTestingClusterTestEventEvent {
	rv := objc.Send[MTRUnitTestingClusterTestEventEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestEventEvent) Autorelease() MTRUnitTestingClusterTestEventEvent {
	rv := objc.Send[MTRUnitTestingClusterTestEventEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestEventEvent creates a new MTRUnitTestingClusterTestEventEvent instance.
func NewMTRUnitTestingClusterTestEventEvent() MTRUnitTestingClusterTestEventEvent {
	return getMTRUnitTestingClusterTestEventEventClass().New()
}




