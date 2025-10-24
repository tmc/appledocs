// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

/* debug [enums.gen.go]: Generating 12 enums for CoreMotion */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum CMAttitudeReferenceFrame (4 cases) */
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

/* debug [enums.gen.go]: Processing enum CMDeviceMotionSensorLocation (3 cases) */
// CMDeviceMotionSensorLocation - Defines the device’s sensor locations.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/SensorLocation-swift.enum
type CMDeviceMotionSensorLocation uint

const (
	// CMDeviceMotionSensorLocationDefault - The default sensor location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/SensorLocation-swift.enum/default
	CMDeviceMotionSensorLocationDefault CMDeviceMotionSensorLocation = 0
	// CMDeviceMotionSensorLocationHeadphoneLeft - The sensor is in the left headphone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/SensorLocation-swift.enum/headphoneLeft
	CMDeviceMotionSensorLocationHeadphoneLeft CMDeviceMotionSensorLocation = 0
	// CMDeviceMotionSensorLocationHeadphoneRight - The sensor is in the right headphone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMDeviceMotion/SensorLocation-swift.enum/headphoneRight
	CMDeviceMotionSensorLocationHeadphoneRight CMDeviceMotionSensorLocation = 0
)

/* debug [enums.gen.go]: Processing enum CMFallDetectionEventUserResolution (4 cases) */
// CMFallDetectionEventUserResolution - User resolutions for fall detection events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent/UserResolution
type CMFallDetectionEventUserResolution uint

const (
	// CMFallDetectionEventUserResolutionConfirmed - The user confirmed the event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent/UserResolution/confirmed
	CMFallDetectionEventUserResolutionConfirmed CMFallDetectionEventUserResolution = 0
	// CMFallDetectionEventUserResolutionDismissed - The user dismissed the fall event alert, but didn’t explicitly confirm or reject the event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent/UserResolution/dismissed
	CMFallDetectionEventUserResolutionDismissed CMFallDetectionEventUserResolution = 0
	// CMFallDetectionEventUserResolutionRejected - The user rejected the fall event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent/UserResolution/rejected
	CMFallDetectionEventUserResolutionRejected CMFallDetectionEventUserResolution = 0
	// CMFallDetectionEventUserResolutionUnresponsive - The user didn’t respond to the fall event and the system hasn’t detected recovery motions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent/UserResolution/unresponsive
	CMFallDetectionEventUserResolutionUnresponsive CMFallDetectionEventUserResolution = 0
)

/* debug [enums.gen.go]: Processing enum CMHeadphoneActivityStatus (2 cases) */
// CMHeadphoneActivityStatus - Headphone connection status updates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneActivityManager/Status
type CMHeadphoneActivityStatus uint

const (
	// CMHeadphoneActivityStatusConnected - A compatible set of headphones is connected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneActivityManager/Status/connected
	CMHeadphoneActivityStatusConnected CMHeadphoneActivityStatus = 0
	// CMHeadphoneActivityStatusDisconnected - The headphones disconnected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHeadphoneActivityManager/Status/disconnected
	CMHeadphoneActivityStatusDisconnected CMHeadphoneActivityStatus = 0
)

/* debug [enums.gen.go]: Processing enum CMWaterSubmersionState (3 cases) */
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
	// CMWaterSubmersionStateUnknown - The submersion state is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionEvent/State-swift.enum/unknown
	CMWaterSubmersionStateUnknown CMWaterSubmersionState = 0
)

/* debug [enums.gen.go]: Processing enum CMWaterSubmersionDepthState (7 cases) */
// CMWaterSubmersionDepthState - A state based on the device’s depth under water.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState
type CMWaterSubmersionDepthState uint

const (
	// CMWaterSubmersionDepthStateApproachingMaxDepth - The device is approaching the maximum safe diving depth.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState/approachingMaxDepth
	CMWaterSubmersionDepthStateApproachingMaxDepth CMWaterSubmersionDepthState = 0
	// CMWaterSubmersionDepthStateNotSubmerged - The device is not submerged in water.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState/notSubmerged
	CMWaterSubmersionDepthStateNotSubmerged CMWaterSubmersionDepthState = 0
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
	// CMWaterSubmersionDepthStateSubmergedShallow - The device is submerged, but less than 1 meter under water.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState/submergedShallow
	CMWaterSubmersionDepthStateSubmergedShallow CMWaterSubmersionDepthState = 0
	// CMWaterSubmersionDepthStateUnknown - The device’s depth state is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement/DepthState/unknown
	CMWaterSubmersionDepthStateUnknown CMWaterSubmersionDepthState = 0
)

/* debug [enums.gen.go]: Processing enum CMAuthorizationStatus (4 cases) */
// CMAuthorizationStatus - The authorization status for motion-related features.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAuthorizationStatus
type CMAuthorizationStatus uint

const (
	// CMAuthorizationStatusAuthorized - Access was granted by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAuthorizationStatus/authorized
	CMAuthorizationStatusAuthorized CMAuthorizationStatus = 0
	// CMAuthorizationStatusDenied - Access was denied by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAuthorizationStatus/denied
	CMAuthorizationStatusDenied CMAuthorizationStatus = 0
	// CMAuthorizationStatusNotDetermined - The status has not yet been determined.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAuthorizationStatus/notDetermined
	CMAuthorizationStatusNotDetermined CMAuthorizationStatus = 0
	// CMAuthorizationStatusRestricted - Access is denied due to system-wide restrictions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAuthorizationStatus/restricted
	CMAuthorizationStatusRestricted CMAuthorizationStatus = 0
)

/* debug [enums.gen.go]: Processing enum CMHighFrequencyHeartRateDataConfidence (4 cases) */
// CMHighFrequencyHeartRateDataConfidence - The level of confidence in the accuracy of the heart rate data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateDataConfidence
type CMHighFrequencyHeartRateDataConfidence uint

const (
	// CMHighFrequencyHeartRateDataConfidenceHigh - A high level of confidence in the heart rate data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateDataConfidence/high
	CMHighFrequencyHeartRateDataConfidenceHigh CMHighFrequencyHeartRateDataConfidence = 0
	// CMHighFrequencyHeartRateDataConfidenceHighest - The highest level of confidence in the heart rate data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateDataConfidence/highest
	CMHighFrequencyHeartRateDataConfidenceHighest CMHighFrequencyHeartRateDataConfidence = 0
	// CMHighFrequencyHeartRateDataConfidenceLow - A low level of confidence in the heart rate data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateDataConfidence/low
	CMHighFrequencyHeartRateDataConfidenceLow CMHighFrequencyHeartRateDataConfidence = 0
	// CMHighFrequencyHeartRateDataConfidenceMedium - A medium level of confidence in the heart rate data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMHighFrequencyHeartRateDataConfidence/medium
	CMHighFrequencyHeartRateDataConfidenceMedium CMHighFrequencyHeartRateDataConfidence = 0
)

/* debug [enums.gen.go]: Processing enum CMMagneticFieldCalibrationAccuracy (4 cases) */
// CMMagneticFieldCalibrationAccuracy - Indicates the calibration accuracy of a magnetic field estimate
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMagneticFieldCalibrationAccuracy
type CMMagneticFieldCalibrationAccuracy uint

const (
	// CMMagneticFieldCalibrationAccuracyHigh - The accuracy of the magnetic field calibration is high.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMagneticFieldCalibrationAccuracy/high
	CMMagneticFieldCalibrationAccuracyHigh CMMagneticFieldCalibrationAccuracy = 0
	// CMMagneticFieldCalibrationAccuracyLow - The accuracy of the magnetic field calibration is low.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMagneticFieldCalibrationAccuracy/low
	CMMagneticFieldCalibrationAccuracyLow CMMagneticFieldCalibrationAccuracy = 0
	// CMMagneticFieldCalibrationAccuracyMedium - The accuracy of the magnetic field calibration is medium.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMagneticFieldCalibrationAccuracy/medium
	CMMagneticFieldCalibrationAccuracyMedium CMMagneticFieldCalibrationAccuracy = 0
	// CMMagneticFieldCalibrationAccuracyUncalibrated - The magnetic field estimate is not calibrated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMagneticFieldCalibrationAccuracy/uncalibrated
	CMMagneticFieldCalibrationAccuracyUncalibrated CMMagneticFieldCalibrationAccuracy = 0
)

/* debug [enums.gen.go]: Processing enum CMMotionActivityConfidence (3 cases) */
// CMMotionActivityConfidence - The confidence that the motion data is accurate.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityConfidence
type CMMotionActivityConfidence uint

const (
	// CMMotionActivityConfidenceHigh - Confidence is high.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityConfidence/high
	CMMotionActivityConfidenceHigh CMMotionActivityConfidence = 0
	// CMMotionActivityConfidenceLow - Confidence is low.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityConfidence/low
	CMMotionActivityConfidenceLow CMMotionActivityConfidence = 0
	// CMMotionActivityConfidenceMedium - Confidence is good.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMotionActivityConfidence/medium
	CMMotionActivityConfidenceMedium CMMotionActivityConfidence = 0
)

/* debug [enums.gen.go]: Processing enum CMOdometerOriginDevice (3 cases) */
// CMOdometerOriginDevice - The device that the odometer sample originates from.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerOriginDevice
type CMOdometerOriginDevice uint

const (
	// CMOdometerOriginDeviceLocal - The origin of the odometer sample comes from the same device that requests the sample.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerOriginDevice/local
	CMOdometerOriginDeviceLocal CMOdometerOriginDevice = 0
	// CMOdometerOriginDeviceRemote - The origin of the odometer sample comes from a device that’s paired with the local device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerOriginDevice/remote
	CMOdometerOriginDeviceRemote CMOdometerOriginDevice = 0
	// CMOdometerOriginDeviceUnknown - The origin of the odometer sample is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMOdometerOriginDevice/unknown
	CMOdometerOriginDeviceUnknown CMOdometerOriginDevice = 0
)

/* debug [enums.gen.go]: Processing enum CMPedometerEventType (2 cases) */
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


