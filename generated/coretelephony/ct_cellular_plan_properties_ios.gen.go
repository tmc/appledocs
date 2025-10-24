//go:build darwin && ios

// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CellularPlanProperties


// iOS-only properties

// The integrated circuit card identifier (ICCID) that identifies a SIM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProperties/associatedIccid
func (c_ CellularPlanProperties) AssociatedIccid() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("associatedIccid"))
	return rv
}
func (c_ CellularPlanProperties) SetAssociatedIccid(value objc.IObject /* cross-framework: NSString */) {
	c_.ID.Send(objc.RegisterName("setAssociatedIccid:"), value)
}

// The available type of cellular plan that your eSIM supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProperties/simCapability
func (c_ CellularPlanProperties) SimCapability() CellularPlanCapability {
	rv := objc.Send[CellularPlanCapability](c_.ID, objc.Sel("simCapability"))
	return rv
}
func (c_ CellularPlanProperties) SetSimCapability(value CellularPlanCapability) {
	c_.ID.Send(objc.RegisterName("setSimCapability:"), value)
}

// The available regions that your eSIM supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProperties/supportedRegionCodes-5elox
func (c_ CellularPlanProperties) SupportedRegionCodes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("supportedRegionCodes"))
	return rv
}
func (c_ CellularPlanProperties) SetSupportedRegionCodes(value []string) {
	c_.ID.Send(objc.RegisterName("setSupportedRegionCodes:"), value)
}





