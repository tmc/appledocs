// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent] class.
var (
	MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEventClass     _MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEventClass
	MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEventClassOnce sync.Once
)

func getMTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEventClass() _MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEventClass {
	MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEventClassOnce.Do(func() {
		MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEventClass = _MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEventClass{objc.GetClass("MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent")}
	})
	return MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEventClass
}

type _MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent] class.
type IMTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent
type MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent struct {
	objectivec.Object
}

// MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEventFrom constructs a [MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent] from an unsafe.Pointer.
func MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEventFrom(ptr unsafe.Pointer) MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent {
	return MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEventClass) Alloc() MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEventClass) New() MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent) Init() MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent) Autorelease() MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent creates a new MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent instance.
func NewMTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent() MTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEvent {
	return getMTRElectricalPowerMeasurementClusterMeasurementPeriodRangesEventClass().New()
}




