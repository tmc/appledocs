// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ImageRequestHandler] class.
var (
	ImageRequestHandlerClass     _ImageRequestHandlerClass
	ImageRequestHandlerClassOnce sync.Once
)

func getImageRequestHandlerClass() _ImageRequestHandlerClass {
	ImageRequestHandlerClassOnce.Do(func() {
		ImageRequestHandlerClass = _ImageRequestHandlerClass{objc.GetClass("VNImageRequestHandler")}
	})
	return ImageRequestHandlerClass
}

type _ImageRequestHandlerClass struct {
	class objc.Class
}

// An interface definition for the [ImageRequestHandler] class.
type IImageRequestHandler interface {
	objectivec.IObject
	PerformRequestsError(requests unsafe.Pointer, error_ unsafe.Pointer) bool
}

// An object that processes one or more image-analysis request pertaining to a single image.
//
// Instantiate this handler to perform Vision requests on a single image. You specify the image and, optionally, a completion handler at the time of creation, and call to begin executing the request.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler
type ImageRequestHandler struct {
	objectivec.Object
}

// ImageRequestHandlerFrom constructs a [ImageRequestHandler] from an unsafe.Pointer.
//
// An object that processes one or more image-analysis request pertaining to a single image.
func ImageRequestHandlerFrom(ptr unsafe.Pointer) ImageRequestHandler {
	return ImageRequestHandler{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageRequestHandlerClass) Alloc() ImageRequestHandler {
	rv := objc.Send[ImageRequestHandler](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageRequestHandlerClass) New() ImageRequestHandler {
	rv := objc.Send[ImageRequestHandler](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageRequestHandler) Init() ImageRequestHandler {
	rv := objc.Send[ImageRequestHandler](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageRequestHandler) Autorelease() ImageRequestHandler {
	rv := objc.Send[ImageRequestHandler](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageRequestHandler creates a new ImageRequestHandler instance.
func NewImageRequestHandler() ImageRequestHandler {
	return getImageRequestHandlerClass().New()
}


// Creates a request handler that performs requests on an image of a specified orientation contained within a sample buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cmSampleBuffer:orientation:options:)
func NewImageRequestHandlerWithCMSampleBufferOrientationOptions(sampleBuffer unsafe.Pointer, orientation unsafe.Pointer, options unsafe.Pointer) ImageRequestHandler {
	instance := getImageRequestHandlerClass().Alloc()
	rv := objc.Send[ImageRequestHandler](instance.ID, objc.Sel("initWithCMSampleBuffer:orientation:options:"), sampleBuffer, orientation, options)
	rv.Autorelease()
	return rv
}

// Creates a handler for performing requests on a Core Video pixel buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cvPixelBuffer:options:)
func NewImageRequestHandlerWithCVPixelBufferOptions(pixelBuffer unsafe.Pointer, options unsafe.Pointer) ImageRequestHandler {
	instance := getImageRequestHandlerClass().Alloc()
	rv := objc.Send[ImageRequestHandler](instance.ID, objc.Sel("initWithCVPixelBuffer:options:"), pixelBuffer, options)
	rv.Autorelease()
	return rv
}

// Creates a handler to be used for performing requests on an image with known orientation, at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(url:orientation:options:)
func NewImageRequestHandlerWithURLOrientationOptions(imageURL unsafe.Pointer, orientation unsafe.Pointer, options unsafe.Pointer) ImageRequestHandler {
	instance := getImageRequestHandlerClass().Alloc()
	rv := objc.Send[ImageRequestHandler](instance.ID, objc.Sel("initWithURL:orientation:options:"), imageURL, orientation, options)
	rv.Autorelease()
	return rv
}

// Creates a handler to be used for performing requests on Core Graphics images.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cgImage:options:)
func NewImageRequestHandlerWithCGImageOptions(image CGImageRef, options unsafe.Pointer) ImageRequestHandler {
	instance := getImageRequestHandlerClass().Alloc()
	rv := objc.Send[ImageRequestHandler](instance.ID, objc.Sel("initWithCGImage:options:"), image, options)
	rv.Autorelease()
	return rv
}

// Creates a handler to be used for performing requests on a Core Graphics image with known orientation.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cgImage:orientation:options:)
func NewImageRequestHandlerWithCGImageOrientationOptions(image CGImageRef, orientation unsafe.Pointer, options unsafe.Pointer) ImageRequestHandler {
	instance := getImageRequestHandlerClass().Alloc()
	rv := objc.Send[ImageRequestHandler](instance.ID, objc.Sel("initWithCGImage:orientation:options:"), image, orientation, options)
	rv.Autorelease()
	return rv
}

// Creates a handler to use for performing requests on Core Image image data.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(ciImage:options:)
func NewImageRequestHandlerWithCIImageOptions(image unsafe.Pointer, options unsafe.Pointer) ImageRequestHandler {
	instance := getImageRequestHandlerClass().Alloc()
	rv := objc.Send[ImageRequestHandler](instance.ID, objc.Sel("initWithCIImage:options:"), image, options)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cvPixelBuffer:depthData:orientation:options:)
func NewImageRequestHandlerWithCVPixelBufferDepthDataOrientationOptions(pixelBuffer unsafe.Pointer, depthData unsafe.Pointer, orientation unsafe.Pointer, options unsafe.Pointer) ImageRequestHandler {
	instance := getImageRequestHandlerClass().Alloc()
	rv := objc.Send[ImageRequestHandler](instance.ID, objc.Sel("initWithCVPixelBuffer:depthData:orientation:options:"), pixelBuffer, depthData, orientation, options)
	rv.Autorelease()
	return rv
}

// Creates a handler for performing requests on a Core Video pixel buffer of a known orientation.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cvPixelBuffer:orientation:options:)
func NewImageRequestHandlerWithCVPixelBufferOrientationOptions(pixelBuffer unsafe.Pointer, orientation unsafe.Pointer, options unsafe.Pointer) ImageRequestHandler {
	instance := getImageRequestHandlerClass().Alloc()
	rv := objc.Send[ImageRequestHandler](instance.ID, objc.Sel("initWithCVPixelBuffer:orientation:options:"), pixelBuffer, orientation, options)
	rv.Autorelease()
	return rv
}

// Creates a handler to use for performing requests on an image in a data object.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(data:options:)
func NewImageRequestHandlerWithDataOptions(imageData unsafe.Pointer, options unsafe.Pointer) ImageRequestHandler {
	instance := getImageRequestHandlerClass().Alloc()
	rv := objc.Send[ImageRequestHandler](instance.ID, objc.Sel("initWithData:options:"), imageData, options)
	rv.Autorelease()
	return rv
}

// Creates a handler to use for performing requests on an image of known orientation.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(data:orientation:options:)
func NewImageRequestHandlerWithDataOrientationOptions(imageData unsafe.Pointer, orientation unsafe.Pointer, options unsafe.Pointer) ImageRequestHandler {
	instance := getImageRequestHandlerClass().Alloc()
	rv := objc.Send[ImageRequestHandler](instance.ID, objc.Sel("initWithData:orientation:options:"), imageData, orientation, options)
	rv.Autorelease()
	return rv
}

// Creates a handler to be used for performing requests on an image at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(url:options:)
func NewImageRequestHandlerWithURLOptions(imageURL unsafe.Pointer, options unsafe.Pointer) ImageRequestHandler {
	instance := getImageRequestHandlerClass().Alloc()
	rv := objc.Send[ImageRequestHandler](instance.ID, objc.Sel("initWithURL:options:"), imageURL, options)
	rv.Autorelease()
	return rv
}

// Creates a handler to be used for performing requests on Core Image image data of a known orientation.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(ciImage:orientation:options:)
func NewImageRequestHandlerWithCIImageOrientationOptions(image unsafe.Pointer, orientation unsafe.Pointer, options unsafe.Pointer) ImageRequestHandler {
	instance := getImageRequestHandlerClass().Alloc()
	rv := objc.Send[ImageRequestHandler](instance.ID, objc.Sel("initWithCIImage:orientation:options:"), image, orientation, options)
	rv.Autorelease()
	return rv
}

// Creates a request handler that performs requests on an image in a sample buffer that contains depth data.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cmSampleBuffer:depthData:orientation:options:)
func NewImageRequestHandlerWithCMSampleBufferDepthDataOrientationOptions(sampleBuffer unsafe.Pointer, depthData unsafe.Pointer, orientation unsafe.Pointer, options unsafe.Pointer) ImageRequestHandler {
	instance := getImageRequestHandlerClass().Alloc()
	rv := objc.Send[ImageRequestHandler](instance.ID, objc.Sel("initWithCMSampleBuffer:depthData:orientation:options:"), sampleBuffer, depthData, orientation, options)
	rv.Autorelease()
	return rv
}

// Creates a request handler that performs requests on an image contained within a sample buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(cmSampleBuffer:options:)
func NewImageRequestHandlerWithCMSampleBufferOptions(sampleBuffer unsafe.Pointer, options unsafe.Pointer) ImageRequestHandler {
	instance := getImageRequestHandlerClass().Alloc()
	rv := objc.Send[ImageRequestHandler](instance.ID, objc.Sel("initWithCMSampleBuffer:options:"), sampleBuffer, options)
	rv.Autorelease()
	return rv
}


// Schedules Vision requests to perform.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/perform(_:)
func (i_ ImageRequestHandler) PerformRequestsError(requests unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("performRequests:error:"), requests, error_)
	return rv
}


