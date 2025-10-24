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
	// properties:
	EndSystime() objc.IObject /* cross-framework: NSNumber */
	SetEndSystime(value objc.IObject /* cross-framework: NSNumber */)
	EndTimestamp() objc.IObject /* cross-framework: NSNumber */
	SetEndTimestamp(value objc.IObject /* cross-framework: NSNumber */)
	Max() objc.IObject /* cross-framework: NSNumber */
	SetMax(value objc.IObject /* cross-framework: NSNumber */)
	MaxSystime() objc.IObject /* cross-framework: NSNumber */
	SetMaxSystime(value objc.IObject /* cross-framework: NSNumber */)
	MaxTimestamp() objc.IObject /* cross-framework: NSNumber */
	SetMaxTimestamp(value objc.IObject /* cross-framework: NSNumber */)
	MeasurementType() objc.IObject /* cross-framework: NSNumber */
	SetMeasurementType(value objc.IObject /* cross-framework: NSNumber */)
	Min() objc.IObject /* cross-framework: NSNumber */
	SetMin(value objc.IObject /* cross-framework: NSNumber */)
	MinSystime() objc.IObject /* cross-framework: NSNumber */
	SetMinSystime(value objc.IObject /* cross-framework: NSNumber */)
	MinTimestamp() objc.IObject /* cross-framework: NSNumber */
	SetMinTimestamp(value objc.IObject /* cross-framework: NSNumber */)
	StartSystime() objc.IObject /* cross-framework: NSNumber */
	SetStartSystime(value objc.IObject /* cross-framework: NSNumber */)
	StartTimestamp() objc.IObject /* cross-framework: NSNumber */
	SetStartTimestamp(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/endsystime
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) EndSystime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endSystime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/endsystime
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetEndSystime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndSystime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/endtimestamp
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) EndTimestamp() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endTimestamp"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/endtimestamp
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetEndTimestamp(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTimestamp:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/max
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) Max() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("max"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/max
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetMax(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMax:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/maxsystime
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) MaxSystime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxSystime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/maxsystime
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetMaxSystime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxSystime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/maxtimestamp
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) MaxTimestamp() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxTimestamp"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/maxtimestamp
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetMaxTimestamp(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTimestamp:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/measurementtype
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) MeasurementType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("measurementType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/measurementtype
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetMeasurementType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMeasurementType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/min
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) Min() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("min"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/min
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetMin(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMin:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/minsystime
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) MinSystime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minSystime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/minsystime
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetMinSystime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinSystime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/mintimestamp
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) MinTimestamp() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minTimestamp"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/mintimestamp
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetMinTimestamp(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinTimestamp:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/startsystime
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) StartSystime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startSystime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/startsystime
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetStartSystime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartSystime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/starttimestamp
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) StartTimestamp() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startTimestamp"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalpowermeasurementclustermeasurementrangestruct/starttimestamp
func (m_ MTRElectricalPowerMeasurementClusterMeasurementRangeStruct) SetStartTimestamp(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTimestamp:"), value)
}



