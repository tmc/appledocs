// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioEngine */


/* debug [class_header]: Header for AVAudioEngine */
// The class instance for the [AudioEngine] class.
var (
	AudioEngineClass     _AudioEngineClass
	AudioEngineClassOnce sync.Once
)

func getAudioEngineClass() _AudioEngineClass {
	AudioEngineClassOnce.Do(func() {
		AudioEngineClass = _AudioEngineClass{objc.GetClass("AVAudioEngine")}
	})
	return AudioEngineClass
}

type _AudioEngineClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioEngine */
// An interface definition for the [AudioEngine] class.
type IAudioEngine interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioEngine */
	// properties:
	AttachedNodes() unsafe.Pointer
	InputNode() IAVAudioInputNode
	AutoShutdownEnabled() bool
	SetAutoShutdownEnabled(value bool)
	IsInManualRenderingMode() bool
	Running() bool
	MainMixerNode() IAVAudioMixerNode
	ManualRenderingBlock() AudioEngineManualRenderingBlock /* not a class type */
	ManualRenderingFormat() IAVAudioFormat
	ManualRenderingMaximumFrameCount() AudioFrameCount /* typedef */
	ManualRenderingMode() AudioEngineManualRenderingMode
	ManualRenderingSampleTime() AudioFramePosition /* typedef */
	MusicSequence() objectivec.IObject
	SetMusicSequence(value objectivec.IObject)
	OutputNode() IAVAudioOutputNode
	IsAutoShutdownEnabled() bool
	SetIsAutoShutdownEnabled(value bool)
	IsRunning() bool
	SetIsRunning(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioEngine */
	// methods:
	AttachNode(node IAVAudioNode)
	ConnectToFormat(node1 IAVAudioNode, node2 IAVAudioNode, format IAVAudioFormat)
	ConnectToConnectionPointsFromBusFormat(sourceNode IAVAudioNode, destNodes []AudioConnectionPoint, sourceBus AudioNodeBus /* typedef */, format IAVAudioFormat)
	ConnectToFromBusToBusFormat(node1 IAVAudioNode, node2 IAVAudioNode, bus1 AudioNodeBus /* typedef */, bus2 AudioNodeBus /* typedef */, format IAVAudioFormat)
	ConnectMIDIToFormatEventListBlock(sourceNode IAVAudioNode, destinationNode IAVAudioNode, format IAVAudioFormat, tapBlock MIDIEventListBlock /* not a class type */)
	ConnectMIDIToNodesFormatEventListBlock(sourceNode IAVAudioNode, destinationNodes []AudioNode, format IAVAudioFormat, tapBlock MIDIEventListBlock /* not a class type */)
	DetachNode(node IAVAudioNode)
	DisableManualRenderingMode()
	DisconnectMIDIFrom(sourceNode IAVAudioNode, destinationNode IAVAudioNode)
	DisconnectMIDIFromNodes(sourceNode IAVAudioNode, destinationNodes []AudioNode)
	DisconnectMIDIInput(node IAVAudioNode)
	DisconnectMIDIOutput(node IAVAudioNode)
	DisconnectNodeInput(node IAVAudioNode)
	DisconnectNodeInputBus(node IAVAudioNode, bus AudioNodeBus /* typedef */)
	DisconnectNodeOutput(node IAVAudioNode)
	DisconnectNodeOutputBus(node IAVAudioNode, bus AudioNodeBus /* typedef */)
	EnableManualRenderingModeFormatMaximumFrameCountError(mode AudioEngineManualRenderingMode, pcmFormat IAVAudioFormat, maximumFrameCount AudioFrameCount /* typedef */, outError objectivec.IObject) bool
	InputConnectionPointForNodeInputBus(node IAVAudioNode, bus AudioNodeBus /* typedef */) IAudioConnectionPoint
	OutputConnectionPointsForNodeOutputBus(node IAVAudioNode, bus AudioNodeBus /* typedef */) []AudioConnectionPoint
	Pause()
	Prepare()
	RenderOfflineToBufferError(numberOfFrames AudioFrameCount /* typedef */, buffer IAVAudioPCMBuffer, outError objectivec.IObject) AudioEngineManualRenderingStatus
	Reset()
	StartAndReturnError(outError objectivec.IObject) bool
	Stop()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioEngine */
// Alloc allocates a new instance without initialization.
func (ac _AudioEngineClass) Alloc() AudioEngine {
	rv := objc.Send[AudioEngine](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioEngineClass) New() AudioEngine {
	rv := objc.Send[AudioEngine](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioEngine) Init() AudioEngine {
	rv := objc.Send[AudioEngine](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioEngine) Autorelease() AudioEngine {
	rv := objc.Send[AudioEngine](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioEngine creates a new AudioEngine instance.
func NewAudioEngine() AudioEngine {
	return getAudioEngineClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioEngine */
// An object that manages a graph of audio nodes, controls playback, and configures real-time rendering constraints.
//
// An audio engine object contains a group of instances that you attach to form an audio processing chain. You can connect, disconnect, and remove audio nodes during runtime with minor limitations. Removing an audio node that has differing channel counts, or that’s a mixer, can break the graph. Reconnect audio nodes only when they’re upstream of a mixer. By default, Audio Engine renders to a connected audio device in real time. You can configure the engine to operate in manual rendering mode when you need to render at, or faster than, real time. In that mode, the engine disconnects from audio devices and your app drives the rendering.


// An object that manages a graph of audio nodes, controls playback, and configures real-time rendering constraints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine
type AudioEngine struct {
	objectivec.Object
}

// AudioEngineFrom constructs a [AudioEngine] from an unsafe.Pointer.
//
// An object that manages a graph of audio nodes, controls playback, and configures real-time rendering constraints.
func AudioEngineFrom(ptr unsafe.Pointer) AudioEngine {
	return AudioEngine{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioEngine */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioEngine */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioEngine */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioEngine */

// Attaches an audio node to the audio engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/attach(_:)
func (a_ AudioEngine) AttachNode(node IAVAudioNode) {
	objc.Send[objc.ID](a_.ID, objc.Sel("attachNode:"), node)
}/* debug [instance_methods/method]: AttachNode */


// Establishes a connection between two nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/connect(_:to:format:)
func (a_ AudioEngine) ConnectToFormat(node1 IAVAudioNode, node2 IAVAudioNode, format IAVAudioFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("connect:to:format:"), node1, node2, format)
}/* debug [instance_methods/method]: ConnectToFormat */


// Establishes a connection between a source node and multiple destination nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/connect(_:to:fromBus:format:)
func (a_ AudioEngine) ConnectToConnectionPointsFromBusFormat(sourceNode IAVAudioNode, destNodes []AudioConnectionPoint, sourceBus AudioNodeBus /* typedef */, format IAVAudioFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("connect:toConnectionPoints:fromBus:format:"), sourceNode, destNodes, sourceBus, format)
}/* debug [instance_methods/method]: ConnectToConnectionPointsFromBusFormat */


// Establishes a connection between two nodes, specifying the input and output busses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/connect(_:to:fromBus:toBus:format:)
func (a_ AudioEngine) ConnectToFromBusToBusFormat(node1 IAVAudioNode, node2 IAVAudioNode, bus1 AudioNodeBus /* typedef */, bus2 AudioNodeBus /* typedef */, format IAVAudioFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("connect:to:fromBus:toBus:format:"), node1, node2, bus1, bus2, format)
}/* debug [instance_methods/method]: ConnectToFromBusToBusFormat */


// Establishes a MIDI connection between two nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/connectMIDI(_:to:format:eventListBlock:)-73cd1
func (a_ AudioEngine) ConnectMIDIToFormatEventListBlock(sourceNode IAVAudioNode, destinationNode IAVAudioNode, format IAVAudioFormat, tapBlock MIDIEventListBlock /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("connectMIDI:to:format:eventListBlock:"), sourceNode, destinationNode, format, tapBlock)
}/* debug [instance_methods/method]: ConnectMIDIToFormatEventListBlock */


// Establishes a MIDI connection between a source node and multiple destination nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/connectMIDI(_:to:format:eventListBlock:)-7qtd5
func (a_ AudioEngine) ConnectMIDIToNodesFormatEventListBlock(sourceNode IAVAudioNode, destinationNodes []AudioNode, format IAVAudioFormat, tapBlock MIDIEventListBlock /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("connectMIDI:toNodes:format:eventListBlock:"), sourceNode, destinationNodes, format, tapBlock)
}/* debug [instance_methods/method]: ConnectMIDIToNodesFormatEventListBlock */


// Detaches an audio node from the audio engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/detach(_:)
func (a_ AudioEngine) DetachNode(node IAVAudioNode) {
	objc.Send[objc.ID](a_.ID, objc.Sel("detachNode:"), node)
}/* debug [instance_methods/method]: DetachNode */


// Sets the engine to render to or from an audio device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disableManualRenderingMode()
func (a_ AudioEngine) DisableManualRenderingMode() {
	objc.Send[objc.ID](a_.ID, objc.Sel("disableManualRenderingMode"))
}/* debug [instance_methods/method]: DisableManualRenderingMode */


// Removes a MIDI connection between two nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disconnectMIDI(_:from:)-1kssy
func (a_ AudioEngine) DisconnectMIDIFrom(sourceNode IAVAudioNode, destinationNode IAVAudioNode) {
	objc.Send[objc.ID](a_.ID, objc.Sel("disconnectMIDI:from:"), sourceNode, destinationNode)
}/* debug [instance_methods/method]: DisconnectMIDIFrom */


// Removes a MIDI connection between one source node and multiple destination nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disconnectMIDI(_:from:)-7oaab
func (a_ AudioEngine) DisconnectMIDIFromNodes(sourceNode IAVAudioNode, destinationNodes []AudioNode) {
	objc.Send[objc.ID](a_.ID, objc.Sel("disconnectMIDI:fromNodes:"), sourceNode, destinationNodes)
}/* debug [instance_methods/method]: DisconnectMIDIFromNodes */


// Disconnects all input MIDI connections from a node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disconnectMIDIInput(_:)
func (a_ AudioEngine) DisconnectMIDIInput(node IAVAudioNode) {
	objc.Send[objc.ID](a_.ID, objc.Sel("disconnectMIDIInput:"), node)
}/* debug [instance_methods/method]: DisconnectMIDIInput */


// Disconnects all output MIDI connections from a node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disconnectMIDIOutput(_:)
func (a_ AudioEngine) DisconnectMIDIOutput(node IAVAudioNode) {
	objc.Send[objc.ID](a_.ID, objc.Sel("disconnectMIDIOutput:"), node)
}/* debug [instance_methods/method]: DisconnectMIDIOutput */


// Removes all input connections of the node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disconnectNodeInput(_:)
func (a_ AudioEngine) DisconnectNodeInput(node IAVAudioNode) {
	objc.Send[objc.ID](a_.ID, objc.Sel("disconnectNodeInput:"), node)
}/* debug [instance_methods/method]: DisconnectNodeInput */


// Removes the input connection of a node on the specified bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disconnectNodeInput(_:bus:)
func (a_ AudioEngine) DisconnectNodeInputBus(node IAVAudioNode, bus AudioNodeBus /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("disconnectNodeInput:bus:"), node, bus)
}/* debug [instance_methods/method]: DisconnectNodeInputBus */


// Removes all output connections of a node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disconnectNodeOutput(_:)
func (a_ AudioEngine) DisconnectNodeOutput(node IAVAudioNode) {
	objc.Send[objc.ID](a_.ID, objc.Sel("disconnectNodeOutput:"), node)
}/* debug [instance_methods/method]: DisconnectNodeOutput */


// Removes the output connection of a node on the specified bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/disconnectNodeOutput(_:bus:)
func (a_ AudioEngine) DisconnectNodeOutputBus(node IAVAudioNode, bus AudioNodeBus /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("disconnectNodeOutput:bus:"), node, bus)
}/* debug [instance_methods/method]: DisconnectNodeOutputBus */


// Sets the engine to operate in manual rendering mode with the render format and maximum frame count you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/enableManualRenderingMode(_:format:maximumFrameCount:)
func (a_ AudioEngine) EnableManualRenderingModeFormatMaximumFrameCountError(mode AudioEngineManualRenderingMode, pcmFormat IAVAudioFormat, maximumFrameCount AudioFrameCount /* typedef */, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("enableManualRenderingMode:format:maximumFrameCount:error:"), mode, pcmFormat, maximumFrameCount, outError)
	return rv
}/* debug [instance_methods/method]: EnableManualRenderingModeFormatMaximumFrameCountError */


// Returns connection information about a node’s input bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/inputConnectionPoint(for:inputBus:)
func (a_ AudioEngine) InputConnectionPointForNodeInputBus(node IAVAudioNode, bus AudioNodeBus /* typedef */) IAudioConnectionPoint {
	rv := objc.Send[AudioConnectionPoint](a_.ID, objc.Sel("inputConnectionPointForNode:inputBus:"), node, bus)
	return rv
}/* debug [instance_methods/method]: InputConnectionPointForNodeInputBus */


// Returns connection information about a node’s output bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/outputConnectionPoints(for:outputBus:)
func (a_ AudioEngine) OutputConnectionPointsForNodeOutputBus(node IAVAudioNode, bus AudioNodeBus /* typedef */) []AudioConnectionPoint {
	rv := objc.Send[[]AudioConnectionPoint](a_.ID, objc.Sel("outputConnectionPointsForNode:outputBus:"), node, bus)
	return rv
}/* debug [instance_methods/method]: OutputConnectionPointsForNodeOutputBus */


// Pauses the audio engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/pause()
func (a_ AudioEngine) Pause() {
	objc.Send[objc.ID](a_.ID, objc.Sel("pause"))
}/* debug [instance_methods/method]: Pause */


// Prepares the audio engine for starting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/prepare()
func (a_ AudioEngine) Prepare() {
	objc.Send[objc.ID](a_.ID, objc.Sel("prepare"))
}/* debug [instance_methods/method]: Prepare */


// Makes a render call to the engine operating in the offline manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/renderOffline(_:to:)
func (a_ AudioEngine) RenderOfflineToBufferError(numberOfFrames AudioFrameCount /* typedef */, buffer IAVAudioPCMBuffer, outError objectivec.IObject) AudioEngineManualRenderingStatus {
	rv := objc.Send[AudioEngineManualRenderingStatus](a_.ID, objc.Sel("renderOffline:toBuffer:error:"), numberOfFrames, buffer, outError)
	return rv
}/* debug [instance_methods/method]: RenderOfflineToBufferError */


// Resets all audio nodes in the audio engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/reset()
func (a_ AudioEngine) Reset() {
	objc.Send[objc.ID](a_.ID, objc.Sel("reset"))
}/* debug [instance_methods/method]: Reset */


// Starts the audio engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/start()
func (a_ AudioEngine) StartAndReturnError(outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("startAndReturnError:"), outError)
	return rv
}/* debug [instance_methods/method]: StartAndReturnError */


// Stops the audio engine and releases any previously prepared resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/stop()
func (a_ AudioEngine) Stop() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stop"))
}/* debug [instance_methods/method]: Stop */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioEngine */

// A read-only set that contains the nodes you attach to the audio engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/attachedNodes
func (a_ AudioEngine) AttachedNodes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("attachedNodes"))
	return rv
}/* debug [instance_properties/getter]: attachedNodes */


// The audio engine’s singleton input audio node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/inputNode
func (a_ AudioEngine) InputNode() IAVAudioInputNode {
	rv := objc.Send[AudioInputNode](a_.ID, objc.Sel("inputNode"))
	return rv
}/* debug [instance_properties/getter]: inputNode */


// A Boolean value that indicates whether autoshutdown is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/isAutoShutdownEnabled
func (a_ AudioEngine) AutoShutdownEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("autoShutdownEnabled"))
	return rv
}/* debug [instance_properties/getter]: autoShutdownEnabled */


// A Boolean value that indicates whether autoshutdown is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/isAutoShutdownEnabled
func (a_ AudioEngine) SetAutoShutdownEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAutoShutdownEnabled:"), value)
}/* debug [instance_properties/setter]: autoShutdownEnabled */


// A Boolean value that indicates whether the engine is operating in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/isInManualRenderingMode
func (a_ AudioEngine) IsInManualRenderingMode() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isInManualRenderingMode"))
	return rv
}/* debug [instance_properties/getter]: isInManualRenderingMode */


// A Boolean value that indicates whether the audio engine is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/isRunning
func (a_ AudioEngine) Running() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("running"))
	return rv
}/* debug [instance_properties/getter]: running */


// The audio engine’s optional singleton main mixer node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/mainMixerNode
func (a_ AudioEngine) MainMixerNode() IAVAudioMixerNode {
	rv := objc.Send[AudioMixerNode](a_.ID, objc.Sel("mainMixerNode"))
	return rv
}/* debug [instance_properties/getter]: mainMixerNode */


// The block that renders the engine when operating in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/manualRenderingBlock
func (a_ AudioEngine) ManualRenderingBlock() AudioEngineManualRenderingBlock /* not a class type */ {
	rv := objc.Send[AudioEngineManualRenderingBlock](a_.ID, objc.Sel("manualRenderingBlock"))
	return rv
}/* debug [instance_properties/getter]: manualRenderingBlock */


// The render format of the engine in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/manualRenderingFormat
func (a_ AudioEngine) ManualRenderingFormat() IAVAudioFormat {
	rv := objc.Send[AudioFormat](a_.ID, objc.Sel("manualRenderingFormat"))
	return rv
}/* debug [instance_properties/getter]: manualRenderingFormat */


// The maximum number of PCM sample frames the engine produces in any single render call in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/manualRenderingMaximumFrameCount
func (a_ AudioEngine) ManualRenderingMaximumFrameCount() AudioFrameCount /* typedef */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("manualRenderingMaximumFrameCount"))
	return rv
}/* debug [instance_properties/getter]: manualRenderingMaximumFrameCount */


// The manual rendering mode configured on the engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/manualRenderingMode
func (a_ AudioEngine) ManualRenderingMode() AudioEngineManualRenderingMode {
	rv := objc.Send[AudioEngineManualRenderingMode](a_.ID, objc.Sel("manualRenderingMode"))
	return rv
}/* debug [instance_properties/getter]: manualRenderingMode */


// An indication of where the engine is on its render timeline in manual rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/manualRenderingSampleTime
func (a_ AudioEngine) ManualRenderingSampleTime() AudioFramePosition /* typedef */ {
	rv := objc.Send[int64](a_.ID, objc.Sel("manualRenderingSampleTime"))
	return rv
}/* debug [instance_properties/getter]: manualRenderingSampleTime */


// The music sequence instance that you attach to the audio engine, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/musicSequence
func (a_ AudioEngine) MusicSequence() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("musicSequence"))
	return rv
}/* debug [instance_properties/getter]: musicSequence */


// The music sequence instance that you attach to the audio engine, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/musicSequence
func (a_ AudioEngine) SetMusicSequence(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMusicSequence:"), value)
}/* debug [instance_properties/setter]: musicSequence */


// The audio engine’s singleton output audio node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/outputNode
func (a_ AudioEngine) OutputNode() IAVAudioOutputNode {
	rv := objc.Send[AudioOutputNode](a_.ID, objc.Sel("outputNode"))
	return rv
}/* debug [instance_properties/getter]: outputNode */


// A Boolean value that indicates whether autoshutdown is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/isautoshutdownenabled
func (a_ AudioEngine) IsAutoShutdownEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isAutoShutdownEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAutoShutdownEnabled */


// A Boolean value that indicates whether autoshutdown is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/isautoshutdownenabled
func (a_ AudioEngine) SetIsAutoShutdownEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsAutoShutdownEnabled:"), value)
}/* debug [instance_properties/setter]: isAutoShutdownEnabled */


// A Boolean value that indicates whether the audio engine is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/isrunning
func (a_ AudioEngine) IsRunning() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRunning"))
	return rv
}/* debug [instance_properties/getter]: isRunning */


// A Boolean value that indicates whether the audio engine is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioengine/isrunning
func (a_ AudioEngine) SetIsRunning(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRunning:"), value)
}/* debug [instance_properties/setter]: isRunning */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioEngine */


