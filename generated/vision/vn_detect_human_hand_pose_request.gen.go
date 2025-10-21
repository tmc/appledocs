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
	SupportedJointNamesAndReturnError(error_ unsafe.Pointer) []string
	SupportedJointsGroupNamesAndReturnError(error_ unsafe.Pointer) []string
}

// A request that detects a human hand pose.
//
// The framework provides the detected hand pose as a .
//
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


// Retrieves the supported joint names for a revision.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/supportedJointNames(forRevision:)
func (dc _DetectHumanHandPoseRequestClass) SupportedJointNamesForRevisionError(revision uint, error_ unsafe.Pointer) []string {
	rv := objc.Send[[]string](objc.ID(dc.class), objc.Sel("supportedJointNamesForRevision:error:"), revision, error_)
	return rv
}

// Retrieves the supported joint group names for a revision.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/supportedJointsGroupNames(forRevision:)
func (dc _DetectHumanHandPoseRequestClass) SupportedJointsGroupNamesForRevisionError(revision uint, error_ unsafe.Pointer) []string {
	rv := objc.Send[[]string](objc.ID(dc.class), objc.Sel("supportedJointsGroupNamesForRevision:error:"), revision, error_)
	return rv
}

// Retrieves the supported joint names.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/supportedJointNamesAndReturnError:
func (d_ DetectHumanHandPoseRequest) SupportedJointNamesAndReturnError(error_ unsafe.Pointer) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("supportedJointNamesAndReturnError:"), error_)
	return rv
}

// Retrieves the supported joint group names.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/supportedJointsGroupNamesAndReturnError:
func (d_ DetectHumanHandPoseRequest) SupportedJointsGroupNamesAndReturnError(error_ unsafe.Pointer) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("supportedJointsGroupNamesAndReturnError:"), error_)
	return rv
}

// A constant for specifying revision 1 of the hand pose detection request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequestrevision1
func (d_ DetectHumanHandPoseRequest) VNDetectHumanHandPoseRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectHumanHandPoseRequestRevision1"))
	return rv
}

// Retrieves the supported joint group names.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/supportedjointsgroupnames
func (d_ DetectHumanHandPoseRequest) SupportedJointsGroupNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("supportedJointsGroupNames"))
	return rv
}


// SetSupportedJointsGroupNames sets the value of the supportedJointsGroupNames property.
// Retrieves the supported joint group names.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/supportedjointsgroupnames
func (d_ DetectHumanHandPoseRequest) SetSupportedJointsGroupNames(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointsGroupNames:"), value)
}

// Retrieves the supported joint names.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/supportedjointnames
func (d_ DetectHumanHandPoseRequest) SupportedJointNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("supportedJointNames"))
	return rv
}


// SetSupportedJointNames sets the value of the supportedJointNames property.
// Retrieves the supported joint names.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/supportedjointnames
func (d_ DetectHumanHandPoseRequest) SetSupportedJointNames(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointNames:"), value)
}

// The maximum number of hands to detect in an image.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/maximumHandCount
func (d_ DetectHumanHandPoseRequest) MaximumHandCount() uint {
	rv := objc.Send[uint](d_.ID, objc.Sel("maximumHandCount"))
	return rv
}


// SetMaximumHandCount sets the value of the maximumHandCount property.
// The maximum number of hands to detect in an image.

//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/maximumHandCount
func (d_ DetectHumanHandPoseRequest) SetMaximumHandCount(value uint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaximumHandCount:"), value)
}

// The observed hand poses.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/results
func (d_ DetectHumanHandPoseRequest) Results() []HumanHandPoseObservation {
	rv := objc.Send[[]HumanHandPoseObservation](d_.ID, objc.Sel("results"))
	return rv
}



