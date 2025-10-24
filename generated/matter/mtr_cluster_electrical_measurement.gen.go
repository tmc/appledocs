// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterElectricalMeasurement] class.
var (
	MTRClusterElectricalMeasurementClass     _MTRClusterElectricalMeasurementClass
	MTRClusterElectricalMeasurementClassOnce sync.Once
)

func getMTRClusterElectricalMeasurementClass() _MTRClusterElectricalMeasurementClass {
	MTRClusterElectricalMeasurementClassOnce.Do(func() {
		MTRClusterElectricalMeasurementClass = _MTRClusterElectricalMeasurementClass{objc.GetClass("MTRClusterElectricalMeasurement")}
	})
	return MTRClusterElectricalMeasurementClass
}

type _MTRClusterElectricalMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterElectricalMeasurement] class.
type IMTRClusterElectricalMeasurement interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterElectricalMeasurement
type MTRClusterElectricalMeasurement struct {
	MTRGenericCluster
}

// MTRClusterElectricalMeasurementFrom constructs a [MTRClusterElectricalMeasurement] from an unsafe.Pointer.
func MTRClusterElectricalMeasurementFrom(ptr unsafe.Pointer) MTRClusterElectricalMeasurement {
	return MTRClusterElectricalMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterElectricalMeasurementClass) Alloc() MTRClusterElectricalMeasurement {
	rv := objc.Send[MTRClusterElectricalMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterElectricalMeasurementClass) New() MTRClusterElectricalMeasurement {
	rv := objc.Send[MTRClusterElectricalMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterElectricalMeasurement) Init() MTRClusterElectricalMeasurement {
	rv := objc.Send[MTRClusterElectricalMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterElectricalMeasurement) Autorelease() MTRClusterElectricalMeasurement {
	rv := objc.Send[MTRClusterElectricalMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterElectricalMeasurement creates a new MTRClusterElectricalMeasurement instance.
func NewMTRClusterElectricalMeasurement() MTRClusterElectricalMeasurement {
	return getMTRClusterElectricalMeasurementClass().New()
}
