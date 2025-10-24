// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterIlluminanceMeasurement] class.
var (
	MTRClusterIlluminanceMeasurementClass     _MTRClusterIlluminanceMeasurementClass
	MTRClusterIlluminanceMeasurementClassOnce sync.Once
)

func getMTRClusterIlluminanceMeasurementClass() _MTRClusterIlluminanceMeasurementClass {
	MTRClusterIlluminanceMeasurementClassOnce.Do(func() {
		MTRClusterIlluminanceMeasurementClass = _MTRClusterIlluminanceMeasurementClass{objc.GetClass("MTRClusterIlluminanceMeasurement")}
	})
	return MTRClusterIlluminanceMeasurementClass
}

type _MTRClusterIlluminanceMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterIlluminanceMeasurement] class.
type IMTRClusterIlluminanceMeasurement interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterIlluminanceMeasurement
type MTRClusterIlluminanceMeasurement struct {
	MTRGenericCluster
}

// MTRClusterIlluminanceMeasurementFrom constructs a [MTRClusterIlluminanceMeasurement] from an unsafe.Pointer.
func MTRClusterIlluminanceMeasurementFrom(ptr unsafe.Pointer) MTRClusterIlluminanceMeasurement {
	return MTRClusterIlluminanceMeasurement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterIlluminanceMeasurementClass) Alloc() MTRClusterIlluminanceMeasurement {
	rv := objc.Send[MTRClusterIlluminanceMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterIlluminanceMeasurementClass) New() MTRClusterIlluminanceMeasurement {
	rv := objc.Send[MTRClusterIlluminanceMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterIlluminanceMeasurement) Init() MTRClusterIlluminanceMeasurement {
	rv := objc.Send[MTRClusterIlluminanceMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterIlluminanceMeasurement) Autorelease() MTRClusterIlluminanceMeasurement {
	rv := objc.Send[MTRClusterIlluminanceMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterIlluminanceMeasurement creates a new MTRClusterIlluminanceMeasurement instance.
func NewMTRClusterIlluminanceMeasurement() MTRClusterIlluminanceMeasurement {
	return getMTRClusterIlluminanceMeasurementClass().New()
}
