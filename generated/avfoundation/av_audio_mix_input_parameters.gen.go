// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVAudioMixInputParameters] class.
var aVAudioMixInputParametersClass = _AVAudioMixInputParametersClass{objc.GetClass("AVAudioMixInputParameters")}

type _AVAudioMixInputParametersClass struct {
	class objc.Class
}

// An interface definition for the [AVAudioMixInputParameters] class.
type IAVAudioMixInputParameters interface {
	objectivec.IObject
}

// An object that represents the parameters that you apply when adding an audio track to a mix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAudioMixInputParameters

type AVAudioMixInputParameters struct {
	objectivec.Object
}

// AVAudioMixInputParametersFrom constructs a [AVAudioMixInputParameters] from an unsafe.Pointer.
//
// An object that represents the parameters that you apply when adding an audio track to a mix.
func AVAudioMixInputParametersFrom(ptr unsafe.Pointer) AVAudioMixInputParameters {
	return AVAudioMixInputParameters{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AVAudioMixInputParametersClass) Alloc() AVAudioMixInputParameters {
	rv := objc.Send[AVAudioMixInputParameters](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVAudioMixInputParametersClass) New() AVAudioMixInputParameters {
	rv := objc.Send[AVAudioMixInputParameters](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVAudioMixInputParameters) Init() AVAudioMixInputParameters {
	rv := objc.Send[AVAudioMixInputParameters](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVAudioMixInputParameters) Autorelease() AVAudioMixInputParameters {
	rv := objc.Send[AVAudioMixInputParameters](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioMixInputParameters creates a new AVAudioMixInputParameters instance.
func NewAVAudioMixInputParameters() AVAudioMixInputParameters {
	return aVAudioMixInputParametersClass.New()
}




