// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVSampleBufferAudioRenderer */


/* debug [class_header]: Header for AVSampleBufferAudioRenderer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SampleBufferAudioRenderer */
// An interface definition for the [SampleBufferAudioRenderer] class.
type ISampleBufferAudioRenderer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SampleBufferAudioRenderer */
	// properties:
	AllowedAudioSpatializationFormats() AudioSpatializationFormats
	SetAllowedAudioSpatializationFormats(value AudioSpatializationFormats)
	AudioOutputDeviceUniqueID() objc.IObject /* cross-framework: NSString */
	SetAudioOutputDeviceUniqueID(value objc.IObject /* cross-framework: NSString */)
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* typedef */
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* typedef */)
	Error() Error
	Muted() bool
	SetMuted(value bool)
	Status() QueuedSampleBufferRenderingStatus
	Volume() float32
	SetVolume(value float32)
	IsMuted() bool
	SetIsMuted(value bool)
	AVSampleBufferAudioRendererFlushTimeKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SampleBufferAudioRenderer */
	// methods:
	FlushFromSourceTimeCompletionHandler(time objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SampleBufferAudioRenderer */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SampleBufferAudioRenderer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SampleBufferAudioRenderer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SampleBufferAudioRenderer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SampleBufferAudioRenderer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SampleBufferAudioRenderer */

// Flushes queued sample buffers with presentation time stamps later than or equal to the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/flush(fromSourceTime:completionHandler:)
func (s_ SampleBufferAudioRenderer) FlushFromSourceTimeCompletionHandler(time objc.IObject /* cross-framework: Time */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("flushFromSourceTime:completionHandler:"), time, completionHandler)
}/* debug [instance_methods/method]: FlushFromSourceTimeCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SampleBufferAudioRenderer */

// The source audio channel layouts the audio renderer supports for spatialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/allowedAudioSpatializationFormats
func (s_ SampleBufferAudioRenderer) AllowedAudioSpatializationFormats() AudioSpatializationFormats {
	rv := objc.Send[AudioSpatializationFormats](s_.ID, objc.Sel("allowedAudioSpatializationFormats"))
	return rv
}/* debug [instance_properties/getter]: allowedAudioSpatializationFormats */


// The source audio channel layouts the audio renderer supports for spatialization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/allowedAudioSpatializationFormats
func (s_ SampleBufferAudioRenderer) SetAllowedAudioSpatializationFormats(value AudioSpatializationFormats) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowedAudioSpatializationFormats:"), value)
}/* debug [instance_properties/setter]: allowedAudioSpatializationFormats */


// The unique identifier of the output device used to play audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/audioOutputDeviceUniqueID
func (s_ SampleBufferAudioRenderer) AudioOutputDeviceUniqueID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("audioOutputDeviceUniqueID"))
	return rv
}/* debug [instance_properties/getter]: audioOutputDeviceUniqueID */


// The unique identifier of the output device used to play audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/audioOutputDeviceUniqueID
func (s_ SampleBufferAudioRenderer) SetAudioOutputDeviceUniqueID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAudioOutputDeviceUniqueID:"), value)
}/* debug [instance_properties/setter]: audioOutputDeviceUniqueID */


// The processing algorithm used to manage audio pitch at different rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/audioTimePitchAlgorithm
func (s_ SampleBufferAudioRenderer) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* typedef */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}/* debug [instance_properties/getter]: audioTimePitchAlgorithm */


// The processing algorithm used to manage audio pitch at different rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/audioTimePitchAlgorithm
func (s_ SampleBufferAudioRenderer) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* typedef */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}/* debug [instance_properties/setter]: audioTimePitchAlgorithm */


// The error that caused the renderer to no longer render sample buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/error
func (s_ SampleBufferAudioRenderer) Error() Error {
	rv := objc.Send[Error](s_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// A Boolean value that indicates whether audio for the renderer is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/isMuted
func (s_ SampleBufferAudioRenderer) Muted() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("muted"))
	return rv
}/* debug [instance_properties/getter]: muted */


// A Boolean value that indicates whether audio for the renderer is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/isMuted
func (s_ SampleBufferAudioRenderer) SetMuted(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMuted:"), value)
}/* debug [instance_properties/setter]: muted */


// The status of the audio renderer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/status
func (s_ SampleBufferAudioRenderer) Status() QueuedSampleBufferRenderingStatus {
	rv := objc.Send[QueuedSampleBufferRenderingStatus](s_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// The current audio volume for the audio renderer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/volume
func (s_ SampleBufferAudioRenderer) Volume() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("volume"))
	return rv
}/* debug [instance_properties/getter]: volume */


// The current audio volume for the audio renderer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/volume
func (s_ SampleBufferAudioRenderer) SetVolume(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVolume:"), value)
}/* debug [instance_properties/setter]: volume */


// A Boolean value that indicates whether audio for the renderer is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/ismuted
func (s_ SampleBufferAudioRenderer) IsMuted() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isMuted"))
	return rv
}/* debug [instance_properties/getter]: isMuted */


// A Boolean value that indicates whether audio for the renderer is in a muted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorenderer/ismuted
func (s_ SampleBufferAudioRenderer) SetIsMuted(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsMuted:"), value)
}/* debug [instance_properties/setter]: isMuted */


// The key that indicates the presentation timestamp of the first queued sample that was flushed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorendererflushtimekey
func (s_ SampleBufferAudioRenderer) AVSampleBufferAudioRendererFlushTimeKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("AVSampleBufferAudioRendererFlushTimeKey"))
	return rv
}/* debug [instance_properties/getter]: AVSampleBufferAudioRendererFlushTimeKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVSampleBufferAudioRenderer */



