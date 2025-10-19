// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GlassEffectView] class.
var (
	glassEffectViewClass     _GlassEffectViewClass
	glassEffectViewClassOnce sync.Once
)

func getGlassEffectViewClass() _GlassEffectViewClass {
	glassEffectViewClassOnce.Do(func() {
		glassEffectViewClass = _GlassEffectViewClass{objc.GetClass("NSGlassEffectView")}
	})
	return glassEffectViewClass
}

type _GlassEffectViewClass struct {
	class objc.Class
}

// An interface definition for the [GlassEffectView] class.
type IGlassEffectView interface {
	IView
}

// A view that embeds its content view in a dynamic glass effect. [Full Topic]
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




