// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRUnitTestingClusterTestFabricScopedEventEvent] class.
var (
	MTRUnitTestingClusterTestFabricScopedEventEventClass     _MTRUnitTestingClusterTestFabricScopedEventEventClass
	MTRUnitTestingClusterTestFabricScopedEventEventClassOnce sync.Once
)

func getMTRUnitTestingClusterTestFabricScopedEventEventClass() _MTRUnitTestingClusterTestFabricScopedEventEventClass {
	MTRUnitTestingClusterTestFabricScopedEventEventClassOnce.Do(func() {
		MTRUnitTestingClusterTestFabricScopedEventEventClass = _MTRUnitTestingClusterTestFabricScopedEventEventClass{objc.GetClass("MTRUnitTestingClusterTestFabricScopedEventEvent")}
	})
	return MTRUnitTestingClusterTestFabricScopedEventEventClass
}

type _MTRUnitTestingClusterTestFabricScopedEventEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestFabricScopedEventEvent] class.
type IMTRUnitTestingClusterTestFabricScopedEventEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScopedEventEvent
type MTRUnitTestingClusterTestFabricScopedEventEvent struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestFabricScopedEventEventFrom constructs a [MTRUnitTestingClusterTestFabricScopedEventEvent] from an unsafe.Pointer.
func MTRUnitTestingClusterTestFabricScopedEventEventFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestFabricScopedEventEvent {
	return MTRUnitTestingClusterTestFabricScopedEventEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestFabricScopedEventEventClass) Alloc() MTRUnitTestingClusterTestFabricScopedEventEvent {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScopedEventEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestFabricScopedEventEventClass) New() MTRUnitTestingClusterTestFabricScopedEventEvent {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScopedEventEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestFabricScopedEventEvent) Init() MTRUnitTestingClusterTestFabricScopedEventEvent {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScopedEventEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestFabricScopedEventEvent) Autorelease() MTRUnitTestingClusterTestFabricScopedEventEvent {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScopedEventEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestFabricScopedEventEvent creates a new MTRUnitTestingClusterTestFabricScopedEventEvent instance.
func NewMTRUnitTestingClusterTestFabricScopedEventEvent() MTRUnitTestingClusterTestFabricScopedEventEvent {
	return getMTRUnitTestingClusterTestFabricScopedEventEventClass().New()
}




