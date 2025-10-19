// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVSampleBufferAudioRenderer] class.
var aVSampleBufferAudioRendererClass = _AVSampleBufferAudioRendererClass{objc.GetClass("AVSampleBufferAudioRenderer")}

type _AVSampleBufferAudioRendererClass struct {
	class objc.Class
}

// An interface definition for the [AVSampleBufferAudioRenderer] class.
type IAVSampleBufferAudioRenderer interface {
	objectivec.IObject
}

// An object used to decompress audio and play compressed or uncompressed audio. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer

type AVSampleBufferAudioRenderer struct {
	objectivec.Object
}

// AVSampleBufferAudioRendererFrom constructs a [AVSampleBufferAudioRenderer] from an unsafe.Pointer.
//
// An object used to decompress audio and play compressed or uncompressed audio.
func AVSampleBufferAudioRendererFrom(ptr unsafe.Pointer) AVSampleBufferAudioRenderer {
	return AVSampleBufferAudioRenderer{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AVSampleBufferAudioRendererClass) Alloc() AVSampleBufferAudioRenderer {
	rv := objc.Send[AVSampleBufferAudioRenderer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVSampleBufferAudioRendererClass) New() AVSampleBufferAudioRenderer {
	rv := objc.Send[AVSampleBufferAudioRenderer](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVSampleBufferAudioRenderer) Init() AVSampleBufferAudioRenderer {
	rv := objc.Send[AVSampleBufferAudioRenderer](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVSampleBufferAudioRenderer) Autorelease() AVSampleBufferAudioRenderer {
	rv := objc.Send[AVSampleBufferAudioRenderer](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSampleBufferAudioRenderer creates a new AVSampleBufferAudioRenderer instance.
func NewAVSampleBufferAudioRenderer() AVSampleBufferAudioRenderer {
	return aVSampleBufferAudioRendererClass.New()
}




