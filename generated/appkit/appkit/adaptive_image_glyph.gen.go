// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AdaptiveImageGlyph] class.
var AdaptiveImageGlyphClass objc.Class

func init() {
	AdaptiveImageGlyphClass = objc.GetClass("NSAdaptiveImageGlyph")
}

type AdaptiveImageGlyph struct {
	objc.ID
}

func AdaptiveImageGlyphFrom(ptr unsafe.Pointer) AdaptiveImageGlyph {
	return AdaptiveImageGlyph{
		ID: objc.ID(ptr),
	}
}




