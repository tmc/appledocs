// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNDetectHumanHandPoseRequest */


/* debug [class_header]: Header for VNDetectHumanHandPoseRequest */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DetectHumanHandPoseRequest */
// An interface definition for the [DetectHumanHandPoseRequest] class.
type IDetectHumanHandPoseRequest interface {
	IImageBasedRequest
	
/* debug [class_interface_properties]: Properties for DetectHumanHandPoseRequest */
	// properties:
	MaximumHandCount() uint
	SetMaximumHandCount(value uint)
	Results() []HumanHandPoseObservation
	SupportedJointNames() objectivec.IObject
	SetSupportedJointNames(value objectivec.IObject)
	SupportedJointsGroupNames() objectivec.IObject
	SetSupportedJointsGroupNames(value objectivec.IObject)
	VNDetectHumanHandPoseRequestRevision1() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DetectHumanHandPoseRequest */
	// methods:
	SupportedJointNamesAndReturnError(error_ objectivec.IObject) []string
	SupportedJointsGroupNamesAndReturnError(error_ objectivec.IObject) []string
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DetectHumanHandPoseRequest */
// Alloc allocates a new instance without initialization.
func (dc _DetectHumanHandPoseRequestClass) Alloc() DetectHumanHandPoseRequest {
	rv := objc.Send[DetectHumanHandPoseRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DetectHumanHandPoseRequest */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DetectHumanHandPoseRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DetectHumanHandPoseRequest */

// Retrieves the supported joint names for a revision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/supportedJointNames(forRevision:)
func (dc _DetectHumanHandPoseRequestClass) SupportedJointNamesForRevisionError(revision uint, error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](objc.ID(dc.class), objc.Sel("supportedJointNamesForRevision:error:"), revision, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportedJointNamesForRevisionError) */


// Retrieves the supported joint group names for a revision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/supportedJointsGroupNames(forRevision:)
func (dc _DetectHumanHandPoseRequestClass) SupportedJointsGroupNamesForRevisionError(revision uint, error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](objc.ID(dc.class), objc.Sel("supportedJointsGroupNamesForRevision:error:"), revision, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportedJointsGroupNamesForRevisionError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DetectHumanHandPoseRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DetectHumanHandPoseRequest */

// Retrieves the supported joint names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/supportedJointNamesAndReturnError:
func (d_ DetectHumanHandPoseRequest) SupportedJointNamesAndReturnError(error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("supportedJointNamesAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: SupportedJointNamesAndReturnError */


// Retrieves the supported joint group names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/supportedJointsGroupNamesAndReturnError:
func (d_ DetectHumanHandPoseRequest) SupportedJointsGroupNamesAndReturnError(error_ objectivec.IObject) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("supportedJointsGroupNamesAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: SupportedJointsGroupNamesAndReturnError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DetectHumanHandPoseRequest */

// The maximum number of hands to detect in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/maximumHandCount
func (d_ DetectHumanHandPoseRequest) MaximumHandCount() uint {
	rv := objc.Send[uint](d_.ID, objc.Sel("maximumHandCount"))
	return rv
}/* debug [instance_properties/getter]: maximumHandCount */


// The maximum number of hands to detect in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/maximumHandCount
func (d_ DetectHumanHandPoseRequest) SetMaximumHandCount(value uint) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaximumHandCount:"), value)
}/* debug [instance_properties/setter]: maximumHandCount */


// The observed hand poses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectHumanHandPoseRequest/results
func (d_ DetectHumanHandPoseRequest) Results() []HumanHandPoseObservation {
	rv := objc.Send[[]HumanHandPoseObservation](d_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// Retrieves the supported joint names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/supportedjointnames
func (d_ DetectHumanHandPoseRequest) SupportedJointNames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("supportedJointNames"))
	return rv
}/* debug [instance_properties/getter]: supportedJointNames */


// Retrieves the supported joint names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/supportedjointnames
func (d_ DetectHumanHandPoseRequest) SetSupportedJointNames(value objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointNames:"), value)
}/* debug [instance_properties/setter]: supportedJointNames */


// Retrieves the supported joint group names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/supportedjointsgroupnames
func (d_ DetectHumanHandPoseRequest) SupportedJointsGroupNames() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("supportedJointsGroupNames"))
	return rv
}/* debug [instance_properties/getter]: supportedJointsGroupNames */


// Retrieves the supported joint group names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequest/supportedjointsgroupnames
func (d_ DetectHumanHandPoseRequest) SetSupportedJointsGroupNames(value objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSupportedJointsGroupNames:"), value)
}/* debug [instance_properties/setter]: supportedJointsGroupNames */


// A constant for specifying revision 1 of the hand pose detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecthumanhandposerequestrevision1
func (d_ DetectHumanHandPoseRequest) VNDetectHumanHandPoseRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectHumanHandPoseRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNDetectHumanHandPoseRequestRevision1 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNDetectHumanHandPoseRequest */



