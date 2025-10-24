// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AsynchronousVideoCompositionRequest] class.
var (
	AsynchronousVideoCompositionRequestClass     _AsynchronousVideoCompositionRequestClass
	AsynchronousVideoCompositionRequestClassOnce sync.Once
)

func getAsynchronousVideoCompositionRequestClass() _AsynchronousVideoCompositionRequestClass {
	AsynchronousVideoCompositionRequestClassOnce.Do(func() {
		AsynchronousVideoCompositionRequestClass = _AsynchronousVideoCompositionRequestClass{objc.GetClass("AVAsynchronousVideoCompositionRequest")}
	})
	return AsynchronousVideoCompositionRequestClass
}

type _AsynchronousVideoCompositionRequestClass struct {
	class objc.Class
}





// An interface definition for the [AsynchronousVideoCompositionRequest] class.
type IAsynchronousVideoCompositionRequest interface {
	objectivec.IObject
	

	// properties:
	CompositionTime() objc.IObject /* cross-framework: Time */
	RenderContext() IAVVideoCompositionRenderContext
	SourceSampleDataTrackIDs() []foundation.Number
	SourceTrackIDs() []foundation.Number
	VideoCompositionInstruction() unsafe.Pointer


	

	// methods:
	AttachSpatialVideoConfigurationToPixelBuffer(spatialVideoConfiguration IAVSpatialVideoConfiguration, pixelBuffer PixelBufferRef /* not a class type */)
	FinishWithError(error_ Error)
	FinishWithComposedTaggedBufferGroup(taggedBufferGroup TaggedBufferGroupRef /* not a class type */)
	FinishCancelledRequest()
	SourceTaggedBufferGroupByTrackID(trackID PersistentTrackID /* not a class type */) TaggedBufferGroupRef /* not a class type */
	SourceTimedMetadataByTrackID(trackID PersistentTrackID /* not a class type */) ITimedMetadataGroup


}





// Alloc allocates a new instance without initialization.
func (ac _AsynchronousVideoCompositionRequestClass) Alloc() AsynchronousVideoCompositionRequest {
	rv := objc.Send[AsynchronousVideoCompositionRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AsynchronousVideoCompositionRequestClass) New() AsynchronousVideoCompositionRequest {
	rv := objc.Send[AsynchronousVideoCompositionRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AsynchronousVideoCompositionRequest) Init() AsynchronousVideoCompositionRequest {
	rv := objc.Send[AsynchronousVideoCompositionRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AsynchronousVideoCompositionRequest) Autorelease() AsynchronousVideoCompositionRequest {
	rv := objc.Send[AsynchronousVideoCompositionRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAsynchronousVideoCompositionRequest creates a new AsynchronousVideoCompositionRequest instance.
func NewAsynchronousVideoCompositionRequest() AsynchronousVideoCompositionRequest {
	return getAsynchronousVideoCompositionRequestClass().New()
}





// An object that contains information a video compositor needs to render an output pixel buffer.
//
// The video compositor must adopt the protocol.


// An object that contains information a video compositor needs to render an output pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest
type AsynchronousVideoCompositionRequest struct {
	objectivec.Object
}

// AsynchronousVideoCompositionRequestFrom constructs a [AsynchronousVideoCompositionRequest] from an unsafe.Pointer.
//
// An object that contains information a video compositor needs to render an output pixel buffer.
func AsynchronousVideoCompositionRequestFrom(ptr unsafe.Pointer) AsynchronousVideoCompositionRequest {
	return AsynchronousVideoCompositionRequest{objectivec.Object{objc.ID(ptr)}}
}




















// Associates the pixel buffer with the specified spatial configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/attachSpatialVideoConfiguration:toPixelBuffer:
func (a_ AsynchronousVideoCompositionRequest) AttachSpatialVideoConfigurationToPixelBuffer(spatialVideoConfiguration IAVSpatialVideoConfiguration, pixelBuffer PixelBufferRef /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("attachSpatialVideoConfiguration:toPixelBuffer:"), spatialVideoConfiguration, pixelBuffer)
}


// Finishes the request with an error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/finish(with:)
func (a_ AsynchronousVideoCompositionRequest) FinishWithError(error_ Error) {
	objc.Send[objc.ID](a_.ID, objc.Sel("finishWithError:"), error_)
}


// The method that the custom compositor calls when composition succeeds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/finishWithComposedTaggedBufferGroup:
func (a_ AsynchronousVideoCompositionRequest) FinishWithComposedTaggedBufferGroup(taggedBufferGroup TaggedBufferGroupRef /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("finishWithComposedTaggedBufferGroup:"), taggedBufferGroup)
}


// Cancels the request to compose a video frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/finishCancelledRequest()
func (a_ AsynchronousVideoCompositionRequest) FinishCancelledRequest() {
	objc.Send[objc.ID](a_.ID, objc.Sel("finishCancelledRequest"))
}


// Returns the source CMTaggedBufferGroupRef for the given track ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/sourceTaggedBufferGroupByTrackID:
func (a_ AsynchronousVideoCompositionRequest) SourceTaggedBufferGroupByTrackID(trackID PersistentTrackID /* not a class type */) TaggedBufferGroupRef /* not a class type */ {
	rv := objc.Send[TaggedBufferGroupRef](a_.ID, objc.Sel("sourceTaggedBufferGroupByTrackID:"), trackID)
	return rv
}


// Returns a source timed metadata group for the track that contains the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/sourceTimedMetadata(byTrackID:)
func (a_ AsynchronousVideoCompositionRequest) SourceTimedMetadataByTrackID(trackID PersistentTrackID /* not a class type */) ITimedMetadataGroup {
	rv := objc.Send[TimedMetadataGroup](a_.ID, objc.Sel("sourceTimedMetadataByTrackID:"), trackID)
	return rv
}







// A time for which to compose the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/compositionTime
func (a_ AsynchronousVideoCompositionRequest) CompositionTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](a_.ID, objc.Sel("compositionTime"))
	return rv
}


// The rendering context of the video composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/renderContext
func (a_ AsynchronousVideoCompositionRequest) RenderContext() IAVVideoCompositionRenderContext {
	rv := objc.Send[VideoCompositionRenderContext](a_.ID, objc.Sel("renderContext"))
	return rv
}


// The identifiers of tracks that contain source metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/sourceSampleDataTrackIDs-9vxz5
func (a_ AsynchronousVideoCompositionRequest) SourceSampleDataTrackIDs() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("sourceSampleDataTrackIDs"))
	return rv
}


// The identifiers of tracks that contain source video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/sourceTrackIDs
func (a_ AsynchronousVideoCompositionRequest) SourceTrackIDs() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("sourceTrackIDs"))
	return rv
}


// A video composition instruction that indicates how to compose the frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/videoCompositionInstruction
func (a_ AsynchronousVideoCompositionRequest) VideoCompositionInstruction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("videoCompositionInstruction"))
	return rv
}








