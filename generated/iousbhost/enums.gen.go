// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

/* debug [enums.gen.go]: Generating 5 enums for IOUSBHost */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum IOUSBHostAbortOption (2 cases) */
// IOUSBHostAbortOption - Options for aborting pending input/output requests.
//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostAbortOption
type IOUSBHostAbortOption uint

const (
	// IOUSBHostAbortOptionAsynchronous - The option to abort input/output requests asynchronously.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostAbortOption/asynchronous
	IOUSBHostAbortOptionAsynchronous IOUSBHostAbortOption = 0
	// IOUSBHostAbortOptionSynchronous - The option to abort input/output requests synchronously.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostAbortOption/synchronous
	IOUSBHostAbortOptionSynchronous IOUSBHostAbortOption = 0
)

/* debug [enums.gen.go]: Processing enum IOUSBHostIsochronousTransactionOptions (2 cases) */
// IOUSBHostIsochronousTransactionOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIsochronousTransactionOptions
type IOUSBHostIsochronousTransactionOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIsochronousTransactionOptions/IOUSBHostIsochronousTransactionOptionsNone
	IOUSBHostIsochronousTransactionOptionsNone IOUSBHostIsochronousTransactionOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIsochronousTransactionOptions/wrap
	IOUSBHostIsochronousTransactionOptionsWrap IOUSBHostIsochronousTransactionOptions = 0
)

/* debug [enums.gen.go]: Processing enum IOUSBHostIsochronousTransferOptions (1 cases) */
// IOUSBHostIsochronousTransferOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIsochronousTransferOptions
type IOUSBHostIsochronousTransferOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIsochronousTransferOptions/IOUSBHostIsochronousTransferOptionsNone
	IOUSBHostIsochronousTransferOptionsNone IOUSBHostIsochronousTransferOptions = 0
)

/* debug [enums.gen.go]: Processing enum IOUSBHostObjectDestroyOptions (2 cases) */
// IOUSBHostObjectDestroyOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObjectDestroyOptions
type IOUSBHostObjectDestroyOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObjectDestroyOptions/deviceSurrender
	IOUSBHostObjectDestroyOptionsDeviceSurrender IOUSBHostObjectDestroyOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObjectDestroyOptions/IOUSBHostObjectDestroyOptionsNone
	IOUSBHostObjectDestroyOptionsNone IOUSBHostObjectDestroyOptions = 0
)

/* debug [enums.gen.go]: Processing enum IOUSBHostObjectInitOptions (3 cases) */
// IOUSBHostObjectInitOptions - Options for initializing the host object.
//
// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObjectInitOptions
type IOUSBHostObjectInitOptions uint

const (
	// IOUSBHostObjectInitOptionsDeviceCapture - The option to capture the device and terminate existing drivers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObjectInitOptions/deviceCapture
	IOUSBHostObjectInitOptionsDeviceCapture IOUSBHostObjectInitOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObjectInitOptions/deviceSeize
	IOUSBHostObjectInitOptionsDeviceSeize IOUSBHostObjectInitOptions = 0
	// IOUSBHostObjectInitOptionsNone - The default argument for initializing the host object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObjectInitOptions/IOUSBHostObjectInitOptionsNone
	IOUSBHostObjectInitOptionsNone IOUSBHostObjectInitOptions = 0
)


