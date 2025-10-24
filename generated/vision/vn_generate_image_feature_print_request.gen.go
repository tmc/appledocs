// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [GenerateImageFeaturePrintRequest] class.
type IGenerateImageFeaturePrintRequest interface {
	IImageBasedRequest
	// properties:
	ImageCropAndScaleOption() ImageCropAndScaleOption /* not a class type */
	SetImageCropAndScaleOption(value ImageCropAndScaleOption /* not a class type */)
	Results() IVNFeaturePrintObservation
	SetResults(value IVNFeaturePrintObservation)
	VNGenerateImageFeaturePrintRequestRevision1() int
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (gc _GenerateImageFeaturePrintRequestClass) Alloc() GenerateImageFeaturePrintRequest {
	rv := objc.Send[GenerateImageFeaturePrintRequest](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// An optional setting that tells the algorithm how to scale an input image before generating the feature print.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateimagefeatureprintrequest/imagecropandscaleoption
func (g_ GenerateImageFeaturePrintRequest) ImageCropAndScaleOption() ImageCropAndScaleOption /* not a class type */ {
	rv := objc.Send[ImageCropAndScaleOption](g_.ID, objc.Sel("imageCropAndScaleOption"))
	return rv
}


// An optional setting that tells the algorithm how to scale an input image before generating the feature print.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateimagefeatureprintrequest/imagecropandscaleoption
func (g_ GenerateImageFeaturePrintRequest) SetImageCropAndScaleOption(value ImageCropAndScaleOption /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setImageCropAndScaleOption:"), value)
}


// The results of the feature print request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateimagefeatureprintrequest/results
func (g_ GenerateImageFeaturePrintRequest) Results() IVNFeaturePrintObservation {
	rv := objc.Send[FeaturePrintObservation](g_.ID, objc.Sel("results"))
	return rv
}


// The results of the feature print request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateimagefeatureprintrequest/results
func (g_ GenerateImageFeaturePrintRequest) SetResults(value IVNFeaturePrintObservation) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResults:"), value)
}


// A constant for specifying the first revision of the feature-print request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateimagefeatureprintrequestrevision1
func (g_ GenerateImageFeaturePrintRequest) VNGenerateImageFeaturePrintRequestRevision1() int {
	rv := objc.Send[int](g_.ID, objc.Sel("VNGenerateImageFeaturePrintRequestRevision1"))
	return rv
}



