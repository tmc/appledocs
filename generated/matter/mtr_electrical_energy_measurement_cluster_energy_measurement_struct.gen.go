// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	EndSystime() objc.IObject /* cross-framework: NSNumber */
	SetEndSystime(value objc.IObject /* cross-framework: NSNumber */)
	EndTimestamp() objc.IObject /* cross-framework: NSNumber */
	SetEndTimestamp(value objc.IObject /* cross-framework: NSNumber */)
	Energy() objc.IObject /* cross-framework: NSNumber */
	SetEnergy(value objc.IObject /* cross-framework: NSNumber */)
	StartSystime() objc.IObject /* cross-framework: NSNumber */
	SetStartSystime(value objc.IObject /* cross-framework: NSNumber */)
	StartTimestamp() objc.IObject /* cross-framework: NSNumber */
	SetStartTimestamp(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/endsystime
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) EndSystime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endSystime"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/endsystime
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) SetEndSystime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndSystime:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/endtimestamp
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) EndTimestamp() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endTimestamp"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/endtimestamp
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) SetEndTimestamp(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTimestamp:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/energy
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) Energy() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("energy"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/energy
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) SetEnergy(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnergy:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/startsystime
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) StartSystime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startSystime"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/startsystime
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) SetStartSystime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartSystime:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/starttimestamp
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) StartTimestamp() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startTimestamp"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrelectricalenergymeasurementclusterenergymeasurementstruct/starttimestamp
func (m_ MTRElectricalEnergyMeasurementClusterEnergyMeasurementStruct) SetStartTimestamp(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTimestamp:"), value)
}
