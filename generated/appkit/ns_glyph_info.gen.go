// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSGlyphInfo */


/* debug [class_header]: Header for NSGlyphInfo */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GlyphInfo */
// An interface definition for the [GlyphInfo] class.
type IGlyphInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GlyphInfo */
	// properties:
	BaseString() objc.IObject /* cross-framework: NSString */
	CharacterCollection() CharacterCollection
	CharacterIdentifier() uint
	GlyphID() Glyph /* typedef */
	GlyphName() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GlyphInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GlyphInfo */
// Alloc allocates a new instance without initialization.
func (gc _GlyphInfoClass) Alloc() GlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GlyphInfo */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GlyphInfo */

// Creates a glyph info object from the specified glyph identifier and font informaton.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(cgGlyph:for:baseString:)
func NewGlyphInfoWithCGGlyphForFontBaseString(glyph Glyph /* typedef */, font IFont, string_ objc.IObject /* cross-framework: NSString */) GlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(getGlyphInfoClass().class), objc.Sel("glyphInfoWithCGGlyph:forFont:baseString:"), glyph, font, string_)
	return rv
}/* debug [class_init_methods/constructor]: NewGlyphInfoWithCGGlyphForFontBaseString */


// Instantiates and returns an object using a character identifier and a character collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(characterIdentifier:collection:baseString:)
func NewGlyphInfoWithCharacterIdentifierCollectionBaseString(cid uint, characterCollection CharacterCollection, string_ objc.IObject /* cross-framework: NSString */) GlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(getGlyphInfoClass().class), objc.Sel("glyphInfoWithCharacterIdentifier:collection:baseString:"), cid, characterCollection, string_)
	return rv
}/* debug [class_init_methods/constructor]: NewGlyphInfoWithCharacterIdentifierCollectionBaseString */


// Instantiates and returns a glyph information object using a glyph index and a specified font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(glyph:forFont:baseString:)
func NewGlyphInfoWithGlyphForFontBaseString(glyph Glyph /* typedef */, font IFont, string_ objc.IObject /* cross-framework: NSString */) GlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(getGlyphInfoClass().class), objc.Sel("glyphInfoWithGlyph:forFont:baseString:"), glyph, font, string_)
	return rv
}/* debug [class_init_methods/constructor]: NewGlyphInfoWithGlyphForFontBaseString */


// Instantiates and returns a glyph information object using a glyph name and a specified font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(glyphName:forFont:baseString:)
func NewGlyphInfoWithGlyphNameForFontBaseString(glyphName objc.IObject /* cross-framework: NSString */, font IFont, string_ objc.IObject /* cross-framework: NSString */) GlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(getGlyphInfoClass().class), objc.Sel("glyphInfoWithGlyphName:forFont:baseString:"), glyphName, font, string_)
	return rv
}/* debug [class_init_methods/constructor]: NewGlyphInfoWithGlyphNameForFontBaseString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GlyphInfo */

// Creates a glyph info object from the specified glyph identifier and font informaton.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(cgGlyph:for:baseString:)
func (gc _GlyphInfoClass) GlyphInfoWithCGGlyphForFontBaseString(glyph Glyph /* typedef */, font IFont, string_ objc.IObject /* cross-framework: NSString */) IGlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(gc.class), objc.Sel("glyphInfoWithCGGlyph:forFont:baseString:"), glyph, font, string_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GlyphInfoWithCGGlyphForFontBaseString) */


// Instantiates and returns an object using a character identifier and a character collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(characterIdentifier:collection:baseString:)
func (gc _GlyphInfoClass) GlyphInfoWithCharacterIdentifierCollectionBaseString(cid uint, characterCollection CharacterCollection, string_ objc.IObject /* cross-framework: NSString */) IGlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(gc.class), objc.Sel("glyphInfoWithCharacterIdentifier:collection:baseString:"), cid, characterCollection, string_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GlyphInfoWithCharacterIdentifierCollectionBaseString) */


// Instantiates and returns a glyph information object using a glyph index and a specified font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(glyph:forFont:baseString:)
func (gc _GlyphInfoClass) GlyphInfoWithGlyphForFontBaseString(glyph Glyph /* typedef */, font IFont, string_ objc.IObject /* cross-framework: NSString */) IGlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(gc.class), objc.Sel("glyphInfoWithGlyph:forFont:baseString:"), glyph, font, string_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GlyphInfoWithGlyphForFontBaseString) */


// Instantiates and returns a glyph information object using a glyph name and a specified font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/init(glyphName:forFont:baseString:)
func (gc _GlyphInfoClass) GlyphInfoWithGlyphNameForFontBaseString(glyphName objc.IObject /* cross-framework: NSString */, font IFont, string_ objc.IObject /* cross-framework: NSString */) IGlyphInfo {
	rv := objc.Send[GlyphInfo](objc.ID(gc.class), objc.Sel("glyphInfoWithGlyphName:forFont:baseString:"), glyphName, font, string_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GlyphInfoWithGlyphNameForFontBaseString) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GlyphInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GlyphInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GlyphInfo */

// The string containing the character represented by the glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/baseString
func (g_ GlyphInfo) BaseString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("baseString"))
	return rv
}/* debug [instance_properties/getter]: baseString */


// A value specifying the glyph–to–character identifier mapping of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/characterCollection
func (g_ GlyphInfo) CharacterCollection() CharacterCollection {
	rv := objc.Send[CharacterCollection](g_.ID, objc.Sel("characterCollection"))
	return rv
}/* debug [instance_properties/getter]: characterCollection */


// The receiver’s character identifier (CID).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/characterIdentifier
func (g_ GlyphInfo) CharacterIdentifier() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("characterIdentifier"))
	return rv
}/* debug [instance_properties/getter]: characterIdentifier */


// The glyph identifier, specified as the index into the internal glyph table of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/glyphID
func (g_ GlyphInfo) GlyphID() Glyph /* typedef */ {
	rv := objc.Send[uint32](g_.ID, objc.Sel("glyphID"))
	return rv
}/* debug [instance_properties/getter]: glyphID */


// The receiver’s glyph name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInfo/glyphName
func (g_ GlyphInfo) GlyphName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("glyphName"))
	return rv
}/* debug [instance_properties/getter]: glyphName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSGlyphInfo */


