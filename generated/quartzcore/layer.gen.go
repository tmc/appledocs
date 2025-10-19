// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [Layer] class.
var (
	layerClass     _LayerClass
	layerClassOnce sync.Once
)

func getLayerClass() _LayerClass {
	layerClassOnce.Do(func() {
		layerClass = _LayerClass{objc.GetClass("CALayer")}
	})
	return layerClass
}

type _LayerClass struct {
	class objc.Class
}

// An interface definition for the [Layer] class.
type ILayer interface {
	objectivec.IObject
	DrawInContext(ctx coregraphics.CGContextRef)
	HitTest(p unsafe.Pointer) unsafe.Pointer
	SetNeedsDisplay()
}

// An object that manages image-based content and allows you to perform animations on that content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer
type Layer struct {
	objectivec.Object
}

// LayerFrom constructs a [Layer] from an unsafe.Pointer.
//
// An object that manages image-based content and allows you to perform animations on that content.
func LayerFrom(ptr unsafe.Pointer) Layer {
	return Layer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LayerClass) Alloc() Layer {
	rv := objc.Send[Layer](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LayerClass) New() Layer {
	rv := objc.Send[Layer](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ Layer) Init() Layer {
	rv := objc.Send[Layer](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ Layer) Autorelease() Layer {
	rv := objc.Send[Layer](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayer creates a new Layer instance.
func NewLayer() Layer {
	return getLayerClass().New()
}


// Draws the layer’s content using the specified graphics context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/draw(in:)
func (l_ Layer) DrawInContext(ctx coregraphics.CGContextRef) {
	objc.Send[objc.ID](l_.ID, objc.Sel("drawInContext:"), ctx)
}
// Returns the farthest descendant of the receiver in the layer hierarchy (including itself) that contains the specified point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/hitTest(_:)
func (l_ Layer) HitTest(p unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("hitTest:"), p)
	return rv
}
// Marks the layer’s contents as needing to be updated. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/setNeedsDisplay()
func (l_ Layer) SetNeedsDisplay() {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNeedsDisplay"))
}


