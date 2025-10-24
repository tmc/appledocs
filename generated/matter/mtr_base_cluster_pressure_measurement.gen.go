// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterPressureMeasurement] class.
var (
	MTRBaseClusterPressureMeasurementClass     _MTRBaseClusterPressureMeasurementClass
	MTRBaseClusterPressureMeasurementClassOnce sync.Once
)

func getMTRBaseClusterPressureMeasurementClass() _MTRBaseClusterPressureMeasurementClass {
	MTRBaseClusterPressureMeasurementClassOnce.Do(func() {
		MTRBaseClusterPressureMeasurementClass = _MTRBaseClusterPressureMeasurementClass{objc.GetClass("MTRBaseClusterPressureMeasurement")}
	})
	return MTRBaseClusterPressureMeasurementClass
}

type _MTRBaseClusterPressureMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterPressureMeasurement] class.
type IMTRBaseClusterPressureMeasurement interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPressureMeasurement
type MTRBaseClusterPressureMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterPressureMeasurementFrom constructs a [MTRBaseClusterPressureMeasurement] from an unsafe.Pointer.
func MTRBaseClusterPressureMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterPressureMeasurement {
	return MTRBaseClusterPressureMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterPressureMeasurementClass) Alloc() MTRBaseClusterPressureMeasurement {
	rv := objc.Send[MTRBaseClusterPressureMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterPressureMeasurementClass) New() MTRBaseClusterPressureMeasurement {
	rv := objc.Send[MTRBaseClusterPressureMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterPressureMeasurement) Init() MTRBaseClusterPressureMeasurement {
	rv := objc.Send[MTRBaseClusterPressureMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterPressureMeasurement) Autorelease() MTRBaseClusterPressureMeasurement {
	rv := objc.Send[MTRBaseClusterPressureMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterPressureMeasurement creates a new MTRBaseClusterPressureMeasurement instance.
func NewMTRBaseClusterPressureMeasurement() MTRBaseClusterPressureMeasurement {
	return getMTRBaseClusterPressureMeasurementClass().New()
}




