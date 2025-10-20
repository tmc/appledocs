// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioSessionPortDescription] class.
var (
	AudioSessionPortDescriptionClass     _AudioSessionPortDescriptionClass
	AudioSessionPortDescriptionClassOnce sync.Once
)

func getAudioSessionPortDescriptionClass() _AudioSessionPortDescriptionClass {
	AudioSessionPortDescriptionClassOnce.Do(func() {
		AudioSessionPortDescriptionClass = _AudioSessionPortDescriptionClass{objc.GetClass("AVAudioSessionPortDescription")}
	})
	return AudioSessionPortDescriptionClass
}

type _AudioSessionPortDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [AudioSessionPortDescription] class.
type IAudioSessionPortDescription interface {
	objectivec.IObject
}

// Information about the capabilities of the port and the hardware channels it supports.
//
// A port description object describes a single input or output port associated with an audio route. Examples of audio ports include a device’s built-in speaker, a microphone on a wired headset, and a Bluetooth device supporting the Advanced Audio Distribution Profile (A2DP). You can query the audio session’s property to get information about the active set of input and output ports. To change the current audio routing, call the method. For example, on a device with a wired headset attached, the audio session’s array may contain two port descriptions: one for the headset microphone and one for the device’s built-in microphone. You can use the audio session’s method to select the headset or built-in microphone for audio input.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortDescription
type AudioSessionPortDescription struct {
	objectivec.Object
}

// AudioSessionPortDescriptionFrom constructs a [AudioSessionPortDescription] from an unsafe.Pointer.
//
// Information about the capabilities of the port and the hardware channels it supports.
func AudioSessionPortDescriptionFrom(ptr unsafe.Pointer) AudioSessionPortDescription {
	return AudioSessionPortDescription{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioSessionPortDescriptionClass) Alloc() AudioSessionPortDescription {
	rv := objc.Send[AudioSessionPortDescription](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioSessionPortDescriptionClass) New() AudioSessionPortDescription {
	rv := objc.Send[AudioSessionPortDescription](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSessionPortDescription) Init() AudioSessionPortDescription {
	rv := objc.Send[AudioSessionPortDescription](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSessionPortDescription) Autorelease() AudioSessionPortDescription {
	rv := objc.Send[AudioSessionPortDescription](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSessionPortDescription creates a new AudioSessionPortDescription instance.
func NewAudioSessionPortDescription() AudioSessionPortDescription {
	return getAudioSessionPortDescriptionClass().New()
}


// The preferred audio data source for the port.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionPortDescription/preferredDataSource
func (a_ AudioSessionPortDescription) PreferredDataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("preferredDataSource"))
	return rv
}



