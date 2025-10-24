// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PKStroke */


/* debug [class_header]: Header for PKStroke */
// The class instance for the [Stroke] class.
var (
	StrokeClass     _StrokeClass
	StrokeClassOnce sync.Once
)

func getStrokeClass() _StrokeClass {
	StrokeClassOnce.Do(func() {
		StrokeClass = _StrokeClass{objc.GetClass("PKStroke")}
	})
	return StrokeClass
}

type _StrokeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Stroke */
// An interface definition for the [Stroke] class.
type IStroke interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Stroke */
	// properties:
	Ink() IPKInk
	Mask() appkit.BezierPath
	MaskedPathRanges() []FloatRange
	Path() IPKStrokePath
	RandomSeed() uint32 /* not a class type */
	RenderBounds() corefoundation.CGRect
	RequiredContentVersion() ContentVersion
	Transform() corefoundation.CGAffineTransform
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Stroke */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Stroke */
// Alloc allocates a new instance without initialization.
func (sc _StrokeClass) Alloc() Stroke {
	rv := objc.Send[Stroke](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StrokeClass) New() Stroke {
	rv := objc.Send[Stroke](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Stroke) Init() Stroke {
	rv := objc.Send[Stroke](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Stroke) Autorelease() Stroke {
	rv := objc.Send[Stroke](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStroke creates a new Stroke instance.
func NewStroke() Stroke {
	return getStrokeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Stroke */
// A class that represents the paths, boundaries and other properties of a stroke drawn on a canvas.


// A class that represents the paths, boundaries and other properties of a stroke drawn on a canvas.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference
type Stroke struct {
	objectivec.Object
}

// StrokeFrom constructs a [Stroke] from an unsafe.Pointer.
//
// A class that represents the paths, boundaries and other properties of a stroke drawn on a canvas.
func StrokeFrom(ptr unsafe.Pointer) Stroke {
	return Stroke{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Stroke */

// Creates a stroke with the line properties, path, transform, and mask that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/init(ink:strokePath:transform:mask:)
func NewStrokeWithInkStrokePathTransformMask(ink IPKInk, strokePath IPKStrokePath, transform corefoundation.CGAffineTransform, mask appkit.BezierPath) Stroke {
	instance := getStrokeClass().Alloc()
	rv := objc.Send[Stroke](instance.ID, objc.Sel("initWithInk:strokePath:transform:mask:"), ink, strokePath, transform, mask)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewStrokeWithInkStrokePathTransformMask */


// Creates a stroke with the line properties, path, transform, mask, and random seed that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/init(ink:strokePath:transform:mask:randomSeed:)
func NewStrokeWithInkStrokePathTransformMaskRandomSeed(ink IPKInk, strokePath IPKStrokePath, transform corefoundation.CGAffineTransform, mask appkit.BezierPath, randomSeed uint32 /* not a class type */) Stroke {
	instance := getStrokeClass().Alloc()
	rv := objc.Send[Stroke](instance.ID, objc.Sel("initWithInk:strokePath:transform:mask:randomSeed:"), ink, strokePath, transform, mask, randomSeed)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewStrokeWithInkStrokePathTransformMaskRandomSeed */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Stroke */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Stroke */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Stroke */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Stroke */

// The line properties used to render this stroke.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/ink
func (s_ Stroke) Ink() IPKInk {
	rv := objc.Send[Ink](s_.ID, objc.Sel("ink"))
	return rv
}/* debug [instance_properties/getter]: ink */


// The pretransform mask used to clip the rendering of the stroke.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/mask
func (s_ Stroke) Mask() appkit.BezierPath {
	rv := objc.Send[appkit.BezierPath](s_.ID, objc.Sel("mask"))
	return rv
}/* debug [instance_properties/getter]: mask */


// The range of points in the stroke path reference that intersect the stroke’s mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/maskedPathRanges
func (s_ Stroke) MaskedPathRanges() []FloatRange {
	rv := objc.Send[[]FloatRange](s_.ID, objc.Sel("maskedPathRanges"))
	return rv
}/* debug [instance_properties/getter]: maskedPathRanges */


// The B-spline path that describes this stroke.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/path
func (s_ Stroke) Path() IPKStrokePath {
	rv := objc.Send[StrokePath](s_.ID, objc.Sel("path"))
	return rv
}/* debug [instance_properties/getter]: path */


// An unsigned 32-bit integer to use as a random seed for drawing strokes that use randomized effects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/randomSeed
func (s_ Stroke) RandomSeed() uint32 /* not a class type */ {
	rv := objc.Send[uint32](s_.ID, objc.Sel("randomSeed"))
	return rv
}/* debug [instance_properties/getter]: randomSeed */


// The bounds of the rendered stroke, including the width and line properties of the stroke after applying the transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/renderBounds
func (s_ Stroke) RenderBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s_.ID, objc.Sel("renderBounds"))
	return rv
}/* debug [instance_properties/getter]: renderBounds */


// The version of PencilKit necessary to use the stroke.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/requiredContentVersion
func (s_ Stroke) RequiredContentVersion() ContentVersion {
	rv := objc.Send[ContentVersion](s_.ID, objc.Sel("requiredContentVersion"))
	return rv
}/* debug [instance_properties/getter]: requiredContentVersion */


// The affine transform of the stroke after rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/transform
func (s_ Stroke) Transform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](s_.ID, objc.Sel("transform"))
	return rv
}/* debug [instance_properties/getter]: transform */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKStroke */


