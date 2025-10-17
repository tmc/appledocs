// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GlyphInfo] class.
var glyphInfoClass = _GlyphInfoClass{objc.GetClass("NSGlyphInfo")}

type _GlyphInfoClass struct {
	class objc.Class
}

// An interface definition for the [GlyphInfo] class.
type IGlyphInfo interface {
	objectivec.IObject
}

// A glyph attribute in an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo

type GlyphInfo struct {
	objectivec.Object
}

// GlyphInfoFrom constructs a [GlyphInfo] from an unsafe.Pointer.
//
// A glyph attribute in an attributed string.
func GlyphInfoFrom(ptr unsafe.Pointer) GlyphInfo {
	return GlyphInfo{objectivec.Object{objc.ID(ptr)}}
}



