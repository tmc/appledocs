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

// An interface definition for the [ScrollLayer] class.
type IScrollLayer interface {
	ILayer
	ScrollToRect(r unsafe.Pointer)
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
// Alloc allocates a new instance without initialization.
func (sc _ScrollLayerClass) Alloc() ScrollLayer {
	rv := objc.Send[ScrollLayer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
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
	return scrollLayerClass.New()
}


// Scroll the contents of the receiver to ensure that the rectangle is visible. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAScrollLayer/scroll(to:)-782vd
func (s_ ScrollLayer) ScrollToRect(r unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("scrollToRect:"), r)
}


