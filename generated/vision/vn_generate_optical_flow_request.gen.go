// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VNGenerateOpticalFlowRequest */


/* debug [class_header]: Header for VNGenerateOpticalFlowRequest */
// The class instance for the [GenerateOpticalFlowRequest] class.
var (
	GenerateOpticalFlowRequestClass     _GenerateOpticalFlowRequestClass
	GenerateOpticalFlowRequestClassOnce sync.Once
)

func getGenerateOpticalFlowRequestClass() _GenerateOpticalFlowRequestClass {
	GenerateOpticalFlowRequestClassOnce.Do(func() {
		GenerateOpticalFlowRequestClass = _GenerateOpticalFlowRequestClass{objc.GetClass("VNGenerateOpticalFlowRequest")}
	})
	return GenerateOpticalFlowRequestClass
}

type _GenerateOpticalFlowRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GenerateOpticalFlowRequest */
// An interface definition for the [GenerateOpticalFlowRequest] class.
type IGenerateOpticalFlowRequest interface {
	ITargetedImageRequest
	
/* debug [class_interface_properties]: Properties for GenerateOpticalFlowRequest */
	// properties:
	ComputationAccuracy() GenerateOpticalFlowRequestComputationAccuracy
	SetComputationAccuracy(value GenerateOpticalFlowRequestComputationAccuracy)
	KeepNetworkOutput() bool
	SetKeepNetworkOutput(value bool)
	OutputPixelFormat() uint32 /* not a class type */
	SetOutputPixelFormat(value uint32 /* not a class type */)
	Results() []PixelBufferObservation
	VNGenerateOpticalFlowRequestRevision1() int
	VNGenerateOpticalFlowRequestRevision2() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GenerateOpticalFlowRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GenerateOpticalFlowRequest */
// Alloc allocates a new instance without initialization.
func (gc _GenerateOpticalFlowRequestClass) Alloc() GenerateOpticalFlowRequest {
	rv := objc.Send[GenerateOpticalFlowRequest](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GenerateOpticalFlowRequestClass) New() GenerateOpticalFlowRequest {
	rv := objc.Send[GenerateOpticalFlowRequest](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GenerateOpticalFlowRequest) Init() GenerateOpticalFlowRequest {
	rv := objc.Send[GenerateOpticalFlowRequest](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GenerateOpticalFlowRequest) Autorelease() GenerateOpticalFlowRequest {
	rv := objc.Send[GenerateOpticalFlowRequest](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGenerateOpticalFlowRequest creates a new GenerateOpticalFlowRequest instance.
func NewGenerateOpticalFlowRequest() GenerateOpticalFlowRequest {
	return getGenerateOpticalFlowRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GenerateOpticalFlowRequest */
// An object that generates directional change vectors for each pixel in the targeted image.
//
// This request operates at a pixel level, so both images need to have the same dimensions to successfully perform the analysis. Setting a region of interest limits the region in which the analysis occurs. However, the system reports the resulting observation at full resolution. Optical flow requests are resource-intensive, so create only one request at a time, and release it immediately after generating optical flows.


// An object that generates directional change vectors for each pixel in the targeted image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest
type GenerateOpticalFlowRequest struct {
	TargetedImageRequest
}

// GenerateOpticalFlowRequestFrom constructs a [GenerateOpticalFlowRequest] from an unsafe.Pointer.
//
// An object that generates directional change vectors for each pixel in the targeted image.
func GenerateOpticalFlowRequestFrom(ptr unsafe.Pointer) GenerateOpticalFlowRequest {
	return GenerateOpticalFlowRequest{
		TargetedImageRequest: TargetedImageRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GenerateOpticalFlowRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GenerateOpticalFlowRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GenerateOpticalFlowRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GenerateOpticalFlowRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GenerateOpticalFlowRequest */

// The accuracy level for computing optical flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest/computationAccuracy-swift.property
func (g_ GenerateOpticalFlowRequest) ComputationAccuracy() GenerateOpticalFlowRequestComputationAccuracy {
	rv := objc.Send[GenerateOpticalFlowRequestComputationAccuracy](g_.ID, objc.Sel("computationAccuracy"))
	return rv
}/* debug [instance_properties/getter]: computationAccuracy */


// The accuracy level for computing optical flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest/computationAccuracy-swift.property
func (g_ GenerateOpticalFlowRequest) SetComputationAccuracy(value GenerateOpticalFlowRequestComputationAccuracy) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setComputationAccuracy:"), value)
}/* debug [instance_properties/setter]: computationAccuracy */


// A Boolean value that indicates whether to keep the raw pixel buffer coming from the machine learning network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest/keepNetworkOutput
func (g_ GenerateOpticalFlowRequest) KeepNetworkOutput() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("keepNetworkOutput"))
	return rv
}/* debug [instance_properties/getter]: keepNetworkOutput */


// A Boolean value that indicates whether to keep the raw pixel buffer coming from the machine learning network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest/keepNetworkOutput
func (g_ GenerateOpticalFlowRequest) SetKeepNetworkOutput(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setKeepNetworkOutput:"), value)
}/* debug [instance_properties/setter]: keepNetworkOutput */


// The output buffer’s pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest/outputPixelFormat
func (g_ GenerateOpticalFlowRequest) OutputPixelFormat() uint32 /* not a class type */ {
	rv := objc.Send[uint32](g_.ID, objc.Sel("outputPixelFormat"))
	return rv
}/* debug [instance_properties/getter]: outputPixelFormat */


// The output buffer’s pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest/outputPixelFormat
func (g_ GenerateOpticalFlowRequest) SetOutputPixelFormat(value uint32 /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputPixelFormat:"), value)
}/* debug [instance_properties/setter]: outputPixelFormat */


// The results of the request to generate optical flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGenerateOpticalFlowRequest/results
func (g_ GenerateOpticalFlowRequest) Results() []PixelBufferObservation {
	rv := objc.Send[[]PixelBufferObservation](g_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// A constant for specifying revision 1 of the optical flow generation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequestrevision1
func (g_ GenerateOpticalFlowRequest) VNGenerateOpticalFlowRequestRevision1() int {
	rv := objc.Send[int](g_.ID, objc.Sel("VNGenerateOpticalFlowRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNGenerateOpticalFlowRequestRevision1 */


// A constant for specifying revision 2 of the optical flow generation request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequestrevision2
func (g_ GenerateOpticalFlowRequest) VNGenerateOpticalFlowRequestRevision2() int {
	rv := objc.Send[int](g_.ID, objc.Sel("VNGenerateOpticalFlowRequestRevision2"))
	return rv
}/* debug [instance_properties/getter]: VNGenerateOpticalFlowRequestRevision2 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNGenerateOpticalFlowRequest */



