// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNDetectHumanBodyPoseRequest */


/* debug [class_header]: Header for VNDetectHumanBodyPoseRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DetectHumanBodyPoseRequest */
// An interface definition for the [DetectHumanBodyPoseRequest] class.
type IDetectHumanBodyPoseRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for DetectHumanBodyPoseRequest */
	// properties:
	Results() []HumanBodyPoseObservation
	SupportedJointNames() objectivec.IObject
	SetSupportedJointNames(value objectivec.IObject)
	SupportedJointsGroupNames() objectivec.IObject
	SetSupportedJointsGroupNames(value objectivec.IObject)
	VNDetectHumanBodyPoseRequestRevision1() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DetectHumanBodyPoseRequest */
	// methods:
	SupportedJointNamesAndReturnError(error_ objectivec.IObject) []string
	SupportedJointsGroupNamesAndReturnError(error_ objectivec.IObject) []string
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DetectHumanBodyPoseRequest */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DetectHumanBodyPoseRequest */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DetectHumanBodyPoseRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DetectHumanBodyPoseRequest */

// Retrieves the supported joint names for a revision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPoseRequest/supportedJointNames(forRevision:)
func (dc _DetectHumanBodyPoseRequestClass) SupportedJointNamesForRevisionError(revision uint, error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](objc.ID(dc.class), objc.Sel("supportedJointNamesForRevision:error:"), revision, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportedJointNamesForRevisionError) */


// Retrieves the supported joint group names for a revision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPoseRequest/supportedJointsGroupNames(forRevision:)
func (dc _DetectHumanBodyPoseRequestClass) SupportedJointsGroupNamesForRevisionError(revision uint, error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](objc.ID(dc.class), objc.Sel("supportedJointsGroupNamesForRevision:error:"), revision, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportedJointsGroupNamesForRevisionError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DetectHumanBodyPoseRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DetectHumanBodyPoseRequest */

// Retrieves the supported joint names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPoseRequest/supportedJointNamesAndReturnError:
func (d_ DetectHumanBodyPoseRequest) SupportedJointNamesAndReturnError(error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("supportedJointNamesAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: SupportedJointNamesAndReturnError */


// Retrieves the supported joint group names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPoseRequest/supportedJointsGroupNamesAndReturnError:
func (d_ DetectHumanBodyPoseRequest) SupportedJointsGroupNamesAndReturnError(error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("supportedJointsGroupNamesAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: SupportedJointsGroupNamesAndReturnError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DetectHumanBodyPoseRequest */

// The observed body poses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanBodyPoseRequest/results
func (d_ DetectHumanBodyPoseRequest) Results() []HumanBodyPoseObservation {
	rv := objc.Send[[]HumanBodyPoseObservation](d_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// Retrieves the supported joint names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodyposerequest/supportedjointnames
func (d_ DetectHumanBodyPoseRequest) SupportedJointNames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("supportedJointNames"))
	return rv
}/* debug [instance_properties/getter]: supportedJointNames */


// Retrieves the supported joint names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodyposerequest/supportedjointnames
func (d_ DetectHumanBodyPoseRequest) SetSupportedJointNames(value objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointNames:"), value)
}/* debug [instance_properties/setter]: supportedJointNames */


// Retrieves the supported joint group names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodyposerequest/supportedjointsgroupnames
func (d_ DetectHumanBodyPoseRequest) SupportedJointsGroupNames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("supportedJointsGroupNames"))
	return rv
}/* debug [instance_properties/getter]: supportedJointsGroupNames */


// Retrieves the supported joint group names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodyposerequest/supportedjointsgroupnames
func (d_ DetectHumanBodyPoseRequest) SetSupportedJointsGroupNames(value objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointsGroupNames:"), value)
}/* debug [instance_properties/setter]: supportedJointsGroupNames */


// A constant for specifying revision 1 of the body pose detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanbodyposerequestrevision1
func (d_ DetectHumanBodyPoseRequest) VNDetectHumanBodyPoseRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectHumanBodyPoseRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNDetectHumanBodyPoseRequestRevision1 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNDetectHumanBodyPoseRequest */



