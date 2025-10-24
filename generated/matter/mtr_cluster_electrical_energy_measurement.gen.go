// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterElectricalEnergyMeasurement] class.
var (
	MTRClusterElectricalEnergyMeasurementClass     _MTRClusterElectricalEnergyMeasurementClass
	MTRClusterElectricalEnergyMeasurementClassOnce sync.Once
)

func getMTRClusterElectricalEnergyMeasurementClass() _MTRClusterElectricalEnergyMeasurementClass {
	MTRClusterElectricalEnergyMeasurementClassOnce.Do(func() {
		MTRClusterElectricalEnergyMeasurementClass = _MTRClusterElectricalEnergyMeasurementClass{objc.GetClass("MTRClusterElectricalEnergyMeasurement")}
	})
	return MTRClusterElectricalEnergyMeasurementClass
}

type _MTRClusterElectricalEnergyMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterElectricalEnergyMeasurement] class.
type IMTRClusterElectricalEnergyMeasurement interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterElectricalEnergyMeasurement
type MTRClusterElectricalEnergyMeasurement struct {
	MTRGenericCluster
}

// MTRClusterElectricalEnergyMeasurementFrom constructs a [MTRClusterElectricalEnergyMeasurement] from an unsafe.Pointer.
func MTRClusterElectricalEnergyMeasurementFrom(ptr unsafe.Pointer) MTRClusterElectricalEnergyMeasurement {
	return MTRClusterElectricalEnergyMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterElectricalEnergyMeasurementClass) Alloc() MTRClusterElectricalEnergyMeasurement {
	rv := objc.Send[MTRClusterElectricalEnergyMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterElectricalEnergyMeasurementClass) New() MTRClusterElectricalEnergyMeasurement {
	rv := objc.Send[MTRClusterElectricalEnergyMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterElectricalEnergyMeasurement) Init() MTRClusterElectricalEnergyMeasurement {
	rv := objc.Send[MTRClusterElectricalEnergyMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterElectricalEnergyMeasurement) Autorelease() MTRClusterElectricalEnergyMeasurement {
	rv := objc.Send[MTRClusterElectricalEnergyMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterElectricalEnergyMeasurement creates a new MTRClusterElectricalEnergyMeasurement instance.
func NewMTRClusterElectricalEnergyMeasurement() MTRClusterElectricalEnergyMeasurement {
	return getMTRClusterElectricalEnergyMeasurementClass().New()
}




