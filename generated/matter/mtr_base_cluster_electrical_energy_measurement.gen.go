// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterElectricalEnergyMeasurement] class.
var (
	MTRBaseClusterElectricalEnergyMeasurementClass     _MTRBaseClusterElectricalEnergyMeasurementClass
	MTRBaseClusterElectricalEnergyMeasurementClassOnce sync.Once
)

func getMTRBaseClusterElectricalEnergyMeasurementClass() _MTRBaseClusterElectricalEnergyMeasurementClass {
	MTRBaseClusterElectricalEnergyMeasurementClassOnce.Do(func() {
		MTRBaseClusterElectricalEnergyMeasurementClass = _MTRBaseClusterElectricalEnergyMeasurementClass{objc.GetClass("MTRBaseClusterElectricalEnergyMeasurement")}
	})
	return MTRBaseClusterElectricalEnergyMeasurementClass
}

type _MTRBaseClusterElectricalEnergyMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterElectricalEnergyMeasurement] class.
type IMTRBaseClusterElectricalEnergyMeasurement interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterElectricalEnergyMeasurement
type MTRBaseClusterElectricalEnergyMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterElectricalEnergyMeasurementFrom constructs a [MTRBaseClusterElectricalEnergyMeasurement] from an unsafe.Pointer.
func MTRBaseClusterElectricalEnergyMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterElectricalEnergyMeasurement {
	return MTRBaseClusterElectricalEnergyMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterElectricalEnergyMeasurementClass) Alloc() MTRBaseClusterElectricalEnergyMeasurement {
	rv := objc.Send[MTRBaseClusterElectricalEnergyMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterElectricalEnergyMeasurementClass) New() MTRBaseClusterElectricalEnergyMeasurement {
	rv := objc.Send[MTRBaseClusterElectricalEnergyMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterElectricalEnergyMeasurement) Init() MTRBaseClusterElectricalEnergyMeasurement {
	rv := objc.Send[MTRBaseClusterElectricalEnergyMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterElectricalEnergyMeasurement) Autorelease() MTRBaseClusterElectricalEnergyMeasurement {
	rv := objc.Send[MTRBaseClusterElectricalEnergyMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterElectricalEnergyMeasurement creates a new MTRBaseClusterElectricalEnergyMeasurement instance.
func NewMTRBaseClusterElectricalEnergyMeasurement() MTRBaseClusterElectricalEnergyMeasurement {
	return getMTRBaseClusterElectricalEnergyMeasurementClass().New()
}




