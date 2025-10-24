// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	DetectedPoints() []Point
	EquationCoefficients() objectivec.IObject
	MovingAverageRadius() float64
	ProjectedPoints() []Point
	Results() IVNTrajectoryObservation
	SetResults(value IVNTrajectoryObservation)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (tc _TrajectoryObservationClass) Alloc() TrajectoryObservation {
	rv := objc.Send[TrajectoryObservation](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An observation that describes a detected trajectory.


// An observation that describes a detected trajectory.
//
// [Full Topic]
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

























// The centroid points of the detected contour along the trajectory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrajectoryObservation/detectedPoints
func (t_ TrajectoryObservation) DetectedPoints() []Point {
	rv := objc.Send[[]Point](t_.ID, objc.Sel("detectedPoints"))
	return rv
}


// The coefficients of the parabolic equation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrajectoryObservation/equationCoefficients
func (t_ TrajectoryObservation) EquationCoefficients() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("equationCoefficients"))
	return rv
}


// The moving average radius of the object the request is tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrajectoryObservation/movingAverageRadius
func (t_ TrajectoryObservation) MovingAverageRadius() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("movingAverageRadius"))
	return rv
}


// The centroids of the calculated trajectory from the detected points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNTrajectoryObservation/projectedPoints
func (t_ TrajectoryObservation) ProjectedPoints() []Point {
	rv := objc.Send[[]Point](t_.ID, objc.Sel("projectedPoints"))
	return rv
}


// The array of detected trajectory observations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/results
func (t_ TrajectoryObservation) Results() IVNTrajectoryObservation {
	rv := objc.Send[TrajectoryObservation](t_.ID, objc.Sel("results"))
	return rv
}


// The array of detected trajectory observations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vndetecttrajectoriesrequest/results
func (t_ TrajectoryObservation) SetResults(value IVNTrajectoryObservation) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setResults:"), value)
}








