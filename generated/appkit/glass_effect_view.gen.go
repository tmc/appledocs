// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GlassEffectView] class.
var glassEffectViewClass = _GlassEffectViewClass{objc.GetClass("NSGlassEffectView")}

type _GlassEffectViewClass struct {
	class objc.Class
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



