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




