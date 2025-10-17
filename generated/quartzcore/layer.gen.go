// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Layer] class.
var layerClass = _LayerClass{objc.GetClass("CALayer")}

type _LayerClass struct {
	class objc.Class
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

// Draws the layer’s content using the specified graphics context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/draw(in:)
func (l_ Layer) DrawInContext(ctx unsafe.Pointer) {
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


