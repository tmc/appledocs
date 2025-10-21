// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [AudioPlayerNode] class.
type IAudioPlayerNode interface {
	IAudioNode
}

// An object for scheduling the playback of buffers or segments of audio files.
//
// This audio node supports scheduling the playback of instances, or segments of audio files that you open through . You can schedule buffers and segments to play at specific points in time or to play immediately following preceding segments. Generally, you want to configure the node’s output format with the same number of channels as in the files and buffers. Otherwise, the node drops or adds channels as necessary. It’s usually preferable to use an for this configuration. Similarly, when playing file segments, the node makes sample rate conversions, if necessary. It’s preferable to configure the node’s output sample rate to match that of the files, and to use a mixer to perform the rate conversion. When playing buffers, there’s an implicit assumption that the buffers are at the same sample rate as the node’s output format. The method unschedules all previously scheduled buffers and file segments, and returns the player timeline to sample time .
//
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

// Alloc allocates a new instance without initialization.
func (ac _AudioPlayerNodeClass) Alloc() AudioPlayerNode {
	rv := objc.Send[AudioPlayerNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The most recent render time.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/lastrendertime
func (a_ AudioPlayerNode) LastRenderTime() AVAudioTime {
	rv := objc.Send[AVAudioTime](a_.ID, objc.Sel("lastRenderTime"))
	return rv
}


// SetLastRenderTime sets the value of the lastRenderTime property.
// The most recent render time.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/lastrendertime
func (a_ AudioPlayerNode) SetLastRenderTime(value IAVAudioTime) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLastRenderTime:"), value)
}

// The processing latency of the node, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/latency
func (a_ AudioPlayerNode) Latency() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("latency"))
	return rv
}


// SetLatency sets the value of the latency property.
// The processing latency of the node, in seconds.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/latency
func (a_ AudioPlayerNode) SetLatency(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLatency:"), value)
}

// The maximum render pipeline latency downstream of the node, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/outputpresentationlatency
func (a_ AudioPlayerNode) OutputPresentationLatency() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("outputPresentationLatency"))
	return rv
}


// SetOutputPresentationLatency sets the value of the outputPresentationLatency property.
// The maximum render pipeline latency downstream of the node, in seconds.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudionode/outputpresentationlatency
func (a_ AudioPlayerNode) SetOutputPresentationLatency(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputPresentationLatency:"), value)
}

// A Boolean value that indicates whether the player is playing.
//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioplayernode/isplaying
func (a_ AudioPlayerNode) IsPlaying() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isPlaying"))
	return rv
}


// SetIsPlaying sets the value of the isPlaying property.
// A Boolean value that indicates whether the player is playing.

//
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioplayernode/isplaying
func (a_ AudioPlayerNode) SetIsPlaying(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsPlaying:"), value)
}



