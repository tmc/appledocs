// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Stroke] class.
type IStroke interface {
	objectivec.IObject
}

// A class that represents the paths, boundaries and other properties of a stroke drawn on a canvas.
//
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

// Alloc allocates a new instance without initialization.
func (sc _StrokeClass) Alloc() Stroke {
	rv := objc.Send[Stroke](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a stroke with the line properties, path, transform, and mask that you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/init(ink:strokePath:transform:mask:)
func NewStrokeWithInkStrokePathTransformMask(ink unsafe.Pointer, strokePath unsafe.Pointer, transform coregraphics.CGAffineTransform, mask unsafe.Pointer) Stroke {
	instance := getStrokeClass().Alloc()
	rv := objc.Send[Stroke](instance.ID, objc.Sel("initWithInk:strokePath:transform:mask:"), ink, strokePath, transform, mask)
	rv.Autorelease()
	return rv
}



// Creates a stroke with the line properties, path, transform, mask, and random seed that you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/init(ink:strokePath:transform:mask:randomSeed:)
func NewStrokeWithInkStrokePathTransformMaskRandomSeed(ink unsafe.Pointer, strokePath unsafe.Pointer, transform coregraphics.CGAffineTransform, mask unsafe.Pointer, randomSeed unsafe.Pointer) Stroke {
	instance := getStrokeClass().Alloc()
	rv := objc.Send[Stroke](instance.ID, objc.Sel("initWithInk:strokePath:transform:mask:randomSeed:"), ink, strokePath, transform, mask, randomSeed)
	rv.Autorelease()
	return rv
}


// The line properties used to render this stroke.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/ink
func (s_ Stroke) Ink() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("ink"))
	return rv
}

// The pretransform mask used to clip the rendering of the stroke.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/mask
func (s_ Stroke) Mask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("mask"))
	return rv
}

// The range of points in the stroke path reference that intersect the stroke’s mask.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/maskedPathRanges
func (s_ Stroke) MaskedPathRanges() []FloatRange {
	rv := objc.Send[[]FloatRange](s_.ID, objc.Sel("maskedPathRanges"))
	return rv
}

// The B-spline path that describes this stroke.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/path
func (s_ Stroke) Path() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("path"))
	return rv
}

// An unsigned 32-bit integer to use as a random seed for drawing strokes that use randomized effects.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/randomSeed
func (s_ Stroke) RandomSeed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("randomSeed"))
	return rv
}

// The bounds of the rendered stroke, including the width and line properties of the stroke after applying the transform.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/renderBounds
func (s_ Stroke) RenderBounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("renderBounds"))
	return rv
}

// The version of PencilKit necessary to use the stroke.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/requiredContentVersion
func (s_ Stroke) RequiredContentVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("requiredContentVersion"))
	return rv
}

// The affine transform of the stroke after rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKStrokeReference/transform
func (s_ Stroke) Transform() coregraphics.CGAffineTransform {
	rv := objc.Send[coregraphics.CGAffineTransform](s_.ID, objc.Sel("transform"))
	return rv
}


