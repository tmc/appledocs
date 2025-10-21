// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterIlluminanceMeasurement] class.
var (
	MTRBaseClusterIlluminanceMeasurementClass     _MTRBaseClusterIlluminanceMeasurementClass
	MTRBaseClusterIlluminanceMeasurementClassOnce sync.Once
)

func getMTRBaseClusterIlluminanceMeasurementClass() _MTRBaseClusterIlluminanceMeasurementClass {
	MTRBaseClusterIlluminanceMeasurementClassOnce.Do(func() {
		MTRBaseClusterIlluminanceMeasurementClass = _MTRBaseClusterIlluminanceMeasurementClass{objc.GetClass("MTRBaseClusterIlluminanceMeasurement")}
	})
	return MTRBaseClusterIlluminanceMeasurementClass
}

type _MTRBaseClusterIlluminanceMeasurementClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterIlluminanceMeasurement] class.
type IMTRBaseClusterIlluminanceMeasurement interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterIlluminanceMeasurement
type MTRBaseClusterIlluminanceMeasurement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterIlluminanceMeasurementFrom constructs a [MTRBaseClusterIlluminanceMeasurement] from an unsafe.Pointer.
func MTRBaseClusterIlluminanceMeasurementFrom(ptr unsafe.Pointer) MTRBaseClusterIlluminanceMeasurement {
	return MTRBaseClusterIlluminanceMeasurement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterIlluminanceMeasurementClass) Alloc() MTRBaseClusterIlluminanceMeasurement {
	rv := objc.Send[MTRBaseClusterIlluminanceMeasurement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterIlluminanceMeasurementClass) New() MTRBaseClusterIlluminanceMeasurement {
	rv := objc.Send[MTRBaseClusterIlluminanceMeasurement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterIlluminanceMeasurement) Init() MTRBaseClusterIlluminanceMeasurement {
	rv := objc.Send[MTRBaseClusterIlluminanceMeasurement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterIlluminanceMeasurement) Autorelease() MTRBaseClusterIlluminanceMeasurement {
	rv := objc.Send[MTRBaseClusterIlluminanceMeasurement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterIlluminanceMeasurement creates a new MTRBaseClusterIlluminanceMeasurement instance.
func NewMTRBaseClusterIlluminanceMeasurement() MTRBaseClusterIlluminanceMeasurement {
	return getMTRBaseClusterIlluminanceMeasurementClass().New()
}




