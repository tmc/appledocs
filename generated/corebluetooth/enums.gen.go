// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

// Enum types and constants
// CBATTError - The possible errors returned by a GATT server (a remote peripheral) during Bluetooth low energy ATT transactions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code
type CBATTError uint

const (
// CBATTErrorAttributeNotFound - The attribute wasn’t found within the specified attribute handle range.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/attributeNotFound
CBATTErrorAttributeNotFound CBATTError = 0
// CBATTErrorAttributeNotLong - The ATT read blob request can’t read or write the attribute.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/attributeNotLong
CBATTErrorAttributeNotLong CBATTError = 0
// CBATTErrorInsufficientAuthentication - Reading or writing the attribute’s value failed for lack of authentication.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/insufficientAuthentication
CBATTErrorInsufficientAuthentication CBATTError = 0
// CBATTErrorInsufficientAuthorization - Reading or writing the attribute’s value failed for lack of authorization.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/insufficientAuthorization
CBATTErrorInsufficientAuthorization CBATTError = 0
// CBATTErrorInsufficientEncryption - Reading or writing the attribute’s value failed for lack of encryption.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/insufficientEncryption
CBATTErrorInsufficientEncryption CBATTError = 0
// CBATTErrorInsufficientEncryptionKeySize - The encryption key size used for encrypting this link is insufficient.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/insufficientEncryptionKeySize
CBATTErrorInsufficientEncryptionKeySize CBATTError = 0
// CBATTErrorInsufficientResources - Resources are insufficient to complete the ATT request.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/insufficientResources
CBATTErrorInsufficientResources CBATTError = 0
// CBATTErrorInvalidAttributeValueLength - The length of the attribute’s value is invalid for the intended operation.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/invalidAttributeValueLength
CBATTErrorInvalidAttributeValueLength CBATTError = 0
// CBATTErrorInvalidHandle - The attribute handle is invalid on this peripheral.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/invalidHandle
CBATTErrorInvalidHandle CBATTError = 0
// CBATTErrorInvalidOffset - The specified offset value was past the end of the attribute’s value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/invalidOffset
CBATTErrorInvalidOffset CBATTError = 0
// CBATTErrorInvalidPdu - The attribute Protocol Data Unit (PDU) is invalid.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/invalidPdu
CBATTErrorInvalidPdu CBATTError = 0
// CBATTErrorPrepareQueueFull - The prepare queue is full, as a result of there being too many write requests in the queue.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/prepareQueueFull
CBATTErrorPrepareQueueFull CBATTError = 0
// CBATTErrorReadNotPermitted - The permissions prohibit reading the attribute’s value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/readNotPermitted
CBATTErrorReadNotPermitted CBATTError = 0
// CBATTErrorRequestNotSupported - The attribute server doesn’t support the request received from the client.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/requestNotSupported
CBATTErrorRequestNotSupported CBATTError = 0
// CBATTErrorSuccess - The ATT command or request successfully completed.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/success
CBATTErrorSuccess CBATTError = 0
// CBATTErrorUnlikelyError - The ATT request encountered an unlikely error and wasn’t completed.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/unlikelyError
CBATTErrorUnlikelyError CBATTError = 0
// CBATTErrorUnsupportedGroupType - The attribute type isn’t a supported grouping attribute as defined by a higher-layer specification.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/unsupportedGroupType
CBATTErrorUnsupportedGroupType CBATTError = 0
// CBATTErrorWriteNotPermitted - The permissions prohibit writing the attribute’s value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBATTError-swift.struct/Code/writeNotPermitted
CBATTErrorWriteNotPermitted CBATTError = 0
)

// CBAttributePermissions - Values that represent the read, write, and encryption permissions for a characteristic’s value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBAttributePermissions
type CBAttributePermissions uint

const (
// CBAttributePermissionsReadEncryptionRequired - A permission that indicates only trusted devices can read the attribute’s value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBAttributePermissions/readEncryptionRequired
CBAttributePermissionsReadEncryptionRequired CBAttributePermissions = 0
// CBAttributePermissionsReadable - A permission that indicates a peripheral can read the attribute’s value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBAttributePermissions/readable
CBAttributePermissionsReadable CBAttributePermissions = 0
// CBAttributePermissionsWriteEncryptionRequired - A permission that indicates only trusted devices can write the attribute’s value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBAttributePermissions/writeEncryptionRequired
CBAttributePermissionsWriteEncryptionRequired CBAttributePermissions = 0
// CBAttributePermissionsWriteable - A permission that indicates a peripheral can write the attribute’s value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBAttributePermissions/writeable
CBAttributePermissionsWriteable CBAttributePermissions = 0
)

// CBCentralManagerFeature - An option set of device-specific features.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/Feature
type CBCentralManagerFeature uint

const (
// CBCentralManagerFeatureExtendedScanAndConnect - The hardware supports extended scans and enhanced connection creation.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/Feature/extendedScanAndConnect
CBCentralManagerFeatureExtendedScanAndConnect CBCentralManagerFeature = 0
)

// CBCentralManagerState - Values that represent the current state of a central manager object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerState
type CBCentralManagerState uint

const (
// CBCentralManagerStatePoweredOff - A state that indicates Bluetooth is currently powered off.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerState/poweredOff
CBCentralManagerStatePoweredOff CBCentralManagerState = 0
// CBCentralManagerStatePoweredOn - A state that indicates Bluetooth is currently powered on and available to use.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerState/poweredOn
CBCentralManagerStatePoweredOn CBCentralManagerState = 0
// CBCentralManagerStateResetting - A state that indicates the connection with the system service was momentarily lost.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerState/resetting
CBCentralManagerStateResetting CBCentralManagerState = 0
// CBCentralManagerStateUnauthorized - A state that indicates the application isn’t authorized to use the Bluetooth low energy role.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerState/unauthorized
CBCentralManagerStateUnauthorized CBCentralManagerState = 0
// CBCentralManagerStateUnknown - The manager’s state is unknown.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerState/unknown
CBCentralManagerStateUnknown CBCentralManagerState = 0
// CBCentralManagerStateUnsupported - A state that indicates this device doesn’t support the Bluetooth low energy central or client role.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManagerState/unsupported
CBCentralManagerStateUnsupported CBCentralManagerState = 0
)

// CBCharacteristicProperties - Values that represent the possible properties of a characteristic.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristicProperties
type CBCharacteristicProperties uint

// CBCharacteristicWriteType - Values representing the possible write types to a characteristic’s value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristicWriteType
type CBCharacteristicWriteType uint

const (
// CBCharacteristicWriteWithResponse - Write a characteristic value, with a response from the peripheral to indicate whether the write was successful.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristicWriteType/withResponse
CBCharacteristicWriteWithResponse CBCharacteristicWriteType = 0
// CBCharacteristicWriteWithoutResponse - Write a characteristic value, without any response from the peripheral to indicate whether the write was successful.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBCharacteristicWriteType/withoutResponse
CBCharacteristicWriteWithoutResponse CBCharacteristicWriteType = 0
)

// CBConnectionEvent - A change to the connection state of a peer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBConnectionEvent
type CBConnectionEvent uint

// CBError - The codes for errors that Core Bluetooth returns during Bluetooth transactions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code
type CBError uint

const (
// CBErrorUnknownDevice - The device is unknown.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-c.enum/CBErrorUnknownDevice
CBErrorUnknownDevice CBError = 0
// CBErrorAlreadyAdvertising - The peripheral is already advertising.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/alreadyAdvertising
CBErrorAlreadyAdvertising CBError = 0
// CBErrorConnectionFailed - The connection failed.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/connectionFailed
CBErrorConnectionFailed CBError = 0
// CBErrorConnectionLimitReached - The device already has the maximum number of connections.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/connectionLimitReached
CBErrorConnectionLimitReached CBError = 0
// CBErrorConnectionTimeout - The connection timed out.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/connectionTimeout
CBErrorConnectionTimeout CBError = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/encryptionTimedOut
CBErrorEncryptionTimedOut CBError = 0
// CBErrorInvalidHandle - The specified attribute handle is invalid.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/invalidHandle
CBErrorInvalidHandle CBError = 0
// CBErrorInvalidParameters - The specified parameters are invalid.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/invalidParameters
CBErrorInvalidParameters CBError = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/leGattExceededBackgroundNotificationLimit
CBErrorLeGattExceededBackgroundNotificationLimit CBError = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/leGattNearBackgroundNotificationLimit
CBErrorLeGattNearBackgroundNotificationLimit CBError = 0
// CBErrorNotConnected - The device isn’t currently connected.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/notConnected
CBErrorNotConnected CBError = 0
// CBErrorOperationCancelled - The error represents a canceled operation.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/operationCancelled
CBErrorOperationCancelled CBError = 0
// CBErrorOperationNotSupported - The operation isn’t supported.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/operationNotSupported
CBErrorOperationNotSupported CBError = 0
// CBErrorOutOfSpace - The device has run out of space to complete the intended operation.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/outOfSpace
CBErrorOutOfSpace CBError = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/peerRemovedPairingInformation
CBErrorPeerRemovedPairingInformation CBError = 0
// CBErrorPeripheralDisconnected - The peripheral disconnected.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/peripheralDisconnected
CBErrorPeripheralDisconnected CBError = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/tooManyLEPairedDevices
CBErrorTooManyLEPairedDevices CBError = 0
// CBErrorUnknown - An unknown error occurred.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/unknown
CBErrorUnknown CBError = 0
// CBErrorUnkownDevice - A misspelled version of the unknown device error code.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/unkownDevice
CBErrorUnkownDevice CBError = 0
// CBErrorUUIDNotAllowed - The specified UUID isn’t permitted.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBError-swift.struct/Code/uuidNotAllowed
CBErrorUUIDNotAllowed CBError = 0
)

// CBManagerAuthorization - The current authorization state of a Core Bluetooth manager.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManagerAuthorization
type CBManagerAuthorization uint

const (
// CBManagerAuthorizationAllowedAlways - A state that indicates the user has authorized Bluetooth at any time.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManagerAuthorization/allowedAlways
CBManagerAuthorizationAllowedAlways CBManagerAuthorization = 0
// CBManagerAuthorizationDenied - A state that indicates the user explicitly denied Bluetooth access for this app.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManagerAuthorization/denied
CBManagerAuthorizationDenied CBManagerAuthorization = 0
// CBManagerAuthorizationNotDetermined - A state that indicates the user has yet to authorize Bluetooth for this app.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManagerAuthorization/notDetermined
CBManagerAuthorizationNotDetermined CBManagerAuthorization = 0
)

// CBManagerState - The possible states of a Core Bluetooth manager.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManagerState
type CBManagerState uint

const (
// CBManagerStatePoweredOff - A state that indicates Bluetooth is currently powered off.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManagerState/poweredOff
CBManagerStatePoweredOff CBManagerState = 0
// CBManagerStatePoweredOn - A state that indicates Bluetooth is currently powered on and available to use.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManagerState/poweredOn
CBManagerStatePoweredOn CBManagerState = 0
// CBManagerStateResetting - A state that indicates the connection with the system service was momentarily lost.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManagerState/resetting
CBManagerStateResetting CBManagerState = 0
// CBManagerStateUnauthorized - A state that indicates the application isn’t authorized to use the Bluetooth low energy role.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManagerState/unauthorized
CBManagerStateUnauthorized CBManagerState = 0
// CBManagerStateUnknown - The manager’s state is unknown.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManagerState/unknown
CBManagerStateUnknown CBManagerState = 0
// CBManagerStateUnsupported - A state that indicates this device doesn’t support the Bluetooth low energy central or client role.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManagerState/unsupported
CBManagerStateUnsupported CBManagerState = 0
)

// CBPeripheralManagerAuthorizationStatus - Values representing the current authorization state of the peripheral manager.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerAuthorizationStatus
type CBPeripheralManagerAuthorizationStatus uint

// CBPeripheralManagerConnectionLatency - Values representing the connection latency of the peripheral manager.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerConnectionLatency
type CBPeripheralManagerConnectionLatency uint

// CBPeripheralManagerState - Values that represent the current state of the peripheral manager.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerState
type CBPeripheralManagerState uint

const (
// CBPeripheralManagerStatePoweredOff - A manager state that indicates Bluetooth is currently powered off.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerState/poweredOff
CBPeripheralManagerStatePoweredOff CBPeripheralManagerState = 0
// CBPeripheralManagerStatePoweredOn - A manager state that indicates Bluetooth is currently powered on and is available to use.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerState/poweredOn
CBPeripheralManagerStatePoweredOn CBPeripheralManagerState = 0
// CBPeripheralManagerStateResetting - A manager state that indicates the connection with the system service was momentarily lost.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerState/resetting
CBPeripheralManagerStateResetting CBPeripheralManagerState = 0
// CBPeripheralManagerStateUnauthorized - A manager state that indicates the app isn’t authorized to use the Bluetooth low energy peripheral/server role.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerState/unauthorized
CBPeripheralManagerStateUnauthorized CBPeripheralManagerState = 0
// CBPeripheralManagerStateUnknown - A manager state that indicates the current state of the peripheral manager is unknown.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerState/unknown
CBPeripheralManagerStateUnknown CBPeripheralManagerState = 0
// CBPeripheralManagerStateUnsupported - A manager state that indicates the platform doesn’t support the Bluetooth low energy peripheral/server role.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralManagerState/unsupported
CBPeripheralManagerStateUnsupported CBPeripheralManagerState = 0
)

// CBPeripheralState - Values representing the connection state of a peripheral.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheralState
type CBPeripheralState uint


