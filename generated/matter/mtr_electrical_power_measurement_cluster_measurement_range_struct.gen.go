// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRElectricalPowerMeasurementClusterMeasurementRangeStruct] class.
var (
	MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass     _MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass
	MTRElectricalPowerMeasurementClusterMeasurementRangeStructClassOnce sync.Once
)

func getMTRElectricalPowerMeasurementClusterMeasurementRangeStructClass() _MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass {
	MTRElectricalPowerMeasurementClusterMeasurementRangeStructClassOnce.Do(func() {
		MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass = _MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass{objc.GetClass("MTRElectricalPowerMeasurementClusterMeasurementRangeStruct")}
	})
	return MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass
}

type _MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalPowerMeasurementClusterMeasurementRangeStruct] class.
type IMTRElectricalPowerMeasurementClusterMeasurementRangeStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalPowerMeasurementClusterMeasurementRangeStruct
type MTRElectricalPowerMeasurementClusterMeasurementRangeStruct struct {
	objectivec.Object
}

// MTRElectricalPowerMeasurementClusterMeasurementRangeStructFrom constructs a [MTRElectricalPowerMeasurementClusterMeasurementRangeStruct] from an unsafe.Pointer.
func MTRElectricalPowerMeasurementClusterMeasurementRangeStructFrom(ptr unsafe.Pointer) MTRElectricalPowerMeasurementClusterMeasurementRangeStruct {
	return MTRElectricalPowerMeasurementClusterMeasurementRangeStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass) Alloc() MTRElectricalPowerMeasurementClusterMeasurementRangeStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementRangeStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass) New() MTRElectricalPowerMeasurementClusterMeasurementRangeStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementRangeStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) Init() MTRElectricalPowerMeasurementClusterMeasurementRangeStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementRangeStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) Autorelease() MTRElectricalPowerMeasurementClusterMeasurementRangeStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementRangeStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalPowerMeasurementClusterMeasurementRangeStruct creates a new MTRElectricalPowerMeasurementClusterMeasurementRangeStruct instance.
func NewMTRElectricalPowerMeasurementClusterMeasurementRangeStruct() MTRElectricalPowerMeasurementClusterMeasurementRangeStruct {
	return getMTRElectricalPowerMeasurementClusterMeasurementRangeStructClass().New()
}




