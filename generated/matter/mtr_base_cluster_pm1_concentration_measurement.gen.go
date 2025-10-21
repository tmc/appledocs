// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterPM1ConcentrationMeasurement] class.
var (
	MTRBaseClusterPM1ConcentrationMeasurementClass     _MTRBaseClusterPM1ConcentrationMeasurementClass
	MTRBaseClusterPM1ConcentrationMeasurementClassOnce sync.Once
)

func getMTRBaseClusterPM1ConcentrationMeasurementClass() _MTRBaseClusterPM1ConcentrationMeasurementClass {
	MTRBaseClusterPM1ConcentrationMeasurementClassOnce.Do(func() {
		MTRBaseClusterPM1ConcentrationMeasurementClass = _MTRBaseClusterPM1ConcentrationMeasurementClass{objc.GetClass("MTRBaseClusterPM1ConcentrationMeasurement")}
	})
	return MTRBaseClusterPM1ConcentrationMeasurementClass
}

type _MTRBaseClusterPM1ConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterPM1ConcentrationMeasurement] class.
type IMTRBaseClusterPM1ConcentrationMeasurement interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPM1ConcentrationMeasurement
type MTRBaseClusterPM1ConcentrationMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterPM1ConcentrationMeasurementFrom constructs a [MTRBaseClusterPM1ConcentrationMeasurement] from an unsafe.Pointer.
func MTRBaseClusterPM1ConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterPM1ConcentrationMeasurement {
	return MTRBaseClusterPM1ConcentrationMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterPM1ConcentrationMeasurementClass) Alloc() MTRBaseClusterPM1ConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterPM1ConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterPM1ConcentrationMeasurementClass) New() MTRBaseClusterPM1ConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterPM1ConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterPM1ConcentrationMeasurement) Init() MTRBaseClusterPM1ConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterPM1ConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterPM1ConcentrationMeasurement) Autorelease() MTRBaseClusterPM1ConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterPM1ConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterPM1ConcentrationMeasurement creates a new MTRBaseClusterPM1ConcentrationMeasurement instance.
func NewMTRBaseClusterPM1ConcentrationMeasurement() MTRBaseClusterPM1ConcentrationMeasurement {
	return getMTRBaseClusterPM1ConcentrationMeasurementClass().New()
}




