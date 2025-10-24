// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

/* debug [enums.gen.go]: Generating 3 enums for CoreTelephony */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum CTCellularPlanCapability (2 cases) */
// CTCellularPlanCapability - The type of cellular plan available for an eSIM.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanCapability
type CTCellularPlanCapability uint

const (
	// CTCellularPlanCapabilityDataAndVoice - The cellular plan is available for data and voice.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanCapability/dataAndVoice
	CTCellularPlanCapabilityDataAndVoice CTCellularPlanCapability = 0
	// CTCellularPlanCapabilityDataOnly - The cellular plan is available for data only.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanCapability/dataOnly
	CTCellularPlanCapabilityDataOnly CTCellularPlanCapability = 0
)

/* debug [enums.gen.go]: Processing enum CTCellularDataRestrictedState (3 cases) */
// CTCellularDataRestrictedState - The possible states of the cellular data policy.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularDataRestrictedState
type CTCellularDataRestrictedState uint

const (
	// kCTCellularDataNotRestricted - A state that allows access to cellular data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularDataRestrictedState/notRestricted
	kCTCellularDataNotRestricted CTCellularDataRestrictedState = 0
	// kCTCellularDataRestricted - A state that denies access to cellular data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularDataRestrictedState/restricted
	kCTCellularDataRestricted CTCellularDataRestrictedState = 0
	// kCTCellularDataRestrictedStateUnknown - A state whose access to cellular data is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularDataRestrictedState/restrictedStateUnknown
	kCTCellularDataRestrictedStateUnknown CTCellularDataRestrictedState = 0
)

/* debug [enums.gen.go]: Processing enum CTCellularPlanProvisioningAddPlanResult (4 cases) */
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


