// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioMixingDestination] class.
var (
	AudioMixingDestinationClass     _AudioMixingDestinationClass
	AudioMixingDestinationClassOnce sync.Once
)

func getAudioMixingDestinationClass() _AudioMixingDestinationClass {
	AudioMixingDestinationClassOnce.Do(func() {
		AudioMixingDestinationClass = _AudioMixingDestinationClass{objc.GetClass("AVAudioMixingDestination")}
	})
	return AudioMixingDestinationClass
}

type _AudioMixingDestinationClass struct {
	class objc.Class
}

// An interface definition for the [AudioMixingDestination] class.
type IAudioMixingDestination interface {
	objectivec.IObject
	ConnectionPoint() unsafe.Pointer
	SetConnectionPoint(value unsafe.Pointer)
}

// An object that represents a connection to a mixer node from a node that conforms to the audio mixing protocol.
//
// You can only use a destination instance when a source node provides it. You can’t use it as a standalone instance.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioMixingDestination
type AudioMixingDestination struct {
	objectivec.Object
}

// AudioMixingDestinationFrom constructs a [AudioMixingDestination] from an unsafe.Pointer.
//
// An object that represents a connection to a mixer node from a node that conforms to the audio mixing protocol.
func AudioMixingDestinationFrom(ptr unsafe.Pointer) AudioMixingDestination {
	return AudioMixingDestination{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioMixingDestinationClass) Alloc() AudioMixingDestination {
	rv := objc.Send[AudioMixingDestination](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioMixingDestinationClass) New() AudioMixingDestination {
	rv := objc.Send[AudioMixingDestination](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioMixingDestination) Init() AudioMixingDestination {
	rv := objc.Send[AudioMixingDestination](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioMixingDestination) Autorelease() AudioMixingDestination {
	rv := objc.Send[AudioMixingDestination](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioMixingDestination creates a new AudioMixingDestination instance.
func NewAudioMixingDestination() AudioMixingDestination {
	return getAudioMixingDestinationClass().New()
}


// The underlying mixer connection point.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiomixingdestination/connectionpoint
func (a_ AudioMixingDestination) ConnectionPoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("connectionPoint"))
	return rv
}


// SetConnectionPoint sets the value of the connectionPoint property.
// The underlying mixer connection point.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiomixingdestination/connectionpoint
func (a_ AudioMixingDestination) SetConnectionPoint(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setConnectionPoint:"), value)
}



