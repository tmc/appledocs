// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GlyphInfo] class.
var (
	GlyphInfoClass     _GlyphInfoClass
	GlyphInfoClassOnce sync.Once
)

func getGlyphInfoClass() _GlyphInfoClass {
	GlyphInfoClassOnce.Do(func() {
		GlyphInfoClass = _GlyphInfoClass{objc.GetClass("NSGlyphInfo")}
	})
	return GlyphInfoClass
}

type _GlyphInfoClass struct {
	class objc.Class
}

// An interface definition for the [GlyphInfo] class.
type IGlyphInfo interface {
	objectivec.IObject
	CharacterIdentifier() uint
	BaseString() string
	SetBaseString(value string)
	CharacterCollection() unsafe.Pointer
	SetCharacterCollection(value unsafe.Pointer)
	GlyphID() Glyph
	SetGlyphID(value IGlyph)
	GlyphName() string
	SetGlyphName(value string)
}

// A glyph attribute in an attributed string.
//
// Glyphs are the graphic representations of characters, stored in a font, that the text system draws on a display or printed page. Before text can be laid out, the layout manager (< ) generates a stream of glyphs, using the character and font information specified by the attributed string and contained in the font file. represents a glyph attribute value ( ) in an attributed string ( ) and provides a means to override the standard glyph generation process and substitute a specified glyph over the attribute’s range. Glyph attributes are integer values that the layout manager uses to denote special handling for particular glyphs during rendering. enables you to override a font’s built-in mapping from a Unicode character code to a corresponding glyph ID. Overriding the mapping allows you to specify a variant glyph for a given character if the font contains multiple variations for that character or to specify a glyph that doesn’t have a standard mapping (such as some ligature glyphs).
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

// Alloc allocates a new instance without initialization.
func (gc _GlyphInfoClass) Alloc() GlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GlyphInfoClass) New() GlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GlyphInfo) Init() GlyphInfo {
	rv := objc.Send[GlyphInfo](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GlyphInfo) Autorelease() GlyphInfo {
	rv := objc.Send[GlyphInfo](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGlyphInfo creates a new GlyphInfo instance.
func NewGlyphInfo() GlyphInfo {
	return getGlyphInfoClass().New()
}


// The receiver’s character identifier (CID).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/characterIdentifier
func (g_ GlyphInfo) CharacterIdentifier() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("characterIdentifier"))
	return rv
}

// The string containing the character represented by the glyph.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/basestring
func (g_ GlyphInfo) BaseString() string {
	rv := objc.Send[string](g_.ID, objc.Sel("baseString"))
	return rv
}


// SetBaseString sets the value of the baseString property.
// The string containing the character represented by the glyph.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/basestring
func (g_ GlyphInfo) SetBaseString(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setBaseString:"), objc.String(value))
}

// A value specifying the glyph–to–character identifier mapping of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/charactercollection
func (g_ GlyphInfo) CharacterCollection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("characterCollection"))
	return rv
}


// SetCharacterCollection sets the value of the characterCollection property.
// A value specifying the glyph–to–character identifier mapping of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/charactercollection
func (g_ GlyphInfo) SetCharacterCollection(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCharacterCollection:"), value)
}

// The glyph identifier, specified as the index into the internal glyph table of the font.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/glyphid
func (g_ GlyphInfo) GlyphID() Glyph {
	rv := objc.Send[Glyph](g_.ID, objc.Sel("glyphID"))
	return rv
}


// SetGlyphID sets the value of the glyphID property.
// The glyph identifier, specified as the index into the internal glyph table of the font.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/glyphid
func (g_ GlyphInfo) SetGlyphID(value IGlyph) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGlyphID:"), value)
}

// The receiver’s glyph name.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/glyphname
func (g_ GlyphInfo) GlyphName() string {
	rv := objc.Send[string](g_.ID, objc.Sel("glyphName"))
	return rv
}


// SetGlyphName sets the value of the glyphName property.
// The receiver’s glyph name.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/glyphname
func (g_ GlyphInfo) SetGlyphName(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGlyphName:"), objc.String(value))
}



