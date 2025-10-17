// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GlassEffectContainerView] class.
var GlassEffectContainerViewClass objc.Class

func init() {
	GlassEffectContainerViewClass = objc.GetClass("NSGlassEffectContainerView")
}

type GlassEffectContainerView struct {
	objc.ID
}

func GlassEffectContainerViewFrom(ptr unsafe.Pointer) GlassEffectContainerView {
	return GlassEffectContainerView{
		ID: objc.ID(ptr),
	}
}



