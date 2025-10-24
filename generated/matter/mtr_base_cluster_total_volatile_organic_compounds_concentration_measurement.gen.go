// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement] class.
var (
	MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass     _MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass
	MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClassOnce sync.Once
)

func getMTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass() _MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass {
	MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClassOnce.Do(func() {
		MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass = _MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass{objc.GetClass("MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement")}
	})
	return MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass
}

type _MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement] class.
type IMTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement
type MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurementFrom constructs a [MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement] from an unsafe.Pointer.
func MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement {
	return MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass) Alloc() MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass) New() MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement) Init() MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement) Autorelease() MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement creates a new MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement instance.
func NewMTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement() MTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurement {
	return getMTRBaseClusterTotalVolatileOrganicCompoundsConcentrationMeasurementClass().New()
}




