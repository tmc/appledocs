// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterRelativeHumidityMeasurement] class.
var (
	MTRClusterRelativeHumidityMeasurementClass     _MTRClusterRelativeHumidityMeasurementClass
	MTRClusterRelativeHumidityMeasurementClassOnce sync.Once
)

func getMTRClusterRelativeHumidityMeasurementClass() _MTRClusterRelativeHumidityMeasurementClass {
	MTRClusterRelativeHumidityMeasurementClassOnce.Do(func() {
		MTRClusterRelativeHumidityMeasurementClass = _MTRClusterRelativeHumidityMeasurementClass{objc.GetClass("MTRClusterRelativeHumidityMeasurement")}
	})
	return MTRClusterRelativeHumidityMeasurementClass
}

type _MTRClusterRelativeHumidityMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterRelativeHumidityMeasurement] class.
type IMTRClusterRelativeHumidityMeasurement interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRelativeHumidityMeasurement
type MTRClusterRelativeHumidityMeasurement struct {
	MTRGenericCluster
}

// MTRClusterRelativeHumidityMeasurementFrom constructs a [MTRClusterRelativeHumidityMeasurement] from an unsafe.Pointer.
func MTRClusterRelativeHumidityMeasurementFrom(ptr unsafe.Pointer) MTRClusterRelativeHumidityMeasurement {
	return MTRClusterRelativeHumidityMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterRelativeHumidityMeasurementClass) Alloc() MTRClusterRelativeHumidityMeasurement {
	rv := objc.Send[MTRClusterRelativeHumidityMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterRelativeHumidityMeasurementClass) New() MTRClusterRelativeHumidityMeasurement {
	rv := objc.Send[MTRClusterRelativeHumidityMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterRelativeHumidityMeasurement) Init() MTRClusterRelativeHumidityMeasurement {
	rv := objc.Send[MTRClusterRelativeHumidityMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterRelativeHumidityMeasurement) Autorelease() MTRClusterRelativeHumidityMeasurement {
	rv := objc.Send[MTRClusterRelativeHumidityMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterRelativeHumidityMeasurement creates a new MTRClusterRelativeHumidityMeasurement instance.
func NewMTRClusterRelativeHumidityMeasurement() MTRClusterRelativeHumidityMeasurement {
	return getMTRClusterRelativeHumidityMeasurementClass().New()
}
