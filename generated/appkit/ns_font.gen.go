// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Font] class.
var (
	FontClass     _FontClass
	FontClassOnce sync.Once
)

func getFontClass() _FontClass {
	FontClassOnce.Do(func() {
		FontClass = _FontClass{objc.GetClass("NSFont")}
	})
	return FontClass
}

type _FontClass struct {
	class objc.Class
}

// An interface definition for the [Font] class.
type IFont interface {
	objectivec.IObject
	// properties:
	Ascender() float64
	BoundingRectForFont() objc.IObject /* cross-framework: Rect */
	CapHeight() float64
	CoveredCharacterSet() foundation.CharacterSet
	Descender() float64
	DisplayName() objc.IObject /* cross-framework: NSString */
	FamilyName() objc.IObject /* cross-framework: NSString */
	FontDescriptor() IFontDescriptor
	FontName() objc.IObject /* cross-framework: NSString */
	FixedPitch() bool
	Vertical() bool
	ItalicAngle() float64
	Leading() float64
	Matrix() corefoundation.CGFloat
	MaximumAdvancement() objc.IObject /* cross-framework: Size */
	MostCompatibleStringEncoding() StringEncoding /* not a class type */
	NumberOfGlyphs() uint
	PointSize() float64
	PrinterFont() IFont
	RenderingMode() FontRenderingMode
	ScreenFont() IFont
	TextTransform() objc.IObject /* cross-framework: AffineTransform */
	UnderlinePosition() float64
	UnderlineThickness() float64
	VerticalFont() IFont
	XHeight() float64
	NSControlGlyph() int
	SetNSControlGlyph(value int)
	IsFixedPitch() bool
	SetIsFixedPitch(value bool)
	IsVertical() bool
	SetIsVertical(value bool)
	Printer() IFont
	SetPrinter(value IFont)
	Screen() IFont
	SetScreen(value IFont)
	NSNullGlyph() int
	SetNSNullGlyph(value int)
	// methods:
	AdvancementForCGGlyph(glyph objc.IObject /* cross-framework: Glyph */) objc.IObject /* cross-framework: Size */
	AdvancementForGlyph(glyph objc.IObject /* cross-framework: Glyph */) objc.IObject /* cross-framework: Size */
	BoundingRectForCGGlyph(glyph objc.IObject /* cross-framework: Glyph */) objc.IObject /* cross-framework: Rect */
	BoundingRectForGlyph(glyph objc.IObject /* cross-framework: Glyph */) objc.IObject /* cross-framework: Rect */
	GetAdvancementsForCGGlyphsCount(advancements SizeArray /* not a class type */, glyphs objc.IObject /* cross-framework: Glyph */, glyphCount uint)
	GetAdvancementsForGlyphsCount(advancements SizeArray /* not a class type */, glyphs objc.IObject /* cross-framework: Glyph */, glyphCount uint)
	GetAdvancementsForPackedGlyphsLength(advancements SizeArray /* not a class type */, packedGlyphs unsafe.Pointer, length uint)
	GetBoundingRectsForCGGlyphsCount(bounds RectArray /* not a class type */, glyphs objc.IObject /* cross-framework: Glyph */, glyphCount uint)
	GetBoundingRectsForGlyphsCount(bounds RectArray /* not a class type */, glyphs objc.IObject /* cross-framework: Glyph */, glyphCount uint)
	GlyphWithName(name objc.IObject /* cross-framework: NSString */) objc.IObject /* cross-framework: Glyph */
	ScreenFontWithRenderingMode(renderingMode FontRenderingMode) IFont
	Set()
	SetInContext(graphicsContext IGraphicsContext)
	FontWithSize(fontSize float64) IFont
}

// The representation of a font in an app.
//
// objects represent fonts to an app, providing access to characteristics of the font and assistance in laying out glyphs relative to one another. Font objects are also used to establish the current font for drawing text directly into a graphics context, using the method. You don’t create objects using the and methods. Instead, you use either or to look up an available font and alter its size or matrix to your needs. These methods check for an existing font object with the specified characteristics, returning it if there is one. Otherwise, they look up the font data requested and create the appropriate object. also defines a number of methods for getting standard system fonts, such as , , and . To request the default size for these standard fonts, pass a negative number or as the font size. See for more information about system fonts.


// The representation of a font in an app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont
type Font struct {
	objectivec.Object
}

// FontFrom constructs a [Font] from an unsafe.Pointer.
//
// The representation of a font in an app.
func FontFrom(ptr unsafe.Pointer) Font {
	return Font{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FontClass) Alloc() Font {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FontClass) New() Font {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ Font) Init() Font {
	rv := objc.Send[Font](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ Font) Autorelease() Font {
	rv := objc.Send[Font](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFont creates a new Font instance.
func NewFont() Font {
	return getFontClass().New()
}



// Returns a font object for the specified font descriptor and font size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/init(descriptor:size:)
func NewFontWithDescriptorSize(fontDescriptor IFontDescriptor, fontSize float64) Font {
	rv := objc.Send[Font](objc.ID(getFontClass().class), objc.Sel("fontWithDescriptor:size:"), fontDescriptor, fontSize)
	return rv
}


// Returns a font object for the specified font descriptor and text transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/init(descriptor:textTransform:)
func NewFontWithDescriptorTextTransform(fontDescriptor IFontDescriptor, textTransform objc.IObject /* cross-framework: AffineTransform */) Font {
	rv := objc.Send[Font](objc.ID(getFontClass().class), objc.Sel("fontWithDescriptor:textTransform:"), fontDescriptor, textTransform)
	return rv
}


// Returns a font object for the specified font name and matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/init(name:matrix:)
func NewFontWithNameMatrix(fontName objc.IObject /* cross-framework: NSString */, fontMatrix corefoundation.CGFloat) Font {
	rv := objc.Send[Font](objc.ID(getFontClass().class), objc.Sel("fontWithName:matrix:"), fontName, fontMatrix)
	return rv
}


// Creates a font object for the specified font name and font size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/init(name:size:)
func NewFontWithNameSize(fontName objc.IObject /* cross-framework: NSString */, fontSize float64) Font {
	rv := objc.Send[Font](objc.ID(getFontClass().class), objc.Sel("fontWithName:size:"), fontName, fontSize)
	return rv
}



// Returns the standard system font in boldface type with the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/boldSystemFont(ofSize:)
func (fc _FontClass) BoldSystemFontOfSize(fontSize float64) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("boldSystemFontOfSize:"), fontSize)
	return rv
}


// Returns the font used for the content of controls in the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/controlContentFont(ofSize:)
func (fc _FontClass) ControlContentFontOfSize(fontSize float64) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("controlContentFontOfSize:"), fontSize)
	return rv
}


// Returns a font object for the specified font descriptor and font size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/init(descriptor:size:)
func (fc _FontClass) FontWithDescriptorSize(fontDescriptor IFontDescriptor, fontSize float64) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("fontWithDescriptor:size:"), fontDescriptor, fontSize)
	return rv
}


// Returns a font object for the specified font descriptor and text transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/init(descriptor:textTransform:)
func (fc _FontClass) FontWithDescriptorTextTransform(fontDescriptor IFontDescriptor, textTransform objc.IObject /* cross-framework: AffineTransform */) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("fontWithDescriptor:textTransform:"), fontDescriptor, textTransform)
	return rv
}


// Returns a font object for the specified font name and matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/init(name:matrix:)
func (fc _FontClass) FontWithNameMatrix(fontName objc.IObject /* cross-framework: NSString */, fontMatrix corefoundation.CGFloat) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("fontWithName:matrix:"), fontName, fontMatrix)
	return rv
}


// Creates a font object for the specified font name and font size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/init(name:size:)
func (fc _FontClass) FontWithNameSize(fontName objc.IObject /* cross-framework: NSString */, fontSize float64) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("fontWithName:size:"), fontName, fontSize)
	return rv
}


// Returns the font used for standard interface labels in the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/labelFont(ofSize:)
func (fc _FontClass) LabelFontOfSize(fontSize float64) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("labelFontOfSize:"), fontSize)
	return rv
}


// Returns the font used for menu bar items, in the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/menuBarFont(ofSize:)
func (fc _FontClass) MenuBarFontOfSize(fontSize float64) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("menuBarFontOfSize:"), fontSize)
	return rv
}


// Returns the font used for menu items, in the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/menuFont(ofSize:)
func (fc _FontClass) MenuFontOfSize(fontSize float64) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("menuFontOfSize:"), fontSize)
	return rv
}


// Returns the font used for standard interface items, such as button labels, menu items, and so on, in the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/messageFont(ofSize:)
func (fc _FontClass) MessageFontOfSize(fontSize float64) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("messageFontOfSize:"), fontSize)
	return rv
}


// Returns a version of the standard system font that contains monospaced digit glyphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/monospacedDigitSystemFont(ofSize:weight:)
func (fc _FontClass) MonospacedDigitSystemFontOfSizeWeight(fontSize float64, weight objc.IObject /* cross-framework: FontWeight */) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("monospacedDigitSystemFontOfSize:weight:"), fontSize, weight)
	return rv
}


// Returns a monospace version of the system font with the specified size and weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/monospacedSystemFont(ofSize:weight:)
func (fc _FontClass) MonospacedSystemFontOfSizeWeight(fontSize float64, weight objc.IObject /* cross-framework: FontWeight */) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("monospacedSystemFontOfSize:weight:"), fontSize, weight)
	return rv
}


// Returns the font used for palette window title bars, in the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/paletteFont(ofSize:)
func (fc _FontClass) PaletteFontOfSize(fontSize float64) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("paletteFontOfSize:"), fontSize)
	return rv
}


// Returns the font associated with the text style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/preferredFont(forTextStyle:options:)
func (fc _FontClass) PreferredFontForTextStyleOptions(style objc.IObject /* cross-framework: FontTextStyle */, options foundation.IDictionary) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("preferredFontForTextStyle:options:"), style, options)
	return rv
}


// Sets the font used by default for documents and other text under the user’s control to the specified font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/setUser(_:)
func (fc _FontClass) SetUserFont(font IFont) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("setUserFont:"), font)
}


// Sets the font used by default for documents and other text under the user’s control, when that font should be fixed-pitch, to the specified font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/setUserFixedPitch(_:)
func (fc _FontClass) SetUserFixedPitchFont(font IFont) {
	objc.Send[objc.ID](objc.ID(fc.class), objc.Sel("setUserFixedPitchFont:"), font)
}


// Returns the standard system font with the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/systemFont(ofSize:)
func (fc _FontClass) SystemFontOfSize(fontSize float64) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("systemFontOfSize:"), fontSize)
	return rv
}


// Returns the standard system font with the specified size and weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/systemFont(ofSize:weight:)
func (fc _FontClass) SystemFontOfSizeWeight(fontSize float64, weight objc.IObject /* cross-framework: FontWeight */) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("systemFontOfSize:weight:"), fontSize, weight)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/systemFont(ofSize:weight:width:)
func (fc _FontClass) SystemFontOfSizeWeightWidth(fontSize float64, weight objc.IObject /* cross-framework: FontWeight */, width objc.IObject /* cross-framework: FontWidth */) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("systemFontOfSize:weight:width:"), fontSize, weight, width)
	return rv
}


// Returns the font size used for the specified control size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/systemFontSize(for:)
func (fc _FontClass) SystemFontSizeForControlSize(controlSize ControlSize) float64 {
	rv := objc.Send[float64](objc.ID(fc.class), objc.Sel("systemFontSizeForControlSize:"), controlSize)
	return rv
}


// Returns the font used for window title bars, in the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/titleBarFont(ofSize:)
func (fc _FontClass) TitleBarFontOfSize(fontSize float64) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("titleBarFontOfSize:"), fontSize)
	return rv
}


// Returns the font used for tool tips labels, in the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/toolTipsFont(ofSize:)
func (fc _FontClass) ToolTipsFontOfSize(fontSize float64) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("toolTipsFontOfSize:"), fontSize)
	return rv
}


// Returns the font used by default for documents and other text under the user’s control (that is, text whose font the user can normally change), when that font should be fixed-pitch, in the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/userFixedPitchFont(ofSize:)
func (fc _FontClass) UserFixedPitchFontOfSize(fontSize float64) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("userFixedPitchFontOfSize:"), fontSize)
	return rv
}


// Returns the font used by default for documents and other text under the user’s control (that is, text whose font the user can normally change), in the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/userFont(ofSize:)
func (fc _FontClass) UserFontOfSize(fontSize float64) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("userFontOfSize:"), fontSize)
	return rv
}


// Returns the size of the standard label font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/labelFontSize
func (fc _FontClass) LabelFontSize() float64 {
	rv := objc.Send[float64](objc.ID(fc.class), objc.Sel("labelFontSize"))
	return rv
}

// Returns the size of the standard small system font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/smallSystemFontSize
func (fc _FontClass) SmallSystemFontSize() float64 {
	rv := objc.Send[float64](objc.ID(fc.class), objc.Sel("smallSystemFontSize"))
	return rv
}

// Returns the size of the standard system font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/systemFontSize
func (fc _FontClass) SystemFontSize() float64 {
	rv := objc.Send[float64](objc.ID(fc.class), objc.Sel("systemFontSize"))
	return rv
}

// Returns the nominal spacing for the given glyph—the distance the current point moves after showing the glyph—accounting for the receiver’s size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/advancement(forCGGlyph:)
func (f_ Font) AdvancementForCGGlyph(glyph objc.IObject /* cross-framework: Glyph */) objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](f_.ID, objc.Sel("advancementForCGGlyph:"), glyph)
	return rv
}


// Returns the nominal spacing for the given glyph—the distance the current point moves after showing the glyph—accounting for the receiver’s size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/advancement(forGlyph:)
func (f_ Font) AdvancementForGlyph(glyph objc.IObject /* cross-framework: Glyph */) objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](f_.ID, objc.Sel("advancementForGlyph:"), glyph)
	return rv
}


// Returns the bounding rectangle for the specified glyph, scaled to the receiver’s size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/boundingRect(forCGGlyph:)
func (f_ Font) BoundingRectForCGGlyph(glyph objc.IObject /* cross-framework: Glyph */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](f_.ID, objc.Sel("boundingRectForCGGlyph:"), glyph)
	return rv
}


// Returns the bounding rectangle for the specified glyph, scaled to the receiver’s size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/boundingRect(forGlyph:)
func (f_ Font) BoundingRectForGlyph(glyph objc.IObject /* cross-framework: Glyph */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](f_.ID, objc.Sel("boundingRectForGlyph:"), glyph)
	return rv
}


// Returns an array of the advancements for the specified glyphs rendered by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/getAdvancements(_:forCGGlyphs:count:)
func (f_ Font) GetAdvancementsForCGGlyphsCount(advancements SizeArray /* not a class type */, glyphs objc.IObject /* cross-framework: Glyph */, glyphCount uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("getAdvancements:forCGGlyphs:count:"), advancements, glyphs, glyphCount)
}


// Returns an array of the advancements for the specified glyphs rendered by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/getAdvancements(_:forGlyphs:count:)
func (f_ Font) GetAdvancementsForGlyphsCount(advancements SizeArray /* not a class type */, glyphs objc.IObject /* cross-framework: Glyph */, glyphCount uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("getAdvancements:forGlyphs:count:"), advancements, glyphs, glyphCount)
}


// Returns an array of the advancements for the specified packed glyphs and rendered by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/getAdvancements(_:forPackedGlyphs:length:)
func (f_ Font) GetAdvancementsForPackedGlyphsLength(advancements SizeArray /* not a class type */, packedGlyphs unsafe.Pointer, length uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("getAdvancements:forPackedGlyphs:length:"), advancements, packedGlyphs, length)
}


// Returns an array of the bounding rectangles for the specified glyphs rendered by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/getBoundingRects(_:forCGGlyphs:count:)
func (f_ Font) GetBoundingRectsForCGGlyphsCount(bounds RectArray /* not a class type */, glyphs objc.IObject /* cross-framework: Glyph */, glyphCount uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("getBoundingRects:forCGGlyphs:count:"), bounds, glyphs, glyphCount)
}


// Returns an array of the bounding rectangles for the specified glyphs rendered by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/getBoundingRects(_:forGlyphs:count:)
func (f_ Font) GetBoundingRectsForGlyphsCount(bounds RectArray /* not a class type */, glyphs objc.IObject /* cross-framework: Glyph */, glyphCount uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("getBoundingRects:forGlyphs:count:"), bounds, glyphs, glyphCount)
}


// Returns the named encoded glyph, or –1 if the receiver contains no such glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/glyph(withName:)
func (f_ Font) GlyphWithName(name objc.IObject /* cross-framework: NSString */) objc.IObject /* cross-framework: Glyph */ {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("glyphWithName:"), name)
	return rv
}


// Returns a bitmapped screen font, when sent to a font object representing a scalable PostScript font, with the specified rendering mode, matching the receiver in typeface and matrix (or size), or if such a font can’t be found.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/screenFont(with:)
func (f_ Font) ScreenFontWithRenderingMode(renderingMode FontRenderingMode) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("screenFontWithRenderingMode:"), renderingMode)
	return rv
}


// Sets this font as the font for the current graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/set()
func (f_ Font) Set() {
	objc.Send[objc.ID](f_.ID, objc.Sel("set"))
}


// Sets this font as the font for the specified graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/set(in:)
func (f_ Font) SetInContext(graphicsContext IGraphicsContext) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInContext:"), graphicsContext)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/withSize(_:)
func (f_ Font) FontWithSize(fontSize float64) IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("fontWithSize:"), fontSize)
	return rv
}


// The top y-coordinate, offset from the baseline, of the font’s longest ascender.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/ascender
func (f_ Font) Ascender() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("ascender"))
	return rv
}


// The font’s bounding rectangle, scaled to the font’s size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/boundingRectForFont
func (f_ Font) BoundingRectForFont() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](f_.ID, objc.Sel("boundingRectForFont"))
	return rv
}


// The cap height of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/capHeight
func (f_ Font) CapHeight() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("capHeight"))
	return rv
}


// The character set containing all of the nominal characters that the font can render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/coveredCharacterSet
func (f_ Font) CoveredCharacterSet() foundation.CharacterSet {
	rv := objc.Send[foundation.CharacterSet](f_.ID, objc.Sel("coveredCharacterSet"))
	return rv
}


// The bottom y-coordinate, offset from the baseline, of the font’s longest descender.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/descender
func (f_ Font) Descender() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("descender"))
	return rv
}


// The name of the font, including family and face names, to use when displaying the font information to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/displayName
func (f_ Font) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("displayName"))
	return rv
}


// The family name of the font—for example, “Times” or “Helvetica.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/familyName
func (f_ Font) FamilyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("familyName"))
	return rv
}


// The font descriptor object for the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/fontDescriptor
func (f_ Font) FontDescriptor() IFontDescriptor {
	rv := objc.Send[FontDescriptor](f_.ID, objc.Sel("fontDescriptor"))
	return rv
}


// The full name of the font, as used in PostScript language code—for example, “Times-Roman” or “Helvetica-Oblique.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/fontName
func (f_ Font) FontName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("fontName"))
	return rv
}


// A Boolean value indicating whether all glyphs in the font have the same advancement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/isFixedPitch
func (f_ Font) FixedPitch() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("fixedPitch"))
	return rv
}


// A Boolean value indicating whether the font is a vertical font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/isVertical
func (f_ Font) Vertical() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("vertical"))
	return rv
}


// The number of degrees that the font is slanted counterclockwise from the vertical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/italicAngle
func (f_ Font) ItalicAngle() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("italicAngle"))
	return rv
}


// Returns the size of the standard label font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/labelFontSize
func (f_ Font) LabelFontSize() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("labelFontSize"))
	return rv
}


// The leading value of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/leading
func (f_ Font) Leading() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("leading"))
	return rv
}


// The transformation matrix associated with the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/matrix
func (f_ Font) Matrix() corefoundation.CGFloat {
	rv := objc.Send[corefoundation.CGFloat](f_.ID, objc.Sel("matrix"))
	return rv
}


// The maximum advance of any of the font’s glyphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/maximumAdvancement
func (f_ Font) MaximumAdvancement() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](f_.ID, objc.Sel("maximumAdvancement"))
	return rv
}


// The string encoding that works best with the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/mostCompatibleStringEncoding
func (f_ Font) MostCompatibleStringEncoding() StringEncoding /* not a class type */ {
	rv := objc.Send[StringEncoding](f_.ID, objc.Sel("mostCompatibleStringEncoding"))
	return rv
}


// The number of glyphs in the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/numberOfGlyphs
func (f_ Font) NumberOfGlyphs() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("numberOfGlyphs"))
	return rv
}


// The point size of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/pointSize
func (f_ Font) PointSize() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("pointSize"))
	return rv
}


// The scalable PostScript font corresponding to current font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/printer
func (f_ Font) PrinterFont() IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("printerFont"))
	return rv
}


// The rendering mode of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/renderingMode
func (f_ Font) RenderingMode() FontRenderingMode {
	rv := objc.Send[FontRenderingMode](f_.ID, objc.Sel("renderingMode"))
	return rv
}


// The bitmapped screen font for the current font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/screen
func (f_ Font) ScreenFont() IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("screenFont"))
	return rv
}


// Returns the size of the standard small system font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/smallSystemFontSize
func (f_ Font) SmallSystemFontSize() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("smallSystemFontSize"))
	return rv
}


// Returns the size of the standard system font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/systemFontSize
func (f_ Font) SystemFontSize() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("systemFontSize"))
	return rv
}


// The current transformation matrix of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/textTransform
func (f_ Font) TextTransform() objc.IObject /* cross-framework: AffineTransform */ {
	rv := objc.Send[corefoundation.AffineTransform](f_.ID, objc.Sel("textTransform"))
	return rv
}


// The baseline offset to use when drawing underlines with the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/underlinePosition
func (f_ Font) UnderlinePosition() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("underlinePosition"))
	return rv
}


// The thickness to use when drawing underlines with the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/underlineThickness
func (f_ Font) UnderlineThickness() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("underlineThickness"))
	return rv
}


// A vertical version of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/vertical-6ym79
func (f_ Font) VerticalFont() IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("verticalFont"))
	return rv
}


// The x-height of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/xHeight
func (f_ Font) XHeight() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("xHeight"))
	return rv
}


// The reserved code for a control glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrolglyph
func (f_ Font) NSControlGlyph() int {
	rv := objc.Send[int](f_.ID, objc.Sel("NSControlGlyph"))
	return rv
}


// The reserved code for a control glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrolglyph
func (f_ Font) SetNSControlGlyph(value int) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setNSControlGlyph:"), value)
}


// A Boolean value indicating whether all glyphs in the font have the same advancement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/isfixedpitch
func (f_ Font) IsFixedPitch() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isFixedPitch"))
	return rv
}


// A Boolean value indicating whether all glyphs in the font have the same advancement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/isfixedpitch
func (f_ Font) SetIsFixedPitch(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsFixedPitch:"), value)
}


// A Boolean value indicating whether the font is a vertical font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/isvertical
func (f_ Font) IsVertical() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isVertical"))
	return rv
}


// A Boolean value indicating whether the font is a vertical font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/isvertical
func (f_ Font) SetIsVertical(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsVertical:"), value)
}


// The scalable PostScript font corresponding to current font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/printer
func (f_ Font) Printer() IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("printer"))
	return rv
}


// The scalable PostScript font corresponding to current font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/printer
func (f_ Font) SetPrinter(value IFont) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPrinter:"), value)
}


// The bitmapped screen font for the current font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/screen
func (f_ Font) Screen() IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("screen"))
	return rv
}


// The bitmapped screen font for the current font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/screen
func (f_ Font) SetScreen(value IFont) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setScreen:"), value)
}


// The reserved code for a null glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnullglyph
func (f_ Font) NSNullGlyph() int {
	rv := objc.Send[int](f_.ID, objc.Sel("NSNullGlyph"))
	return rv
}


// The reserved code for a null glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnullglyph
func (f_ Font) SetNSNullGlyph(value int) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setNSNullGlyph:"), value)
}


