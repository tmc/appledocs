
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GlassEffectView] class.
var GlassEffectViewClass _GlassEffectViewClass

func init() {
	GlassEffectViewClass = _GlassEffectViewClass{objc.GetClass("NSGlassEffectView")}
}

type _GlassEffectViewClass struct {
	objc.Class
}

// An interface definition for the [GlassEffectView] class.
type IGlassEffectView interface {
	ID() objc.ID
}

type GlassEffectView struct {
	id objc.ID
}

func GlassEffectViewFrom(ptr unsafe.Pointer) GlassEffectView {
	return GlassEffectView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (g_ GlassEffectView) ID() objc.ID {
	return g_.id
}

// Alloc allocates a new instance without initialization.
func (gc _GlassEffectViewClass) Alloc() GlassEffectView {
	rv := objc.Send[GlassEffectView](objc.ID(gc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (gc _GlassEffectViewClass) New() GlassEffectView {
	rv := objc.Send[GlassEffectView](objc.ID(gc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewGlassEffectView creates and returns a new initialized instance.
func NewGlassEffectView() GlassEffectView {
	return GlassEffectViewClass.New()
}

// Init initializes the instance.
func (g_ GlassEffectView) Init() GlassEffectView {
	rv := objc.Send[GlassEffectView](g_.ID(), selInit)
	return rv
}
// The view to embed in glass. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGlassEffectView/contentView
func (g_ GlassEffectView) ContentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("contentView"))
	return rv
}
// SetContentView sets the value of the contentView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGlassEffectView/contentView
func (g_ GlassEffectView) SetContentView(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("setContentView:"), value)
}
// The amount of curvature for all corners of the glass. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGlassEffectView/cornerRadius
func (g_ GlassEffectView) CornerRadius() float64 {
	rv := objc.Send[float64](g_.ID(), objc.RegisterName("cornerRadius"))
	return rv
}
// SetCornerRadius sets the value of the cornerRadius property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGlassEffectView/cornerRadius
func (g_ GlassEffectView) SetCornerRadius(value float64) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("setCornerRadius:"), value)
}
// The style of glass this view uses. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGlassEffectView/style-swift.property
func (g_ GlassEffectView) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("style"))
	return rv
}
// SetStyle sets the value of the style property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGlassEffectView/style-swift.property
func (g_ GlassEffectView) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("setStyle:"), value)
}
// The color the glass effect view uses to tint the background and glass effect toward. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGlassEffectView/tintColor
func (g_ GlassEffectView) TintColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID(), objc.RegisterName("tintColor"))
	return rv
}
// SetTintColor sets the value of the tintColor property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSGlassEffectView/tintColor
func (g_ GlassEffectView) SetTintColor(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID(), objc.RegisterName("setTintColor:"), value)
}
