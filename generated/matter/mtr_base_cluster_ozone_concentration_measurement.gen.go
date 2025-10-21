// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterOzoneConcentrationMeasurement] class.
var (
	MTRBaseClusterOzoneConcentrationMeasurementClass     _MTRBaseClusterOzoneConcentrationMeasurementClass
	MTRBaseClusterOzoneConcentrationMeasurementClassOnce sync.Once
)

func getMTRBaseClusterOzoneConcentrationMeasurementClass() _MTRBaseClusterOzoneConcentrationMeasurementClass {
	MTRBaseClusterOzoneConcentrationMeasurementClassOnce.Do(func() {
		MTRBaseClusterOzoneConcentrationMeasurementClass = _MTRBaseClusterOzoneConcentrationMeasurementClass{objc.GetClass("MTRBaseClusterOzoneConcentrationMeasurement")}
	})
	return MTRBaseClusterOzoneConcentrationMeasurementClass
}

type _MTRBaseClusterOzoneConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterOzoneConcentrationMeasurement] class.
type IMTRBaseClusterOzoneConcentrationMeasurement interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterOzoneConcentrationMeasurement
type MTRBaseClusterOzoneConcentrationMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterOzoneConcentrationMeasurementFrom constructs a [MTRBaseClusterOzoneConcentrationMeasurement] from an unsafe.Pointer.
func MTRBaseClusterOzoneConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterOzoneConcentrationMeasurement {
	return MTRBaseClusterOzoneConcentrationMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterOzoneConcentrationMeasurementClass) Alloc() MTRBaseClusterOzoneConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterOzoneConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterOzoneConcentrationMeasurementClass) New() MTRBaseClusterOzoneConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterOzoneConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterOzoneConcentrationMeasurement) Init() MTRBaseClusterOzoneConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterOzoneConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterOzoneConcentrationMeasurement) Autorelease() MTRBaseClusterOzoneConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterOzoneConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterOzoneConcentrationMeasurement creates a new MTRBaseClusterOzoneConcentrationMeasurement instance.
func NewMTRBaseClusterOzoneConcentrationMeasurement() MTRBaseClusterOzoneConcentrationMeasurement {
	return getMTRBaseClusterOzoneConcentrationMeasurementClass().New()
}




