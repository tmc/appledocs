// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioFormat] class.
var (
	AudioFormatClass     _AudioFormatClass
	AudioFormatClassOnce sync.Once
)

func getAudioFormatClass() _AudioFormatClass {
	AudioFormatClassOnce.Do(func() {
		AudioFormatClass = _AudioFormatClass{objc.GetClass("AVAudioFormat")}
	})
	return AudioFormatClass
}

type _AudioFormatClass struct {
	class objc.Class
}

// An interface definition for the [AudioFormat] class.
type IAudioFormat interface {
	objectivec.IObject
}

// An object that describes the representation of an audio format.
//
// The class wraps Core Audio’s , and includes convenience initializers and accessors for common formats, including Core Audio’s standard deinterleaved 32-bit floating point format. Instances of this class are immutable.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat
type AudioFormat struct {
	objectivec.Object
}

// AudioFormatFrom constructs a [AudioFormat] from an unsafe.Pointer.
//
// An object that describes the representation of an audio format.
func AudioFormatFrom(ptr unsafe.Pointer) AudioFormat {
	return AudioFormat{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioFormatClass) Alloc() AudioFormat {
	rv := objc.Send[AudioFormat](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioFormatClass) New() AudioFormat {
	rv := objc.Send[AudioFormat](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioFormat) Init() AudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioFormat) Autorelease() AudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioFormat creates a new AudioFormat instance.
func NewAudioFormat() AudioFormat {
	return getAudioFormatClass().New()
}


// The audio format description to use with Core Media APIs.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/formatDescription
func (a_ AudioFormat) FormatDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("formatDescription"))
	return rv
}



