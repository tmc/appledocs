// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterPM25ConcentrationMeasurement] class.
var (
	MTRClusterPM25ConcentrationMeasurementClass     _MTRClusterPM25ConcentrationMeasurementClass
	MTRClusterPM25ConcentrationMeasurementClassOnce sync.Once
)

func getMTRClusterPM25ConcentrationMeasurementClass() _MTRClusterPM25ConcentrationMeasurementClass {
	MTRClusterPM25ConcentrationMeasurementClassOnce.Do(func() {
		MTRClusterPM25ConcentrationMeasurementClass = _MTRClusterPM25ConcentrationMeasurementClass{objc.GetClass("MTRClusterPM25ConcentrationMeasurement")}
	})
	return MTRClusterPM25ConcentrationMeasurementClass
}

type _MTRClusterPM25ConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterPM25ConcentrationMeasurement] class.
type IMTRClusterPM25ConcentrationMeasurement interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPM25ConcentrationMeasurement
type MTRClusterPM25ConcentrationMeasurement struct {
	MTRGenericCluster
}

// MTRClusterPM25ConcentrationMeasurementFrom constructs a [MTRClusterPM25ConcentrationMeasurement] from an unsafe.Pointer.
func MTRClusterPM25ConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRClusterPM25ConcentrationMeasurement {
	return MTRClusterPM25ConcentrationMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterPM25ConcentrationMeasurementClass) Alloc() MTRClusterPM25ConcentrationMeasurement {
	rv := objc.Send[MTRClusterPM25ConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterPM25ConcentrationMeasurementClass) New() MTRClusterPM25ConcentrationMeasurement {
	rv := objc.Send[MTRClusterPM25ConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterPM25ConcentrationMeasurement) Init() MTRClusterPM25ConcentrationMeasurement {
	rv := objc.Send[MTRClusterPM25ConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterPM25ConcentrationMeasurement) Autorelease() MTRClusterPM25ConcentrationMeasurement {
	rv := objc.Send[MTRClusterPM25ConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterPM25ConcentrationMeasurement creates a new MTRClusterPM25ConcentrationMeasurement instance.
func NewMTRClusterPM25ConcentrationMeasurement() MTRClusterPM25ConcentrationMeasurement {
	return getMTRClusterPM25ConcentrationMeasurementClass().New()
}




