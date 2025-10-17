// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GlyphInfo] class.
var GlyphInfoClass objc.Class

func init() {
	GlyphInfoClass = objc.GetClass("NSGlyphInfo")
}

type GlyphInfo struct {
	objc.ID
}

func GlyphInfoFrom(ptr unsafe.Pointer) GlyphInfo {
	return GlyphInfo{
		ID: objc.ID(ptr),
	}
}



