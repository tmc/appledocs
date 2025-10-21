// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterFlowMeasurement] class.
var (
	MTRClusterFlowMeasurementClass     _MTRClusterFlowMeasurementClass
	MTRClusterFlowMeasurementClassOnce sync.Once
)

func getMTRClusterFlowMeasurementClass() _MTRClusterFlowMeasurementClass {
	MTRClusterFlowMeasurementClassOnce.Do(func() {
		MTRClusterFlowMeasurementClass = _MTRClusterFlowMeasurementClass{objc.GetClass("MTRClusterFlowMeasurement")}
	})
	return MTRClusterFlowMeasurementClass
}

type _MTRClusterFlowMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterFlowMeasurement] class.
type IMTRClusterFlowMeasurement interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterFlowMeasurement
type MTRClusterFlowMeasurement struct {
	MTRGenericCluster
}

// MTRClusterFlowMeasurementFrom constructs a [MTRClusterFlowMeasurement] from an unsafe.Pointer.
func MTRClusterFlowMeasurementFrom(ptr unsafe.Pointer) MTRClusterFlowMeasurement {
	return MTRClusterFlowMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterFlowMeasurementClass) Alloc() MTRClusterFlowMeasurement {
	rv := objc.Send[MTRClusterFlowMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterFlowMeasurementClass) New() MTRClusterFlowMeasurement {
	rv := objc.Send[MTRClusterFlowMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterFlowMeasurement) Init() MTRClusterFlowMeasurement {
	rv := objc.Send[MTRClusterFlowMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterFlowMeasurement) Autorelease() MTRClusterFlowMeasurement {
	rv := objc.Send[MTRClusterFlowMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterFlowMeasurement creates a new MTRClusterFlowMeasurement instance.
func NewMTRClusterFlowMeasurement() MTRClusterFlowMeasurement {
	return getMTRClusterFlowMeasurementClass().New()
}




