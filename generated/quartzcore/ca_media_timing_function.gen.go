// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MediaTimingFunction] class.
var (
	MediaTimingFunctionClass     _MediaTimingFunctionClass
	MediaTimingFunctionClassOnce sync.Once
)

func getMediaTimingFunctionClass() _MediaTimingFunctionClass {
	MediaTimingFunctionClassOnce.Do(func() {
		MediaTimingFunctionClass = _MediaTimingFunctionClass{objc.GetClass("CAMediaTimingFunction")}
	})
	return MediaTimingFunctionClass
}

type _MediaTimingFunctionClass struct {
	class objc.Class
}

// An interface definition for the [MediaTimingFunction] class.
type IMediaTimingFunction interface {
	objectivec.IObject
	GetControlPointAtIndexValues(idx uintptr, ptr unsafe.Pointer)
}

// A function that defines the pacing of an animation as a timing curve.
//
// represents one segment of a function that defines the pacing of an animation as a timing curve. The function maps an input time normalized to the range to an output time also in the range . You can create a media timing function by supplying your own cubic Bézier curve control points using the method or by using one of the predefined timing functions.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction
type MediaTimingFunction struct {
	objectivec.Object
}

// MediaTimingFunctionFrom constructs a [MediaTimingFunction] from an unsafe.Pointer.
//
// A function that defines the pacing of an animation as a timing curve.
func MediaTimingFunctionFrom(ptr unsafe.Pointer) MediaTimingFunction {
	return MediaTimingFunction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaTimingFunctionClass) Alloc() MediaTimingFunction {
	rv := objc.Send[MediaTimingFunction](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaTimingFunctionClass) New() MediaTimingFunction {
	rv := objc.Send[MediaTimingFunction](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaTimingFunction) Init() MediaTimingFunction {
	rv := objc.Send[MediaTimingFunction](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaTimingFunction) Autorelease() MediaTimingFunction {
	rv := objc.Send[MediaTimingFunction](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaTimingFunction creates a new MediaTimingFunction instance.
func NewMediaTimingFunction() MediaTimingFunction {
	return getMediaTimingFunctionClass().New()
}


// Returns an initialized timing function modeled as a cubic Bézier curve using the specified control points.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction/init(controlPoints:_:_:_:)
func NewMediaTimingFunctionWithControlPoints(c1x float32, c1y float32, c2x float32, c2y float32) MediaTimingFunction {
	instance := getMediaTimingFunctionClass().Alloc()
	rv := objc.Send[MediaTimingFunction](instance.ID, objc.Sel("initWithControlPoints::::"), c1x, c1y, c2x, c2y)
	rv.Autorelease()
	return rv
}

// Creates and returns a new instance of configured with the predefined timing function specified by .
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction/init(name:)
func NewMediaTimingFunctionWithName(name unsafe.Pointer) MediaTimingFunction {
	rv := objc.Send[MediaTimingFunction](objc.ID(getMediaTimingFunctionClass().class), objc.Sel("functionWithName:"), name)
	return rv
}


// Creates and returns a new instance of timing function modeled as a cubic Bézier curve using the specified control points.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction/functionWithControlPoints::::
func (mc _MediaTimingFunctionClass) FunctionWithControlPoints(c1x float32, c1y float32, c2x float32, c2y float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("functionWithControlPoints::::"), c1x, c1y, c2x, c2y)
	return rv
}

// Creates and returns a new instance of configured with the predefined timing function specified by .
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction/init(name:)
func (mc _MediaTimingFunctionClass) FunctionWithName(name unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("functionWithName:"), name)
	return rv
}

// Returns the control point for the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction/getControlPoint(at:values:)
func (m_ MediaTimingFunction) GetControlPointAtIndexValues(idx uintptr, ptr unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getControlPointAtIndex:values:"), idx, ptr)
}


