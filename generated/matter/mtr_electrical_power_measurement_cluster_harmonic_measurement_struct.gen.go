// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct] class.
var (
	MTRElectricalPowerMeasurementClusterHarmonicMeasurementStructClass     _MTRElectricalPowerMeasurementClusterHarmonicMeasurementStructClass
	MTRElectricalPowerMeasurementClusterHarmonicMeasurementStructClassOnce sync.Once
)

func getMTRElectricalPowerMeasurementClusterHarmonicMeasurementStructClass() _MTRElectricalPowerMeasurementClusterHarmonicMeasurementStructClass {
	MTRElectricalPowerMeasurementClusterHarmonicMeasurementStructClassOnce.Do(func() {
		MTRElectricalPowerMeasurementClusterHarmonicMeasurementStructClass = _MTRElectricalPowerMeasurementClusterHarmonicMeasurementStructClass{objc.GetClass("MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct")}
	})
	return MTRElectricalPowerMeasurementClusterHarmonicMeasurementStructClass
}

type _MTRElectricalPowerMeasurementClusterHarmonicMeasurementStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct] class.
type IMTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct
type MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct struct {
	objectivec.Object
}

// MTRElectricalPowerMeasurementClusterHarmonicMeasurementStructFrom constructs a [MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct] from an unsafe.Pointer.
func MTRElectricalPowerMeasurementClusterHarmonicMeasurementStructFrom(ptr unsafe.Pointer) MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct {
	return MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalPowerMeasurementClusterHarmonicMeasurementStructClass) Alloc() MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalPowerMeasurementClusterHarmonicMeasurementStructClass) New() MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct) Init() MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct) Autorelease() MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct creates a new MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct instance.
func NewMTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct() MTRElectricalPowerMeasurementClusterHarmonicMeasurementStruct {
	return getMTRElectricalPowerMeasurementClusterHarmonicMeasurementStructClass().New()
}




