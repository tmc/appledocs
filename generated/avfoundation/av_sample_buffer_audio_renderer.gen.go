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
}

// An object used to decompress audio and play compressed or uncompressed audio.
//
// You must add an instance of this class to an before queuing the first sample buffer.
//
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

// Alloc allocates a new instance without initialization.
func (sc _SampleBufferAudioRendererClass) Alloc() SampleBufferAudioRenderer {
	rv := objc.Send[SampleBufferAudioRenderer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The unique identifier of the output device used to play audio.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/audioOutputDeviceUniqueID
func (s_ SampleBufferAudioRenderer) AudioOutputDeviceUniqueID() string {
	rv := objc.Send[string](s_.ID, objc.Sel("audioOutputDeviceUniqueID"))
	return rv
}


// SetAudioOutputDeviceUniqueID sets the value of the audioOutputDeviceUniqueID property.
// The unique identifier of the output device used to play audio.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/audioOutputDeviceUniqueID
func (s_ SampleBufferAudioRenderer) SetAudioOutputDeviceUniqueID(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAudioOutputDeviceUniqueID:"), objc.String(value))
}

// The source audio channel layouts the audio renderer supports for spatialization.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/allowedaudiospatializationformats
func (s_ SampleBufferAudioRenderer) AllowedAudioSpatializationFormats() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("allowedAudioSpatializationFormats"))
	return rv
}


// SetAllowedAudioSpatializationFormats sets the value of the allowedAudioSpatializationFormats property.
// The source audio channel layouts the audio renderer supports for spatialization.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/allowedaudiospatializationformats
func (s_ SampleBufferAudioRenderer) SetAllowedAudioSpatializationFormats(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowedAudioSpatializationFormats:"), value)
}

// The processing algorithm used to manage audio pitch at different rates.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/audiotimepitchalgorithm
func (s_ SampleBufferAudioRenderer) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm {
	rv := objc.Send[AudioTimePitchAlgorithm](s_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}


// SetAudioTimePitchAlgorithm sets the value of the audioTimePitchAlgorithm property.
// The processing algorithm used to manage audio pitch at different rates.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/audiotimepitchalgorithm
func (s_ SampleBufferAudioRenderer) SetAudioTimePitchAlgorithm(value IAudioTimePitchAlgorithm) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}

// The error that caused the renderer to no longer render sample buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/error
func (s_ SampleBufferAudioRenderer) Error() Error {
	rv := objc.Send[Error](s_.ID, objc.Sel("error"))
	return rv
}


// SetError sets the value of the error property.
// The error that caused the renderer to no longer render sample buffers.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/error
func (s_ SampleBufferAudioRenderer) SetError(value IError) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setError:"), value)
}

// A Boolean value that indicates whether audio for the renderer is in a muted state.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/ismuted
func (s_ SampleBufferAudioRenderer) IsMuted() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isMuted"))
	return rv
}


// SetIsMuted sets the value of the isMuted property.
// A Boolean value that indicates whether audio for the renderer is in a muted state.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/ismuted
func (s_ SampleBufferAudioRenderer) SetIsMuted(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsMuted:"), value)
}

// The status of the audio renderer.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/status
func (s_ SampleBufferAudioRenderer) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
// The status of the audio renderer.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/status
func (s_ SampleBufferAudioRenderer) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStatus:"), value)
}

// The current audio volume for the audio renderer.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/volume
func (s_ SampleBufferAudioRenderer) Volume() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("volume"))
	return rv
}


// SetVolume sets the value of the volume property.
// The current audio volume for the audio renderer.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/volume
func (s_ SampleBufferAudioRenderer) SetVolume(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVolume:"), value)
}

// The key that indicates the presentation timestamp of the first queued sample that was flushed.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorendererflushtimekey
func (s_ SampleBufferAudioRenderer) AVSampleBufferAudioRendererFlushTimeKey() string {
	rv := objc.Send[string](s_.ID, objc.Sel("AVSampleBufferAudioRendererFlushTimeKey"))
	return rv
}



