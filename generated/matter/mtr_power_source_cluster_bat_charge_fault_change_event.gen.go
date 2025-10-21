// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPowerSourceClusterBatChargeFaultChangeEvent] class.
var (
	MTRPowerSourceClusterBatChargeFaultChangeEventClass     _MTRPowerSourceClusterBatChargeFaultChangeEventClass
	MTRPowerSourceClusterBatChargeFaultChangeEventClassOnce sync.Once
)

func getMTRPowerSourceClusterBatChargeFaultChangeEventClass() _MTRPowerSourceClusterBatChargeFaultChangeEventClass {
	MTRPowerSourceClusterBatChargeFaultChangeEventClassOnce.Do(func() {
		MTRPowerSourceClusterBatChargeFaultChangeEventClass = _MTRPowerSourceClusterBatChargeFaultChangeEventClass{objc.GetClass("MTRPowerSourceClusterBatChargeFaultChangeEvent")}
	})
	return MTRPowerSourceClusterBatChargeFaultChangeEventClass
}

type _MTRPowerSourceClusterBatChargeFaultChangeEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRPowerSourceClusterBatChargeFaultChangeEvent] class.
type IMTRPowerSourceClusterBatChargeFaultChangeEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatChargeFaultChangeEvent
type MTRPowerSourceClusterBatChargeFaultChangeEvent struct {
	objectivec.Object
}

// MTRPowerSourceClusterBatChargeFaultChangeEventFrom constructs a [MTRPowerSourceClusterBatChargeFaultChangeEvent] from an unsafe.Pointer.
func MTRPowerSourceClusterBatChargeFaultChangeEventFrom(ptr unsafe.Pointer) MTRPowerSourceClusterBatChargeFaultChangeEvent {
	return MTRPowerSourceClusterBatChargeFaultChangeEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPowerSourceClusterBatChargeFaultChangeEventClass) Alloc() MTRPowerSourceClusterBatChargeFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterBatChargeFaultChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPowerSourceClusterBatChargeFaultChangeEventClass) New() MTRPowerSourceClusterBatChargeFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterBatChargeFaultChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPowerSourceClusterBatChargeFaultChangeEvent) Init() MTRPowerSourceClusterBatChargeFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterBatChargeFaultChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPowerSourceClusterBatChargeFaultChangeEvent) Autorelease() MTRPowerSourceClusterBatChargeFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterBatChargeFaultChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPowerSourceClusterBatChargeFaultChangeEvent creates a new MTRPowerSourceClusterBatChargeFaultChangeEvent instance.
func NewMTRPowerSourceClusterBatChargeFaultChangeEvent() MTRPowerSourceClusterBatChargeFaultChangeEvent {
	return getMTRPowerSourceClusterBatChargeFaultChangeEventClass().New()
}




