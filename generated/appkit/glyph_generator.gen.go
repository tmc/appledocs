// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GlyphGenerator] class.
var glyphGeneratorClass = _GlyphGeneratorClass{objc.GetClass("NSGlyphGenerator")}

type _GlyphGeneratorClass struct {
	class objc.Class
}

// An interface definition for the [GlyphGenerator] class.
type IGlyphGenerator interface {
	objectivec.IObject
}

// An object that performs the initial, nominal glyph generation phase in the layout process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphGenerator

type GlyphGenerator struct {
	objectivec.Object
}

// GlyphGeneratorFrom constructs a [GlyphGenerator] from an unsafe.Pointer.
//
// An object that performs the initial, nominal glyph generation phase in the layout process.
func GlyphGeneratorFrom(ptr unsafe.Pointer) GlyphGenerator {
	return GlyphGenerator{objectivec.Object{objc.ID(ptr)}}
}



