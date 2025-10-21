// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AudioPCMBuffer] class.
var (
	AudioPCMBufferClass     _AudioPCMBufferClass
	AudioPCMBufferClassOnce sync.Once
)

func getAudioPCMBufferClass() _AudioPCMBufferClass {
	AudioPCMBufferClassOnce.Do(func() {
		AudioPCMBufferClass = _AudioPCMBufferClass{objc.GetClass("AVAudioPCMBuffer")}
	})
	return AudioPCMBufferClass
}

type _AudioPCMBufferClass struct {
	class objc.Class
}

// An interface definition for the [AudioPCMBuffer] class.
type IAudioPCMBuffer interface {
	IAudioBuffer
}

// An object that represents an audio buffer you use with PCM audio formats.
//
// The PCM buffer class provides methods that are useful for manipulating buffers of audio in PCM format.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer
type AudioPCMBuffer struct {
	AudioBuffer
}

// AudioPCMBufferFrom constructs a [AudioPCMBuffer] from an unsafe.Pointer.
//
// An object that represents an audio buffer you use with PCM audio formats.
func AudioPCMBufferFrom(ptr unsafe.Pointer) AudioPCMBuffer {
	return AudioPCMBuffer{
		AudioBuffer: AudioBufferFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioPCMBufferClass) Alloc() AudioPCMBuffer {
	rv := objc.Send[AudioPCMBuffer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioPCMBufferClass) New() AudioPCMBuffer {
	rv := objc.Send[AudioPCMBuffer](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioPCMBuffer) Init() AudioPCMBuffer {
	rv := objc.Send[AudioPCMBuffer](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioPCMBuffer) Autorelease() AudioPCMBuffer {
	rv := objc.Send[AudioPCMBuffer](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioPCMBuffer creates a new AudioPCMBuffer instance.
func NewAudioPCMBuffer() AudioPCMBuffer {
	return getAudioPCMBufferClass().New()
}




// Creates a PCM audio buffer instance without copying samples, for PCM audio data, with a specified buffer list and a deallocator closure.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/init(pcmFormat:bufferListNoCopy:deallocator:)
func NewAudioPCMBufferWithPCMFormatBufferListNoCopyDeallocator(format AVAudioFormat, bufferList unsafe.Pointer, deallocator unsafe.Pointer) AudioPCMBuffer {
	instance := getAudioPCMBufferClass().Alloc()
	rv := objc.Send[AudioPCMBuffer](instance.ID, objc.Sel("initWithPCMFormat:bufferListNoCopy:deallocator:"), format, bufferList, deallocator)
	rv.Autorelease()
	return rv
}



// Creates a PCM audio buffer instance for PCM audio data.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/init(pcmFormat:frameCapacity:)
func NewAudioPCMBufferWithPCMFormatFrameCapacity(format AVAudioFormat, frameCapacity IAudioFrameCount) AudioPCMBuffer {
	instance := getAudioPCMBufferClass().Alloc()
	rv := objc.Send[AudioPCMBuffer](instance.ID, objc.Sel("initWithPCMFormat:frameCapacity:"), format, frameCapacity)
	rv.Autorelease()
	return rv
}


// The buffer’s audio samples as floating point values.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/floatChannelData
func (a_ AudioPCMBuffer) FloatChannelData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("floatChannelData"))
	return rv
}

// The buffer’s capacity, in audio sample frames.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/frameCapacity
func (a_ AudioPCMBuffer) FrameCapacity() AudioFrameCount {
	rv := objc.Send[AudioFrameCount](a_.ID, objc.Sel("frameCapacity"))
	return rv
}

// The current number of valid sample frames in the buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/frameLength
func (a_ AudioPCMBuffer) FrameLength() AudioFrameCount {
	rv := objc.Send[AudioFrameCount](a_.ID, objc.Sel("frameLength"))
	return rv
}


// SetFrameLength sets the value of the frameLength property.
// The current number of valid sample frames in the buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/frameLength
func (a_ AudioPCMBuffer) SetFrameLength(value IAudioFrameCount) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFrameLength:"), value)
}

// The buffer’s 16-bit integer audio samples.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/int16ChannelData
func (a_ AudioPCMBuffer) Int16ChannelData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("int16ChannelData"))
	return rv
}

// The buffer’s 32-bit integer audio samples.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/int32ChannelData
func (a_ AudioPCMBuffer) Int32ChannelData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("int32ChannelData"))
	return rv
}

// The buffer’s number of interleaved channels.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/stride
func (a_ AudioPCMBuffer) Stride() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("stride"))
	return rv
}


