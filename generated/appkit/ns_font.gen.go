// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	NSControlGlyph() int
	SetNSControlGlyph(value int)
	CoveredCharacterSet() foundation.CharacterSet
	SetCoveredCharacterSet(value foundation.CharacterSet)
	DisplayName() string
	SetDisplayName(value string)
	FamilyName() string
	SetFamilyName(value string)
	FontDescriptor() FontDescriptor
	SetFontDescriptor(value FontDescriptor)
	FontName() string
	SetFontName(value string)
	IsFixedPitch() bool
	SetIsFixedPitch(value bool)
	IsVertical() bool
	SetIsVertical(value bool)
	MostCompatibleStringEncoding() uint
	SetMostCompatibleStringEncoding(value uint)
	NumberOfGlyphs() int
	SetNumberOfGlyphs(value int)
	PointSize() float64
	SetPointSize(value float64)
	Printer() IFont
	SetPrinter(value IFont)
	RenderingMode() unsafe.Pointer
	SetRenderingMode(value unsafe.Pointer)
	Screen() IFont
	SetScreen(value IFont)
	Vertical() IFont
	SetVertical(value IFont)
	NSNullGlyph() int
	SetNSNullGlyph(value int)
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/systemFont(ofSize:weight:width:)
func (fc _FontClass) SystemFontOfSizeWeightWidth(fontSize float64, weight unsafe.Pointer, width FontWidth) IFont {
	rv := objc.Send[Font](objc.ID(fc.class), objc.Sel("systemFontOfSize:weight:width:"), fontSize, weight, width)
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


// The character set containing all of the nominal characters that the font can render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/coveredcharacterset
func (f_ Font) CoveredCharacterSet() foundation.CharacterSet {
	rv := objc.Send[foundation.CharacterSet](f_.ID, objc.Sel("coveredCharacterSet"))
	return rv
}


// The character set containing all of the nominal characters that the font can render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/coveredcharacterset
func (f_ Font) SetCoveredCharacterSet(value foundation.CharacterSet) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setCoveredCharacterSet:"), value)
}


// The name of the font, including family and face names, to use when displaying the font information to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/displayname
func (f_ Font) DisplayName() string {
	rv := objc.Send[string](f_.ID, objc.Sel("displayName"))
	return rv
}


// The name of the font, including family and face names, to use when displaying the font information to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/displayname
func (f_ Font) SetDisplayName(value string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDisplayName:"), objc.String(value))
}


// The family name of the font—for example, “Times” or “Helvetica.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/familyname
func (f_ Font) FamilyName() string {
	rv := objc.Send[string](f_.ID, objc.Sel("familyName"))
	return rv
}


// The family name of the font—for example, “Times” or “Helvetica.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/familyname
func (f_ Font) SetFamilyName(value string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFamilyName:"), objc.String(value))
}


// The font descriptor object for the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/fontdescriptor
func (f_ Font) FontDescriptor() FontDescriptor {
	rv := objc.Send[FontDescriptor](f_.ID, objc.Sel("fontDescriptor"))
	return rv
}


// The font descriptor object for the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/fontdescriptor
func (f_ Font) SetFontDescriptor(value FontDescriptor) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFontDescriptor:"), value)
}


// The full name of the font, as used in PostScript language code—for example, “Times-Roman” or “Helvetica-Oblique.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/fontname
func (f_ Font) FontName() string {
	rv := objc.Send[string](f_.ID, objc.Sel("fontName"))
	return rv
}


// The full name of the font, as used in PostScript language code—for example, “Times-Roman” or “Helvetica-Oblique.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/fontname
func (f_ Font) SetFontName(value string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFontName:"), objc.String(value))
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


// The string encoding that works best with the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/mostcompatiblestringencoding
func (f_ Font) MostCompatibleStringEncoding() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("mostCompatibleStringEncoding"))
	return rv
}


// The string encoding that works best with the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/mostcompatiblestringencoding
func (f_ Font) SetMostCompatibleStringEncoding(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMostCompatibleStringEncoding:"), value)
}


// The number of glyphs in the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/numberofglyphs
func (f_ Font) NumberOfGlyphs() int {
	rv := objc.Send[int](f_.ID, objc.Sel("numberOfGlyphs"))
	return rv
}


// The number of glyphs in the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/numberofglyphs
func (f_ Font) SetNumberOfGlyphs(value int) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setNumberOfGlyphs:"), value)
}


// The point size of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/pointsize
func (f_ Font) PointSize() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("pointSize"))
	return rv
}


// The point size of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/pointsize
func (f_ Font) SetPointSize(value float64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPointSize:"), value)
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


// The rendering mode of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/renderingmode
func (f_ Font) RenderingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("renderingMode"))
	return rv
}


// The rendering mode of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/renderingmode
func (f_ Font) SetRenderingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setRenderingMode:"), value)
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


// A vertical version of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/vertical-6ym79
func (f_ Font) Vertical() IFont {
	rv := objc.Send[Font](f_.ID, objc.Sel("vertical"))
	return rv
}


// A vertical version of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfont/vertical-6ym79
func (f_ Font) SetVertical(value IFont) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setVertical:"), value)
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



