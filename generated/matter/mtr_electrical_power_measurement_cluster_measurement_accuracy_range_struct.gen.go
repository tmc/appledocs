// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/percentmax
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) PercentMax() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("percentMax"))
	return rv
}


// SetPercentMax sets the value of the percentMax property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/percentmax
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) SetPercentMax(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentMax:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/rangemin
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) RangeMin() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rangeMin"))
	return rv
}


// SetRangeMin sets the value of the rangeMin property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/rangemin
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) SetRangeMin(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRangeMin:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/fixedtypical
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) FixedTypical() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fixedTypical"))
	return rv
}


// SetFixedTypical sets the value of the fixedTypical property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/fixedtypical
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) SetFixedTypical(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFixedTypical:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/percenttypical
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) PercentTypical() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("percentTypical"))
	return rv
}


// SetPercentTypical sets the value of the percentTypical property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/percenttypical
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) SetPercentTypical(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentTypical:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/percentmin
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) PercentMin() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("percentMin"))
	return rv
}


// SetPercentMin sets the value of the percentMin property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/percentmin
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) SetPercentMin(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentMin:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/fixedmin
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) FixedMin() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fixedMin"))
	return rv
}


// SetFixedMin sets the value of the fixedMin property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/fixedmin
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) SetFixedMin(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFixedMin:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/fixedmax
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) FixedMax() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fixedMax"))
	return rv
}


// SetFixedMax sets the value of the fixedMax property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/fixedmax
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) SetFixedMax(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFixedMax:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/rangemax
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) RangeMax() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rangeMax"))
	return rv
}


// SetRangeMax sets the value of the rangeMax property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/rangemax
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) SetRangeMax(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRangeMax:"), value)
}



