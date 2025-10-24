// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNGenerateImageFeaturePrintRequest */


/* debug [class_header]: Header for VNGenerateImageFeaturePrintRequest */
// The class instance for the [GenerateImageFeaturePrintRequest] class.
var (
	GenerateImageFeaturePrintRequestClass     _GenerateImageFeaturePrintRequestClass
	GenerateImageFeaturePrintRequestClassOnce sync.Once
)

func getGenerateImageFeaturePrintRequestClass() _GenerateImageFeaturePrintRequestClass {
	GenerateImageFeaturePrintRequestClassOnce.Do(func() {
		GenerateImageFeaturePrintRequestClass = _GenerateImageFeaturePrintRequestClass{objc.GetClass("VNGenerateImageFeaturePrintRequest")}
	})
	return GenerateImageFeaturePrintRequestClass
}

type _GenerateImageFeaturePrintRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GenerateImageFeaturePrintRequest */
// An interface definition for the [GenerateImageFeaturePrintRequest] class.
type IGenerateImageFeaturePrintRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for GenerateImageFeaturePrintRequest */
	// properties:
	ImageCropAndScaleOption() ImageCropAndScaleOption
	SetImageCropAndScaleOption(value ImageCropAndScaleOption)
	Results() []FeaturePrintObservation
	VNGenerateImageFeaturePrintRequestRevision1() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GenerateImageFeaturePrintRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GenerateImageFeaturePrintRequest */
// Alloc allocates a new instance without initialization.
func (gc _GenerateImageFeaturePrintRequestClass) Alloc() GenerateImageFeaturePrintRequest {
	rv := objc.Send[GenerateImageFeaturePrintRequest](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GenerateImageFeaturePrintRequestClass) New() GenerateImageFeaturePrintRequest {
	rv := objc.Send[GenerateImageFeaturePrintRequest](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GenerateImageFeaturePrintRequest) Init() GenerateImageFeaturePrintRequest {
	rv := objc.Send[GenerateImageFeaturePrintRequest](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GenerateImageFeaturePrintRequest) Autorelease() GenerateImageFeaturePrintRequest {
	rv := objc.Send[GenerateImageFeaturePrintRequest](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGenerateImageFeaturePrintRequest creates a new GenerateImageFeaturePrintRequest instance.
func NewGenerateImageFeaturePrintRequest() GenerateImageFeaturePrintRequest {
	return getGenerateImageFeaturePrintRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GenerateImageFeaturePrintRequest */
// An image-based request to generate feature prints from an image.
//
// This request returns the feature print data it generates as an array of objects.


// An image-based request to generate feature prints from an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateImageFeaturePrintRequest
type GenerateImageFeaturePrintRequest struct {
	ImageBasedRequest
}

// GenerateImageFeaturePrintRequestFrom constructs a [GenerateImageFeaturePrintRequest] from an unsafe.Pointer.
//
// An image-based request to generate feature prints from an image.
func GenerateImageFeaturePrintRequestFrom(ptr unsafe.Pointer) GenerateImageFeaturePrintRequest {
	return GenerateImageFeaturePrintRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GenerateImageFeaturePrintRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GenerateImageFeaturePrintRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GenerateImageFeaturePrintRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GenerateImageFeaturePrintRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GenerateImageFeaturePrintRequest */

// An optional setting that tells the algorithm how to scale an input image before generating the feature print.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateImageFeaturePrintRequest/imageCropAndScaleOption
func (g_ GenerateImageFeaturePrintRequest) ImageCropAndScaleOption() ImageCropAndScaleOption {
	rv := objc.Send[ImageCropAndScaleOption](g_.ID, objc.Sel("imageCropAndScaleOption"))
	return rv
}/* debug [instance_properties/getter]: imageCropAndScaleOption */


// An optional setting that tells the algorithm how to scale an input image before generating the feature print.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateImageFeaturePrintRequest/imageCropAndScaleOption
func (g_ GenerateImageFeaturePrintRequest) SetImageCropAndScaleOption(value ImageCropAndScaleOption) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setImageCropAndScaleOption:"), value)
}/* debug [instance_properties/setter]: imageCropAndScaleOption */


// The results of the feature print request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateImageFeaturePrintRequest/results
func (g_ GenerateImageFeaturePrintRequest) Results() []FeaturePrintObservation {
	rv := objc.Send[[]FeaturePrintObservation](g_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// A constant for specifying the first revision of the feature-print request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateimagefeatureprintrequestrevision1
func (g_ GenerateImageFeaturePrintRequest) VNGenerateImageFeaturePrintRequestRevision1() int {
	rv := objc.Send[int](g_.ID, objc.Sel("VNGenerateImageFeaturePrintRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNGenerateImageFeaturePrintRequestRevision1 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNGenerateImageFeaturePrintRequest */



