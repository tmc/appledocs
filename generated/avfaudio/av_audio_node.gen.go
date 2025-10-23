// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioNode] class.
var (
	AudioNodeClass     _AudioNodeClass
	AudioNodeClassOnce sync.Once
)

func getAudioNodeClass() _AudioNodeClass {
	AudioNodeClassOnce.Do(func() {
		AudioNodeClass = _AudioNodeClass{objc.GetClass("AVAudioNode")}
	})
	return AudioNodeClass
}

type _AudioNodeClass struct {
	class objc.Class
}

// An interface definition for the [AudioNode] class.
type IAudioNode interface {
	objectivec.IObject
	AuAudioUnit() IAudioUnit
	SetAuAudioUnit(value IAudioUnit)
	Engine() IAVAudioEngine
	SetEngine(value IAVAudioEngine)
	LastRenderTime() IAVAudioTime
	SetLastRenderTime(value IAVAudioTime)
	Latency() unsafe.Pointer
	SetLatency(value unsafe.Pointer)
	NumberOfInputs() int
	SetNumberOfInputs(value int)
	NumberOfOutputs() int
	SetNumberOfOutputs(value int)
	OutputPresentationLatency() unsafe.Pointer
	SetOutputPresentationLatency(value unsafe.Pointer)
}

// An object you use for audio generation, processing, or an I/O block.
//
// An object contains instances of audio nodes that you attach, and this base class provides common functionality. Instances of this class don’t provide useful functionality until you attach them to an engine. Nodes have input and output busses that serve as connection points. For example, an effect has one input bus and one output bus, and a mixer has multiple input busses and one output bus. A bus contains a format the framework expresses in terms of sample rate and channel count. Formats must match exactly when making connections between nodes, excluding and .


// An object you use for audio generation, processing, or an I/O block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode
type AudioNode struct {
	objectivec.Object
}

// AudioNodeFrom constructs a [AudioNode] from an unsafe.Pointer.
//
// An object you use for audio generation, processing, or an I/O block.
func AudioNodeFrom(ptr unsafe.Pointer) AudioNode {
	return AudioNode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioNodeClass) Alloc() AudioNode {
	rv := objc.Send[AudioNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioNodeClass) New() AudioNode {
	rv := objc.Send[AudioNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioNode) Init() AudioNode {
	rv := objc.Send[AudioNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioNode) Autorelease() AudioNode {
	rv := objc.Send[AudioNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioNode creates a new AudioNode instance.
func NewAudioNode() AudioNode {
	return getAudioNodeClass().New()
}



// An audio unit object that wraps or underlies the implementation’s audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/auaudiounit
func (a_ AudioNode) AuAudioUnit() IAudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("auAudioUnit"))
	return rv
}


// An audio unit object that wraps or underlies the implementation’s audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/auaudiounit
func (a_ AudioNode) SetAuAudioUnit(value IAudioUnit) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAuAudioUnit:"), value)
}


// The audio engine that manages the node, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/engine
func (a_ AudioNode) Engine() IAVAudioEngine {
	rv := objc.Send[AudioEngine](a_.ID, objc.Sel("engine"))
	return rv
}


// The audio engine that manages the node, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/engine
func (a_ AudioNode) SetEngine(value IAVAudioEngine) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEngine:"), value)
}


// The most recent render time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/lastrendertime
func (a_ AudioNode) LastRenderTime() IAVAudioTime {
	rv := objc.Send[AudioTime](a_.ID, objc.Sel("lastRenderTime"))
	return rv
}


// The most recent render time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/lastrendertime
func (a_ AudioNode) SetLastRenderTime(value IAVAudioTime) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLastRenderTime:"), value)
}


// The processing latency of the node, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/latency
func (a_ AudioNode) Latency() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("latency"))
	return rv
}


// The processing latency of the node, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/latency
func (a_ AudioNode) SetLatency(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLatency:"), value)
}


// The number of input busses for the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/numberofinputs
func (a_ AudioNode) NumberOfInputs() int {
	rv := objc.Send[int](a_.ID, objc.Sel("numberOfInputs"))
	return rv
}


// The number of input busses for the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/numberofinputs
func (a_ AudioNode) SetNumberOfInputs(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNumberOfInputs:"), value)
}


// The number of output busses for the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/numberofoutputs
func (a_ AudioNode) NumberOfOutputs() int {
	rv := objc.Send[int](a_.ID, objc.Sel("numberOfOutputs"))
	return rv
}


// The number of output busses for the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/numberofoutputs
func (a_ AudioNode) SetNumberOfOutputs(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNumberOfOutputs:"), value)
}


// The maximum render pipeline latency downstream of the node, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/outputpresentationlatency
func (a_ AudioNode) OutputPresentationLatency() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("outputPresentationLatency"))
	return rv
}


// The maximum render pipeline latency downstream of the node, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/outputpresentationlatency
func (a_ AudioNode) SetOutputPresentationLatency(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputPresentationLatency:"), value)
}



