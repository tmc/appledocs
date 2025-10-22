// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	AccuracyRanges() unsafe.Pointer
	SetAccuracyRanges(value unsafe.Pointer)
	MaxMeasuredValue() foundation.Number
	SetMaxMeasuredValue(value foundation.INumber)
	Measured() foundation.Number
	SetMeasured(value foundation.INumber)
	MeasurementType() foundation.Number
	SetMeasurementType(value foundation.INumber)
	MinMeasuredValue() foundation.Number
	SetMinMeasuredValue(value foundation.INumber)
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracystruct/accuracyranges
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct) AccuracyRanges() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("accuracyRanges"))
	return rv
}


// SetAccuracyRanges sets the value of the accuracyRanges property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracystruct/accuracyranges
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct) SetAccuracyRanges(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAccuracyRanges:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracystruct/maxmeasuredvalue
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct) MaxMeasuredValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("maxMeasuredValue"))
	return rv
}


// SetMaxMeasuredValue sets the value of the maxMeasuredValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracystruct/maxmeasuredvalue
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct) SetMaxMeasuredValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxMeasuredValue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracystruct/measured
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct) Measured() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("measured"))
	return rv
}


// SetMeasured sets the value of the measured property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracystruct/measured
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct) SetMeasured(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeasured:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracystruct/measurementtype
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct) MeasurementType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("measurementType"))
	return rv
}


// SetMeasurementType sets the value of the measurementType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracystruct/measurementtype
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct) SetMeasurementType(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeasurementType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracystruct/minmeasuredvalue
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct) MinMeasuredValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("minMeasuredValue"))
	return rv
}


// SetMinMeasuredValue sets the value of the minMeasuredValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustermeasurementaccuracystruct/minmeasuredvalue
func (m_ MTRElectricalEnergyMeasurementClusterMeasurementAccuracyStruct) SetMinMeasuredValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinMeasuredValue:"), value)
}



