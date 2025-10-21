// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterPM1ConcentrationMeasurement] class.
var (
	MTRClusterPM1ConcentrationMeasurementClass     _MTRClusterPM1ConcentrationMeasurementClass
	MTRClusterPM1ConcentrationMeasurementClassOnce sync.Once
)

func getMTRClusterPM1ConcentrationMeasurementClass() _MTRClusterPM1ConcentrationMeasurementClass {
	MTRClusterPM1ConcentrationMeasurementClassOnce.Do(func() {
		MTRClusterPM1ConcentrationMeasurementClass = _MTRClusterPM1ConcentrationMeasurementClass{objc.GetClass("MTRClusterPM1ConcentrationMeasurement")}
	})
	return MTRClusterPM1ConcentrationMeasurementClass
}

type _MTRClusterPM1ConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterPM1ConcentrationMeasurement] class.
type IMTRClusterPM1ConcentrationMeasurement interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPM1ConcentrationMeasurement
type MTRClusterPM1ConcentrationMeasurement struct {
	MTRGenericCluster
}

// MTRClusterPM1ConcentrationMeasurementFrom constructs a [MTRClusterPM1ConcentrationMeasurement] from an unsafe.Pointer.
func MTRClusterPM1ConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRClusterPM1ConcentrationMeasurement {
	return MTRClusterPM1ConcentrationMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterPM1ConcentrationMeasurementClass) Alloc() MTRClusterPM1ConcentrationMeasurement {
	rv := objc.Send[MTRClusterPM1ConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterPM1ConcentrationMeasurementClass) New() MTRClusterPM1ConcentrationMeasurement {
	rv := objc.Send[MTRClusterPM1ConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterPM1ConcentrationMeasurement) Init() MTRClusterPM1ConcentrationMeasurement {
	rv := objc.Send[MTRClusterPM1ConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterPM1ConcentrationMeasurement) Autorelease() MTRClusterPM1ConcentrationMeasurement {
	rv := objc.Send[MTRClusterPM1ConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterPM1ConcentrationMeasurement creates a new MTRClusterPM1ConcentrationMeasurement instance.
func NewMTRClusterPM1ConcentrationMeasurement() MTRClusterPM1ConcentrationMeasurement {
	return getMTRClusterPM1ConcentrationMeasurementClass().New()
}




