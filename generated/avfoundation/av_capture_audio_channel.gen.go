// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaptureAudioChannel] class.
var (
	aVCaptureAudioChannelClass     _AVCaptureAudioChannelClass
	aVCaptureAudioChannelClassOnce sync.Once
)

func getAVCaptureAudioChannelClass() _AVCaptureAudioChannelClass {
	aVCaptureAudioChannelClassOnce.Do(func() {
		aVCaptureAudioChannelClass = _AVCaptureAudioChannelClass{objc.GetClass("AVCaptureAudioChannel")}
	})
	return aVCaptureAudioChannelClass
}

type _AVCaptureAudioChannelClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureAudioChannel] class.
type IAVCaptureAudioChannel interface {
	objectivec.IObject
}

// An object that monitors average and peak power levels for an audio channel in a capture connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioChannel
type AVCaptureAudioChannel struct {
	objectivec.Object
}

// AVCaptureAudioChannelFrom constructs a [AVCaptureAudioChannel] from an unsafe.Pointer.
//
// An object that monitors average and peak power levels for an audio channel in a capture connection.
func AVCaptureAudioChannelFrom(ptr unsafe.Pointer) AVCaptureAudioChannel {
	return AVCaptureAudioChannel{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVCaptureAudioChannelClass) Alloc() AVCaptureAudioChannel {
	rv := objc.Send[AVCaptureAudioChannel](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVCaptureAudioChannelClass) New() AVCaptureAudioChannel {
	rv := objc.Send[AVCaptureAudioChannel](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureAudioChannel) Init() AVCaptureAudioChannel {
	rv := objc.Send[AVCaptureAudioChannel](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureAudioChannel) Autorelease() AVCaptureAudioChannel {
	rv := objc.Send[AVCaptureAudioChannel](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureAudioChannel creates a new AVCaptureAudioChannel instance.
func NewAVCaptureAudioChannel() AVCaptureAudioChannel {
	return getAVCaptureAudioChannelClass().New()
}




