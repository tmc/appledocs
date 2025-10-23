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
	// properties:
	BaseString() string /* primitive/slice/pointer. */
	CharacterCollection() CharacterCollection
	CharacterIdentifier() uint /* primitive/slice/pointer. */
	GlyphID() objc.IObject /* cross-framework: Glyph */
	GlyphName() string /* primitive/slice/pointer. */
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



// Creates a glyph info object from the specified glyph identifier and font informaton.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(cgGlyph:for:baseString:)
func NewGlyphInfoWithCGGlyphForFontBaseString(glyph objc.IObject /* cross-framework Glyph */, font IFont, string_ string /* primitive/slice/pointer. */) GlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(getGlyphInfoClass().class), objc.Sel("glyphInfoWithCGGlyph:forFont:baseString:"), glyph, font, objc.String(string_))
	return rv
}


// Instantiates and returns an object using a character identifier and a character collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(characterIdentifier:collection:baseString:)
func NewGlyphInfoWithCharacterIdentifierCollectionBaseString(cid uint /* primitive/slice/pointer. */, characterCollection CharacterCollection, string_ string /* primitive/slice/pointer. */) GlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(getGlyphInfoClass().class), objc.Sel("glyphInfoWithCharacterIdentifier:collection:baseString:"), cid, characterCollection, objc.String(string_))
	return rv
}


// Instantiates and returns a glyph information object using a glyph index and a specified font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(glyph:forFont:baseString:)
func NewGlyphInfoWithGlyphForFontBaseString(glyph objc.IObject /* cross-framework Glyph */, font IFont, string_ string /* primitive/slice/pointer. */) GlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(getGlyphInfoClass().class), objc.Sel("glyphInfoWithGlyph:forFont:baseString:"), glyph, font, objc.String(string_))
	return rv
}


// Instantiates and returns a glyph information object using a glyph name and a specified font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(glyphName:forFont:baseString:)
func NewGlyphInfoWithGlyphNameForFontBaseString(glyphName string /* primitive/slice/pointer. */, font IFont, string_ string /* primitive/slice/pointer. */) GlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(getGlyphInfoClass().class), objc.Sel("glyphInfoWithGlyphName:forFont:baseString:"), objc.String(glyphName), font, objc.String(string_))
	return rv
}



// Creates a glyph info object from the specified glyph identifier and font informaton.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(cgGlyph:for:baseString:)
func (gc _GlyphInfoClass) GlyphInfoWithCGGlyphForFontBaseString(glyph objc.IObject /* cross-framework Glyph */, font IFont, string_ string /* primitive/slice/pointer. */) IGlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(gc.class), objc.Sel("glyphInfoWithCGGlyph:forFont:baseString:"), glyph, font, objc.String(string_))
	return rv
}


// Instantiates and returns an object using a character identifier and a character collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(characterIdentifier:collection:baseString:)
func (gc _GlyphInfoClass) GlyphInfoWithCharacterIdentifierCollectionBaseString(cid uint /* primitive/slice/pointer. */, characterCollection CharacterCollection, string_ string /* primitive/slice/pointer. */) IGlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(gc.class), objc.Sel("glyphInfoWithCharacterIdentifier:collection:baseString:"), cid, characterCollection, objc.String(string_))
	return rv
}


// Instantiates and returns a glyph information object using a glyph index and a specified font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(glyph:forFont:baseString:)
func (gc _GlyphInfoClass) GlyphInfoWithGlyphForFontBaseString(glyph objc.IObject /* cross-framework Glyph */, font IFont, string_ string /* primitive/slice/pointer. */) IGlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(gc.class), objc.Sel("glyphInfoWithGlyph:forFont:baseString:"), glyph, font, objc.String(string_))
	return rv
}


// Instantiates and returns a glyph information object using a glyph name and a specified font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(glyphName:forFont:baseString:)
func (gc _GlyphInfoClass) GlyphInfoWithGlyphNameForFontBaseString(glyphName string /* primitive/slice/pointer. */, font IFont, string_ string /* primitive/slice/pointer. */) IGlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(gc.class), objc.Sel("glyphInfoWithGlyphName:forFont:baseString:"), objc.String(glyphName), font, objc.String(string_))
	return rv
}


// The string containing the character represented by the glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/baseString
func (g_ GlyphInfo) BaseString() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](g_.ID, objc.Sel("baseString"))
	return rv
}


// A value specifying the glyph–to–character identifier mapping of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/characterCollection
func (g_ GlyphInfo) CharacterCollection() CharacterCollection {
	rv := objc.Send[CharacterCollection](g_.ID, objc.Sel("characterCollection"))
	return rv
}


// The receiver’s character identifier (CID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/characterIdentifier
func (g_ GlyphInfo) CharacterIdentifier() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](g_.ID, objc.Sel("characterIdentifier"))
	return rv
}


// The glyph identifier, specified as the index into the internal glyph table of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/glyphID
func (g_ GlyphInfo) GlyphID() objc.IObject /* cross-framework: Glyph */ {
	rv := objc.Send[Glyph](g_.ID, objc.Sel("glyphID"))
	return rv
}


// The receiver’s glyph name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/glyphName
func (g_ GlyphInfo) GlyphName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](g_.ID, objc.Sel("glyphName"))
	return rv
}


