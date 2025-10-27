// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [SampleBufferAudioRenderer] class.
var (
	SampleBufferAudioRendererClass     _SampleBufferAudioRendererClass
	SampleBufferAudioRendererClassOnce sync.Once
)

func getSampleBufferAudioRendererClass() _SampleBufferAudioRendererClass {
	SampleBufferAudioRendererClassOnce.Do(func() {
		SampleBufferAudioRendererClass = _SampleBufferAudioRendererClass{objc.GetClass("AVSampleBufferAudioRenderer")}
	})
	return SampleBufferAudioRendererClass
}

type _SampleBufferAudioRendererClass struct {
	class objc.Class
}





// An interface definition for the [SampleBufferAudioRenderer] class.
type ISampleBufferAudioRenderer interface {
	objectivec.IObject
	

	// properties:
	AllowedAudioSpatializationFormats() AudioSpatializationFormats
	SetAllowedAudioSpatializationFormats(value AudioSpatializationFormats)
	AudioOutputDeviceUniqueID() foundation.foundation.INSString
	SetAudioOutputDeviceUniqueID(value foundation.foundation.INSString)
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm)
	Error() foundation.foundation.INSError
	Muted() bool
	SetMuted(value bool)
	Status() QueuedSampleBufferRenderingStatus
	Volume() float32
	SetVolume(value float32)
	IsMuted() bool
	SetIsMuted(value bool)
	AVSampleBufferAudioRendererFlushTimeKey() foundation.foundation.INSString


	

	// methods:
	FlushFromSourceTimeCompletionHandler(time objectivec.IObject, completionHandler unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (sc _SampleBufferAudioRendererClass) Alloc() SampleBufferAudioRenderer {
	rv := objc.Send[SampleBufferAudioRenderer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SampleBufferAudioRendererClass) New() SampleBufferAudioRenderer {
	rv := objc.Send[SampleBufferAudioRenderer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SampleBufferAudioRenderer) Init() SampleBufferAudioRenderer {
	rv := objc.Send[SampleBufferAudioRenderer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SampleBufferAudioRenderer) Autorelease() SampleBufferAudioRenderer {
	rv := objc.Send[SampleBufferAudioRenderer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSampleBufferAudioRenderer creates a new SampleBufferAudioRenderer instance.
func NewSampleBufferAudioRenderer() SampleBufferAudioRenderer {
	return getSampleBufferAudioRendererClass().New()
}





// An object used to decompress audio and play compressed or uncompressed audio.
//
// You must add an instance of this class to an before queuing the first sample buffer.


// An object used to decompress audio and play compressed or uncompressed audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer
type SampleBufferAudioRenderer struct {
	objectivec.Object
}

// SampleBufferAudioRendererFrom constructs a [SampleBufferAudioRenderer] from an unsafe.Pointer.
//
// An object used to decompress audio and play compressed or uncompressed audio.
func SampleBufferAudioRendererFrom(ptr unsafe.Pointer) SampleBufferAudioRenderer {
	return SampleBufferAudioRenderer{objectivec.Object{objc.ID(ptr)}}
}




















// Flushes queued sample buffers with presentation time stamps later than or equal to the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/flush(fromSourceTime:completionHandler:)
func (s_ SampleBufferAudioRenderer) FlushFromSourceTimeCompletionHandler(time objectivec.IObject, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("flushFromSourceTime:completionHandler:"), time, completionHandler)
}







// The source audio channel layouts the audio renderer supports for spatialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/allowedAudioSpatializationFormats
func (s_ SampleBufferAudioRenderer) AllowedAudioSpatializationFormats() AudioSpatializationFormats {
	rv := objc.Send[AudioSpatializationFormats](s_.ID, objc.Sel("allowedAudioSpatializationFormats"))
	return rv
}


// The source audio channel layouts the audio renderer supports for spatialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/allowedAudioSpatializationFormats
func (s_ SampleBufferAudioRenderer) SetAllowedAudioSpatializationFormats(value AudioSpatializationFormats) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowedAudioSpatializationFormats:"), value)
}


// The unique identifier of the output device used to play audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/audioOutputDeviceUniqueID
func (s_ SampleBufferAudioRenderer) AudioOutputDeviceUniqueID() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("audioOutputDeviceUniqueID"))
	return rv
}


// The unique identifier of the output device used to play audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/audioOutputDeviceUniqueID
func (s_ SampleBufferAudioRenderer) SetAudioOutputDeviceUniqueID(value foundation.foundation.INSString) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAudioOutputDeviceUniqueID:"), value)
}


// The processing algorithm used to manage audio pitch at different rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/audioTimePitchAlgorithm
func (s_ SampleBufferAudioRenderer) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm {
	rv := objc.Send[AudioTimePitchAlgorithm](s_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}


// The processing algorithm used to manage audio pitch at different rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/audioTimePitchAlgorithm
func (s_ SampleBufferAudioRenderer) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}


// The error that caused the renderer to no longer render sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/error
func (s_ SampleBufferAudioRenderer) Error() foundation.foundation.INSError {
	rv := objc.Send[foundation.NSError](s_.ID, objc.Sel("error"))
	return rv
}


// A Boolean value that indicates whether audio for the renderer is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/isMuted
func (s_ SampleBufferAudioRenderer) Muted() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("muted"))
	return rv
}


// A Boolean value that indicates whether audio for the renderer is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/isMuted
func (s_ SampleBufferAudioRenderer) SetMuted(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMuted:"), value)
}


// The status of the audio renderer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/status
func (s_ SampleBufferAudioRenderer) Status() QueuedSampleBufferRenderingStatus {
	rv := objc.Send[QueuedSampleBufferRenderingStatus](s_.ID, objc.Sel("status"))
	return rv
}


// The current audio volume for the audio renderer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/volume
func (s_ SampleBufferAudioRenderer) Volume() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("volume"))
	return rv
}


// The current audio volume for the audio renderer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/volume
func (s_ SampleBufferAudioRenderer) SetVolume(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVolume:"), value)
}


// A Boolean value that indicates whether audio for the renderer is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/ismuted
func (s_ SampleBufferAudioRenderer) IsMuted() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isMuted"))
	return rv
}


// A Boolean value that indicates whether audio for the renderer is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/ismuted
func (s_ SampleBufferAudioRenderer) SetIsMuted(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsMuted:"), value)
}


// The key that indicates the presentation timestamp of the first queued sample that was flushed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorendererflushtimekey
func (s_ SampleBufferAudioRenderer) AVSampleBufferAudioRendererFlushTimeKey() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("AVSampleBufferAudioRendererFlushTimeKey"))
	return rv
}








