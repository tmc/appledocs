// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	InputFormatForBus(bus unsafe.Pointer) unsafe.Pointer
	InstallTapOnBusBufferSizeFormatBlock(bus unsafe.Pointer, bufferSize unsafe.Pointer, format unsafe.Pointer, tapBlock unsafe.Pointer)
	NameForInputBus(bus unsafe.Pointer) string
	NameForOutputBus(bus unsafe.Pointer) string
	OutputFormatForBus(bus unsafe.Pointer) unsafe.Pointer
	RemoveTapOnBus(bus unsafe.Pointer)
	Reset()
}

// An object you use for audio generation, processing, or an I/O block.
//
// An object contains instances of audio nodes that you attach, and this base class provides common functionality. Instances of this class don’t provide useful functionality until you attach them to an engine. Nodes have input and output busses that serve as connection points. For example, an effect has one input bus and one output bus, and a mixer has multiple input busses and one output bus. A bus contains a format the framework expresses in terms of sample rate and channel count. Formats must match exactly when making connections between nodes, excluding and .
//
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


// Gets the input format for the bus you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/inputFormat(forBus:)
func (a_ AudioNode) InputFormatForBus(bus unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("inputFormatForBus:"), bus)
	return rv
}

// Installs an audio tap on a bus you specify to record, monitor, and observe the output of the node.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/installTap(onBus:bufferSize:format:block:)
func (a_ AudioNode) InstallTapOnBusBufferSizeFormatBlock(bus unsafe.Pointer, bufferSize unsafe.Pointer, format unsafe.Pointer, tapBlock unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("installTapOnBus:bufferSize:format:block:"), bus, bufferSize, format, tapBlock)
}

// Gets the name of the input bus you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/name(forInputBus:)
func (a_ AudioNode) NameForInputBus(bus unsafe.Pointer) string {
	rv := objc.Send[string](a_.ID, objc.Sel("nameForInputBus:"), bus)
	return rv
}

// Retrieves the name of the output bus you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/name(forOutputBus:)
func (a_ AudioNode) NameForOutputBus(bus unsafe.Pointer) string {
	rv := objc.Send[string](a_.ID, objc.Sel("nameForOutputBus:"), bus)
	return rv
}

// Retrieves the output format for the bus you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/outputFormat(forBus:)
func (a_ AudioNode) OutputFormatForBus(bus unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("outputFormatForBus:"), bus)
	return rv
}

// Removes an audio tap on a bus you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/removeTap(onBus:)
func (a_ AudioNode) RemoveTapOnBus(bus unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeTapOnBus:"), bus)
}

// Clears a unit’s previous processing state.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/reset()
func (a_ AudioNode) Reset() {
	objc.Send[objc.ID](a_.ID, objc.Sel("reset"))
}

// An audio unit object that wraps or underlies the implementation’s audio unit.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/auAudioUnit
func (a_ AudioNode) AUAudioUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("AUAudioUnit"))
	return rv
}

// The audio engine that manages the node, if any.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/engine
func (a_ AudioNode) Engine() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("engine"))
	return rv
}

// The most recent render time.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/lastRenderTime
func (a_ AudioNode) LastRenderTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("lastRenderTime"))
	return rv
}

// The processing latency of the node, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/latency
func (a_ AudioNode) Latency() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](a_.ID, objc.Sel("latency"))
	return rv
}

// The number of input busses for the node.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/numberOfInputs
func (a_ AudioNode) NumberOfInputs() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("numberOfInputs"))
	return rv
}

// The number of output busses for the node.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/numberOfOutputs
func (a_ AudioNode) NumberOfOutputs() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("numberOfOutputs"))
	return rv
}

// The maximum render pipeline latency downstream of the node, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/outputPresentationLatency
func (a_ AudioNode) OutputPresentationLatency() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](a_.ID, objc.Sel("outputPresentationLatency"))
	return rv
}



