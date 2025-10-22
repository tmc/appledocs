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
	EnergyExported() MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct
	SetEnergyExported(value IMTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct)
	EnergyImported() MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct
	SetEnergyImported(value IMTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct)
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustercumulativeenergymeasuredevent/energyexported
func (m_ MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent) EnergyExported() MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct](m_.ID, objc.Sel("energyExported"))
	return rv
}


// SetEnergyExported sets the value of the energyExported property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustercumulativeenergymeasuredevent/energyexported
func (m_ MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent) SetEnergyExported(value IMTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnergyExported:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustercumulativeenergymeasuredevent/energyimported
func (m_ MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent) EnergyImported() MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct](m_.ID, objc.Sel("energyImported"))
	return rv
}


// SetEnergyImported sets the value of the energyImported property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustercumulativeenergymeasuredevent/energyimported
func (m_ MTRElectricalEnergyMeasurementClusterCumulativeEnergyMeasuredEvent) SetEnergyImported(value IMTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnergyImported:"), value)
}



