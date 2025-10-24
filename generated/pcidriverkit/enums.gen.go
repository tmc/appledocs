// Code generated from Apple documentation for PCIDriverKit. DO NOT EDIT.

package pcidriverkit

/* debug [enums.gen.go]: Generating 9 enums for PCIDriverKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum IOPCIBARType (5 cases) */
// IOPCIBARType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIBARType
type IOPCIBARType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIBARType/kPCIBARTypeIO
	kPCIBARTypeIO IOPCIBARType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIBARType/kPCIBARTypeM32
	kPCIBARTypeM32 IOPCIBARType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIBARType/kPCIBARTypeM32PF
	kPCIBARTypeM32PF IOPCIBARType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIBARType/kPCIBARTypeM64
	kPCIBARTypeM64 IOPCIBARType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIBARType/kPCIBARTypeM64PF
	kPCIBARTypeM64PF IOPCIBARType = 0
)

/* debug [enums.gen.go]: Processing enum IOPCILinkSpeed (5 cases) */
// IOPCILinkSpeed enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCILinkSpeed
type IOPCILinkSpeed uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCILinkSpeed/kPCILinkSpeed_16_GTs
	kPCILinkSpeed_16_GTs IOPCILinkSpeed = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCILinkSpeed/kPCILinkSpeed_2_5_GTs
	kPCILinkSpeed_2_5_GTs IOPCILinkSpeed = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCILinkSpeed/kPCILinkSpeed_32_GTs
	kPCILinkSpeed_32_GTs IOPCILinkSpeed = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCILinkSpeed/kPCILinkSpeed_5_GTs
	kPCILinkSpeed_5_GTs IOPCILinkSpeed = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCILinkSpeed/kPCILinkSpeed_8_GTs
	kPCILinkSpeed_8_GTs IOPCILinkSpeed = 0
)

/* debug [enums.gen.go]: Processing enum IOPCIMemoryRange (7 cases) */
// IOPCIMemoryRange enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIMemoryRange
type IOPCIMemoryRange uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIMemoryRange/kPCIMemoryRangeBAR0
	kPCIMemoryRangeBAR0 IOPCIMemoryRange = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIMemoryRange/kPCIMemoryRangeBAR1
	kPCIMemoryRangeBAR1 IOPCIMemoryRange = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIMemoryRange/kPCIMemoryRangeBAR2
	kPCIMemoryRangeBAR2 IOPCIMemoryRange = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIMemoryRange/kPCIMemoryRangeBAR3
	kPCIMemoryRangeBAR3 IOPCIMemoryRange = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIMemoryRange/kPCIMemoryRangeBAR4
	kPCIMemoryRangeBAR4 IOPCIMemoryRange = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIMemoryRange/kPCIMemoryRangeBAR5
	kPCIMemoryRangeBAR5 IOPCIMemoryRange = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCIMemoryRange/kPCIMemoryRangeExpansionROM
	kPCIMemoryRangeExpansionROM IOPCIMemoryRange = 0
)

/* debug [enums.gen.go]: Processing enum IOPCISaveDeviceStateOptions (1 cases) */
// IOPCISaveDeviceStateOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCISaveDeviceStateOptions
type IOPCISaveDeviceStateOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/IOPCISaveDeviceStateOptions/kPCIConfigShadowPermanent
	kPCIConfigShadowPermanent IOPCISaveDeviceStateOptions = 0
)

/* debug [enums.gen.go]: Processing enum tIOPCIAccessOptions (1 cases) */
// tIOPCIAccessOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCIAccessOptions
type tIOPCIAccessOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCIAccessOptions/kIOPCIAccessLatencyTolerantHint
	kIOPCIAccessLatencyTolerantHint tIOPCIAccessOptions = 0
)

/* debug [enums.gen.go]: Processing enum tIOPCIDeviceResetOptions (2 cases) */
// tIOPCIDeviceResetOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCIDeviceResetOptions
type tIOPCIDeviceResetOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCIDeviceResetOptions/kIOPCIDeviceResetOptionNone
	kIOPCIDeviceResetOptionNone tIOPCIDeviceResetOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCIDeviceResetOptions/kIOPCIDeviceResetOptionTerminate
	kIOPCIDeviceResetOptionTerminate tIOPCIDeviceResetOptions = 0
)

/* debug [enums.gen.go]: Processing enum tIOPCIDeviceResetTypes (5 cases) */
// tIOPCIDeviceResetTypes enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCIDeviceResetTypes
type tIOPCIDeviceResetTypes uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCIDeviceResetTypes/kIOPCIDeviceResetTypeFunctionReset
	kIOPCIDeviceResetTypeFunctionReset tIOPCIDeviceResetTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCIDeviceResetTypes/kIOPCIDeviceResetTypeHotReset
	kIOPCIDeviceResetTypeHotReset tIOPCIDeviceResetTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCIDeviceResetTypes/kIOPCIDeviceResetTypeWarmReset
	kIOPCIDeviceResetTypeWarmReset tIOPCIDeviceResetTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCIDeviceResetTypes/kIOPCIDeviceResetTypeWarmResetDisable
	kIOPCIDeviceResetTypeWarmResetDisable tIOPCIDeviceResetTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCIDeviceResetTypes/kIOPCIDeviceResetTypeWarmResetEnable
	kIOPCIDeviceResetTypeWarmResetEnable tIOPCIDeviceResetTypes = 0
)

/* debug [enums.gen.go]: Processing enum tIOPCILinkControlASPMBits (4 cases) */
// tIOPCILinkControlASPMBits enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCILinkControlASPMBits
type tIOPCILinkControlASPMBits uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCILinkControlASPMBits/kIOPCILinkControlASPMBitsDisabled
	kIOPCILinkControlASPMBitsDisabled tIOPCILinkControlASPMBits = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCILinkControlASPMBits/kIOPCILinkControlASPMBitsL0s
	kIOPCILinkControlASPMBitsL0s tIOPCILinkControlASPMBits = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCILinkControlASPMBits/kIOPCILinkControlASPMBitsL0sL1
	kIOPCILinkControlASPMBitsL0sL1 tIOPCILinkControlASPMBits = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCILinkControlASPMBits/kIOPCILinkControlASPMBitsL1
	kIOPCILinkControlASPMBitsL1 tIOPCILinkControlASPMBits = 0
)

/* debug [enums.gen.go]: Processing enum tIOPCILinkSpeed (5 cases) */
// tIOPCILinkSpeed enum type
//
// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCILinkSpeed
type tIOPCILinkSpeed uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCILinkSpeed/kIOPCILinkSpeed_16_GTs
	kIOPCILinkSpeed_16_GTs tIOPCILinkSpeed = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCILinkSpeed/kIOPCILinkSpeed_2_5_GTs
	kIOPCILinkSpeed_2_5_GTs tIOPCILinkSpeed = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCILinkSpeed/kIOPCILinkSpeed_32_GTs
	kIOPCILinkSpeed_32_GTs tIOPCILinkSpeed = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCILinkSpeed/kIOPCILinkSpeed_5_GTs
	kIOPCILinkSpeed_5_GTs tIOPCILinkSpeed = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PCIDriverKit/tIOPCILinkSpeed/kIOPCILinkSpeed_8_GTs
	kIOPCILinkSpeed_8_GTs tIOPCILinkSpeed = 0
)
