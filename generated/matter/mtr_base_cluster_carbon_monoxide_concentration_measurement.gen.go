// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterCarbonMonoxideConcentrationMeasurement] class.
var (
	MTRBaseClusterCarbonMonoxideConcentrationMeasurementClass     _MTRBaseClusterCarbonMonoxideConcentrationMeasurementClass
	MTRBaseClusterCarbonMonoxideConcentrationMeasurementClassOnce sync.Once
)

func getMTRBaseClusterCarbonMonoxideConcentrationMeasurementClass() _MTRBaseClusterCarbonMonoxideConcentrationMeasurementClass {
	MTRBaseClusterCarbonMonoxideConcentrationMeasurementClassOnce.Do(func() {
		MTRBaseClusterCarbonMonoxideConcentrationMeasurementClass = _MTRBaseClusterCarbonMonoxideConcentrationMeasurementClass{objc.GetClass("MTRBaseClusterCarbonMonoxideConcentrationMeasurement")}
	})
	return MTRBaseClusterCarbonMonoxideConcentrationMeasurementClass
}

type _MTRBaseClusterCarbonMonoxideConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterCarbonMonoxideConcentrationMeasurement] class.
type IMTRBaseClusterCarbonMonoxideConcentrationMeasurement interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterCarbonMonoxideConcentrationMeasurement
type MTRBaseClusterCarbonMonoxideConcentrationMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterCarbonMonoxideConcentrationMeasurementFrom constructs a [MTRBaseClusterCarbonMonoxideConcentrationMeasurement] from an unsafe.Pointer.
func MTRBaseClusterCarbonMonoxideConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterCarbonMonoxideConcentrationMeasurement {
	return MTRBaseClusterCarbonMonoxideConcentrationMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterCarbonMonoxideConcentrationMeasurementClass) Alloc() MTRBaseClusterCarbonMonoxideConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterCarbonMonoxideConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterCarbonMonoxideConcentrationMeasurementClass) New() MTRBaseClusterCarbonMonoxideConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterCarbonMonoxideConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterCarbonMonoxideConcentrationMeasurement) Init() MTRBaseClusterCarbonMonoxideConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterCarbonMonoxideConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterCarbonMonoxideConcentrationMeasurement) Autorelease() MTRBaseClusterCarbonMonoxideConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterCarbonMonoxideConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterCarbonMonoxideConcentrationMeasurement creates a new MTRBaseClusterCarbonMonoxideConcentrationMeasurement instance.
func NewMTRBaseClusterCarbonMonoxideConcentrationMeasurement() MTRBaseClusterCarbonMonoxideConcentrationMeasurement {
	return getMTRBaseClusterCarbonMonoxideConcentrationMeasurementClass().New()
}




