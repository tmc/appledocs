// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct] class.
var (
	MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStructClass     _MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStructClass
	MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStructClassOnce sync.Once
)

func getMTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStructClass() _MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStructClass {
	MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStructClassOnce.Do(func() {
		MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStructClass = _MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStructClass{objc.GetClass("MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct")}
	})
	return MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStructClass
}

type _MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct] class.
type IMTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct
type MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct struct {
	objectivec.Object
}

// MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStructFrom constructs a [MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct] from an unsafe.Pointer.
func MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStructFrom(ptr unsafe.Pointer) MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct {
	return MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStructClass) Alloc() MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStructClass) New() MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) Init() MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) Autorelease() MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct creates a new MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct instance.
func NewMTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct() MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct {
	return getMTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStructClass().New()
}




