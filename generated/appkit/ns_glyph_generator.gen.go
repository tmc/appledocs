// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GlyphGenerator] class.
var (
	GlyphGeneratorClass     _GlyphGeneratorClass
	GlyphGeneratorClassOnce sync.Once
)

func getGlyphGeneratorClass() _GlyphGeneratorClass {
	GlyphGeneratorClassOnce.Do(func() {
		GlyphGeneratorClass = _GlyphGeneratorClass{objc.GetClass("NSGlyphGenerator")}
	})
	return GlyphGeneratorClass
}

type _GlyphGeneratorClass struct {
	class objc.Class
}

// An interface definition for the [GlyphGenerator] class.
type IGlyphGenerator interface {
	objectivec.IObject
}

// An object that performs the initial, nominal glyph generation phase in the layout process.
//
// The nominal glyph generation pass essentially generates one glyph per character; the typesetter may later make substitutions in the glyph stream, for example, changing an acute accent glyph followed by an “e” glyph into a single acute-accented “é” glyph. communicates via the protocol. An example of a class that conforms to the protocol is .


// An object that performs the initial, nominal glyph generation phase in the layout process.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getGlyphGeneratorClass().New()
}




