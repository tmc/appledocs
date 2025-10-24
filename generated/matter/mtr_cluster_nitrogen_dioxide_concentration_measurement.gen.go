// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterNitrogenDioxideConcentrationMeasurement] class.
var (
	MTRClusterNitrogenDioxideConcentrationMeasurementClass     _MTRClusterNitrogenDioxideConcentrationMeasurementClass
	MTRClusterNitrogenDioxideConcentrationMeasurementClassOnce sync.Once
)

func getMTRClusterNitrogenDioxideConcentrationMeasurementClass() _MTRClusterNitrogenDioxideConcentrationMeasurementClass {
	MTRClusterNitrogenDioxideConcentrationMeasurementClassOnce.Do(func() {
		MTRClusterNitrogenDioxideConcentrationMeasurementClass = _MTRClusterNitrogenDioxideConcentrationMeasurementClass{objc.GetClass("MTRClusterNitrogenDioxideConcentrationMeasurement")}
	})
	return MTRClusterNitrogenDioxideConcentrationMeasurementClass
}

type _MTRClusterNitrogenDioxideConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterNitrogenDioxideConcentrationMeasurement] class.
type IMTRClusterNitrogenDioxideConcentrationMeasurement interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterNitrogenDioxideConcentrationMeasurement
type MTRClusterNitrogenDioxideConcentrationMeasurement struct {
	MTRGenericCluster
}

// MTRClusterNitrogenDioxideConcentrationMeasurementFrom constructs a [MTRClusterNitrogenDioxideConcentrationMeasurement] from an unsafe.Pointer.
func MTRClusterNitrogenDioxideConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRClusterNitrogenDioxideConcentrationMeasurement {
	return MTRClusterNitrogenDioxideConcentrationMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterNitrogenDioxideConcentrationMeasurementClass) Alloc() MTRClusterNitrogenDioxideConcentrationMeasurement {
	rv := objc.Send[MTRClusterNitrogenDioxideConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterNitrogenDioxideConcentrationMeasurementClass) New() MTRClusterNitrogenDioxideConcentrationMeasurement {
	rv := objc.Send[MTRClusterNitrogenDioxideConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterNitrogenDioxideConcentrationMeasurement) Init() MTRClusterNitrogenDioxideConcentrationMeasurement {
	rv := objc.Send[MTRClusterNitrogenDioxideConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterNitrogenDioxideConcentrationMeasurement) Autorelease() MTRClusterNitrogenDioxideConcentrationMeasurement {
	rv := objc.Send[MTRClusterNitrogenDioxideConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterNitrogenDioxideConcentrationMeasurement creates a new MTRClusterNitrogenDioxideConcentrationMeasurement instance.
func NewMTRClusterNitrogenDioxideConcentrationMeasurement() MTRClusterNitrogenDioxideConcentrationMeasurement {
	return getMTRClusterNitrogenDioxideConcentrationMeasurementClass().New()
}




