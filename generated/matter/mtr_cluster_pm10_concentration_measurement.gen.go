// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterPM10ConcentrationMeasurement] class.
var (
	MTRClusterPM10ConcentrationMeasurementClass     _MTRClusterPM10ConcentrationMeasurementClass
	MTRClusterPM10ConcentrationMeasurementClassOnce sync.Once
)

func getMTRClusterPM10ConcentrationMeasurementClass() _MTRClusterPM10ConcentrationMeasurementClass {
	MTRClusterPM10ConcentrationMeasurementClassOnce.Do(func() {
		MTRClusterPM10ConcentrationMeasurementClass = _MTRClusterPM10ConcentrationMeasurementClass{objc.GetClass("MTRClusterPM10ConcentrationMeasurement")}
	})
	return MTRClusterPM10ConcentrationMeasurementClass
}

type _MTRClusterPM10ConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterPM10ConcentrationMeasurement] class.
type IMTRClusterPM10ConcentrationMeasurement interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPM10ConcentrationMeasurement
type MTRClusterPM10ConcentrationMeasurement struct {
	MTRGenericCluster
}

// MTRClusterPM10ConcentrationMeasurementFrom constructs a [MTRClusterPM10ConcentrationMeasurement] from an unsafe.Pointer.
func MTRClusterPM10ConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRClusterPM10ConcentrationMeasurement {
	return MTRClusterPM10ConcentrationMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterPM10ConcentrationMeasurementClass) Alloc() MTRClusterPM10ConcentrationMeasurement {
	rv := objc.Send[MTRClusterPM10ConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterPM10ConcentrationMeasurementClass) New() MTRClusterPM10ConcentrationMeasurement {
	rv := objc.Send[MTRClusterPM10ConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterPM10ConcentrationMeasurement) Init() MTRClusterPM10ConcentrationMeasurement {
	rv := objc.Send[MTRClusterPM10ConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterPM10ConcentrationMeasurement) Autorelease() MTRClusterPM10ConcentrationMeasurement {
	rv := objc.Send[MTRClusterPM10ConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterPM10ConcentrationMeasurement creates a new MTRClusterPM10ConcentrationMeasurement instance.
func NewMTRClusterPM10ConcentrationMeasurement() MTRClusterPM10ConcentrationMeasurement {
	return getMTRClusterPM10ConcentrationMeasurementClass().New()
}




