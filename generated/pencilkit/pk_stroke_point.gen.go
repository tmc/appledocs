// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PKStrokePoint */


/* debug [class_header]: Header for PKStrokePoint */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StrokePoint */
// An interface definition for the [StrokePoint] class.
type IStrokePoint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for StrokePoint */
	// properties:
	Altitude() float64
	Azimuth() float64
	Force() float64
	Location() corefoundation.CGPoint
	Opacity() float64
	SecondaryScale() float64
	Size() corefoundation.CGSize
	Threshold() float64
	TimeOffset() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StrokePoint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StrokePoint */
// Alloc allocates a new instance without initialization.
func (sc _StrokePointClass) Alloc() StrokePoint {
	rv := objc.Send[StrokePoint](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StrokePoint */
// A class that represents the properties of a specific point along a stroke’s path.


// A class that represents the properties of a specific point along a stroke’s path.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StrokePoint */

// Creates a new point with the provided properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/init(location:timeOffset:size:opacity:force:azimuth:altitude:)
func NewStrokePointWithLocationTimeOffsetSizeOpacityForceAzimuthAltitude(location corefoundation.CGPoint, timeOffset float64, size corefoundation.CGSize, opacity float64, force float64, azimuth float64, altitude float64) StrokePoint {
	instance := getStrokePointClass().Alloc()
	rv := objc.Send[StrokePoint](instance.ID, objc.Sel("initWithLocation:timeOffset:size:opacity:force:azimuth:altitude:"), location, timeOffset, size, opacity, force, azimuth, altitude)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewStrokePointWithLocationTimeOffsetSizeOpacityForceAzimuthAltitude */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/init(location:timeOffset:size:opacity:force:azimuth:altitude:secondaryScale:)
func NewStrokePointWithLocationTimeOffsetSizeOpacityForceAzimuthAltitudeSecondaryScale(location corefoundation.CGPoint, timeOffset float64, size corefoundation.CGSize, opacity float64, force float64, azimuth float64, altitude float64, secondaryScale float64) StrokePoint {
	instance := getStrokePointClass().Alloc()
	rv := objc.Send[StrokePoint](instance.ID, objc.Sel("initWithLocation:timeOffset:size:opacity:force:azimuth:altitude:secondaryScale:"), location, timeOffset, size, opacity, force, azimuth, altitude, secondaryScale)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewStrokePointWithLocationTimeOffsetSizeOpacityForceAzimuthAltitudeSecondaryScale */


// Create a new point with the provided properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/init(location:timeOffset:size:opacity:force:azimuth:altitude:secondaryScale:threshold:)
func NewStrokePointWithLocationTimeOffsetSizeOpacityForceAzimuthAltitudeSecondaryScaleThreshold(location corefoundation.CGPoint, timeOffset float64, size corefoundation.CGSize, opacity float64, force float64, azimuth float64, altitude float64, secondaryScale float64, threshold float64) StrokePoint {
	instance := getStrokePointClass().Alloc()
	rv := objc.Send[StrokePoint](instance.ID, objc.Sel("initWithLocation:timeOffset:size:opacity:force:azimuth:altitude:secondaryScale:threshold:"), location, timeOffset, size, opacity, force, azimuth, altitude, secondaryScale, threshold)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewStrokePointWithLocationTimeOffsetSizeOpacityForceAzimuthAltitudeSecondaryScaleThreshold */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StrokePoint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StrokePoint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StrokePoint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StrokePoint */

// The altitude of this point in radians.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/altitude
func (s_ StrokePoint) Altitude() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("altitude"))
	return rv
}/* debug [instance_properties/getter]: altitude */


// The azimuth of this point in radians.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/azimuth
func (s_ StrokePoint) Azimuth() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("azimuth"))
	return rv
}/* debug [instance_properties/getter]: azimuth */


// The amount of force applied by the touch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/force
func (s_ StrokePoint) Force() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("force"))
	return rv
}/* debug [instance_properties/getter]: force */


// The location of this point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/location
func (s_ StrokePoint) Location() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](s_.ID, objc.Sel("location"))
	return rv
}/* debug [instance_properties/getter]: location */


// Opacity of the point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/opacity
func (s_ StrokePoint) Opacity() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("opacity"))
	return rv
}/* debug [instance_properties/getter]: opacity */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/secondaryScale
func (s_ StrokePoint) SecondaryScale() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("secondaryScale"))
	return rv
}/* debug [instance_properties/getter]: secondaryScale */


// The size of the point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/size
func (s_ StrokePoint) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */


// The threshold for clipping the stroke rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/threshold
func (s_ StrokePoint) Threshold() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("threshold"))
	return rv
}/* debug [instance_properties/getter]: threshold */


// The time offset since the start of the stroke path in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokePointReference/timeOffset
func (s_ StrokePoint) TimeOffset() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("timeOffset"))
	return rv
}/* debug [instance_properties/getter]: timeOffset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKStrokePoint */


