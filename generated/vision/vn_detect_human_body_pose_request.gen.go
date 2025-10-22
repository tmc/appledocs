// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DetectHumanBodyPoseRequest] class.
var (
	DetectHumanBodyPoseRequestClass     _DetectHumanBodyPoseRequestClass
	DetectHumanBodyPoseRequestClassOnce sync.Once
)

func getDetectHumanBodyPoseRequestClass() _DetectHumanBodyPoseRequestClass {
	DetectHumanBodyPoseRequestClassOnce.Do(func() {
		DetectHumanBodyPoseRequestClass = _DetectHumanBodyPoseRequestClass{objc.GetClass("VNDetectHumanBodyPoseRequest")}
	})
	return DetectHumanBodyPoseRequestClass
}

type _DetectHumanBodyPoseRequestClass struct {
	class objc.Class
}

// An interface definition for the [DetectHumanBodyPoseRequest] class.
type IDetectHumanBodyPoseRequest interface {
	IImageBasedRequest
	Results() []HumanBodyPoseObservation
	SupportedJointNames() unsafe.Pointer
	SetSupportedJointNames(value unsafe.Pointer)
	SupportedJointsGroupNames() unsafe.Pointer
	SetSupportedJointsGroupNames(value unsafe.Pointer)
	VNDetectHumanBodyPoseRequestRevision1() int
}

// A request that detects a human body pose.
//
// The framework provides the detected body pose as a .
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPoseRequest
type DetectHumanBodyPoseRequest struct {
	ImageBasedRequest
}

// DetectHumanBodyPoseRequestFrom constructs a [DetectHumanBodyPoseRequest] from an unsafe.Pointer.
//
// A request that detects a human body pose.
func DetectHumanBodyPoseRequestFrom(ptr unsafe.Pointer) DetectHumanBodyPoseRequest {
	return DetectHumanBodyPoseRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DetectHumanBodyPoseRequestClass) Alloc() DetectHumanBodyPoseRequest {
	rv := objc.Send[DetectHumanBodyPoseRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DetectHumanBodyPoseRequestClass) New() DetectHumanBodyPoseRequest {
	rv := objc.Send[DetectHumanBodyPoseRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectHumanBodyPoseRequest) Init() DetectHumanBodyPoseRequest {
	rv := objc.Send[DetectHumanBodyPoseRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectHumanBodyPoseRequest) Autorelease() DetectHumanBodyPoseRequest {
	rv := objc.Send[DetectHumanBodyPoseRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectHumanBodyPoseRequest creates a new DetectHumanBodyPoseRequest instance.
func NewDetectHumanBodyPoseRequest() DetectHumanBodyPoseRequest {
	return getDetectHumanBodyPoseRequestClass().New()
}


// The observed body poses.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPoseRequest/results
func (d_ DetectHumanBodyPoseRequest) Results() []HumanBodyPoseObservation {
	rv := objc.Send[[]HumanBodyPoseObservation](d_.ID, objc.Sel("results"))
	return rv
}

// Retrieves the supported joint names.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodyposerequest/supportedjointnames
func (d_ DetectHumanBodyPoseRequest) SupportedJointNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("supportedJointNames"))
	return rv
}


// SetSupportedJointNames sets the value of the supportedJointNames property.
// Retrieves the supported joint names.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodyposerequest/supportedjointnames
func (d_ DetectHumanBodyPoseRequest) SetSupportedJointNames(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointNames:"), value)
}

// Retrieves the supported joint group names.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodyposerequest/supportedjointsgroupnames
func (d_ DetectHumanBodyPoseRequest) SupportedJointsGroupNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("supportedJointsGroupNames"))
	return rv
}


// SetSupportedJointsGroupNames sets the value of the supportedJointsGroupNames property.
// Retrieves the supported joint group names.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodyposerequest/supportedjointsgroupnames
func (d_ DetectHumanBodyPoseRequest) SetSupportedJointsGroupNames(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointsGroupNames:"), value)
}

// A constant for specifying revision 1 of the body pose detection request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodyposerequestrevision1
func (d_ DetectHumanBodyPoseRequest) VNDetectHumanBodyPoseRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectHumanBodyPoseRequestRevision1"))
	return rv
}



