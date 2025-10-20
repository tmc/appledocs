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
	FollowExternalSyncDeviceVideoFrameDurationDelegate(externalSyncDevice unsafe.Pointer, frameDuration unsafe.Pointer, delegate objc.ID)
}

// An object that provides media input from a capture device to a capture session.
//
// This class is a concrete subclass of that you use to connect a capture device to a capture session.
//
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
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/init(device:)
func NewCaptureDeviceInputWithDeviceError(device unsafe.Pointer, outError unsafe.Pointer) CaptureDeviceInput {
	instance := getCaptureDeviceInputClass().Alloc()
	rv := objc.Send[CaptureDeviceInput](instance.ID, objc.Sel("initWithDevice:error:"), device, outError)
	rv.Autorelease()
	return rv
}


// Configures the the device input to follow an external sync device at the given frame duration.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/follow(_:videoFrameDuration:delegate:)
func (c_ CaptureDeviceInput) FollowExternalSyncDeviceVideoFrameDurationDelegate(externalSyncDevice unsafe.Pointer, frameDuration unsafe.Pointer, delegate objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("followExternalSyncDevice:videoFrameDuration:delegate:"), externalSyncDevice, frameDuration, delegate)
}

// The receiver’s locked frame duration (the reciprocal of its frame rate). Setting this property guarantees the intra-frame duration delivered by the device input is precisely the frame duration you request.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/activeLockedVideoFrameDuration
func (c_ CaptureDeviceInput) ActiveLockedVideoFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("activeLockedVideoFrameDuration"))
	return rv
}


// SetActiveLockedVideoFrameDuration sets the value of the activeLockedVideoFrameDuration property.
// The receiver’s locked frame duration (the reciprocal of its frame rate). Setting this property guarantees the intra-frame duration delivered by the device input is precisely the frame duration you request.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/activeLockedVideoFrameDuration
func (c_ CaptureDeviceInput) SetActiveLockedVideoFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveLockedVideoFrameDuration:"), value)
}
// The external sync device currently being followed by this input.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/externalSyncDevice
func (c_ CaptureDeviceInput) ExternalSyncDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("externalSyncDevice"))
	return rv
}

// A BOOL value specifying whether Cinematic Video capture is supported.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isCinematicVideoCaptureSupported
func (c_ CaptureDeviceInput) CinematicVideoCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("cinematicVideoCaptureSupported"))
	return rv
}

// Indicates whether the device input supports locked frame durations.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isLockedVideoFrameDurationSupported
func (c_ CaptureDeviceInput) LockedVideoFrameDurationSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("lockedVideoFrameDurationSupported"))
	return rv
}

// The multichannel audio mode to apply when recording audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/multichannelAudioMode
func (c_ CaptureDeviceInput) MultichannelAudioMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("multichannelAudioMode"))
	return rv
}


// SetMultichannelAudioMode sets the value of the multichannelAudioMode property.
// The multichannel audio mode to apply when recording audio.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/multichannelAudioMode
func (c_ CaptureDeviceInput) SetMultichannelAudioMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMultichannelAudioMode:"), value)
}
// A Boolean value that indicates whether the input enables unified auto-exposure defaults.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/unifiedAutoExposureDefaultsEnabled
func (c_ CaptureDeviceInput) UnifiedAutoExposureDefaultsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("unifiedAutoExposureDefaultsEnabled"))
	return rv
}


// SetUnifiedAutoExposureDefaultsEnabled sets the value of the unifiedAutoExposureDefaultsEnabled property.
// A Boolean value that indicates whether the input enables unified auto-exposure defaults.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/unifiedAutoExposureDefaultsEnabled
func (c_ CaptureDeviceInput) SetUnifiedAutoExposureDefaultsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUnifiedAutoExposureDefaultsEnabled:"), value)
}
// A time value that acts as a modifier to a capture device’s active video minimum frame duration.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/videoMinFrameDurationOverride
func (c_ CaptureDeviceInput) VideoMinFrameDurationOverride() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("videoMinFrameDurationOverride"))
	return rv
}


// SetVideoMinFrameDurationOverride sets the value of the videoMinFrameDurationOverride property.
// A time value that acts as a modifier to a capture device’s active video minimum frame duration.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/videoMinFrameDurationOverride
func (c_ CaptureDeviceInput) SetVideoMinFrameDurationOverride(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMinFrameDurationOverride:"), value)
}

