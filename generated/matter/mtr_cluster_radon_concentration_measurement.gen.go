// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterRadonConcentrationMeasurement] class.
var (
	MTRClusterRadonConcentrationMeasurementClass     _MTRClusterRadonConcentrationMeasurementClass
	MTRClusterRadonConcentrationMeasurementClassOnce sync.Once
)

func getMTRClusterRadonConcentrationMeasurementClass() _MTRClusterRadonConcentrationMeasurementClass {
	MTRClusterRadonConcentrationMeasurementClassOnce.Do(func() {
		MTRClusterRadonConcentrationMeasurementClass = _MTRClusterRadonConcentrationMeasurementClass{objc.GetClass("MTRClusterRadonConcentrationMeasurement")}
	})
	return MTRClusterRadonConcentrationMeasurementClass
}

type _MTRClusterRadonConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterRadonConcentrationMeasurement] class.
type IMTRClusterRadonConcentrationMeasurement interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRadonConcentrationMeasurement
type MTRClusterRadonConcentrationMeasurement struct {
	MTRGenericCluster
}

// MTRClusterRadonConcentrationMeasurementFrom constructs a [MTRClusterRadonConcentrationMeasurement] from an unsafe.Pointer.
func MTRClusterRadonConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRClusterRadonConcentrationMeasurement {
	return MTRClusterRadonConcentrationMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterRadonConcentrationMeasurementClass) Alloc() MTRClusterRadonConcentrationMeasurement {
	rv := objc.Send[MTRClusterRadonConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterRadonConcentrationMeasurementClass) New() MTRClusterRadonConcentrationMeasurement {
	rv := objc.Send[MTRClusterRadonConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterRadonConcentrationMeasurement) Init() MTRClusterRadonConcentrationMeasurement {
	rv := objc.Send[MTRClusterRadonConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterRadonConcentrationMeasurement) Autorelease() MTRClusterRadonConcentrationMeasurement {
	rv := objc.Send[MTRClusterRadonConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterRadonConcentrationMeasurement creates a new MTRClusterRadonConcentrationMeasurement instance.
func NewMTRClusterRadonConcentrationMeasurement() MTRClusterRadonConcentrationMeasurement {
	return getMTRClusterRadonConcentrationMeasurementClass().New()
}
