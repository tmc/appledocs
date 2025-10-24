// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterTemperatureMeasurement] class.
var (
	MTRBaseClusterTemperatureMeasurementClass     _MTRBaseClusterTemperatureMeasurementClass
	MTRBaseClusterTemperatureMeasurementClassOnce sync.Once
)

func getMTRBaseClusterTemperatureMeasurementClass() _MTRBaseClusterTemperatureMeasurementClass {
	MTRBaseClusterTemperatureMeasurementClassOnce.Do(func() {
		MTRBaseClusterTemperatureMeasurementClass = _MTRBaseClusterTemperatureMeasurementClass{objc.GetClass("MTRBaseClusterTemperatureMeasurement")}
	})
	return MTRBaseClusterTemperatureMeasurementClass
}

type _MTRBaseClusterTemperatureMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterTemperatureMeasurement] class.
type IMTRBaseClusterTemperatureMeasurement interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTemperatureMeasurement
type MTRBaseClusterTemperatureMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterTemperatureMeasurementFrom constructs a [MTRBaseClusterTemperatureMeasurement] from an unsafe.Pointer.
func MTRBaseClusterTemperatureMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterTemperatureMeasurement {
	return MTRBaseClusterTemperatureMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterTemperatureMeasurementClass) Alloc() MTRBaseClusterTemperatureMeasurement {
	rv := objc.Send[MTRBaseClusterTemperatureMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterTemperatureMeasurementClass) New() MTRBaseClusterTemperatureMeasurement {
	rv := objc.Send[MTRBaseClusterTemperatureMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterTemperatureMeasurement) Init() MTRBaseClusterTemperatureMeasurement {
	rv := objc.Send[MTRBaseClusterTemperatureMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterTemperatureMeasurement) Autorelease() MTRBaseClusterTemperatureMeasurement {
	rv := objc.Send[MTRBaseClusterTemperatureMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterTemperatureMeasurement creates a new MTRBaseClusterTemperatureMeasurement instance.
func NewMTRBaseClusterTemperatureMeasurement() MTRBaseClusterTemperatureMeasurement {
	return getMTRBaseClusterTemperatureMeasurementClass().New()
}




