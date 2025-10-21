// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent] class.
var (
	MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEventClass     _MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEventClass
	MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEventClassOnce sync.Once
)

func getMTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEventClass() _MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEventClass {
	MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEventClassOnce.Do(func() {
		MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEventClass = _MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEventClass{objc.GetClass("MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent")}
	})
	return MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEventClass
}

type _MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent] class.
type IMTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent
type MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent struct {
	objectivec.Object
}

// MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEventFrom constructs a [MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent] from an unsafe.Pointer.
func MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEventFrom(ptr unsafe.Pointer) MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent {
	return MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEventClass) Alloc() MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEventClass) New() MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent) Init() MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent) Autorelease() MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent creates a new MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent instance.
func NewMTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent() MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent {
	return getMTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEventClass().New()
}




