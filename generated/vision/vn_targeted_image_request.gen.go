// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [TargetedImageRequest] class.
var (
	TargetedImageRequestClass     _TargetedImageRequestClass
	TargetedImageRequestClassOnce sync.Once
)

func getTargetedImageRequestClass() _TargetedImageRequestClass {
	TargetedImageRequestClassOnce.Do(func() {
		TargetedImageRequestClass = _TargetedImageRequestClass{objc.GetClass("VNTargetedImageRequest")}
	})
	return TargetedImageRequestClass
}

type _TargetedImageRequestClass struct {
	class objc.Class
}





// An interface definition for the [TargetedImageRequest] class.
type ITargetedImageRequest interface {
	IImageBasedRequest
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (tc _TargetedImageRequestClass) Alloc() TargetedImageRequest {
	rv := objc.Send[TargetedImageRequest](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TargetedImageRequestClass) New() TargetedImageRequest {
	rv := objc.Send[TargetedImageRequest](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TargetedImageRequest) Init() TargetedImageRequest {
	rv := objc.Send[TargetedImageRequest](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TargetedImageRequest) Autorelease() TargetedImageRequest {
	rv := objc.Send[TargetedImageRequest](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTargetedImageRequest creates a new TargetedImageRequest instance.
func NewTargetedImageRequest() TargetedImageRequest {
	return getTargetedImageRequestClass().New()
}





// The abstract superclass for image analysis requests that operate on both the processed image and a secondary image.
//
// Other Vision request handlers that operate on both the processed image and a secondary image inherit from this abstract base class. Instantiate one of its subclasses to perform image analysis, and pass in auxiliary image data by filling in the dictionary at initialization.


// The abstract superclass for image analysis requests that operate on both the processed image and a secondary image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest
type TargetedImageRequest struct {
	ImageBasedRequest
}

// TargetedImageRequestFrom constructs a [TargetedImageRequest] from an unsafe.Pointer.
//
// The abstract superclass for image analysis requests that operate on both the processed image and a secondary image.
func TargetedImageRequestFrom(ptr unsafe.Pointer) TargetedImageRequest {
	return TargetedImageRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}






// Creates a new request targeting a Core Graphics image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCGImage:options:
func NewTargetedImageRequestWithTargetedCGImageOptions(cgImage ImageRef /* not a class type */, options foundation.IDictionary) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedCGImage:options:"), cgImage, options)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting a Core Graphics image, executing the completion handler when done.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCGImage:options:completionHandler:)
func NewTargetedImageRequestWithTargetedCGImageOptionsCompletionHandler(cgImage ImageRef /* not a class type */, options foundation.IDictionary, completionHandler RequestCompletionHandler /* not a class type */) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedCGImage:options:completionHandler:"), cgImage, options, completionHandler)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting a Core Graphics image of known orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCGImage:orientation:options:
func NewTargetedImageRequestWithTargetedCGImageOrientationOptions(cgImage ImageRef /* not a class type */, orientation ImagePropertyOrientation /* not a class type */, options foundation.IDictionary) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedCGImage:orientation:options:"), cgImage, orientation, options)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting a Core Graphics image of known orientation, executing the completion handler when done.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCGImage:orientation:options:completionHandler:)
func NewTargetedImageRequestWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage ImageRef /* not a class type */, orientation ImagePropertyOrientation /* not a class type */, options foundation.IDictionary, completionHandler RequestCompletionHandler /* not a class type */) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedCGImage:orientation:options:completionHandler:"), cgImage, orientation, options, completionHandler)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting a Core Image image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCIImage:options:
func NewTargetedImageRequestWithTargetedCIImageOptions(ciImage objectivec.IObject, options foundation.IDictionary) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedCIImage:options:"), ciImage, options)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting a Core Image image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCIImage:options:completionHandler:)
func NewTargetedImageRequestWithTargetedCIImageOptionsCompletionHandler(ciImage objectivec.IObject, options foundation.IDictionary, completionHandler RequestCompletionHandler /* not a class type */) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedCIImage:options:completionHandler:"), ciImage, options, completionHandler)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting a Core Image image of known orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCIImage:orientation:options:
func NewTargetedImageRequestWithTargetedCIImageOrientationOptions(ciImage objectivec.IObject, orientation ImagePropertyOrientation /* not a class type */, options foundation.IDictionary) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedCIImage:orientation:options:"), ciImage, orientation, options)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting a Core Image image of known orientation, executing the completion handler when done.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCIImage:orientation:options:completionHandler:)
func NewTargetedImageRequestWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage objectivec.IObject, orientation ImagePropertyOrientation /* not a class type */, options foundation.IDictionary, completionHandler RequestCompletionHandler /* not a class type */) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedCIImage:orientation:options:completionHandler:"), ciImage, orientation, options, completionHandler)
	rv.Autorelease()
	return rv
}


// Creates a new request that targets an image in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCMSampleBuffer:options:
func NewTargetedImageRequestWithTargetedCMSampleBufferOptions(sampleBuffer SampleBufferRef /* not a class type */, options foundation.IDictionary) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:options:"), sampleBuffer, options)
	rv.Autorelease()
	return rv
}


// Creates a new request with a completion handler that targets an image in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCMSampleBuffer:options:completionHandler:)
func NewTargetedImageRequestWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer SampleBufferRef /* not a class type */, options foundation.IDictionary, completionHandler RequestCompletionHandler /* not a class type */) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:options:completionHandler:"), sampleBuffer, options, completionHandler)
	rv.Autorelease()
	return rv
}


// Creates a new request that targets an image of a known orientation in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCMSampleBuffer:orientation:options:
func NewTargetedImageRequestWithTargetedCMSampleBufferOrientationOptions(sampleBuffer SampleBufferRef /* not a class type */, orientation ImagePropertyOrientation /* not a class type */, options foundation.IDictionary) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:"), sampleBuffer, orientation, options)
	rv.Autorelease()
	return rv
}


// Creates a new request with a completion handler that targets an image of a known orientation in a sample buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCMSampleBuffer:orientation:options:completionHandler:)
func NewTargetedImageRequestWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer SampleBufferRef /* not a class type */, orientation ImagePropertyOrientation /* not a class type */, options foundation.IDictionary, completionHandler RequestCompletionHandler /* not a class type */) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:completionHandler:"), sampleBuffer, orientation, options, completionHandler)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting an image in a pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCVPixelBuffer:options:
func NewTargetedImageRequestWithTargetedCVPixelBufferOptions(pixelBuffer PixelBufferRef /* not a class type */, options foundation.IDictionary) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:options:"), pixelBuffer, options)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting an image in a pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCVPixelBuffer:options:completionHandler:)
func NewTargetedImageRequestWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer PixelBufferRef /* not a class type */, options foundation.IDictionary, completionHandler RequestCompletionHandler /* not a class type */) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:options:completionHandler:"), pixelBuffer, options, completionHandler)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting an image in a pixel buffer of known orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedCVPixelBuffer:orientation:options:
func NewTargetedImageRequestWithTargetedCVPixelBufferOrientationOptions(pixelBuffer PixelBufferRef /* not a class type */, orientation ImagePropertyOrientation /* not a class type */, options foundation.IDictionary) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:"), pixelBuffer, orientation, options)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting an image in a pixel buffer of known orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedCVPixelBuffer:orientation:options:completionHandler:)
func NewTargetedImageRequestWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer PixelBufferRef /* not a class type */, orientation ImagePropertyOrientation /* not a class type */, options foundation.IDictionary, completionHandler RequestCompletionHandler /* not a class type */) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:completionHandler:"), pixelBuffer, orientation, options, completionHandler)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting an image as raw data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageData:options:
func NewTargetedImageRequestWithTargetedImageDataOptions(imageData foundation.foundation.INSData, options foundation.IDictionary) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedImageData:options:"), imageData, options)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting an image as raw data, executing the completion handler when done.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedImageData:options:completionHandler:)
func NewTargetedImageRequestWithTargetedImageDataOptionsCompletionHandler(imageData foundation.foundation.INSData, options foundation.IDictionary, completionHandler RequestCompletionHandler /* not a class type */) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedImageData:options:completionHandler:"), imageData, options, completionHandler)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting a raw data image of known orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageData:orientation:options:
func NewTargetedImageRequestWithTargetedImageDataOrientationOptions(imageData foundation.foundation.INSData, orientation ImagePropertyOrientation /* not a class type */, options foundation.IDictionary) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedImageData:orientation:options:"), imageData, orientation, options)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting a raw data image of known orientation, executing the completion handler when done.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedImageData:orientation:options:completionHandler:)
func NewTargetedImageRequestWithTargetedImageDataOrientationOptionsCompletionHandler(imageData foundation.foundation.INSData, orientation ImagePropertyOrientation /* not a class type */, options foundation.IDictionary, completionHandler RequestCompletionHandler /* not a class type */) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedImageData:orientation:options:completionHandler:"), imageData, orientation, options, completionHandler)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting an image at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageURL:options:
func NewTargetedImageRequestWithTargetedImageURLOptions(imageURL foundation.foundation.INSURL, options foundation.IDictionary) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedImageURL:options:"), imageURL, options)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting an image at the specified URL, executing the completion handler when done.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedImageURL:options:completionHandler:)
func NewTargetedImageRequestWithTargetedImageURLOptionsCompletionHandler(imageURL foundation.foundation.INSURL, options foundation.IDictionary, completionHandler RequestCompletionHandler /* not a class type */) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedImageURL:options:completionHandler:"), imageURL, options, completionHandler)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting an image of known orientation, at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/initWithTargetedImageURL:orientation:options:
func NewTargetedImageRequestWithTargetedImageURLOrientationOptions(imageURL foundation.foundation.INSURL, orientation ImagePropertyOrientation /* not a class type */, options foundation.IDictionary) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedImageURL:orientation:options:"), imageURL, orientation, options)
	rv.Autorelease()
	return rv
}


// Creates a new request targeting an image of known orientation, at the specified URL, executing the completion handler when done.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTargetedImageRequest/init(targetedImageURL:orientation:options:completionHandler:)
func NewTargetedImageRequestWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL foundation.foundation.INSURL, orientation ImagePropertyOrientation /* not a class type */, options foundation.IDictionary, completionHandler RequestCompletionHandler /* not a class type */) TargetedImageRequest {
	instance := getTargetedImageRequestClass().Alloc()
	rv := objc.Send[TargetedImageRequest](instance.ID, objc.Sel("initWithTargetedImageURL:orientation:options:completionHandler:"), imageURL, orientation, options, completionHandler)
	rv.Autorelease()
	return rv
}



























