// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AudioUnitV2Bridge] class.
var (
	AudioUnitV2BridgeClass     _AudioUnitV2BridgeClass
	AudioUnitV2BridgeClassOnce sync.Once
)

func getAudioUnitV2BridgeClass() _AudioUnitV2BridgeClass {
	AudioUnitV2BridgeClassOnce.Do(func() {
		AudioUnitV2BridgeClass = _AudioUnitV2BridgeClass{objc.GetClass("AUAudioUnitV2Bridge")}
	})
	return AudioUnitV2BridgeClass
}

type _AudioUnitV2BridgeClass struct {
	class objc.Class
}

// An interface definition for the [AudioUnitV2Bridge] class.
type IAudioUnitV2Bridge interface {
	IAudioUnit
}

// A class that wraps a version 2 audio unit as version 3 audio unit.
//
// A version 3 audio unit may subclass the class. If so, the audio unit’s component description should refer to a registered component with a version 2 implementation by using a factory function. The bridge will instantiate the version 2 audio unit via the factory function and communicate with it using version 2 audio unit APIs. Hosts should not access this class; it will be instantiated if needed when creating an audio unit.


// A class that wraps a version 2 audio unit as version 3 audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitV2Bridge
type AudioUnitV2Bridge struct {
	AudioUnit
}

// AudioUnitV2BridgeFrom constructs a [AudioUnitV2Bridge] from an unsafe.Pointer.
//
// A class that wraps a version 2 audio unit as version 3 audio unit.
func AudioUnitV2BridgeFrom(ptr unsafe.Pointer) AudioUnitV2Bridge {
	return AudioUnitV2Bridge{
		AudioUnit: AudioUnitFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioUnitV2BridgeClass) Alloc() AudioUnitV2Bridge {
	rv := objc.Send[AudioUnitV2Bridge](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioUnitV2BridgeClass) New() AudioUnitV2Bridge {
	rv := objc.Send[AudioUnitV2Bridge](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitV2Bridge) Init() AudioUnitV2Bridge {
	rv := objc.Send[AudioUnitV2Bridge](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitV2Bridge) Autorelease() AudioUnitV2Bridge {
	rv := objc.Send[AudioUnitV2Bridge](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitV2Bridge creates a new AudioUnitV2Bridge instance.
func NewAudioUnitV2Bridge() AudioUnitV2Bridge {
	return getAudioUnitV2BridgeClass().New()
}




