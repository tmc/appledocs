// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterPM10ConcentrationMeasurement] class.
var (
	MTRBaseClusterPM10ConcentrationMeasurementClass     _MTRBaseClusterPM10ConcentrationMeasurementClass
	MTRBaseClusterPM10ConcentrationMeasurementClassOnce sync.Once
)

func getMTRBaseClusterPM10ConcentrationMeasurementClass() _MTRBaseClusterPM10ConcentrationMeasurementClass {
	MTRBaseClusterPM10ConcentrationMeasurementClassOnce.Do(func() {
		MTRBaseClusterPM10ConcentrationMeasurementClass = _MTRBaseClusterPM10ConcentrationMeasurementClass{objc.GetClass("MTRBaseClusterPM10ConcentrationMeasurement")}
	})
	return MTRBaseClusterPM10ConcentrationMeasurementClass
}

type _MTRBaseClusterPM10ConcentrationMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterPM10ConcentrationMeasurement] class.
type IMTRBaseClusterPM10ConcentrationMeasurement interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPM10ConcentrationMeasurement
type MTRBaseClusterPM10ConcentrationMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterPM10ConcentrationMeasurementFrom constructs a [MTRBaseClusterPM10ConcentrationMeasurement] from an unsafe.Pointer.
func MTRBaseClusterPM10ConcentrationMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterPM10ConcentrationMeasurement {
	return MTRBaseClusterPM10ConcentrationMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterPM10ConcentrationMeasurementClass) Alloc() MTRBaseClusterPM10ConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterPM10ConcentrationMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterPM10ConcentrationMeasurementClass) New() MTRBaseClusterPM10ConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterPM10ConcentrationMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterPM10ConcentrationMeasurement) Init() MTRBaseClusterPM10ConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterPM10ConcentrationMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterPM10ConcentrationMeasurement) Autorelease() MTRBaseClusterPM10ConcentrationMeasurement {
	rv := objc.Send[MTRBaseClusterPM10ConcentrationMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterPM10ConcentrationMeasurement creates a new MTRBaseClusterPM10ConcentrationMeasurement instance.
func NewMTRBaseClusterPM10ConcentrationMeasurement() MTRBaseClusterPM10ConcentrationMeasurement {
	return getMTRBaseClusterPM10ConcentrationMeasurementClass().New()
}
