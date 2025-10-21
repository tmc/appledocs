// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TrajectoryObservation] class.
var (
	TrajectoryObservationClass     _TrajectoryObservationClass
	TrajectoryObservationClassOnce sync.Once
)

func getTrajectoryObservationClass() _TrajectoryObservationClass {
	TrajectoryObservationClassOnce.Do(func() {
		TrajectoryObservationClass = _TrajectoryObservationClass{objc.GetClass("VNTrajectoryObservation")}
	})
	return TrajectoryObservationClass
}

type _TrajectoryObservationClass struct {
	class objc.Class
}

// An interface definition for the [TrajectoryObservation] class.
type ITrajectoryObservation interface {
	IObservation
}

// An observation that describes a detected trajectory.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrajectoryObservation
type TrajectoryObservation struct {
	Observation
}

// TrajectoryObservationFrom constructs a [TrajectoryObservation] from an unsafe.Pointer.
//
// An observation that describes a detected trajectory.
func TrajectoryObservationFrom(ptr unsafe.Pointer) TrajectoryObservation {
	return TrajectoryObservation{
		Observation: ObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TrajectoryObservationClass) Alloc() TrajectoryObservation {
	rv := objc.Send[TrajectoryObservation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TrajectoryObservationClass) New() TrajectoryObservation {
	rv := objc.Send[TrajectoryObservation](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TrajectoryObservation) Init() TrajectoryObservation {
	rv := objc.Send[TrajectoryObservation](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TrajectoryObservation) Autorelease() TrajectoryObservation {
	rv := objc.Send[TrajectoryObservation](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTrajectoryObservation creates a new TrajectoryObservation instance.
func NewTrajectoryObservation() TrajectoryObservation {
	return getTrajectoryObservationClass().New()
}


// The coefficients of the parabolic equation.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrajectoryobservation/equationcoefficients
func (t_ TrajectoryObservation) EquationCoefficients() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("equationCoefficients"))
	return rv
}


// SetEquationCoefficients sets the value of the equationCoefficients property.
// The coefficients of the parabolic equation.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrajectoryobservation/equationcoefficients
func (t_ TrajectoryObservation) SetEquationCoefficients(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEquationCoefficients:"), value)
}

// The centroid points of the detected contour along the trajectory.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrajectoryobservation/detectedpoints
func (t_ TrajectoryObservation) DetectedPoints() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("detectedPoints"))
	return rv
}


// SetDetectedPoints sets the value of the detectedPoints property.
// The centroid points of the detected contour along the trajectory.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrajectoryobservation/detectedpoints
func (t_ TrajectoryObservation) SetDetectedPoints(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDetectedPoints:"), value)
}

// The centroids of the calculated trajectory from the detected points.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrajectoryobservation/projectedpoints
func (t_ TrajectoryObservation) ProjectedPoints() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("projectedPoints"))
	return rv
}


// SetProjectedPoints sets the value of the projectedPoints property.
// The centroids of the calculated trajectory from the detected points.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrajectoryobservation/projectedpoints
func (t_ TrajectoryObservation) SetProjectedPoints(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setProjectedPoints:"), value)
}

// The array of detected trajectory observations.
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/results
func (t_ TrajectoryObservation) Results() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("results"))
	return rv
}


// SetResults sets the value of the results property.
// The array of detected trajectory observations.

//
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/results
func (t_ TrajectoryObservation) SetResults(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResults:"), value)
}

// The moving average radius of the object the request is tracking.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrajectoryObservation/movingAverageRadius
func (t_ TrajectoryObservation) MovingAverageRadius() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("movingAverageRadius"))
	return rv
}



