// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DetectHumanHandPoseRequest] class.
var (
	DetectHumanHandPoseRequestClass     _DetectHumanHandPoseRequestClass
	DetectHumanHandPoseRequestClassOnce sync.Once
)

func getDetectHumanHandPoseRequestClass() _DetectHumanHandPoseRequestClass {
	DetectHumanHandPoseRequestClassOnce.Do(func() {
		DetectHumanHandPoseRequestClass = _DetectHumanHandPoseRequestClass{objc.GetClass("VNDetectHumanHandPoseRequest")}
	})
	return DetectHumanHandPoseRequestClass
}

type _DetectHumanHandPoseRequestClass struct {
	class objc.Class
}

// An interface definition for the [DetectHumanHandPoseRequest] class.
type IDetectHumanHandPoseRequest interface {
	IImageBasedRequest
	// properties:
	MaximumHandCount() int
	SetMaximumHandCount(value int)
	Results() IVNHumanHandPoseObservation
	SetResults(value IVNHumanHandPoseObservation)
	SupportedJointNames() unsafe.Pointer
	SetSupportedJointNames(value unsafe.Pointer)
	SupportedJointsGroupNames() unsafe.Pointer
	SetSupportedJointsGroupNames(value unsafe.Pointer)
	VNDetectHumanHandPoseRequestRevision1() int
	// methods:
}

// A request that detects a human hand pose.
//
// The framework provides the detected hand pose as a .


// A request that detects a human hand pose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest
type DetectHumanHandPoseRequest struct {
	ImageBasedRequest
}

// DetectHumanHandPoseRequestFrom constructs a [DetectHumanHandPoseRequest] from an unsafe.Pointer.
//
// A request that detects a human hand pose.
func DetectHumanHandPoseRequestFrom(ptr unsafe.Pointer) DetectHumanHandPoseRequest {
	return DetectHumanHandPoseRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DetectHumanHandPoseRequestClass) Alloc() DetectHumanHandPoseRequest {
	rv := objc.Send[DetectHumanHandPoseRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DetectHumanHandPoseRequestClass) New() DetectHumanHandPoseRequest {
	rv := objc.Send[DetectHumanHandPoseRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectHumanHandPoseRequest) Init() DetectHumanHandPoseRequest {
	rv := objc.Send[DetectHumanHandPoseRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectHumanHandPoseRequest) Autorelease() DetectHumanHandPoseRequest {
	rv := objc.Send[DetectHumanHandPoseRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectHumanHandPoseRequest creates a new DetectHumanHandPoseRequest instance.
func NewDetectHumanHandPoseRequest() DetectHumanHandPoseRequest {
	return getDetectHumanHandPoseRequestClass().New()
}



// The maximum number of hands to detect in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/maximumhandcount
func (d_ DetectHumanHandPoseRequest) MaximumHandCount() int {
	rv := objc.Send[int](d_.ID, objc.Sel("maximumHandCount"))
	return rv
}


// The maximum number of hands to detect in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/maximumhandcount
func (d_ DetectHumanHandPoseRequest) SetMaximumHandCount(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaximumHandCount:"), value)
}


// The observed hand poses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/results
func (d_ DetectHumanHandPoseRequest) Results() IVNHumanHandPoseObservation {
	rv := objc.Send[HumanHandPoseObservation](d_.ID, objc.Sel("results"))
	return rv
}


// The observed hand poses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/results
func (d_ DetectHumanHandPoseRequest) SetResults(value IVNHumanHandPoseObservation) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setResults:"), value)
}


// Retrieves the supported joint names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/supportedjointnames
func (d_ DetectHumanHandPoseRequest) SupportedJointNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("supportedJointNames"))
	return rv
}


// Retrieves the supported joint names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/supportedjointnames
func (d_ DetectHumanHandPoseRequest) SetSupportedJointNames(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointNames:"), value)
}


// Retrieves the supported joint group names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/supportedjointsgroupnames
func (d_ DetectHumanHandPoseRequest) SupportedJointsGroupNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("supportedJointsGroupNames"))
	return rv
}


// Retrieves the supported joint group names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/supportedjointsgroupnames
func (d_ DetectHumanHandPoseRequest) SetSupportedJointsGroupNames(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointsGroupNames:"), value)
}


// A constant for specifying revision 1 of the hand pose detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequestrevision1
func (d_ DetectHumanHandPoseRequest) VNDetectHumanHandPoseRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectHumanHandPoseRequestRevision1"))
	return rv
}



