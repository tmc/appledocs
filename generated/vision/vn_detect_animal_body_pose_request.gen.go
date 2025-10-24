// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [DetectAnimalBodyPoseRequest] class.
var (
	DetectAnimalBodyPoseRequestClass     _DetectAnimalBodyPoseRequestClass
	DetectAnimalBodyPoseRequestClassOnce sync.Once
)

func getDetectAnimalBodyPoseRequestClass() _DetectAnimalBodyPoseRequestClass {
	DetectAnimalBodyPoseRequestClassOnce.Do(func() {
		DetectAnimalBodyPoseRequestClass = _DetectAnimalBodyPoseRequestClass{objc.GetClass("VNDetectAnimalBodyPoseRequest")}
	})
	return DetectAnimalBodyPoseRequestClass
}

type _DetectAnimalBodyPoseRequestClass struct {
	class objc.Class
}





// An interface definition for the [DetectAnimalBodyPoseRequest] class.
type IDetectAnimalBodyPoseRequest interface {
	IImageBasedRequest
	

	// properties:
	Results() []AnimalBodyPoseObservation
	SupportedJointNames() objectivec.IObject
	SetSupportedJointNames(value objectivec.IObject)
	SupportedJointsGroupNames() objectivec.IObject
	SetSupportedJointsGroupNames(value objectivec.IObject)


	

	// methods:
	SupportedJointNamesAndReturnError(error_ objectivec.IObject) []string
	SupportedJointsGroupNamesAndReturnError(error_ objectivec.IObject) []string


}





// Alloc allocates a new instance without initialization.
func (dc _DetectAnimalBodyPoseRequestClass) Alloc() DetectAnimalBodyPoseRequest {
	rv := objc.Send[DetectAnimalBodyPoseRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DetectAnimalBodyPoseRequestClass) New() DetectAnimalBodyPoseRequest {
	rv := objc.Send[DetectAnimalBodyPoseRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectAnimalBodyPoseRequest) Init() DetectAnimalBodyPoseRequest {
	rv := objc.Send[DetectAnimalBodyPoseRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectAnimalBodyPoseRequest) Autorelease() DetectAnimalBodyPoseRequest {
	rv := objc.Send[DetectAnimalBodyPoseRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectAnimalBodyPoseRequest creates a new DetectAnimalBodyPoseRequest instance.
func NewDetectAnimalBodyPoseRequest() DetectAnimalBodyPoseRequest {
	return getDetectAnimalBodyPoseRequestClass().New()
}





// A request that detects an animal body pose.


// A request that detects an animal body pose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectAnimalBodyPoseRequest
type DetectAnimalBodyPoseRequest struct {
	ImageBasedRequest
}

// DetectAnimalBodyPoseRequestFrom constructs a [DetectAnimalBodyPoseRequest] from an unsafe.Pointer.
//
// A request that detects an animal body pose.
func DetectAnimalBodyPoseRequestFrom(ptr unsafe.Pointer) DetectAnimalBodyPoseRequest {
	return DetectAnimalBodyPoseRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}




















// Retrieves the joint names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectAnimalBodyPoseRequest/supportedJointNamesAndReturnError:
func (d_ DetectAnimalBodyPoseRequest) SupportedJointNamesAndReturnError(error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("supportedJointNamesAndReturnError:"), error_)
	return rv
}


// Retrieves the joint group names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectAnimalBodyPoseRequest/supportedJointsGroupNamesAndReturnError:
func (d_ DetectAnimalBodyPoseRequest) SupportedJointsGroupNamesAndReturnError(error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("supportedJointsGroupNamesAndReturnError:"), error_)
	return rv
}







// The animal body pose the request observes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectAnimalBodyPoseRequest/results
func (d_ DetectAnimalBodyPoseRequest) Results() []AnimalBodyPoseObservation {
	rv := objc.Send[[]AnimalBodyPoseObservation](d_.ID, objc.Sel("results"))
	return rv
}


// Retrieves the joint names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectanimalbodyposerequest/supportedjointnames
func (d_ DetectAnimalBodyPoseRequest) SupportedJointNames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("supportedJointNames"))
	return rv
}


// Retrieves the joint names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectanimalbodyposerequest/supportedjointnames
func (d_ DetectAnimalBodyPoseRequest) SetSupportedJointNames(value objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointNames:"), value)
}


// Retrieves the joint group names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectanimalbodyposerequest/supportedjointsgroupnames
func (d_ DetectAnimalBodyPoseRequest) SupportedJointsGroupNames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("supportedJointsGroupNames"))
	return rv
}


// Retrieves the joint group names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectanimalbodyposerequest/supportedjointsgroupnames
func (d_ DetectAnimalBodyPoseRequest) SetSupportedJointsGroupNames(value objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointsGroupNames:"), value)
}








