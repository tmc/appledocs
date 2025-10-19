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
// Alloc allocates a new instance without initialization.
func (gc _GlyphGeneratorClass) Alloc() GlyphGenerator {
	rv := objc.Send[GlyphGenerator](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (gc _GlyphGeneratorClass) New() GlyphGenerator {
	rv := objc.Send[GlyphGenerator](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GlyphGenerator) Init() GlyphGenerator {
	rv := objc.Send[GlyphGenerator](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GlyphGenerator) Autorelease() GlyphGenerator {
	rv := objc.Send[GlyphGenerator](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGlyphGenerator creates a new GlyphGenerator instance.
func NewGlyphGenerator() GlyphGenerator {
	return glyphGeneratorClass.New()
}




