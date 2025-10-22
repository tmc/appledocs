// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	IsFocusModeSupported(focusMode unsafe.Pointer) bool
	LockForConfiguration(outError unsafe.Pointer) bool
	RampToVideoZoomFactorWithRate(factor float64, rate float32)
	SetCameraLensSmudgeDetectionEnabledDetectionInterval(cameraLensSmudgeDetectionEnabled bool, detectionInterval unsafe.Pointer)
	SetCinematicVideoFixedFocusAtPointFocusMode(point coregraphics.CGPoint, focusMode CaptureCinematicVideoFocusMode)
	SetCinematicVideoTrackingFocusAtPointFocusMode(point coregraphics.CGPoint, focusMode CaptureCinematicVideoFocusMode)
	SetDynamicAspectRatioCompletionHandler(dynamicAspectRatio unsafe.Pointer, handler unsafe.Pointer)
	SetPrimaryConstituentDeviceSwitchingBehaviorRestrictedSwitchingBehaviorConditions(switchingBehavior CapturePrimaryConstituentDeviceSwitchingBehavior, restrictedSwitchingBehaviorConditions ICapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions)
	SetWhiteBalanceModeLockedWithDeviceWhiteBalanceGainsCompletionHandler(whiteBalanceGains unsafe.Pointer, handler unsafe.Pointer)
	UnlockForConfiguration()
}

// An object that represents a hardware or virtual capture device like a camera or microphone.
//
// Capture devices provide media data to capture session inputs that you connect to an . An individual device can provide one or more streams of media of a particular type. You don’t create capture device instances directly. Instead, retrieve them using an instance of , or by calling the method. A capture device provides several configuration options. Before attempting to configure device properties, such as its focus mode, exposure mode, and so on, you must first acquire a lock on the device by calling the method. You should also query the device’s capabilities to ensure that the new modes you intend to set are valid for the device. You can then set the properties and release the lock using the method. You may hold the lock if you want all settable device properties to remain unchanged. However, holding the device lock unnecessarily may degrade capture quality in other apps sharing the device and isn’t recommended.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CaptureDeviceClass) Alloc() CaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates an object that represents a device with the specified identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/init(uniqueID:)
func NewCaptureDeviceWithUniqueID(deviceUniqueID string) CaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(getCaptureDeviceClass().class), objc.Sel("deviceWithUniqueID:"), objc.String(deviceUniqueID))
	return rv
}


// Returns an authorization status that indicates whether the user grants the app permission to capture media of a particular type.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/authorizationStatus(for:)
func (cc _CaptureDeviceClass) AuthorizationStatusForMediaType(mediaType MediaType) AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](objc.ID(cc.class), objc.Sel("authorizationStatusForMediaType:"), mediaType)
	return rv
}

// Returns the default device for the specified device type, media type, and position.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/default(_:for:position:)
func (cc _CaptureDeviceClass) DefaultDeviceWithDeviceTypeMediaTypePosition(deviceType unsafe.Pointer, mediaType MediaType, position CaptureDevicePosition) CaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(cc.class), objc.Sel("defaultDeviceWithDeviceType:mediaType:position:"), deviceType, mediaType, position)
	return rv
}

// Returns the default device that captures the specified media type.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/default(for:)
func (cc _CaptureDeviceClass) DefaultDeviceWithMediaType(mediaType MediaType) CaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(cc.class), objc.Sel("defaultDeviceWithMediaType:"), mediaType)
	return rv
}

// Returns devices capable of capturing media of the specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/devices(for:)
func (cc _CaptureDeviceClass) DevicesWithMediaType(mediaType MediaType) []CaptureDevice {
	rv := objc.Send[[]CaptureDevice](objc.ID(cc.class), objc.Sel("devicesWithMediaType:"), mediaType)
	return rv
}

// Creates an object that represents a device with the specified identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/init(uniqueID:)
func (cc _CaptureDeviceClass) DeviceWithUniqueID(deviceUniqueID string) CaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(cc.class), objc.Sel("deviceWithUniqueID:"), objc.String(deviceUniqueID))
	return rv
}

// Requests the user’s permission to allow the app to capture media of a particular type.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/requestAccess(for:completionHandler:)
func (cc _CaptureDeviceClass) RequestAccessForMediaTypeCompletionHandler(mediaType MediaType, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("requestAccessForMediaType:completionHandler:"), mediaType, handler)
}

// The device’s active microphone mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeMicrophoneMode
func (cc _CaptureDeviceClass) ActiveMicrophoneMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("activeMicrophoneMode"))
	return rv
}
// A class property that indicates whether a person enables the Background Replacement feature for this app.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isBackgroundReplacementEnabled
func (cc _CaptureDeviceClass) BackgroundReplacementEnabled() bool {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("backgroundReplacementEnabled"))
	return rv
}
// A camera the system prefers to use for video and photo capture.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/systemPreferredCamera
func (cc _CaptureDeviceClass) SystemPreferredCamera() CaptureDevice {
	rv := objc.Send[AVCaptureDevice](objc.ID(cc.class), objc.Sel("systemPreferredCamera"))
	return rv
}
// A camera the user prefers to use for video and photo capture.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/userPreferredCamera
func (cc _CaptureDeviceClass) UserPreferredCamera() CaptureDevice {
	rv := objc.Send[AVCaptureDevice](objc.ID(cc.class), objc.Sel("userPreferredCamera"))
	return rv
}
// Returns a Boolean value that indicates whether the device supports the specified focus mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFocusModeSupported(_:)
func (c_ CaptureDevice) IsFocusModeSupported(focusMode unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFocusModeSupported:"), focusMode)
	return rv
}

// Requests exclusive access to configure device hardware properties.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/lockForConfiguration()
func (c_ CaptureDevice) LockForConfiguration(outError unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("lockForConfiguration:"), outError)
	return rv
}

// Begins a smooth transition from the current zoom factor to another.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/ramp(toVideoZoomFactor:withRate:)
func (c_ CaptureDevice) RampToVideoZoomFactorWithRate(factor float64, rate float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("rampToVideoZoomFactor:withRate:"), factor, rate)
}

// Specify whether to enable camera lens smudge detection, and the interval time between each run of detections.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setCameraLensSmudgeDetectionEnabled(_:detectionInterval:)
func (c_ CaptureDevice) SetCameraLensSmudgeDetectionEnabledDetectionInterval(cameraLensSmudgeDetectionEnabled bool, detectionInterval unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCameraLensSmudgeDetectionEnabled:detectionInterval:"), cameraLensSmudgeDetectionEnabled, detectionInterval)
}

// Fix focus at a distance.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setCinematicVideoFixedFocus(at:focusMode:)
func (c_ CaptureDevice) SetCinematicVideoFixedFocusAtPointFocusMode(point coregraphics.CGPoint, focusMode CaptureCinematicVideoFocusMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCinematicVideoFixedFocusAtPoint:focusMode:"), point, focusMode)
}

// Focus on and start tracking an object if it can be detected at the region specified by the point.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setCinematicVideoTrackingFocus(at:focusMode:)
func (c_ CaptureDevice) SetCinematicVideoTrackingFocusAtPointFocusMode(point coregraphics.CGPoint, focusMode CaptureCinematicVideoFocusMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCinematicVideoTrackingFocusAtPoint:focusMode:"), point, focusMode)
}

// Updates the dynamic aspect ratio of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setDynamicAspectRatio(_:completionHandler:)
func (c_ CaptureDevice) SetDynamicAspectRatioCompletionHandler(dynamicAspectRatio unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDynamicAspectRatio:completionHandler:"), dynamicAspectRatio, handler)
}

// Sets the switching behavior of the primary constituent device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setPrimaryConstituentDeviceSwitchingBehavior(_:restrictedSwitchingBehaviorConditions:)
func (c_ CaptureDevice) SetPrimaryConstituentDeviceSwitchingBehaviorRestrictedSwitchingBehaviorConditions(switchingBehavior CapturePrimaryConstituentDeviceSwitchingBehavior, restrictedSwitchingBehaviorConditions ICapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryConstituentDeviceSwitchingBehavior:restrictedSwitchingBehaviorConditions:"), switchingBehavior, restrictedSwitchingBehaviorConditions)
}

// Sets the white balance to locked mode with the specified white balance gains.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setWhiteBalanceModeLocked(with:completionHandler:)
func (c_ CaptureDevice) SetWhiteBalanceModeLockedWithDeviceWhiteBalanceGainsCompletionHandler(whiteBalanceGains unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWhiteBalanceModeLockedWithDeviceWhiteBalanceGains:completionHandler:"), whiteBalanceGains, handler)
}

// Releases exclusive control over device hardware properties.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/unlockForConfiguration()
func (c_ CaptureDevice) UnlockForConfiguration() {
	objc.Send[objc.ID](c_.ID, objc.Sel("unlockForConfiguration"))
}

// The minimum frame duration of depth data.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeDepthDataMinFrameDuration
func (c_ CaptureDevice) ActiveDepthDataMinFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("activeDepthDataMinFrameDuration"))
	return rv
}


// SetActiveDepthDataMinFrameDuration sets the value of the activeDepthDataMinFrameDuration property.
// The minimum frame duration of depth data.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeDepthDataMinFrameDuration
func (c_ CaptureDevice) SetActiveDepthDataMinFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveDepthDataMinFrameDuration:"), value)
}

// The capture format in use by the device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeFormat
func (c_ CaptureDevice) ActiveFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("activeFormat"))
	return rv
}


// SetActiveFormat sets the value of the activeFormat property.
// The capture format in use by the device.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeFormat
func (c_ CaptureDevice) SetActiveFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveFormat:"), value)
}

// The currently active input source of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeInputSource
func (c_ CaptureDevice) ActiveInputSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("activeInputSource"))
	return rv
}


// SetActiveInputSource sets the value of the activeInputSource property.
// The currently active input source of the device.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeInputSource
func (c_ CaptureDevice) SetActiveInputSource(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveInputSource:"), value)
}

// The device’s active microphone mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeMicrophoneMode
func (c_ CaptureDevice) ActiveMicrophoneMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("activeMicrophoneMode"))
	return rv
}

// A virtual device’s active primary constituent device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activePrimaryConstituent
func (c_ CaptureDevice) ActivePrimaryConstituentDevice() AVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](c_.ID, objc.Sel("activePrimaryConstituentDevice"))
	return rv
}

// The conditions that restrict camera switching behavior for the active primary constituent device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions
func (c_ CaptureDevice) ActivePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions {
	rv := objc.Send[CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions](c_.ID, objc.Sel("activePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions"))
	return rv
}

// The switching behavior of the active constituent device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activePrimaryConstituentDeviceSwitchingBehavior
func (c_ CaptureDevice) ActivePrimaryConstituentDeviceSwitchingBehavior() CapturePrimaryConstituentDeviceSwitchingBehavior {
	rv := objc.Send[CapturePrimaryConstituentDeviceSwitchingBehavior](c_.ID, objc.Sel("activePrimaryConstituentDeviceSwitchingBehavior"))
	return rv
}

// The camera lens smudge detection interval.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/cameraLensSmudgeDetectionInterval
func (c_ CaptureDevice) CameraLensSmudgeDetectionInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("cameraLensSmudgeDetectionInterval"))
	return rv
}

// A value specifying the status of camera lens smudge detection.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/cameraLensSmudgeDetectionStatus
func (c_ CaptureDevice) CameraLensSmudgeDetectionStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("cameraLensSmudgeDetectionStatus"))
	return rv
}

// A Boolean value that indicates whether you can perform reaction effects on a capture device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/canPerformReactionEffects
func (c_ CaptureDevice) CanPerformReactionEffects() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canPerformReactionEffects"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/centerStageRectOfInterestSupported
func (c_ CaptureDevice) CenterStageRectOfInterestSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("centerStageRectOfInterestSupported"))
	return rv
}

// A Desk View camera associated with a device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/companionDeskViewCamera
func (c_ CaptureDevice) CompanionDeskViewCamera() AVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](c_.ID, objc.Sel("companionDeskViewCamera"))
	return rv
}

// An array of physical devices that make up a virtual device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/constituentDevices
func (c_ CaptureDevice) ConstituentDevices() []CaptureDevice {
	rv := objc.Send[[]CaptureDevice](c_.ID, objc.Sel("constituentDevices"))
	return rv
}

// A key-value observable property indicating the current aspect ratio for a device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/dynamicAspectRatio
func (c_ CaptureDevice) DynamicAspectRatio() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("dynamicAspectRatio"))
	return rv
}

// A key-value observable property describing the output dimensions of the video buffer based on the device’s dynamic aspect ratio.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/dynamicDimensions
func (c_ CaptureDevice) DynamicDimensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("dynamicDimensions"))
	return rv
}

// The exposure mode for the device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposureMode-swift.property
func (c_ CaptureDevice) ExposureMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("exposureMode"))
	return rv
}


// SetExposureMode sets the value of the exposureMode property.
// The exposure mode for the device.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposureMode-swift.property
func (c_ CaptureDevice) SetExposureMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureMode:"), value)
}

// The fallback devices to use when a constituent device with a longer focal length becomes limited by its light sensitivity or minimum focus distance.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/fallbackPrimaryConstituentDevices
func (c_ CaptureDevice) FallbackPrimaryConstituentDevices() []CaptureDevice {
	rv := objc.Send[[]CaptureDevice](c_.ID, objc.Sel("fallbackPrimaryConstituentDevices"))
	return rv
}


// SetFallbackPrimaryConstituentDevices sets the value of the fallbackPrimaryConstituentDevices property.
// The fallback devices to use when a constituent device with a longer focal length becomes limited by its light sensitivity or minimum focus distance.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/fallbackPrimaryConstituentDevices
func (c_ CaptureDevice) SetFallbackPrimaryConstituentDevices(value []CaptureDevice) {
	// Convert Go slice to NSArray
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

// The capture device’s focus mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/focusMode-swift.property
func (c_ CaptureDevice) FocusMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("focusMode"))
	return rv
}


// SetFocusMode sets the value of the focusMode property.
// The capture device’s focus mode.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/focusMode-swift.property
func (c_ CaptureDevice) SetFocusMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocusMode:"), value)
}

// The device’s current focus rectangle of interest, if it has one.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/focusRectOfInterest
func (c_ CaptureDevice) FocusRectOfInterest() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](c_.ID, objc.Sel("focusRectOfInterest"))
	return rv
}


// SetFocusRectOfInterest sets the value of the focusRectOfInterest property.
// The device’s current focus rectangle of interest, if it has one.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/focusRectOfInterest
func (c_ CaptureDevice) SetFocusRectOfInterest(value coregraphics.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocusRectOfInterest:"), value)
}

// A Boolean value that indicates whether the capture device performs automatic video frame rate adjustments.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isAutoVideoFrameRateEnabled
func (c_ CaptureDevice) AutoVideoFrameRateEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoVideoFrameRateEnabled"))
	return rv
}


// SetAutoVideoFrameRateEnabled sets the value of the autoVideoFrameRateEnabled property.
// A Boolean value that indicates whether the capture device performs automatic video frame rate adjustments.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isAutoVideoFrameRateEnabled
func (c_ CaptureDevice) SetAutoVideoFrameRateEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutoVideoFrameRateEnabled:"), value)
}

// A class property that indicates whether a person enables the Background Replacement feature for this app.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isBackgroundReplacementEnabled
func (c_ CaptureDevice) BackgroundReplacementEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("backgroundReplacementEnabled"))
	return rv
}

// Whether camera lens smudge detection is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isCameraLensSmudgeDetectionEnabled
func (c_ CaptureDevice) CameraLensSmudgeDetectionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("cameraLensSmudgeDetectionEnabled"))
	return rv
}

// A Boolean value that indicates whether Center Stage is active on a device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isCenterStageActive
func (c_ CaptureDevice) CenterStageActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("centerStageActive"))
	return rv
}

// A Boolean value that indicates whether a device is currently connected to the system and available for use.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isConnected
func (c_ CaptureDevice) Connected() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("connected"))
	return rv
}

// A Boolean value that indicates whether the device is a Continuity Camera.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isContinuityCamera
func (c_ CaptureDevice) ContinuityCamera() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("continuityCamera"))
	return rv
}

// A Boolean value that indicates whether the flash is currently available for use.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFlashAvailable
func (c_ CaptureDevice) FlashAvailable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("flashAvailable"))
	return rv
}

// Whether the device is following an external sync device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFollowingExternalSyncDevice
func (c_ CaptureDevice) FollowingExternalSyncDevice() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("followingExternalSyncDevice"))
	return rv
}

// A Boolean value that indicates whether the Portrait video effect is active on a device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isPortraitEffectActive
func (c_ CaptureDevice) PortraitEffectActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("portraitEffectActive"))
	return rv
}

// A Boolean value that indicates whether the device supports smooth autofocus.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isSmoothAutoFocusSupported
func (c_ CaptureDevice) SmoothAutoFocusSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("smoothAutoFocusSupported"))
	return rv
}

// A Boolean value that indicates whether the device monitors the subject area for changes.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isSubjectAreaChangeMonitoringEnabled
func (c_ CaptureDevice) SubjectAreaChangeMonitoringEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("subjectAreaChangeMonitoringEnabled"))
	return rv
}


// SetSubjectAreaChangeMonitoringEnabled sets the value of the subjectAreaChangeMonitoringEnabled property.
// A Boolean value that indicates whether the device monitors the subject area for changes.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isSubjectAreaChangeMonitoringEnabled
func (c_ CaptureDevice) SetSubjectAreaChangeMonitoringEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubjectAreaChangeMonitoringEnabled:"), value)
}

// A Boolean value that indicates whether the device is in a suspended state.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isSuspended
func (c_ CaptureDevice) Suspended() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("suspended"))
	return rv
}

// Whether the device’s video frame rate (expressed as a duration) is currently locked.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isVideoFrameDurationLocked
func (c_ CaptureDevice) VideoFrameDurationLocked() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("videoFrameDurationLocked"))
	return rv
}

// A Boolean value that indicates whether the device streams high dynamic range video buffers, also known as extended dynamic range (EDR).
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isVideoHDREnabled
func (c_ CaptureDevice) VideoHDREnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("videoHDREnabled"))
	return rv
}


// SetVideoHDREnabled sets the value of the videoHDREnabled property.
// A Boolean value that indicates whether the device streams high dynamic range video buffers, also known as extended dynamic range (EDR).

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isVideoHDREnabled
func (c_ CaptureDevice) SetVideoHDREnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoHDREnabled:"), value)
}

// A Boolean value that indicates whether the device consists of two or more physical devices.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isVirtualDevice
func (c_ CaptureDevice) VirtualDevice() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("virtualDevice"))
	return rv
}

// The current focus position of the lens.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/lensPosition
func (c_ CaptureDevice) LensPosition() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("lensPosition"))
	return rv
}

// An array of capture devices that are physically linked to a device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/linkedDevices
func (c_ CaptureDevice) LinkedDevices() []CaptureDevice {
	rv := objc.Send[[]CaptureDevice](c_.ID, objc.Sel("linkedDevices"))
	return rv
}

// The maximum zoom factor allowed in the current capture configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/maxAvailableVideoZoomFactor
func (c_ CaptureDevice) MaxAvailableVideoZoomFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("maxAvailableVideoZoomFactor"))
	return rv
}

// The minimum zoom factor allowed in the current capture configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minAvailableVideoZoomFactor
func (c_ CaptureDevice) MinAvailableVideoZoomFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("minAvailableVideoZoomFactor"))
	return rv
}

// The minimum frame duration that can be passed as the when directing your device input to follow an external sync device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minSupportedExternalSyncFrameDuration
func (c_ CaptureDevice) MinSupportedExternalSyncFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("minSupportedExternalSyncFrameDuration"))
	return rv
}

// The capture device’s minimum focus distance in millimeters.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minimumFocusDistance
func (c_ CaptureDevice) MinimumFocusDistance() int {
	rv := objc.Send[int](c_.ID, objc.Sel("minimumFocusDistance"))
	return rv
}

// The nominal 35mm equivalent focal length of the capture device’s lens.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/nominalFocalLengthIn35mmFilm
func (c_ CaptureDevice) NominalFocalLengthIn35mmFilm() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("nominalFocalLengthIn35mmFilm"))
	return rv
}

// The conditions that restrict the primary constituent device’s switching behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/primaryConstituentDeviceRestrictedSwitchingBehaviorConditions-swift.property
func (c_ CaptureDevice) PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions {
	rv := objc.Send[CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions](c_.ID, objc.Sel("primaryConstituentDeviceRestrictedSwitchingBehaviorConditions"))
	return rv
}

// The switching behavior for the primary constituent device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/primaryConstituentDeviceSwitchingBehavior-swift.property
func (c_ CaptureDevice) PrimaryConstituentDeviceSwitchingBehavior() CapturePrimaryConstituentDeviceSwitchingBehavior {
	rv := objc.Send[CapturePrimaryConstituentDeviceSwitchingBehavior](c_.ID, objc.Sel("primaryConstituentDeviceSwitchingBehavior"))
	return rv
}

// A monitor owned by the device that recommends an optimal framing based on the content in the scene.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/smartFramingMonitor
func (c_ CaptureDevice) SmartFramingMonitor() AVCaptureSmartFramingMonitor {
	rv := objc.Send[AVCaptureSmartFramingMonitor](c_.ID, objc.Sel("smartFramingMonitor"))
	return rv
}

// The constituent devices available to select as a fallback for a longer focal length primary constituent device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/supportedFallbackPrimaryConstituentDevices
func (c_ CaptureDevice) SupportedFallbackPrimaryConstituentDevices() []CaptureDevice {
	rv := objc.Send[[]CaptureDevice](c_.ID, objc.Sel("supportedFallbackPrimaryConstituentDevices"))
	return rv
}

// A camera the system prefers to use for video and photo capture.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/systemPreferredCamera
func (c_ CaptureDevice) SystemPreferredCamera() AVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](c_.ID, objc.Sel("systemPreferredCamera"))
	return rv
}

// A value that indicates the capture device’s current system pressure state.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/systemPressureState-swift.property
func (c_ CaptureDevice) SystemPressureState() AVCaptureSystemPressureState {
	rv := objc.Send[AVCaptureSystemPressureState](c_.ID, objc.Sel("systemPressureState"))
	return rv
}

// The transport type of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/transportType
func (c_ CaptureDevice) TransportType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("transportType"))
	return rv
}

// An identifier that uniquely identifies the device.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/uniqueID
func (c_ CaptureDevice) UniqueID() string {
	rv := objc.Send[string](c_.ID, objc.Sel("uniqueID"))
	return rv
}

// A camera the user prefers to use for video and photo capture.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/userPreferredCamera
func (c_ CaptureDevice) UserPreferredCamera() AVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](c_.ID, objc.Sel("userPreferredCamera"))
	return rv
}


// SetUserPreferredCamera sets the value of the userPreferredCamera property.
// A camera the user prefers to use for video and photo capture.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/userPreferredCamera
func (c_ CaptureDevice) SetUserPreferredCamera(value IAVCaptureDevice) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserPreferredCamera:"), value)
}

// A virtual device’s active primary constituent device.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeprimaryconstituent
func (c_ CaptureDevice) ActivePrimaryConstituent() AVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](c_.ID, objc.Sel("activePrimaryConstituent"))
	return rv
}


// SetActivePrimaryConstituent sets the value of the activePrimaryConstituent property.
// A virtual device’s active primary constituent device.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeprimaryconstituent
func (c_ CaptureDevice) SetActivePrimaryConstituent(value IAVCaptureDevice) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActivePrimaryConstituent:"), value)
}

// The current scene monitoring statuses related to Cinematic Video capture.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/cinematicvideocapturescenemonitoringstatuses
func (c_ CaptureDevice) CinematicVideoCaptureSceneMonitoringStatuses() CaptureSceneMonitoringStatus {
	rv := objc.Send[CaptureSceneMonitoringStatus](c_.ID, objc.Sel("cinematicVideoCaptureSceneMonitoringStatuses"))
	return rv
}


// SetCinematicVideoCaptureSceneMonitoringStatuses sets the value of the cinematicVideoCaptureSceneMonitoringStatuses property.
// The current scene monitoring statuses related to Cinematic Video capture.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/cinematicvideocapturescenemonitoringstatuses
func (c_ CaptureDevice) SetCinematicVideoCaptureSceneMonitoringStatuses(value CaptureSceneMonitoringStatus) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCinematicVideoCaptureSceneMonitoringStatuses:"), value)
}

// The type of device, such as a built-in microphone or wide-angle camera.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/devicetype-swift.property
func (c_ CaptureDevice) DeviceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("deviceType"))
	return rv
}


// SetDeviceType sets the value of the deviceType property.
// The type of device, such as a built-in microphone or wide-angle camera.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/devicetype-swift.property
func (c_ CaptureDevice) SetDeviceType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeviceType:"), value)
}

// A Boolean value that indicates whether the capture device performs automatic video frame rate adjustments.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isautovideoframerateenabled
func (c_ CaptureDevice) IsAutoVideoFrameRateEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoVideoFrameRateEnabled"))
	return rv
}


// SetIsAutoVideoFrameRateEnabled sets the value of the isAutoVideoFrameRateEnabled property.
// A Boolean value that indicates whether the capture device performs automatic video frame rate adjustments.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isautovideoframerateenabled
func (c_ CaptureDevice) SetIsAutoVideoFrameRateEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoVideoFrameRateEnabled:"), value)
}

// Whether camera lens smudge detection is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/iscameralenssmudgedetectionenabled
func (c_ CaptureDevice) IsCameraLensSmudgeDetectionEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraLensSmudgeDetectionEnabled"))
	return rv
}


// SetIsCameraLensSmudgeDetectionEnabled sets the value of the isCameraLensSmudgeDetectionEnabled property.
// Whether camera lens smudge detection is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/iscameralenssmudgedetectionenabled
func (c_ CaptureDevice) SetIsCameraLensSmudgeDetectionEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraLensSmudgeDetectionEnabled:"), value)
}

// A Boolean value that indicates whether a device is currently connected to the system and available for use.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isconnected
func (c_ CaptureDevice) IsConnected() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConnected"))
	return rv
}


// SetIsConnected sets the value of the isConnected property.
// A Boolean value that indicates whether a device is currently connected to the system and available for use.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isconnected
func (c_ CaptureDevice) SetIsConnected(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConnected:"), value)
}

// A Boolean value that indicates whether the device is a Continuity Camera.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/iscontinuitycamera
func (c_ CaptureDevice) IsContinuityCamera() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isContinuityCamera"))
	return rv
}


// SetIsContinuityCamera sets the value of the isContinuityCamera property.
// A Boolean value that indicates whether the device is a Continuity Camera.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/iscontinuitycamera
func (c_ CaptureDevice) SetIsContinuityCamera(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsContinuityCamera:"), value)
}

// Whether the device is following an external sync device.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isfollowingexternalsyncdevice
func (c_ CaptureDevice) IsFollowingExternalSyncDevice() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFollowingExternalSyncDevice"))
	return rv
}


// SetIsFollowingExternalSyncDevice sets the value of the isFollowingExternalSyncDevice property.
// Whether the device is following an external sync device.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isfollowingexternalsyncdevice
func (c_ CaptureDevice) SetIsFollowingExternalSyncDevice(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFollowingExternalSyncDevice:"), value)
}

// A Boolean value that indicates whether another app is using the device.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isinusebyanotherapplication
func (c_ CaptureDevice) IsInUseByAnotherApplication() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isInUseByAnotherApplication"))
	return rv
}


// SetIsInUseByAnotherApplication sets the value of the isInUseByAnotherApplication property.
// A Boolean value that indicates whether another app is using the device.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isinusebyanotherapplication
func (c_ CaptureDevice) SetIsInUseByAnotherApplication(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsInUseByAnotherApplication:"), value)
}

// A Boolean value that indicates whether the device monitors the subject area for changes.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/issubjectareachangemonitoringenabled
func (c_ CaptureDevice) IsSubjectAreaChangeMonitoringEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSubjectAreaChangeMonitoringEnabled"))
	return rv
}


// SetIsSubjectAreaChangeMonitoringEnabled sets the value of the isSubjectAreaChangeMonitoringEnabled property.
// A Boolean value that indicates whether the device monitors the subject area for changes.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/issubjectareachangemonitoringenabled
func (c_ CaptureDevice) SetIsSubjectAreaChangeMonitoringEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSubjectAreaChangeMonitoringEnabled:"), value)
}

// A Boolean value that indicates whether the device is in a suspended state.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/issuspended
func (c_ CaptureDevice) IsSuspended() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSuspended"))
	return rv
}


// SetIsSuspended sets the value of the isSuspended property.
// A Boolean value that indicates whether the device is in a suspended state.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/issuspended
func (c_ CaptureDevice) SetIsSuspended(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSuspended:"), value)
}

// Whether the device’s video frame rate (expressed as a duration) is currently locked.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isvideoframedurationlocked
func (c_ CaptureDevice) IsVideoFrameDurationLocked() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoFrameDurationLocked"))
	return rv
}


// SetIsVideoFrameDurationLocked sets the value of the isVideoFrameDurationLocked property.
// Whether the device’s video frame rate (expressed as a duration) is currently locked.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isvideoframedurationlocked
func (c_ CaptureDevice) SetIsVideoFrameDurationLocked(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoFrameDurationLocked:"), value)
}

// A Boolean value that indicates whether the device consists of two or more physical devices.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isvirtualdevice
func (c_ CaptureDevice) IsVirtualDevice() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVirtualDevice"))
	return rv
}


// SetIsVirtualDevice sets the value of the isVirtualDevice property.
// A Boolean value that indicates whether the device consists of two or more physical devices.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isvirtualdevice
func (c_ CaptureDevice) SetIsVirtualDevice(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVirtualDevice:"), value)
}

// A localized device name for display in the user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/localizedname
func (c_ CaptureDevice) LocalizedName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("localizedName"))
	return rv
}


// SetLocalizedName sets the value of the localizedName property.
// A localized device name for display in the user interface.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/localizedname
func (c_ CaptureDevice) SetLocalizedName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocalizedName:"), objc.String(value))
}

// A human-readable string for the manufacturer of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/manufacturer
func (c_ CaptureDevice) Manufacturer() string {
	rv := objc.Send[string](c_.ID, objc.Sel("manufacturer"))
	return rv
}


// SetManufacturer sets the value of the manufacturer property.
// A human-readable string for the manufacturer of the device.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/manufacturer
func (c_ CaptureDevice) SetManufacturer(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setManufacturer:"), objc.String(value))
}

// The maximum frame rate (expressed as a minimum duration) that can be set on an input associated with this device.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/minsupportedlockedvideoframeduration
func (c_ CaptureDevice) MinSupportedLockedVideoFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("minSupportedLockedVideoFrameDuration"))
	return rv
}


// SetMinSupportedLockedVideoFrameDuration sets the value of the minSupportedLockedVideoFrameDuration property.
// The maximum frame rate (expressed as a minimum duration) that can be set on an input associated with this device.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/minsupportedlockedvideoframeduration
func (c_ CaptureDevice) SetMinSupportedLockedVideoFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinSupportedLockedVideoFrameDuration:"), value)
}

// A model identifier for the device.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/modelid
func (c_ CaptureDevice) ModelID() string {
	rv := objc.Send[string](c_.ID, objc.Sel("modelID"))
	return rv
}


// SetModelID sets the value of the modelID property.
// A model identifier for the device.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/modelid
func (c_ CaptureDevice) SetModelID(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModelID:"), objc.String(value))
}

// The physical position of the capture device hardware.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/position-swift.property
func (c_ CaptureDevice) Position() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("position"))
	return rv
}


// SetPosition sets the value of the position property.
// The physical position of the capture device hardware.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/position-swift.property
func (c_ CaptureDevice) SetPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPosition:"), value)
}

// Reasons why current environmental conditions aren’t suitable to capturing spatial videos that are comfortable to view.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/spatialcapturediscomfortreasons
func (c_ CaptureDevice) SpatialCaptureDiscomfortReasons() SpatialCaptureDiscomfortReason {
	rv := objc.Send[SpatialCaptureDiscomfortReason](c_.ID, objc.Sel("spatialCaptureDiscomfortReasons"))
	return rv
}


// SetSpatialCaptureDiscomfortReasons sets the value of the spatialCaptureDiscomfortReasons property.
// Reasons why current environmental conditions aren’t suitable to capturing spatial videos that are comfortable to view.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/spatialcapturediscomfortreasons
func (c_ CaptureDevice) SetSpatialCaptureDiscomfortReasons(value ISpatialCaptureDiscomfortReason) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSpatialCaptureDiscomfortReasons:"), value)
}

// A key to retrieve a state value that indicates the system pressure level and contributing factors that caused the interruption.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesessioninterruptionsystempressurestatekey
func (c_ CaptureDevice) AVCaptureSessionInterruptionSystemPressureStateKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("AVCaptureSessionInterruptionSystemPressureStateKey"))
	return rv
}


