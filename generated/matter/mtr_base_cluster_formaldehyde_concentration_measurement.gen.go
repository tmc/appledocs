// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterFormaldehydeConcentrationMeasurement] class.
var (
	MTRBaseClusterFormaldehydeConcentrationMeasurementClass     _MTRBaseClusterFormaldehydeConcentrationMeasurementClass
	MTRBaseClusterFormaldehydeConcentrationMeasurementClassOnce sync.Once
)

func getMTRBaseClusterFormaldehydeConcentrationMeasurementClass() _MTRBaseClusterFormaldehydeConcentrationMeasurementClass {
	MTRBaseClusterFormaldehydeConcentrationMeasurementClassOnce.Do(func() {
		MTRBaseClusterFormaldehydeConcentrationMeasurementClass = _MTRBaseClusterFormaldehydeConcentrationMeasurementClass{objc.GetClass("MTRBaseClusterFormaldehydeConcentrationMeasurement")}
	})
	return MTRBaseClusterFormaldehydeConcentrationMeasurementClass
}

type _MTRBaseClusterFormaldehydeConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterFormaldehydeConcentrationMeasurement] class.
type IMTRBaseClusterFormaldehydeConcentrationMeasurement interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterFormaldehydeConcentrationMeasurement
type MTRBaseClusterFormaldehydeConcentrationMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterFormaldehydeConcentrationMeasurementFrom constructs a [MTRBaseClusterFormaldehydeConcentrationMeasurement] from an unsafe.Pointer.
func MTRBaseClusterFormaldehydeConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterFormaldehydeConcentrationMeasurement {
	return MTRBaseClusterFormaldehydeConcentrationMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterFormaldehydeConcentrationMeasurementClass) Alloc() MTRBaseClusterFormaldehydeConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterFormaldehydeConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterFormaldehydeConcentrationMeasurementClass) New() MTRBaseClusterFormaldehydeConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterFormaldehydeConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterFormaldehydeConcentrationMeasurement) Init() MTRBaseClusterFormaldehydeConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterFormaldehydeConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterFormaldehydeConcentrationMeasurement) Autorelease() MTRBaseClusterFormaldehydeConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterFormaldehydeConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterFormaldehydeConcentrationMeasurement creates a new MTRBaseClusterFormaldehydeConcentrationMeasurement instance.
func NewMTRBaseClusterFormaldehydeConcentrationMeasurement() MTRBaseClusterFormaldehydeConcentrationMeasurement {
	return getMTRBaseClusterFormaldehydeConcentrationMeasurementClass().New()
}
