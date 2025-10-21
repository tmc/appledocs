// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct] class.
var (
	MTRElectricalEnergyMeasurementClusterEnergyMeasurementStructClass     _MTRElectricalEnergyMeasurementClusterEnergyMeasurementStructClass
	MTRElectricalEnergyMeasurementClusterEnergyMeasurementStructClassOnce sync.Once
)

func getMTRElectricalEnergyMeasurementClusterEnergyMeasurementStructClass() _MTRElectricalEnergyMeasurementClusterEnergyMeasurementStructClass {
	MTRElectricalEnergyMeasurementClusterEnergyMeasurementStructClassOnce.Do(func() {
		MTRElectricalEnergyMeasurementClusterEnergyMeasurementStructClass = _MTRElectricalEnergyMeasurementClusterEnergyMeasurementStructClass{objc.GetClass("MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct")}
	})
	return MTRElectricalEnergyMeasurementClusterEnergyMeasurementStructClass
}

type _MTRElectricalEnergyMeasurementClusterEnergyMeasurementStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct] class.
type IMTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct
type MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct struct {
	objectivec.Object
}

// MTRElectricalEnergyMeasurementClusterEnergyMeasurementStructFrom constructs a [MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct] from an unsafe.Pointer.
func MTRElectricalEnergyMeasurementClusterEnergyMeasurementStructFrom(ptr unsafe.Pointer) MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct {
	return MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalEnergyMeasurementClusterEnergyMeasurementStructClass) Alloc() MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalEnergyMeasurementClusterEnergyMeasurementStructClass) New() MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) Init() MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) Autorelease() MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct creates a new MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct instance.
func NewMTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct() MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct {
	return getMTRElectricalEnergyMeasurementClusterEnergyMeasurementStructClass().New()
}




