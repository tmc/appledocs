// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestFabricScopedEventEvent] class.
var (
	MTRTestClusterClusterTestFabricScopedEventEventClass     _MTRTestClusterClusterTestFabricScopedEventEventClass
	MTRTestClusterClusterTestFabricScopedEventEventClassOnce sync.Once
)

func getMTRTestClusterClusterTestFabricScopedEventEventClass() _MTRTestClusterClusterTestFabricScopedEventEventClass {
	MTRTestClusterClusterTestFabricScopedEventEventClassOnce.Do(func() {
		MTRTestClusterClusterTestFabricScopedEventEventClass = _MTRTestClusterClusterTestFabricScopedEventEventClass{objc.GetClass("MTRTestClusterClusterTestFabricScopedEventEvent")}
	})
	return MTRTestClusterClusterTestFabricScopedEventEventClass
}

type _MTRTestClusterClusterTestFabricScopedEventEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestFabricScopedEventEvent] class.
type IMTRTestClusterClusterTestFabricScopedEventEvent interface {
	IMTRUnitTestingClusterTestFabricScopedEventEvent
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestFabricScopedEventEvent
type MTRTestClusterClusterTestFabricScopedEventEvent struct {
	MTRUnitTestingClusterTestFabricScopedEventEvent
}

// MTRTestClusterClusterTestFabricScopedEventEventFrom constructs a [MTRTestClusterClusterTestFabricScopedEventEvent] from an unsafe.Pointer.
func MTRTestClusterClusterTestFabricScopedEventEventFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestFabricScopedEventEvent {
	return MTRTestClusterClusterTestFabricScopedEventEvent{
		MTRUnitTestingClusterTestFabricScopedEventEvent: MTRUnitTestingClusterTestFabricScopedEventEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestFabricScopedEventEventClass) Alloc() MTRTestClusterClusterTestFabricScopedEventEvent {
	rv := objc.Send[MTRTestClusterClusterTestFabricScopedEventEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestFabricScopedEventEventClass) New() MTRTestClusterClusterTestFabricScopedEventEvent {
	rv := objc.Send[MTRTestClusterClusterTestFabricScopedEventEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestFabricScopedEventEvent) Init() MTRTestClusterClusterTestFabricScopedEventEvent {
	rv := objc.Send[MTRTestClusterClusterTestFabricScopedEventEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestFabricScopedEventEvent) Autorelease() MTRTestClusterClusterTestFabricScopedEventEvent {
	rv := objc.Send[MTRTestClusterClusterTestFabricScopedEventEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestFabricScopedEventEvent creates a new MTRTestClusterClusterTestFabricScopedEventEvent instance.
func NewMTRTestClusterClusterTestFabricScopedEventEvent() MTRTestClusterClusterTestFabricScopedEventEvent {
	return getMTRTestClusterClusterTestFabricScopedEventEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscopedeventevent/fabricindex
func (m_ MTRTestClusterClusterTestFabricScopedEventEvent) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscopedeventevent/fabricindex
func (m_ MTRTestClusterClusterTestFabricScopedEventEvent) SetFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}



