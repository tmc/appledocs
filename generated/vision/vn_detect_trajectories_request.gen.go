// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNDetectTrajectoriesRequest */


/* debug [class_header]: Header for VNDetectTrajectoriesRequest */
// The class instance for the [DetectTrajectoriesRequest] class.
var (
	DetectTrajectoriesRequestClass     _DetectTrajectoriesRequestClass
	DetectTrajectoriesRequestClassOnce sync.Once
)

func getDetectTrajectoriesRequestClass() _DetectTrajectoriesRequestClass {
	DetectTrajectoriesRequestClassOnce.Do(func() {
		DetectTrajectoriesRequestClass = _DetectTrajectoriesRequestClass{objc.GetClass("VNDetectTrajectoriesRequest")}
	})
	return DetectTrajectoriesRequestClass
}

type _DetectTrajectoriesRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DetectTrajectoriesRequest */
// An interface definition for the [DetectTrajectoriesRequest] class.
type IDetectTrajectoriesRequest interface {
	IStatefulRequest
	
/* debug [class_interface_properties]: Properties for DetectTrajectoriesRequest */
	// properties:
	MaximumObjectSize() float32
	SetMaximumObjectSize(value float32)
	MinimumObjectSize() float32
	SetMinimumObjectSize(value float32)
	ObjectMaximumNormalizedRadius() float32
	SetObjectMaximumNormalizedRadius(value float32)
	ObjectMinimumNormalizedRadius() float32
	SetObjectMinimumNormalizedRadius(value float32)
	Results() []TrajectoryObservation
	TargetFrameTime() objc.IObject /* cross-framework: Time */
	SetTargetFrameTime(value objc.IObject /* cross-framework: Time */)
	TrajectoryLength() int
	VNDetectTrajectoriesRequestRevision1() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DetectTrajectoriesRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DetectTrajectoriesRequest */
// Alloc allocates a new instance without initialization.
func (dc _DetectTrajectoriesRequestClass) Alloc() DetectTrajectoriesRequest {
	rv := objc.Send[DetectTrajectoriesRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DetectTrajectoriesRequestClass) New() DetectTrajectoriesRequest {
	rv := objc.Send[DetectTrajectoriesRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DetectTrajectoriesRequest) Init() DetectTrajectoriesRequest {
	rv := objc.Send[DetectTrajectoriesRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DetectTrajectoriesRequest) Autorelease() DetectTrajectoriesRequest {
	rv := objc.Send[DetectTrajectoriesRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetectTrajectoriesRequest creates a new DetectTrajectoriesRequest instance.
func NewDetectTrajectoriesRequest() DetectTrajectoriesRequest {
	return getDetectTrajectoriesRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DetectTrajectoriesRequest */
// A request that detects the trajectories of shapes moving along a parabolic path.
//
// After the request detects a trajectory, it produces an observation that contains the shape’s detected points and an equation describing the parabola.


// A request that detects the trajectories of shapes moving along a parabolic path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest
type DetectTrajectoriesRequest struct {
	StatefulRequest
}

// DetectTrajectoriesRequestFrom constructs a [DetectTrajectoriesRequest] from an unsafe.Pointer.
//
// A request that detects the trajectories of shapes moving along a parabolic path.
func DetectTrajectoriesRequestFrom(ptr unsafe.Pointer) DetectTrajectoriesRequest {
	return DetectTrajectoriesRequest{
		StatefulRequest: StatefulRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DetectTrajectoriesRequest */

// Creates a new request to detect trajectories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/init(frameAnalysisSpacing:trajectoryLength:completionHandler:)
func NewDetectTrajectoriesRequestWithFrameAnalysisSpacingTrajectoryLengthCompletionHandler(frameAnalysisSpacing objc.IObject /* cross-framework: Time */, trajectoryLength int, completionHandler RequestCompletionHandler /* not a class type */) DetectTrajectoriesRequest {
	instance := getDetectTrajectoriesRequestClass().Alloc()
	rv := objc.Send[DetectTrajectoriesRequest](instance.ID, objc.Sel("initWithFrameAnalysisSpacing:trajectoryLength:completionHandler:"), frameAnalysisSpacing, trajectoryLength, completionHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDetectTrajectoriesRequestWithFrameAnalysisSpacingTrajectoryLengthCompletionHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DetectTrajectoriesRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DetectTrajectoriesRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DetectTrajectoriesRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DetectTrajectoriesRequest */

// The maximum radius of the tracked shape’s bounding circle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/maximumObjectSize
func (d_ DetectTrajectoriesRequest) MaximumObjectSize() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("maximumObjectSize"))
	return rv
}/* debug [instance_properties/getter]: maximumObjectSize */


// The maximum radius of the tracked shape’s bounding circle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/maximumObjectSize
func (d_ DetectTrajectoriesRequest) SetMaximumObjectSize(value float32) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaximumObjectSize:"), value)
}/* debug [instance_properties/setter]: maximumObjectSize */


// The minimum radius of the tracked shape’s bounding circle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/minimumObjectSize
func (d_ DetectTrajectoriesRequest) MinimumObjectSize() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("minimumObjectSize"))
	return rv
}/* debug [instance_properties/getter]: minimumObjectSize */


// The minimum radius of the tracked shape’s bounding circle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/minimumObjectSize
func (d_ DetectTrajectoriesRequest) SetMinimumObjectSize(value float32) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinimumObjectSize:"), value)
}/* debug [instance_properties/setter]: minimumObjectSize */


// The maximum radius of the bounding circle of the object to track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/objectMaximumNormalizedRadius
func (d_ DetectTrajectoriesRequest) ObjectMaximumNormalizedRadius() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("objectMaximumNormalizedRadius"))
	return rv
}/* debug [instance_properties/getter]: objectMaximumNormalizedRadius */


// The maximum radius of the bounding circle of the object to track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/objectMaximumNormalizedRadius
func (d_ DetectTrajectoriesRequest) SetObjectMaximumNormalizedRadius(value float32) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setObjectMaximumNormalizedRadius:"), value)
}/* debug [instance_properties/setter]: objectMaximumNormalizedRadius */


// The minimum radius of the bounding circle of the object to track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/objectMinimumNormalizedRadius
func (d_ DetectTrajectoriesRequest) ObjectMinimumNormalizedRadius() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("objectMinimumNormalizedRadius"))
	return rv
}/* debug [instance_properties/getter]: objectMinimumNormalizedRadius */


// The minimum radius of the bounding circle of the object to track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/objectMinimumNormalizedRadius
func (d_ DetectTrajectoriesRequest) SetObjectMinimumNormalizedRadius(value float32) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setObjectMinimumNormalizedRadius:"), value)
}/* debug [instance_properties/setter]: objectMinimumNormalizedRadius */


// The array of detected trajectory observations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/results
func (d_ DetectTrajectoriesRequest) Results() []TrajectoryObservation {
	rv := objc.Send[[]TrajectoryObservation](d_.ID, objc.Sel("results"))
	return rv
}/* debug [instance_properties/getter]: results */


// The requested target frame time for processing trajectory detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/targetFrameTime
func (d_ DetectTrajectoriesRequest) TargetFrameTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](d_.ID, objc.Sel("targetFrameTime"))
	return rv
}/* debug [instance_properties/getter]: targetFrameTime */


// The requested target frame time for processing trajectory detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/targetFrameTime
func (d_ DetectTrajectoriesRequest) SetTargetFrameTime(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTargetFrameTime:"), value)
}/* debug [instance_properties/setter]: targetFrameTime */


// The number of points to detect before calculating a trajectory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/trajectoryLength
func (d_ DetectTrajectoriesRequest) TrajectoryLength() int {
	rv := objc.Send[int](d_.ID, objc.Sel("trajectoryLength"))
	return rv
}/* debug [instance_properties/getter]: trajectoryLength */


// A constant for specifying revision 1 of the trajectories detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequestrevision1
func (d_ DetectTrajectoriesRequest) VNDetectTrajectoriesRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectTrajectoriesRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNDetectTrajectoriesRequestRevision1 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNDetectTrajectoriesRequest */


