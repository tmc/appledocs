// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CellularPlanProperties] class.
var (
	CellularPlanPropertiesClass     _CellularPlanPropertiesClass
	CellularPlanPropertiesClassOnce sync.Once
)

func getCellularPlanPropertiesClass() _CellularPlanPropertiesClass {
	CellularPlanPropertiesClassOnce.Do(func() {
		CellularPlanPropertiesClass = _CellularPlanPropertiesClass{objc.GetClass("CTCellularPlanProperties")}
	})
	return CellularPlanPropertiesClass
}

type _CellularPlanPropertiesClass struct {
	class objc.Class
}

// An interface definition for the [CellularPlanProperties] class.
type ICellularPlanProperties interface {
	objectivec.IObject
	AssociatedIccid() string
	SetAssociatedIccid(value string)
	SimCapability() CTCellularPlanCapability
	SetSimCapability(value CTCellularPlanCapability)
	SupportedRegionCodes() corelocation.Region
	SetSupportedRegionCodes(value corelocation.Region)
}

// An object you use for an eSIM.


// An object you use for an eSIM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProperties
type CellularPlanProperties struct {
	objectivec.Object
}

// CellularPlanPropertiesFrom constructs a [CellularPlanProperties] from an unsafe.Pointer.
//
// An object you use for an eSIM.
func CellularPlanPropertiesFrom(ptr unsafe.Pointer) CellularPlanProperties {
	return CellularPlanProperties{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CellularPlanPropertiesClass) Alloc() CellularPlanProperties {
	rv := objc.Send[CellularPlanProperties](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CellularPlanPropertiesClass) New() CellularPlanProperties {
	rv := objc.Send[CellularPlanProperties](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CellularPlanProperties) Init() CellularPlanProperties {
	rv := objc.Send[CellularPlanProperties](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CellularPlanProperties) Autorelease() CellularPlanProperties {
	rv := objc.Send[CellularPlanProperties](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCellularPlanProperties creates a new CellularPlanProperties instance.
func NewCellularPlanProperties() CellularPlanProperties {
	return getCellularPlanPropertiesClass().New()
}



// The integrated circuit card identifier (ICCID) that identifies a SIM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProperties/associatedIccid
func (c_ CellularPlanProperties) AssociatedIccid() string {
	rv := objc.Send[string](c_.ID, objc.Sel("associatedIccid"))
	return rv
}


// The integrated circuit card identifier (ICCID) that identifies a SIM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProperties/associatedIccid
func (c_ CellularPlanProperties) SetAssociatedIccid(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAssociatedIccid:"), objc.String(value))
}


// The available type of cellular plan that your eSIM supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProperties/simCapability
func (c_ CellularPlanProperties) SimCapability() CTCellularPlanCapability {
	rv := objc.Send[CTCellularPlanCapability](c_.ID, objc.Sel("simCapability"))
	return rv
}


// The available type of cellular plan that your eSIM supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProperties/simCapability
func (c_ CellularPlanProperties) SetSimCapability(value CTCellularPlanCapability) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSimCapability:"), value)
}


// The available regions that your eSIM supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanproperties/supportedregioncodes-yhu5
func (c_ CellularPlanProperties) SupportedRegionCodes() corelocation.Region {
	rv := objc.Send[corelocation.Region](c_.ID, objc.Sel("supportedRegionCodes"))
	return rv
}


// The available regions that your eSIM supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctcellularplanproperties/supportedregioncodes-yhu5
func (c_ CellularPlanProperties) SetSupportedRegionCodes(value corelocation.Region) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedRegionCodes:"), value)
}



