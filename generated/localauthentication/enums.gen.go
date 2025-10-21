// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

// Enum types and constants
// LAAccessControlOperation - Operations to be evaluated for access control.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAccessControlOperation
type AccessControlOperation uint

const (
// AccessControlOperationCreateItem - Specifies that access control is used for item creation.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAccessControlOperation/createItem
AccessControlOperationCreateItem AccessControlOperation = 0
// AccessControlOperationCreateKey - Specifies that access control is used for key creation.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAccessControlOperation/createKey
AccessControlOperationCreateKey AccessControlOperation = 0
// AccessControlOperationUseItem - Specifies that access control is used for accessing an existing item.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAccessControlOperation/useItem
AccessControlOperationUseItem AccessControlOperation = 0
// AccessControlOperationUseKeyDecrypt - Specifies that access control is used for data decryption using existing key.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAccessControlOperation/useKeyDecrypt
AccessControlOperationUseKeyDecrypt AccessControlOperation = 0
// AccessControlOperationUseKeyKeyExchange - Specifies that access control is used for key exchange.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAccessControlOperation/useKeyKeyExchange
AccessControlOperationUseKeyKeyExchange AccessControlOperation = 0
// AccessControlOperationUseKeySign - Specifies that access control is used for accessing an existing key.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAccessControlOperation/useKeySign
AccessControlOperationUseKeySign AccessControlOperation = 0
)

// LABiometryType - The set of available biometric authentication types.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LABiometryType
type BiometryType uint

// LACompanionType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LACompanionType
type CompanionType uint

const (
// CompanionTypeMac - Paired Mac
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LACompanionType/mac
CompanionTypeMac CompanionType = 0
// CompanionTypeVision - Paired Vision Pro
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LACompanionType/vision
CompanionTypeVision CompanionType = 0
// CompanionTypeWatch - Paired Apple Watch
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LACompanionType/watch
CompanionTypeWatch CompanionType = 0
)

// LACredentialType - The types of credentials to be used for authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LACredentialType
type CredentialType uint

// LAError - Errors issued by the LocalAuthentication framework.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code
type Error uint

const (
// ErrorBiometryLockout - Biometry is locked because there were too many failed attempts.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-c.enum/LAErrorBiometryLockout
ErrorBiometryLockout Error = 0
// ErrorBiometryNotAvailable - Biometry is not available on the device.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-c.enum/LAErrorBiometryNotAvailable
ErrorBiometryNotAvailable Error = 0
// ErrorBiometryNotEnrolled - The user has no enrolled biometric identities.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-c.enum/LAErrorBiometryNotEnrolled
ErrorBiometryNotEnrolled Error = 0
// ErrorAppCancel - The app canceled authentication.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/appCancel
ErrorAppCancel Error = 0
// ErrorAuthenticationFailed - The user failed to provide valid credentials.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/authenticationFailed
ErrorAuthenticationFailed Error = 0
// ErrorBiometryDisconnected - The device supports biometry only using a removable accessory, but the paired accessory isn’t connected.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/biometryDisconnected
ErrorBiometryDisconnected Error = 0
// ErrorBiometryNotPaired - The device supports biometry only using a removable accessory, but no accessory is paired.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/biometryNotPaired
ErrorBiometryNotPaired Error = 0
// ErrorCompanionNotAvailable - Authentication could not start because there was no paired companion device nearby.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/companionNotAvailable-swift.enum.case
ErrorCompanionNotAvailable Error = 0
// ErrorInvalidContext - The context was previously invalidated.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/invalidContext
ErrorInvalidContext Error = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/invalidDimensions
ErrorInvalidDimensions Error = 0
// ErrorNotInteractive - Displaying the required authentication user interface is forbidden.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/notInteractive
ErrorNotInteractive Error = 0
// ErrorPasscodeNotSet - A passcode isn’t set on the device.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/passcodeNotSet
ErrorPasscodeNotSet Error = 0
// ErrorSystemCancel - The system canceled authentication.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/systemCancel
ErrorSystemCancel Error = 0
// ErrorTouchIDLockout - Touch ID is locked because there were too many failed attempts.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/touchIDLockout
ErrorTouchIDLockout Error = 0
// ErrorTouchIDNotAvailable - Touch ID is not available on the device.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/touchIDNotAvailable
ErrorTouchIDNotAvailable Error = 0
// ErrorTouchIDNotEnrolled - The user has no enrolled Touch ID fingers.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/touchIDNotEnrolled
ErrorTouchIDNotEnrolled Error = 0
// ErrorUserCancel - The user tapped the cancel button in the authentication dialog.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/userCancel
ErrorUserCancel Error = 0
// ErrorUserFallback - The user tapped the fallback button in the authentication dialog, but no fallback is available for the authentication policy.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/userFallback
ErrorUserFallback Error = 0
// ErrorWatchNotAvailable - An attempt to authenticate with Apple Watch failed.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/watchNotAvailable
ErrorWatchNotAvailable Error = 0
)

// LAPolicy - The set of available local authentication policies.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy
type Policy uint

const (
// PolicyDeviceOwnerAuthentication - User authentication with biometry, Apple Watch, or the device passcode.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy/deviceOwnerAuthentication
PolicyDeviceOwnerAuthentication Policy = 0
// PolicyDeviceOwnerAuthenticationWithBiometrics - User authentication with biometry.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy/deviceOwnerAuthenticationWithBiometrics
PolicyDeviceOwnerAuthenticationWithBiometrics Policy = 0
// PolicyDeviceOwnerAuthenticationWithBiometricsOrCompanion - Device owner will be authenticated by biometry or a companion device e.g. Watch, Mac, etc.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy/deviceOwnerAuthenticationWithBiometricsOrCompanion
PolicyDeviceOwnerAuthenticationWithBiometricsOrCompanion Policy = 0
// PolicyDeviceOwnerAuthenticationWithBiometricsOrWatch - User authentication with either biometry or Apple Watch.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy/deviceOwnerAuthenticationWithBiometricsOrWatch
PolicyDeviceOwnerAuthenticationWithBiometricsOrWatch Policy = 0
// PolicyDeviceOwnerAuthenticationWithCompanion - Device owner will be authenticated by a companion device e.g. Watch, Mac, etc.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy/deviceOwnerAuthenticationWithCompanion
PolicyDeviceOwnerAuthenticationWithCompanion Policy = 0
// PolicyDeviceOwnerAuthenticationWithWatch - User authentication with Apple Watch.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy/deviceOwnerAuthenticationWithWatch
PolicyDeviceOwnerAuthenticationWithWatch Policy = 0
// PolicyDeviceOwnerAuthenticationWithWristDetection - User authentication with wrist detection on watchOS.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy/deviceOwnerAuthenticationWithWristDetection
PolicyDeviceOwnerAuthenticationWithWristDetection Policy = 0
)

// LARightState - The possible states for a right during authorization.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/State-swift.enum
type RightState uint

const (
// RightStateAuthorized - The authorization completed successfully.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/State-swift.enum/authorized
RightStateAuthorized RightState = 0
// RightStateAuthorizing - The authorization is in progress but not completed.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/State-swift.enum/authorizing
RightStateAuthorizing RightState = 0
// RightStateNotAuthorized - The authorization failed.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/State-swift.enum/notAuthorized
RightStateNotAuthorized RightState = 0
// RightStateUnknown - The authorization is in an unknown state.
//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/State-swift.enum/unknown
RightStateUnknown RightState = 0
)


