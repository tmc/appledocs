// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GlassEffectView] class.
var (
	GlassEffectViewClass     _GlassEffectViewClass
	GlassEffectViewClassOnce sync.Once
)

func getGlassEffectViewClass() _GlassEffectViewClass {
	GlassEffectViewClassOnce.Do(func() {
		GlassEffectViewClass = _GlassEffectViewClass{objc.GetClass("NSGlassEffectView")}
	})
	return GlassEffectViewClass
}

type _GlassEffectViewClass struct {
	class objc.Class
}

// An interface definition for the [GlassEffectView] class.
type IGlassEffectView interface {
	IView
}

// A view that embeds its content view in a dynamic glass effect.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView
type GlassEffectView struct {
	View
}

// GlassEffectViewFrom constructs a [GlassEffectView] from an unsafe.Pointer.
//
// A view that embeds its content view in a dynamic glass effect.
func GlassEffectViewFrom(ptr unsafe.Pointer) GlassEffectView {
	return GlassEffectView{
		View: ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GlassEffectViewClass) Alloc() GlassEffectView {
	rv := objc.Send[GlassEffectView](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GlassEffectViewClass) New() GlassEffectView {
	rv := objc.Send[GlassEffectView](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GlassEffectView) Init() GlassEffectView {
	rv := objc.Send[GlassEffectView](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GlassEffectView) Autorelease() GlassEffectView {
	rv := objc.Send[GlassEffectView](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGlassEffectView creates a new GlassEffectView instance.
func NewGlassEffectView() GlassEffectView {
	return getGlassEffectViewClass().New()
}


// The view to embed in glass.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/contentView
func (g_ GlassEffectView) ContentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("contentView"))
	return rv
}


// SetContentView sets the value of the contentView property.
// The view to embed in glass.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/contentView
func (g_ GlassEffectView) SetContentView(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setContentView:"), value)
}
// The amount of curvature for all corners of the glass.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/cornerRadius
func (g_ GlassEffectView) CornerRadius() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("cornerRadius"))
	return rv
}


// SetCornerRadius sets the value of the cornerRadius property.
// The amount of curvature for all corners of the glass.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/cornerRadius
func (g_ GlassEffectView) SetCornerRadius(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCornerRadius:"), value)
}
// The style of glass this view uses.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/style-swift.property
func (g_ GlassEffectView) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("style"))
	return rv
}


// SetStyle sets the value of the style property.
// The style of glass this view uses.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/style-swift.property
func (g_ GlassEffectView) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setStyle:"), value)
}
// The color the glass effect view uses to tint the background and glass effect toward.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/tintColor
func (g_ GlassEffectView) TintColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("tintColor"))
	return rv
}


// SetTintColor sets the value of the tintColor property.
// The color the glass effect view uses to tint the background and glass effect toward.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/tintColor
func (g_ GlassEffectView) SetTintColor(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTintColor:"), value)
}


