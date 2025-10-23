// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	ActiveExternalSyncVideoFrameDuration() Time /* not a class type */
	ActiveLockedVideoFrameDuration() Time /* not a class type */
	SetActiveLockedVideoFrameDuration(value Time /* not a class type */)
	Device() IAVCaptureDevice
	ExternalSyncDevice() ExternalSyncDevice /* not a class type */
	CinematicVideoCaptureEnabled() bool /* primitive/slice/pointer. */
	SetCinematicVideoCaptureEnabled(value bool /* primitive/slice/pointer. */)
	CinematicVideoCaptureSupported() bool /* primitive/slice/pointer. */
	ExternalSyncSupported() bool /* primitive/slice/pointer. */
	LockedVideoFrameDurationSupported() bool /* primitive/slice/pointer. */
	WindNoiseRemovalEnabled() bool /* primitive/slice/pointer. */
	SetWindNoiseRemovalEnabled(value bool /* primitive/slice/pointer. */)
	WindNoiseRemovalSupported() bool /* primitive/slice/pointer. */
	MultichannelAudioMode() CaptureMultichannelAudioMode
	SetMultichannelAudioMode(value CaptureMultichannelAudioMode)
	SimulatedAperture() float32 /* primitive/slice/pointer. */
	SetSimulatedAperture(value float32 /* primitive/slice/pointer. */)
	UnifiedAutoExposureDefaultsEnabled() bool /* primitive/slice/pointer. */
	SetUnifiedAutoExposureDefaultsEnabled(value bool /* primitive/slice/pointer. */)
	VideoMinFrameDurationOverride() Time /* not a class type */
	SetVideoMinFrameDurationOverride(value Time /* not a class type */)
	IsCinematicVideoCaptureEnabled() bool /* primitive/slice/pointer. */
	SetIsCinematicVideoCaptureEnabled(value bool /* primitive/slice/pointer. */)
	IsCinematicVideoCaptureSupported() bool /* primitive/slice/pointer. */
	SetIsCinematicVideoCaptureSupported(value bool /* primitive/slice/pointer. */)
	IsExternalSyncSupported() bool /* primitive/slice/pointer. */
	SetIsExternalSyncSupported(value bool /* primitive/slice/pointer. */)
	IsLockedVideoFrameDurationSupported() bool /* primitive/slice/pointer. */
	SetIsLockedVideoFrameDurationSupported(value bool /* primitive/slice/pointer. */)
	IsWindNoiseRemovalEnabled() bool /* primitive/slice/pointer. */
	SetIsWindNoiseRemovalEnabled(value bool /* primitive/slice/pointer. */)
	IsWindNoiseRemovalSupported() bool /* primitive/slice/pointer. */
	SetIsWindNoiseRemovalSupported(value bool /* primitive/slice/pointer. */)
	// methods:
	FollowExternalSyncDeviceVideoFrameDurationDelegate(externalSyncDevice ExternalSyncDevice /* not a class type */, frameDuration Time /* not a class type */, delegate objectivec.IObject)
	IsMultichannelAudioModeSupported(multichannelAudioMode CaptureMultichannelAudioMode) bool /* primitive/slice/pointer. */
	PortsWithMediaTypeSourceDeviceTypeSourceDevicePosition(mediaType MediaType /* not a class type */, sourceDeviceType objc.IObject /* cross-framework CaptureDeviceType */, sourceDevicePosition CaptureDevicePosition) []CaptureInputPort /* primitive/slice/pointer. */
	UnfollowExternalSyncDevice()
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



// Creates an input for the specified capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/init(device:)
func NewCaptureDeviceInputWithDeviceError(device IAVCaptureDevice, outError unsafe.Pointer) CaptureDeviceInput {
	instance := getCaptureDeviceInputClass().Alloc()
	rv := objc.Send[CaptureDeviceInput](instance.ID, objc.Sel("initWithDevice:error:"), device, outError)
	rv.Autorelease()
	return rv
}



// Returns a new input for the specified capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/deviceInputWithDevice:error:
func (cc _CaptureDeviceInputClass) DeviceInputWithDeviceError(device IAVCaptureDevice, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("deviceInputWithDevice:error:"), device, outError)
	return rv
}


// Configures the the device input to follow an external sync device at the given frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/follow(_:videoFrameDuration:delegate:)
func (c_ CaptureDeviceInput) FollowExternalSyncDeviceVideoFrameDurationDelegate(externalSyncDevice ExternalSyncDevice /* not a class type */, frameDuration Time /* not a class type */, delegate objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("followExternalSyncDevice:videoFrameDuration:delegate:"), externalSyncDevice, frameDuration, delegate)
}


// A Boolean value that indicates whether the input supports the specified multichannel audio mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isMultichannelAudioModeSupported(_:)
func (c_ CaptureDeviceInput) IsMultichannelAudioModeSupported(multichannelAudioMode CaptureMultichannelAudioMode) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isMultichannelAudioModeSupported:"), multichannelAudioMode)
	return rv
}


// Retrieves a virtual device’s constituent device ports for use in a multi-camera session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/ports(for:sourceDeviceType:sourceDevicePosition:)
func (c_ CaptureDeviceInput) PortsWithMediaTypeSourceDeviceTypeSourceDevicePosition(mediaType MediaType /* not a class type */, sourceDeviceType objc.IObject /* cross-framework CaptureDeviceType */, sourceDevicePosition CaptureDevicePosition) []CaptureInputPort /* primitive/slice/pointer. */ {
	rv := objc.Send[[]CaptureInputPort](c_.ID, objc.Sel("portsWithMediaType:sourceDeviceType:sourceDevicePosition:"), mediaType, sourceDeviceType, sourceDevicePosition)
	return rv
}


// Discontinues external sync.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/unfollowExternalSyncDevice()
func (c_ CaptureDeviceInput) UnfollowExternalSyncDevice() {
	objc.Send[objc.ID](c_.ID, objc.Sel("unfollowExternalSyncDevice"))
}


// The receiver’s external sync frame duration (the reciprocal of its frame rate) when being driven by an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/activeExternalSyncVideoFrameDuration
func (c_ CaptureDeviceInput) ActiveExternalSyncVideoFrameDuration() Time /* not a class type */ {
	rv := objc.Send[Time](c_.ID, objc.Sel("activeExternalSyncVideoFrameDuration"))
	return rv
}


// The receiver’s locked frame duration (the reciprocal of its frame rate). Setting this property guarantees the intra-frame duration delivered by the device input is precisely the frame duration you request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/activeLockedVideoFrameDuration
func (c_ CaptureDeviceInput) ActiveLockedVideoFrameDuration() Time /* not a class type */ {
	rv := objc.Send[Time](c_.ID, objc.Sel("activeLockedVideoFrameDuration"))
	return rv
}


// The receiver’s locked frame duration (the reciprocal of its frame rate). Setting this property guarantees the intra-frame duration delivered by the device input is precisely the frame duration you request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/activeLockedVideoFrameDuration
func (c_ CaptureDeviceInput) SetActiveLockedVideoFrameDuration(value Time /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveLockedVideoFrameDuration:"), value)
}


// A capture device associated with this input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/device
func (c_ CaptureDeviceInput) Device() IAVCaptureDevice {
	rv := objc.Send[CaptureDevice](c_.ID, objc.Sel("device"))
	return rv
}


// The external sync device currently being followed by this input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/externalSyncDevice
func (c_ CaptureDeviceInput) ExternalSyncDevice() ExternalSyncDevice /* not a class type */ {
	rv := objc.Send[ExternalSyncDevice](c_.ID, objc.Sel("externalSyncDevice"))
	return rv
}


// A BOOL value specifying whether the Cinematic Video effect is being applied to any movie file output, video data output, metadata output, or video preview layer added to the capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isCinematicVideoCaptureEnabled
func (c_ CaptureDeviceInput) CinematicVideoCaptureEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("cinematicVideoCaptureEnabled"))
	return rv
}


// A BOOL value specifying whether the Cinematic Video effect is being applied to any movie file output, video data output, metadata output, or video preview layer added to the capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isCinematicVideoCaptureEnabled
func (c_ CaptureDeviceInput) SetCinematicVideoCaptureEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCinematicVideoCaptureEnabled:"), value)
}


// A BOOL value specifying whether Cinematic Video capture is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isCinematicVideoCaptureSupported
func (c_ CaptureDeviceInput) CinematicVideoCaptureSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("cinematicVideoCaptureSupported"))
	return rv
}


// Indicates whether the device input supports being configured to follow an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isExternalSyncSupported
func (c_ CaptureDeviceInput) ExternalSyncSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("externalSyncSupported"))
	return rv
}


// Indicates whether the device input supports locked frame durations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isLockedVideoFrameDurationSupported
func (c_ CaptureDeviceInput) LockedVideoFrameDurationSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("lockedVideoFrameDurationSupported"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isWindNoiseRemovalEnabled
func (c_ CaptureDeviceInput) WindNoiseRemovalEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("windNoiseRemovalEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isWindNoiseRemovalEnabled
func (c_ CaptureDeviceInput) SetWindNoiseRemovalEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWindNoiseRemovalEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isWindNoiseRemovalSupported
func (c_ CaptureDeviceInput) WindNoiseRemovalSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("windNoiseRemovalSupported"))
	return rv
}


// The multichannel audio mode to apply when recording audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/multichannelAudioMode
func (c_ CaptureDeviceInput) MultichannelAudioMode() CaptureMultichannelAudioMode {
	rv := objc.Send[CaptureMultichannelAudioMode](c_.ID, objc.Sel("multichannelAudioMode"))
	return rv
}


// The multichannel audio mode to apply when recording audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/multichannelAudioMode
func (c_ CaptureDeviceInput) SetMultichannelAudioMode(value CaptureMultichannelAudioMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMultichannelAudioMode:"), value)
}


// Shallow depth of field simulated aperture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/simulatedAperture
func (c_ CaptureDeviceInput) SimulatedAperture() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](c_.ID, objc.Sel("simulatedAperture"))
	return rv
}


// Shallow depth of field simulated aperture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/simulatedAperture
func (c_ CaptureDeviceInput) SetSimulatedAperture(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSimulatedAperture:"), value)
}


// A Boolean value that indicates whether the input enables unified auto-exposure defaults.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/unifiedAutoExposureDefaultsEnabled
func (c_ CaptureDeviceInput) UnifiedAutoExposureDefaultsEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("unifiedAutoExposureDefaultsEnabled"))
	return rv
}


// A Boolean value that indicates whether the input enables unified auto-exposure defaults.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/unifiedAutoExposureDefaultsEnabled
func (c_ CaptureDeviceInput) SetUnifiedAutoExposureDefaultsEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUnifiedAutoExposureDefaultsEnabled:"), value)
}


// A time value that acts as a modifier to a capture device’s active video minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/videoMinFrameDurationOverride
func (c_ CaptureDeviceInput) VideoMinFrameDurationOverride() Time /* not a class type */ {
	rv := objc.Send[Time](c_.ID, objc.Sel("videoMinFrameDurationOverride"))
	return rv
}


// A time value that acts as a modifier to a capture device’s active video minimum frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/videoMinFrameDurationOverride
func (c_ CaptureDeviceInput) SetVideoMinFrameDurationOverride(value Time /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMinFrameDurationOverride:"), value)
}


// A BOOL value specifying whether the Cinematic Video effect is being applied to any movie file output, video data output, metadata output, or video preview layer added to the capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iscinematicvideocaptureenabled
func (c_ CaptureDeviceInput) IsCinematicVideoCaptureEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCinematicVideoCaptureEnabled"))
	return rv
}


// A BOOL value specifying whether the Cinematic Video effect is being applied to any movie file output, video data output, metadata output, or video preview layer added to the capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iscinematicvideocaptureenabled
func (c_ CaptureDeviceInput) SetIsCinematicVideoCaptureEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCinematicVideoCaptureEnabled:"), value)
}


// A BOOL value specifying whether Cinematic Video capture is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iscinematicvideocapturesupported
func (c_ CaptureDeviceInput) IsCinematicVideoCaptureSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCinematicVideoCaptureSupported"))
	return rv
}


// A BOOL value specifying whether Cinematic Video capture is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iscinematicvideocapturesupported
func (c_ CaptureDeviceInput) SetIsCinematicVideoCaptureSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCinematicVideoCaptureSupported:"), value)
}


// Indicates whether the device input supports being configured to follow an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/isexternalsyncsupported
func (c_ CaptureDeviceInput) IsExternalSyncSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isExternalSyncSupported"))
	return rv
}


// Indicates whether the device input supports being configured to follow an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/isexternalsyncsupported
func (c_ CaptureDeviceInput) SetIsExternalSyncSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsExternalSyncSupported:"), value)
}


// Indicates whether the device input supports locked frame durations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/islockedvideoframedurationsupported
func (c_ CaptureDeviceInput) IsLockedVideoFrameDurationSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLockedVideoFrameDurationSupported"))
	return rv
}


// Indicates whether the device input supports locked frame durations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/islockedvideoframedurationsupported
func (c_ CaptureDeviceInput) SetIsLockedVideoFrameDurationSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLockedVideoFrameDurationSupported:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iswindnoiseremovalenabled
func (c_ CaptureDeviceInput) IsWindNoiseRemovalEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isWindNoiseRemovalEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iswindnoiseremovalenabled
func (c_ CaptureDeviceInput) SetIsWindNoiseRemovalEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsWindNoiseRemovalEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iswindnoiseremovalsupported
func (c_ CaptureDeviceInput) IsWindNoiseRemovalSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isWindNoiseRemovalSupported"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iswindnoiseremovalsupported
func (c_ CaptureDeviceInput) SetIsWindNoiseRemovalSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsWindNoiseRemovalSupported:"), value)
}


