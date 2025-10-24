// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	// methods:
	PerformRequestsError(requests []IRequest, error_ unsafe.Pointer) bool
}

// An object that processes one or more image-analysis request pertaining to a single image.
//
// Instantiate this handler to perform Vision requests on a single image. You specify the image and, optionally, a completion handler at the time of creation, and call to begin executing the request.


// An object that processes one or more image-analysis request pertaining to a single image.
//
// [Full Topic]
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



// Creates a handler to be used for performing requests on an image with known orientation, at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/init(url:orientation:options:)
func NewImageRequestHandlerWithURLOrientationOptions(imageURL objc.IObject /* cross-framework: NSURL */, orientation ImagePropertyOrientation /* not a class type */, options foundation.IDictionary) ImageRequestHandler {
	instance := getImageRequestHandlerClass().Alloc()
	rv := objc.Send[ImageRequestHandler](instance.ID, objc.Sel("initWithURL:orientation:options:"), imageURL, orientation, options)
	rv.Autorelease()
	return rv
}



// Schedules Vision requests to perform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRequestHandler/perform(_:)
func (i_ ImageRequestHandler) PerformRequestsError(requests []IRequest, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("performRequests:error:"), requests, error_)
	return rv
}


