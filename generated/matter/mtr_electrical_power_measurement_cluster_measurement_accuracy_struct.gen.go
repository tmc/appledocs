// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AccuracyRanges() unsafe.Pointer
	SetAccuracyRanges(value unsafe.Pointer)
	MaxMeasuredValue() objc.IObject /* cross-framework: NSNumber */
	SetMaxMeasuredValue(value objc.IObject /* cross-framework: NSNumber */)
	Measured() objc.IObject /* cross-framework: NSNumber */
	SetMeasured(value objc.IObject /* cross-framework: NSNumber */)
	MeasurementType() objc.IObject /* cross-framework: NSNumber */
	SetMeasurementType(value objc.IObject /* cross-framework: NSNumber */)
	MinMeasuredValue() objc.IObject /* cross-framework: NSNumber */
	SetMinMeasuredValue(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracystruct/accuracyranges
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct) AccuracyRanges() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("accuracyRanges"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracystruct/accuracyranges
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct) SetAccuracyRanges(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAccuracyRanges:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracystruct/maxmeasuredvalue
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct) MaxMeasuredValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxMeasuredValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracystruct/maxmeasuredvalue
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct) SetMaxMeasuredValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxMeasuredValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracystruct/measured
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct) Measured() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("measured"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracystruct/measured
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct) SetMeasured(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeasured:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracystruct/measurementtype
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct) MeasurementType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("measurementType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracystruct/measurementtype
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct) SetMeasurementType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeasurementType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracystruct/minmeasuredvalue
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct) MinMeasuredValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minMeasuredValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementaccuracystruct/minmeasuredvalue
func (m_ MTRElectricalPowerMeasurementClusterMeasurementAccuracyStruct) SetMinMeasuredValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinMeasuredValue:"), value)
}



