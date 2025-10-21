// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [StrokePoint] class.
var (
	StrokePointClass     _StrokePointClass
	StrokePointClassOnce sync.Once
)

func getStrokePointClass() _StrokePointClass {
	StrokePointClassOnce.Do(func() {
		StrokePointClass = _StrokePointClass{objc.GetClass("PKStrokePoint")}
	})
	return StrokePointClass
}

type _StrokePointClass struct {
	class objc.Class
}

// An interface definition for the [StrokePoint] class.
type IStrokePoint interface {
	objectivec.IObject
}

// A class that represents the properties of a specific point along a stroke’s path.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference
type StrokePoint struct {
	objectivec.Object
}

// StrokePointFrom constructs a [StrokePoint] from an unsafe.Pointer.
//
// A class that represents the properties of a specific point along a stroke’s path.
func StrokePointFrom(ptr unsafe.Pointer) StrokePoint {
	return StrokePoint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StrokePointClass) Alloc() StrokePoint {
	rv := objc.Send[StrokePoint](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StrokePointClass) New() StrokePoint {
	rv := objc.Send[StrokePoint](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StrokePoint) Init() StrokePoint {
	rv := objc.Send[StrokePoint](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StrokePoint) Autorelease() StrokePoint {
	rv := objc.Send[StrokePoint](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStrokePoint creates a new StrokePoint instance.
func NewStrokePoint() StrokePoint {
	return getStrokePointClass().New()
}




// Creates a new point with the provided properties.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/init(location:timeOffset:size:opacity:force:azimuth:altitude:)
func NewStrokePointWithLocationTimeOffsetSizeOpacityForceAzimuthAltitude(location coregraphics.CGPoint, timeOffset TimeInterval, size coregraphics.CGSize, opacity float64, force float64, azimuth float64, altitude float64) StrokePoint {
	instance := getStrokePointClass().Alloc()
	rv := objc.Send[StrokePoint](instance.ID, objc.Sel("initWithLocation:timeOffset:size:opacity:force:azimuth:altitude:"), location, timeOffset, size, opacity, force, azimuth, altitude)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/init(location:timeOffset:size:opacity:force:azimuth:altitude:secondaryScale:)
func NewStrokePointWithLocationTimeOffsetSizeOpacityForceAzimuthAltitudeSecondaryScale(location coregraphics.CGPoint, timeOffset TimeInterval, size coregraphics.CGSize, opacity float64, force float64, azimuth float64, altitude float64, secondaryScale float64) StrokePoint {
	instance := getStrokePointClass().Alloc()
	rv := objc.Send[StrokePoint](instance.ID, objc.Sel("initWithLocation:timeOffset:size:opacity:force:azimuth:altitude:secondaryScale:"), location, timeOffset, size, opacity, force, azimuth, altitude, secondaryScale)
	rv.Autorelease()
	return rv
}



// Create a new point with the provided properties.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/init(location:timeOffset:size:opacity:force:azimuth:altitude:secondaryScale:threshold:)
func NewStrokePointWithLocationTimeOffsetSizeOpacityForceAzimuthAltitudeSecondaryScaleThreshold(location coregraphics.CGPoint, timeOffset TimeInterval, size coregraphics.CGSize, opacity float64, force float64, azimuth float64, altitude float64, secondaryScale float64, threshold float64) StrokePoint {
	instance := getStrokePointClass().Alloc()
	rv := objc.Send[StrokePoint](instance.ID, objc.Sel("initWithLocation:timeOffset:size:opacity:force:azimuth:altitude:secondaryScale:threshold:"), location, timeOffset, size, opacity, force, azimuth, altitude, secondaryScale, threshold)
	rv.Autorelease()
	return rv
}


// The altitude of this point in radians.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/altitude
func (s_ StrokePoint) Altitude() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("altitude"))
	return rv
}

// The azimuth of this point in radians.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/azimuth
func (s_ StrokePoint) Azimuth() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("azimuth"))
	return rv
}

// The amount of force applied by the touch.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/force
func (s_ StrokePoint) Force() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("force"))
	return rv
}

// The location of this point.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/location
func (s_ StrokePoint) Location() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](s_.ID, objc.Sel("location"))
	return rv
}

// Opacity of the point.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/opacity
func (s_ StrokePoint) Opacity() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("opacity"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/secondaryScale
func (s_ StrokePoint) SecondaryScale() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("secondaryScale"))
	return rv
}

// The size of the point.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/size
func (s_ StrokePoint) Size() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](s_.ID, objc.Sel("size"))
	return rv
}

// The threshold for clipping the stroke rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/threshold
func (s_ StrokePoint) Threshold() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("threshold"))
	return rv
}

// The time offset since the start of the stroke path in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/timeOffset
func (s_ StrokePoint) TimeOffset() TimeInterval {
	rv := objc.Send[TimeInterval](s_.ID, objc.Sel("timeOffset"))
	return rv
}


