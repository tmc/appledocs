// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct] class.
var (
	MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass     _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass
	MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClassOnce sync.Once
)

func getMTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass() _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass {
	MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClassOnce.Do(func() {
		MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass = _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass{objc.GetClass("MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct")}
	})
	return MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass
}

type _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct] class.
type IMTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct interface {
	objectivec.IObject
	FixedMax() foundation.Number
	SetFixedMax(value foundation.INumber)
	FixedMin() foundation.Number
	SetFixedMin(value foundation.INumber)
	FixedTypical() foundation.Number
	SetFixedTypical(value foundation.INumber)
	PercentMax() foundation.Number
	SetPercentMax(value foundation.INumber)
	PercentMin() foundation.Number
	SetPercentMin(value foundation.INumber)
	PercentTypical() foundation.Number
	SetPercentTypical(value foundation.INumber)
	RangeMax() foundation.Number
	SetRangeMax(value foundation.INumber)
	RangeMin() foundation.Number
	SetRangeMin(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct
type MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct struct {
	objectivec.Object
}

// MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructFrom constructs a [MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct] from an unsafe.Pointer.
func MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructFrom(ptr unsafe.Pointer) MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct {
	return MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass) Alloc() MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass) New() MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) Init() MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) Autorelease() MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct creates a new MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct instance.
func NewMTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct() MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct {
	return getMTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/fixedMax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) FixedMax() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fixedMax"))
	return rv
}


// SetFixedMax sets the value of the fixedMax property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/fixedMax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetFixedMax(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFixedMax:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/fixedMin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) FixedMin() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fixedMin"))
	return rv
}


// SetFixedMin sets the value of the fixedMin property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/fixedMin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetFixedMin(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFixedMin:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/fixedTypical
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) FixedTypical() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fixedTypical"))
	return rv
}


// SetFixedTypical sets the value of the fixedTypical property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/fixedTypical
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetFixedTypical(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFixedTypical:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/percentMax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) PercentMax() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("percentMax"))
	return rv
}


// SetPercentMax sets the value of the percentMax property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/percentMax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetPercentMax(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentMax:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/percentMin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) PercentMin() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("percentMin"))
	return rv
}


// SetPercentMin sets the value of the percentMin property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/percentMin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetPercentMin(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentMin:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/percentTypical
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) PercentTypical() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("percentTypical"))
	return rv
}


// SetPercentTypical sets the value of the percentTypical property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/percentTypical
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetPercentTypical(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentTypical:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/rangeMax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) RangeMax() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rangeMax"))
	return rv
}


// SetRangeMax sets the value of the rangeMax property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/rangeMax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetRangeMax(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRangeMax:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/rangeMin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) RangeMin() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rangeMin"))
	return rv
}


// SetRangeMin sets the value of the rangeMin property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/rangeMin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetRangeMin(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRangeMin:"), value)
}



