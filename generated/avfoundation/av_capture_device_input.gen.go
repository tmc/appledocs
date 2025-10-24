// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureDeviceInput */


/* debug [class_header]: Header for AVCaptureDeviceInput */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureDeviceInput */
// An interface definition for the [CaptureDeviceInput] class.
type ICaptureDeviceInput interface {
	ICaptureInput
	
/* debug [class_interface_properties]: Properties for CaptureDeviceInput */
	// properties:
	ActiveExternalSyncVideoFrameDuration() objc.IObject /* cross-framework: Time */
	ActiveLockedVideoFrameDuration() objc.IObject /* cross-framework: Time */
	SetActiveLockedVideoFrameDuration(value objc.IObject /* cross-framework: Time */)
	Device() IAVCaptureDevice
	ExternalSyncDevice() IAVExternalSyncDevice
	CinematicVideoCaptureEnabled() bool
	SetCinematicVideoCaptureEnabled(value bool)
	CinematicVideoCaptureSupported() bool
	ExternalSyncSupported() bool
	LockedVideoFrameDurationSupported() bool
	WindNoiseRemovalEnabled() bool
	SetWindNoiseRemovalEnabled(value bool)
	WindNoiseRemovalSupported() bool
	MultichannelAudioMode() CaptureMultichannelAudioMode
	SetMultichannelAudioMode(value CaptureMultichannelAudioMode)
	SimulatedAperture() float32
	SetSimulatedAperture(value float32)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureDeviceInput */
	// methods:
	FollowExternalSyncDeviceVideoFrameDurationDelegate(externalSyncDevice IAVExternalSyncDevice, frameDuration objc.IObject /* cross-framework: Time */, delegate unsafe.Pointer)
	IsMultichannelAudioModeSupported(multichannelAudioMode CaptureMultichannelAudioMode) bool
	UnfollowExternalSyncDevice()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureDeviceInput */
// Alloc allocates a new instance without initialization.
func (cc _CaptureDeviceInputClass) Alloc() CaptureDeviceInput {
	rv := objc.Send[CaptureDeviceInput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureDeviceInput */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureDeviceInput */

// Creates an input for the specified capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/init(device:)
func NewCaptureDeviceInputWithDeviceError(device IAVCaptureDevice, outError objectivec.IObject) CaptureDeviceInput {
	instance := getCaptureDeviceInputClass().Alloc()
	rv := objc.Send[CaptureDeviceInput](instance.ID, objc.Sel("initWithDevice:error:"), device, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureDeviceInputWithDeviceError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureDeviceInput */

// Returns a new input for the specified capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/deviceInputWithDevice:error:
func (cc _CaptureDeviceInputClass) DeviceInputWithDeviceError(device IAVCaptureDevice, outError objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("deviceInputWithDevice:error:"), device, outError)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DeviceInputWithDeviceError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureDeviceInput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureDeviceInput */

// Configures the the device input to follow an external sync device at the given frame duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/follow(_:videoFrameDuration:delegate:)
func (c_ CaptureDeviceInput) FollowExternalSyncDeviceVideoFrameDurationDelegate(externalSyncDevice IAVExternalSyncDevice, frameDuration objc.IObject /* cross-framework: Time */, delegate unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("followExternalSyncDevice:videoFrameDuration:delegate:"), externalSyncDevice, frameDuration, delegate)
}/* debug [instance_methods/method]: FollowExternalSyncDeviceVideoFrameDurationDelegate */


// A Boolean value that indicates whether the input supports the specified multichannel audio mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isMultichannelAudioModeSupported(_:)
func (c_ CaptureDeviceInput) IsMultichannelAudioModeSupported(multichannelAudioMode CaptureMultichannelAudioMode) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isMultichannelAudioModeSupported:"), multichannelAudioMode)
	return rv
}/* debug [instance_methods/method]: IsMultichannelAudioModeSupported */


// Discontinues external sync.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/unfollowExternalSyncDevice()
func (c_ CaptureDeviceInput) UnfollowExternalSyncDevice() {
	objc.Send[objc.ID](c_.ID, objc.Sel("unfollowExternalSyncDevice"))
}/* debug [instance_methods/method]: UnfollowExternalSyncDevice */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureDeviceInput */

// The receiver’s external sync frame duration (the reciprocal of its frame rate) when being driven by an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/activeExternalSyncVideoFrameDuration
func (c_ CaptureDeviceInput) ActiveExternalSyncVideoFrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("activeExternalSyncVideoFrameDuration"))
	return rv
}/* debug [instance_properties/getter]: activeExternalSyncVideoFrameDuration */


// The receiver’s locked frame duration (the reciprocal of its frame rate). Setting this property guarantees the intra-frame duration delivered by the device input is precisely the frame duration you request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/activeLockedVideoFrameDuration
func (c_ CaptureDeviceInput) ActiveLockedVideoFrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("activeLockedVideoFrameDuration"))
	return rv
}/* debug [instance_properties/getter]: activeLockedVideoFrameDuration */


// The receiver’s locked frame duration (the reciprocal of its frame rate). Setting this property guarantees the intra-frame duration delivered by the device input is precisely the frame duration you request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/activeLockedVideoFrameDuration
func (c_ CaptureDeviceInput) SetActiveLockedVideoFrameDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveLockedVideoFrameDuration:"), value)
}/* debug [instance_properties/setter]: activeLockedVideoFrameDuration */


// A capture device associated with this input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/device
func (c_ CaptureDeviceInput) Device() IAVCaptureDevice {
	rv := objc.Send[CaptureDevice](c_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// The external sync device currently being followed by this input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/externalSyncDevice
func (c_ CaptureDeviceInput) ExternalSyncDevice() IAVExternalSyncDevice {
	rv := objc.Send[ExternalSyncDevice](c_.ID, objc.Sel("externalSyncDevice"))
	return rv
}/* debug [instance_properties/getter]: externalSyncDevice */


// A BOOL value specifying whether the Cinematic Video effect is being applied to any movie file output, video data output, metadata output, or video preview layer added to the capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isCinematicVideoCaptureEnabled
func (c_ CaptureDeviceInput) CinematicVideoCaptureEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("cinematicVideoCaptureEnabled"))
	return rv
}/* debug [instance_properties/getter]: cinematicVideoCaptureEnabled */


// A BOOL value specifying whether the Cinematic Video effect is being applied to any movie file output, video data output, metadata output, or video preview layer added to the capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isCinematicVideoCaptureEnabled
func (c_ CaptureDeviceInput) SetCinematicVideoCaptureEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCinematicVideoCaptureEnabled:"), value)
}/* debug [instance_properties/setter]: cinematicVideoCaptureEnabled */


// A BOOL value specifying whether Cinematic Video capture is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isCinematicVideoCaptureSupported
func (c_ CaptureDeviceInput) CinematicVideoCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("cinematicVideoCaptureSupported"))
	return rv
}/* debug [instance_properties/getter]: cinematicVideoCaptureSupported */


// Indicates whether the device input supports being configured to follow an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isExternalSyncSupported
func (c_ CaptureDeviceInput) ExternalSyncSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("externalSyncSupported"))
	return rv
}/* debug [instance_properties/getter]: externalSyncSupported */


// Indicates whether the device input supports locked frame durations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isLockedVideoFrameDurationSupported
func (c_ CaptureDeviceInput) LockedVideoFrameDurationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("lockedVideoFrameDurationSupported"))
	return rv
}/* debug [instance_properties/getter]: lockedVideoFrameDurationSupported */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isWindNoiseRemovalEnabled
func (c_ CaptureDeviceInput) WindNoiseRemovalEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("windNoiseRemovalEnabled"))
	return rv
}/* debug [instance_properties/getter]: windNoiseRemovalEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isWindNoiseRemovalEnabled
func (c_ CaptureDeviceInput) SetWindNoiseRemovalEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWindNoiseRemovalEnabled:"), value)
}/* debug [instance_properties/setter]: windNoiseRemovalEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isWindNoiseRemovalSupported
func (c_ CaptureDeviceInput) WindNoiseRemovalSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("windNoiseRemovalSupported"))
	return rv
}/* debug [instance_properties/getter]: windNoiseRemovalSupported */


// The multichannel audio mode to apply when recording audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/multichannelAudioMode
func (c_ CaptureDeviceInput) MultichannelAudioMode() CaptureMultichannelAudioMode {
	rv := objc.Send[CaptureMultichannelAudioMode](c_.ID, objc.Sel("multichannelAudioMode"))
	return rv
}/* debug [instance_properties/getter]: multichannelAudioMode */


// The multichannel audio mode to apply when recording audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/multichannelAudioMode
func (c_ CaptureDeviceInput) SetMultichannelAudioMode(value CaptureMultichannelAudioMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMultichannelAudioMode:"), value)
}/* debug [instance_properties/setter]: multichannelAudioMode */


// Shallow depth of field simulated aperture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/simulatedAperture
func (c_ CaptureDeviceInput) SimulatedAperture() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("simulatedAperture"))
	return rv
}/* debug [instance_properties/getter]: simulatedAperture */


// Shallow depth of field simulated aperture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/simulatedAperture
func (c_ CaptureDeviceInput) SetSimulatedAperture(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSimulatedAperture:"), value)
}/* debug [instance_properties/setter]: simulatedAperture */


// A BOOL value specifying whether the Cinematic Video effect is being applied to any movie file output, video data output, metadata output, or video preview layer added to the capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iscinematicvideocaptureenabled
func (c_ CaptureDeviceInput) IsCinematicVideoCaptureEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCinematicVideoCaptureEnabled"))
	return rv
}/* debug [instance_properties/getter]: isCinematicVideoCaptureEnabled */


// A BOOL value specifying whether the Cinematic Video effect is being applied to any movie file output, video data output, metadata output, or video preview layer added to the capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iscinematicvideocaptureenabled
func (c_ CaptureDeviceInput) SetIsCinematicVideoCaptureEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCinematicVideoCaptureEnabled:"), value)
}/* debug [instance_properties/setter]: isCinematicVideoCaptureEnabled */


// A BOOL value specifying whether Cinematic Video capture is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iscinematicvideocapturesupported
func (c_ CaptureDeviceInput) IsCinematicVideoCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCinematicVideoCaptureSupported"))
	return rv
}/* debug [instance_properties/getter]: isCinematicVideoCaptureSupported */


// A BOOL value specifying whether Cinematic Video capture is supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iscinematicvideocapturesupported
func (c_ CaptureDeviceInput) SetIsCinematicVideoCaptureSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsCinematicVideoCaptureSupported:"), value)
}/* debug [instance_properties/setter]: isCinematicVideoCaptureSupported */


// Indicates whether the device input supports being configured to follow an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/isexternalsyncsupported
func (c_ CaptureDeviceInput) IsExternalSyncSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isExternalSyncSupported"))
	return rv
}/* debug [instance_properties/getter]: isExternalSyncSupported */


// Indicates whether the device input supports being configured to follow an external sync device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/isexternalsyncsupported
func (c_ CaptureDeviceInput) SetIsExternalSyncSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsExternalSyncSupported:"), value)
}/* debug [instance_properties/setter]: isExternalSyncSupported */


// Indicates whether the device input supports locked frame durations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/islockedvideoframedurationsupported
func (c_ CaptureDeviceInput) IsLockedVideoFrameDurationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isLockedVideoFrameDurationSupported"))
	return rv
}/* debug [instance_properties/getter]: isLockedVideoFrameDurationSupported */


// Indicates whether the device input supports locked frame durations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/islockedvideoframedurationsupported
func (c_ CaptureDeviceInput) SetIsLockedVideoFrameDurationSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsLockedVideoFrameDurationSupported:"), value)
}/* debug [instance_properties/setter]: isLockedVideoFrameDurationSupported */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iswindnoiseremovalenabled
func (c_ CaptureDeviceInput) IsWindNoiseRemovalEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isWindNoiseRemovalEnabled"))
	return rv
}/* debug [instance_properties/getter]: isWindNoiseRemovalEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iswindnoiseremovalenabled
func (c_ CaptureDeviceInput) SetIsWindNoiseRemovalEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsWindNoiseRemovalEnabled:"), value)
}/* debug [instance_properties/setter]: isWindNoiseRemovalEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iswindnoiseremovalsupported
func (c_ CaptureDeviceInput) IsWindNoiseRemovalSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isWindNoiseRemovalSupported"))
	return rv
}/* debug [instance_properties/getter]: isWindNoiseRemovalSupported */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceinput/iswindnoiseremovalsupported
func (c_ CaptureDeviceInput) SetIsWindNoiseRemovalSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsWindNoiseRemovalSupported:"), value)
}/* debug [instance_properties/setter]: isWindNoiseRemovalSupported */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureDeviceInput */


