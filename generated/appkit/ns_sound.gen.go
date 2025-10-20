// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Sound] class.
var (
	SoundClass     _SoundClass
	SoundClassOnce sync.Once
)

func getSoundClass() _SoundClass {
	SoundClassOnce.Do(func() {
		SoundClass = _SoundClass{objc.GetClass("NSSound")}
	})
	return SoundClass
}

type _SoundClass struct {
	class objc.Class
}

// An interface definition for the [Sound] class.
type ISound interface {
	objectivec.IObject
	Play() bool
}

// A simple interface for loading and playing audio files.
//
// You create a sound object with an audio file or data, which can be in any format that Core Audio supports. Customize the sound by configuring its properties, such as setting its playback volume and looping behavior. Call the sound’s method to begin playback. The system executes this call asynchronously so that it doesn’t interrupt the functioning of your app. If you want to play the system beep sound, use the (Swift) or (Objective-C) function.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound
type Sound struct {
	objectivec.Object
}

// SoundFrom constructs a [Sound] from an unsafe.Pointer.
//
// A simple interface for loading and playing audio files.
func SoundFrom(ptr unsafe.Pointer) Sound {
	return Sound{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SoundClass) Alloc() Sound {
	rv := objc.Send[Sound](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SoundClass) New() Sound {
	rv := objc.Send[Sound](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Sound) Init() Sound {
	rv := objc.Send[Sound](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Sound) Autorelease() Sound {
	rv := objc.Send[Sound](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSound creates a new Sound instance.
func NewSound() Sound {
	return getSoundClass().New()
}


// Provides the list of file types the class understands.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/soundUnfilteredFileTypes
func (sc _SoundClass) SoundUnfilteredFileTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("soundUnfilteredFileTypes"))
	return rv
}

// Initiates audio playback.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSound/play()
func (s_ Sound) Play() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("play"))
	return rv
}



