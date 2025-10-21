// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASEPushStreamNode] class.
var (
	PHASEPushStreamNodeClass     _PHASEPushStreamNodeClass
	PHASEPushStreamNodeClassOnce sync.Once
)

func getPHASEPushStreamNodeClass() _PHASEPushStreamNodeClass {
	PHASEPushStreamNodeClassOnce.Do(func() {
		PHASEPushStreamNodeClass = _PHASEPushStreamNodeClass{objc.GetClass("PHASEPushStreamNode")}
	})
	return PHASEPushStreamNodeClass
}

type _PHASEPushStreamNodeClass struct {
	class objc.Class
}

// An interface definition for the [PHASEPushStreamNode] class.
type IPHASEPushStreamNode interface {
	IPHASEStreamNode
}

// An audio stream you manage to provide a sound buffer data.
//
// A sound event’s dictionary populates with an instance of this class when PHASE invokes a in your event node tree. Your app provides the audio data that the sound event plays by calling one or more of this class’s buffer-scheduling functions, for example, .
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNode
type PHASEPushStreamNode struct {
	PHASEStreamNode
}

// PHASEPushStreamNodeFrom constructs a [PHASEPushStreamNode] from an unsafe.Pointer.
//
// An audio stream you manage to provide a sound buffer data.
func PHASEPushStreamNodeFrom(ptr unsafe.Pointer) PHASEPushStreamNode {
	return PHASEPushStreamNode{
		PHASEStreamNode: PHASEStreamNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEPushStreamNodeClass) Alloc() PHASEPushStreamNode {
	rv := objc.Send[PHASEPushStreamNode](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEPushStreamNodeClass) New() PHASEPushStreamNode {
	rv := objc.Send[PHASEPushStreamNode](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEPushStreamNode) Init() PHASEPushStreamNode {
	rv := objc.Send[PHASEPushStreamNode](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEPushStreamNode) Autorelease() PHASEPushStreamNode {
	rv := objc.Send[PHASEPushStreamNode](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEPushStreamNode creates a new PHASEPushStreamNode instance.
func NewPHASEPushStreamNode() PHASEPushStreamNode {
	return getPHASEPushStreamNodeClass().New()
}


// The format of the audio stream data.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasepushstreamnode/format
func (p_ PHASEPushStreamNode) Format() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("format"))
	return rv
}


// SetFormat sets the value of the format property.
// The format of the audio stream data.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasepushstreamnode/format
func (p_ PHASEPushStreamNode) SetFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFormat:"), value)
}

// A meta parameter for dynamic loudness control.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasepushstreamnode/gainmetaparameter
func (p_ PHASEPushStreamNode) GainMetaParameter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("gainMetaParameter"))
	return rv
}


// SetGainMetaParameter sets the value of the gainMetaParameter property.
// A meta parameter for dynamic loudness control.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasepushstreamnode/gainmetaparameter
func (p_ PHASEPushStreamNode) SetGainMetaParameter(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGainMetaParameter:"), value)
}

// The audio stream’s output pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasepushstreamnode/mixer
func (p_ PHASEPushStreamNode) Mixer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("mixer"))
	return rv
}


// SetMixer sets the value of the mixer property.
// The audio stream’s output pipeline.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasepushstreamnode/mixer
func (p_ PHASEPushStreamNode) SetMixer(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMixer:"), value)
}

// A meta parameter for dynamic rate control.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasepushstreamnode/ratemetaparameter
func (p_ PHASEPushStreamNode) RateMetaParameter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("rateMetaParameter"))
	return rv
}


// SetRateMetaParameter sets the value of the rateMetaParameter property.
// A meta parameter for dynamic rate control.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasepushstreamnode/ratemetaparameter
func (p_ PHASEPushStreamNode) SetRateMetaParameter(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRateMetaParameter:"), value)
}

// A collection of audio streams for playback.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/pushstreamnodes
func (p_ PHASEPushStreamNode) PushStreamNodes() string {
	rv := objc.Send[string](p_.ID, objc.Sel("pushStreamNodes"))
	return rv
}


// SetPushStreamNodes sets the value of the pushStreamNodes property.
// A collection of audio streams for playback.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesoundevent/pushstreamnodes
func (p_ PHASEPushStreamNode) SetPushStreamNodes(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPushStreamNodes:"), objc.String(value))
}



