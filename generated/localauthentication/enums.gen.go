// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

// Enum types and constants
// LAAccessControlOperation - Operations to be evaluated for access control.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAAccessControlOperation
type LAAccessControlOperation uint

// LABiometryType - The set of available biometric authentication types.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LABiometryType
type LABiometryType uint

// LACompanionType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LACompanionType
type LACompanionType uint

const (
	// LACompanionTypeMac - Paired Mac
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LACompanionType/mac
	LACompanionTypeMac LACompanionType = 0
	// LACompanionTypeVision - Paired Vision Pro
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LACompanionType/vision
	LACompanionTypeVision LACompanionType = 0
	// LACompanionTypeWatch - Paired Apple Watch
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LACompanionType/watch
	LACompanionTypeWatch LACompanionType = 0
)

// LACredentialType - The types of credentials to be used for authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LACredentialType
type LACredentialType uint

// LAError - Errors issued by the LocalAuthentication framework.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code
type LAError uint

const (
	// LAErrorBiometryLockout - Biometry is locked because there were too many failed attempts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-c.enum/LAErrorBiometryLockout
	LAErrorBiometryLockout LAError = 0
	// LAErrorBiometryNotAvailable - Biometry is not available on the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-c.enum/LAErrorBiometryNotAvailable
	LAErrorBiometryNotAvailable LAError = 0
	// LAErrorBiometryNotEnrolled - The user has no enrolled biometric identities.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-c.enum/LAErrorBiometryNotEnrolled
	LAErrorBiometryNotEnrolled LAError = 0
	// LAErrorAppCancel - The app canceled authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/appCancel
	LAErrorAppCancel LAError = 0
	// LAErrorAuthenticationFailed - The user failed to provide valid credentials.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/authenticationFailed
	LAErrorAuthenticationFailed LAError = 0
	// LAErrorBiometryDisconnected - The device supports biometry only using a removable accessory, but the paired accessory isn’t connected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/biometryDisconnected
	LAErrorBiometryDisconnected LAError = 0
	// LAErrorBiometryNotPaired - The device supports biometry only using a removable accessory, but no accessory is paired.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/biometryNotPaired
	LAErrorBiometryNotPaired LAError = 0
	// LAErrorCompanionNotAvailable - Authentication could not start because there was no paired companion device nearby.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/companionNotAvailable-swift.enum.case
	LAErrorCompanionNotAvailable LAError = 0
	// LAErrorInvalidContext - The context was previously invalidated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/invalidContext
	LAErrorInvalidContext LAError = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/invalidDimensions
	LAErrorInvalidDimensions LAError = 0
	// LAErrorNotInteractive - Displaying the required authentication user interface is forbidden.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/notInteractive
	LAErrorNotInteractive LAError = 0
	// LAErrorPasscodeNotSet - A passcode isn’t set on the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/passcodeNotSet
	LAErrorPasscodeNotSet LAError = 0
	// LAErrorSystemCancel - The system canceled authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/systemCancel
	LAErrorSystemCancel LAError = 0
	// LAErrorTouchIDLockout - Touch ID is locked because there were too many failed attempts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/touchIDLockout
	LAErrorTouchIDLockout LAError = 0
	// LAErrorTouchIDNotAvailable - Touch ID is not available on the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/touchIDNotAvailable
	LAErrorTouchIDNotAvailable LAError = 0
	// LAErrorTouchIDNotEnrolled - The user has no enrolled Touch ID fingers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/touchIDNotEnrolled
	LAErrorTouchIDNotEnrolled LAError = 0
	// LAErrorUserCancel - The user tapped the cancel button in the authentication dialog.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/userCancel
	LAErrorUserCancel LAError = 0
	// LAErrorUserFallback - The user tapped the fallback button in the authentication dialog, but no fallback is available for the authentication policy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/userFallback
	LAErrorUserFallback LAError = 0
	// LAErrorWatchNotAvailable - An attempt to authenticate with Apple Watch failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAError-swift.struct/Code/watchNotAvailable
	LAErrorWatchNotAvailable LAError = 0
)

// LAPolicy - The set of available local authentication policies.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy
type LAPolicy uint

const (
	// LAPolicyDeviceOwnerAuthentication - User authentication with biometry, Apple Watch, or the device passcode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy/deviceOwnerAuthentication
	LAPolicyDeviceOwnerAuthentication LAPolicy = 0
	// LAPolicyDeviceOwnerAuthenticationWithBiometrics - User authentication with biometry.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy/deviceOwnerAuthenticationWithBiometrics
	LAPolicyDeviceOwnerAuthenticationWithBiometrics LAPolicy = 0
	// LAPolicyDeviceOwnerAuthenticationWithBiometricsOrWatch - User authentication with either biometry or Apple Watch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy/deviceOwnerAuthenticationWithBiometricsOrWatch
	LAPolicyDeviceOwnerAuthenticationWithBiometricsOrWatch LAPolicy = 0
	// LAPolicyDeviceOwnerAuthenticationWithWatch - User authentication with Apple Watch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy/deviceOwnerAuthenticationWithWatch
	LAPolicyDeviceOwnerAuthenticationWithWatch LAPolicy = 0
	// LAPolicyDeviceOwnerAuthenticationWithWristDetection - User authentication with wrist detection on watchOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPolicy/deviceOwnerAuthenticationWithWristDetection
	LAPolicyDeviceOwnerAuthenticationWithWristDetection LAPolicy = 0
)

// LARightState - The possible states for a right during authorization.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/State-swift.enum
type LARightState uint

const (
	// LARightStateAuthorized - The authorization completed successfully.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/State-swift.enum/authorized
	LARightStateAuthorized LARightState = 0
	// LARightStateAuthorizing - The authorization is in progress but not completed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/State-swift.enum/authorizing
	LARightStateAuthorizing LARightState = 0
	// LARightStateNotAuthorized - The authorization failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/State-swift.enum/notAuthorized
	LARightStateNotAuthorized LARightState = 0
	// LARightStateUnknown - The authorization is in an unknown state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LARight/State-swift.enum/unknown
	LARightStateUnknown LARightState = 0
)


