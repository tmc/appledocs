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


// The array of detected trajectory observations.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNDetectTrajectoriesRequest/results
func (d_ DetectTrajectoriesRequest) Results() []TrajectoryObservation {
	rv := objc.Send[[]TrajectoryObservation](d_.ID, objc.Sel("results"))
	return rv
}



