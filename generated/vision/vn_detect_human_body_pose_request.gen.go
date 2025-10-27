// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	Results() []HumanBodyPoseObservation
	SupportedJointNames() objectivec.IObject
	SetSupportedJointNames(value objectivec.IObject)
	SupportedJointsGroupNames() objectivec.IObject
	SetSupportedJointsGroupNames(value objectivec.IObject)
	VNDetectHumanBodyPoseRequestRevision1() int


	

	// methods:
	SupportedJointNamesAndReturnError(error_ foundation.foundation.INSError) []string
	SupportedJointsGroupNamesAndReturnError(error_ foundation.foundation.INSError) []string


}





// Alloc allocates a new instance without initialization.
func (dc _DetectHumanBodyPoseRequestClass) Alloc() DetectHumanBodyPoseRequest {
	rv := objc.Send[DetectHumanBodyPoseRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A request that detects a human body pose.
//
// The framework provides the detected body pose as a .


// A request that detects a human body pose.
//
// [Full Topic]
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










// Retrieves the supported joint names for a revision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPoseRequest/supportedJointNames(forRevision:)
func (dc _DetectHumanBodyPoseRequestClass) SupportedJointNamesForRevisionError(revision uint, error_ foundation.foundation.INSError) []string {
	rv := objc.Send[[]string](objc.ID(dc.class), objc.Sel("supportedJointNamesForRevision:error:"), revision, error_)
	return rv
}


// Retrieves the supported joint group names for a revision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPoseRequest/supportedJointsGroupNames(forRevision:)
func (dc _DetectHumanBodyPoseRequestClass) SupportedJointsGroupNamesForRevisionError(revision uint, error_ foundation.foundation.INSError) []string {
	rv := objc.Send[[]string](objc.ID(dc.class), objc.Sel("supportedJointsGroupNamesForRevision:error:"), revision, error_)
	return rv
}












// Retrieves the supported joint names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPoseRequest/supportedJointNamesAndReturnError:
func (d_ DetectHumanBodyPoseRequest) SupportedJointNamesAndReturnError(error_ foundation.foundation.INSError) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("supportedJointNamesAndReturnError:"), error_)
	return rv
}


// Retrieves the supported joint group names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPoseRequest/supportedJointsGroupNamesAndReturnError:
func (d_ DetectHumanBodyPoseRequest) SupportedJointsGroupNamesAndReturnError(error_ foundation.foundation.INSError) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("supportedJointsGroupNamesAndReturnError:"), error_)
	return rv
}







// The observed body poses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPoseRequest/results
func (d_ DetectHumanBodyPoseRequest) Results() []HumanBodyPoseObservation {
	rv := objc.Send[[]HumanBodyPoseObservation](d_.ID, objc.Sel("results"))
	return rv
}


// Retrieves the supported joint names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodyposerequest/supportedjointnames
func (d_ DetectHumanBodyPoseRequest) SupportedJointNames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("supportedJointNames"))
	return rv
}


// Retrieves the supported joint names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodyposerequest/supportedjointnames
func (d_ DetectHumanBodyPoseRequest) SetSupportedJointNames(value objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointNames:"), value)
}


// Retrieves the supported joint group names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodyposerequest/supportedjointsgroupnames
func (d_ DetectHumanBodyPoseRequest) SupportedJointsGroupNames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("supportedJointsGroupNames"))
	return rv
}


// Retrieves the supported joint group names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodyposerequest/supportedjointsgroupnames
func (d_ DetectHumanBodyPoseRequest) SetSupportedJointsGroupNames(value objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointsGroupNames:"), value)
}


// A constant for specifying revision 1 of the body pose detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodyposerequestrevision1
func (d_ DetectHumanBodyPoseRequest) VNDetectHumanBodyPoseRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectHumanBodyPoseRequestRevision1"))
	return rv
}








