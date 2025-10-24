// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterElectricalPowerMeasurement] class.
var (
	MTRClusterElectricalPowerMeasurementClass     _MTRClusterElectricalPowerMeasurementClass
	MTRClusterElectricalPowerMeasurementClassOnce sync.Once
)

func getMTRClusterElectricalPowerMeasurementClass() _MTRClusterElectricalPowerMeasurementClass {
	MTRClusterElectricalPowerMeasurementClassOnce.Do(func() {
		MTRClusterElectricalPowerMeasurementClass = _MTRClusterElectricalPowerMeasurementClass{objc.GetClass("MTRClusterElectricalPowerMeasurement")}
	})
	return MTRClusterElectricalPowerMeasurementClass
}

type _MTRClusterElectricalPowerMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterElectricalPowerMeasurement] class.
type IMTRClusterElectricalPowerMeasurement interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterElectricalPowerMeasurement
type MTRClusterElectricalPowerMeasurement struct {
	MTRGenericCluster
}

// MTRClusterElectricalPowerMeasurementFrom constructs a [MTRClusterElectricalPowerMeasurement] from an unsafe.Pointer.
func MTRClusterElectricalPowerMeasurementFrom(ptr unsafe.Pointer) MTRClusterElectricalPowerMeasurement {
	return MTRClusterElectricalPowerMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterElectricalPowerMeasurementClass) Alloc() MTRClusterElectricalPowerMeasurement {
	rv := objc.Send[MTRClusterElectricalPowerMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterElectricalPowerMeasurementClass) New() MTRClusterElectricalPowerMeasurement {
	rv := objc.Send[MTRClusterElectricalPowerMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterElectricalPowerMeasurement) Init() MTRClusterElectricalPowerMeasurement {
	rv := objc.Send[MTRClusterElectricalPowerMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterElectricalPowerMeasurement) Autorelease() MTRClusterElectricalPowerMeasurement {
	rv := objc.Send[MTRClusterElectricalPowerMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterElectricalPowerMeasurement creates a new MTRClusterElectricalPowerMeasurement instance.
func NewMTRClusterElectricalPowerMeasurement() MTRClusterElectricalPowerMeasurement {
	return getMTRClusterElectricalPowerMeasurementClass().New()
}




