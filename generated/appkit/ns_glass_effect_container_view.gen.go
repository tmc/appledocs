// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GlassEffectContainerView] class.
var (
	GlassEffectContainerViewClass     _GlassEffectContainerViewClass
	GlassEffectContainerViewClassOnce sync.Once
)

func getGlassEffectContainerViewClass() _GlassEffectContainerViewClass {
	GlassEffectContainerViewClassOnce.Do(func() {
		GlassEffectContainerViewClass = _GlassEffectContainerViewClass{objc.GetClass("NSGlassEffectContainerView")}
	})
	return GlassEffectContainerViewClass
}

type _GlassEffectContainerViewClass struct {
	class objc.Class
}

// An interface definition for the [GlassEffectContainerView] class.
type IGlassEffectContainerView interface {
	IView
	ContentView() IView
	SetContentView(value IView)
	Spacing() float64
	SetSpacing(value float64)
}

// A view that efficiently merges descendant glass effect views together when they are within a specified proximity to each other.


// A view that efficiently merges descendant glass effect views together when they are within a specified proximity to each other.
//
// [Full Topic]
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



// The view that contains descendant views to merge together when in proximity to each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectContainerView/contentView
func (g_ GlassEffectContainerView) ContentView() IView {
	rv := objc.Send[View](g_.ID, objc.Sel("contentView"))
	return rv
}


// The view that contains descendant views to merge together when in proximity to each other.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectContainerView/contentView
func (g_ GlassEffectContainerView) SetContentView(value IView) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setContentView:"), value)
}


// The proximity at which the glass effect container view begins merging eligible descendent glass effect views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectContainerView/spacing
func (g_ GlassEffectContainerView) Spacing() float64 {
	rv := objc.Send[float64](g_.ID, objc.Sel("spacing"))
	return rv
}


// The proximity at which the glass effect container view begins merging eligible descendent glass effect views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectContainerView/spacing
func (g_ GlassEffectContainerView) SetSpacing(value float64) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSpacing:"), value)
}



