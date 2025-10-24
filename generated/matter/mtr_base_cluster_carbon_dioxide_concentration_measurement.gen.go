// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterCarbonDioxideConcentrationMeasurement] class.
var (
	MTRBaseClusterCarbonDioxideConcentrationMeasurementClass     _MTRBaseClusterCarbonDioxideConcentrationMeasurementClass
	MTRBaseClusterCarbonDioxideConcentrationMeasurementClassOnce sync.Once
)

func getMTRBaseClusterCarbonDioxideConcentrationMeasurementClass() _MTRBaseClusterCarbonDioxideConcentrationMeasurementClass {
	MTRBaseClusterCarbonDioxideConcentrationMeasurementClassOnce.Do(func() {
		MTRBaseClusterCarbonDioxideConcentrationMeasurementClass = _MTRBaseClusterCarbonDioxideConcentrationMeasurementClass{objc.GetClass("MTRBaseClusterCarbonDioxideConcentrationMeasurement")}
	})
	return MTRBaseClusterCarbonDioxideConcentrationMeasurementClass
}

type _MTRBaseClusterCarbonDioxideConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterCarbonDioxideConcentrationMeasurement] class.
type IMTRBaseClusterCarbonDioxideConcentrationMeasurement interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterCarbonDioxideConcentrationMeasurement
type MTRBaseClusterCarbonDioxideConcentrationMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterCarbonDioxideConcentrationMeasurementFrom constructs a [MTRBaseClusterCarbonDioxideConcentrationMeasurement] from an unsafe.Pointer.
func MTRBaseClusterCarbonDioxideConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterCarbonDioxideConcentrationMeasurement {
	return MTRBaseClusterCarbonDioxideConcentrationMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterCarbonDioxideConcentrationMeasurementClass) Alloc() MTRBaseClusterCarbonDioxideConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterCarbonDioxideConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterCarbonDioxideConcentrationMeasurementClass) New() MTRBaseClusterCarbonDioxideConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterCarbonDioxideConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterCarbonDioxideConcentrationMeasurement) Init() MTRBaseClusterCarbonDioxideConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterCarbonDioxideConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterCarbonDioxideConcentrationMeasurement) Autorelease() MTRBaseClusterCarbonDioxideConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterCarbonDioxideConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterCarbonDioxideConcentrationMeasurement creates a new MTRBaseClusterCarbonDioxideConcentrationMeasurement instance.
func NewMTRBaseClusterCarbonDioxideConcentrationMeasurement() MTRBaseClusterCarbonDioxideConcentrationMeasurement {
	return getMTRBaseClusterCarbonDioxideConcentrationMeasurementClass().New()
}




