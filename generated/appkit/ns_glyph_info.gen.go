// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	BaseString() objc.IObject /* cross-framework: NSString */
	SetBaseString(value objc.IObject /* cross-framework: NSString */)
	CharacterCollection() CharacterCollection /* not a class type */
	SetCharacterCollection(value CharacterCollection /* not a class type */)
	CharacterIdentifier() int
	SetCharacterIdentifier(value int)
	GlyphID() objc.IObject /* cross-framework: Glyph */
	SetGlyphID(value objc.IObject /* cross-framework: Glyph */)
	GlyphName() objc.IObject /* cross-framework: NSString */
	SetGlyphName(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A glyph attribute in an attributed string.
//
// Glyphs are the graphic representations of characters, stored in a font, that the text system draws on a display or printed page. Before text can be laid out, the layout manager (< ) generates a stream of glyphs, using the character and font information specified by the attributed string and contained in the font file. represents a glyph attribute value ( ) in an attributed string ( ) and provides a means to override the standard glyph generation process and substitute a specified glyph over the attribute’s range. Glyph attributes are integer values that the layout manager uses to denote special handling for particular glyphs during rendering. enables you to override a font’s built-in mapping from a Unicode character code to a corresponding glyph ID. Overriding the mapping allows you to specify a variant glyph for a given character if the font contains multiple variations for that character or to specify a glyph that doesn’t have a standard mapping (such as some ligature glyphs).


// A glyph attribute in an attributed string.
//
// [Full Topic]
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



// The string containing the character represented by the glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/basestring
func (g_ GlyphInfo) BaseString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("baseString"))
	return rv
}


// The string containing the character represented by the glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/basestring
func (g_ GlyphInfo) SetBaseString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setBaseString:"), value)
}


// A value specifying the glyph–to–character identifier mapping of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/charactercollection
func (g_ GlyphInfo) CharacterCollection() CharacterCollection /* not a class type */ {
	rv := objc.Send[CharacterCollection](g_.ID, objc.Sel("characterCollection"))
	return rv
}


// A value specifying the glyph–to–character identifier mapping of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/charactercollection
func (g_ GlyphInfo) SetCharacterCollection(value CharacterCollection /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCharacterCollection:"), value)
}


// The receiver’s character identifier (CID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/characteridentifier
func (g_ GlyphInfo) CharacterIdentifier() int {
	rv := objc.Send[int](g_.ID, objc.Sel("characterIdentifier"))
	return rv
}


// The receiver’s character identifier (CID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/characteridentifier
func (g_ GlyphInfo) SetCharacterIdentifier(value int) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCharacterIdentifier:"), value)
}


// The glyph identifier, specified as the index into the internal glyph table of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/glyphid
func (g_ GlyphInfo) GlyphID() objc.IObject /* cross-framework: Glyph */ {
	rv := objc.Send[Glyph](g_.ID, objc.Sel("glyphID"))
	return rv
}


// The glyph identifier, specified as the index into the internal glyph table of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/glyphid
func (g_ GlyphInfo) SetGlyphID(value objc.IObject /* cross-framework: Glyph */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGlyphID:"), value)
}


// The receiver’s glyph name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/glyphname
func (g_ GlyphInfo) GlyphName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("glyphName"))
	return rv
}


// The receiver’s glyph name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsglyphinfo/glyphname
func (g_ GlyphInfo) SetGlyphName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGlyphName:"), value)
}



