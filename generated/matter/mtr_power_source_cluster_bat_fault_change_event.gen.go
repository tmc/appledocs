// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPowerSourceClusterBatFaultChangeEvent] class.
var (
	MTRPowerSourceClusterBatFaultChangeEventClass     _MTRPowerSourceClusterBatFaultChangeEventClass
	MTRPowerSourceClusterBatFaultChangeEventClassOnce sync.Once
)

func getMTRPowerSourceClusterBatFaultChangeEventClass() _MTRPowerSourceClusterBatFaultChangeEventClass {
	MTRPowerSourceClusterBatFaultChangeEventClassOnce.Do(func() {
		MTRPowerSourceClusterBatFaultChangeEventClass = _MTRPowerSourceClusterBatFaultChangeEventClass{objc.GetClass("MTRPowerSourceClusterBatFaultChangeEvent")}
	})
	return MTRPowerSourceClusterBatFaultChangeEventClass
}

type _MTRPowerSourceClusterBatFaultChangeEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPowerSourceClusterBatFaultChangeEvent] class.
type IMTRPowerSourceClusterBatFaultChangeEvent interface {
	objectivec.IObject
	Current() unsafe.Pointer
	SetCurrent(value unsafe.Pointer)
	Previous() unsafe.Pointer
	SetPrevious(value unsafe.Pointer)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatFaultChangeEvent
type MTRPowerSourceClusterBatFaultChangeEvent struct {
	objectivec.Object
}

// MTRPowerSourceClusterBatFaultChangeEventFrom constructs a [MTRPowerSourceClusterBatFaultChangeEvent] from an unsafe.Pointer.
func MTRPowerSourceClusterBatFaultChangeEventFrom(ptr unsafe.Pointer) MTRPowerSourceClusterBatFaultChangeEvent {
	return MTRPowerSourceClusterBatFaultChangeEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPowerSourceClusterBatFaultChangeEventClass) Alloc() MTRPowerSourceClusterBatFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterBatFaultChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPowerSourceClusterBatFaultChangeEventClass) New() MTRPowerSourceClusterBatFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterBatFaultChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPowerSourceClusterBatFaultChangeEvent) Init() MTRPowerSourceClusterBatFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterBatFaultChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPowerSourceClusterBatFaultChangeEvent) Autorelease() MTRPowerSourceClusterBatFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterBatFaultChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPowerSourceClusterBatFaultChangeEvent creates a new MTRPowerSourceClusterBatFaultChangeEvent instance.
func NewMTRPowerSourceClusterBatFaultChangeEvent() MTRPowerSourceClusterBatFaultChangeEvent {
	return getMTRPowerSourceClusterBatFaultChangeEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrpowersourceclusterbatfaultchangeevent/current
func (m_ MTRPowerSourceClusterBatFaultChangeEvent) Current() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("current"))
	return rv
}


// SetCurrent sets the value of the current property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrpowersourceclusterbatfaultchangeevent/current
func (m_ MTRPowerSourceClusterBatFaultChangeEvent) SetCurrent(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrpowersourceclusterbatfaultchangeevent/previous
func (m_ MTRPowerSourceClusterBatFaultChangeEvent) Previous() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("previous"))
	return rv
}


// SetPrevious sets the value of the previous property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrpowersourceclusterbatfaultchangeevent/previous
func (m_ MTRPowerSourceClusterBatFaultChangeEvent) SetPrevious(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrevious:"), value)
}



