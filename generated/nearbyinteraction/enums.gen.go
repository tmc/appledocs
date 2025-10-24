// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

// Enum types and constants
// NIAlgorithmConvergenceStatus - Expose algorithm state to make it possible for apps to coach users.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIAlgorithmConvergenceStatus-2fbmj
type NIAlgorithmConvergenceStatus int

const (
	// NIAlgorithmConvergenceStatusNotConverged - A status that indicates the framework’s Camera Assistance feature requires action from the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIAlgorithmConvergenceStatus-2fbmj/NIAlgorithmConvergenceStatusNotConverged
	NIAlgorithmConvergenceStatusNotConverged NIAlgorithmConvergenceStatus = 0
	// NIAlgorithmConvergenceStatusUnknown - An indication that the framework is unsure of the Camera Assistance status.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIAlgorithmConvergenceStatus-2fbmj/NIAlgorithmConvergenceStatusUnknown
	NIAlgorithmConvergenceStatusUnknown NIAlgorithmConvergenceStatus = 0
)

// NIDLTDOACoordinatesType - The possible coordinate types for Downlink Time-Difference-of-Arrival measurement updates.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOACoordinatesType
type NIDLTDOACoordinatesType uint

// NIDLTDOAMeasurementType - The possible phases of downlink positioning signals.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAMeasurementType
type NIDLTDOAMeasurementType uint

// NIErrorCode - Codes that identify errors in Nearby Interaction.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIError/Code
type NIErrorCode uint

const (
	// NIErrorCodeAccessoryPeerDeviceUnavailable - An error that indicates the peer Bluetooth accessory isn’t connected or paired.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIError/Code/accessoryPeerDeviceUnavailable
	NIErrorCodeAccessoryPeerDeviceUnavailable NIErrorCode = 0
	// NIErrorCodeInvalidConfiguration - An error code that indicates that the nearby-interaction configuration isn’t valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIError/Code/invalidConfiguration
	NIErrorCodeInvalidConfiguration NIErrorCode = 0
	// NIErrorCodeUserDidNotAllow - An error code that indicates that the user declined the request to share their relative position with nearby devices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIError/Code/userDidNotAllow
	NIErrorCodeUserDidNotAllow NIErrorCode = 0
)

// NINearbyObjectRemovalReason - The reason a session removed a nearby object.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/RemovalReason
type NINearbyObjectRemovalReason uint


