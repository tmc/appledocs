// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

// Enum types and constants
// CTCellularDataRestrictedState - The possible states of the cellular data policy.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularDataRestrictedState
type CTCellularDataRestrictedState uint

// CTCellularPlanCapability - The type of cellular plan available for an eSIM.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanCapability
type CTCellularPlanCapability uint

const (
	// CTCellularPlanCapabilityDataOnly - The cellular plan is available for data only.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanCapability/dataOnly
	CTCellularPlanCapabilityDataOnly CTCellularPlanCapability = 0
)

// CTCellularPlanProvisioningAddPlanResult - The result from attempting to provision an eSIM.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioningAddPlanResult
type CTCellularPlanProvisioningAddPlanResult uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioningAddPlanResult/cancel
	CTCellularPlanProvisioningAddPlanResultCancel CTCellularPlanProvisioningAddPlanResult = 0
	// CTCellularPlanProvisioningAddPlanResultFail - The requested eSIM provisioning failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioningAddPlanResult/fail
	CTCellularPlanProvisioningAddPlanResultFail CTCellularPlanProvisioningAddPlanResult = 0
	// CTCellularPlanProvisioningAddPlanResultSuccess - The requested eSIM provisioning succeeded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioningAddPlanResult/success
	CTCellularPlanProvisioningAddPlanResultSuccess CTCellularPlanProvisioningAddPlanResult = 0
	// CTCellularPlanProvisioningAddPlanResultUnknown - The result of the requested eSIM provisioning is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanProvisioningAddPlanResult/unknown
	CTCellularPlanProvisioningAddPlanResultUnknown CTCellularPlanProvisioningAddPlanResult = 0
)


