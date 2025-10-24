// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coremedia"
)

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

// An interface definition for the [DetectTrajectoriesRequest] class.
type IDetectTrajectoriesRequest interface {
	IStatefulRequest
	// properties:
	MaximumObjectSize() float32
	SetMaximumObjectSize(value float32)
	MinimumObjectSize() float32
	SetMinimumObjectSize(value float32)
	ObjectMaximumNormalizedRadius() float32
	SetObjectMaximumNormalizedRadius(value float32)
	ObjectMinimumNormalizedRadius() float32
	SetObjectMinimumNormalizedRadius(value float32)
	Results() objc.IObject /* cross-framework: TrajectoryObservation */
	SetResults(value objc.IObject /* cross-framework: TrajectoryObservation */)
	TargetFrameTime() objc.IObject /* cross-framework: Time */
	SetTargetFrameTime(value objc.IObject /* cross-framework: Time */)
	TrajectoryLength() int
	SetTrajectoryLength(value int)
	VNDetectTrajectoriesRequestRevision1() int
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (dc _DetectTrajectoriesRequestClass) Alloc() DetectTrajectoriesRequest {
	rv := objc.Send[DetectTrajectoriesRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The maximum radius of the tracked shape’s bounding circle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/maximumobjectsize
func (d_ DetectTrajectoriesRequest) MaximumObjectSize() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("maximumObjectSize"))
	return rv
}


// The maximum radius of the tracked shape’s bounding circle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/maximumobjectsize
func (d_ DetectTrajectoriesRequest) SetMaximumObjectSize(value float32) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaximumObjectSize:"), value)
}


// The minimum radius of the tracked shape’s bounding circle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/minimumobjectsize
func (d_ DetectTrajectoriesRequest) MinimumObjectSize() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("minimumObjectSize"))
	return rv
}


// The minimum radius of the tracked shape’s bounding circle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/minimumobjectsize
func (d_ DetectTrajectoriesRequest) SetMinimumObjectSize(value float32) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinimumObjectSize:"), value)
}


// The maximum radius of the bounding circle of the object to track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/objectmaximumnormalizedradius
func (d_ DetectTrajectoriesRequest) ObjectMaximumNormalizedRadius() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("objectMaximumNormalizedRadius"))
	return rv
}


// The maximum radius of the bounding circle of the object to track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/objectmaximumnormalizedradius
func (d_ DetectTrajectoriesRequest) SetObjectMaximumNormalizedRadius(value float32) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setObjectMaximumNormalizedRadius:"), value)
}


// The minimum radius of the bounding circle of the object to track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/objectminimumnormalizedradius
func (d_ DetectTrajectoriesRequest) ObjectMinimumNormalizedRadius() float32 {
	rv := objc.Send[float32](d_.ID, objc.Sel("objectMinimumNormalizedRadius"))
	return rv
}


// The minimum radius of the bounding circle of the object to track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/objectminimumnormalizedradius
func (d_ DetectTrajectoriesRequest) SetObjectMinimumNormalizedRadius(value float32) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setObjectMinimumNormalizedRadius:"), value)
}


// The array of detected trajectory observations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/results
func (d_ DetectTrajectoriesRequest) Results() objc.IObject /* cross-framework: TrajectoryObservation */ {
	rv := objc.Send[TrajectoryObservation](d_.ID, objc.Sel("results"))
	return rv
}


// The array of detected trajectory observations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/results
func (d_ DetectTrajectoriesRequest) SetResults(value objc.IObject /* cross-framework: TrajectoryObservation */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setResults:"), value)
}


// The requested target frame time for processing trajectory detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/targetframetime
func (d_ DetectTrajectoriesRequest) TargetFrameTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[coremedia.Time](d_.ID, objc.Sel("targetFrameTime"))
	return rv
}


// The requested target frame time for processing trajectory detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/targetframetime
func (d_ DetectTrajectoriesRequest) SetTargetFrameTime(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTargetFrameTime:"), value)
}


// The number of points to detect before calculating a trajectory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/trajectorylength
func (d_ DetectTrajectoriesRequest) TrajectoryLength() int {
	rv := objc.Send[int](d_.ID, objc.Sel("trajectoryLength"))
	return rv
}


// The number of points to detect before calculating a trajectory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/trajectorylength
func (d_ DetectTrajectoriesRequest) SetTrajectoryLength(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTrajectoryLength:"), value)
}


// A constant for specifying revision 1 of the trajectories detection request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequestrevision1
func (d_ DetectTrajectoriesRequest) VNDetectTrajectoriesRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectTrajectoriesRequestRevision1"))
	return rv
}



