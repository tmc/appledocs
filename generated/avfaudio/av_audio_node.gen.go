// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioNode */


/* debug [class_header]: Header for AVAudioNode */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioNode */
// An interface definition for the [AudioNode] class.
type IAudioNode interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioNode */
	// properties:
	AUAudioUnit() IAudioUnit
	Engine() IAVAudioEngine
	LastRenderTime() IAVAudioTime
	Latency() float64
	NumberOfInputs() uint
	NumberOfOutputs() uint
	OutputPresentationLatency() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioNode */
	// methods:
	InputFormatForBus(bus AudioNodeBus /* typedef */) IAudioFormat
	InstallTapOnBusBufferSizeFormatBlock(bus AudioNodeBus /* typedef */, bufferSize AudioFrameCount /* typedef */, format IAVAudioFormat, tapBlock AudioNodeTapBlock /* not a class type */)
	NameForInputBus(bus AudioNodeBus /* typedef */) foundation.String
	NameForOutputBus(bus AudioNodeBus /* typedef */) foundation.String
	OutputFormatForBus(bus AudioNodeBus /* typedef */) IAudioFormat
	RemoveTapOnBus(bus AudioNodeBus /* typedef */)
	Reset()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioNode */
// Alloc allocates a new instance without initialization.
func (ac _AudioNodeClass) Alloc() AudioNode {
	rv := objc.Send[AudioNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioNode */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioNode */

// Gets the input format for the bus you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/inputFormat(forBus:)
func (a_ AudioNode) InputFormatForBus(bus AudioNodeBus /* typedef */) IAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("inputFormatForBus:"), bus)
	return rv
}/* debug [instance_methods/method]: InputFormatForBus */


// Installs an audio tap on a bus you specify to record, monitor, and observe the output of the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/installTap(onBus:bufferSize:format:block:)
func (a_ AudioNode) InstallTapOnBusBufferSizeFormatBlock(bus AudioNodeBus /* typedef */, bufferSize AudioFrameCount /* typedef */, format IAVAudioFormat, tapBlock AudioNodeTapBlock /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("installTapOnBus:bufferSize:format:block:"), bus, bufferSize, format, tapBlock)
}/* debug [instance_methods/method]: InstallTapOnBusBufferSizeFormatBlock */


// Gets the name of the input bus you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/name(forInputBus:)
func (a_ AudioNode) NameForInputBus(bus AudioNodeBus /* typedef */) foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("nameForInputBus:"), bus)
	return rv
}/* debug [instance_methods/method]: NameForInputBus */


// Retrieves the name of the output bus you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/name(forOutputBus:)
func (a_ AudioNode) NameForOutputBus(bus AudioNodeBus /* typedef */) foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("nameForOutputBus:"), bus)
	return rv
}/* debug [instance_methods/method]: NameForOutputBus */


// Retrieves the output format for the bus you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/outputFormat(forBus:)
func (a_ AudioNode) OutputFormatForBus(bus AudioNodeBus /* typedef */) IAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("outputFormatForBus:"), bus)
	return rv
}/* debug [instance_methods/method]: OutputFormatForBus */


// Removes an audio tap on a bus you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/removeTap(onBus:)
func (a_ AudioNode) RemoveTapOnBus(bus AudioNodeBus /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeTapOnBus:"), bus)
}/* debug [instance_methods/method]: RemoveTapOnBus */


// Clears a unit’s previous processing state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/reset()
func (a_ AudioNode) Reset() {
	objc.Send[objc.ID](a_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioNode */

// An audio unit object that wraps or underlies the implementation’s audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/auAudioUnit
func (a_ AudioNode) AUAudioUnit() IAudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("AUAudioUnit"))
	return rv
}/* debug [instance_properties/getter]: AUAudioUnit */


// The audio engine that manages the node, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/engine
func (a_ AudioNode) Engine() IAVAudioEngine {
	rv := objc.Send[AudioEngine](a_.ID, objc.Sel("engine"))
	return rv
}/* debug [instance_properties/getter]: engine */


// The most recent render time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/lastRenderTime
func (a_ AudioNode) LastRenderTime() IAVAudioTime {
	rv := objc.Send[AudioTime](a_.ID, objc.Sel("lastRenderTime"))
	return rv
}/* debug [instance_properties/getter]: lastRenderTime */


// The processing latency of the node, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/latency
func (a_ AudioNode) Latency() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("latency"))
	return rv
}/* debug [instance_properties/getter]: latency */


// The number of input busses for the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/numberOfInputs
func (a_ AudioNode) NumberOfInputs() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("numberOfInputs"))
	return rv
}/* debug [instance_properties/getter]: numberOfInputs */


// The number of output busses for the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/numberOfOutputs
func (a_ AudioNode) NumberOfOutputs() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("numberOfOutputs"))
	return rv
}/* debug [instance_properties/getter]: numberOfOutputs */


// The maximum render pipeline latency downstream of the node, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/outputPresentationLatency
func (a_ AudioNode) OutputPresentationLatency() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("outputPresentationLatency"))
	return rv
}/* debug [instance_properties/getter]: outputPresentationLatency */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioNode */



