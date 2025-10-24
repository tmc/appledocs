// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterCarbonMonoxideConcentrationMeasurement] class.
var (
	MTRClusterCarbonMonoxideConcentrationMeasurementClass     _MTRClusterCarbonMonoxideConcentrationMeasurementClass
	MTRClusterCarbonMonoxideConcentrationMeasurementClassOnce sync.Once
)

func getMTRClusterCarbonMonoxideConcentrationMeasurementClass() _MTRClusterCarbonMonoxideConcentrationMeasurementClass {
	MTRClusterCarbonMonoxideConcentrationMeasurementClassOnce.Do(func() {
		MTRClusterCarbonMonoxideConcentrationMeasurementClass = _MTRClusterCarbonMonoxideConcentrationMeasurementClass{objc.GetClass("MTRClusterCarbonMonoxideConcentrationMeasurement")}
	})
	return MTRClusterCarbonMonoxideConcentrationMeasurementClass
}

type _MTRClusterCarbonMonoxideConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterCarbonMonoxideConcentrationMeasurement] class.
type IMTRClusterCarbonMonoxideConcentrationMeasurement interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterCarbonMonoxideConcentrationMeasurement
type MTRClusterCarbonMonoxideConcentrationMeasurement struct {
	MTRGenericCluster
}

// MTRClusterCarbonMonoxideConcentrationMeasurementFrom constructs a [MTRClusterCarbonMonoxideConcentrationMeasurement] from an unsafe.Pointer.
func MTRClusterCarbonMonoxideConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRClusterCarbonMonoxideConcentrationMeasurement {
	return MTRClusterCarbonMonoxideConcentrationMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterCarbonMonoxideConcentrationMeasurementClass) Alloc() MTRClusterCarbonMonoxideConcentrationMeasurement {
	rv := objc.Send[MTRClusterCarbonMonoxideConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterCarbonMonoxideConcentrationMeasurementClass) New() MTRClusterCarbonMonoxideConcentrationMeasurement {
	rv := objc.Send[MTRClusterCarbonMonoxideConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterCarbonMonoxideConcentrationMeasurement) Init() MTRClusterCarbonMonoxideConcentrationMeasurement {
	rv := objc.Send[MTRClusterCarbonMonoxideConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterCarbonMonoxideConcentrationMeasurement) Autorelease() MTRClusterCarbonMonoxideConcentrationMeasurement {
	rv := objc.Send[MTRClusterCarbonMonoxideConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterCarbonMonoxideConcentrationMeasurement creates a new MTRClusterCarbonMonoxideConcentrationMeasurement instance.
func NewMTRClusterCarbonMonoxideConcentrationMeasurement() MTRClusterCarbonMonoxideConcentrationMeasurement {
	return getMTRClusterCarbonMonoxideConcentrationMeasurementClass().New()
}




