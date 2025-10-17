
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GlassEffectContainerView] class.
var GlassEffectContainerViewClass _GlassEffectContainerViewClass

func init() {
	GlassEffectContainerViewClass = _GlassEffectContainerViewClass{objc.GetClass("NSGlassEffectContainerView")}
}

type _GlassEffectContainerViewClass struct {
	objc.Class
}

// An interface definition for the [GlassEffectContainerView] class.
type IGlassEffectContainerView interface {
	ID() objc.ID
}

type GlassEffectContainerView struct {
	id objc.ID
}

func GlassEffectContainerViewFrom(ptr unsafe.Pointer) GlassEffectContainerView {
	return GlassEffectContainerView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (g_ GlassEffectContainerView) ID() objc.ID {
	return g_.id
}

// Alloc allocates a new instance without initialization.
func (gc _GlassEffectContainerViewClass) Alloc() GlassEffectContainerView {
	rv := objc.Send[GlassEffectContainerView](objc.ID(gc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (gc _GlassEffectContainerViewClass) New() GlassEffectContainerView {
	rv := objc.Send[GlassEffectContainerView](objc.ID(gc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewGlassEffectContainerView creates and returns a new initialized instance.
func NewGlassEffectContainerView() GlassEffectContainerView {
	return GlassEffectContainerViewClass.New()
}

// Init initializes the instance.
func (g_ GlassEffectContainerView) Init() GlassEffectContainerView {
	rv := objc.Send[GlassEffectContainerView](g_.ID(), selInit)
	return rv
}
// The view that contains descendant views to merge together when in proximity to each other. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGlassEffectContainerView/contentView
func (g_ GlassEffectContainerView) ContentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("contentView"))
	return rv
}
// SetContentView sets the value of the contentView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGlassEffectContainerView/contentView
func (g_ GlassEffectContainerView) SetContentView(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("setContentView:"), value)
}
// The proximity at which the glass effect container view begins merging eligible descendent glass effect views. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGlassEffectContainerView/spacing
func (g_ GlassEffectContainerView) Spacing() float64 {
	rv := objc.Send[float64](g_.ID(), objc.RegisterName("spacing"))
	return rv
}
// SetSpacing sets the value of the spacing property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGlassEffectContainerView/spacing
func (g_ GlassEffectContainerView) SetSpacing(value float64) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("setSpacing:"), value)
}
