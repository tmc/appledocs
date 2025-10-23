// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CaptureEventSound] class.
var (
	CaptureEventSoundClass     _CaptureEventSoundClass
	CaptureEventSoundClassOnce sync.Once
)

func getCaptureEventSoundClass() _CaptureEventSoundClass {
	CaptureEventSoundClassOnce.Do(func() {
		CaptureEventSoundClass = _CaptureEventSoundClass{objc.GetClass("AVCaptureEventSound")}
	})
	return CaptureEventSoundClass
}

type _CaptureEventSoundClass struct {
	class objc.Class
}

// An interface definition for the [CaptureEventSound] class.
type ICaptureEventSound interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A sound object for a capture event.


// A sound object for a capture event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventSound
type CaptureEventSound struct {
	objectivec.Object
}

// CaptureEventSoundFrom constructs a [CaptureEventSound] from an unsafe.Pointer.
//
// A sound object for a capture event.
func CaptureEventSoundFrom(ptr unsafe.Pointer) CaptureEventSound {
	return CaptureEventSound{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureEventSoundClass) Alloc() CaptureEventSound {
	rv := objc.Send[CaptureEventSound](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureEventSoundClass) New() CaptureEventSound {
	rv := objc.Send[CaptureEventSound](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureEventSound) Init() CaptureEventSound {
	rv := objc.Send[CaptureEventSound](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureEventSound) Autorelease() CaptureEventSound {
	rv := objc.Send[CaptureEventSound](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureEventSound creates a new CaptureEventSound instance.
func NewCaptureEventSound() CaptureEventSound {
	return getCaptureEventSoundClass().New()
}



// Creates a sound object for a capture event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventSound/init(url:)
func NewCaptureEventSoundWithURLError(url foundation.objc.IObject /* cross-framework URL */, error_ unsafe.Pointer) CaptureEventSound {
	instance := getCaptureEventSoundClass().Alloc()
	rv := objc.Send[CaptureEventSound](instance.ID, objc.Sel("initWithURL:error:"), url, error_)
	rv.Autorelease()
	return rv
}



// The default sound for starting a video recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventSound/beginVideoRecording
func (cc _CaptureEventSoundClass) BeginVideoRecordingSound() CaptureEventSound {
	rv := objc.Send[CaptureEventSound](objc.ID(cc.class), objc.Sel("beginVideoRecordingSound"))
	return rv
}

// The default sound for photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventSound/cameraShutter
func (cc _CaptureEventSoundClass) CameraShutterSound() CaptureEventSound {
	rv := objc.Send[CaptureEventSound](objc.ID(cc.class), objc.Sel("cameraShutterSound"))
	return rv
}

// The default sound for ending a video recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventSound/endVideoRecording
func (cc _CaptureEventSoundClass) EndVideoRecordingSound() CaptureEventSound {
	rv := objc.Send[CaptureEventSound](objc.ID(cc.class), objc.Sel("endVideoRecordingSound"))
	return rv
}

// The default sound for starting a video recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventSound/beginVideoRecording
func (c_ CaptureEventSound) BeginVideoRecordingSound() IAVCaptureEventSound {
	rv := objc.Send[CaptureEventSound](c_.ID, objc.Sel("beginVideoRecordingSound"))
	return rv
}


// The default sound for photo capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventSound/cameraShutter
func (c_ CaptureEventSound) CameraShutterSound() IAVCaptureEventSound {
	rv := objc.Send[CaptureEventSound](c_.ID, objc.Sel("cameraShutterSound"))
	return rv
}


// The default sound for ending a video recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventSound/endVideoRecording
func (c_ CaptureEventSound) EndVideoRecordingSound() IAVCaptureEventSound {
	rv := objc.Send[CaptureEventSound](c_.ID, objc.Sel("endVideoRecordingSound"))
	return rv
}


