// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioSessionRouteDescription] class.
var (
	AudioSessionRouteDescriptionClass     _AudioSessionRouteDescriptionClass
	AudioSessionRouteDescriptionClassOnce sync.Once
)

func getAudioSessionRouteDescriptionClass() _AudioSessionRouteDescriptionClass {
	AudioSessionRouteDescriptionClassOnce.Do(func() {
		AudioSessionRouteDescriptionClass = _AudioSessionRouteDescriptionClass{objc.GetClass("AVAudioSessionRouteDescription")}
	})
	return AudioSessionRouteDescriptionClass
}

type _AudioSessionRouteDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [AudioSessionRouteDescription] class.
type IAudioSessionRouteDescription interface {
	objectivec.IObject
}

// An object that describes the input and output ports associated with a session’s audio route.
//
// You don’t create instances of this class yourself. Instead, you retrieve the current audio route from your app’s object.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionRouteDescription
type AudioSessionRouteDescription struct {
	objectivec.Object
}

// AudioSessionRouteDescriptionFrom constructs a [AudioSessionRouteDescription] from an unsafe.Pointer.
//
// An object that describes the input and output ports associated with a session’s audio route.
func AudioSessionRouteDescriptionFrom(ptr unsafe.Pointer) AudioSessionRouteDescription {
	return AudioSessionRouteDescription{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioSessionRouteDescriptionClass) Alloc() AudioSessionRouteDescription {
	rv := objc.Send[AudioSessionRouteDescription](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioSessionRouteDescriptionClass) New() AudioSessionRouteDescription {
	rv := objc.Send[AudioSessionRouteDescription](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioSessionRouteDescription) Init() AudioSessionRouteDescription {
	rv := objc.Send[AudioSessionRouteDescription](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioSessionRouteDescription) Autorelease() AudioSessionRouteDescription {
	rv := objc.Send[AudioSessionRouteDescription](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioSessionRouteDescription creates a new AudioSessionRouteDescription instance.
func NewAudioSessionRouteDescription() AudioSessionRouteDescription {
	return getAudioSessionRouteDescriptionClass().New()
}


// An array of audio input port descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioSessionRouteDescription/inputs
func (a_ AudioSessionRouteDescription) Inputs() []AudioSessionPortDescription {
	rv := objc.Send[[]AudioSessionPortDescription](a_.ID, objc.Sel("inputs"))
	return rv
}



