// Code generated from Apple documentation for DeviceCheck. DO NOT EDIT.

package devicecheck

// Enum types and constants
// DCError - DeviceCheck error codes.
//
// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCError-swift.struct/Code
type DCError uint

const (
// DCErrorFeatureUnsupported - DeviceCheck is unavailable on this device.
//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCError-swift.struct/Code/featureUnsupported
DCErrorFeatureUnsupported DCError = 0
// DCErrorInvalidInput - An error code that indicates when your app provides data that isn’t   formatted correctly.
//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCError-swift.struct/Code/invalidInput
DCErrorInvalidInput DCError = 0
// DCErrorInvalidKey - An error caused by a failed attempt to use the App Attest key.
//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCError-swift.struct/Code/invalidKey
DCErrorInvalidKey DCError = 0
// DCErrorServerUnavailable - An error that indicates a failed attempt to contact the App Attest service   during an attestation.
//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCError-swift.struct/Code/serverUnavailable
DCErrorServerUnavailable DCError = 0
// DCErrorUnknownSystemFailure - A failure has occurred, such as the failure to generate a token.
//
	// [Full Topic]: https://developer.apple.com/documentation/DeviceCheck/DCError-swift.struct/Code/unknownSystemFailure
DCErrorUnknownSystemFailure DCError = 0
)


