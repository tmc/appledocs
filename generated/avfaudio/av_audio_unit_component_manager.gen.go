// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioUnitComponentManager] class.
var (
	AudioUnitComponentManagerClass     _AudioUnitComponentManagerClass
	AudioUnitComponentManagerClassOnce sync.Once
)

func getAudioUnitComponentManagerClass() _AudioUnitComponentManagerClass {
	AudioUnitComponentManagerClassOnce.Do(func() {
		AudioUnitComponentManagerClass = _AudioUnitComponentManagerClass{objc.GetClass("AVAudioUnitComponentManager")}
	})
	return AudioUnitComponentManagerClass
}

type _AudioUnitComponentManagerClass struct {
	class objc.Class
}

// An interface definition for the [AudioUnitComponentManager] class.
type IAudioUnitComponentManager interface {
	objectivec.IObject
	StandardLocalizedTagNames() string
	SetStandardLocalizedTagNames(value string)
	TagNames() string
	SetTagNames(value string)
}

// An object that provides a way to search and query audio components that the system registers.
//
// The component manager has methods to find various information about the audio components without opening them. Currently, you can only search audio components that are audio units. The class supports system tags and arbitrary user tags. You can tag each audio unit as part of its definition. Audio unit hosts, such as Logic or GarageBand, can present groupings of audio units according to the tags. You can search for audio units in the following ways: Using a instance that contains search strings for tags or descriptions Using a block to match on a custom criteria Using an


// An object that provides a way to search and query audio components that the system registers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager

type AudioUnitComponentManager struct {
	objectivec.Object
}

// AudioUnitComponentManagerFrom constructs a [AudioUnitComponentManager] from an unsafe.Pointer.
//
// An object that provides a way to search and query audio components that the system registers.
func AudioUnitComponentManagerFrom(ptr unsafe.Pointer) AudioUnitComponentManager {
	return AudioUnitComponentManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioUnitComponentManagerClass) Alloc() AudioUnitComponentManager {
	rv := objc.Send[AudioUnitComponentManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioUnitComponentManagerClass) New() AudioUnitComponentManager {
	rv := objc.Send[AudioUnitComponentManager](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitComponentManager) Init() AudioUnitComponentManager {
	rv := objc.Send[AudioUnitComponentManager](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitComponentManager) Autorelease() AudioUnitComponentManager {
	rv := objc.Send[AudioUnitComponentManager](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitComponentManager creates a new AudioUnitComponentManager instance.
func NewAudioUnitComponentManager() AudioUnitComponentManager {
	return getAudioUnitComponentManagerClass().New()
}



// An array of the localized standard system tags the audio units define.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponentmanager/standardlocalizedtagnames

func (a_ AudioUnitComponentManager) StandardLocalizedTagNames() string {
	rv := objc.Send[string](a_.ID, objc.Sel("standardLocalizedTagNames"))
	return rv
}


// An array of the localized standard system tags the audio units define.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponentmanager/standardlocalizedtagnames

func (a_ AudioUnitComponentManager) SetStandardLocalizedTagNames(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStandardLocalizedTagNames:"), objc.String(value))
}


// An array of all tags the audio unit associates with the current user, and the system tags the audio units define.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponentmanager/tagnames

func (a_ AudioUnitComponentManager) TagNames() string {
	rv := objc.Send[string](a_.ID, objc.Sel("tagNames"))
	return rv
}


// An array of all tags the audio unit associates with the current user, and the system tags the audio units define.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudiounitcomponentmanager/tagnames

func (a_ AudioUnitComponentManager) SetTagNames(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTagNames:"), objc.String(value))
}



