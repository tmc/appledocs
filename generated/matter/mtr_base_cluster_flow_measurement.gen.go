// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterFlowMeasurement] class.
var (
	MTRBaseClusterFlowMeasurementClass     _MTRBaseClusterFlowMeasurementClass
	MTRBaseClusterFlowMeasurementClassOnce sync.Once
)

func getMTRBaseClusterFlowMeasurementClass() _MTRBaseClusterFlowMeasurementClass {
	MTRBaseClusterFlowMeasurementClassOnce.Do(func() {
		MTRBaseClusterFlowMeasurementClass = _MTRBaseClusterFlowMeasurementClass{objc.GetClass("MTRBaseClusterFlowMeasurement")}
	})
	return MTRBaseClusterFlowMeasurementClass
}

type _MTRBaseClusterFlowMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterFlowMeasurement] class.
type IMTRBaseClusterFlowMeasurement interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterFlowMeasurement
type MTRBaseClusterFlowMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterFlowMeasurementFrom constructs a [MTRBaseClusterFlowMeasurement] from an unsafe.Pointer.
func MTRBaseClusterFlowMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterFlowMeasurement {
	return MTRBaseClusterFlowMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterFlowMeasurementClass) Alloc() MTRBaseClusterFlowMeasurement {
	rv := objc.Send[MTRBaseClusterFlowMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterFlowMeasurementClass) New() MTRBaseClusterFlowMeasurement {
	rv := objc.Send[MTRBaseClusterFlowMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterFlowMeasurement) Init() MTRBaseClusterFlowMeasurement {
	rv := objc.Send[MTRBaseClusterFlowMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterFlowMeasurement) Autorelease() MTRBaseClusterFlowMeasurement {
	rv := objc.Send[MTRBaseClusterFlowMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterFlowMeasurement creates a new MTRBaseClusterFlowMeasurement instance.
func NewMTRBaseClusterFlowMeasurement() MTRBaseClusterFlowMeasurement {
	return getMTRBaseClusterFlowMeasurementClass().New()
}




