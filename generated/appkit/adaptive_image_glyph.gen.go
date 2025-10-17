// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AdaptiveImageGlyph] class.
var adaptiveImageGlyphClass = _AdaptiveImageGlyphClass{objc.GetClass("NSAdaptiveImageGlyph")}

type _AdaptiveImageGlyphClass struct {
	class objc.Class
}

// An interface definition for the [AdaptiveImageGlyph] class.
type IAdaptiveImageGlyph interface {
	objectivec.IObject
}

// A data object for an emoji-like image that can appear in attributed text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph

type AdaptiveImageGlyph struct {
	objectivec.Object
}

// AdaptiveImageGlyphFrom constructs a [AdaptiveImageGlyph] from an unsafe.Pointer.
//
// A data object for an emoji-like image that can appear in attributed text.
func AdaptiveImageGlyphFrom(ptr unsafe.Pointer) AdaptiveImageGlyph {
	return AdaptiveImageGlyph{objectivec.Object{objc.ID(ptr)}}
}



