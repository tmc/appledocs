// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [ScrollLayer] class.
var (
	ScrollLayerClass     _ScrollLayerClass
	ScrollLayerClassOnce sync.Once
)

func getScrollLayerClass() _ScrollLayerClass {
	ScrollLayerClassOnce.Do(func() {
		ScrollLayerClass = _ScrollLayerClass{objc.GetClass("CAScrollLayer")}
	})
	return ScrollLayerClass
}

type _ScrollLayerClass struct {
	class objc.Class
}

// An interface definition for the [ScrollLayer] class.
type IScrollLayer interface {
	ILayer
	ScrollToPoint(p coregraphics.CGPoint)
	ScrollToRect(r coregraphics.CGRect)
}

// A layer that displays scrollable content larger than its own bounds.
//
// The class is a subclass of that simplifies displaying a portion of a layer. The extent of the scrollable area of the is defined by the layout of its sublayers. The visible portion of the layer content is set by specifying the origin as a point or a rectangular area of the contents to be displayed. does not provide keyboard or mouse event-handling, nor does it provide visible scrollers.
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

// Alloc allocates a new instance without initialization.
func (sc _ScrollLayerClass) Alloc() ScrollLayer {
	rv := objc.Send[ScrollLayer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScrollLayerClass) New() ScrollLayer {
	rv := objc.Send[ScrollLayer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrollLayer) Init() ScrollLayer {
	rv := objc.Send[ScrollLayer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrollLayer) Autorelease() ScrollLayer {
	rv := objc.Send[ScrollLayer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrollLayer creates a new ScrollLayer instance.
func NewScrollLayer() ScrollLayer {
	return getScrollLayerClass().New()
}


// Changes the origin of the receiver to the specified point.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAScrollLayer/scroll(to:)-37q0p
func (s_ ScrollLayer) ScrollToPoint(p coregraphics.CGPoint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("scrollToPoint:"), p)
}

// Scroll the contents of the receiver to ensure that the rectangle is visible.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAScrollLayer/scroll(to:)-782vd
func (s_ ScrollLayer) ScrollToRect(r coregraphics.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("scrollToRect:"), r)
}

// Defines the axes in which the layer may be scrolled.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAScrollLayer/scrollMode
func (s_ ScrollLayer) ScrollMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("scrollMode"))
	return rv
}


// SetScrollMode sets the value of the scrollMode property.
// Defines the axes in which the layer may be scrolled.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAScrollLayer/scrollMode
func (s_ ScrollLayer) SetScrollMode(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrollMode:"), value)
}


