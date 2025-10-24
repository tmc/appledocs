// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterPressureMeasurement] class.
var (
	MTRClusterPressureMeasurementClass     _MTRClusterPressureMeasurementClass
	MTRClusterPressureMeasurementClassOnce sync.Once
)

func getMTRClusterPressureMeasurementClass() _MTRClusterPressureMeasurementClass {
	MTRClusterPressureMeasurementClassOnce.Do(func() {
		MTRClusterPressureMeasurementClass = _MTRClusterPressureMeasurementClass{objc.GetClass("MTRClusterPressureMeasurement")}
	})
	return MTRClusterPressureMeasurementClass
}

type _MTRClusterPressureMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterPressureMeasurement] class.
type IMTRClusterPressureMeasurement interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPressureMeasurement
type MTRClusterPressureMeasurement struct {
	MTRGenericCluster
}

// MTRClusterPressureMeasurementFrom constructs a [MTRClusterPressureMeasurement] from an unsafe.Pointer.
func MTRClusterPressureMeasurementFrom(ptr unsafe.Pointer) MTRClusterPressureMeasurement {
	return MTRClusterPressureMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterPressureMeasurementClass) Alloc() MTRClusterPressureMeasurement {
	rv := objc.Send[MTRClusterPressureMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterPressureMeasurementClass) New() MTRClusterPressureMeasurement {
	rv := objc.Send[MTRClusterPressureMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterPressureMeasurement) Init() MTRClusterPressureMeasurement {
	rv := objc.Send[MTRClusterPressureMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterPressureMeasurement) Autorelease() MTRClusterPressureMeasurement {
	rv := objc.Send[MTRClusterPressureMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterPressureMeasurement creates a new MTRClusterPressureMeasurement instance.
func NewMTRClusterPressureMeasurement() MTRClusterPressureMeasurement {
	return getMTRClusterPressureMeasurementClass().New()
}




