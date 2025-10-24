// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNDetectAnimalBodyPoseRequest */


/* debug [class_header]: Header for VNDetectAnimalBodyPoseRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DetectAnimalBodyPoseRequest */
// An interface definition for the [DetectAnimalBodyPoseRequest] class.
type IDetectAnimalBodyPoseRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for DetectAnimalBodyPoseRequest */
	// properties:
	Results() []AnimalBodyPoseObservation
	SupportedJointNames() objectivec.IObject
	SetSupportedJointNames(value objectivec.IObject)
	SupportedJointsGroupNames() objectivec.IObject
	SetSupportedJointsGroupNames(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DetectAnimalBodyPoseRequest */
	// methods:
	SupportedJointNamesAndReturnError(error_ objectivec.IObject) []string
	SupportedJointsGroupNamesAndReturnError(error_ objectivec.IObject) []string
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DetectAnimalBodyPoseRequest */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DetectAnimalBodyPoseRequest */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DetectAnimalBodyPoseRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DetectAnimalBodyPoseRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DetectAnimalBodyPoseRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DetectAnimalBodyPoseRequest */

// Retrieves the joint names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectAnimalBodyPoseRequest/supportedJointNamesAndReturnError:
func (d_ DetectAnimalBodyPoseRequest) SupportedJointNamesAndReturnError(error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("supportedJointNamesAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: SupportedJointNamesAndReturnError */


// Retrieves the joint group names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectAnimalBodyPoseRequest/supportedJointsGroupNamesAndReturnError:
func (d_ DetectAnimalBodyPoseRequest) SupportedJointsGroupNamesAndReturnError(error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("supportedJointsGroupNamesAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: SupportedJointsGroupNamesAndReturnError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DetectAnimalBodyPoseRequest */

// The animal body pose the request observes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectAnimalBodyPoseRequest/results
func (d_ DetectAnimalBodyPoseRequest) Results() []AnimalBodyPoseObservation {
	rv := objc.Send[[]AnimalBodyPoseObservation](d_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// Retrieves the joint names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectanimalbodyposerequest/supportedjointnames
func (d_ DetectAnimalBodyPoseRequest) SupportedJointNames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("supportedJointNames"))
	return rv
}/* debug [instance_properties/getter]: supportedJointNames */


// Retrieves the joint names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectanimalbodyposerequest/supportedjointnames
func (d_ DetectAnimalBodyPoseRequest) SetSupportedJointNames(value objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointNames:"), value)
}/* debug [instance_properties/setter]: supportedJointNames */


// Retrieves the joint group names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectanimalbodyposerequest/supportedjointsgroupnames
func (d_ DetectAnimalBodyPoseRequest) SupportedJointsGroupNames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("supportedJointsGroupNames"))
	return rv
}/* debug [instance_properties/getter]: supportedJointsGroupNames */


// Retrieves the joint group names the request supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetectanimalbodyposerequest/supportedjointsgroupnames
func (d_ DetectAnimalBodyPoseRequest) SetSupportedJointsGroupNames(value objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointsGroupNames:"), value)
}/* debug [instance_properties/setter]: supportedJointsGroupNames */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNDetectAnimalBodyPoseRequest */



