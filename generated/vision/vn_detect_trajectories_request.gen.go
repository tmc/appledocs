// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// A request that detects the trajectories of shapes moving along a parabolic path.
//
// After the request detects a trajectory, it produces an observation that contains the shape’s detected points and an equation describing the parabola.
//
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


// The minimum radius of the bounding circle of the object to track.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/objectminimumnormalizedradius
func (d_ DetectTrajectoriesRequest) ObjectMinimumNormalizedRadius() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("objectMinimumNormalizedRadius"))
	return rv
}


// SetObjectMinimumNormalizedRadius sets the value of the objectMinimumNormalizedRadius property.
// The minimum radius of the bounding circle of the object to track.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/objectminimumnormalizedradius
func (d_ DetectTrajectoriesRequest) SetObjectMinimumNormalizedRadius(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setObjectMinimumNormalizedRadius:"), value)
}

// The maximum radius of the bounding circle of the object to track.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/objectmaximumnormalizedradius
func (d_ DetectTrajectoriesRequest) ObjectMaximumNormalizedRadius() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("objectMaximumNormalizedRadius"))
	return rv
}


// SetObjectMaximumNormalizedRadius sets the value of the objectMaximumNormalizedRadius property.
// The maximum radius of the bounding circle of the object to track.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/objectmaximumnormalizedradius
func (d_ DetectTrajectoriesRequest) SetObjectMaximumNormalizedRadius(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setObjectMaximumNormalizedRadius:"), value)
}

// The number of points to detect before calculating a trajectory.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/trajectorylength
func (d_ DetectTrajectoriesRequest) TrajectoryLength() int {
	rv := objc.Send[int](d_.ID, objc.Sel("trajectoryLength"))
	return rv
}


// SetTrajectoryLength sets the value of the trajectoryLength property.
// The number of points to detect before calculating a trajectory.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/trajectorylength
func (d_ DetectTrajectoriesRequest) SetTrajectoryLength(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTrajectoryLength:"), value)
}

// The requested target frame time for processing trajectory detection.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/targetframetime
func (d_ DetectTrajectoriesRequest) TargetFrameTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("targetFrameTime"))
	return rv
}


// SetTargetFrameTime sets the value of the targetFrameTime property.
// The requested target frame time for processing trajectory detection.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/targetframetime
func (d_ DetectTrajectoriesRequest) SetTargetFrameTime(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTargetFrameTime:"), value)
}

// The maximum radius of the tracked shape’s bounding circle.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/maximumobjectsize
func (d_ DetectTrajectoriesRequest) MaximumObjectSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("maximumObjectSize"))
	return rv
}


// SetMaximumObjectSize sets the value of the maximumObjectSize property.
// The maximum radius of the tracked shape’s bounding circle.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/maximumobjectsize
func (d_ DetectTrajectoriesRequest) SetMaximumObjectSize(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaximumObjectSize:"), value)
}

// The minimum radius of the tracked shape’s bounding circle.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/minimumobjectsize
func (d_ DetectTrajectoriesRequest) MinimumObjectSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("minimumObjectSize"))
	return rv
}


// SetMinimumObjectSize sets the value of the minimumObjectSize property.
// The minimum radius of the tracked shape’s bounding circle.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/minimumobjectsize
func (d_ DetectTrajectoriesRequest) SetMinimumObjectSize(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinimumObjectSize:"), value)
}

// A constant for specifying revision 1 of the trajectories detection request.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequestrevision1
func (d_ DetectTrajectoriesRequest) VNDetectTrajectoriesRequestRevision1() int {
	rv := objc.Send[int](d_.ID, objc.Sel("VNDetectTrajectoriesRequestRevision1"))
	return rv
}

// The array of detected trajectory observations.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/results
func (d_ DetectTrajectoriesRequest) Results() []TrajectoryObservation {
	rv := objc.Send[[]TrajectoryObservation](d_.ID, objc.Sel("results"))
	return rv
}



