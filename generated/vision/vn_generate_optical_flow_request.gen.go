// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [GenerateOpticalFlowRequest] class.
type IGenerateOpticalFlowRequest interface {
	ITargetedImageRequest
}

// An object that generates directional change vectors for each pixel in the targeted image.
//
// This request operates at a pixel level, so both images need to have the same dimensions to successfully perform the analysis. Setting a region of interest limits the region in which the analysis occurs. However, the system reports the resulting observation at full resolution. Optical flow requests are resource-intensive, so create only one request at a time, and release it immediately after generating optical flows.
//
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

// Alloc allocates a new instance without initialization.
func (gc _GenerateOpticalFlowRequestClass) Alloc() GenerateOpticalFlowRequest {
	rv := objc.Send[GenerateOpticalFlowRequest](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A constant for specifying revision 2 of the optical flow generation request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequestrevision2
func (g_ GenerateOpticalFlowRequest) VNGenerateOpticalFlowRequestRevision2() int {
	rv := objc.Send[int](g_.ID, objc.Sel("VNGenerateOpticalFlowRequestRevision2"))
	return rv
}

// The results of the request to generate optical flow.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequest/results
func (g_ GenerateOpticalFlowRequest) Results() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("results"))
	return rv
}


// SetResults sets the value of the results property.
// The results of the request to generate optical flow.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequest/results
func (g_ GenerateOpticalFlowRequest) SetResults(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setResults:"), value)
}

// A constant for specifying revision 1 of the optical flow generation request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequestrevision1
func (g_ GenerateOpticalFlowRequest) VNGenerateOpticalFlowRequestRevision1() int {
	rv := objc.Send[int](g_.ID, objc.Sel("VNGenerateOpticalFlowRequestRevision1"))
	return rv
}

// A Boolean value that indicates whether to keep the raw pixel buffer coming from the machine learning network.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequest/keepnetworkoutput
func (g_ GenerateOpticalFlowRequest) KeepNetworkOutput() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("keepNetworkOutput"))
	return rv
}


// SetKeepNetworkOutput sets the value of the keepNetworkOutput property.
// A Boolean value that indicates whether to keep the raw pixel buffer coming from the machine learning network.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequest/keepnetworkoutput
func (g_ GenerateOpticalFlowRequest) SetKeepNetworkOutput(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setKeepNetworkOutput:"), value)
}

// The output buffer’s pixel format.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequest/outputpixelformat
func (g_ GenerateOpticalFlowRequest) OutputPixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("outputPixelFormat"))
	return rv
}


// SetOutputPixelFormat sets the value of the outputPixelFormat property.
// The output buffer’s pixel format.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequest/outputpixelformat
func (g_ GenerateOpticalFlowRequest) SetOutputPixelFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setOutputPixelFormat:"), value)
}

// The accuracy level for computing optical flow.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequest/computationaccuracy-swift.property
func (g_ GenerateOpticalFlowRequest) ComputationAccuracy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("computationAccuracy"))
	return rv
}


// SetComputationAccuracy sets the value of the computationAccuracy property.
// The accuracy level for computing optical flow.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateopticalflowrequest/computationaccuracy-swift.property
func (g_ GenerateOpticalFlowRequest) SetComputationAccuracy(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setComputationAccuracy:"), value)
}



