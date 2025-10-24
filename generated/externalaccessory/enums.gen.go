// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

/* debug [enums.gen.go]: Generating 4 enums for ExternalAccessory */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum EABluetoothAccessoryPickerErrorCode (4 cases) */
// EABluetoothAccessoryPickerErrorCode - The error codes that may be passed in an error object for the Bluetooth picker completion block.
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EABluetoothAccessoryPickerError/Code
type EABluetoothAccessoryPickerErrorCode uint

const (
	// EABluetoothAccessoryPickerAlreadyConnected - The specified accessory was already connected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EABluetoothAccessoryPickerError/Code/alreadyConnected
	EABluetoothAccessoryPickerAlreadyConnected EABluetoothAccessoryPickerErrorCode = 0
	// EABluetoothAccessoryPickerResultCancelled - The user canceled the picker alert.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EABluetoothAccessoryPickerError/Code/resultCancelled
	EABluetoothAccessoryPickerResultCancelled EABluetoothAccessoryPickerErrorCode = 0
	// EABluetoothAccessoryPickerResultFailed - Selecting an accessory failed for an unknown reason.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EABluetoothAccessoryPickerError/Code/resultFailed
	EABluetoothAccessoryPickerResultFailed EABluetoothAccessoryPickerErrorCode = 0
	// EABluetoothAccessoryPickerResultNotFound - The specified accessory could not be found, perhaps because it was turned off prior to connection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EABluetoothAccessoryPickerError/Code/resultNotFound
	EABluetoothAccessoryPickerResultNotFound EABluetoothAccessoryPickerErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum EAWiFiUnconfiguredAccessoryBrowserState (4 cases) */
// EAWiFiUnconfiguredAccessoryBrowserState - The possible states of an accessory browser.
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowserState
type EAWiFiUnconfiguredAccessoryBrowserState uint

const (
	// EAWiFiUnconfiguredAccessoryBrowserStateConfiguring - The browser is actively configuring an accessory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowserState/configuring
	EAWiFiUnconfiguredAccessoryBrowserStateConfiguring EAWiFiUnconfiguredAccessoryBrowserState = 0
	// EAWiFiUnconfiguredAccessoryBrowserStateSearching - The browser is actively searching for unconfigured accessory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowserState/searching
	EAWiFiUnconfiguredAccessoryBrowserStateSearching EAWiFiUnconfiguredAccessoryBrowserState = 0
	// EAWiFiUnconfiguredAccessoryBrowserStateStopped - The browser is not actively searching for unconfigured accessories.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowserState/stopped
	EAWiFiUnconfiguredAccessoryBrowserStateStopped EAWiFiUnconfiguredAccessoryBrowserState = 0
	// EAWiFiUnconfiguredAccessoryBrowserStateWiFiUnavailable - Wi-Fi is unavailable, typically because the user has placed the device in Airplane Mode or explicitly turned off Wi-Fi.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowserState/wiFiUnavailable
	EAWiFiUnconfiguredAccessoryBrowserStateWiFiUnavailable EAWiFiUnconfiguredAccessoryBrowserState = 0
)

/* debug [enums.gen.go]: Processing enum EAWiFiUnconfiguredAccessoryConfigurationStatus (3 cases) */
// EAWiFiUnconfiguredAccessoryConfigurationStatus - Values that represent the state of the configuration process for an 
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryConfigurationStatus
type EAWiFiUnconfiguredAccessoryConfigurationStatus uint

const (
	// EAWiFiUnconfiguredAccessoryConfigurationStatusFailed - The configuration failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryConfigurationStatus/failed
	EAWiFiUnconfiguredAccessoryConfigurationStatusFailed EAWiFiUnconfiguredAccessoryConfigurationStatus = 0
	// EAWiFiUnconfiguredAccessoryConfigurationStatusSuccess - The configuration of the accessory succeeded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryConfigurationStatus/success
	EAWiFiUnconfiguredAccessoryConfigurationStatusSuccess EAWiFiUnconfiguredAccessoryConfigurationStatus = 0
	// EAWiFiUnconfiguredAccessoryConfigurationStatusUserCancelledConfiguration - The user cancelled the configuration process.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryConfigurationStatus/userCancelledConfiguration
	EAWiFiUnconfiguredAccessoryConfigurationStatusUserCancelledConfiguration EAWiFiUnconfiguredAccessoryConfigurationStatus = 0
)

/* debug [enums.gen.go]: Processing enum EAWiFiUnconfiguredAccessoryProperties (3 cases) */
// EAWiFiUnconfiguredAccessoryProperties - Options that can be combined using the C bitwise 
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryProperties
type EAWiFiUnconfiguredAccessoryProperties uint

const (
	// EAWiFiUnconfiguredAccessoryPropertySupportsAirPlay - The accessory indicates that it supports AirPlay.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryProperties/propertySupportsAirPlay
	EAWiFiUnconfiguredAccessoryPropertySupportsAirPlay EAWiFiUnconfiguredAccessoryProperties = 0
	// EAWiFiUnconfiguredAccessoryPropertySupportsAirPrint - The accessory indicates that it supports AirPrint.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryProperties/propertySupportsAirPrint
	EAWiFiUnconfiguredAccessoryPropertySupportsAirPrint EAWiFiUnconfiguredAccessoryProperties = 0
	// EAWiFiUnconfiguredAccessoryPropertySupportsHomeKit - The accessory indicates that it supports HomeKit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryProperties/propertySupportsHomeKit
	EAWiFiUnconfiguredAccessoryPropertySupportsHomeKit EAWiFiUnconfiguredAccessoryProperties = 0
)


