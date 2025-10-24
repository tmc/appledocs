// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterCarbonDioxideConcentrationMeasurement] class.
var (
	MTRClusterCarbonDioxideConcentrationMeasurementClass     _MTRClusterCarbonDioxideConcentrationMeasurementClass
	MTRClusterCarbonDioxideConcentrationMeasurementClassOnce sync.Once
)

func getMTRClusterCarbonDioxideConcentrationMeasurementClass() _MTRClusterCarbonDioxideConcentrationMeasurementClass {
	MTRClusterCarbonDioxideConcentrationMeasurementClassOnce.Do(func() {
		MTRClusterCarbonDioxideConcentrationMeasurementClass = _MTRClusterCarbonDioxideConcentrationMeasurementClass{objc.GetClass("MTRClusterCarbonDioxideConcentrationMeasurement")}
	})
	return MTRClusterCarbonDioxideConcentrationMeasurementClass
}

type _MTRClusterCarbonDioxideConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterCarbonDioxideConcentrationMeasurement] class.
type IMTRClusterCarbonDioxideConcentrationMeasurement interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterCarbonDioxideConcentrationMeasurement
type MTRClusterCarbonDioxideConcentrationMeasurement struct {
	MTRGenericCluster
}

// MTRClusterCarbonDioxideConcentrationMeasurementFrom constructs a [MTRClusterCarbonDioxideConcentrationMeasurement] from an unsafe.Pointer.
func MTRClusterCarbonDioxideConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRClusterCarbonDioxideConcentrationMeasurement {
	return MTRClusterCarbonDioxideConcentrationMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterCarbonDioxideConcentrationMeasurementClass) Alloc() MTRClusterCarbonDioxideConcentrationMeasurement {
	rv := objc.Send[MTRClusterCarbonDioxideConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterCarbonDioxideConcentrationMeasurementClass) New() MTRClusterCarbonDioxideConcentrationMeasurement {
	rv := objc.Send[MTRClusterCarbonDioxideConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterCarbonDioxideConcentrationMeasurement) Init() MTRClusterCarbonDioxideConcentrationMeasurement {
	rv := objc.Send[MTRClusterCarbonDioxideConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterCarbonDioxideConcentrationMeasurement) Autorelease() MTRClusterCarbonDioxideConcentrationMeasurement {
	rv := objc.Send[MTRClusterCarbonDioxideConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterCarbonDioxideConcentrationMeasurement creates a new MTRClusterCarbonDioxideConcentrationMeasurement instance.
func NewMTRClusterCarbonDioxideConcentrationMeasurement() MTRClusterCarbonDioxideConcentrationMeasurement {
	return getMTRClusterCarbonDioxideConcentrationMeasurementClass().New()
}
