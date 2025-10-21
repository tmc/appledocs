// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// A request that detects points on human bodies in 3D space, relative to the camera.
//
// This request generates a collection of objects that describe the position of each body the request detects. If the system allows it, the request uses information to improve the accuracy.
//
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

// Alloc allocates a new instance without initialization.
func (dc _DetectHumanBodyPose3DRequestClass) Alloc() DetectHumanBodyPose3DRequest {
	rv := objc.Send[DetectHumanBodyPose3DRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns the joint group names the request supports.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodypose3drequest/supportedjointnames
func (d_ DetectHumanBodyPose3DRequest) SupportedJointNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("supportedJointNames"))
	return rv
}


// SetSupportedJointNames sets the value of the supportedJointNames property.
// Returns the joint group names the request supports.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodypose3drequest/supportedjointnames
func (d_ DetectHumanBodyPose3DRequest) SetSupportedJointNames(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointNames:"), value)
}

// The 3D body pose the request observes.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodypose3drequest/results
func (d_ DetectHumanBodyPose3DRequest) Results() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("results"))
	return rv
}


// SetResults sets the value of the results property.
// The 3D body pose the request observes.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodypose3drequest/results
func (d_ DetectHumanBodyPose3DRequest) SetResults(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setResults:"), value)
}

// Returns the joint names the request supports.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodypose3drequest/supportedjointsgroupnames
func (d_ DetectHumanBodyPose3DRequest) SupportedJointsGroupNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("supportedJointsGroupNames"))
	return rv
}


// SetSupportedJointsGroupNames sets the value of the supportedJointsGroupNames property.
// Returns the joint names the request supports.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodypose3drequest/supportedjointsgroupnames
func (d_ DetectHumanBodyPose3DRequest) SetSupportedJointsGroupNames(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointsGroupNames:"), value)
}



