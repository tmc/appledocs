// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioEnvironmentReverbParameters] class.
var (
	AudioEnvironmentReverbParametersClass     _AudioEnvironmentReverbParametersClass
	AudioEnvironmentReverbParametersClassOnce sync.Once
)

func getAudioEnvironmentReverbParametersClass() _AudioEnvironmentReverbParametersClass {
	AudioEnvironmentReverbParametersClassOnce.Do(func() {
		AudioEnvironmentReverbParametersClass = _AudioEnvironmentReverbParametersClass{objc.GetClass("AVAudioEnvironmentReverbParameters")}
	})
	return AudioEnvironmentReverbParametersClass
}

type _AudioEnvironmentReverbParametersClass struct {
	class objc.Class
}

// An interface definition for the [AudioEnvironmentReverbParameters] class.
type IAudioEnvironmentReverbParameters interface {
	objectivec.IObject
	Enable() bool
	SetEnable(value bool)
	FilterParameters() unsafe.Pointer
	SetFilterParameters(value unsafe.Pointer)
	Level() float32
	SetLevel(value float32)
}

// A class that encapsulates the parameters that you use to control the reverb of the environment node class.
//
// Use reverberation to simulate the acoustic characteristics of an environment. The class has a built-in reverb that describe the space that the listener is in. The reverb has a single filter that sits at the end of the chain. You use this filter to shape the overall sound of the reverb. For instance, select one of the reverb presets to simulate the general space, and then use the filter to brighten or darken the overall sound. You can’t create a standalone instance of . Only an instance vended by a source object is valid, such as an instance.


// A class that encapsulates the parameters that you use to control the reverb of the environment node class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentReverbParameters
type AudioEnvironmentReverbParameters struct {
	objectivec.Object
}

// AudioEnvironmentReverbParametersFrom constructs a [AudioEnvironmentReverbParameters] from an unsafe.Pointer.
//
// A class that encapsulates the parameters that you use to control the reverb of the environment node class.
func AudioEnvironmentReverbParametersFrom(ptr unsafe.Pointer) AudioEnvironmentReverbParameters {
	return AudioEnvironmentReverbParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioEnvironmentReverbParametersClass) Alloc() AudioEnvironmentReverbParameters {
	rv := objc.Send[AudioEnvironmentReverbParameters](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioEnvironmentReverbParametersClass) New() AudioEnvironmentReverbParameters {
	rv := objc.Send[AudioEnvironmentReverbParameters](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioEnvironmentReverbParameters) Init() AudioEnvironmentReverbParameters {
	rv := objc.Send[AudioEnvironmentReverbParameters](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioEnvironmentReverbParameters) Autorelease() AudioEnvironmentReverbParameters {
	rv := objc.Send[AudioEnvironmentReverbParameters](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioEnvironmentReverbParameters creates a new AudioEnvironmentReverbParameters instance.
func NewAudioEnvironmentReverbParameters() AudioEnvironmentReverbParameters {
	return getAudioEnvironmentReverbParametersClass().New()
}



// A Boolean value that indicates whether reverberation is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentreverbparameters/enable
func (a_ AudioEnvironmentReverbParameters) Enable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("enable"))
	return rv
}


// A Boolean value that indicates whether reverberation is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentreverbparameters/enable
func (a_ AudioEnvironmentReverbParameters) SetEnable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEnable:"), value)
}


// A filter that the system applies to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentreverbparameters/filterparameters
func (a_ AudioEnvironmentReverbParameters) FilterParameters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("filterParameters"))
	return rv
}


// A filter that the system applies to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentreverbparameters/filterparameters
func (a_ AudioEnvironmentReverbParameters) SetFilterParameters(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFilterParameters:"), value)
}


// Controls the amount of reverb, in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentreverbparameters/level
func (a_ AudioEnvironmentReverbParameters) Level() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("level"))
	return rv
}


// Controls the amount of reverb, in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentreverbparameters/level
func (a_ AudioEnvironmentReverbParameters) SetLevel(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLevel:"), value)
}



