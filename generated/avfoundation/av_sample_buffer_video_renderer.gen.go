// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SampleBufferVideoRenderer] class.
var (
	SampleBufferVideoRendererClass     _SampleBufferVideoRendererClass
	SampleBufferVideoRendererClassOnce sync.Once
)

func getSampleBufferVideoRendererClass() _SampleBufferVideoRendererClass {
	SampleBufferVideoRendererClassOnce.Do(func() {
		SampleBufferVideoRendererClass = _SampleBufferVideoRendererClass{objc.GetClass("AVSampleBufferVideoRenderer")}
	})
	return SampleBufferVideoRendererClass
}

type _SampleBufferVideoRendererClass struct {
	class objc.Class
}

// An interface definition for the [SampleBufferVideoRenderer] class.
type ISampleBufferVideoRenderer interface {
	objectivec.IObject
	CopyDisplayedPixelBuffer() unsafe.Pointer
}

// An object that enqueues video sample buffers for rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer
type SampleBufferVideoRenderer struct {
	objectivec.Object
}

// SampleBufferVideoRendererFrom constructs a [SampleBufferVideoRenderer] from an unsafe.Pointer.
//
// An object that enqueues video sample buffers for rendering.
func SampleBufferVideoRendererFrom(ptr unsafe.Pointer) SampleBufferVideoRenderer {
	return SampleBufferVideoRenderer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SampleBufferVideoRendererClass) Alloc() SampleBufferVideoRenderer {
	rv := objc.Send[SampleBufferVideoRenderer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SampleBufferVideoRendererClass) New() SampleBufferVideoRenderer {
	rv := objc.Send[SampleBufferVideoRenderer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SampleBufferVideoRenderer) Init() SampleBufferVideoRenderer {
	rv := objc.Send[SampleBufferVideoRenderer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SampleBufferVideoRenderer) Autorelease() SampleBufferVideoRenderer {
	rv := objc.Send[SampleBufferVideoRenderer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSampleBufferVideoRenderer creates a new SampleBufferVideoRenderer instance.
func NewSampleBufferVideoRenderer() SampleBufferVideoRenderer {
	return getSampleBufferVideoRendererClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/displayedPixelBuffer()
func (s_ SampleBufferVideoRenderer) CopyDisplayedPixelBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("copyDisplayedPixelBuffer"))
	return rv
}

// An object the describes the error that caused the rendering failure.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffervideorenderer/error
func (s_ SampleBufferVideoRenderer) Error() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("error"))
	return rv
}


// SetError sets the value of the error property.
// An object the describes the error that caused the rendering failure.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffervideorenderer/error
func (s_ SampleBufferVideoRenderer) SetError(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setError:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffervideorenderer/presentationtimeexpectation-swift.property
func (s_ SampleBufferVideoRenderer) PresentationTimeExpectation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("presentationTimeExpectation"))
	return rv
}


// SetPresentationTimeExpectation sets the value of the presentationTimeExpectation property.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffervideorenderer/presentationtimeexpectation-swift.property
func (s_ SampleBufferVideoRenderer) SetPresentationTimeExpectation(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPresentationTimeExpectation:"), value)
}

// Recommended pixel buffer attributes for optimal performance when using CMSampleBuffers containing CVPixelbuffers.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffervideorenderer/recommendedpixelbufferattributes-6zrqb
func (s_ SampleBufferVideoRenderer) RecommendedPixelBufferAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("recommendedPixelBufferAttributes"))
	return rv
}


// SetRecommendedPixelBufferAttributes sets the value of the recommendedPixelBufferAttributes property.
// Recommended pixel buffer attributes for optimal performance when using CMSampleBuffers containing CVPixelbuffers.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffervideorenderer/recommendedpixelbufferattributes-6zrqb
func (s_ SampleBufferVideoRenderer) SetRecommendedPixelBufferAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRecommendedPixelBufferAttributes:"), value)
}

// A Boolean value that Indicates whether the renderer requires flushing to continue decoding frames.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffervideorenderer/requiresflushtoresumedecoding
func (s_ SampleBufferVideoRenderer) RequiresFlushToResumeDecoding() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("requiresFlushToResumeDecoding"))
	return rv
}


// SetRequiresFlushToResumeDecoding sets the value of the requiresFlushToResumeDecoding property.
// A Boolean value that Indicates whether the renderer requires flushing to continue decoding frames.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffervideorenderer/requiresflushtoresumedecoding
func (s_ SampleBufferVideoRenderer) SetRequiresFlushToResumeDecoding(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRequiresFlushToResumeDecoding:"), value)
}

// A status value that indicates whether this object can enqueue and render sample buffers.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffervideorenderer/status
func (s_ SampleBufferVideoRenderer) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
// A status value that indicates whether this object can enqueue and render sample buffers.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebuffervideorenderer/status
func (s_ SampleBufferVideoRenderer) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStatus:"), value)
}



