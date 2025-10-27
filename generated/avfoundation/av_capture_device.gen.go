// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptureDevice] class.
var (
	CaptureDeviceClass     _CaptureDeviceClass
	CaptureDeviceClassOnce sync.Once
)

func getCaptureDeviceClass() _CaptureDeviceClass {
	CaptureDeviceClassOnce.Do(func() {
		CaptureDeviceClass = _CaptureDeviceClass{objc.GetClass("AVCaptureDevice")}
	})
	return CaptureDeviceClass
}

type _CaptureDeviceClass struct {
	class objc.Class
}





// An interface definition for the [CaptureDevice] class.
type ICaptureDevice interface {
	objectivec.IObject
	

	// properties:
	ActiveColorSpace() CaptureColorSpace
	SetActiveColorSpace(value CaptureColorSpace)
	ActiveFormat() IAVCaptureDeviceFormat
	SetActiveFormat(value IAVCaptureDeviceFormat)
	ActiveInputSource() IAVCaptureDeviceInputSource
	SetActiveInputSource(value IAVCaptureDeviceInputSource)
	ActivePrimaryConstituentDevice() IAVCaptureDevice
	ActivePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions
	ActivePrimaryConstituentDeviceSwitchingBehavior() CapturePrimaryConstituentDeviceSwitchingBehavior
	ActiveVideoMaxFrameDuration() objectivec.IObject
	SetActiveVideoMaxFrameDuration(value objectivec.IObject)
	ActiveVideoMinFrameDuration() objectivec.IObject
	SetActiveVideoMinFrameDuration(value objectivec.IObject)
	AvailableReactionTypes() unsafe.Pointer
	CameraLensSmudgeDetectionInterval() objectivec.IObject
	CameraLensSmudgeDetectionStatus() CaptureCameraLensSmudgeDetectionStatus
	CanPerformReactionEffects() bool
	CenterStageRectOfInterestSupported() bool
	CenterStageRectOfInterest() corefoundation.CGRect
	SetCenterStageRectOfInterest(value corefoundation.CGRect)
	CinematicVideoCaptureSceneMonitoringStatuses() unsafe.Pointer
	CompanionDeskViewCamera() IAVCaptureDevice
	DeviceType() CaptureDeviceType
	DisplayVideoZoomFactorMultiplier() float64
	ExposureMode() CaptureExposureMode
	SetExposureMode(value CaptureExposureMode)
	ExposurePointOfInterest() corefoundation.CGPoint
	SetExposurePointOfInterest(value corefoundation.CGPoint)
	ExposureRectOfInterest() corefoundation.CGRect
	SetExposureRectOfInterest(value corefoundation.CGRect)
	FallbackPrimaryConstituentDevices() []CaptureDevice
	SetFallbackPrimaryConstituentDevices(value []CaptureDevice)
	FlashMode() CaptureFlashMode
	SetFlashMode(value CaptureFlashMode)
	FocusMode() CaptureFocusMode
	SetFocusMode(value CaptureFocusMode)
	FocusPointOfInterest() corefoundation.CGPoint
	SetFocusPointOfInterest(value corefoundation.CGPoint)
	FocusRectOfInterest() corefoundation.CGRect
	SetFocusRectOfInterest(value corefoundation.CGRect)
	Formats() []CaptureDeviceFormat
	HasFlash() bool
	HasTorch() bool
	InputSources() []CaptureDeviceInputSource
	AdjustingExposure() bool
	AdjustingFocus() bool
	AdjustingWhiteBalance() bool
	AutoVideoFrameRateEnabled() bool
	SetAutoVideoFrameRateEnabled(value bool)
	BackgroundReplacementActive() bool
	CameraLensSmudgeDetectionEnabled() bool
	CenterStageActive() bool
	Connected() bool
	ContinuityCamera() bool
	ExposurePointOfInterestSupported() bool
	ExposureRectOfInterestSupported() bool
	FlashAvailable() bool
	FocusPointOfInterestSupported() bool
	FocusRectOfInterestSupported() bool
	FollowingExternalSyncDevice() bool
	InUseByAnotherApplication() bool
	PortraitEffectActive() bool
	StudioLightActive() bool
	Suspended() bool
	TorchActive() bool
	TorchAvailable() bool
	VideoFrameDurationLocked() bool
	LinkedDevices() []CaptureDevice
	LocalizedName() foundation.foundation.INSString
	Manufacturer() foundation.foundation.INSString
	MinExposureRectOfInterestSize() corefoundation.CGSize
	MinFocusRectOfInterestSize() corefoundation.CGSize
	MinimumFocusDistance() int
	MinSupportedExternalSyncFrameDuration() objectivec.IObject
	MinSupportedLockedVideoFrameDuration() objectivec.IObject
	ModelID() foundation.foundation.INSString
	Position() CaptureDevicePosition
	PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions
	PrimaryConstituentDeviceSwitchingBehavior() CapturePrimaryConstituentDeviceSwitchingBehavior
	ReactionEffectsInProgress() []CaptureReactionEffectState
	SpatialCaptureDiscomfortReasons() unsafe.Pointer
	SupportedFallbackPrimaryConstituentDevices() []CaptureDevice
	TorchLevel() float32
	TorchMode() CaptureTorchMode
	SetTorchMode(value CaptureTorchMode)
	TransportControlsPlaybackMode() CaptureDeviceTransportControlsPlaybackMode
	TransportControlsSpeed() CaptureDeviceTransportControlsSpeed
	TransportControlsSupported() bool
	TransportType() int32 /* not a class type */
	UniqueID() foundation.foundation.INSString
	WhiteBalanceMode() CaptureWhiteBalanceMode
	SetWhiteBalanceMode(value CaptureWhiteBalanceMode)
	ActivePrimaryConstituent() IAVCaptureDevice
	SetActivePrimaryConstituent(value IAVCaptureDevice)
	IsAutoVideoFrameRateEnabled() bool
	SetIsAutoVideoFrameRateEnabled(value bool)
	IsCameraLensSmudgeDetectionEnabled() bool
	SetIsCameraLensSmudgeDetectionEnabled(value bool)
	IsConnected() bool
	SetIsConnected(value bool)
	IsContinuityCamera() bool
	SetIsContinuityCamera(value bool)
	IsFollowingExternalSyncDevice() bool
	SetIsFollowingExternalSyncDevice(value bool)
	IsInUseByAnotherApplication() bool
	SetIsInUseByAnotherApplication(value bool)
	IsSubjectAreaChangeMonitoringEnabled() bool
	SetIsSubjectAreaChangeMonitoringEnabled(value bool)
	IsSuspended() bool
	SetIsSuspended(value bool)
	IsVideoFrameDurationLocked() bool
	SetIsVideoFrameDurationLocked(value bool)
	IsVirtualDevice() bool
	SetIsVirtualDevice(value bool)
	AVCaptureSessionInterruptionSystemPressureStateKey() foundation.foundation.INSString


	

	// methods:
	DefaultRectForExposurePointOfInterest(pointOfInterest corefoundation.CGPoint) corefoundation.CGRect
	DefaultRectForFocusPointOfInterest(pointOfInterest corefoundation.CGPoint) corefoundation.CGRect
	HasMediaType(mediaType MediaType) bool
	IsExposureModeSupported(exposureMode CaptureExposureMode) bool
	IsFocusModeSupported(focusMode CaptureFocusMode) bool
	IsTorchModeSupported(torchMode CaptureTorchMode) bool
	IsWhiteBalanceModeSupported(whiteBalanceMode CaptureWhiteBalanceMode) bool
	LockForConfiguration(outError foundation.foundation.INSError) bool
	PerformEffectForReaction(reactionType CaptureReactionType)
	SetCameraLensSmudgeDetectionEnabledDetectionInterval(cameraLensSmudgeDetectionEnabled bool, detectionInterval objectivec.IObject)
	SetCinematicVideoFixedFocusAtPointFocusMode(point corefoundation.CGPoint, focusMode CaptureCinematicVideoFocusMode)
	SetCinematicVideoTrackingFocusAtPointFocusMode(point corefoundation.CGPoint, focusMode CaptureCinematicVideoFocusMode)
	SetCinematicVideoTrackingFocusWithDetectedObjectIDFocusMode(detectedObjectID int, focusMode CaptureCinematicVideoFocusMode)
	SetPrimaryConstituentDeviceSwitchingBehaviorRestrictedSwitchingBehaviorConditions(switchingBehavior CapturePrimaryConstituentDeviceSwitchingBehavior, restrictedSwitchingBehaviorConditions CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions)
	SetTorchModeOnWithLevelError(torchLevel float32, outError foundation.foundation.INSError) bool
	SetTransportControlsPlaybackModeSpeed(mode CaptureDeviceTransportControlsPlaybackMode, speed CaptureDeviceTransportControlsSpeed)
	SupportsAVCaptureSessionPreset(preset CaptureSessionPreset) bool
	UnlockForConfiguration()


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureDeviceClass) Alloc() CaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureDeviceClass) New() CaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureDevice) Init() CaptureDevice {
	rv := objc.Send[CaptureDevice](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureDevice) Autorelease() CaptureDevice {
	rv := objc.Send[CaptureDevice](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureDevice creates a new CaptureDevice instance.
func NewCaptureDevice() CaptureDevice {
	return getCaptureDeviceClass().New()
}





// An object that represents a hardware or virtual capture device like a camera or microphone.
//
// Capture devices provide media data to capture session inputs that you connect to an . An individual device can provide one or more streams of media of a particular type. You don’t create capture device instances directly. Instead, retrieve them using an instance of , or by calling the method. A capture device provides several configuration options. Before attempting to configure device properties, such as its focus mode, exposure mode, and so on, you must first acquire a lock on the device by calling the method. You should also query the device’s capabilities to ensure that the new modes you intend to set are valid for the device. You can then set the properties and release the lock using the method. You may hold the lock if you want all settable device properties to remain unchanged. However, holding the device lock unnecessarily may degrade capture quality in other apps sharing the device and isn’t recommended.


// An object that represents a hardware or virtual capture device like a camera or microphone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice
type CaptureDevice struct {
	objectivec.Object
}

// CaptureDeviceFrom constructs a [CaptureDevice] from an unsafe.Pointer.
//
// An object that represents a hardware or virtual capture device like a camera or microphone.
func CaptureDeviceFrom(ptr unsafe.Pointer) CaptureDevice {
	return CaptureDevice{objectivec.Object{objc.ID(ptr)}}
}






// Creates an object that represents a device with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/init(uniqueID:)
func NewCaptureDeviceWithUniqueID(deviceUniqueID foundation.foundation.INSString) CaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(getCaptureDeviceClass().class), objc.Sel("deviceWithUniqueID:"), deviceUniqueID)
	return rv
}







// Returns an authorization status that indicates whether the user grants the app permission to capture media of a particular type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/authorizationStatus(for:)
func (cc _CaptureDeviceClass) AuthorizationStatusForMediaType(mediaType MediaType) AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](objc.ID(cc.class), objc.Sel("authorizationStatusForMediaType:"), mediaType)
	return rv
}


// Returns the default device for the specified device type, media type, and position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/default(_:for:position:)
func (cc _CaptureDeviceClass) DefaultDeviceWithDeviceTypeMediaTypePosition(deviceType CaptureDeviceType, mediaType MediaType, position CaptureDevicePosition) ICaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(cc.class), objc.Sel("defaultDeviceWithDeviceType:mediaType:position:"), deviceType, mediaType, position)
	return rv
}


// Returns the default device that captures the specified media type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/default(for:)
func (cc _CaptureDeviceClass) DefaultDeviceWithMediaType(mediaType MediaType) ICaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(cc.class), objc.Sel("defaultDeviceWithMediaType:"), mediaType)
	return rv
}


// Returns all available capture devices on the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/devices()
func (cc _CaptureDeviceClass) Devices() []CaptureDevice {
	rv := objc.Send[[]CaptureDevice](objc.ID(cc.class), objc.Sel("devices"))
	return rv
}


// Returns devices capable of capturing media of the specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/devices(for:)
func (cc _CaptureDeviceClass) DevicesWithMediaType(mediaType MediaType) []CaptureDevice {
	rv := objc.Send[[]CaptureDevice](objc.ID(cc.class), objc.Sel("devicesWithMediaType:"), mediaType)
	return rv
}


// Returns the relative extrinsic matrix from one capture device to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/extrinsicMatrix(from:to:)
func (cc _CaptureDeviceClass) ExtrinsicMatrixFromDeviceToDevice(fromDevice IAVCaptureDevice, toDevice IAVCaptureDevice) foundation.Data {
	rv := objc.Send[foundation.Data](objc.ID(cc.class), objc.Sel("extrinsicMatrixFromDevice:toDevice:"), fromDevice, toDevice)
	return rv
}


// Creates an object that represents a device with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/init(uniqueID:)
func (cc _CaptureDeviceClass) DeviceWithUniqueID(deviceUniqueID foundation.foundation.INSString) ICaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(cc.class), objc.Sel("deviceWithUniqueID:"), deviceUniqueID)
	return rv
}


// Requests the user’s permission to allow the app to capture media of a particular type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/requestAccess(for:completionHandler:)
func (cc _CaptureDeviceClass) RequestAccessForMediaTypeCompletionHandler(mediaType MediaType, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("requestAccessForMediaType:completionHandler:"), mediaType, handler)
}


// Displays the system’s user interface to configure video effects or microphone modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/showSystemUserInterface(_:)
func (cc _CaptureDeviceClass) ShowSystemUserInterface(systemUserInterface CaptureSystemUserInterface) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("showSystemUserInterface:"), systemUserInterface)
}







// The device’s active microphone mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeMicrophoneMode
func (cc _CaptureDeviceClass) ActiveMicrophoneMode() CaptureMicrophoneMode {
	rv := objc.Send[CaptureMicrophoneMode](objc.ID(cc.class), objc.Sel("activeMicrophoneMode"))
	return rv
}

// A value that indicates the current mode of Center Stage control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/centerStageControlMode-swift.type.property
func (cc _CaptureDeviceClass) CenterStageControlMode() CaptureCenterStageControlMode {
	rv := objc.Send[CaptureCenterStageControlMode](objc.ID(cc.class), objc.Sel("centerStageControlMode"))
	return rv
}

// A class property that indicates whether a person enables the Background Replacement feature for this app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isBackgroundReplacementEnabled
func (cc _CaptureDeviceClass) BackgroundReplacementEnabled() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("backgroundReplacementEnabled"))
	return rv
}

// A Boolean value that indicates whether a user or an app enabled Center Stage on a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isCenterStageEnabled
func (cc _CaptureDeviceClass) CenterStageEnabled() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("centerStageEnabled"))
	return rv
}

// A Boolean value that indicates whether the user enabled the Portrait video effect in Control Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isPortraitEffectEnabled
func (cc _CaptureDeviceClass) PortraitEffectEnabled() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("portraitEffectEnabled"))
	return rv
}

// A Boolean value that indicates whether a user enabled Studio Light on a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isStudioLightEnabled
func (cc _CaptureDeviceClass) StudioLightEnabled() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("studioLightEnabled"))
	return rv
}

// The microphone mode that the user selects in Control Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/preferredMicrophoneMode
func (cc _CaptureDeviceClass) PreferredMicrophoneMode() CaptureMicrophoneMode {
	rv := objc.Send[CaptureMicrophoneMode](objc.ID(cc.class), objc.Sel("preferredMicrophoneMode"))
	return rv
}

// A Boolean value that indicates whether gesture detection triggers reaction effects on the video stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/reactionEffectGesturesEnabled
func (cc _CaptureDeviceClass) ReactionEffectGesturesEnabled() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("reactionEffectGesturesEnabled"))
	return rv
}

// A Boolean value that indicates whether the app supports performing reaction effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/reactionEffectsEnabled
func (cc _CaptureDeviceClass) ReactionEffectsEnabled() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("reactionEffectsEnabled"))
	return rv
}

// A camera the system prefers to use for video and photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/systemPreferredCamera
func (cc _CaptureDeviceClass) SystemPreferredCamera() CaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(cc.class), objc.Sel("systemPreferredCamera"))
	return rv
}

// A camera the user prefers to use for video and photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/userPreferredCamera
func (cc _CaptureDeviceClass) UserPreferredCamera() CaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(cc.class), objc.Sel("userPreferredCamera"))
	return rv
}






// The default rectangle of interest used for a given exposure point of interest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/defaultRectForExposurePoint(ofInterest:)
func (c_ CaptureDevice) DefaultRectForExposurePointOfInterest(pointOfInterest corefoundation.CGPoint) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("defaultRectForExposurePointOfInterest:"), pointOfInterest)
	return rv
}


// The default rectangle of interest used for a given focus point of interest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/defaultRectForFocusPoint(ofInterest:)
func (c_ CaptureDevice) DefaultRectForFocusPointOfInterest(pointOfInterest corefoundation.CGPoint) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("defaultRectForFocusPointOfInterest:"), pointOfInterest)
	return rv
}


// Returns a Boolean value that indicates whether the device captures media of a particular type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/hasMediaType(_:)
func (c_ CaptureDevice) HasMediaType(mediaType MediaType) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasMediaType:"), mediaType)
	return rv
}


// Returns a Boolean value that indicates whether a device supports the specified exposure mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isExposureModeSupported(_:)
func (c_ CaptureDevice) IsExposureModeSupported(exposureMode CaptureExposureMode) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isExposureModeSupported:"), exposureMode)
	return rv
}


// Returns a Boolean value that indicates whether the device supports the specified focus mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFocusModeSupported(_:)
func (c_ CaptureDevice) IsFocusModeSupported(focusMode CaptureFocusMode) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFocusModeSupported:"), focusMode)
	return rv
}


// Returns a Boolean value that indicates whether the device supports the specified torch mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isTorchModeSupported(_:)
func (c_ CaptureDevice) IsTorchModeSupported(torchMode CaptureTorchMode) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isTorchModeSupported:"), torchMode)
	return rv
}


// Returns a Boolean value that indicates whether the device supports the specified white balance mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isWhiteBalanceModeSupported(_:)
func (c_ CaptureDevice) IsWhiteBalanceModeSupported(whiteBalanceMode CaptureWhiteBalanceMode) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isWhiteBalanceModeSupported:"), whiteBalanceMode)
	return rv
}


// Requests exclusive access to configure device hardware properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/lockForConfiguration()
func (c_ CaptureDevice) LockForConfiguration(outError foundation.foundation.INSError) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("lockForConfiguration:"), outError)
	return rv
}


// Performs the specified reaction type on the video stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/performEffect(for:)
func (c_ CaptureDevice) PerformEffectForReaction(reactionType CaptureReactionType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("performEffectForReaction:"), reactionType)
}


// Specify whether to enable camera lens smudge detection, and the interval time between each run of detections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setCameraLensSmudgeDetectionEnabled(_:detectionInterval:)
func (c_ CaptureDevice) SetCameraLensSmudgeDetectionEnabledDetectionInterval(cameraLensSmudgeDetectionEnabled bool, detectionInterval objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCameraLensSmudgeDetectionEnabled:detectionInterval:"), cameraLensSmudgeDetectionEnabled, detectionInterval)
}


// Fix focus at a distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setCinematicVideoFixedFocus(at:focusMode:)
func (c_ CaptureDevice) SetCinematicVideoFixedFocusAtPointFocusMode(point corefoundation.CGPoint, focusMode CaptureCinematicVideoFocusMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCinematicVideoFixedFocusAtPoint:focusMode:"), point, focusMode)
}


// Focus on and start tracking an object if it can be detected at the region specified by the point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setCinematicVideoTrackingFocus(at:focusMode:)
func (c_ CaptureDevice) SetCinematicVideoTrackingFocusAtPointFocusMode(point corefoundation.CGPoint, focusMode CaptureCinematicVideoFocusMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCinematicVideoTrackingFocusAtPoint:focusMode:"), point, focusMode)
}


// Focus on and start tracking a detected object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setCinematicVideoTrackingFocus(detectedObjectID:focusMode:)
func (c_ CaptureDevice) SetCinematicVideoTrackingFocusWithDetectedObjectIDFocusMode(detectedObjectID int, focusMode CaptureCinematicVideoFocusMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCinematicVideoTrackingFocusWithDetectedObjectID:focusMode:"), detectedObjectID, focusMode)
}


// Sets the switching behavior of the primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setPrimaryConstituentDeviceSwitchingBehavior(_:restrictedSwitchingBehaviorConditions:)
func (c_ CaptureDevice) SetPrimaryConstituentDeviceSwitchingBehaviorRestrictedSwitchingBehaviorConditions(switchingBehavior CapturePrimaryConstituentDeviceSwitchingBehavior, restrictedSwitchingBehaviorConditions CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryConstituentDeviceSwitchingBehavior:restrictedSwitchingBehaviorConditions:"), switchingBehavior, restrictedSwitchingBehaviorConditions)
}


// Sets the illumination level when in torch mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setTorchModeOn(level:)
func (c_ CaptureDevice) SetTorchModeOnWithLevelError(torchLevel float32, outError foundation.foundation.INSError) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("setTorchModeOnWithLevel:error:"), torchLevel, outError)
	return rv
}


// Sets the transport control’s playback mode and speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setTransportControlsPlaybackMode(_:speed:)
func (c_ CaptureDevice) SetTransportControlsPlaybackModeSpeed(mode CaptureDeviceTransportControlsPlaybackMode, speed CaptureDeviceTransportControlsSpeed) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTransportControlsPlaybackMode:speed:"), mode, speed)
}


// Returns a Boolean value that indicates whether you can use the device with capture session configured with the specified preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/supportsSessionPreset(_:)
func (c_ CaptureDevice) SupportsAVCaptureSessionPreset(preset CaptureSessionPreset) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsAVCaptureSessionPreset:"), preset)
	return rv
}


// Releases exclusive control over device hardware properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/unlockForConfiguration()
func (c_ CaptureDevice) UnlockForConfiguration() {
	objc.Send[objc.ID](c_.ID, objc.Sel("unlockForConfiguration"))
}







// The currently active color space for capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeColorSpace
func (c_ CaptureDevice) ActiveColorSpace() CaptureColorSpace {
	rv := objc.Send[CaptureColorSpace](c_.ID, objc.Sel("activeColorSpace"))
	return rv
}


// The currently active color space for capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeColorSpace
func (c_ CaptureDevice) SetActiveColorSpace(value CaptureColorSpace) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveColorSpace:"), value)
}


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeFormat
func (c_ CaptureDevice) ActiveFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("activeFormat"))
	return rv
}


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeFormat
func (c_ CaptureDevice) SetActiveFormat(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveFormat:"), value)
}


// The currently active input source of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeInputSource
func (c_ CaptureDevice) ActiveInputSource() IAVCaptureDeviceInputSource {
	rv := objc.Send[CaptureDeviceInputSource](c_.ID, objc.Sel("activeInputSource"))
	return rv
}


// The currently active input source of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeInputSource
func (c_ CaptureDevice) SetActiveInputSource(value IAVCaptureDeviceInputSource) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveInputSource:"), value)
}


// The device’s active microphone mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeMicrophoneMode
func (c_ CaptureDevice) ActiveMicrophoneMode() CaptureMicrophoneMode {
	rv := objc.Send[CaptureMicrophoneMode](c_.ID, objc.Sel("activeMicrophoneMode"))
	return rv
}


// A virtual device’s active primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activePrimaryConstituent
func (c_ CaptureDevice) ActivePrimaryConstituentDevice() IAVCaptureDevice {
	rv := objc.Send[CaptureDevice](c_.ID, objc.Sel("activePrimaryConstituentDevice"))
	return rv
}


// The conditions that restrict camera switching behavior for the active primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions
func (c_ CaptureDevice) ActivePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions {
	rv := objc.Send[CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions](c_.ID, objc.Sel("activePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions"))
	return rv
}


// The switching behavior of the active constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activePrimaryConstituentDeviceSwitchingBehavior
func (c_ CaptureDevice) ActivePrimaryConstituentDeviceSwitchingBehavior() CapturePrimaryConstituentDeviceSwitchingBehavior {
	rv := objc.Send[CapturePrimaryConstituentDeviceSwitchingBehavior](c_.ID, objc.Sel("activePrimaryConstituentDeviceSwitchingBehavior"))
	return rv
}


// The currently active maximum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeVideoMaxFrameDuration
func (c_ CaptureDevice) ActiveVideoMaxFrameDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("activeVideoMaxFrameDuration"))
	return rv
}


// The currently active maximum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeVideoMaxFrameDuration
func (c_ CaptureDevice) SetActiveVideoMaxFrameDuration(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveVideoMaxFrameDuration:"), value)
}


// The currently active minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeVideoMinFrameDuration
func (c_ CaptureDevice) ActiveVideoMinFrameDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("activeVideoMinFrameDuration"))
	return rv
}


// The currently active minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeVideoMinFrameDuration
func (c_ CaptureDevice) SetActiveVideoMinFrameDuration(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveVideoMinFrameDuration:"), value)
}


// A set of reactions types that a device supports performing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/availableReactionTypes
func (c_ CaptureDevice) AvailableReactionTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availableReactionTypes"))
	return rv
}


// The camera lens smudge detection interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/cameraLensSmudgeDetectionInterval
func (c_ CaptureDevice) CameraLensSmudgeDetectionInterval() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("cameraLensSmudgeDetectionInterval"))
	return rv
}


// A value specifying the status of camera lens smudge detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/cameraLensSmudgeDetectionStatus
func (c_ CaptureDevice) CameraLensSmudgeDetectionStatus() CaptureCameraLensSmudgeDetectionStatus {
	rv := objc.Send[CaptureCameraLensSmudgeDetectionStatus](c_.ID, objc.Sel("cameraLensSmudgeDetectionStatus"))
	return rv
}


// A Boolean value that indicates whether you can perform reaction effects on a capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/canPerformReactionEffects
func (c_ CaptureDevice) CanPerformReactionEffects() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canPerformReactionEffects"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/centerStageRectOfInterestSupported
func (c_ CaptureDevice) CenterStageRectOfInterestSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("centerStageRectOfInterestSupported"))
	return rv
}


// A value that indicates the current mode of Center Stage control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/centerStageControlMode-swift.type.property
func (c_ CaptureDevice) CenterStageControlMode() CaptureCenterStageControlMode {
	rv := objc.Send[CaptureCenterStageControlMode](c_.ID, objc.Sel("centerStageControlMode"))
	return rv
}


// A value that indicates the current mode of Center Stage control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/centerStageControlMode-swift.type.property
func (c_ CaptureDevice) SetCenterStageControlMode(value CaptureCenterStageControlMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCenterStageControlMode:"), value)
}


// The effective region within the output pixel buffer to perform Center Stage framing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/centerStageRectOfInterest
func (c_ CaptureDevice) CenterStageRectOfInterest() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("centerStageRectOfInterest"))
	return rv
}


// The effective region within the output pixel buffer to perform Center Stage framing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/centerStageRectOfInterest
func (c_ CaptureDevice) SetCenterStageRectOfInterest(value corefoundation.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCenterStageRectOfInterest:"), value)
}


// The current scene monitoring statuses related to Cinematic Video capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/cinematicVideoCaptureSceneMonitoringStatuses
func (c_ CaptureDevice) CinematicVideoCaptureSceneMonitoringStatuses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("cinematicVideoCaptureSceneMonitoringStatuses"))
	return rv
}


// A Desk View camera associated with a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/companionDeskViewCamera
func (c_ CaptureDevice) CompanionDeskViewCamera() IAVCaptureDevice {
	rv := objc.Send[CaptureDevice](c_.ID, objc.Sel("companionDeskViewCamera"))
	return rv
}


// The type of device, such as a built-in microphone or wide-angle camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/deviceType-swift.property
func (c_ CaptureDevice) DeviceType() CaptureDeviceType {
	rv := objc.Send[CaptureDeviceType](c_.ID, objc.Sel("deviceType"))
	return rv
}


// A video zoom factor multiplier to use when displaying zoom information in a user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/displayVideoZoomFactorMultiplier
func (c_ CaptureDevice) DisplayVideoZoomFactorMultiplier() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("displayVideoZoomFactorMultiplier"))
	return rv
}


// The exposure mode for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposureMode-swift.property
func (c_ CaptureDevice) ExposureMode() CaptureExposureMode {
	rv := objc.Send[CaptureExposureMode](c_.ID, objc.Sel("exposureMode"))
	return rv
}


// The exposure mode for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposureMode-swift.property
func (c_ CaptureDevice) SetExposureMode(value CaptureExposureMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureMode:"), value)
}


// The point of interest for exposure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposurePointOfInterest
func (c_ CaptureDevice) ExposurePointOfInterest() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c_.ID, objc.Sel("exposurePointOfInterest"))
	return rv
}


// The point of interest for exposure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposurePointOfInterest
func (c_ CaptureDevice) SetExposurePointOfInterest(value corefoundation.CGPoint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposurePointOfInterest:"), value)
}


// The device’s current exposure rectangle of interest, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposureRectOfInterest
func (c_ CaptureDevice) ExposureRectOfInterest() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("exposureRectOfInterest"))
	return rv
}


// The device’s current exposure rectangle of interest, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposureRectOfInterest
func (c_ CaptureDevice) SetExposureRectOfInterest(value corefoundation.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureRectOfInterest:"), value)
}


// The fallback devices to use when a constituent device with a longer focal length becomes limited by its light sensitivity or minimum focus distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/fallbackPrimaryConstituentDevices
func (c_ CaptureDevice) FallbackPrimaryConstituentDevices() []CaptureDevice {
	rv := objc.Send[[]CaptureDevice](c_.ID, objc.Sel("fallbackPrimaryConstituentDevices"))
	return rv
}


// The fallback devices to use when a constituent device with a longer focal length becomes limited by its light sensitivity or minimum focus distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/fallbackPrimaryConstituentDevices
func (c_ CaptureDevice) SetFallbackPrimaryConstituentDevices(value []CaptureDevice) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setFallbackPrimaryConstituentDevices:"), nsArray)
}


// The device’s current flash mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/flashMode-swift.property
func (c_ CaptureDevice) FlashMode() CaptureFlashMode {
	rv := objc.Send[CaptureFlashMode](c_.ID, objc.Sel("flashMode"))
	return rv
}


// The device’s current flash mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/flashMode-swift.property
func (c_ CaptureDevice) SetFlashMode(value CaptureFlashMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlashMode:"), value)
}


// The capture device’s focus mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/focusMode-swift.property
func (c_ CaptureDevice) FocusMode() CaptureFocusMode {
	rv := objc.Send[CaptureFocusMode](c_.ID, objc.Sel("focusMode"))
	return rv
}


// The capture device’s focus mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/focusMode-swift.property
func (c_ CaptureDevice) SetFocusMode(value CaptureFocusMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocusMode:"), value)
}


// The point of interest for focusing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/focusPointOfInterest
func (c_ CaptureDevice) FocusPointOfInterest() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c_.ID, objc.Sel("focusPointOfInterest"))
	return rv
}


// The point of interest for focusing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/focusPointOfInterest
func (c_ CaptureDevice) SetFocusPointOfInterest(value corefoundation.CGPoint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocusPointOfInterest:"), value)
}


// The device’s current focus rectangle of interest, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/focusRectOfInterest
func (c_ CaptureDevice) FocusRectOfInterest() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("focusRectOfInterest"))
	return rv
}


// The device’s current focus rectangle of interest, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/focusRectOfInterest
func (c_ CaptureDevice) SetFocusRectOfInterest(value corefoundation.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocusRectOfInterest:"), value)
}


// The capture formats a device supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/formats
func (c_ CaptureDevice) Formats() []CaptureDeviceFormat {
	rv := objc.Send[[]CaptureDeviceFormat](c_.ID, objc.Sel("formats"))
	return rv
}


// A Boolean value that indicates whether the capture device has a flash.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/hasFlash
func (c_ CaptureDevice) HasFlash() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasFlash"))
	return rv
}


// A Boolean value that specifies whether the capture device has a torch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/hasTorch
func (c_ CaptureDevice) HasTorch() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasTorch"))
	return rv
}


// An array of input sources that the device supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/inputSources
func (c_ CaptureDevice) InputSources() []CaptureDeviceInputSource {
	rv := objc.Send[[]CaptureDeviceInputSource](c_.ID, objc.Sel("inputSources"))
	return rv
}


// A Boolean value that indicates whether the device is currently adjusting its exposure setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isAdjustingExposure
func (c_ CaptureDevice) AdjustingExposure() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("adjustingExposure"))
	return rv
}


// A Boolean value that indicates whether the device is currently adjusting its focus setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isAdjustingFocus
func (c_ CaptureDevice) AdjustingFocus() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("adjustingFocus"))
	return rv
}


// A Boolean value that indicates whether the device is currently adjusting the white balance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isAdjustingWhiteBalance
func (c_ CaptureDevice) AdjustingWhiteBalance() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("adjustingWhiteBalance"))
	return rv
}


// A Boolean value that indicates whether the capture device performs automatic video frame rate adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isAutoVideoFrameRateEnabled
func (c_ CaptureDevice) AutoVideoFrameRateEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoVideoFrameRateEnabled"))
	return rv
}


// A Boolean value that indicates whether the capture device performs automatic video frame rate adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isAutoVideoFrameRateEnabled
func (c_ CaptureDevice) SetAutoVideoFrameRateEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutoVideoFrameRateEnabled:"), value)
}


// A Boolean value that indicates whether Background Replacement is currently active on a capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isBackgroundReplacementActive
func (c_ CaptureDevice) BackgroundReplacementActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("backgroundReplacementActive"))
	return rv
}


// A class property that indicates whether a person enables the Background Replacement feature for this app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isBackgroundReplacementEnabled
func (c_ CaptureDevice) BackgroundReplacementEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("backgroundReplacementEnabled"))
	return rv
}


// Whether camera lens smudge detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isCameraLensSmudgeDetectionEnabled
func (c_ CaptureDevice) CameraLensSmudgeDetectionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("cameraLensSmudgeDetectionEnabled"))
	return rv
}


// A Boolean value that indicates whether Center Stage is active on a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isCenterStageActive
func (c_ CaptureDevice) CenterStageActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("centerStageActive"))
	return rv
}


// A Boolean value that indicates whether a user or an app enabled Center Stage on a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isCenterStageEnabled
func (c_ CaptureDevice) CenterStageEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("centerStageEnabled"))
	return rv
}


// A Boolean value that indicates whether a user or an app enabled Center Stage on a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isCenterStageEnabled
func (c_ CaptureDevice) SetCenterStageEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCenterStageEnabled:"), value)
}


// A Boolean value that indicates whether a device is currently connected to the system and available for use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isConnected
func (c_ CaptureDevice) Connected() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("connected"))
	return rv
}


// A Boolean value that indicates whether the device is a Continuity Camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isContinuityCamera
func (c_ CaptureDevice) ContinuityCamera() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("continuityCamera"))
	return rv
}


// A Boolean value that indicates whether the device supports a point of interest for exposure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isExposurePointOfInterestSupported
func (c_ CaptureDevice) ExposurePointOfInterestSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("exposurePointOfInterestSupported"))
	return rv
}


// Whether the device supports exposure rectangles of interest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isExposureRectOfInterestSupported
func (c_ CaptureDevice) ExposureRectOfInterestSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("exposureRectOfInterestSupported"))
	return rv
}


// A Boolean value that indicates whether the flash is currently available for use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFlashAvailable
func (c_ CaptureDevice) FlashAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("flashAvailable"))
	return rv
}


// A Boolean value that indicates whether the device supports a point of interest for focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFocusPointOfInterestSupported
func (c_ CaptureDevice) FocusPointOfInterestSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("focusPointOfInterestSupported"))
	return rv
}


// Whether the receiver supports focus rectangles of interest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFocusRectOfInterestSupported
func (c_ CaptureDevice) FocusRectOfInterestSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("focusRectOfInterestSupported"))
	return rv
}


// Whether the device is following an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFollowingExternalSyncDevice
func (c_ CaptureDevice) FollowingExternalSyncDevice() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("followingExternalSyncDevice"))
	return rv
}


// A Boolean value that indicates whether another app is using the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isInUseByAnotherApplication
func (c_ CaptureDevice) InUseByAnotherApplication() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("inUseByAnotherApplication"))
	return rv
}


// A Boolean value that indicates whether the Portrait video effect is active on a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isPortraitEffectActive
func (c_ CaptureDevice) PortraitEffectActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("portraitEffectActive"))
	return rv
}


// A Boolean value that indicates whether the user enabled the Portrait video effect in Control Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isPortraitEffectEnabled
func (c_ CaptureDevice) PortraitEffectEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("portraitEffectEnabled"))
	return rv
}


// A Boolean value that indicates whether Studio Light is active on a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isStudioLightActive
func (c_ CaptureDevice) StudioLightActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("studioLightActive"))
	return rv
}


// A Boolean value that indicates whether a user enabled Studio Light on a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isStudioLightEnabled
func (c_ CaptureDevice) StudioLightEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("studioLightEnabled"))
	return rv
}


// A Boolean value that indicates whether the device is in a suspended state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isSuspended
func (c_ CaptureDevice) Suspended() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("suspended"))
	return rv
}


// A Boolean value that indicates whether the device’s torch is currently active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isTorchActive
func (c_ CaptureDevice) TorchActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("torchActive"))
	return rv
}


// A Boolean value that indicates whether the torch is currently available for use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isTorchAvailable
func (c_ CaptureDevice) TorchAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("torchAvailable"))
	return rv
}


// Whether the device’s video frame rate (expressed as a duration) is currently locked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isVideoFrameDurationLocked
func (c_ CaptureDevice) VideoFrameDurationLocked() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("videoFrameDurationLocked"))
	return rv
}


// An array of capture devices that are physically linked to a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/linkedDevices
func (c_ CaptureDevice) LinkedDevices() []CaptureDevice {
	rv := objc.Send[[]CaptureDevice](c_.ID, objc.Sel("linkedDevices"))
	return rv
}


// A localized device name for display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/localizedName
func (c_ CaptureDevice) LocalizedName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedName"))
	return rv
}


// A human-readable string for the manufacturer of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/manufacturer
func (c_ CaptureDevice) Manufacturer() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("manufacturer"))
	return rv
}


// The minimum size you may use when specifying a rectangle of interest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minExposureRectOfInterestSize
func (c_ CaptureDevice) MinExposureRectOfInterestSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("minExposureRectOfInterestSize"))
	return rv
}


// The minimum size you may use when specifying a rectangle of interest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minFocusRectOfInterestSize
func (c_ CaptureDevice) MinFocusRectOfInterestSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("minFocusRectOfInterestSize"))
	return rv
}


// The capture device’s minimum focus distance in millimeters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minimumFocusDistance
func (c_ CaptureDevice) MinimumFocusDistance() int {
	rv := objc.Send[int](c_.ID, objc.Sel("minimumFocusDistance"))
	return rv
}


// The minimum frame duration that can be passed as the when directing your device input to follow an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minSupportedExternalSyncFrameDuration
func (c_ CaptureDevice) MinSupportedExternalSyncFrameDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("minSupportedExternalSyncFrameDuration"))
	return rv
}


// The maximum frame rate (expressed as a minimum duration) that can be set on an input associated with this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minSupportedLockedVideoFrameDuration
func (c_ CaptureDevice) MinSupportedLockedVideoFrameDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("minSupportedLockedVideoFrameDuration"))
	return rv
}


// A model identifier for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/modelID
func (c_ CaptureDevice) ModelID() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("modelID"))
	return rv
}


// The physical position of the capture device hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/position-swift.property
func (c_ CaptureDevice) Position() CaptureDevicePosition {
	rv := objc.Send[CaptureDevicePosition](c_.ID, objc.Sel("position"))
	return rv
}


// The microphone mode that the user selects in Control Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/preferredMicrophoneMode
func (c_ CaptureDevice) PreferredMicrophoneMode() CaptureMicrophoneMode {
	rv := objc.Send[CaptureMicrophoneMode](c_.ID, objc.Sel("preferredMicrophoneMode"))
	return rv
}


// The conditions that restrict the primary constituent device’s switching behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/primaryConstituentDeviceRestrictedSwitchingBehaviorConditions-swift.property
func (c_ CaptureDevice) PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions {
	rv := objc.Send[CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions](c_.ID, objc.Sel("primaryConstituentDeviceRestrictedSwitchingBehaviorConditions"))
	return rv
}


// The switching behavior for the primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/primaryConstituentDeviceSwitchingBehavior-swift.property
func (c_ CaptureDevice) PrimaryConstituentDeviceSwitchingBehavior() CapturePrimaryConstituentDeviceSwitchingBehavior {
	rv := objc.Send[CapturePrimaryConstituentDeviceSwitchingBehavior](c_.ID, objc.Sel("primaryConstituentDeviceSwitchingBehavior"))
	return rv
}


// A Boolean value that indicates whether gesture detection triggers reaction effects on the video stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/reactionEffectGesturesEnabled
func (c_ CaptureDevice) ReactionEffectGesturesEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("reactionEffectGesturesEnabled"))
	return rv
}


// A Boolean value that indicates whether the app supports performing reaction effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/reactionEffectsEnabled
func (c_ CaptureDevice) ReactionEffectsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("reactionEffectsEnabled"))
	return rv
}


// An array of reaction effects that the device is currently performing, sorted by timestamp.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/reactionEffectsInProgress
func (c_ CaptureDevice) ReactionEffectsInProgress() []CaptureReactionEffectState {
	rv := objc.Send[[]CaptureReactionEffectState](c_.ID, objc.Sel("reactionEffectsInProgress"))
	return rv
}


// Reasons why current environmental conditions aren’t suitable to capturing spatial videos that are comfortable to view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/spatialCaptureDiscomfortReasons
func (c_ CaptureDevice) SpatialCaptureDiscomfortReasons() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("spatialCaptureDiscomfortReasons"))
	return rv
}


// The constituent devices available to select as a fallback for a longer focal length primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/supportedFallbackPrimaryConstituentDevices
func (c_ CaptureDevice) SupportedFallbackPrimaryConstituentDevices() []CaptureDevice {
	rv := objc.Send[[]CaptureDevice](c_.ID, objc.Sel("supportedFallbackPrimaryConstituentDevices"))
	return rv
}


// A camera the system prefers to use for video and photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/systemPreferredCamera
func (c_ CaptureDevice) SystemPreferredCamera() IAVCaptureDevice {
	rv := objc.Send[CaptureDevice](c_.ID, objc.Sel("systemPreferredCamera"))
	return rv
}


// The current torch brightness level.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/torchLevel
func (c_ CaptureDevice) TorchLevel() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("torchLevel"))
	return rv
}


// The current torch mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/torchMode-swift.property
func (c_ CaptureDevice) TorchMode() CaptureTorchMode {
	rv := objc.Send[CaptureTorchMode](c_.ID, objc.Sel("torchMode"))
	return rv
}


// The current torch mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/torchMode-swift.property
func (c_ CaptureDevice) SetTorchMode(value CaptureTorchMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTorchMode:"), value)
}


// The current playback mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/transportControlsPlaybackMode-swift.property
func (c_ CaptureDevice) TransportControlsPlaybackMode() CaptureDeviceTransportControlsPlaybackMode {
	rv := objc.Send[CaptureDeviceTransportControlsPlaybackMode](c_.ID, objc.Sel("transportControlsPlaybackMode"))
	return rv
}


// The current playback speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/transportControlsSpeed-swift.property
func (c_ CaptureDevice) TransportControlsSpeed() CaptureDeviceTransportControlsSpeed {
	rv := objc.Send[CaptureDeviceTransportControlsSpeed](c_.ID, objc.Sel("transportControlsSpeed"))
	return rv
}


// A Boolean value that indicates whether the device supports transport control commands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/transportControlsSupported
func (c_ CaptureDevice) TransportControlsSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("transportControlsSupported"))
	return rv
}


// The transport type of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/transportType
func (c_ CaptureDevice) TransportType() int32 /* not a class type */ {
	rv := objc.Send[int32](c_.ID, objc.Sel("transportType"))
	return rv
}


// An identifier that uniquely identifies the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/uniqueID
func (c_ CaptureDevice) UniqueID() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("uniqueID"))
	return rv
}


// A camera the user prefers to use for video and photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/userPreferredCamera
func (c_ CaptureDevice) UserPreferredCamera() IAVCaptureDevice {
	rv := objc.Send[CaptureDevice](c_.ID, objc.Sel("userPreferredCamera"))
	return rv
}


// A camera the user prefers to use for video and photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/userPreferredCamera
func (c_ CaptureDevice) SetUserPreferredCamera(value IAVCaptureDevice) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserPreferredCamera:"), value)
}


// The current white balance mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/whiteBalanceMode-swift.property
func (c_ CaptureDevice) WhiteBalanceMode() CaptureWhiteBalanceMode {
	rv := objc.Send[CaptureWhiteBalanceMode](c_.ID, objc.Sel("whiteBalanceMode"))
	return rv
}


// The current white balance mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/whiteBalanceMode-swift.property
func (c_ CaptureDevice) SetWhiteBalanceMode(value CaptureWhiteBalanceMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWhiteBalanceMode:"), value)
}


// A virtual device’s active primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeprimaryconstituent
func (c_ CaptureDevice) ActivePrimaryConstituent() IAVCaptureDevice {
	rv := objc.Send[CaptureDevice](c_.ID, objc.Sel("activePrimaryConstituent"))
	return rv
}


// A virtual device’s active primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeprimaryconstituent
func (c_ CaptureDevice) SetActivePrimaryConstituent(value IAVCaptureDevice) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActivePrimaryConstituent:"), value)
}


// A Boolean value that indicates whether the capture device performs automatic video frame rate adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isautovideoframerateenabled
func (c_ CaptureDevice) IsAutoVideoFrameRateEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoVideoFrameRateEnabled"))
	return rv
}


// A Boolean value that indicates whether the capture device performs automatic video frame rate adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isautovideoframerateenabled
func (c_ CaptureDevice) SetIsAutoVideoFrameRateEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoVideoFrameRateEnabled:"), value)
}


// Whether camera lens smudge detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/iscameralenssmudgedetectionenabled
func (c_ CaptureDevice) IsCameraLensSmudgeDetectionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraLensSmudgeDetectionEnabled"))
	return rv
}


// Whether camera lens smudge detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/iscameralenssmudgedetectionenabled
func (c_ CaptureDevice) SetIsCameraLensSmudgeDetectionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraLensSmudgeDetectionEnabled:"), value)
}


// A Boolean value that indicates whether a device is currently connected to the system and available for use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isconnected
func (c_ CaptureDevice) IsConnected() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConnected"))
	return rv
}


// A Boolean value that indicates whether a device is currently connected to the system and available for use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isconnected
func (c_ CaptureDevice) SetIsConnected(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConnected:"), value)
}


// A Boolean value that indicates whether the device is a Continuity Camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/iscontinuitycamera
func (c_ CaptureDevice) IsContinuityCamera() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isContinuityCamera"))
	return rv
}


// A Boolean value that indicates whether the device is a Continuity Camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/iscontinuitycamera
func (c_ CaptureDevice) SetIsContinuityCamera(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsContinuityCamera:"), value)
}


// Whether the device is following an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isfollowingexternalsyncdevice
func (c_ CaptureDevice) IsFollowingExternalSyncDevice() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFollowingExternalSyncDevice"))
	return rv
}


// Whether the device is following an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isfollowingexternalsyncdevice
func (c_ CaptureDevice) SetIsFollowingExternalSyncDevice(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFollowingExternalSyncDevice:"), value)
}


// A Boolean value that indicates whether another app is using the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isinusebyanotherapplication
func (c_ CaptureDevice) IsInUseByAnotherApplication() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isInUseByAnotherApplication"))
	return rv
}


// A Boolean value that indicates whether another app is using the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isinusebyanotherapplication
func (c_ CaptureDevice) SetIsInUseByAnotherApplication(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsInUseByAnotherApplication:"), value)
}


// A Boolean value that indicates whether the device monitors the subject area for changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/issubjectareachangemonitoringenabled
func (c_ CaptureDevice) IsSubjectAreaChangeMonitoringEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSubjectAreaChangeMonitoringEnabled"))
	return rv
}


// A Boolean value that indicates whether the device monitors the subject area for changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/issubjectareachangemonitoringenabled
func (c_ CaptureDevice) SetIsSubjectAreaChangeMonitoringEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSubjectAreaChangeMonitoringEnabled:"), value)
}


// A Boolean value that indicates whether the device is in a suspended state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/issuspended
func (c_ CaptureDevice) IsSuspended() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSuspended"))
	return rv
}


// A Boolean value that indicates whether the device is in a suspended state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/issuspended
func (c_ CaptureDevice) SetIsSuspended(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSuspended:"), value)
}


// Whether the device’s video frame rate (expressed as a duration) is currently locked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isvideoframedurationlocked
func (c_ CaptureDevice) IsVideoFrameDurationLocked() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoFrameDurationLocked"))
	return rv
}


// Whether the device’s video frame rate (expressed as a duration) is currently locked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isvideoframedurationlocked
func (c_ CaptureDevice) SetIsVideoFrameDurationLocked(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoFrameDurationLocked:"), value)
}


// A Boolean value that indicates whether the device consists of two or more physical devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isvirtualdevice
func (c_ CaptureDevice) IsVirtualDevice() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVirtualDevice"))
	return rv
}


// A Boolean value that indicates whether the device consists of two or more physical devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isvirtualdevice
func (c_ CaptureDevice) SetIsVirtualDevice(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVirtualDevice:"), value)
}


// A key to retrieve a state value that indicates the system pressure level and contributing factors that caused the interruption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesessioninterruptionsystempressurestatekey
func (c_ CaptureDevice) AVCaptureSessionInterruptionSystemPressureStateKey() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("AVCaptureSessionInterruptionSystemPressureStateKey"))
	return rv
}







