// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterPM25ConcentrationMeasurement] class.
var (
	MTRBaseClusterPM25ConcentrationMeasurementClass     _MTRBaseClusterPM25ConcentrationMeasurementClass
	MTRBaseClusterPM25ConcentrationMeasurementClassOnce sync.Once
)

func getMTRBaseClusterPM25ConcentrationMeasurementClass() _MTRBaseClusterPM25ConcentrationMeasurementClass {
	MTRBaseClusterPM25ConcentrationMeasurementClassOnce.Do(func() {
		MTRBaseClusterPM25ConcentrationMeasurementClass = _MTRBaseClusterPM25ConcentrationMeasurementClass{objc.GetClass("MTRBaseClusterPM25ConcentrationMeasurement")}
	})
	return MTRBaseClusterPM25ConcentrationMeasurementClass
}

type _MTRBaseClusterPM25ConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterPM25ConcentrationMeasurement] class.
type IMTRBaseClusterPM25ConcentrationMeasurement interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPM25ConcentrationMeasurement
type MTRBaseClusterPM25ConcentrationMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterPM25ConcentrationMeasurementFrom constructs a [MTRBaseClusterPM25ConcentrationMeasurement] from an unsafe.Pointer.
func MTRBaseClusterPM25ConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterPM25ConcentrationMeasurement {
	return MTRBaseClusterPM25ConcentrationMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterPM25ConcentrationMeasurementClass) Alloc() MTRBaseClusterPM25ConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterPM25ConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterPM25ConcentrationMeasurementClass) New() MTRBaseClusterPM25ConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterPM25ConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterPM25ConcentrationMeasurement) Init() MTRBaseClusterPM25ConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterPM25ConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterPM25ConcentrationMeasurement) Autorelease() MTRBaseClusterPM25ConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterPM25ConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterPM25ConcentrationMeasurement creates a new MTRBaseClusterPM25ConcentrationMeasurement instance.
func NewMTRBaseClusterPM25ConcentrationMeasurement() MTRBaseClusterPM25ConcentrationMeasurement {
	return getMTRBaseClusterPM25ConcentrationMeasurementClass().New()
}




