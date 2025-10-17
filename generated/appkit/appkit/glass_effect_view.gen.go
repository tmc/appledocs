// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GlassEffectView] class.
var GlassEffectViewClass objc.Class

func init() {
	GlassEffectViewClass = objc.GetClass("NSGlassEffectView")
}

type GlassEffectView struct {
	objc.ID
}

func GlassEffectViewFrom(ptr unsafe.Pointer) GlassEffectView {
	return GlassEffectView{
		ID: objc.ID(ptr),
	}
}




