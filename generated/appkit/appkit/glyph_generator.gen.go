// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GlyphGenerator] class.
var GlyphGeneratorClass objc.Class

func init() {
	GlyphGeneratorClass = objc.GetClass("NSGlyphGenerator")
}

type GlyphGenerator struct {
	objc.ID
}

func GlyphGeneratorFrom(ptr unsafe.Pointer) GlyphGenerator {
	return GlyphGenerator{
		ID: objc.ID(ptr),
	}
}




