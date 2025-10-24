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
	BoundingRectForFont() objc.IObject /* cross-framework: Rect */
	CoveredCharacterSet() objc.IObject /* cross-framework: CharacterSet */
	DisplayName() objc.IObject /* cross-framework: NSString */
	FontName() objc.IObject /* cross-framework: NSString */
	Vertical() bool
	MostCompatibleStringEncoding() StringEncoding /* not a class type */
	NumberOfGlyphs() uint
	PointSize() float64
	PrinterFont() IFont
	RenderingMode() FontRenderingMode
	ScreenFont() IFont
	NSControlGlyph() int
	SetNSControlGlyph(value int)
	FamilyName() objc.IObject /* cross-framework: NSString */
	SetFamilyName(value objc.IObject /* cross-framework: NSString */)
	FontDescriptor() IFontDescriptor
	SetFontDescriptor(value IFontDescriptor)
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
	BoundingRectForCGGlyph(glyph objc.IObject /* cross-framework: Glyph */) objc.IObject /* cross-framework: Rect */
	BoundingRectForGlyph(glyph objc.IObject /* cross-framework: Glyph */) objc.IObject /* cross-framework: Rect */
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
func (fc _FontClass) MonospacedDigitSystemFontOfSizeWeight(fontSize float64, weight FontWeight /* not a class type */) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("monospacedDigitSystemFontOfSize:weight:"), fontSize, weight)
	return rv
}


// Returns a monospace version of the system font with the specified size and weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/monospacedSystemFont(ofSize:weight:)
func (fc _FontClass) MonospacedSystemFontOfSizeWeight(fontSize float64, weight FontWeight /* not a class type */) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("monospacedSystemFontOfSize:weight:"), fontSize, weight)
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
func (fc _FontClass) SystemFontOfSizeWeight(fontSize float64, weight FontWeight /* not a class type */) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("systemFontOfSize:weight:"), fontSize, weight)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/systemFont(ofSize:weight:width:)
func (fc _FontClass) SystemFontOfSizeWeightWidth(fontSize float64, weight FontWeight /* not a class type */, width objc.IObject /* cross-framework: FontWidth */) IFont {
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


// Returns the named encoded glyph, or –1 if the receiver contains no such glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/glyph(withName:)
func (f_ Font) GlyphWithName(name objc.IObject /* cross-framework: NSString */) objc.IObject /* cross-framework: Glyph */ {
	rv := objc.Send[Glyph](f_.ID, objc.Sel("glyphWithName:"), name)
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


// The font’s bounding rectangle, scaled to the font’s size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/boundingRectForFont
func (f_ Font) BoundingRectForFont() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](f_.ID, objc.Sel("boundingRectForFont"))
	return rv
}


// The character set containing all of the nominal characters that the font can render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/coveredCharacterSet
func (f_ Font) CoveredCharacterSet() objc.IObject /* cross-framework: CharacterSet */ {
	rv := objc.Send[foundation.CharacterSet](f_.ID, objc.Sel("coveredCharacterSet"))
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


// The full name of the font, as used in PostScript language code—for example, “Times-Roman” or “Helvetica-Oblique.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/fontName
func (f_ Font) FontName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("fontName"))
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


// Returns the size of the standard label font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/labelFontSize
func (f_ Font) LabelFontSize() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("labelFontSize"))
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


// The family name of the font—for example, “Times” or “Helvetica.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/familyname
func (f_ Font) FamilyName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("familyName"))
	return rv
}


// The family name of the font—for example, “Times” or “Helvetica.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/familyname
func (f_ Font) SetFamilyName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFamilyName:"), value)
}


// The font descriptor object for the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/fontdescriptor
func (f_ Font) FontDescriptor() IFontDescriptor {
	rv := objc.Send[FontDescriptor](f_.ID, objc.Sel("fontDescriptor"))
	return rv
}


// The font descriptor object for the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/fontdescriptor
func (f_ Font) SetFontDescriptor(value IFontDescriptor) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFontDescriptor:"), value)
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


