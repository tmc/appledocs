// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterElectricalPowerMeasurement] class.
var (
	MTRBaseClusterElectricalPowerMeasurementClass     _MTRBaseClusterElectricalPowerMeasurementClass
	MTRBaseClusterElectricalPowerMeasurementClassOnce sync.Once
)

func getMTRBaseClusterElectricalPowerMeasurementClass() _MTRBaseClusterElectricalPowerMeasurementClass {
	MTRBaseClusterElectricalPowerMeasurementClassOnce.Do(func() {
		MTRBaseClusterElectricalPowerMeasurementClass = _MTRBaseClusterElectricalPowerMeasurementClass{objc.GetClass("MTRBaseClusterElectricalPowerMeasurement")}
	})
	return MTRBaseClusterElectricalPowerMeasurementClass
}

type _MTRBaseClusterElectricalPowerMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterElectricalPowerMeasurement] class.
type IMTRBaseClusterElectricalPowerMeasurement interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterElectricalPowerMeasurement
type MTRBaseClusterElectricalPowerMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterElectricalPowerMeasurementFrom constructs a [MTRBaseClusterElectricalPowerMeasurement] from an unsafe.Pointer.
func MTRBaseClusterElectricalPowerMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterElectricalPowerMeasurement {
	return MTRBaseClusterElectricalPowerMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterElectricalPowerMeasurementClass) Alloc() MTRBaseClusterElectricalPowerMeasurement {
	rv := objc.Send[MTRBaseClusterElectricalPowerMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterElectricalPowerMeasurementClass) New() MTRBaseClusterElectricalPowerMeasurement {
	rv := objc.Send[MTRBaseClusterElectricalPowerMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterElectricalPowerMeasurement) Init() MTRBaseClusterElectricalPowerMeasurement {
	rv := objc.Send[MTRBaseClusterElectricalPowerMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterElectricalPowerMeasurement) Autorelease() MTRBaseClusterElectricalPowerMeasurement {
	rv := objc.Send[MTRBaseClusterElectricalPowerMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterElectricalPowerMeasurement creates a new MTRBaseClusterElectricalPowerMeasurement instance.
func NewMTRBaseClusterElectricalPowerMeasurement() MTRBaseClusterElectricalPowerMeasurement {
	return getMTRBaseClusterElectricalPowerMeasurementClass().New()
}




