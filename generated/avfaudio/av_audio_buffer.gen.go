// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AudioBuffer] class.
var (
	AudioBufferClass     _AudioBufferClass
	AudioBufferClassOnce sync.Once
)

func getAudioBufferClass() _AudioBufferClass {
	AudioBufferClassOnce.Do(func() {
		AudioBufferClass = _AudioBufferClass{objc.GetClass("AVAudioBuffer")}
	})
	return AudioBufferClass
}

type _AudioBufferClass struct {
	class objc.Class
}





// An interface definition for the [AudioBuffer] class.
type IAudioBuffer interface {
	objectivec.IObject
	

	// properties:
	AudioBufferList() objc.IObject
	Format() IAVAudioFormat
	MutableAudioBufferList() objc.IObject


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AudioBufferClass) Alloc() AudioBuffer {
	rv := objc.Send[AudioBuffer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioBufferClass) New() AudioBuffer {
	rv := objc.Send[AudioBuffer](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioBuffer) Init() AudioBuffer {
	rv := objc.Send[AudioBuffer](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioBuffer) Autorelease() AudioBuffer {
	rv := objc.Send[AudioBuffer](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioBuffer creates a new AudioBuffer instance.
func NewAudioBuffer() AudioBuffer {
	return getAudioBufferClass().New()
}





// An object that represents a buffer of audio data with a format.


// An object that represents a buffer of audio data with a format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer
type AudioBuffer struct {
	objectivec.Object
}

// AudioBufferFrom constructs a [AudioBuffer] from an unsafe.Pointer.
//
// An object that represents a buffer of audio data with a format.
func AudioBufferFrom(ptr unsafe.Pointer) AudioBuffer {
	return AudioBuffer{objectivec.Object{objc.ID(ptr)}}
}

























// The buffer’s underlying audio buffer list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer/audioBufferList
func (a_ AudioBuffer) AudioBufferList() objc.IObject {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("audioBufferList"))
	return rv
}


// The format of the audio in the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer/format
func (a_ AudioBuffer) Format() IAVAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("format"))
	return rv
}


// A mutable version of the buffer’s underlying audio buffer list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer/mutableAudioBufferList
func (a_ AudioBuffer) MutableAudioBufferList() objc.IObject {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("mutableAudioBufferList"))
	return rv
}








