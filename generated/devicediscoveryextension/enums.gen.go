// Code generated from Apple documentation for DeviceDiscoveryExtension. DO NOT EDIT.

package devicediscoveryextension

// Enum types and constants
// DDDeviceCategory - An option that determines the icon for the device in the picker UI.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/Category-swift.enum
type DDDeviceCategory uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/Category-swift.enum/accessorySetup
	DDDeviceCategoryAccessorySetup DDDeviceCategory = 0
	// DDDeviceCategoryDesktopComputer - An icon that depicts a desktop computer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/Category-swift.enum/desktopComputer
	DDDeviceCategoryDesktopComputer DDDeviceCategory = 0
	// DDDeviceCategoryHiFiSpeaker - An icon that depicts a high-fidelity speaker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/Category-swift.enum/hifiSpeaker
	DDDeviceCategoryHiFiSpeaker DDDeviceCategory = 0
	// DDDeviceCategoryHiFiSpeakerMultiple - An icon that depicts multiple high-fidelity speakers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/Category-swift.enum/hifiSpeakerMultiple
	DDDeviceCategoryHiFiSpeakerMultiple DDDeviceCategory = 0
	// DDDeviceCategoryLaptopComputer - An icon that depicts a laptop computer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/Category-swift.enum/laptopComputer
	DDDeviceCategoryLaptopComputer DDDeviceCategory = 0
	// DDDeviceCategoryTV - An icon that depicts a television.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/Category-swift.enum/tv
	DDDeviceCategoryTV DDDeviceCategory = 0
	// DDDeviceCategoryTVWithMediaBox - An icon that depicts a TV with a set-top box.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/Category-swift.enum/tvWithMediaBox
	DDDeviceCategoryTVWithMediaBox DDDeviceCategory = 0
)

// DDDeviceMediaPlaybackState - States that indicate the status of a device’s media playback.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/MediaPlaybackState-swift.enum
type DDDeviceMediaPlaybackState uint

const (
	// DDDeviceMediaPlaybackStateNoContent - A state that indicates when the device plays no content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/MediaPlaybackState-swift.enum/noContent
	DDDeviceMediaPlaybackStateNoContent DDDeviceMediaPlaybackState = 0
	// DDDeviceMediaPlaybackStatePaused - A state that indicates when content playback for the device pauses.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/MediaPlaybackState-swift.enum/paused
	DDDeviceMediaPlaybackStatePaused DDDeviceMediaPlaybackState = 0
	// DDDeviceMediaPlaybackStatePlaying - A state that indicates when the device plays media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/MediaPlaybackState-swift.enum/playing
	DDDeviceMediaPlaybackStatePlaying DDDeviceMediaPlaybackState = 0
)

// DDDeviceProtocol - An identifier for the manner in which an app interacts with a device.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/Protocol-swift.enum
type DDDeviceProtocol uint

const (
	// DDDeviceProtocolDIAL - A protocol for client devices that stream media to a TV or set-top box.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/Protocol-swift.enum/dial
	DDDeviceProtocolDIAL DDDeviceProtocol = 0
	// DDDeviceProtocolInvalid - A default value for a device protocol.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/Protocol-swift.enum/invalid
	DDDeviceProtocolInvalid DDDeviceProtocol = 0
)

// DDDeviceWiFiAwareServiceRole enum type
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/WiFiAwareServiceRole-swift.enum
type DDDeviceWiFiAwareServiceRole uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/WiFiAwareServiceRole-swift.enum/publisher
	DDDeviceWiFiAwareServiceRolePublisher DDDeviceWiFiAwareServiceRole = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDevice/WiFiAwareServiceRole-swift.enum/subscriber
	DDDeviceWiFiAwareServiceRoleSubscriber DDDeviceWiFiAwareServiceRole = 0
)

// DDEventType - Identifiers for the types of events that occur in the device discovery life cycle.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceEvent/EventType-swift.enum
type DDEventType uint

const (
	// DDEventTypeDeviceChanged - A status that indicates when the device of interest changes configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceEvent/EventType-swift.enum/deviceChanged
	DDEventTypeDeviceChanged DDEventType = 0
	// DDEventTypeDeviceFound - A status that indicates when the extension finds the device of interest.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceEvent/EventType-swift.enum/deviceFound
	DDEventTypeDeviceFound DDEventType = 0
	// DDEventTypeDeviceLost - A status that indicates when the extension loses a connection to the device of interest.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceEvent/EventType-swift.enum/deviceLost
	DDEventTypeDeviceLost DDEventType = 0
	// DDEventTypeUnknown - A value for uninitialized event types.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceEvent/EventType-swift.enum/unknown
	DDEventTypeUnknown DDEventType = 0
)

// DDDeviceState - A state that represents the level of user interaction with the device.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceState
type DDDeviceState uint

const (
	// DDDeviceStateActivated - A state that indicates when the user authorizes the device and the app connects to the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceState/activated
	DDDeviceStateActivated DDDeviceState = 0
	// DDDeviceStateActivating - A state that indicates when the user selects the device in the picker UI.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceState/activating
	DDDeviceStateActivating DDDeviceState = 0
	// DDDeviceStateAuthorized - A state that indicates when the user authorizes the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceState/authorized
	DDDeviceStateAuthorized DDDeviceState = 0
	// DDDeviceStateInvalid - A state that indicates the device is invalid or that the user disapproves of the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceState/invalid
	DDDeviceStateInvalid DDDeviceState = 0
	// DDDeviceStateInvalidating - A state that indicates that the device is soon to be invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceState/invalidating
	DDDeviceStateInvalidating DDDeviceState = 0
)

// DDDeviceSupports enum type
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceSupports
type DDDeviceSupports uint

const (
	// DDDeviceSupportsBluetoothHID - Device supports bring up of classic transport profiles when low energy transport for peripheral is connected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceSupports/bluetoothHID
	DDDeviceSupportsBluetoothHID DDDeviceSupports = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceSupports/bluetoothPairingLE
	DDDeviceSupportsBluetoothPairingLE DDDeviceSupports = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDDeviceSupports/bluetoothTransportBridging
	DDDeviceSupportsBluetoothTransportBridging DDDeviceSupports = 0
)

// DDErrorCode - Codes that identify errors that can occur during the framework’s use.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDError/Code
type DDErrorCode uint

const (
	// DDErrorCodeBadParameter - An error that indicates the framework doesn’t support a parameter that the extension provides.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDError/Code/badParameter
	DDErrorCodeBadParameter DDErrorCode = 0
	// DDErrorCodeInternal - An error that indicates a problem of internal origin.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDError/Code/internal
	DDErrorCodeInternal DDErrorCode = 0
	// DDErrorCodeMissingEntitlement - An error that indicates that the app extension lacks a required entitlement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDError/Code/missingEntitlement
	DDErrorCodeMissingEntitlement DDErrorCode = 0
	// DDErrorCodeNext - An error the framework reserves for future use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDError/Code/next
	DDErrorCodeNext DDErrorCode = 0
	// DDErrorCodePermission - An error that indicates the app extension lacks necessary permissions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDError/Code/permission
	DDErrorCodePermission DDErrorCode = 0
	// DDErrorCodeSuccess - An error that indicates an operation succeeds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDError/Code/success
	DDErrorCodeSuccess DDErrorCode = 0
	// DDErrorCodeTimeout - An error that indicates a timeout occurs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDError/Code/timeout
	DDErrorCodeTimeout DDErrorCode = 0
	// DDErrorCodeUnknown - An error that indicates an uncategorized problem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDError/Code/unknown
	DDErrorCodeUnknown DDErrorCode = 0
	// DDErrorCodeUnsupported - An error that indicates an unsupported configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceDiscoveryExtension/DDError/Code/unsupported
	DDErrorCodeUnsupported DDErrorCode = 0
)


