
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GlyphGenerator] class.
var GlyphGeneratorClass _GlyphGeneratorClass

func init() {
	GlyphGeneratorClass = _GlyphGeneratorClass{objc.GetClass("NSGlyphGenerator")}
}

type _GlyphGeneratorClass struct {
	objc.Class
}

// An interface definition for the [GlyphGenerator] class.
type IGlyphGenerator interface {
	ID() objc.ID
}

type GlyphGenerator struct {
	id objc.ID
}

func GlyphGeneratorFrom(ptr unsafe.Pointer) GlyphGenerator {
	return GlyphGenerator{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (g_ GlyphGenerator) ID() objc.ID {
	return g_.id
}

// Alloc allocates a new instance without initialization.
func (gc _GlyphGeneratorClass) Alloc() GlyphGenerator {
	rv := objc.Send[GlyphGenerator](objc.ID(gc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (gc _GlyphGeneratorClass) New() GlyphGenerator {
	rv := objc.Send[GlyphGenerator](objc.ID(gc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewGlyphGenerator creates and returns a new initialized instance.
func NewGlyphGenerator() GlyphGenerator {
	return GlyphGeneratorClass.New()
}

// Init initializes the instance.
func (g_ GlyphGenerator) Init() GlyphGenerator {
	rv := objc.Send[GlyphGenerator](g_.ID(), selInit)
	return rv
}
