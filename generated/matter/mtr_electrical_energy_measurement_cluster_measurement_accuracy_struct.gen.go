// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct] class.
var (
	MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStructClass     _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStructClass
	MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStructClassOnce sync.Once
)

func getMTRElectricalEnergyMeasurementClusterMeasurementAccuracyStructClass() _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStructClass {
	MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStructClassOnce.Do(func() {
		MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStructClass = _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStructClass{objc.GetClass("MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct")}
	})
	return MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStructClass
}

type _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct] class.
type IMTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct
type MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct struct {
	objectivec.Object
}

// MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStructFrom constructs a [MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct] from an unsafe.Pointer.
func MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStructFrom(ptr unsafe.Pointer) MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct {
	return MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStructClass) Alloc() MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStructClass) New() MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct) Init() MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct) Autorelease() MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct creates a new MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct instance.
func NewMTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct() MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct {
	return getMTRElectricalEnergyMeasurementClusterMeasurementAccuracyStructClass().New()
}




