// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/startsystime
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) StartSystime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("startSystime"))
	return rv
}


// SetStartSystime sets the value of the startSystime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/startsystime
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) SetStartSystime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartSystime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/endsystime
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) EndSystime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endSystime"))
	return rv
}


// SetEndSystime sets the value of the endSystime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/endsystime
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) SetEndSystime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndSystime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/starttimestamp
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) StartTimestamp() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("startTimestamp"))
	return rv
}


// SetStartTimestamp sets the value of the startTimestamp property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/starttimestamp
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) SetStartTimestamp(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTimestamp:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/energy
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) Energy() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("energy"))
	return rv
}


// SetEnergy sets the value of the energy property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/energy
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) SetEnergy(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnergy:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/endtimestamp
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) EndTimestamp() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endTimestamp"))
	return rv
}


// SetEndTimestamp sets the value of the endTimestamp property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/endtimestamp
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) SetEndTimestamp(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTimestamp:"), value)
}



