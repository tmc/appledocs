// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
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
	AudioOutputDeviceUniqueID() objc.IObject /* cross-framework: NSString */
	SetAudioOutputDeviceUniqueID(value objc.IObject /* cross-framework: NSString */)
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* not a class type */
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* not a class type */)
	Error() coretelephony.Error
	SetError(value coretelephony.Error)
	IsMuted() bool
	SetIsMuted(value bool)
	Status() QueuedSampleBufferRenderingStatus /* not a class type */
	SetStatus(value QueuedSampleBufferRenderingStatus /* not a class type */)
	Volume() float32
	SetVolume(value float32)
	AVSampleBufferAudioRendererFlushTimeKey() objc.IObject /* cross-framework: NSString */
	// methods:
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



// The source audio channel layouts the audio renderer supports for spatialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/allowedaudiospatializationformats
func (s_ SampleBufferAudioRenderer) AllowedAudioSpatializationFormats() AudioSpatializationFormats {
	rv := objc.Send[AudioSpatializationFormats](s_.ID, objc.Sel("allowedAudioSpatializationFormats"))
	return rv
}


// The source audio channel layouts the audio renderer supports for spatialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/allowedaudiospatializationformats
func (s_ SampleBufferAudioRenderer) SetAllowedAudioSpatializationFormats(value AudioSpatializationFormats) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowedAudioSpatializationFormats:"), value)
}


// The unique identifier of the output device used to play audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/audiooutputdeviceuniqueid
func (s_ SampleBufferAudioRenderer) AudioOutputDeviceUniqueID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("audioOutputDeviceUniqueID"))
	return rv
}


// The unique identifier of the output device used to play audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/audiooutputdeviceuniqueid
func (s_ SampleBufferAudioRenderer) SetAudioOutputDeviceUniqueID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAudioOutputDeviceUniqueID:"), value)
}


// The processing algorithm used to manage audio pitch at different rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/audiotimepitchalgorithm
func (s_ SampleBufferAudioRenderer) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* not a class type */ {
	rv := objc.Send[AudioTimePitchAlgorithm](s_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}


// The processing algorithm used to manage audio pitch at different rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/audiotimepitchalgorithm
func (s_ SampleBufferAudioRenderer) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}


// The error that caused the renderer to no longer render sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/error
func (s_ SampleBufferAudioRenderer) Error() coretelephony.Error {
	rv := objc.Send[coretelephony.Error](s_.ID, objc.Sel("error"))
	return rv
}


// The error that caused the renderer to no longer render sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/error
func (s_ SampleBufferAudioRenderer) SetError(value coretelephony.Error) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setError:"), value)
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


// The status of the audio renderer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/status
func (s_ SampleBufferAudioRenderer) Status() QueuedSampleBufferRenderingStatus /* not a class type */ {
	rv := objc.Send[QueuedSampleBufferRenderingStatus](s_.ID, objc.Sel("status"))
	return rv
}


// The status of the audio renderer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/status
func (s_ SampleBufferAudioRenderer) SetStatus(value QueuedSampleBufferRenderingStatus /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStatus:"), value)
}


// The current audio volume for the audio renderer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/volume
func (s_ SampleBufferAudioRenderer) Volume() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("volume"))
	return rv
}


// The current audio volume for the audio renderer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/volume
func (s_ SampleBufferAudioRenderer) SetVolume(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVolume:"), value)
}


// The key that indicates the presentation timestamp of the first queued sample that was flushed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorendererflushtimekey
func (s_ SampleBufferAudioRenderer) AVSampleBufferAudioRendererFlushTimeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("AVSampleBufferAudioRendererFlushTimeKey"))
	return rv
}



