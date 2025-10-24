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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/fixedmax
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) FixedMax() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fixedMax"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/fixedmax
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) SetFixedMax(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFixedMax:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/fixedmin
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) FixedMin() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fixedMin"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/fixedmin
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) SetFixedMin(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFixedMin:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/fixedtypical
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) FixedTypical() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fixedTypical"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/fixedtypical
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) SetFixedTypical(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFixedTypical:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/percentmax
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) PercentMax() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("percentMax"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/percentmax
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) SetPercentMax(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentMax:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/percentmin
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) PercentMin() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("percentMin"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/percentmin
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) SetPercentMin(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentMin:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/percenttypical
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) PercentTypical() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("percentTypical"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/percenttypical
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) SetPercentTypical(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPercentTypical:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/rangemax
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) RangeMax() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rangeMax"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/rangemax
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) SetRangeMax(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRangeMax:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/rangemin
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) RangeMin() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rangeMin"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracyrangestruct/rangemin
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyRangeStruct) SetRangeMin(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRangeMin:"), value)
}



