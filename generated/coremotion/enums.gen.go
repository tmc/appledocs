// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

// Enum types and constants
// CMAttitudeReferenceFrame - Constants that indicate the frame of reference for attitude-related motion data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitudeReferenceFrame
type AttitudeReferenceFrame uint

const (
	// AttitudeReferenceFrameXArbitraryCorrectedZVertical - A reference frame where the Z axis is vertical and has improved rotation accuracy, and the X axis points in an arbitrary direction in the horizontal plane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitudeReferenceFrame/xArbitraryCorrectedZVertical
	AttitudeReferenceFrameXArbitraryCorrectedZVertical AttitudeReferenceFrame = 0
	// AttitudeReferenceFrameXArbitraryZVertical - A reference frame where the Z axis is vertical and the X axis points in an arbitrary direction in the horizontal plane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitudeReferenceFrame/xArbitraryZVertical
	AttitudeReferenceFrameXArbitraryZVertical AttitudeReferenceFrame = 0
	// AttitudeReferenceFrameXMagneticNorthZVertical - A reference frame where the Z axis is vertical and the X axis points to the magnetic north pole.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitudeReferenceFrame/xMagneticNorthZVertical
	AttitudeReferenceFrameXMagneticNorthZVertical AttitudeReferenceFrame = 0
	// AttitudeReferenceFrameXTrueNorthZVertical - A reference frame where the Z axis is vertical and the X axis points to the geographic north pole.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitudeReferenceFrame/xTrueNorthZVertical
	AttitudeReferenceFrameXTrueNorthZVertical AttitudeReferenceFrame = 0
)

// CMAuthorizationStatus - The authorization status for motion-related features.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAuthorizationStatus
type AuthorizationStatus uint

const (
	// AuthorizationStatusAuthorized - Access was granted by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAuthorizationStatus/authorized
	AuthorizationStatusAuthorized AuthorizationStatus = 0
	// AuthorizationStatusDenied - Access was denied by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAuthorizationStatus/denied
	AuthorizationStatusDenied AuthorizationStatus = 0
	// AuthorizationStatusNotDetermined - The status has not yet been determined.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAuthorizationStatus/notDetermined
	AuthorizationStatusNotDetermined AuthorizationStatus = 0
	// AuthorizationStatusRestricted - Access is denied due to system-wide restrictions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAuthorizationStatus/restricted
	AuthorizationStatusRestricted AuthorizationStatus = 0
)

// CMDeviceMotionSensorLocation - Defines the device’s sensor locations.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/SensorLocation-swift.enum
type DeviceMotionSensorLocation uint

// CMFallDetectionEventUserResolution - User resolutions for fall detection events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent/UserResolution
type FallDetectionEventUserResolution uint

const (
	// FallDetectionEventUserResolutionConfirmed - The user confirmed the event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent/UserResolution/confirmed
	FallDetectionEventUserResolutionConfirmed FallDetectionEventUserResolution = 0
	// FallDetectionEventUserResolutionDismissed - The user dismissed the fall event alert, but didn’t explicitly confirm or reject the event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent/UserResolution/dismissed
	FallDetectionEventUserResolutionDismissed FallDetectionEventUserResolution = 0
	// FallDetectionEventUserResolutionRejected - The user rejected the fall event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent/UserResolution/rejected
	FallDetectionEventUserResolutionRejected FallDetectionEventUserResolution = 0
	// FallDetectionEventUserResolutionUnresponsive - The user didn’t respond to the fall event and the system hasn’t detected recovery motions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent/UserResolution/unresponsive
	FallDetectionEventUserResolutionUnresponsive FallDetectionEventUserResolution = 0
)

// CMHeadphoneActivityStatus - Headphone connection status updates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneActivityManager/Status
type HeadphoneActivityStatus uint

// CMHighFrequencyHeartRateDataConfidence - The level of confidence in the accuracy of the heart rate data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateDataConfidence
type HighFrequencyHeartRateDataConfidence uint

const (
	// HighFrequencyHeartRateDataConfidenceHigh - A high level of confidence in the heart rate data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateDataConfidence/high
	HighFrequencyHeartRateDataConfidenceHigh HighFrequencyHeartRateDataConfidence = 0
	// HighFrequencyHeartRateDataConfidenceHighest - The highest level of confidence in the heart rate data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateDataConfidence/highest
	HighFrequencyHeartRateDataConfidenceHighest HighFrequencyHeartRateDataConfidence = 0
	// HighFrequencyHeartRateDataConfidenceLow - A low level of confidence in the heart rate data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateDataConfidence/low
	HighFrequencyHeartRateDataConfidenceLow HighFrequencyHeartRateDataConfidence = 0
	// HighFrequencyHeartRateDataConfidenceMedium - A medium level of confidence in the heart rate data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateDataConfidence/medium
	HighFrequencyHeartRateDataConfidenceMedium HighFrequencyHeartRateDataConfidence = 0
)

// CMMagneticFieldCalibrationAccuracy - Indicates the calibration accuracy of a magnetic field estimate
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMagneticFieldCalibrationAccuracy
type MagneticFieldCalibrationAccuracy uint

const (
	// MagneticFieldCalibrationAccuracyUncalibrated - The magnetic field estimate is not calibrated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMagneticFieldCalibrationAccuracy/uncalibrated
	MagneticFieldCalibrationAccuracyUncalibrated MagneticFieldCalibrationAccuracy = 0
)

// CMMotionActivityConfidence - The confidence that the motion data is accurate.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityConfidence
type MotionActivityConfidence uint

// CMOdometerOriginDevice - The device that the odometer sample originates from.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerOriginDevice
type OdometerOriginDevice uint

// CMPedometerEventType - Constants indicating the change that occurred to the user’s pedestrian activity.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerEventType
type PedometerEventType uint

const (
	// PedometerEventTypePause - The user’s pedestrian activity stopped.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerEventType/pause
	PedometerEventTypePause PedometerEventType = 0
	// PedometerEventTypeResume - The user’s pedestrian activity resumed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerEventType/resume
	PedometerEventTypeResume PedometerEventType = 0
)

// CMWaterSubmersionState - The device’s submersion state.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionEvent/State-swift.enum
type WaterSubmersionState uint

const (
	// WaterSubmersionStateNotSubmerged - The device isn’t submerged in water.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionEvent/State-swift.enum/notSubmerged
	WaterSubmersionStateNotSubmerged WaterSubmersionState = 0
	// WaterSubmersionStateSubmerged - The device is submerged in water.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionEvent/State-swift.enum/submerged
	WaterSubmersionStateSubmerged WaterSubmersionState = 0
	// WaterSubmersionStateUnknown - The submersion state is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionEvent/State-swift.enum/unknown
	WaterSubmersionStateUnknown WaterSubmersionState = 0
)

// CMWaterSubmersionDepthState - A state based on the device’s depth under water.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState
type WaterSubmersionDepthState uint

const (
	// WaterSubmersionDepthStateApproachingMaxDepth - The device is approaching the maximum safe diving depth.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState/approachingMaxDepth
	WaterSubmersionDepthStateApproachingMaxDepth WaterSubmersionDepthState = 0
	// WaterSubmersionDepthStateNotSubmerged - The device is not submerged in water.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState/notSubmerged
	WaterSubmersionDepthStateNotSubmerged WaterSubmersionDepthState = 0
	// WaterSubmersionDepthStatePastMaxDepth - The device has exceeded the maximum safe diving depth.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState/pastMaxDepth
	WaterSubmersionDepthStatePastMaxDepth WaterSubmersionDepthState = 0
	// WaterSubmersionDepthStateSensorDepthError - An error with the depth sensor occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState/sensorDepthError
	WaterSubmersionDepthStateSensorDepthError WaterSubmersionDepthState = 0
	// WaterSubmersionDepthStateSubmergedDeep - The device is submerged at least 1 meter under water.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState/submergedDeep
	WaterSubmersionDepthStateSubmergedDeep WaterSubmersionDepthState = 0
	// WaterSubmersionDepthStateSubmergedShallow - The device is submerged, but less than 1 meter under water.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState/submergedShallow
	WaterSubmersionDepthStateSubmergedShallow WaterSubmersionDepthState = 0
	// WaterSubmersionDepthStateUnknown - The device’s depth state is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState/unknown
	WaterSubmersionDepthStateUnknown WaterSubmersionDepthState = 0
)


