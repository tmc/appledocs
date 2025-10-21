// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CaptureAudioChannel] class.
var (
	CaptureAudioChannelClass     _CaptureAudioChannelClass
	CaptureAudioChannelClassOnce sync.Once
)

func getCaptureAudioChannelClass() _CaptureAudioChannelClass {
	CaptureAudioChannelClassOnce.Do(func() {
		CaptureAudioChannelClass = _CaptureAudioChannelClass{objc.GetClass("AVCaptureAudioChannel")}
	})
	return CaptureAudioChannelClass
}

type _CaptureAudioChannelClass struct {
	class objc.Class
}

// An interface definition for the [CaptureAudioChannel] class.
type ICaptureAudioChannel interface {
	objectivec.IObject
}

// An object that monitors average and peak power levels for an audio channel in a capture connection.
//
// You don’t create instances of this class directly. Instead, an object that connects an audio input to an audio output provides an array of objects, one for each channel of audio available. You can poll for audio levels by iterating through these audio channel objects.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioChannel
type CaptureAudioChannel struct {
	objectivec.Object
}

// CaptureAudioChannelFrom constructs a [CaptureAudioChannel] from an unsafe.Pointer.
//
// An object that monitors average and peak power levels for an audio channel in a capture connection.
func CaptureAudioChannelFrom(ptr unsafe.Pointer) CaptureAudioChannel {
	return CaptureAudioChannel{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureAudioChannelClass) Alloc() CaptureAudioChannel {
	rv := objc.Send[CaptureAudioChannel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureAudioChannelClass) New() CaptureAudioChannel {
	rv := objc.Send[CaptureAudioChannel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureAudioChannel) Init() CaptureAudioChannel {
	rv := objc.Send[CaptureAudioChannel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureAudioChannel) Autorelease() CaptureAudioChannel {
	rv := objc.Send[CaptureAudioChannel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureAudioChannel creates a new CaptureAudioChannel instance.
func NewCaptureAudioChannel() CaptureAudioChannel {
	return getCaptureAudioChannelClass().New()
}




