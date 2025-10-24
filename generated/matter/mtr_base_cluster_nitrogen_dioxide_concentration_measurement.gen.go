// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterNitrogenDioxideConcentrationMeasurement] class.
var (
	MTRBaseClusterNitrogenDioxideConcentrationMeasurementClass     _MTRBaseClusterNitrogenDioxideConcentrationMeasurementClass
	MTRBaseClusterNitrogenDioxideConcentrationMeasurementClassOnce sync.Once
)

func getMTRBaseClusterNitrogenDioxideConcentrationMeasurementClass() _MTRBaseClusterNitrogenDioxideConcentrationMeasurementClass {
	MTRBaseClusterNitrogenDioxideConcentrationMeasurementClassOnce.Do(func() {
		MTRBaseClusterNitrogenDioxideConcentrationMeasurementClass = _MTRBaseClusterNitrogenDioxideConcentrationMeasurementClass{objc.GetClass("MTRBaseClusterNitrogenDioxideConcentrationMeasurement")}
	})
	return MTRBaseClusterNitrogenDioxideConcentrationMeasurementClass
}

type _MTRBaseClusterNitrogenDioxideConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterNitrogenDioxideConcentrationMeasurement] class.
type IMTRBaseClusterNitrogenDioxideConcentrationMeasurement interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterNitrogenDioxideConcentrationMeasurement
type MTRBaseClusterNitrogenDioxideConcentrationMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterNitrogenDioxideConcentrationMeasurementFrom constructs a [MTRBaseClusterNitrogenDioxideConcentrationMeasurement] from an unsafe.Pointer.
func MTRBaseClusterNitrogenDioxideConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterNitrogenDioxideConcentrationMeasurement {
	return MTRBaseClusterNitrogenDioxideConcentrationMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterNitrogenDioxideConcentrationMeasurementClass) Alloc() MTRBaseClusterNitrogenDioxideConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterNitrogenDioxideConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterNitrogenDioxideConcentrationMeasurementClass) New() MTRBaseClusterNitrogenDioxideConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterNitrogenDioxideConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterNitrogenDioxideConcentrationMeasurement) Init() MTRBaseClusterNitrogenDioxideConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterNitrogenDioxideConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterNitrogenDioxideConcentrationMeasurement) Autorelease() MTRBaseClusterNitrogenDioxideConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterNitrogenDioxideConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterNitrogenDioxideConcentrationMeasurement creates a new MTRBaseClusterNitrogenDioxideConcentrationMeasurement instance.
func NewMTRBaseClusterNitrogenDioxideConcentrationMeasurement() MTRBaseClusterNitrogenDioxideConcentrationMeasurement {
	return getMTRBaseClusterNitrogenDioxideConcentrationMeasurementClass().New()
}




