// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterTemperatureMeasurement] class.
var (
	MTRClusterTemperatureMeasurementClass     _MTRClusterTemperatureMeasurementClass
	MTRClusterTemperatureMeasurementClassOnce sync.Once
)

func getMTRClusterTemperatureMeasurementClass() _MTRClusterTemperatureMeasurementClass {
	MTRClusterTemperatureMeasurementClassOnce.Do(func() {
		MTRClusterTemperatureMeasurementClass = _MTRClusterTemperatureMeasurementClass{objc.GetClass("MTRClusterTemperatureMeasurement")}
	})
	return MTRClusterTemperatureMeasurementClass
}

type _MTRClusterTemperatureMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterTemperatureMeasurement] class.
type IMTRClusterTemperatureMeasurement interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterTemperatureMeasurement
type MTRClusterTemperatureMeasurement struct {
	MTRGenericCluster
}

// MTRClusterTemperatureMeasurementFrom constructs a [MTRClusterTemperatureMeasurement] from an unsafe.Pointer.
func MTRClusterTemperatureMeasurementFrom(ptr unsafe.Pointer) MTRClusterTemperatureMeasurement {
	return MTRClusterTemperatureMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterTemperatureMeasurementClass) Alloc() MTRClusterTemperatureMeasurement {
	rv := objc.Send[MTRClusterTemperatureMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterTemperatureMeasurementClass) New() MTRClusterTemperatureMeasurement {
	rv := objc.Send[MTRClusterTemperatureMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterTemperatureMeasurement) Init() MTRClusterTemperatureMeasurement {
	rv := objc.Send[MTRClusterTemperatureMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterTemperatureMeasurement) Autorelease() MTRClusterTemperatureMeasurement {
	rv := objc.Send[MTRClusterTemperatureMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterTemperatureMeasurement creates a new MTRClusterTemperatureMeasurement instance.
func NewMTRClusterTemperatureMeasurement() MTRClusterTemperatureMeasurement {
	return getMTRClusterTemperatureMeasurementClass().New()
}




