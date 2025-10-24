// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNGeneratePersonSegmentationRequest */


/* debug [class_header]: Header for VNGeneratePersonSegmentationRequest */
// The class instance for the [GeneratePersonSegmentationRequest] class.
var (
	GeneratePersonSegmentationRequestClass     _GeneratePersonSegmentationRequestClass
	GeneratePersonSegmentationRequestClassOnce sync.Once
)

func getGeneratePersonSegmentationRequestClass() _GeneratePersonSegmentationRequestClass {
	GeneratePersonSegmentationRequestClassOnce.Do(func() {
		GeneratePersonSegmentationRequestClass = _GeneratePersonSegmentationRequestClass{objc.GetClass("VNGeneratePersonSegmentationRequest")}
	})
	return GeneratePersonSegmentationRequestClass
}

type _GeneratePersonSegmentationRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GeneratePersonSegmentationRequest */
// An interface definition for the [GeneratePersonSegmentationRequest] class.
type IGeneratePersonSegmentationRequest interface {
	IStatefulRequest
	
/* debug [class_interface_properties]: Properties for GeneratePersonSegmentationRequest */
	// properties:
	OutputPixelFormat() uint32 /* not a class type */
	SetOutputPixelFormat(value uint32 /* not a class type */)
	QualityLevel() GeneratePersonSegmentationRequestQualityLevel
	SetQualityLevel(value GeneratePersonSegmentationRequestQualityLevel)
	Results() []PixelBufferObservation
	VNGeneratePersonSegmentationRequestRevision1() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GeneratePersonSegmentationRequest */
	// methods:
	SupportedOutputPixelFormatsAndReturnError(error_ objectivec.IObject) []foundation.Number
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GeneratePersonSegmentationRequest */
// Alloc allocates a new instance without initialization.
func (gc _GeneratePersonSegmentationRequestClass) Alloc() GeneratePersonSegmentationRequest {
	rv := objc.Send[GeneratePersonSegmentationRequest](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GeneratePersonSegmentationRequestClass) New() GeneratePersonSegmentationRequest {
	rv := objc.Send[GeneratePersonSegmentationRequest](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GeneratePersonSegmentationRequest) Init() GeneratePersonSegmentationRequest {
	rv := objc.Send[GeneratePersonSegmentationRequest](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GeneratePersonSegmentationRequest) Autorelease() GeneratePersonSegmentationRequest {
	rv := objc.Send[GeneratePersonSegmentationRequest](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGeneratePersonSegmentationRequest creates a new GeneratePersonSegmentationRequest instance.
func NewGeneratePersonSegmentationRequest() GeneratePersonSegmentationRequest {
	return getGeneratePersonSegmentationRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GeneratePersonSegmentationRequest */
// An object that produces a matte image for a person it finds in the input image.
//
// Perform this request to detect and generate an image mask for a person in an image. The request returns the resulting image mask in an instance of .


// An object that produces a matte image for a person it finds in the input image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest
type GeneratePersonSegmentationRequest struct {
	StatefulRequest
}

// GeneratePersonSegmentationRequestFrom constructs a [GeneratePersonSegmentationRequest] from an unsafe.Pointer.
//
// An object that produces a matte image for a person it finds in the input image.
func GeneratePersonSegmentationRequestFrom(ptr unsafe.Pointer) GeneratePersonSegmentationRequest {
	return GeneratePersonSegmentationRequest{
		StatefulRequest: StatefulRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GeneratePersonSegmentationRequest */

// Creates a generate person segmentation request with a completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/init(completionHandler:)
func NewGeneratePersonSegmentationRequestWithCompletionHandler(completionHandler RequestCompletionHandler /* not a class type */) GeneratePersonSegmentationRequest {
	instance := getGeneratePersonSegmentationRequestClass().Alloc()
	rv := objc.Send[GeneratePersonSegmentationRequest](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGeneratePersonSegmentationRequestWithCompletionHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GeneratePersonSegmentationRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GeneratePersonSegmentationRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GeneratePersonSegmentationRequest */

// Returns a list of output pixel formats that the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/supportedOutputPixelFormats()
func (g_ GeneratePersonSegmentationRequest) SupportedOutputPixelFormatsAndReturnError(error_ objectivec.IObject) []foundation.Number {
	rv := objc.Send[[]foundation.Number](g_.ID, objc.Sel("supportedOutputPixelFormatsAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: SupportedOutputPixelFormatsAndReturnError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GeneratePersonSegmentationRequest */

// The pixel format of the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/outputPixelFormat
func (g_ GeneratePersonSegmentationRequest) OutputPixelFormat() uint32 /* not a class type */ {
	rv := objc.Send[uint32](g_.ID, objc.Sel("outputPixelFormat"))
	return rv
}/* debug [instance_properties/getter]: outputPixelFormat */


// The pixel format of the output image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/outputPixelFormat
func (g_ GeneratePersonSegmentationRequest) SetOutputPixelFormat(value uint32 /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputPixelFormat:"), value)
}/* debug [instance_properties/setter]: outputPixelFormat */


// A value that indicates how the request balances accuracy and performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/qualityLevel-swift.property
func (g_ GeneratePersonSegmentationRequest) QualityLevel() GeneratePersonSegmentationRequestQualityLevel {
	rv := objc.Send[GeneratePersonSegmentationRequestQualityLevel](g_.ID, objc.Sel("qualityLevel"))
	return rv
}/* debug [instance_properties/getter]: qualityLevel */


// A value that indicates how the request balances accuracy and performance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/qualityLevel-swift.property
func (g_ GeneratePersonSegmentationRequest) SetQualityLevel(value GeneratePersonSegmentationRequestQualityLevel) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setQualityLevel:"), value)
}/* debug [instance_properties/setter]: qualityLevel */


// The results of the segmentation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeneratePersonSegmentationRequest/results
func (g_ GeneratePersonSegmentationRequest) Results() []PixelBufferObservation {
	rv := objc.Send[[]PixelBufferObservation](g_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// A constant for specifying revision 1 of the person segmentation generation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vngeneratepersonsegmentationrequestrevision1
func (g_ GeneratePersonSegmentationRequest) VNGeneratePersonSegmentationRequestRevision1() int {
	rv := objc.Send[int](g_.ID, objc.Sel("VNGeneratePersonSegmentationRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNGeneratePersonSegmentationRequestRevision1 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNGeneratePersonSegmentationRequest */


