
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AdaptiveImageGlyph] class.
var AdaptiveImageGlyphClass _AdaptiveImageGlyphClass

func init() {
	AdaptiveImageGlyphClass = _AdaptiveImageGlyphClass{objc.GetClass("NSAdaptiveImageGlyph")}
}

type _AdaptiveImageGlyphClass struct {
	objc.Class
}

// An interface definition for the [AdaptiveImageGlyph] class.
type IAdaptiveImageGlyph interface {
	ID() objc.ID
}

type AdaptiveImageGlyph struct {
	id objc.ID
}

func AdaptiveImageGlyphFrom(ptr unsafe.Pointer) AdaptiveImageGlyph {
	return AdaptiveImageGlyph{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ AdaptiveImageGlyph) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _AdaptiveImageGlyphClass) Alloc() AdaptiveImageGlyph {
	rv := objc.Send[AdaptiveImageGlyph](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _AdaptiveImageGlyphClass) New() AdaptiveImageGlyph {
	rv := objc.Send[AdaptiveImageGlyph](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewAdaptiveImageGlyph creates and returns a new initialized instance.
func NewAdaptiveImageGlyph() AdaptiveImageGlyph {
	return AdaptiveImageGlyphClass.New()
}

// Init initializes the instance.
func (a_ AdaptiveImageGlyph) Init() AdaptiveImageGlyph {
	rv := objc.Send[AdaptiveImageGlyph](a_.ID(), selInit)
	return rv
}
