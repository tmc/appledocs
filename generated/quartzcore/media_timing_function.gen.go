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
	mediaTimingFunctionClass     _MediaTimingFunctionClass
	mediaTimingFunctionClassOnce sync.Once
)

func getMediaTimingFunctionClass() _MediaTimingFunctionClass {
	mediaTimingFunctionClassOnce.Do(func() {
		mediaTimingFunctionClass = _MediaTimingFunctionClass{objc.GetClass("CAMediaTimingFunction")}
	})
	return mediaTimingFunctionClass
}

type _MediaTimingFunctionClass struct {
	class objc.Class
}

// An interface definition for the [MediaTimingFunction] class.
type IMediaTimingFunction interface {
	objectivec.IObject
}

// A function that defines the pacing of an animation as a timing curve.
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



