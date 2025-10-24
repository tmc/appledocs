//go:build darwin && ios

// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CellularPlanProvisioning


// Starts the provisioning process with optional properties for the specified eSIM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioning/addPlan(request:properties:completionHandler:)
func (c_ CellularPlanProvisioning) AddPlanWithRequestPropertiesCompletionHandler(request ICTCellularPlanProvisioningRequest, properties ICTCellularPlanProperties, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addPlanWithRequest:properties:completionHandler:"), request, properties, completionHandler)
}

// Starts the provisioning process for a specified eSIM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioning/addPlan(with:completionHandler:)
func (c_ CellularPlanProvisioning) AddPlanWithCompletionHandler(request ICTCellularPlanProvisioningRequest, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addPlanWith:completionHandler:"), request, completionHandler)
}

// Indicates whether the device supports eSIM and the activation policy allows eSIM installation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioning/supportsCellularPlan()
func (c_ CellularPlanProvisioning) SupportsCellularPlan() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsCellularPlan"))
	return rv
}

// Updates the capability and region availability for an eSIM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioning/update(_:completionHandler:)
func (c_ CellularPlanProvisioning) UpdateCellularPlanPropertiesCompletionHandler(properties ICTCellularPlanProperties, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("updateCellularPlanProperties:completionHandler:"), properties, completionHandler)
}

// iOS-only properties

// A Boolean value that indicates whether the device has hardware eSIM support.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioning/supportsEmbeddedSIM
func (c_ CellularPlanProvisioning) SupportsEmbeddedSIM() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsEmbeddedSIM"))
	return rv
}





