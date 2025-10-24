// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AudioUnit] class.
var (
	AudioUnitClass     _AudioUnitClass
	AudioUnitClassOnce sync.Once
)

func getAudioUnitClass() _AudioUnitClass {
	AudioUnitClassOnce.Do(func() {
		AudioUnitClass = _AudioUnitClass{objc.GetClass("AUAudioUnit")}
	})
	return AudioUnitClass
}

type _AudioUnitClass struct {
	class objc.Class
}





// An interface definition for the [AudioUnit] class.
type IAudioUnit interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AudioUnitClass) Alloc() AudioUnit {
	rv := objc.Send[AudioUnit](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitClass) New() AudioUnit {
	rv := objc.Send[AudioUnit](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnit) Init() AudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnit) Autorelease() AudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnit creates a new AudioUnit instance.
func NewAudioUnit() AudioUnit {
	return getAudioUnitClass().New()
}





// A parent class referenced by other AVFAudio classes.


// A parent class referenced by other AVFAudio classes. [Full Topic]
type AudioUnit struct {
	objectivec.Object
}

// AudioUnitFrom constructs a [AudioUnit] from an unsafe.Pointer.
//
// A parent class referenced by other AVFAudio classes.
func AudioUnitFrom(ptr unsafe.Pointer) AudioUnit {
	return AudioUnit{objectivec.Object{objc.ID(ptr)}}
}































