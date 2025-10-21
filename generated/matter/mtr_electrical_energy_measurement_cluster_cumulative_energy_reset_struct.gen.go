// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct] class.
var (
	MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStructClass     _MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStructClass
	MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStructClassOnce sync.Once
)

func getMTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStructClass() _MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStructClass {
	MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStructClassOnce.Do(func() {
		MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStructClass = _MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStructClass{objc.GetClass("MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct")}
	})
	return MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStructClass
}

type _MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct] class.
type IMTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct
type MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct struct {
	objectivec.Object
}

// MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStructFrom constructs a [MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct] from an unsafe.Pointer.
func MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStructFrom(ptr unsafe.Pointer) MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct {
	return MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStructClass) Alloc() MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStructClass) New() MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct) Init() MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct) Autorelease() MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct {
	rv := objc.Send[MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct creates a new MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct instance.
func NewMTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct() MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct {
	return getMTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustercumulativeenergyresetstruct/importedresetsystime
func (m_ MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct) ImportedResetSystime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("importedResetSystime"))
	return rv
}


// SetImportedResetSystime sets the value of the importedResetSystime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustercumulativeenergyresetstruct/importedresetsystime
func (m_ MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct) SetImportedResetSystime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImportedResetSystime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustercumulativeenergyresetstruct/importedresettimestamp
func (m_ MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct) ImportedResetTimestamp() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("importedResetTimestamp"))
	return rv
}


// SetImportedResetTimestamp sets the value of the importedResetTimestamp property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustercumulativeenergyresetstruct/importedresettimestamp
func (m_ MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct) SetImportedResetTimestamp(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImportedResetTimestamp:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustercumulativeenergyresetstruct/exportedresetsystime
func (m_ MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct) ExportedResetSystime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("exportedResetSystime"))
	return rv
}


// SetExportedResetSystime sets the value of the exportedResetSystime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustercumulativeenergyresetstruct/exportedresetsystime
func (m_ MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct) SetExportedResetSystime(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExportedResetSystime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustercumulativeenergyresetstruct/exportedresettimestamp
func (m_ MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct) ExportedResetTimestamp() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("exportedResetTimestamp"))
	return rv
}


// SetExportedResetTimestamp sets the value of the exportedResetTimestamp property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclustercumulativeenergyresetstruct/exportedresettimestamp
func (m_ MTRElectricalEnergyMeasurementClusterCumulativeEnergyResetStruct) SetExportedResetTimestamp(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExportedResetTimestamp:"), value)
}



