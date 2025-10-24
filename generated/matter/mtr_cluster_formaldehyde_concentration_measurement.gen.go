// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterFormaldehydeConcentrationMeasurement] class.
var (
	MTRClusterFormaldehydeConcentrationMeasurementClass     _MTRClusterFormaldehydeConcentrationMeasurementClass
	MTRClusterFormaldehydeConcentrationMeasurementClassOnce sync.Once
)

func getMTRClusterFormaldehydeConcentrationMeasurementClass() _MTRClusterFormaldehydeConcentrationMeasurementClass {
	MTRClusterFormaldehydeConcentrationMeasurementClassOnce.Do(func() {
		MTRClusterFormaldehydeConcentrationMeasurementClass = _MTRClusterFormaldehydeConcentrationMeasurementClass{objc.GetClass("MTRClusterFormaldehydeConcentrationMeasurement")}
	})
	return MTRClusterFormaldehydeConcentrationMeasurementClass
}

type _MTRClusterFormaldehydeConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterFormaldehydeConcentrationMeasurement] class.
type IMTRClusterFormaldehydeConcentrationMeasurement interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterFormaldehydeConcentrationMeasurement
type MTRClusterFormaldehydeConcentrationMeasurement struct {
	MTRGenericCluster
}

// MTRClusterFormaldehydeConcentrationMeasurementFrom constructs a [MTRClusterFormaldehydeConcentrationMeasurement] from an unsafe.Pointer.
func MTRClusterFormaldehydeConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRClusterFormaldehydeConcentrationMeasurement {
	return MTRClusterFormaldehydeConcentrationMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterFormaldehydeConcentrationMeasurementClass) Alloc() MTRClusterFormaldehydeConcentrationMeasurement {
	rv := objc.Send[MTRClusterFormaldehydeConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterFormaldehydeConcentrationMeasurementClass) New() MTRClusterFormaldehydeConcentrationMeasurement {
	rv := objc.Send[MTRClusterFormaldehydeConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterFormaldehydeConcentrationMeasurement) Init() MTRClusterFormaldehydeConcentrationMeasurement {
	rv := objc.Send[MTRClusterFormaldehydeConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterFormaldehydeConcentrationMeasurement) Autorelease() MTRClusterFormaldehydeConcentrationMeasurement {
	rv := objc.Send[MTRClusterFormaldehydeConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterFormaldehydeConcentrationMeasurement creates a new MTRClusterFormaldehydeConcentrationMeasurement instance.
func NewMTRClusterFormaldehydeConcentrationMeasurement() MTRClusterFormaldehydeConcentrationMeasurement {
	return getMTRClusterFormaldehydeConcentrationMeasurementClass().New()
}




