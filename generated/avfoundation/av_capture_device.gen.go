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
	// properties:
	ActiveFormat() IAVCaptureDeviceFormat
	SetActiveFormat(value IAVCaptureDeviceFormat)
	ActivePrimaryConstituentDevice() IAVCaptureDevice
	ActivePrimaryConstituentDeviceSwitchingBehavior() AVCapturePrimaryConstituentDeviceSwitchingBehavior /* foo */
	CameraLensSmudgeDetectionInterval() CMTime /* foo */
	CinematicVideoCaptureSceneMonitoringStatuses() unsafe.Pointer
	ConstituentDevices() []CaptureDevice /* primitive/slice/pointer */
	DynamicAspectRatio() AVCaptureAspectRatio /* foo */
	AutoVideoFrameRateEnabled() bool /* primitive/slice/pointer */
	SetAutoVideoFrameRateEnabled(value bool /* primitive/slice/pointer */)
	CameraLensSmudgeDetectionEnabled() bool /* primitive/slice/pointer */
	Suspended() bool /* primitive/slice/pointer */
	VideoFrameDurationLocked() bool /* primitive/slice/pointer */
	MinSupportedLockedVideoFrameDuration() CMTime /* foo */
	NominalFocalLengthIn35mmFilm() float32 /* primitive/slice/pointer */
	Position() AVCaptureDevicePosition /* enum */
	SystemPressureState() AVCaptureSystemPressureState /* foo */
	TransportType() unsafe.Pointer
	ActivePrimaryConstituent() IAVCaptureDevice
	SetActivePrimaryConstituent(value IAVCaptureDevice)
	ActivePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() unsafe.Pointer
	SetActivePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions(value unsafe.Pointer)
	CameraLensSmudgeDetectionStatus() AVCaptureCameraLensSmudgeDetectionStatus /* foo */
	SetCameraLensSmudgeDetectionStatus(value AVCaptureCameraLensSmudgeDetectionStatus /* foo */)
	CompanionDeskViewCamera() IAVCaptureDevice
	SetCompanionDeskViewCamera(value IAVCaptureDevice)
	DeviceType() unsafe.Pointer
	SetDeviceType(value unsafe.Pointer)
	DynamicDimensions() CMVideoDimensions /* foo */
	SetDynamicDimensions(value CMVideoDimensions /* foo */)
	FallbackPrimaryConstituentDevices() IAVCaptureDevice
	SetFallbackPrimaryConstituentDevices(value IAVCaptureDevice)
	IsAutoVideoFrameRateEnabled() bool /* primitive/slice/pointer */
	SetIsAutoVideoFrameRateEnabled(value bool /* primitive/slice/pointer */)
	IsCameraLensSmudgeDetectionEnabled() bool /* primitive/slice/pointer */
	SetIsCameraLensSmudgeDetectionEnabled(value bool /* primitive/slice/pointer */)
	IsConnected() bool /* primitive/slice/pointer */
	SetIsConnected(value bool /* primitive/slice/pointer */)
	IsContinuityCamera() bool /* primitive/slice/pointer */
	SetIsContinuityCamera(value bool /* primitive/slice/pointer */)
	IsFollowingExternalSyncDevice() bool /* primitive/slice/pointer */
	SetIsFollowingExternalSyncDevice(value bool /* primitive/slice/pointer */)
	IsInUseByAnotherApplication() bool /* primitive/slice/pointer */
	SetIsInUseByAnotherApplication(value bool /* primitive/slice/pointer */)
	IsSubjectAreaChangeMonitoringEnabled() bool /* primitive/slice/pointer */
	SetIsSubjectAreaChangeMonitoringEnabled(value bool /* primitive/slice/pointer */)
	IsSuspended() bool /* primitive/slice/pointer */
	SetIsSuspended(value bool /* primitive/slice/pointer */)
	IsVideoFrameDurationLocked() bool /* primitive/slice/pointer */
	SetIsVideoFrameDurationLocked(value bool /* primitive/slice/pointer */)
	IsVirtualDevice() bool /* primitive/slice/pointer */
	SetIsVirtualDevice(value bool /* primitive/slice/pointer */)
	LocalizedName() string /* primitive/slice/pointer */
	SetLocalizedName(value string /* primitive/slice/pointer */)
	Manufacturer() string /* primitive/slice/pointer */
	SetManufacturer(value string /* primitive/slice/pointer */)
	MinSupportedExternalSyncFrameDuration() CMTime /* foo */
	SetMinSupportedExternalSyncFrameDuration(value CMTime /* foo */)
	ModelID() string /* primitive/slice/pointer */
	SetModelID(value string /* primitive/slice/pointer */)
	PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() unsafe.Pointer
	SetPrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions(value unsafe.Pointer)
	PrimaryConstituentDeviceSwitchingBehavior() unsafe.Pointer
	SetPrimaryConstituentDeviceSwitchingBehavior(value unsafe.Pointer)
	SmartFramingMonitor() IAVCaptureSmartFramingMonitor
	SetSmartFramingMonitor(value IAVCaptureSmartFramingMonitor)
	SpatialCaptureDiscomfortReasons() AVSpatialCaptureDiscomfortReason /* foo */
	SetSpatialCaptureDiscomfortReasons(value AVSpatialCaptureDiscomfortReason /* foo */)
	SupportedFallbackPrimaryConstituentDevices() IAVCaptureDevice
	SetSupportedFallbackPrimaryConstituentDevices(value IAVCaptureDevice)
	UniqueID() string /* primitive/slice/pointer */
	SetUniqueID(value string /* primitive/slice/pointer */)
	AVCaptureSessionInterruptionSystemPressureStateKey() string /* primitive/slice/pointer */
	// methods:
	HasMediaType(mediaType AVMediaType /* foo */) bool /* primitive/slice/pointer */
	LockForConfiguration(outError unsafe.Pointer) bool /* primitive/slice/pointer */
	SetCinematicVideoFixedFocusAtPointFocusMode(point coregraphics.CGPoint, focusMode AVCaptureCinematicVideoFocusMode /* foo */)
	SetCinematicVideoTrackingFocusAtPointFocusMode(point coregraphics.CGPoint, focusMode AVCaptureCinematicVideoFocusMode /* foo */)
	SetCinematicVideoTrackingFocusWithDetectedObjectIDFocusMode(detectedObjectID int /* primitive/slice/pointer */, focusMode AVCaptureCinematicVideoFocusMode /* foo */)
	UnlockForConfiguration()
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/init(uniqueID:)
func NewCaptureDeviceWithUniqueID(deviceUniqueID string /* primitive/slice/pointer */) CaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(getCaptureDeviceClass().class), objc.Sel("deviceWithUniqueID:"), objc.String(deviceUniqueID))
	return rv
}



// Returns an authorization status that indicates whether the user grants the app permission to capture media of a particular type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/authorizationStatus(for:)
func (cc _CaptureDeviceClass) AuthorizationStatusForMediaType(mediaType AVMediaType /* foo */) AVAuthorizationStatus /* enum */ {
	rv := objc.Send[AVAuthorizationStatus](objc.ID(cc.class), objc.Sel("authorizationStatusForMediaType:"), mediaType)
	return rv
}


// Returns the default device for the specified device type, media type, and position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/default(_:for:position:)
func (cc _CaptureDeviceClass) DefaultDeviceWithDeviceTypeMediaTypePosition(deviceType AVCaptureDeviceType /* typedef */, mediaType AVMediaType /* foo */, position AVCaptureDevicePosition /* enum */) CaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(cc.class), objc.Sel("defaultDeviceWithDeviceType:mediaType:position:"), deviceType, mediaType, position)
	return rv
}


// Returns the default device that captures the specified media type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/default(for:)
func (cc _CaptureDeviceClass) DefaultDeviceWithMediaType(mediaType AVMediaType /* foo */) CaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(cc.class), objc.Sel("defaultDeviceWithMediaType:"), mediaType)
	return rv
}


// Returns all available capture devices on the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/devices()
func (cc _CaptureDeviceClass) Devices() []CaptureDevice /* primitive/slice/pointer */ {
	rv := objc.Send[[]CaptureDevice](objc.ID(cc.class), objc.Sel("devices"))
	return rv
}


// Returns devices capable of capturing media of the specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/devices(for:)
func (cc _CaptureDeviceClass) DevicesWithMediaType(mediaType AVMediaType /* foo */) []CaptureDevice /* primitive/slice/pointer */ {
	rv := objc.Send[[]CaptureDevice](objc.ID(cc.class), objc.Sel("devicesWithMediaType:"), mediaType)
	return rv
}


// Returns the relative extrinsic matrix from one capture device to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/extrinsicMatrix(from:to:)
func (cc _CaptureDeviceClass) ExtrinsicMatrixFromDeviceToDevice(fromDevice IAVCaptureDevice, toDevice IAVCaptureDevice) NSData /* foo */ {
	rv := objc.Send[Data](objc.ID(cc.class), objc.Sel("extrinsicMatrixFromDevice:toDevice:"), fromDevice, toDevice)
	return rv
}


// Creates an object that represents a device with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/init(uniqueID:)
func (cc _CaptureDeviceClass) DeviceWithUniqueID(deviceUniqueID string /* primitive/slice/pointer */) CaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(cc.class), objc.Sel("deviceWithUniqueID:"), objc.String(deviceUniqueID))
	return rv
}


// A class property that indicates whether a person enables the Background Replacement feature for this app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isBackgroundReplacementEnabled
func (cc _CaptureDeviceClass) BackgroundReplacementEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("backgroundReplacementEnabled"))
	return rv
}

// Returns a Boolean value that indicates whether the device captures media of a particular type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/hasMediaType(_:)
func (c_ CaptureDevice) HasMediaType(mediaType AVMediaType /* foo */) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasMediaType:"), mediaType)
	return rv
}


// Requests exclusive access to configure device hardware properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/lockForConfiguration()
func (c_ CaptureDevice) LockForConfiguration(outError unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("lockForConfiguration:"), outError)
	return rv
}


// Fix focus at a distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setCinematicVideoFixedFocus(at:focusMode:)
func (c_ CaptureDevice) SetCinematicVideoFixedFocusAtPointFocusMode(point coregraphics.CGPoint, focusMode AVCaptureCinematicVideoFocusMode /* foo */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCinematicVideoFixedFocusAtPoint:focusMode:"), point, focusMode)
}


// Focus on and start tracking an object if it can be detected at the region specified by the point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setCinematicVideoTrackingFocus(at:focusMode:)
func (c_ CaptureDevice) SetCinematicVideoTrackingFocusAtPointFocusMode(point coregraphics.CGPoint, focusMode AVCaptureCinematicVideoFocusMode /* foo */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCinematicVideoTrackingFocusAtPoint:focusMode:"), point, focusMode)
}


// Focus on and start tracking a detected object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setCinematicVideoTrackingFocus(detectedObjectID:focusMode:)
func (c_ CaptureDevice) SetCinematicVideoTrackingFocusWithDetectedObjectIDFocusMode(detectedObjectID int /* primitive/slice/pointer */, focusMode AVCaptureCinematicVideoFocusMode /* foo */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCinematicVideoTrackingFocusWithDetectedObjectID:focusMode:"), detectedObjectID, focusMode)
}


// Releases exclusive control over device hardware properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/unlockForConfiguration()
func (c_ CaptureDevice) UnlockForConfiguration() {
	objc.Send[objc.ID](c_.ID, objc.Sel("unlockForConfiguration"))
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


// A virtual device’s active primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activePrimaryConstituent
func (c_ CaptureDevice) ActivePrimaryConstituentDevice() IAVCaptureDevice {
	rv := objc.Send[CaptureDevice](c_.ID, objc.Sel("activePrimaryConstituentDevice"))
	return rv
}


// The switching behavior of the active constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activePrimaryConstituentDeviceSwitchingBehavior
func (c_ CaptureDevice) ActivePrimaryConstituentDeviceSwitchingBehavior() AVCapturePrimaryConstituentDeviceSwitchingBehavior /* foo */ {
	rv := objc.Send[CapturePrimaryConstituentDeviceSwitchingBehavior](c_.ID, objc.Sel("activePrimaryConstituentDeviceSwitchingBehavior"))
	return rv
}


// The camera lens smudge detection interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/cameraLensSmudgeDetectionInterval
func (c_ CaptureDevice) CameraLensSmudgeDetectionInterval() CMTime /* foo */ {
	rv := objc.Send[Time](c_.ID, objc.Sel("cameraLensSmudgeDetectionInterval"))
	return rv
}


// The current scene monitoring statuses related to Cinematic Video capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/cinematicVideoCaptureSceneMonitoringStatuses
func (c_ CaptureDevice) CinematicVideoCaptureSceneMonitoringStatuses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("cinematicVideoCaptureSceneMonitoringStatuses"))
	return rv
}


// An array of physical devices that make up a virtual device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/constituentDevices
func (c_ CaptureDevice) ConstituentDevices() []CaptureDevice /* primitive/slice/pointer */ {
	rv := objc.Send[[]CaptureDevice](c_.ID, objc.Sel("constituentDevices"))
	return rv
}


// A key-value observable property indicating the current aspect ratio for a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/dynamicAspectRatio
func (c_ CaptureDevice) DynamicAspectRatio() AVCaptureAspectRatio /* foo */ {
	rv := objc.Send[CaptureAspectRatio](c_.ID, objc.Sel("dynamicAspectRatio"))
	return rv
}


// A Boolean value that indicates whether the capture device performs automatic video frame rate adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isAutoVideoFrameRateEnabled
func (c_ CaptureDevice) AutoVideoFrameRateEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("autoVideoFrameRateEnabled"))
	return rv
}


// A Boolean value that indicates whether the capture device performs automatic video frame rate adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isAutoVideoFrameRateEnabled
func (c_ CaptureDevice) SetAutoVideoFrameRateEnabled(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutoVideoFrameRateEnabled:"), value)
}


// A class property that indicates whether a person enables the Background Replacement feature for this app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isBackgroundReplacementEnabled
func (c_ CaptureDevice) BackgroundReplacementEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("backgroundReplacementEnabled"))
	return rv
}


// Whether camera lens smudge detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isCameraLensSmudgeDetectionEnabled
func (c_ CaptureDevice) CameraLensSmudgeDetectionEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("cameraLensSmudgeDetectionEnabled"))
	return rv
}


// A Boolean value that indicates whether the device is in a suspended state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isSuspended
func (c_ CaptureDevice) Suspended() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("suspended"))
	return rv
}


// Whether the device’s video frame rate (expressed as a duration) is currently locked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isVideoFrameDurationLocked
func (c_ CaptureDevice) VideoFrameDurationLocked() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("videoFrameDurationLocked"))
	return rv
}


// The maximum frame rate (expressed as a minimum duration) that can be set on an input associated with this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minSupportedLockedVideoFrameDuration
func (c_ CaptureDevice) MinSupportedLockedVideoFrameDuration() CMTime /* foo */ {
	rv := objc.Send[Time](c_.ID, objc.Sel("minSupportedLockedVideoFrameDuration"))
	return rv
}


// The nominal 35mm equivalent focal length of the capture device’s lens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/nominalFocalLengthIn35mmFilm
func (c_ CaptureDevice) NominalFocalLengthIn35mmFilm() float32 /* primitive/slice/pointer */ {
	rv := objc.Send[float32](c_.ID, objc.Sel("nominalFocalLengthIn35mmFilm"))
	return rv
}


// The physical position of the capture device hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/position-swift.property
func (c_ CaptureDevice) Position() AVCaptureDevicePosition /* enum */ {
	rv := objc.Send[AVCaptureDevicePosition](c_.ID, objc.Sel("position"))
	return rv
}


// A value that indicates the capture device’s current system pressure state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/systemPressureState-swift.property
func (c_ CaptureDevice) SystemPressureState() AVCaptureSystemPressureState /* foo */ {
	rv := objc.Send[CaptureSystemPressureState](c_.ID, objc.Sel("systemPressureState"))
	return rv
}


// The transport type of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/transportType
func (c_ CaptureDevice) TransportType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("transportType"))
	return rv
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


// The conditions that restrict camera switching behavior for the active primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeprimaryconstituentdevicerestrictedswitchingbehaviorconditions
func (c_ CaptureDevice) ActivePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("activePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions"))
	return rv
}


// The conditions that restrict camera switching behavior for the active primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeprimaryconstituentdevicerestrictedswitchingbehaviorconditions
func (c_ CaptureDevice) SetActivePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActivePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions:"), value)
}


// A value specifying the status of camera lens smudge detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/cameralenssmudgedetectionstatus
func (c_ CaptureDevice) CameraLensSmudgeDetectionStatus() AVCaptureCameraLensSmudgeDetectionStatus /* foo */ {
	rv := objc.Send[CaptureCameraLensSmudgeDetectionStatus](c_.ID, objc.Sel("cameraLensSmudgeDetectionStatus"))
	return rv
}


// A value specifying the status of camera lens smudge detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/cameralenssmudgedetectionstatus
func (c_ CaptureDevice) SetCameraLensSmudgeDetectionStatus(value AVCaptureCameraLensSmudgeDetectionStatus /* foo */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCameraLensSmudgeDetectionStatus:"), value)
}


// A Desk View camera associated with a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/companiondeskviewcamera
func (c_ CaptureDevice) CompanionDeskViewCamera() IAVCaptureDevice {
	rv := objc.Send[CaptureDevice](c_.ID, objc.Sel("companionDeskViewCamera"))
	return rv
}


// A Desk View camera associated with a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/companiondeskviewcamera
func (c_ CaptureDevice) SetCompanionDeskViewCamera(value IAVCaptureDevice) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompanionDeskViewCamera:"), value)
}


// The type of device, such as a built-in microphone or wide-angle camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/devicetype-swift.property
func (c_ CaptureDevice) DeviceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("deviceType"))
	return rv
}


// The type of device, such as a built-in microphone or wide-angle camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/devicetype-swift.property
func (c_ CaptureDevice) SetDeviceType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeviceType:"), value)
}


// A key-value observable property describing the output dimensions of the video buffer based on the device’s dynamic aspect ratio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/dynamicdimensions
func (c_ CaptureDevice) DynamicDimensions() CMVideoDimensions /* foo */ {
	rv := objc.Send[VideoDimensions](c_.ID, objc.Sel("dynamicDimensions"))
	return rv
}


// A key-value observable property describing the output dimensions of the video buffer based on the device’s dynamic aspect ratio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/dynamicdimensions
func (c_ CaptureDevice) SetDynamicDimensions(value CMVideoDimensions /* foo */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDynamicDimensions:"), value)
}


// The fallback devices to use when a constituent device with a longer focal length becomes limited by its light sensitivity or minimum focus distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/fallbackprimaryconstituentdevices
func (c_ CaptureDevice) FallbackPrimaryConstituentDevices() IAVCaptureDevice {
	rv := objc.Send[CaptureDevice](c_.ID, objc.Sel("fallbackPrimaryConstituentDevices"))
	return rv
}


// The fallback devices to use when a constituent device with a longer focal length becomes limited by its light sensitivity or minimum focus distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/fallbackprimaryconstituentdevices
func (c_ CaptureDevice) SetFallbackPrimaryConstituentDevices(value IAVCaptureDevice) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFallbackPrimaryConstituentDevices:"), value)
}


// A Boolean value that indicates whether the capture device performs automatic video frame rate adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isautovideoframerateenabled
func (c_ CaptureDevice) IsAutoVideoFrameRateEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAutoVideoFrameRateEnabled"))
	return rv
}


// A Boolean value that indicates whether the capture device performs automatic video frame rate adjustments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isautovideoframerateenabled
func (c_ CaptureDevice) SetIsAutoVideoFrameRateEnabled(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAutoVideoFrameRateEnabled:"), value)
}


// Whether camera lens smudge detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/iscameralenssmudgedetectionenabled
func (c_ CaptureDevice) IsCameraLensSmudgeDetectionEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCameraLensSmudgeDetectionEnabled"))
	return rv
}


// Whether camera lens smudge detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/iscameralenssmudgedetectionenabled
func (c_ CaptureDevice) SetIsCameraLensSmudgeDetectionEnabled(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCameraLensSmudgeDetectionEnabled:"), value)
}


// A Boolean value that indicates whether a device is currently connected to the system and available for use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isconnected
func (c_ CaptureDevice) IsConnected() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConnected"))
	return rv
}


// A Boolean value that indicates whether a device is currently connected to the system and available for use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isconnected
func (c_ CaptureDevice) SetIsConnected(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConnected:"), value)
}


// A Boolean value that indicates whether the device is a Continuity Camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/iscontinuitycamera
func (c_ CaptureDevice) IsContinuityCamera() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isContinuityCamera"))
	return rv
}


// A Boolean value that indicates whether the device is a Continuity Camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/iscontinuitycamera
func (c_ CaptureDevice) SetIsContinuityCamera(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsContinuityCamera:"), value)
}


// Whether the device is following an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isfollowingexternalsyncdevice
func (c_ CaptureDevice) IsFollowingExternalSyncDevice() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFollowingExternalSyncDevice"))
	return rv
}


// Whether the device is following an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isfollowingexternalsyncdevice
func (c_ CaptureDevice) SetIsFollowingExternalSyncDevice(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFollowingExternalSyncDevice:"), value)
}


// A Boolean value that indicates whether another app is using the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isinusebyanotherapplication
func (c_ CaptureDevice) IsInUseByAnotherApplication() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isInUseByAnotherApplication"))
	return rv
}


// A Boolean value that indicates whether another app is using the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isinusebyanotherapplication
func (c_ CaptureDevice) SetIsInUseByAnotherApplication(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsInUseByAnotherApplication:"), value)
}


// A Boolean value that indicates whether the device monitors the subject area for changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/issubjectareachangemonitoringenabled
func (c_ CaptureDevice) IsSubjectAreaChangeMonitoringEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSubjectAreaChangeMonitoringEnabled"))
	return rv
}


// A Boolean value that indicates whether the device monitors the subject area for changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/issubjectareachangemonitoringenabled
func (c_ CaptureDevice) SetIsSubjectAreaChangeMonitoringEnabled(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSubjectAreaChangeMonitoringEnabled:"), value)
}


// A Boolean value that indicates whether the device is in a suspended state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/issuspended
func (c_ CaptureDevice) IsSuspended() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSuspended"))
	return rv
}


// A Boolean value that indicates whether the device is in a suspended state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/issuspended
func (c_ CaptureDevice) SetIsSuspended(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSuspended:"), value)
}


// Whether the device’s video frame rate (expressed as a duration) is currently locked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isvideoframedurationlocked
func (c_ CaptureDevice) IsVideoFrameDurationLocked() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideoFrameDurationLocked"))
	return rv
}


// Whether the device’s video frame rate (expressed as a duration) is currently locked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isvideoframedurationlocked
func (c_ CaptureDevice) SetIsVideoFrameDurationLocked(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideoFrameDurationLocked:"), value)
}


// A Boolean value that indicates whether the device consists of two or more physical devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isvirtualdevice
func (c_ CaptureDevice) IsVirtualDevice() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVirtualDevice"))
	return rv
}


// A Boolean value that indicates whether the device consists of two or more physical devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/isvirtualdevice
func (c_ CaptureDevice) SetIsVirtualDevice(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVirtualDevice:"), value)
}


// A localized device name for display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/localizedname
func (c_ CaptureDevice) LocalizedName() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](c_.ID, objc.Sel("localizedName"))
	return rv
}


// A localized device name for display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/localizedname
func (c_ CaptureDevice) SetLocalizedName(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocalizedName:"), objc.String(value))
}


// A human-readable string for the manufacturer of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/manufacturer
func (c_ CaptureDevice) Manufacturer() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](c_.ID, objc.Sel("manufacturer"))
	return rv
}


// A human-readable string for the manufacturer of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/manufacturer
func (c_ CaptureDevice) SetManufacturer(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setManufacturer:"), objc.String(value))
}


// The minimum frame duration that can be passed as the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/minsupportedexternalsyncframeduration
func (c_ CaptureDevice) MinSupportedExternalSyncFrameDuration() CMTime /* foo */ {
	rv := objc.Send[Time](c_.ID, objc.Sel("minSupportedExternalSyncFrameDuration"))
	return rv
}


// The minimum frame duration that can be passed as the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/minsupportedexternalsyncframeduration
func (c_ CaptureDevice) SetMinSupportedExternalSyncFrameDuration(value CMTime /* foo */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinSupportedExternalSyncFrameDuration:"), value)
}


// A model identifier for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/modelid
func (c_ CaptureDevice) ModelID() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](c_.ID, objc.Sel("modelID"))
	return rv
}


// A model identifier for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/modelid
func (c_ CaptureDevice) SetModelID(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModelID:"), objc.String(value))
}


// The conditions that restrict the primary constituent device’s switching behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/primaryconstituentdevicerestrictedswitchingbehaviorconditions-swift.property
func (c_ CaptureDevice) PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("primaryConstituentDeviceRestrictedSwitchingBehaviorConditions"))
	return rv
}


// The conditions that restrict the primary constituent device’s switching behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/primaryconstituentdevicerestrictedswitchingbehaviorconditions-swift.property
func (c_ CaptureDevice) SetPrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions:"), value)
}


// The switching behavior for the primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/primaryconstituentdeviceswitchingbehavior-swift.property
func (c_ CaptureDevice) PrimaryConstituentDeviceSwitchingBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("primaryConstituentDeviceSwitchingBehavior"))
	return rv
}


// The switching behavior for the primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/primaryconstituentdeviceswitchingbehavior-swift.property
func (c_ CaptureDevice) SetPrimaryConstituentDeviceSwitchingBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryConstituentDeviceSwitchingBehavior:"), value)
}


// A monitor owned by the device that recommends an optimal framing based on the content in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/smartframingmonitor
func (c_ CaptureDevice) SmartFramingMonitor() IAVCaptureSmartFramingMonitor {
	rv := objc.Send[CaptureSmartFramingMonitor](c_.ID, objc.Sel("smartFramingMonitor"))
	return rv
}


// A monitor owned by the device that recommends an optimal framing based on the content in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/smartframingmonitor
func (c_ CaptureDevice) SetSmartFramingMonitor(value IAVCaptureSmartFramingMonitor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSmartFramingMonitor:"), value)
}


// Reasons why current environmental conditions aren’t suitable to capturing spatial videos that are comfortable to view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/spatialcapturediscomfortreasons
func (c_ CaptureDevice) SpatialCaptureDiscomfortReasons() AVSpatialCaptureDiscomfortReason /* foo */ {
	rv := objc.Send[SpatialCaptureDiscomfortReason](c_.ID, objc.Sel("spatialCaptureDiscomfortReasons"))
	return rv
}


// Reasons why current environmental conditions aren’t suitable to capturing spatial videos that are comfortable to view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/spatialcapturediscomfortreasons
func (c_ CaptureDevice) SetSpatialCaptureDiscomfortReasons(value AVSpatialCaptureDiscomfortReason /* foo */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSpatialCaptureDiscomfortReasons:"), value)
}


// The constituent devices available to select as a fallback for a longer focal length primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/supportedfallbackprimaryconstituentdevices
func (c_ CaptureDevice) SupportedFallbackPrimaryConstituentDevices() IAVCaptureDevice {
	rv := objc.Send[CaptureDevice](c_.ID, objc.Sel("supportedFallbackPrimaryConstituentDevices"))
	return rv
}


// The constituent devices available to select as a fallback for a longer focal length primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/supportedfallbackprimaryconstituentdevices
func (c_ CaptureDevice) SetSupportedFallbackPrimaryConstituentDevices(value IAVCaptureDevice) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedFallbackPrimaryConstituentDevices:"), value)
}


// An identifier that uniquely identifies the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/uniqueid
func (c_ CaptureDevice) UniqueID() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](c_.ID, objc.Sel("uniqueID"))
	return rv
}


// An identifier that uniquely identifies the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/uniqueid
func (c_ CaptureDevice) SetUniqueID(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUniqueID:"), objc.String(value))
}


// A key to retrieve a state value that indicates the system pressure level and contributing factors that caused the interruption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesessioninterruptionsystempressurestatekey
func (c_ CaptureDevice) AVCaptureSessionInterruptionSystemPressureStateKey() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](c_.ID, objc.Sel("AVCaptureSessionInterruptionSystemPressureStateKey"))
	return rv
}


