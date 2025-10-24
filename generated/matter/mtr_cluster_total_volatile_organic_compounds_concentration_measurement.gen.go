// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement] class.
var (
	MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass     _MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass
	MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClassOnce sync.Once
)

func getMTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass() _MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass {
	MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClassOnce.Do(func() {
		MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass = _MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass{objc.GetClass("MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement")}
	})
	return MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass
}

type _MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement] class.
type IMTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement
type MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement struct {
	MTRGenericCluster
}

// MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurementFrom constructs a [MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement] from an unsafe.Pointer.
func MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement {
	return MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass) Alloc() MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement {
	rv := objc.Send[MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass) New() MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement {
	rv := objc.Send[MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement) Init() MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement {
	rv := objc.Send[MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement) Autorelease() MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement {
	rv := objc.Send[MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement creates a new MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement instance.
func NewMTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement() MTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurement {
	return getMTRClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass().New()
}
