// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterRadonConcentrationMeasurement] class.
var (
	MTRBaseClusterRadonConcentrationMeasurementClass     _MTRBaseClusterRadonConcentrationMeasurementClass
	MTRBaseClusterRadonConcentrationMeasurementClassOnce sync.Once
)

func getMTRBaseClusterRadonConcentrationMeasurementClass() _MTRBaseClusterRadonConcentrationMeasurementClass {
	MTRBaseClusterRadonConcentrationMeasurementClassOnce.Do(func() {
		MTRBaseClusterRadonConcentrationMeasurementClass = _MTRBaseClusterRadonConcentrationMeasurementClass{objc.GetClass("MTRBaseClusterRadonConcentrationMeasurement")}
	})
	return MTRBaseClusterRadonConcentrationMeasurementClass
}

type _MTRBaseClusterRadonConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterRadonConcentrationMeasurement] class.
type IMTRBaseClusterRadonConcentrationMeasurement interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRadonConcentrationMeasurement
type MTRBaseClusterRadonConcentrationMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterRadonConcentrationMeasurementFrom constructs a [MTRBaseClusterRadonConcentrationMeasurement] from an unsafe.Pointer.
func MTRBaseClusterRadonConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterRadonConcentrationMeasurement {
	return MTRBaseClusterRadonConcentrationMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterRadonConcentrationMeasurementClass) Alloc() MTRBaseClusterRadonConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterRadonConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterRadonConcentrationMeasurementClass) New() MTRBaseClusterRadonConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterRadonConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterRadonConcentrationMeasurement) Init() MTRBaseClusterRadonConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterRadonConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterRadonConcentrationMeasurement) Autorelease() MTRBaseClusterRadonConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterRadonConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterRadonConcentrationMeasurement creates a new MTRBaseClusterRadonConcentrationMeasurement instance.
func NewMTRBaseClusterRadonConcentrationMeasurement() MTRBaseClusterRadonConcentrationMeasurement {
	return getMTRBaseClusterRadonConcentrationMeasurementClass().New()
}




