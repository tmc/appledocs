// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioBuffer */


/* debug [class_header]: Header for AVAudioBuffer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioBuffer */
// An interface definition for the [AudioBuffer] class.
type IAudioBuffer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioBuffer */
	// properties:
	AudioBufferList() objc.IObject
	Format() IAVAudioFormat
	MutableAudioBufferList() objc.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioBuffer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioBuffer */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioBuffer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioBuffer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioBuffer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioBuffer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioBuffer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioBuffer */

// The buffer’s underlying audio buffer list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer/audioBufferList
func (a_ AudioBuffer) AudioBufferList() objc.IObject {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("audioBufferList"))
	return rv
}/* debug [instance_properties/getter]: audioBufferList */


// The format of the audio in the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer/format
func (a_ AudioBuffer) Format() IAVAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("format"))
	return rv
}/* debug [instance_properties/getter]: format */


// A mutable version of the buffer’s underlying audio buffer list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer/mutableAudioBufferList
func (a_ AudioBuffer) MutableAudioBufferList() objc.IObject {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("mutableAudioBufferList"))
	return rv
}/* debug [instance_properties/getter]: mutableAudioBufferList */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioBuffer */



