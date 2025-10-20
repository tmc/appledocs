// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SampleBufferAudioRenderer] class.
var (
	SampleBufferAudioRendererClass     _SampleBufferAudioRendererClass
	SampleBufferAudioRendererClassOnce sync.Once
)

func getSampleBufferAudioRendererClass() _SampleBufferAudioRendererClass {
	SampleBufferAudioRendererClassOnce.Do(func() {
		SampleBufferAudioRendererClass = _SampleBufferAudioRendererClass{objc.GetClass("AVSampleBufferAudioRenderer")}
	})
	return SampleBufferAudioRendererClass
}

type _SampleBufferAudioRendererClass struct {
	class objc.Class
}

// An interface definition for the [SampleBufferAudioRenderer] class.
type ISampleBufferAudioRenderer interface {
	objectivec.IObject
}

// An object used to decompress audio and play compressed or uncompressed audio.
//
// You must add an instance of this class to an before queuing the first sample buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer
type SampleBufferAudioRenderer struct {
	objectivec.Object
}

// SampleBufferAudioRendererFrom constructs a [SampleBufferAudioRenderer] from an unsafe.Pointer.
//
// An object used to decompress audio and play compressed or uncompressed audio.
func SampleBufferAudioRendererFrom(ptr unsafe.Pointer) SampleBufferAudioRenderer {
	return SampleBufferAudioRenderer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SampleBufferAudioRendererClass) Alloc() SampleBufferAudioRenderer {
	rv := objc.Send[SampleBufferAudioRenderer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SampleBufferAudioRendererClass) New() SampleBufferAudioRenderer {
	rv := objc.Send[SampleBufferAudioRenderer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SampleBufferAudioRenderer) Init() SampleBufferAudioRenderer {
	rv := objc.Send[SampleBufferAudioRenderer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SampleBufferAudioRenderer) Autorelease() SampleBufferAudioRenderer {
	rv := objc.Send[SampleBufferAudioRenderer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSampleBufferAudioRenderer creates a new SampleBufferAudioRenderer instance.
func NewSampleBufferAudioRenderer() SampleBufferAudioRenderer {
	return getSampleBufferAudioRendererClass().New()
}


// The unique identifier of the output device used to play audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/audioOutputDeviceUniqueID
func (s_ SampleBufferAudioRenderer) AudioOutputDeviceUniqueID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("audioOutputDeviceUniqueID"))
	return rv
}


// SetAudioOutputDeviceUniqueID sets the value of the audioOutputDeviceUniqueID property.
// The unique identifier of the output device used to play audio.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/audioOutputDeviceUniqueID
func (s_ SampleBufferAudioRenderer) SetAudioOutputDeviceUniqueID(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAudioOutputDeviceUniqueID:"), value)
}


