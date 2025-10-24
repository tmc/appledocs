// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StrokePath] class.
var (
	StrokePathClass     _StrokePathClass
	StrokePathClassOnce sync.Once
)

func getStrokePathClass() _StrokePathClass {
	StrokePathClassOnce.Do(func() {
		StrokePathClass = _StrokePathClass{objc.GetClass("PKStrokePath")}
	})
	return StrokePathClass
}

type _StrokePathClass struct {
	class objc.Class
}

// An interface definition for the [StrokePath] class.
type IStrokePath interface {
	objectivec.IObject
	// properties:
	Count() uint
	CreationDate() objc.IObject /* cross-framework: NSDate */
	// methods:
	EnumerateInterpolatedPointsInRangeStrideByDistanceUsingBlock(range_ IPKFloatRange, distanceStep float64, block unsafe.Pointer)
	EnumerateInterpolatedPointsInRangeStrideByParametricStepUsingBlock(range_ IPKFloatRange, parametricStep float64, block unsafe.Pointer)
	EnumerateInterpolatedPointsInRangeStrideByTimeUsingBlock(range_ IPKFloatRange, timeStep float64, block unsafe.Pointer)
	InterpolatedLocationAt(parametricValue float64) objc.IObject /* cross-framework: Point */
	InterpolatedPointAt(parametricValue float64) IStrokePoint
	ParametricValueOffsetByDistance(parametricValue float64, distanceStep float64) float64
	ParametricValueOffsetByTime(parametricValue float64, timeStep float64) float64
	PointAtIndex(i uint) IStrokePoint
	ObjectAtIndexedSubscript(i uint) IStrokePoint
}

// A class that captures the components of a stroke and provides methods to find and interpolate points along the stroke’s path.


// A class that captures the components of a stroke and provides methods to find and interpolate points along the stroke’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePathReference
type StrokePath struct {
	objectivec.Object
}

// StrokePathFrom constructs a [StrokePath] from an unsafe.Pointer.
//
// A class that captures the components of a stroke and provides methods to find and interpolate points along the stroke’s path.
func StrokePathFrom(ptr unsafe.Pointer) StrokePath {
	return StrokePath{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StrokePathClass) Alloc() StrokePath {
	rv := objc.Send[StrokePath](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StrokePathClass) New() StrokePath {
	rv := objc.Send[StrokePath](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StrokePath) Init() StrokePath {
	rv := objc.Send[StrokePath](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StrokePath) Autorelease() StrokePath {
	rv := objc.Send[StrokePath](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStrokePath creates a new StrokePath instance.
func NewStrokePath() StrokePath {
	return getStrokePathClass().New()
}



// Creates a stroke path with the cubic B-spline control points and a date that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePathReference/init(controlPoints:creationDate:)
func NewStrokePathWithControlPointsCreationDate(controlPoints []IStrokePoint, creationDate objc.IObject /* cross-framework: NSDate */) StrokePath {
	instance := getStrokePathClass().Alloc()
	rv := objc.Send[StrokePath](instance.ID, objc.Sel("initWithControlPoints:creationDate:"), controlPoints, creationDate)
	rv.Autorelease()
	return rv
}



// Executes a given block using each point in a range with a distance step.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePathReference/enumerateInterpolatedPoints(in:strideByDistance:using:)
func (s_ StrokePath) EnumerateInterpolatedPointsInRangeStrideByDistanceUsingBlock(range_ IPKFloatRange, distanceStep float64, block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateInterpolatedPointsInRange:strideByDistance:usingBlock:"), range_, distanceStep, block)
}


// Executes a given block using each point in a range with a parametric step.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePathReference/enumerateInterpolatedPoints(in:strideByParametricStep:using:)
func (s_ StrokePath) EnumerateInterpolatedPointsInRangeStrideByParametricStepUsingBlock(range_ IPKFloatRange, parametricStep float64, block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateInterpolatedPointsInRange:strideByParametricStep:usingBlock:"), range_, parametricStep, block)
}


// Executes a given block using each point in a range with a time step.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePathReference/enumerateInterpolatedPoints(in:strideByTime:using:)
func (s_ StrokePath) EnumerateInterpolatedPointsInRangeStrideByTimeUsingBlock(range_ IPKFloatRange, timeStep float64, block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateInterpolatedPointsInRange:strideByTime:usingBlock:"), range_, timeStep, block)
}


// Returns the on-curve point for the floating point parametric value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePathReference/interpolatedLocation(at:)
func (s_ StrokePath) InterpolatedLocationAt(parametricValue float64) objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[corefoundation.Point](s_.ID, objc.Sel("interpolatedLocationAt:"), parametricValue)
	return rv
}


// Returns the on-curve point for the provided floating point parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePathReference/interpolatedPoint(at:)
func (s_ StrokePath) InterpolatedPointAt(parametricValue float64) IStrokePoint {
	rv := objc.Send[StrokePoint](s_.ID, objc.Sel("interpolatedPointAt:"), parametricValue)
	return rv
}


// Returns a parametric value on the B-spline that’s a specified distance from the given parametric value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePathReference/parametricValue(_:offsetByDistance:)
func (s_ StrokePath) ParametricValueOffsetByDistance(parametricValue float64, distanceStep float64) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("parametricValue:offsetByDistance:"), parametricValue, distanceStep)
	return rv
}


// Returns a parametric value on the B-spline that’s a specified time from the given parametric value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePathReference/parametricValue(_:offsetByTime:)
func (s_ StrokePath) ParametricValueOffsetByTime(parametricValue float64, timeStep float64) float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("parametricValue:offsetByTime:"), parametricValue, timeStep)
	return rv
}


// Returns the B-spline control point at an index point that you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePathReference/point(at:)
func (s_ StrokePath) PointAtIndex(i uint) IStrokePoint {
	rv := objc.Send[StrokePoint](s_.ID, objc.Sel("pointAtIndex:"), i)
	return rv
}


// Returns the B-spline control point the location index that you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePathReference/subscript(_:)
func (s_ StrokePath) ObjectAtIndexedSubscript(i uint) IStrokePoint {
	rv := objc.Send[StrokePoint](s_.ID, objc.Sel("objectAtIndexedSubscript:"), i)
	return rv
}


// The number of control points in this stroke path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePathReference/count
func (s_ StrokePath) Count() uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("count"))
	return rv
}


// The time at which this stroke path starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePathReference/creationDate
func (s_ StrokePath) CreationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](s_.ID, objc.Sel("creationDate"))
	return rv
}


