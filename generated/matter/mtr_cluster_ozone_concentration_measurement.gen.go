// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterOzoneConcentrationMeasurement] class.
var (
	MTRClusterOzoneConcentrationMeasurementClass     _MTRClusterOzoneConcentrationMeasurementClass
	MTRClusterOzoneConcentrationMeasurementClassOnce sync.Once
)

func getMTRClusterOzoneConcentrationMeasurementClass() _MTRClusterOzoneConcentrationMeasurementClass {
	MTRClusterOzoneConcentrationMeasurementClassOnce.Do(func() {
		MTRClusterOzoneConcentrationMeasurementClass = _MTRClusterOzoneConcentrationMeasurementClass{objc.GetClass("MTRClusterOzoneConcentrationMeasurement")}
	})
	return MTRClusterOzoneConcentrationMeasurementClass
}

type _MTRClusterOzoneConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterOzoneConcentrationMeasurement] class.
type IMTRClusterOzoneConcentrationMeasurement interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterOzoneConcentrationMeasurement
type MTRClusterOzoneConcentrationMeasurement struct {
	MTRGenericCluster
}

// MTRClusterOzoneConcentrationMeasurementFrom constructs a [MTRClusterOzoneConcentrationMeasurement] from an unsafe.Pointer.
func MTRClusterOzoneConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRClusterOzoneConcentrationMeasurement {
	return MTRClusterOzoneConcentrationMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterOzoneConcentrationMeasurementClass) Alloc() MTRClusterOzoneConcentrationMeasurement {
	rv := objc.Send[MTRClusterOzoneConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterOzoneConcentrationMeasurementClass) New() MTRClusterOzoneConcentrationMeasurement {
	rv := objc.Send[MTRClusterOzoneConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterOzoneConcentrationMeasurement) Init() MTRClusterOzoneConcentrationMeasurement {
	rv := objc.Send[MTRClusterOzoneConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterOzoneConcentrationMeasurement) Autorelease() MTRClusterOzoneConcentrationMeasurement {
	rv := objc.Send[MTRClusterOzoneConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterOzoneConcentrationMeasurement creates a new MTRClusterOzoneConcentrationMeasurement instance.
func NewMTRClusterOzoneConcentrationMeasurement() MTRClusterOzoneConcentrationMeasurement {
	return getMTRClusterOzoneConcentrationMeasurementClass().New()
}




