// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioUnitComponent] class.
var (
	AudioUnitComponentClass     _AudioUnitComponentClass
	AudioUnitComponentClassOnce sync.Once
)

func getAudioUnitComponentClass() _AudioUnitComponentClass {
	AudioUnitComponentClassOnce.Do(func() {
		AudioUnitComponentClass = _AudioUnitComponentClass{objc.GetClass("AVAudioUnitComponent")}
	})
	return AudioUnitComponentClass
}

type _AudioUnitComponentClass struct {
	class objc.Class
}

// An interface definition for the [AudioUnitComponent] class.
type IAudioUnitComponent interface {
	objectivec.IObject
}

// An object that provides details about an audio unit.
//
// Details can include information such as type, subtype, manufacturer, and location. An can include user tags, which you can query later for display.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent
type AudioUnitComponent struct {
	objectivec.Object
}

// AudioUnitComponentFrom constructs a [AudioUnitComponent] from an unsafe.Pointer.
//
// An object that provides details about an audio unit.
func AudioUnitComponentFrom(ptr unsafe.Pointer) AudioUnitComponent {
	return AudioUnitComponent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioUnitComponentClass) Alloc() AudioUnitComponent {
	rv := objc.Send[AudioUnitComponent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioUnitComponentClass) New() AudioUnitComponent {
	rv := objc.Send[AudioUnitComponent](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitComponent) Init() AudioUnitComponent {
	rv := objc.Send[AudioUnitComponent](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitComponent) Autorelease() AudioUnitComponent {
	rv := objc.Send[AudioUnitComponent](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitComponent creates a new AudioUnitComponent instance.
func NewAudioUnitComponent() AudioUnitComponent {
	return getAudioUnitComponentClass().New()
}




