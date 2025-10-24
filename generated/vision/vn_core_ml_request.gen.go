// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	ImageCropAndScaleOption() ImageCropAndScaleOption /* not a class type */
	SetImageCropAndScaleOption(value ImageCropAndScaleOption /* not a class type */)
	ModelDescription() objc.IObject /* cross-framework: ModelDescription */
	SetModelDescription(value objc.IObject /* cross-framework: ModelDescription */)
	PredictedFeatureName() objc.IObject /* cross-framework: NSString */
	SetPredictedFeatureName(value objc.IObject /* cross-framework: NSString */)
	Model() IVNCoreMLModel
	SetModel(value IVNCoreMLModel)
	VNCoreMLRequestRevision1() int
	Confidence() Confidence /* not a class type */
	SetConfidence(value Confidence /* not a class type */)
	// methods:
}

// An image-analysis request that uses a Core ML model to process images.
//
// The results array of a Core ML-based image analysis request contains a different observation type, depending on the kind of object you use: If the model predicts a single feature, the model’s object has a non- value for and Vision treats the model as a classifier. The results are objects. If the model’s outputs include at least one output with a feature type of , Vision treats that model as an image-to-image model. The results are objects. Otherwise, Vision treats the model as a general predictor model. The results are objects.


// An image-analysis request that uses a Core ML model to process images.
//
// [Full Topic]
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



// An optional setting that tells the Vision algorithm how to scale an input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/imageCropAndScaleOption
func (c_ CoreMLRequest) ImageCropAndScaleOption() ImageCropAndScaleOption /* not a class type */ {
	rv := objc.Send[ImageCropAndScaleOption](c_.ID, objc.Sel("imageCropAndScaleOption"))
	return rv
}


// An optional setting that tells the Vision algorithm how to scale an input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/imageCropAndScaleOption
func (c_ CoreMLRequest) SetImageCropAndScaleOption(value ImageCropAndScaleOption /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImageCropAndScaleOption:"), value)
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c_ CoreMLRequest) ModelDescription() objc.IObject /* cross-framework: ModelDescription */ {
	rv := objc.Send[coreml.ModelDescription](c_.ID, objc.Sel("modelDescription"))
	return rv
}


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c_ CoreMLRequest) SetModelDescription(value objc.IObject /* cross-framework: ModelDescription */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModelDescription:"), value)
}


// The name of the primary prediction feature output description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c_ CoreMLRequest) PredictedFeatureName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("predictedFeatureName"))
	return rv
}


// The name of the primary prediction feature output description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c_ CoreMLRequest) SetPredictedFeatureName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredictedFeatureName:"), value)
}


// The model to base the image analysis request on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlrequest/model
func (c_ CoreMLRequest) Model() IVNCoreMLModel {
	rv := objc.Send[CoreMLModel](c_.ID, objc.Sel("model"))
	return rv
}


// The model to base the image analysis request on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlrequest/model
func (c_ CoreMLRequest) SetModel(value IVNCoreMLModel) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModel:"), value)
}


// A constant for specifying revision 1 of a Core ML request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlrequestrevision1
func (c_ CoreMLRequest) VNCoreMLRequestRevision1() int {
	rv := objc.Send[int](c_.ID, objc.Sel("VNCoreMLRequestRevision1"))
	return rv
}


// The level of confidence in the observation’s accuracy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnobservation/confidence
func (c_ CoreMLRequest) Confidence() Confidence /* not a class type */ {
	rv := objc.Send[Confidence](c_.ID, objc.Sel("confidence"))
	return rv
}


// The level of confidence in the observation’s accuracy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnobservation/confidence
func (c_ CoreMLRequest) SetConfidence(value Confidence /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfidence:"), value)
}



