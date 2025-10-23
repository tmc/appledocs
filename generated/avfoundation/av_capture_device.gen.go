// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	ActiveFormat() unsafe.Pointer
	SetActiveFormat(value unsafe.Pointer)
	AutoVideoFrameRateEnabled() bool
	SetAutoVideoFrameRateEnabled(value bool)
	SystemPressureState() CaptureSystemPressureState
	ActivePrimaryConstituent() IAVCaptureDevice
	SetActivePrimaryConstituent(value IAVCaptureDevice)
	ActivePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() unsafe.Pointer
	SetActivePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions(value unsafe.Pointer)
	ActivePrimaryConstituentDeviceSwitchingBehavior() unsafe.Pointer
	SetActivePrimaryConstituentDeviceSwitchingBehavior(value unsafe.Pointer)
	CameraLensSmudgeDetectionInterval() unsafe.Pointer
	SetCameraLensSmudgeDetectionInterval(value unsafe.Pointer)
	CameraLensSmudgeDetectionStatus() unsafe.Pointer
	SetCameraLensSmudgeDetectionStatus(value unsafe.Pointer)
	CinematicVideoCaptureSceneMonitoringStatuses() unsafe.Pointer
	SetCinematicVideoCaptureSceneMonitoringStatuses(value unsafe.Pointer)
	CompanionDeskViewCamera() IAVCaptureDevice
	SetCompanionDeskViewCamera(value IAVCaptureDevice)
	ConstituentDevices() IAVCaptureDevice
	SetConstituentDevices(value IAVCaptureDevice)
	DeviceType() unsafe.Pointer
	SetDeviceType(value unsafe.Pointer)
	DynamicAspectRatio() unsafe.Pointer
	SetDynamicAspectRatio(value unsafe.Pointer)
	DynamicDimensions() unsafe.Pointer
	SetDynamicDimensions(value unsafe.Pointer)
	FallbackPrimaryConstituentDevices() IAVCaptureDevice
	SetFallbackPrimaryConstituentDevices(value IAVCaptureDevice)
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
	LocalizedName() string
	SetLocalizedName(value string)
	Manufacturer() string
	SetManufacturer(value string)
	MinSupportedExternalSyncFrameDuration() unsafe.Pointer
	SetMinSupportedExternalSyncFrameDuration(value unsafe.Pointer)
	MinSupportedLockedVideoFrameDuration() unsafe.Pointer
	SetMinSupportedLockedVideoFrameDuration(value unsafe.Pointer)
	ModelID() string
	SetModelID(value string)
	NominalFocalLengthIn35mmFilm() float32
	SetNominalFocalLengthIn35mmFilm(value float32)
	Position() unsafe.Pointer
	SetPosition(value unsafe.Pointer)
	PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() unsafe.Pointer
	SetPrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions(value unsafe.Pointer)
	PrimaryConstituentDeviceSwitchingBehavior() unsafe.Pointer
	SetPrimaryConstituentDeviceSwitchingBehavior(value unsafe.Pointer)
	SmartFramingMonitor() CaptureSmartFramingMonitor
	SetSmartFramingMonitor(value CaptureSmartFramingMonitor)
	SpatialCaptureDiscomfortReasons() unsafe.Pointer
	SetSpatialCaptureDiscomfortReasons(value unsafe.Pointer)
	SupportedFallbackPrimaryConstituentDevices() IAVCaptureDevice
	SetSupportedFallbackPrimaryConstituentDevices(value IAVCaptureDevice)
	TransportType() unsafe.Pointer
	SetTransportType(value unsafe.Pointer)
	UniqueID() string
	SetUniqueID(value string)
	AVCaptureSessionInterruptionSystemPressureStateKey() string
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



// Returns the default device for the specified device type, media type, and position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/default(_:for:position:)
func (cc _CaptureDeviceClass) DefaultDeviceWithDeviceTypeMediaTypePosition(deviceType unsafe.Pointer, mediaType unsafe.Pointer, position unsafe.Pointer) ICaptureDevice {
	rv := objc.Send[CaptureDevice](objc.ID(cc.class), objc.Sel("defaultDeviceWithDeviceType:mediaType:position:"), deviceType, mediaType, position)
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

// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeFormat
func (c_ CaptureDevice) ActiveFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("activeFormat"))
	return rv
}


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeFormat
func (c_ CaptureDevice) SetActiveFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveFormat:"), value)
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


// A class property that indicates whether a person enables the Background Replacement feature for this app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isBackgroundReplacementEnabled
func (c_ CaptureDevice) BackgroundReplacementEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("backgroundReplacementEnabled"))
	return rv
}


// A value that indicates the capture device’s current system pressure state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/systemPressureState-swift.property
func (c_ CaptureDevice) SystemPressureState() CaptureSystemPressureState {
	rv := objc.Send[CaptureSystemPressureState](c_.ID, objc.Sel("systemPressureState"))
	return rv
}


// A virtual device’s active primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeprimaryconstituent
func (c_ CaptureDevice) ActivePrimaryConstituent() IAVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](c_.ID, objc.Sel("activePrimaryConstituent"))
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


// The switching behavior of the active constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeprimaryconstituentdeviceswitchingbehavior
func (c_ CaptureDevice) ActivePrimaryConstituentDeviceSwitchingBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("activePrimaryConstituentDeviceSwitchingBehavior"))
	return rv
}


// The switching behavior of the active constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeprimaryconstituentdeviceswitchingbehavior
func (c_ CaptureDevice) SetActivePrimaryConstituentDeviceSwitchingBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActivePrimaryConstituentDeviceSwitchingBehavior:"), value)
}


// The camera lens smudge detection interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/cameralenssmudgedetectioninterval
func (c_ CaptureDevice) CameraLensSmudgeDetectionInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("cameraLensSmudgeDetectionInterval"))
	return rv
}


// The camera lens smudge detection interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/cameralenssmudgedetectioninterval
func (c_ CaptureDevice) SetCameraLensSmudgeDetectionInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCameraLensSmudgeDetectionInterval:"), value)
}


// A value specifying the status of camera lens smudge detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/cameralenssmudgedetectionstatus
func (c_ CaptureDevice) CameraLensSmudgeDetectionStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("cameraLensSmudgeDetectionStatus"))
	return rv
}


// A value specifying the status of camera lens smudge detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/cameralenssmudgedetectionstatus
func (c_ CaptureDevice) SetCameraLensSmudgeDetectionStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCameraLensSmudgeDetectionStatus:"), value)
}


// The current scene monitoring statuses related to Cinematic Video capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/cinematicvideocapturescenemonitoringstatuses
func (c_ CaptureDevice) CinematicVideoCaptureSceneMonitoringStatuses() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("cinematicVideoCaptureSceneMonitoringStatuses"))
	return rv
}


// The current scene monitoring statuses related to Cinematic Video capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/cinematicvideocapturescenemonitoringstatuses
func (c_ CaptureDevice) SetCinematicVideoCaptureSceneMonitoringStatuses(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCinematicVideoCaptureSceneMonitoringStatuses:"), value)
}


// A Desk View camera associated with a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/companiondeskviewcamera
func (c_ CaptureDevice) CompanionDeskViewCamera() IAVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](c_.ID, objc.Sel("companionDeskViewCamera"))
	return rv
}


// A Desk View camera associated with a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/companiondeskviewcamera
func (c_ CaptureDevice) SetCompanionDeskViewCamera(value IAVCaptureDevice) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompanionDeskViewCamera:"), value)
}


// An array of physical devices that make up a virtual device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/constituentdevices
func (c_ CaptureDevice) ConstituentDevices() IAVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](c_.ID, objc.Sel("constituentDevices"))
	return rv
}


// An array of physical devices that make up a virtual device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/constituentdevices
func (c_ CaptureDevice) SetConstituentDevices(value IAVCaptureDevice) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConstituentDevices:"), value)
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


// A key-value observable property indicating the current aspect ratio for a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/dynamicaspectratio
func (c_ CaptureDevice) DynamicAspectRatio() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("dynamicAspectRatio"))
	return rv
}


// A key-value observable property indicating the current aspect ratio for a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/dynamicaspectratio
func (c_ CaptureDevice) SetDynamicAspectRatio(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDynamicAspectRatio:"), value)
}


// A key-value observable property describing the output dimensions of the video buffer based on the device’s dynamic aspect ratio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/dynamicdimensions
func (c_ CaptureDevice) DynamicDimensions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("dynamicDimensions"))
	return rv
}


// A key-value observable property describing the output dimensions of the video buffer based on the device’s dynamic aspect ratio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/dynamicdimensions
func (c_ CaptureDevice) SetDynamicDimensions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDynamicDimensions:"), value)
}


// The fallback devices to use when a constituent device with a longer focal length becomes limited by its light sensitivity or minimum focus distance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/fallbackprimaryconstituentdevices
func (c_ CaptureDevice) FallbackPrimaryConstituentDevices() IAVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](c_.ID, objc.Sel("fallbackPrimaryConstituentDevices"))
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


// A localized device name for display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/localizedname
func (c_ CaptureDevice) LocalizedName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("localizedName"))
	return rv
}


// A localized device name for display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/localizedname
func (c_ CaptureDevice) SetLocalizedName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocalizedName:"), objc.String(value))
}


// A human-readable string for the manufacturer of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/manufacturer
func (c_ CaptureDevice) Manufacturer() string {
	rv := objc.Send[string](c_.ID, objc.Sel("manufacturer"))
	return rv
}


// A human-readable string for the manufacturer of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/manufacturer
func (c_ CaptureDevice) SetManufacturer(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setManufacturer:"), objc.String(value))
}


// The minimum frame duration that can be passed as the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/minsupportedexternalsyncframeduration
func (c_ CaptureDevice) MinSupportedExternalSyncFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("minSupportedExternalSyncFrameDuration"))
	return rv
}


// The minimum frame duration that can be passed as the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/minsupportedexternalsyncframeduration
func (c_ CaptureDevice) SetMinSupportedExternalSyncFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinSupportedExternalSyncFrameDuration:"), value)
}


// The maximum frame rate (expressed as a minimum duration) that can be set on an input associated with this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/minsupportedlockedvideoframeduration
func (c_ CaptureDevice) MinSupportedLockedVideoFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("minSupportedLockedVideoFrameDuration"))
	return rv
}


// The maximum frame rate (expressed as a minimum duration) that can be set on an input associated with this device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/minsupportedlockedvideoframeduration
func (c_ CaptureDevice) SetMinSupportedLockedVideoFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMinSupportedLockedVideoFrameDuration:"), value)
}


// A model identifier for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/modelid
func (c_ CaptureDevice) ModelID() string {
	rv := objc.Send[string](c_.ID, objc.Sel("modelID"))
	return rv
}


// A model identifier for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/modelid
func (c_ CaptureDevice) SetModelID(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModelID:"), objc.String(value))
}


// The nominal 35mm equivalent focal length of the capture device’s lens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/nominalfocallengthin35mmfilm
func (c_ CaptureDevice) NominalFocalLengthIn35mmFilm() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("nominalFocalLengthIn35mmFilm"))
	return rv
}


// The nominal 35mm equivalent focal length of the capture device’s lens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/nominalfocallengthin35mmfilm
func (c_ CaptureDevice) SetNominalFocalLengthIn35mmFilm(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNominalFocalLengthIn35mmFilm:"), value)
}


// The physical position of the capture device hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/position-swift.property
func (c_ CaptureDevice) Position() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("position"))
	return rv
}


// The physical position of the capture device hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/position-swift.property
func (c_ CaptureDevice) SetPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPosition:"), value)
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
func (c_ CaptureDevice) SmartFramingMonitor() CaptureSmartFramingMonitor {
	rv := objc.Send[CaptureSmartFramingMonitor](c_.ID, objc.Sel("smartFramingMonitor"))
	return rv
}


// A monitor owned by the device that recommends an optimal framing based on the content in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/smartframingmonitor
func (c_ CaptureDevice) SetSmartFramingMonitor(value CaptureSmartFramingMonitor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSmartFramingMonitor:"), value)
}


// Reasons why current environmental conditions aren’t suitable to capturing spatial videos that are comfortable to view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/spatialcapturediscomfortreasons
func (c_ CaptureDevice) SpatialCaptureDiscomfortReasons() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("spatialCaptureDiscomfortReasons"))
	return rv
}


// Reasons why current environmental conditions aren’t suitable to capturing spatial videos that are comfortable to view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/spatialcapturediscomfortreasons
func (c_ CaptureDevice) SetSpatialCaptureDiscomfortReasons(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSpatialCaptureDiscomfortReasons:"), value)
}


// The constituent devices available to select as a fallback for a longer focal length primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/supportedfallbackprimaryconstituentdevices
func (c_ CaptureDevice) SupportedFallbackPrimaryConstituentDevices() IAVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](c_.ID, objc.Sel("supportedFallbackPrimaryConstituentDevices"))
	return rv
}


// The constituent devices available to select as a fallback for a longer focal length primary constituent device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/supportedfallbackprimaryconstituentdevices
func (c_ CaptureDevice) SetSupportedFallbackPrimaryConstituentDevices(value IAVCaptureDevice) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedFallbackPrimaryConstituentDevices:"), value)
}


// The transport type of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/transporttype
func (c_ CaptureDevice) TransportType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("transportType"))
	return rv
}


// The transport type of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/transporttype
func (c_ CaptureDevice) SetTransportType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTransportType:"), value)
}


// An identifier that uniquely identifies the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/uniqueid
func (c_ CaptureDevice) UniqueID() string {
	rv := objc.Send[string](c_.ID, objc.Sel("uniqueID"))
	return rv
}


// An identifier that uniquely identifies the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/uniqueid
func (c_ CaptureDevice) SetUniqueID(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUniqueID:"), objc.String(value))
}


// A key to retrieve a state value that indicates the system pressure level and contributing factors that caused the interruption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesessioninterruptionsystempressurestatekey
func (c_ CaptureDevice) AVCaptureSessionInterruptionSystemPressureStateKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("AVCaptureSessionInterruptionSystemPressureStateKey"))
	return rv
}



