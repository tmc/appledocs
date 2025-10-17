// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScrollLayer] class.
var scrollLayerClass = _ScrollLayerClass{objc.GetClass("CAScrollLayer")}

type _ScrollLayerClass struct {
	class objc.Class
}

// A layer that displays scrollable content larger than its own bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAScrollLayer

type ScrollLayer struct {
	Layer
}

// ScrollLayerFrom constructs a [ScrollLayer] from an unsafe.Pointer.
//
// A layer that displays scrollable content larger than its own bounds.
func ScrollLayerFrom(ptr unsafe.Pointer) ScrollLayer {
	return ScrollLayer{
		Layer: LayerFrom(ptr),
	}
}

// Scroll the contents of the receiver to ensure that the rectangle is visible. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAScrollLayer/scroll(to:)-782vd
func (s_ ScrollLayer) ScrollToRect(r unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("scrollToRect:"), r)
}


