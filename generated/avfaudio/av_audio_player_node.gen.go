// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVAudioPlayerNode */

/* debug [class_header]: Header for AVAudioPlayerNode */
// The class instance for the [AudioPlayerNode] class.
var (
	AudioPlayerNodeClass     _AudioPlayerNodeClass
	AudioPlayerNodeClassOnce sync.Once
)

func getAudioPlayerNodeClass() _AudioPlayerNodeClass {
	AudioPlayerNodeClassOnce.Do(func() {
		AudioPlayerNodeClass = _AudioPlayerNodeClass{objc.GetClass("AVAudioPlayerNode")}
	})
	return AudioPlayerNodeClass
}

type _AudioPlayerNodeClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for AudioPlayerNode */
// An interface definition for the [AudioPlayerNode] class.
type IAudioPlayerNode interface {
	IAudioNode

	/* debug [class_interface_properties]: Properties for AudioPlayerNode */
	// properties:
	Playing() bool
	LastRenderTime() objc.IObject /* cross-framework: AudioTime */
	SetLastRenderTime(value objc.IObject /* cross-framework: AudioTime */)
	Latency() float64
	SetLatency(value float64)
	OutputPresentationLatency() float64
	SetOutputPresentationLatency(value float64)
	IsPlaying() bool
	SetIsPlaying(value bool)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for AudioPlayerNode */
	// methods:
	NodeTimeForPlayerTime(playerTime objc.IObject /* cross-framework: AudioTime */) objc.IObject /* cross-framework: AudioTime */
	Pause()
	Play()
	PlayAtTime(when objc.IObject /* cross-framework: AudioTime */)
	PlayerTimeForNodeTime(nodeTime objc.IObject /* cross-framework: AudioTime */) objc.IObject /* cross-framework: AudioTime */
	PrepareWithFrameCount(frameCount AudioFrameCount /* not a class type */)
	ScheduleBufferAtTimeOptionsCompletionCallbackTypeCompletionHandler(buffer objc.IObject /* cross-framework: AudioPCMBuffer */, when objc.IObject /* cross-framework: AudioTime */, options AudioPlayerNodeBufferOptions, callbackType AudioPlayerNodeCompletionCallbackType, completionHandler AudioPlayerNodeCompletionHandler /* not a class type */)
	ScheduleBufferAtTimeOptionsCompletionHandler(buffer objc.IObject /* cross-framework: AudioPCMBuffer */, when objc.IObject /* cross-framework: AudioTime */, options AudioPlayerNodeBufferOptions, completionHandler AudioNodeCompletionHandler /* not a class type */)
	ScheduleBufferCompletionCallbackTypeCompletionHandler(buffer objc.IObject /* cross-framework: AudioPCMBuffer */, callbackType AudioPlayerNodeCompletionCallbackType, completionHandler AudioPlayerNodeCompletionHandler /* not a class type */)
	ScheduleBufferCompletionHandler(buffer objc.IObject /* cross-framework: AudioPCMBuffer */, completionHandler AudioNodeCompletionHandler /* not a class type */)
	ScheduleFileAtTimeCompletionCallbackTypeCompletionHandler(file objc.IObject /* cross-framework: AudioFile */, when objc.IObject /* cross-framework: AudioTime */, callbackType AudioPlayerNodeCompletionCallbackType, completionHandler AudioPlayerNodeCompletionHandler /* not a class type */)
	ScheduleFileAtTimeCompletionHandler(file objc.IObject /* cross-framework: AudioFile */, when objc.IObject /* cross-framework: AudioTime */, completionHandler AudioNodeCompletionHandler /* not a class type */)
	ScheduleSegmentStartingFrameFrameCountAtTimeCompletionCallbackTypeCompletionHandler(file objc.IObject /* cross-framework: AudioFile */, startFrame AudioFramePosition /* not a class type */, numberFrames AudioFrameCount /* not a class type */, when objc.IObject /* cross-framework: AudioTime */, callbackType AudioPlayerNodeCompletionCallbackType, completionHandler AudioPlayerNodeCompletionHandler /* not a class type */)
	ScheduleSegmentStartingFrameFrameCountAtTimeCompletionHandler(file objc.IObject /* cross-framework: AudioFile */, startFrame AudioFramePosition /* not a class type */, numberFrames AudioFrameCount /* not a class type */, when objc.IObject /* cross-framework: AudioTime */, completionHandler AudioNodeCompletionHandler /* not a class type */)
	Stop()
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for AudioPlayerNode */
// Alloc allocates a new instance without initialization.
func (ac _AudioPlayerNodeClass) Alloc() AudioPlayerNode {
	rv := objc.Send[AudioPlayerNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioPlayerNodeClass) New() AudioPlayerNode {
	rv := objc.Send[AudioPlayerNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioPlayerNode) Init() AudioPlayerNode {
	rv := objc.Send[AudioPlayerNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioPlayerNode) Autorelease() AudioPlayerNode {
	rv := objc.Send[AudioPlayerNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioPlayerNode creates a new AudioPlayerNode instance.
func NewAudioPlayerNode() AudioPlayerNode {
	return getAudioPlayerNodeClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for AudioPlayerNode */
// An object for scheduling the playback of buffers or segments of audio files.
//
// This audio node supports scheduling the playback of instances, or segments of audio files that you open through . You can schedule buffers and segments to play at specific points in time or to play immediately following preceding segments. Generally, you want to configure the node’s output format with the same number of channels as in the files and buffers. Otherwise, the node drops or adds channels as necessary. It’s usually preferable to use an for this configuration. Similarly, when playing file segments, the node makes sample rate conversions, if necessary. It’s preferable to configure the node’s output sample rate to match that of the files, and to use a mixer to perform the rate conversion. When playing buffers, there’s an implicit assumption that the buffers are at the same sample rate as the node’s output format. The method unschedules all previously scheduled buffers and file segments, and returns the player timeline to sample time .

// An object for scheduling the playback of buffers or segments of audio files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode
type AudioPlayerNode struct {
	AudioNode
}

// AudioPlayerNodeFrom constructs a [AudioPlayerNode] from an unsafe.Pointer.
//
// An object for scheduling the playback of buffers or segments of audio files.
func AudioPlayerNodeFrom(ptr unsafe.Pointer) AudioPlayerNode {
	return AudioPlayerNode{
		AudioNode: AudioNodeFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for AudioPlayerNode */
/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for AudioPlayerNode */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for AudioPlayerNode */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for AudioPlayerNode */

// Converts from player time to node time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/nodeTime(forPlayerTime:)
func (a_ AudioPlayerNode) NodeTimeForPlayerTime(playerTime objc.IObject /* cross-framework: AudioTime */) objc.IObject /* cross-framework: AudioTime */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("nodeTimeForPlayerTime:"), playerTime)
	return rv
} /* debug [instance_methods/method]: NodeTimeForPlayerTime */

// Pauses the node’s playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/pause()
func (a_ AudioPlayerNode) Pause() {
	objc.Send[objc.ID](a_.ID, objc.Sel("pause"))
} /* debug [instance_methods/method]: Pause */

// Starts or resumes playback immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/play()
func (a_ AudioPlayerNode) Play() {
	objc.Send[objc.ID](a_.ID, objc.Sel("play"))
} /* debug [instance_methods/method]: Play */

// Starts or resumes playback at a time you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/play(at:)
func (a_ AudioPlayerNode) PlayAtTime(when objc.IObject /* cross-framework: AudioTime */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("playAtTime:"), when)
} /* debug [instance_methods/method]: PlayAtTime */

// Converts from node time to player time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/playerTime(forNodeTime:)
func (a_ AudioPlayerNode) PlayerTimeForNodeTime(nodeTime objc.IObject /* cross-framework: AudioTime */) objc.IObject /* cross-framework: AudioTime */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("playerTimeForNodeTime:"), nodeTime)
	return rv
} /* debug [instance_methods/method]: PlayerTimeForNodeTime */

// Prepares the file regions or buffers you schedule for playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/prepare(withFrameCount:)
func (a_ AudioPlayerNode) PrepareWithFrameCount(frameCount AudioFrameCount /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("prepareWithFrameCount:"), frameCount)
} /* debug [instance_methods/method]: PrepareWithFrameCount */

// Schedules the playing samples from an audio buffer with the playback options you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/scheduleBuffer(_:at:options:completionCallbackType:completionHandler:)
func (a_ AudioPlayerNode) ScheduleBufferAtTimeOptionsCompletionCallbackTypeCompletionHandler(buffer objc.IObject /* cross-framework: AudioPCMBuffer */, when objc.IObject /* cross-framework: AudioTime */, options AudioPlayerNodeBufferOptions, callbackType AudioPlayerNodeCompletionCallbackType, completionHandler AudioPlayerNodeCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("scheduleBuffer:atTime:options:completionCallbackType:completionHandler:"), buffer, when, options, callbackType, completionHandler)
} /* debug [instance_methods/method]: ScheduleBufferAtTimeOptionsCompletionCallbackTypeCompletionHandler */

// Schedules the playing samples from an audio buffer at the time and playback options you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/scheduleBuffer(_:at:options:completionHandler:)
func (a_ AudioPlayerNode) ScheduleBufferAtTimeOptionsCompletionHandler(buffer objc.IObject /* cross-framework: AudioPCMBuffer */, when objc.IObject /* cross-framework: AudioTime */, options AudioPlayerNodeBufferOptions, completionHandler AudioNodeCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("scheduleBuffer:atTime:options:completionHandler:"), buffer, when, options, completionHandler)
} /* debug [instance_methods/method]: ScheduleBufferAtTimeOptionsCompletionHandler */

// Schedules the playing samples from an audio buffer with the callback option you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/scheduleBuffer(_:completionCallbackType:completionHandler:)
func (a_ AudioPlayerNode) ScheduleBufferCompletionCallbackTypeCompletionHandler(buffer objc.IObject /* cross-framework: AudioPCMBuffer */, callbackType AudioPlayerNodeCompletionCallbackType, completionHandler AudioPlayerNodeCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("scheduleBuffer:completionCallbackType:completionHandler:"), buffer, callbackType, completionHandler)
} /* debug [instance_methods/method]: ScheduleBufferCompletionCallbackTypeCompletionHandler */

// Schedules the playing samples from an audio buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/scheduleBuffer(_:completionHandler:)
func (a_ AudioPlayerNode) ScheduleBufferCompletionHandler(buffer objc.IObject /* cross-framework: AudioPCMBuffer */, completionHandler AudioNodeCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("scheduleBuffer:completionHandler:"), buffer, completionHandler)
} /* debug [instance_methods/method]: ScheduleBufferCompletionHandler */

// Schedules the playing of an entire audio file with a callback option you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/scheduleFile(_:at:completionCallbackType:completionHandler:)
func (a_ AudioPlayerNode) ScheduleFileAtTimeCompletionCallbackTypeCompletionHandler(file objc.IObject /* cross-framework: AudioFile */, when objc.IObject /* cross-framework: AudioTime */, callbackType AudioPlayerNodeCompletionCallbackType, completionHandler AudioPlayerNodeCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("scheduleFile:atTime:completionCallbackType:completionHandler:"), file, when, callbackType, completionHandler)
} /* debug [instance_methods/method]: ScheduleFileAtTimeCompletionCallbackTypeCompletionHandler */

// Schedules the playing of an entire audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/scheduleFile(_:at:completionHandler:)
func (a_ AudioPlayerNode) ScheduleFileAtTimeCompletionHandler(file objc.IObject /* cross-framework: AudioFile */, when objc.IObject /* cross-framework: AudioTime */, completionHandler AudioNodeCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("scheduleFile:atTime:completionHandler:"), file, when, completionHandler)
} /* debug [instance_methods/method]: ScheduleFileAtTimeCompletionHandler */

// Schedules the playing of an audio file segment with a callback option you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/scheduleSegment(_:startingFrame:frameCount:at:completionCallbackType:completionHandler:)
func (a_ AudioPlayerNode) ScheduleSegmentStartingFrameFrameCountAtTimeCompletionCallbackTypeCompletionHandler(file objc.IObject /* cross-framework: AudioFile */, startFrame AudioFramePosition /* not a class type */, numberFrames AudioFrameCount /* not a class type */, when objc.IObject /* cross-framework: AudioTime */, callbackType AudioPlayerNodeCompletionCallbackType, completionHandler AudioPlayerNodeCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("scheduleSegment:startingFrame:frameCount:atTime:completionCallbackType:completionHandler:"), file, startFrame, numberFrames, when, callbackType, completionHandler)
} /* debug [instance_methods/method]: ScheduleSegmentStartingFrameFrameCountAtTimeCompletionCallbackTypeCompletionHandler */

// Schedules the playing of an audio file segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/scheduleSegment(_:startingFrame:frameCount:at:completionHandler:)
func (a_ AudioPlayerNode) ScheduleSegmentStartingFrameFrameCountAtTimeCompletionHandler(file objc.IObject /* cross-framework: AudioFile */, startFrame AudioFramePosition /* not a class type */, numberFrames AudioFrameCount /* not a class type */, when objc.IObject /* cross-framework: AudioTime */, completionHandler AudioNodeCompletionHandler /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("scheduleSegment:startingFrame:frameCount:atTime:completionHandler:"), file, startFrame, numberFrames, when, completionHandler)
} /* debug [instance_methods/method]: ScheduleSegmentStartingFrameFrameCountAtTimeCompletionHandler */

// Clears all of the node’s events you schedule and stops playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/stop()
func (a_ AudioPlayerNode) Stop() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stop"))
} /* debug [instance_methods/method]: Stop */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for AudioPlayerNode */

// A Boolean value that indicates whether the player is playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerNode/isPlaying
func (a_ AudioPlayerNode) Playing() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("playing"))
	return rv
} /* debug [instance_properties/getter]: playing */

// The most recent render time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/lastrendertime
func (a_ AudioPlayerNode) LastRenderTime() objc.IObject /* cross-framework: AudioTime */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("lastRenderTime"))
	return rv
} /* debug [instance_properties/getter]: lastRenderTime */

// The most recent render time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/lastrendertime
func (a_ AudioPlayerNode) SetLastRenderTime(value objc.IObject /* cross-framework: AudioTime */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLastRenderTime:"), value)
} /* debug [instance_properties/setter]: lastRenderTime */

// The processing latency of the node, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/latency
func (a_ AudioPlayerNode) Latency() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("latency"))
	return rv
} /* debug [instance_properties/getter]: latency */

// The processing latency of the node, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/latency
func (a_ AudioPlayerNode) SetLatency(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLatency:"), value)
} /* debug [instance_properties/setter]: latency */

// The maximum render pipeline latency downstream of the node, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/outputpresentationlatency
func (a_ AudioPlayerNode) OutputPresentationLatency() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("outputPresentationLatency"))
	return rv
} /* debug [instance_properties/getter]: outputPresentationLatency */

// The maximum render pipeline latency downstream of the node, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/outputpresentationlatency
func (a_ AudioPlayerNode) SetOutputPresentationLatency(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputPresentationLatency:"), value)
} /* debug [instance_properties/setter]: outputPresentationLatency */

// A Boolean value that indicates whether the player is playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioplayernode/isplaying
func (a_ AudioPlayerNode) IsPlaying() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPlaying"))
	return rv
} /* debug [instance_properties/getter]: isPlaying */

// A Boolean value that indicates whether the player is playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioplayernode/isplaying
func (a_ AudioPlayerNode) SetIsPlaying(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPlaying:"), value)
} /* debug [instance_properties/setter]: isPlaying */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class AVAudioPlayerNode */
