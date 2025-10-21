// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent] class.
var (
	MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEventClass     _MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEventClass
	MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEventClassOnce sync.Once
)

func getMTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEventClass() _MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEventClass {
	MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEventClassOnce.Do(func() {
		MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEventClass = _MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEventClass{objc.GetClass("MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent")}
	})
	return MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEventClass
}

type _MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent] class.
type IMTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent
type MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent struct {
	objectivec.Object
}

// MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEventFrom constructs a [MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent] from an unsafe.Pointer.
func MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEventFrom(ptr unsafe.Pointer) MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent {
	return MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEventClass) Alloc() MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEventClass) New() MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent) Init() MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent) Autorelease() MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent creates a new MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent instance.
func NewMTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent() MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent {
	return getMTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterperiodicenergymeasuredevent/energyexported
func (m_ MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent) EnergyExported() MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct](m_.ID, objc.Sel("energyExported"))
	return rv
}


// SetEnergyExported sets the value of the energyExported property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterperiodicenergymeasuredevent/energyexported
func (m_ MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent) SetEnergyExported(value IMTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnergyExported:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterperiodicenergymeasuredevent/energyimported
func (m_ MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent) EnergyImported() MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct](m_.ID, objc.Sel("energyImported"))
	return rv
}


// SetEnergyImported sets the value of the energyImported property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterperiodicenergymeasuredevent/energyimported
func (m_ MTRElectricalEnergyMeasurementClusterPeriodicEnergyMeasuredEvent) SetEnergyImported(value IMTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnergyImported:"), value)
}



