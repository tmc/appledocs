// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterElectricalMeasurement] class.
var (
	MTRBaseClusterElectricalMeasurementClass     _MTRBaseClusterElectricalMeasurementClass
	MTRBaseClusterElectricalMeasurementClassOnce sync.Once
)

func getMTRBaseClusterElectricalMeasurementClass() _MTRBaseClusterElectricalMeasurementClass {
	MTRBaseClusterElectricalMeasurementClassOnce.Do(func() {
		MTRBaseClusterElectricalMeasurementClass = _MTRBaseClusterElectricalMeasurementClass{objc.GetClass("MTRBaseClusterElectricalMeasurement")}
	})
	return MTRBaseClusterElectricalMeasurementClass
}

type _MTRBaseClusterElectricalMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterElectricalMeasurement] class.
type IMTRBaseClusterElectricalMeasurement interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterElectricalMeasurement
type MTRBaseClusterElectricalMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterElectricalMeasurementFrom constructs a [MTRBaseClusterElectricalMeasurement] from an unsafe.Pointer.
func MTRBaseClusterElectricalMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterElectricalMeasurement {
	return MTRBaseClusterElectricalMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterElectricalMeasurementClass) Alloc() MTRBaseClusterElectricalMeasurement {
	rv := objc.Send[MTRBaseClusterElectricalMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterElectricalMeasurementClass) New() MTRBaseClusterElectricalMeasurement {
	rv := objc.Send[MTRBaseClusterElectricalMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterElectricalMeasurement) Init() MTRBaseClusterElectricalMeasurement {
	rv := objc.Send[MTRBaseClusterElectricalMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterElectricalMeasurement) Autorelease() MTRBaseClusterElectricalMeasurement {
	rv := objc.Send[MTRBaseClusterElectricalMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterElectricalMeasurement creates a new MTRBaseClusterElectricalMeasurement instance.
func NewMTRBaseClusterElectricalMeasurement() MTRBaseClusterElectricalMeasurement {
	return getMTRBaseClusterElectricalMeasurementClass().New()
}
