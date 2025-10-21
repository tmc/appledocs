// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct] class.
var (
	MTRElectricalPowerMeasurementClusterMeasurementAccuracyStructClass     _MTRElectricalPowerMeasurementClusterMeasurementAccuracyStructClass
	MTRElectricalPowerMeasurementClusterMeasurementAccuracyStructClassOnce sync.Once
)

func getMTRElectricalPowerMeasurementClusterMeasurementAccuracyStructClass() _MTRElectricalPowerMeasurementClusterMeasurementAccuracyStructClass {
	MTRElectricalPowerMeasurementClusterMeasurementAccuracyStructClassOnce.Do(func() {
		MTRElectricalPowerMeasurementClusterMeasurementAccuracyStructClass = _MTRElectricalPowerMeasurementClusterMeasurementAccuracyStructClass{objc.GetClass("MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct")}
	})
	return MTRElectricalPowerMeasurementClusterMeasurementAccuracyStructClass
}

type _MTRElectricalPowerMeasurementClusterMeasurementAccuracyStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct] class.
type IMTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct
type MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct struct {
	objectivec.Object
}

// MTRElectricalPowerMeasurementClusterMeasurementAccuracyStructFrom constructs a [MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct] from an unsafe.Pointer.
func MTRElectricalPowerMeasurementClusterMeasurementAccuracyStructFrom(ptr unsafe.Pointer) MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct {
	return MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalPowerMeasurementClusterMeasurementAccuracyStructClass) Alloc() MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalPowerMeasurementClusterMeasurementAccuracyStructClass) New() MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct) Init() MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct) Autorelease() MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct creates a new MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct instance.
func NewMTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct() MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct {
	return getMTRElectricalPowerMeasurementClusterMeasurementAccuracyStructClass().New()
}




