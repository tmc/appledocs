// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

// Enum types and constants
// CMAttitudeReferenceFrame - Constants that indicate the frame of reference for attitude-related motion data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitudeReferenceFrame
type CMAttitudeReferenceFrame uint

const (
	// CMAttitudeReferenceFrameXArbitraryCorrectedZVertical - A reference frame where the Z axis is vertical and has improved rotation accuracy, and the X axis points in an arbitrary direction in the horizontal plane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitudeReferenceFrame/xArbitraryCorrectedZVertical
	CMAttitudeReferenceFrameXArbitraryCorrectedZVertical CMAttitudeReferenceFrame = 0
	// CMAttitudeReferenceFrameXArbitraryZVertical - A reference frame where the Z axis is vertical and the X axis points in an arbitrary direction in the horizontal plane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitudeReferenceFrame/xArbitraryZVertical
	CMAttitudeReferenceFrameXArbitraryZVertical CMAttitudeReferenceFrame = 0
	// CMAttitudeReferenceFrameXMagneticNorthZVertical - A reference frame where the Z axis is vertical and the X axis points to the magnetic north pole.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitudeReferenceFrame/xMagneticNorthZVertical
	CMAttitudeReferenceFrameXMagneticNorthZVertical CMAttitudeReferenceFrame = 0
	// CMAttitudeReferenceFrameXTrueNorthZVertical - A reference frame where the Z axis is vertical and the X axis points to the geographic north pole.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAttitudeReferenceFrame/xTrueNorthZVertical
	CMAttitudeReferenceFrameXTrueNorthZVertical CMAttitudeReferenceFrame = 0
)

// CMAuthorizationStatus - The authorization status for motion-related features.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAuthorizationStatus
type CMAuthorizationStatus uint

// CMDeviceMotionSensorLocation - Defines the device’s sensor locations.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/SensorLocation-swift.enum
type CMDeviceMotionSensorLocation uint

// CMFallDetectionEventUserResolution - User resolutions for fall detection events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent/UserResolution
type CMFallDetectionEventUserResolution uint

// CMHeadphoneActivityStatus - Headphone connection status updates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneActivityManager/Status
type CMHeadphoneActivityStatus uint

// CMHighFrequencyHeartRateDataConfidence - The level of confidence in the accuracy of the heart rate data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateDataConfidence
type CMHighFrequencyHeartRateDataConfidence uint

// CMMagneticFieldCalibrationAccuracy - Indicates the calibration accuracy of a magnetic field estimate
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMagneticFieldCalibrationAccuracy
type CMMagneticFieldCalibrationAccuracy uint

// CMMotionActivityConfidence - The confidence that the motion data is accurate.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityConfidence
type CMMotionActivityConfidence uint

// CMOdometerOriginDevice - The device that the odometer sample originates from.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerOriginDevice
type CMOdometerOriginDevice uint

// CMPedometerEventType - Constants indicating the change that occurred to the user’s pedestrian activity.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerEventType
type CMPedometerEventType uint

const (
	// CMPedometerEventTypePause - The user’s pedestrian activity stopped.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerEventType/pause
	CMPedometerEventTypePause CMPedometerEventType = 0
	// CMPedometerEventTypeResume - The user’s pedestrian activity resumed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerEventType/resume
	CMPedometerEventTypeResume CMPedometerEventType = 0
)

// CMWaterSubmersionState - The device’s submersion state.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionEvent/State-swift.enum
type CMWaterSubmersionState uint

const (
	// CMWaterSubmersionStateNotSubmerged - The device isn’t submerged in water.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionEvent/State-swift.enum/notSubmerged
	CMWaterSubmersionStateNotSubmerged CMWaterSubmersionState = 0
	// CMWaterSubmersionStateSubmerged - The device is submerged in water.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionEvent/State-swift.enum/submerged
	CMWaterSubmersionStateSubmerged CMWaterSubmersionState = 0
)

// CMWaterSubmersionDepthState - A state based on the device’s depth under water.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState
type CMWaterSubmersionDepthState uint

const (
	// CMWaterSubmersionDepthStateApproachingMaxDepth - The device is approaching the maximum safe diving depth.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState/approachingMaxDepth
	CMWaterSubmersionDepthStateApproachingMaxDepth CMWaterSubmersionDepthState = 0
	// CMWaterSubmersionDepthStatePastMaxDepth - The device has exceeded the maximum safe diving depth.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState/pastMaxDepth
	CMWaterSubmersionDepthStatePastMaxDepth CMWaterSubmersionDepthState = 0
	// CMWaterSubmersionDepthStateSensorDepthError - An error with the depth sensor occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState/sensorDepthError
	CMWaterSubmersionDepthStateSensorDepthError CMWaterSubmersionDepthState = 0
	// CMWaterSubmersionDepthStateSubmergedDeep - The device is submerged at least 1 meter under water.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState/submergedDeep
	CMWaterSubmersionDepthStateSubmergedDeep CMWaterSubmersionDepthState = 0
)


