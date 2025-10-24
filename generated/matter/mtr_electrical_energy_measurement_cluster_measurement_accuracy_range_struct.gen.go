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
	// properties:
	FixedMax() objc.IObject /* cross-framework: NSNumber */
	SetFixedMax(value objc.IObject /* cross-framework: NSNumber */)
	FixedMin() objc.IObject /* cross-framework: NSNumber */
	SetFixedMin(value objc.IObject /* cross-framework: NSNumber */)
	FixedTypical() objc.IObject /* cross-framework: NSNumber */
	SetFixedTypical(value objc.IObject /* cross-framework: NSNumber */)
	PercentMax() objc.IObject /* cross-framework: NSNumber */
	SetPercentMax(value objc.IObject /* cross-framework: NSNumber */)
	PercentMin() objc.IObject /* cross-framework: NSNumber */
	SetPercentMin(value objc.IObject /* cross-framework: NSNumber */)
	PercentTypical() objc.IObject /* cross-framework: NSNumber */
	SetPercentTypical(value objc.IObject /* cross-framework: NSNumber */)
	RangeMax() objc.IObject /* cross-framework: NSNumber */
	SetRangeMax(value objc.IObject /* cross-framework: NSNumber */)
	RangeMin() objc.IObject /* cross-framework: NSNumber */
	SetRangeMin(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/fixedMax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) FixedMax() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fixedMax"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/fixedMax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetFixedMax(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFixedMax:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/fixedMin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) FixedMin() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fixedMin"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/fixedMin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetFixedMin(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFixedMin:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/fixedTypical
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) FixedTypical() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fixedTypical"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/fixedTypical
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetFixedTypical(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFixedTypical:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/percentMax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) PercentMax() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("percentMax"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/percentMax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetPercentMax(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentMax:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/percentMin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) PercentMin() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("percentMin"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/percentMin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetPercentMin(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentMin:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/percentTypical
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) PercentTypical() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("percentTypical"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/percentTypical
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetPercentTypical(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentTypical:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/rangeMax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) RangeMax() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rangeMax"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/rangeMax
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetRangeMax(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRangeMax:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/rangeMin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) RangeMin() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rangeMin"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct/rangeMin
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyRangeStruct) SetRangeMin(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRangeMin:"), value)
}



