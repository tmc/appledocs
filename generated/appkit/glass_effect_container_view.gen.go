// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GlassEffectContainerView] class.
var (
	glassEffectContainerViewClass     _GlassEffectContainerViewClass
	glassEffectContainerViewClassOnce sync.Once
)

func getGlassEffectContainerViewClass() _GlassEffectContainerViewClass {
	glassEffectContainerViewClassOnce.Do(func() {
		glassEffectContainerViewClass = _GlassEffectContainerViewClass{objc.GetClass("NSGlassEffectContainerView")}
	})
	return glassEffectContainerViewClass
}

type _GlassEffectContainerViewClass struct {
	class objc.Class
}

// An interface definition for the [GlassEffectContainerView] class.
type IGlassEffectContainerView interface {
	IView
}

// A view that efficiently merges descendant glass effect views together when they are within a specified proximity to each other. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectContainerView
type GlassEffectContainerView struct {
	View
}

// GlassEffectContainerViewFrom constructs a [GlassEffectContainerView] from an unsafe.Pointer.
//
// A view that efficiently merges descendant glass effect views together when they are within a specified proximity to each other.
func GlassEffectContainerViewFrom(ptr unsafe.Pointer) GlassEffectContainerView {
	return GlassEffectContainerView{
		View: ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GlassEffectContainerViewClass) Alloc() GlassEffectContainerView {
	rv := objc.Send[GlassEffectContainerView](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GlassEffectContainerViewClass) New() GlassEffectContainerView {
	rv := objc.Send[GlassEffectContainerView](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GlassEffectContainerView) Init() GlassEffectContainerView {
	rv := objc.Send[GlassEffectContainerView](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GlassEffectContainerView) Autorelease() GlassEffectContainerView {
	rv := objc.Send[GlassEffectContainerView](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGlassEffectContainerView creates a new GlassEffectContainerView instance.
func NewGlassEffectContainerView() GlassEffectContainerView {
	return getGlassEffectContainerViewClass().New()
}




