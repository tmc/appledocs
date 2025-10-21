// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AudioChannelLayout] class.
var (
	AudioChannelLayoutClass     _AudioChannelLayoutClass
	AudioChannelLayoutClassOnce sync.Once
)

func getAudioChannelLayoutClass() _AudioChannelLayoutClass {
	AudioChannelLayoutClassOnce.Do(func() {
		AudioChannelLayoutClass = _AudioChannelLayoutClass{objc.GetClass("AVAudioChannelLayout")}
	})
	return AudioChannelLayoutClass
}

type _AudioChannelLayoutClass struct {
	class objc.Class
}

// An interface definition for the [AudioChannelLayout] class.
type IAudioChannelLayout interface {
	objectivec.IObject
}

// An object that describes the roles of a set of audio channels.
//
// The class is a thin wrapper for Core Audio’s .
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioChannelLayout
type AudioChannelLayout struct {
	objectivec.Object
}

// AudioChannelLayoutFrom constructs a [AudioChannelLayout] from an unsafe.Pointer.
//
// An object that describes the roles of a set of audio channels.
func AudioChannelLayoutFrom(ptr unsafe.Pointer) AudioChannelLayout {
	return AudioChannelLayout{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioChannelLayoutClass) Alloc() AudioChannelLayout {
	rv := objc.Send[AudioChannelLayout](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioChannelLayoutClass) New() AudioChannelLayout {
	rv := objc.Send[AudioChannelLayout](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioChannelLayout) Init() AudioChannelLayout {
	rv := objc.Send[AudioChannelLayout](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioChannelLayout) Autorelease() AudioChannelLayout {
	rv := objc.Send[AudioChannelLayout](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioChannelLayout creates a new AudioChannelLayout instance.
func NewAudioChannelLayout() AudioChannelLayout {
	return getAudioChannelLayoutClass().New()
}




