// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

/* debug [enums.gen.go]: Generating 6 enums for NearbyInteraction */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum NIAlgorithmConvergenceStatus (3 cases) */
// NIAlgorithmConvergenceStatus - Expose algorithm state to make it possible for apps to coach users.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIAlgorithmConvergenceStatus-2fbmj
type NIAlgorithmConvergenceStatus int

const (
	// NIAlgorithmConvergenceStatusConverged - A status that indicates the framework’s Camera Assistance feature is operational.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIAlgorithmConvergenceStatus-2fbmj/NIAlgorithmConvergenceStatusConverged
	NIAlgorithmConvergenceStatusConverged NIAlgorithmConvergenceStatus = 0
	// NIAlgorithmConvergenceStatusNotConverged - A status that indicates the framework’s Camera Assistance feature requires action from the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIAlgorithmConvergenceStatus-2fbmj/NIAlgorithmConvergenceStatusNotConverged
	NIAlgorithmConvergenceStatusNotConverged NIAlgorithmConvergenceStatus = 0
	// NIAlgorithmConvergenceStatusUnknown - An indication that the framework is unsure of the Camera Assistance status.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIAlgorithmConvergenceStatus-2fbmj/NIAlgorithmConvergenceStatusUnknown
	NIAlgorithmConvergenceStatusUnknown NIAlgorithmConvergenceStatus = 0
)

/* debug [enums.gen.go]: Processing enum NIErrorCode (10 cases) */
// NIErrorCode - Codes that identify errors in Nearby Interaction.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIError/Code
type NIErrorCode uint

const (
	// NIErrorCodeAccessoryPeerDeviceUnavailable - An error that indicates the peer Bluetooth accessory isn’t connected or paired.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIError/Code/accessoryPeerDeviceUnavailable
	NIErrorCodeAccessoryPeerDeviceUnavailable NIErrorCode = 0
	// NIErrorCodeActiveExtendedDistanceSessionsLimitExceeded - An error that indicates the device exceeds the available number of active extended distance sessions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIError/Code/activeExtendedDistanceSessionsLimitExceeded
	NIErrorCodeActiveExtendedDistanceSessionsLimitExceeded NIErrorCode = 0
	// NIErrorCodeActiveSessionsLimitExceeded - An error code that indicates that the app reached the maximum number of sessions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIError/Code/activeSessionsLimitExceeded
	NIErrorCodeActiveSessionsLimitExceeded NIErrorCode = 0
	// NIErrorCodeIncompatiblePeerDevice - An error that indicates the peer device isn’t compatible with this Nearby Interaction session instance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIError/Code/incompatiblePeerDevice
	NIErrorCodeIncompatiblePeerDevice NIErrorCode = 0
	// NIErrorCodeInvalidARConfiguration - An error that indicates the framework can’t begin Camera Assistance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIError/Code/invalidARConfiguration
	NIErrorCodeInvalidARConfiguration NIErrorCode = 0
	// NIErrorCodeInvalidConfiguration - An error code that indicates that the nearby-interaction configuration isn’t valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIError/Code/invalidConfiguration
	NIErrorCodeInvalidConfiguration NIErrorCode = 0
	// NIErrorCodeResourceUsageTimeout - An error code that indicates that the framework timed out the session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIError/Code/resourceUsageTimeout
	NIErrorCodeResourceUsageTimeout NIErrorCode = 0
	// NIErrorCodeSessionFailed - An error code that indicates that the session failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIError/Code/sessionFailed
	NIErrorCodeSessionFailed NIErrorCode = 0
	// NIErrorCodeUnsupportedPlatform - An error code that indicates that the framework doesn’t support the device platform.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIError/Code/unsupportedPlatform
	NIErrorCodeUnsupportedPlatform NIErrorCode = 0
	// NIErrorCodeUserDidNotAllow - An error code that indicates that the user declined the request to share their relative position with nearby devices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIError/Code/userDidNotAllow
	NIErrorCodeUserDidNotAllow NIErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum NINearbyObjectRemovalReason (2 cases) */
// NINearbyObjectRemovalReason - The reason a session removed a nearby object.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/RemovalReason
type NINearbyObjectRemovalReason uint

const (
	// NINearbyObjectRemovalReasonPeerEnded - The peer ended the session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/RemovalReason/peerEnded
	NINearbyObjectRemovalReasonPeerEnded NINearbyObjectRemovalReason = 0
	// NINearbyObjectRemovalReasonTimeout - NI timed out the session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/RemovalReason/timeout
	NINearbyObjectRemovalReasonTimeout NINearbyObjectRemovalReason = 0
)

/* debug [enums.gen.go]: Processing enum NINearbyObjectVerticalDirectionEstimate (5 cases) */
// NINearbyObjectVerticalDirectionEstimate - Estimations of a nearby object’s vertical position in relation to the user’s device.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/VerticalDirectionEstimate-swift.enum
type NINearbyObjectVerticalDirectionEstimate uint

const (
	// NINearbyObjectVerticalDirectionEstimateAbove - An indication that the nearby object resides at a higher vertical location than the user’s device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/VerticalDirectionEstimate-swift.enum/above
	NINearbyObjectVerticalDirectionEstimateAbove NINearbyObjectVerticalDirectionEstimate = 0
	// NINearbyObjectVerticalDirectionEstimateAboveOrBelow - An indication that the nearby object doesn’t reside at the same vertical location as the user’s device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/VerticalDirectionEstimate-swift.enum/aboveOrBelow
	NINearbyObjectVerticalDirectionEstimateAboveOrBelow NINearbyObjectVerticalDirectionEstimate = 0
	// NINearbyObjectVerticalDirectionEstimateBelow - An indication that the nearby object resides at a lower vertical location than the user’s device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/VerticalDirectionEstimate-swift.enum/below
	NINearbyObjectVerticalDirectionEstimateBelow NINearbyObjectVerticalDirectionEstimate = 0
	// NINearbyObjectVerticalDirectionEstimateSame - An indication that the nearby object resides at an equivalent vertical location as the user’s device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/VerticalDirectionEstimate-swift.enum/same
	NINearbyObjectVerticalDirectionEstimateSame NINearbyObjectVerticalDirectionEstimate = 0
	// NINearbyObjectVerticalDirectionEstimateUnknown - An indication that the nearby object resides at an unknown vertical location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NINearbyObject/VerticalDirectionEstimate-swift.enum/unknown
	NINearbyObjectVerticalDirectionEstimateUnknown NINearbyObjectVerticalDirectionEstimate = 0
)

/* debug [enums.gen.go]: Processing enum NIDLTDOACoordinatesType (2 cases) */
// NIDLTDOACoordinatesType - The possible coordinate types for Downlink Time-Difference-of-Arrival measurement updates.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOACoordinatesType
type NIDLTDOACoordinatesType uint

const (
	// NIDLTDOACoordinatesTypeGeodetic - A coordinate type that specifies a latitude, longitude, and altitude triplet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOACoordinatesType/geodetic
	NIDLTDOACoordinatesTypeGeodetic NIDLTDOACoordinatesType = 0
	// NIDLTDOACoordinatesTypeRelative - A coordinate type that specifies a 3D Cartesian triplet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOACoordinatesType/relative
	NIDLTDOACoordinatesTypeRelative NIDLTDOACoordinatesType = 0
)

/* debug [enums.gen.go]: Processing enum NIDLTDOAMeasurementType (3 cases) */
// NIDLTDOAMeasurementType - The possible phases of downlink positioning signals.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAMeasurementType
type NIDLTDOAMeasurementType uint

const (
	// NIDLTDOAMeasurementTypeFinal - A type that indicates the measurement derives from an initial anchor’s last message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAMeasurementType/final
	NIDLTDOAMeasurementTypeFinal NIDLTDOAMeasurementType = 0
	// NIDLTDOAMeasurementTypePoll - A type that indicates the measurement derives from an initiating anchor’s first message.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAMeasurementType/poll
	NIDLTDOAMeasurementTypePoll NIDLTDOAMeasurementType = 0
	// NIDLTDOAMeasurementTypeResponse - A type that indicates the measurement derives from responder anchors’ messages.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDLTDOAMeasurementType/response
	NIDLTDOAMeasurementTypeResponse NIDLTDOAMeasurementType = 0
)


