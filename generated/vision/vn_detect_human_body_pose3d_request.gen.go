// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [DetectHumanBodyPose3DRequest] class.
var (
	DetectHumanBodyPose3DRequestClass     _DetectHumanBodyPose3DRequestClass
	DetectHumanBodyPose3DRequestClassOnce sync.Once
)

func getDetectHumanBodyPose3DRequestClass() _DetectHumanBodyPose3DRequestClass {
	DetectHumanBodyPose3DRequestClassOnce.Do(func() {
		DetectHumanBodyPose3DRequestClass = _DetectHumanBodyPose3DRequestClass{objc.GetClass("VNDetectHumanBodyPose3DRequest")}
	})
	return DetectHumanBodyPose3DRequestClass
}

type _DetectHumanBodyPose3DRequestClass struct {
	class objc.Class
}





// An interface definition for the [DetectHumanBodyPose3DRequest] class.
type IDetectHumanBodyPose3DRequest interface {
	IStatefulRequest
	

	// properties:
	Results() []HumanBodyPose3DObservation
	SupportedJointNames() objectivec.IObject
	SetSupportedJointNames(value objectivec.IObject)
	SupportedJointsGroupNames() objectivec.IObject
	SetSupportedJointsGroupNames(value objectivec.IObject)


	

	// methods:
	SupportedJointNamesAndReturnError(error_ foundation.foundation.INSError) []string
	SupportedJointsGroupNamesAndReturnError(error_ foundation.foundation.INSError) []string


}





// Alloc allocates a new instance without initialization.
func (dc _DetectHumanBodyPose3DRequestClass) Alloc() DetectHumanBodyPose3DRequest {
	rv := objc.Send[DetectHumanBodyPose3DRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DetectHumanBodyPose3DRequestClass) New() DetectHumanBodyPose3DRequest {
	rv := objc.Send[DetectHumanBodyPose3DRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectHumanBodyPose3DRequest) Init() DetectHumanBodyPose3DRequest {
	rv := objc.Send[DetectHumanBodyPose3DRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectHumanBodyPose3DRequest) Autorelease() DetectHumanBodyPose3DRequest {
	rv := objc.Send[DetectHumanBodyPose3DRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectHumanBodyPose3DRequest creates a new DetectHumanBodyPose3DRequest instance.
func NewDetectHumanBodyPose3DRequest() DetectHumanBodyPose3DRequest {
	return getDetectHumanBodyPose3DRequestClass().New()
}





// A request that detects points on human bodies in 3D space, relative to the camera.
//
// This request generates a collection of objects that describe the position of each body the request detects. If the system allows it, the request uses information to improve the accuracy.


// A request that detects points on human bodies in 3D space, relative to the camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPose3DRequest
type DetectHumanBodyPose3DRequest struct {
	StatefulRequest
}

// DetectHumanBodyPose3DRequestFrom constructs a [DetectHumanBodyPose3DRequest] from an unsafe.Pointer.
//
// A request that detects points on human bodies in 3D space, relative to the camera.
func DetectHumanBodyPose3DRequestFrom(ptr unsafe.Pointer) DetectHumanBodyPose3DRequest {
	return DetectHumanBodyPose3DRequest{
		StatefulRequest: StatefulRequestFrom(ptr),
	}
}






// Creates a new 3D body pose request with a completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPose3DRequest/init(completionHandler:)
func NewDetectHumanBodyPose3DRequestWithCompletionHandler(completionHandler RequestCompletionHandler /* not a class type */) DetectHumanBodyPose3DRequest {
	instance := getDetectHumanBodyPose3DRequestClass().Alloc()
	rv := objc.Send[DetectHumanBodyPose3DRequest](instance.ID, objc.Sel("initWithCompletionHandler:"), completionHandler)
	rv.Autorelease()
	return rv
}

















// Returns the joint names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPose3DRequest/supportedJointNamesAndReturnError:
func (d_ DetectHumanBodyPose3DRequest) SupportedJointNamesAndReturnError(error_ foundation.foundation.INSError) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("supportedJointNamesAndReturnError:"), error_)
	return rv
}


// Returns the joint group names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPose3DRequest/supportedJointsGroupNamesAndReturnError:
func (d_ DetectHumanBodyPose3DRequest) SupportedJointsGroupNamesAndReturnError(error_ foundation.foundation.INSError) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("supportedJointsGroupNamesAndReturnError:"), error_)
	return rv
}







// The 3D body pose the request observes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPose3DRequest/results
func (d_ DetectHumanBodyPose3DRequest) Results() []HumanBodyPose3DObservation {
	rv := objc.Send[[]HumanBodyPose3DObservation](d_.ID, objc.Sel("results"))
	return rv
}


// Returns the joint group names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodypose3drequest/supportedjointnames
func (d_ DetectHumanBodyPose3DRequest) SupportedJointNames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("supportedJointNames"))
	return rv
}


// Returns the joint group names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodypose3drequest/supportedjointnames
func (d_ DetectHumanBodyPose3DRequest) SetSupportedJointNames(value objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointNames:"), value)
}


// Returns the joint names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodypose3drequest/supportedjointsgroupnames
func (d_ DetectHumanBodyPose3DRequest) SupportedJointsGroupNames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("supportedJointsGroupNames"))
	return rv
}


// Returns the joint names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodypose3drequest/supportedjointsgroupnames
func (d_ DetectHumanBodyPose3DRequest) SetSupportedJointsGroupNames(value objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointsGroupNames:"), value)
}







