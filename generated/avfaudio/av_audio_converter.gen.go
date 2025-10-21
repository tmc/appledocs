// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AudioConverter] class.
var (
	AudioConverterClass     _AudioConverterClass
	AudioConverterClassOnce sync.Once
)

func getAudioConverterClass() _AudioConverterClass {
	AudioConverterClassOnce.Do(func() {
		AudioConverterClass = _AudioConverterClass{objc.GetClass("AVAudioConverter")}
	})
	return AudioConverterClass
}

type _AudioConverterClass struct {
	class objc.Class
}

// An interface definition for the [AudioConverter] class.
type IAudioConverter interface {
	objectivec.IObject
	ConvertToBufferErrorWithInputFromBlock(outputBuffer unsafe.Pointer, outError unsafe.Pointer, inputBlock unsafe.Pointer) unsafe.Pointer
}

// An object that converts streams of audio between formats.
//
// The audio converter class transforms audio between file formats and audio encodings. Supported transformations include: PCM float, integer, or bit depth conversions PCM sample rate conversion PCM interleaving and deinterleaving Encoding PCM to compressed formats Decoding compressed formats to PCM A single audio converter instance may perform more than one of the above transformations.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter
type AudioConverter struct {
	objectivec.Object
}

// AudioConverterFrom constructs a [AudioConverter] from an unsafe.Pointer.
//
// An object that converts streams of audio between formats.
func AudioConverterFrom(ptr unsafe.Pointer) AudioConverter {
	return AudioConverter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioConverterClass) Alloc() AudioConverter {
	rv := objc.Send[AudioConverter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioConverterClass) New() AudioConverter {
	rv := objc.Send[AudioConverter](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioConverter) Init() AudioConverter {
	rv := objc.Send[AudioConverter](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioConverter) Autorelease() AudioConverter {
	rv := objc.Send[AudioConverter](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioConverter creates a new AudioConverter instance.
func NewAudioConverter() AudioConverter {
	return getAudioConverterClass().New()
}


// Performs a conversion between audio formats, if the system supports it.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConverter/convert(to:error:withInputFrom:)
func (a_ AudioConverter) ConvertToBufferErrorWithInputFromBlock(outputBuffer unsafe.Pointer, outError unsafe.Pointer, inputBlock unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("convertToBuffer:error:withInputFromBlock:"), outputBuffer, outError, inputBlock)
	return rv
}



