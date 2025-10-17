// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [GlassEffectContainerView] class.
var glassEffectContainerViewClass = _GlassEffectContainerViewClass{objc.GetClass("NSGlassEffectContainerView")}

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



