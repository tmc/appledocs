// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CaptureDeviceInput] class.
var (
	CaptureDeviceInputClass     _CaptureDeviceInputClass
	CaptureDeviceInputClassOnce sync.Once
)

func getCaptureDeviceInputClass() _CaptureDeviceInputClass {
	CaptureDeviceInputClassOnce.Do(func() {
		CaptureDeviceInputClass = _CaptureDeviceInputClass{objc.GetClass("AVCaptureDeviceInput")}
	})
	return CaptureDeviceInputClass
}

type _CaptureDeviceInputClass struct {
	class objc.Class
}

// An interface definition for the [CaptureDeviceInput] class.
type ICaptureDeviceInput interface {
	ICaptureInput
	MultichannelAudioMode() unsafe.Pointer
	SetMultichannelAudioMode(value unsafe.Pointer)
	ActiveExternalSyncVideoFrameDuration() unsafe.Pointer
	SetActiveExternalSyncVideoFrameDuration(value unsafe.Pointer)
	ActiveLockedVideoFrameDuration() unsafe.Pointer
	SetActiveLockedVideoFrameDuration(value unsafe.Pointer)
	Device() IAVCaptureDevice
	SetDevice(value IAVCaptureDevice)
	ExternalSyncDevice() unsafe.Pointer
	SetExternalSyncDevice(value unsafe.Pointer)
	IsCinematicVideoCaptureEnabled() bool
	SetIsCinematicVideoCaptureEnabled(value bool)
	IsCinematicVideoCaptureSupported() bool
	SetIsCinematicVideoCaptureSupported(value bool)
	IsExternalSyncSupported() bool
	SetIsExternalSyncSupported(value bool)
	IsLockedVideoFrameDurationSupported() bool
	SetIsLockedVideoFrameDurationSupported(value bool)
	IsWindNoiseRemovalEnabled() bool
	SetIsWindNoiseRemovalEnabled(value bool)
	IsWindNoiseRemovalSupported() bool
	SetIsWindNoiseRemovalSupported(value bool)
	SimulatedAperture() float32
	SetSimulatedAperture(value float32)
	UnifiedAutoExposureDefaultsEnabled() bool
	SetUnifiedAutoExposureDefaultsEnabled(value bool)
	VideoMinFrameDurationOverride() unsafe.Pointer
	SetVideoMinFrameDurationOverride(value unsafe.Pointer)
}

// An object that provides media input from a capture device to a capture session.
//
// This class is a concrete subclass of that you use to connect a capture device to a capture session.


// An object that provides media input from a capture device to a capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput
type CaptureDeviceInput struct {
	CaptureInput
}

// CaptureDeviceInputFrom constructs a [CaptureDeviceInput] from an unsafe.Pointer.
//
// An object that provides media input from a capture device to a capture session.
func CaptureDeviceInputFrom(ptr unsafe.Pointer) CaptureDeviceInput {
	return CaptureDeviceInput{
		CaptureInput: CaptureInputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureDeviceInputClass) Alloc() CaptureDeviceInput {
	rv := objc.Send[CaptureDeviceInput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureDeviceInputClass) New() CaptureDeviceInput {
	rv := objc.Send[CaptureDeviceInput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureDeviceInput) Init() CaptureDeviceInput {
	rv := objc.Send[CaptureDeviceInput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureDeviceInput) Autorelease() CaptureDeviceInput {
	rv := objc.Send[CaptureDeviceInput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureDeviceInput creates a new CaptureDeviceInput instance.
func NewCaptureDeviceInput() CaptureDeviceInput {
	return getCaptureDeviceInputClass().New()
}



// The multichannel audio mode to apply when recording audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/multichannelAudioMode
func (c_ CaptureDeviceInput) MultichannelAudioMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("multichannelAudioMode"))
	return rv
}


// The multichannel audio mode to apply when recording audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/multichannelAudioMode
func (c_ CaptureDeviceInput) SetMultichannelAudioMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMultichannelAudioMode:"), value)
}


// The receiver’s external sync frame duration (the reciprocal of its frame rate) when being driven by an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/activeexternalsyncvideoframeduration
func (c_ CaptureDeviceInput) ActiveExternalSyncVideoFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("activeExternalSyncVideoFrameDuration"))
	return rv
}


// The receiver’s external sync frame duration (the reciprocal of its frame rate) when being driven by an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/activeexternalsyncvideoframeduration
func (c_ CaptureDeviceInput) SetActiveExternalSyncVideoFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveExternalSyncVideoFrameDuration:"), value)
}


// The receiver’s locked frame duration (the reciprocal of its frame rate). Setting this property guarantees the intra-frame duration delivered by the device input is precisely the frame duration you request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/activelockedvideoframeduration
func (c_ CaptureDeviceInput) ActiveLockedVideoFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("activeLockedVideoFrameDuration"))
	return rv
}


// The receiver’s locked frame duration (the reciprocal of its frame rate). Setting this property guarantees the intra-frame duration delivered by the device input is precisely the frame duration you request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/activelockedvideoframeduration
func (c_ CaptureDeviceInput) SetActiveLockedVideoFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveLockedVideoFrameDuration:"), value)
}


// A capture device associated with this input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/device
func (c_ CaptureDeviceInput) Device() IAVCaptureDevice {
	rv := objc.Send[CaptureDevice](c_.ID, objc.Sel("device"))
	return rv
}


// A capture device associated with this input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/device
func (c_ CaptureDeviceInput) SetDevice(value IAVCaptureDevice) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDevice:"), value)
}


// The external sync device currently being followed by this input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/externalsyncdevice
func (c_ CaptureDeviceInput) ExternalSyncDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("externalSyncDevice"))
	return rv
}


// The external sync device currently being followed by this input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/externalsyncdevice
func (c_ CaptureDeviceInput) SetExternalSyncDevice(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExternalSyncDevice:"), value)
}


// A BOOL value specifying whether the Cinematic Video effect is being applied to any movie file output, video data output, metadata output, or video preview layer added to the capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iscinematicvideocaptureenabled
func (c_ CaptureDeviceInput) IsCinematicVideoCaptureEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCinematicVideoCaptureEnabled"))
	return rv
}


// A BOOL value specifying whether the Cinematic Video effect is being applied to any movie file output, video data output, metadata output, or video preview layer added to the capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iscinematicvideocaptureenabled
func (c_ CaptureDeviceInput) SetIsCinematicVideoCaptureEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCinematicVideoCaptureEnabled:"), value)
}


// A BOOL value specifying whether Cinematic Video capture is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iscinematicvideocapturesupported
func (c_ CaptureDeviceInput) IsCinematicVideoCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCinematicVideoCaptureSupported"))
	return rv
}


// A BOOL value specifying whether Cinematic Video capture is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iscinematicvideocapturesupported
func (c_ CaptureDeviceInput) SetIsCinematicVideoCaptureSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCinematicVideoCaptureSupported:"), value)
}


// Indicates whether the device input supports being configured to follow an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/isexternalsyncsupported
func (c_ CaptureDeviceInput) IsExternalSyncSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isExternalSyncSupported"))
	return rv
}


// Indicates whether the device input supports being configured to follow an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/isexternalsyncsupported
func (c_ CaptureDeviceInput) SetIsExternalSyncSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsExternalSyncSupported:"), value)
}


// Indicates whether the device input supports locked frame durations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/islockedvideoframedurationsupported
func (c_ CaptureDeviceInput) IsLockedVideoFrameDurationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLockedVideoFrameDurationSupported"))
	return rv
}


// Indicates whether the device input supports locked frame durations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/islockedvideoframedurationsupported
func (c_ CaptureDeviceInput) SetIsLockedVideoFrameDurationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLockedVideoFrameDurationSupported:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iswindnoiseremovalenabled
func (c_ CaptureDeviceInput) IsWindNoiseRemovalEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isWindNoiseRemovalEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iswindnoiseremovalenabled
func (c_ CaptureDeviceInput) SetIsWindNoiseRemovalEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsWindNoiseRemovalEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iswindnoiseremovalsupported
func (c_ CaptureDeviceInput) IsWindNoiseRemovalSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isWindNoiseRemovalSupported"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iswindnoiseremovalsupported
func (c_ CaptureDeviceInput) SetIsWindNoiseRemovalSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsWindNoiseRemovalSupported:"), value)
}


// Shallow depth of field simulated aperture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/simulatedaperture
func (c_ CaptureDeviceInput) SimulatedAperture() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("simulatedAperture"))
	return rv
}


// Shallow depth of field simulated aperture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/simulatedaperture
func (c_ CaptureDeviceInput) SetSimulatedAperture(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSimulatedAperture:"), value)
}


// A Boolean value that indicates whether the input enables unified auto-exposure defaults.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/unifiedautoexposuredefaultsenabled
func (c_ CaptureDeviceInput) UnifiedAutoExposureDefaultsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("unifiedAutoExposureDefaultsEnabled"))
	return rv
}


// A Boolean value that indicates whether the input enables unified auto-exposure defaults.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/unifiedautoexposuredefaultsenabled
func (c_ CaptureDeviceInput) SetUnifiedAutoExposureDefaultsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUnifiedAutoExposureDefaultsEnabled:"), value)
}


// A time value that acts as a modifier to a capture device’s active video minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/videominframedurationoverride
func (c_ CaptureDeviceInput) VideoMinFrameDurationOverride() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("videoMinFrameDurationOverride"))
	return rv
}


// A time value that acts as a modifier to a capture device’s active video minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/videominframedurationoverride
func (c_ CaptureDeviceInput) SetVideoMinFrameDurationOverride(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMinFrameDurationOverride:"), value)
}



