// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CoreMLRequest] class.
var (
	CoreMLRequestClass     _CoreMLRequestClass
	CoreMLRequestClassOnce sync.Once
)

func getCoreMLRequestClass() _CoreMLRequestClass {
	CoreMLRequestClassOnce.Do(func() {
		CoreMLRequestClass = _CoreMLRequestClass{objc.GetClass("VNCoreMLRequest")}
	})
	return CoreMLRequestClass
}

type _CoreMLRequestClass struct {
	class objc.Class
}

// An interface definition for the [CoreMLRequest] class.
type ICoreMLRequest interface {
	IImageBasedRequest
}

// An image-analysis request that uses a Core ML model to process images.
//
// The results array of a Core ML-based image analysis request contains a different observation type, depending on the kind of object you use: If the model predicts a single feature, the model’s object has a non- value for and Vision treats the model as a classifier. The results are objects. If the model’s outputs include at least one output with a feature type of , Vision treats that model as an image-to-image model. The results are objects. Otherwise, Vision treats the model as a general predictor model. The results are objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLRequest
type CoreMLRequest struct {
	ImageBasedRequest
}

// CoreMLRequestFrom constructs a [CoreMLRequest] from an unsafe.Pointer.
//
// An image-analysis request that uses a Core ML model to process images.
func CoreMLRequestFrom(ptr unsafe.Pointer) CoreMLRequest {
	return CoreMLRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CoreMLRequestClass) Alloc() CoreMLRequest {
	rv := objc.Send[CoreMLRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CoreMLRequestClass) New() CoreMLRequest {
	rv := objc.Send[CoreMLRequest](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CoreMLRequest) Init() CoreMLRequest {
	rv := objc.Send[CoreMLRequest](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CoreMLRequest) Autorelease() CoreMLRequest {
	rv := objc.Send[CoreMLRequest](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLRequest creates a new CoreMLRequest instance.
func NewCoreMLRequest() CoreMLRequest {
	return getCoreMLRequestClass().New()
}


// Creates a model container to use with an image analysis request based on the model you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/init(model:)
func NewCoreMLRequestWithModel(model unsafe.Pointer) CoreMLRequest {
	instance := getCoreMLRequestClass().Alloc()
	rv := objc.Send[CoreMLRequest](instance.ID, objc.Sel("initWithModel:"), model)
	rv.Autorelease()
	return rv
}

// Creates a model container to use with an image analysis request based on the model you provide, with an optional completion handler.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/init(model:completionHandler:)
func NewCoreMLRequestWithModelCompletionHandler(model unsafe.Pointer, completionHandler unsafe.Pointer) CoreMLRequest {
	instance := getCoreMLRequestClass().Alloc()
	rv := objc.Send[CoreMLRequest](instance.ID, objc.Sel("initWithModel:completionHandler:"), model, completionHandler)
	rv.Autorelease()
	return rv
}


// An optional setting that tells the Vision algorithm how to scale an input image.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/imageCropAndScaleOption
func (c_ CoreMLRequest) ImageCropAndScaleOption() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("imageCropAndScaleOption"))
	return rv
}


// SetImageCropAndScaleOption sets the value of the imageCropAndScaleOption property.
// An optional setting that tells the Vision algorithm how to scale an input image.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/imageCropAndScaleOption
func (c_ CoreMLRequest) SetImageCropAndScaleOption(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImageCropAndScaleOption:"), value)
}
// The model to base the image analysis request on.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/model
func (c_ CoreMLRequest) Model() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("model"))
	return rv
}


