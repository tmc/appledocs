// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRElectricalPowerMeasurementClusterMeasurementRangeStruct] class.
var (
	MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass     _MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass
	MTRElectricalPowerMeasurementClusterMeasurementRangeStructClassOnce sync.Once
)

func getMTRElectricalPowerMeasurementClusterMeasurementRangeStructClass() _MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass {
	MTRElectricalPowerMeasurementClusterMeasurementRangeStructClassOnce.Do(func() {
		MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass = _MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass{objc.GetClass("MTRElectricalPowerMeasurementClusterMeasurementRangeStruct")}
	})
	return MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass
}

type _MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalPowerMeasurementClusterMeasurementRangeStruct] class.
type IMTRElectricalPowerMeasurementClusterMeasurementRangeStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalPowerMeasurementClusterMeasurementRangeStruct
type MTRElectricalPowerMeasurementClusterMeasurementRangeStruct struct {
	objectivec.Object
}

// MTRElectricalPowerMeasurementClusterMeasurementRangeStructFrom constructs a [MTRElectricalPowerMeasurementClusterMeasurementRangeStruct] from an unsafe.Pointer.
func MTRElectricalPowerMeasurementClusterMeasurementRangeStructFrom(ptr unsafe.Pointer) MTRElectricalPowerMeasurementClusterMeasurementRangeStruct {
	return MTRElectricalPowerMeasurementClusterMeasurementRangeStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass) Alloc() MTRElectricalPowerMeasurementClusterMeasurementRangeStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementRangeStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalPowerMeasurementClusterMeasurementRangeStructClass) New() MTRElectricalPowerMeasurementClusterMeasurementRangeStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementRangeStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) Init() MTRElectricalPowerMeasurementClusterMeasurementRangeStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementRangeStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) Autorelease() MTRElectricalPowerMeasurementClusterMeasurementRangeStruct {
	rv := objc.Send[MTRElectricalPowerMeasurementClusterMeasurementRangeStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalPowerMeasurementClusterMeasurementRangeStruct creates a new MTRElectricalPowerMeasurementClusterMeasurementRangeStruct instance.
func NewMTRElectricalPowerMeasurementClusterMeasurementRangeStruct() MTRElectricalPowerMeasurementClusterMeasurementRangeStruct {
	return getMTRElectricalPowerMeasurementClusterMeasurementRangeStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/measurementtype
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) MeasurementType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("measurementType"))
	return rv
}


// SetMeasurementType sets the value of the measurementType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/measurementtype
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetMeasurementType(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeasurementType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/maxsystime
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) MaxSystime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("maxSystime"))
	return rv
}


// SetMaxSystime sets the value of the maxSystime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/maxsystime
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetMaxSystime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxSystime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/max
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) Max() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("max"))
	return rv
}


// SetMax sets the value of the max property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/max
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetMax(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMax:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/endtimestamp
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) EndTimestamp() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endTimestamp"))
	return rv
}


// SetEndTimestamp sets the value of the endTimestamp property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/endtimestamp
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetEndTimestamp(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTimestamp:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/starttimestamp
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) StartTimestamp() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("startTimestamp"))
	return rv
}


// SetStartTimestamp sets the value of the startTimestamp property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/starttimestamp
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetStartTimestamp(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTimestamp:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/minsystime
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) MinSystime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("minSystime"))
	return rv
}


// SetMinSystime sets the value of the minSystime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/minsystime
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetMinSystime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinSystime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/maxtimestamp
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) MaxTimestamp() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("maxTimestamp"))
	return rv
}


// SetMaxTimestamp sets the value of the maxTimestamp property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/maxtimestamp
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetMaxTimestamp(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTimestamp:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/startsystime
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) StartSystime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("startSystime"))
	return rv
}


// SetStartSystime sets the value of the startSystime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/startsystime
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetStartSystime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartSystime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/min
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) Min() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("min"))
	return rv
}


// SetMin sets the value of the min property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/min
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetMin(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMin:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/mintimestamp
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) MinTimestamp() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("minTimestamp"))
	return rv
}


// SetMinTimestamp sets the value of the minTimestamp property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/mintimestamp
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetMinTimestamp(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinTimestamp:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/endsystime
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) EndSystime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endSystime"))
	return rv
}


// SetEndSystime sets the value of the endSystime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/endsystime
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetEndSystime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndSystime:"), value)
}



