// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [SequenceRequestHandler] class.
var (
	SequenceRequestHandlerClass     _SequenceRequestHandlerClass
	SequenceRequestHandlerClassOnce sync.Once
)

func getSequenceRequestHandlerClass() _SequenceRequestHandlerClass {
	SequenceRequestHandlerClassOnce.Do(func() {
		SequenceRequestHandlerClass = _SequenceRequestHandlerClass{objc.GetClass("VNSequenceRequestHandler")}
	})
	return SequenceRequestHandlerClass
}

type _SequenceRequestHandlerClass struct {
	class objc.Class
}





// An interface definition for the [SequenceRequestHandler] class.
type ISequenceRequestHandler interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	PerformRequestsOnCVPixelBufferError(requests []Request, pixelBuffer PixelBufferRef /* not a class type */, error_ foundation.foundation.INSError) bool
	PerformRequestsOnCGImageError(requests []Request, image ImageRef /* not a class type */, error_ foundation.foundation.INSError) bool
	PerformRequestsOnCMSampleBufferError(requests []Request, sampleBuffer SampleBufferRef /* not a class type */, error_ foundation.foundation.INSError) bool
	PerformRequestsOnCIImageError(requests []Request, image objectivec.IObject, error_ foundation.foundation.INSError) bool
	PerformRequestsOnCIImageOrientationError(requests []Request, image objectivec.IObject, orientation ImagePropertyOrientation /* not a class type */, error_ foundation.foundation.INSError) bool
	PerformRequestsOnCVPixelBufferOrientationError(requests []Request, pixelBuffer PixelBufferRef /* not a class type */, orientation ImagePropertyOrientation /* not a class type */, error_ foundation.foundation.INSError) bool
	PerformRequestsOnCGImageOrientationError(requests []Request, image ImageRef /* not a class type */, orientation ImagePropertyOrientation /* not a class type */, error_ foundation.foundation.INSError) bool
	PerformRequestsOnCMSampleBufferOrientationError(requests []Request, sampleBuffer SampleBufferRef /* not a class type */, orientation ImagePropertyOrientation /* not a class type */, error_ foundation.foundation.INSError) bool
	PerformRequestsOnImageDataError(requests []Request, imageData foundation.foundation.INSData, error_ foundation.foundation.INSError) bool
	PerformRequestsOnImageDataOrientationError(requests []Request, imageData foundation.foundation.INSData, orientation ImagePropertyOrientation /* not a class type */, error_ foundation.foundation.INSError) bool
	PerformRequestsOnImageURLError(requests []Request, imageURL foundation.foundation.INSURL, error_ foundation.foundation.INSError) bool
	PerformRequestsOnImageURLOrientationError(requests []Request, imageURL foundation.foundation.INSURL, orientation ImagePropertyOrientation /* not a class type */, error_ foundation.foundation.INSError) bool


}





// Alloc allocates a new instance without initialization.
func (sc _SequenceRequestHandlerClass) Alloc() SequenceRequestHandler {
	rv := objc.Send[SequenceRequestHandler](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SequenceRequestHandlerClass) New() SequenceRequestHandler {
	rv := objc.Send[SequenceRequestHandler](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SequenceRequestHandler) Init() SequenceRequestHandler {
	rv := objc.Send[SequenceRequestHandler](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SequenceRequestHandler) Autorelease() SequenceRequestHandler {
	rv := objc.Send[SequenceRequestHandler](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSequenceRequestHandler creates a new SequenceRequestHandler instance.
func NewSequenceRequestHandler() SequenceRequestHandler {
	return getSequenceRequestHandlerClass().New()
}





// An object that processes image-analysis requests for each frame in a sequence.
//
// Instantiate this handler to perform Vision requests on a series of images. Unlike the , you don’t specify the image on creation. Instead, you supply each image frame one by one as you continue to call one of the methods.


// An object that processes image-analysis requests for each frame in a sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler
type SequenceRequestHandler struct {
	objectivec.Object
}

// SequenceRequestHandlerFrom constructs a [SequenceRequestHandler] from an unsafe.Pointer.
//
// An object that processes image-analysis requests for each frame in a sequence.
func SequenceRequestHandlerFrom(ptr unsafe.Pointer) SequenceRequestHandler {
	return SequenceRequestHandler{objectivec.Object{objc.ID(ptr)}}
}





















// Schedules one or more Vision requests to be performed on a Core Video pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:on:)-3d7nt
func (s_ SequenceRequestHandler) PerformRequestsOnCVPixelBufferError(requests []Request, pixelBuffer PixelBufferRef /* not a class type */, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("performRequests:onCVPixelBuffer:error:"), requests, pixelBuffer, error_)
	return rv
}


// Schedules Vision requests to be performed on a Core Graphics image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:on:)-3zt7l
func (s_ SequenceRequestHandler) PerformRequestsOnCGImageError(requests []Request, image ImageRef /* not a class type */, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("performRequests:onCGImage:error:"), requests, image, error_)
	return rv
}


// Performs one or more requests on an image contained within a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:on:)-45e73
func (s_ SequenceRequestHandler) PerformRequestsOnCMSampleBufferError(requests []Request, sampleBuffer SampleBufferRef /* not a class type */, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("performRequests:onCMSampleBuffer:error:"), requests, sampleBuffer, error_)
	return rv
}


// Schedules one or more Vision requests to be performed on Core Image image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:on:)-9jtgj
func (s_ SequenceRequestHandler) PerformRequestsOnCIImageError(requests []Request, image objectivec.IObject, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("performRequests:onCIImage:error:"), requests, image, error_)
	return rv
}


// Schedules one or more Vision requests to be performed on Core Image image data with known orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:on:orientation:)-1bkm1
func (s_ SequenceRequestHandler) PerformRequestsOnCIImageOrientationError(requests []Request, image objectivec.IObject, orientation ImagePropertyOrientation /* not a class type */, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("performRequests:onCIImage:orientation:error:"), requests, image, orientation, error_)
	return rv
}


// Schedules one or more Vision requests to be performed on a Core Video pixel buffer with known orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:on:orientation:)-2wvt8
func (s_ SequenceRequestHandler) PerformRequestsOnCVPixelBufferOrientationError(requests []Request, pixelBuffer PixelBufferRef /* not a class type */, orientation ImagePropertyOrientation /* not a class type */, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("performRequests:onCVPixelBuffer:orientation:error:"), requests, pixelBuffer, orientation, error_)
	return rv
}


// Schedules one or more Vision requests to be performed on a Core Graphics image with known orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:on:orientation:)-3gcmv
func (s_ SequenceRequestHandler) PerformRequestsOnCGImageOrientationError(requests []Request, image ImageRef /* not a class type */, orientation ImagePropertyOrientation /* not a class type */, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("performRequests:onCGImage:orientation:error:"), requests, image, orientation, error_)
	return rv
}


// Performs one or more requests on an image of a specified orientation contained within a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:on:orientation:)-6b7rk
func (s_ SequenceRequestHandler) PerformRequestsOnCMSampleBufferOrientationError(requests []Request, sampleBuffer SampleBufferRef /* not a class type */, orientation ImagePropertyOrientation /* not a class type */, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("performRequests:onCMSampleBuffer:orientation:error:"), requests, sampleBuffer, orientation, error_)
	return rv
}


// Schedules one or more Vision requests to be performed on raw image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:onImageData:)
func (s_ SequenceRequestHandler) PerformRequestsOnImageDataError(requests []Request, imageData foundation.foundation.INSData, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("performRequests:onImageData:error:"), requests, imageData, error_)
	return rv
}


// Schedules one or more Vision requests to be performed on raw data containing an image with known orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:onImageData:orientation:)
func (s_ SequenceRequestHandler) PerformRequestsOnImageDataOrientationError(requests []Request, imageData foundation.foundation.INSData, orientation ImagePropertyOrientation /* not a class type */, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("performRequests:onImageData:orientation:error:"), requests, imageData, orientation, error_)
	return rv
}


// Schedules one or more Vision requests to be performed on an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:onImageURL:)
func (s_ SequenceRequestHandler) PerformRequestsOnImageURLError(requests []Request, imageURL foundation.foundation.INSURL, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("performRequests:onImageURL:error:"), requests, imageURL, error_)
	return rv
}


// Schedules one or more Vision requests to be performed on an image with known orientation, at a specific URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNSequenceRequestHandler/perform(_:onImageURL:orientation:)
func (s_ SequenceRequestHandler) PerformRequestsOnImageURLOrientationError(requests []Request, imageURL foundation.foundation.INSURL, orientation ImagePropertyOrientation /* not a class type */, error_ foundation.foundation.INSError) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("performRequests:onImageURL:orientation:error:"), requests, imageURL, orientation, error_)
	return rv
}












