// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterRelativeHumidityMeasurement] class.
var (
	MTRBaseClusterRelativeHumidityMeasurementClass     _MTRBaseClusterRelativeHumidityMeasurementClass
	MTRBaseClusterRelativeHumidityMeasurementClassOnce sync.Once
)

func getMTRBaseClusterRelativeHumidityMeasurementClass() _MTRBaseClusterRelativeHumidityMeasurementClass {
	MTRBaseClusterRelativeHumidityMeasurementClassOnce.Do(func() {
		MTRBaseClusterRelativeHumidityMeasurementClass = _MTRBaseClusterRelativeHumidityMeasurementClass{objc.GetClass("MTRBaseClusterRelativeHumidityMeasurement")}
	})
	return MTRBaseClusterRelativeHumidityMeasurementClass
}

type _MTRBaseClusterRelativeHumidityMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterRelativeHumidityMeasurement] class.
type IMTRBaseClusterRelativeHumidityMeasurement interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRelativeHumidityMeasurement
type MTRBaseClusterRelativeHumidityMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterRelativeHumidityMeasurementFrom constructs a [MTRBaseClusterRelativeHumidityMeasurement] from an unsafe.Pointer.
func MTRBaseClusterRelativeHumidityMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterRelativeHumidityMeasurement {
	return MTRBaseClusterRelativeHumidityMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterRelativeHumidityMeasurementClass) Alloc() MTRBaseClusterRelativeHumidityMeasurement {
	rv := objc.Send[MTRBaseClusterRelativeHumidityMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterRelativeHumidityMeasurementClass) New() MTRBaseClusterRelativeHumidityMeasurement {
	rv := objc.Send[MTRBaseClusterRelativeHumidityMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterRelativeHumidityMeasurement) Init() MTRBaseClusterRelativeHumidityMeasurement {
	rv := objc.Send[MTRBaseClusterRelativeHumidityMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterRelativeHumidityMeasurement) Autorelease() MTRBaseClusterRelativeHumidityMeasurement {
	rv := objc.Send[MTRBaseClusterRelativeHumidityMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterRelativeHumidityMeasurement creates a new MTRBaseClusterRelativeHumidityMeasurement instance.
func NewMTRBaseClusterRelativeHumidityMeasurement() MTRBaseClusterRelativeHumidityMeasurement {
	return getMTRBaseClusterRelativeHumidityMeasurementClass().New()
}




