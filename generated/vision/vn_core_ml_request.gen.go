// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNCoreMLRequest */


/* debug [class_header]: Header for VNCoreMLRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CoreMLRequest */
// An interface definition for the [CoreMLRequest] class.
type ICoreMLRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for CoreMLRequest */
	// properties:
	ImageCropAndScaleOption() ImageCropAndScaleOption
	SetImageCropAndScaleOption(value ImageCropAndScaleOption)
	Model() IVNCoreMLModel
	VNCoreMLRequestRevision1() int
	Confidence() Confidence /* typedef */
	SetConfidence(value Confidence /* typedef */)
	ModelDescription() coreml.ModelDescription
	SetModelDescription(value coreml.ModelDescription)
	PredictedFeatureName() objc.IObject /* cross-framework: NSString */
	SetPredictedFeatureName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CoreMLRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CoreMLRequest */
// Alloc allocates a new instance without initialization.
func (cc _CoreMLRequestClass) Alloc() CoreMLRequest {
	rv := objc.Send[CoreMLRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CoreMLRequest */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CoreMLRequest */

// Creates a model container to use with an image analysis request based on the model you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/init(model:)
func NewCoreMLRequestWithModel(model IVNCoreMLModel) CoreMLRequest {
	instance := getCoreMLRequestClass().Alloc()
	rv := objc.Send[CoreMLRequest](instance.ID, objc.Sel("initWithModel:"), model)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCoreMLRequestWithModel */


// Creates a model container to use with an image analysis request based on the model you provide, with an optional completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/init(model:completionHandler:)
func NewCoreMLRequestWithModelCompletionHandler(model IVNCoreMLModel, completionHandler RequestCompletionHandler /* not a class type */) CoreMLRequest {
	instance := getCoreMLRequestClass().Alloc()
	rv := objc.Send[CoreMLRequest](instance.ID, objc.Sel("initWithModel:completionHandler:"), model, completionHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCoreMLRequestWithModelCompletionHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CoreMLRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CoreMLRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CoreMLRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CoreMLRequest */

// An optional setting that tells the Vision algorithm how to scale an input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/imageCropAndScaleOption
func (c_ CoreMLRequest) ImageCropAndScaleOption() ImageCropAndScaleOption {
	rv := objc.Send[ImageCropAndScaleOption](c_.ID, objc.Sel("imageCropAndScaleOption"))
	return rv
}/* debug [instance_properties/getter]: imageCropAndScaleOption */


// An optional setting that tells the Vision algorithm how to scale an input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/imageCropAndScaleOption
func (c_ CoreMLRequest) SetImageCropAndScaleOption(value ImageCropAndScaleOption) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImageCropAndScaleOption:"), value)
}/* debug [instance_properties/setter]: imageCropAndScaleOption */


// The model to base the image analysis request on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNCoreMLRequest/model
func (c_ CoreMLRequest) Model() IVNCoreMLModel {
	rv := objc.Send[CoreMLModel](c_.ID, objc.Sel("model"))
	return rv
}/* debug [instance_properties/getter]: model */


// A constant for specifying revision 1 of a Core ML request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vncoremlrequestrevision1
func (c_ CoreMLRequest) VNCoreMLRequestRevision1() int {
	rv := objc.Send[int](c_.ID, objc.Sel("VNCoreMLRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNCoreMLRequestRevision1 */


// The level of confidence in the observation’s accuracy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnobservation/confidence
func (c_ CoreMLRequest) Confidence() Confidence /* typedef */ {
	rv := objc.Send[float32](c_.ID, objc.Sel("confidence"))
	return rv
}/* debug [instance_properties/getter]: confidence */


// The level of confidence in the observation’s accuracy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vnobservation/confidence
func (c_ CoreMLRequest) SetConfidence(value Confidence /* typedef */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfidence:"), value)
}/* debug [instance_properties/setter]: confidence */


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c_ CoreMLRequest) ModelDescription() coreml.ModelDescription {
	rv := objc.Send[coreml.ModelDescription](c_.ID, objc.Sel("modelDescription"))
	return rv
}/* debug [instance_properties/getter]: modelDescription */


// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c_ CoreMLRequest) SetModelDescription(value coreml.ModelDescription) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModelDescription:"), value)
}/* debug [instance_properties/setter]: modelDescription */


// The name of the primary prediction feature output description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c_ CoreMLRequest) PredictedFeatureName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("predictedFeatureName"))
	return rv
}/* debug [instance_properties/getter]: predictedFeatureName */


// The name of the primary prediction feature output description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c_ CoreMLRequest) SetPredictedFeatureName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPredictedFeatureName:"), value)
}/* debug [instance_properties/setter]: predictedFeatureName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNCoreMLRequest */


