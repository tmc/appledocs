// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPowerSourceClusterWiredFaultChangeEvent] class.
var (
	MTRPowerSourceClusterWiredFaultChangeEventClass     _MTRPowerSourceClusterWiredFaultChangeEventClass
	MTRPowerSourceClusterWiredFaultChangeEventClassOnce sync.Once
)

func getMTRPowerSourceClusterWiredFaultChangeEventClass() _MTRPowerSourceClusterWiredFaultChangeEventClass {
	MTRPowerSourceClusterWiredFaultChangeEventClassOnce.Do(func() {
		MTRPowerSourceClusterWiredFaultChangeEventClass = _MTRPowerSourceClusterWiredFaultChangeEventClass{objc.GetClass("MTRPowerSourceClusterWiredFaultChangeEvent")}
	})
	return MTRPowerSourceClusterWiredFaultChangeEventClass
}

type _MTRPowerSourceClusterWiredFaultChangeEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPowerSourceClusterWiredFaultChangeEvent] class.
type IMTRPowerSourceClusterWiredFaultChangeEvent interface {
	objectivec.IObject
	Current() unsafe.Pointer
	SetCurrent(value unsafe.Pointer)
	Previous() unsafe.Pointer
	SetPrevious(value unsafe.Pointer)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterWiredFaultChangeEvent
type MTRPowerSourceClusterWiredFaultChangeEvent struct {
	objectivec.Object
}

// MTRPowerSourceClusterWiredFaultChangeEventFrom constructs a [MTRPowerSourceClusterWiredFaultChangeEvent] from an unsafe.Pointer.
func MTRPowerSourceClusterWiredFaultChangeEventFrom(ptr unsafe.Pointer) MTRPowerSourceClusterWiredFaultChangeEvent {
	return MTRPowerSourceClusterWiredFaultChangeEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPowerSourceClusterWiredFaultChangeEventClass) Alloc() MTRPowerSourceClusterWiredFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterWiredFaultChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPowerSourceClusterWiredFaultChangeEventClass) New() MTRPowerSourceClusterWiredFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterWiredFaultChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPowerSourceClusterWiredFaultChangeEvent) Init() MTRPowerSourceClusterWiredFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterWiredFaultChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPowerSourceClusterWiredFaultChangeEvent) Autorelease() MTRPowerSourceClusterWiredFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterWiredFaultChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPowerSourceClusterWiredFaultChangeEvent creates a new MTRPowerSourceClusterWiredFaultChangeEvent instance.
func NewMTRPowerSourceClusterWiredFaultChangeEvent() MTRPowerSourceClusterWiredFaultChangeEvent {
	return getMTRPowerSourceClusterWiredFaultChangeEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrpowersourceclusterwiredfaultchangeevent/current
func (m_ MTRPowerSourceClusterWiredFaultChangeEvent) Current() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("current"))
	return rv
}


// SetCurrent sets the value of the current property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrpowersourceclusterwiredfaultchangeevent/current
func (m_ MTRPowerSourceClusterWiredFaultChangeEvent) SetCurrent(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrpowersourceclusterwiredfaultchangeevent/previous
func (m_ MTRPowerSourceClusterWiredFaultChangeEvent) Previous() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("previous"))
	return rv
}


// SetPrevious sets the value of the previous property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrpowersourceclusterwiredfaultchangeevent/previous
func (m_ MTRPowerSourceClusterWiredFaultChangeEvent) SetPrevious(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrevious:"), value)
}



