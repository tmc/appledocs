// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	Set()
}

// The representation of a font in an app.
//
// objects represent fonts to an app, providing access to characteristics of the font and assistance in laying out glyphs relative to one another. Font objects are also used to establish the current font for drawing text directly into a graphics context, using the method. You don’t create objects using the and methods. Instead, you use either or to look up an available font and alter its size or matrix to your needs. These methods check for an existing font object with the specified characteristics, returning it if there is one. Otherwise, they look up the font data requested and create the appropriate object. also defines a number of methods for getting standard system fonts, such as , , and . To request the default size for these standard fonts, pass a negative number or as the font size. See for more information about system fonts.
//
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


// Returns the font used for menu bar items, in the specified size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/menuBarFont(ofSize:)
func (fc _FontClass) MenuBarFontOfSize(fontSize float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("menuBarFontOfSize:"), fontSize)
	return rv
}

// Returns the standard system font with the specified size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/systemFont(ofSize:)
func (fc _FontClass) SystemFontOfSize(fontSize float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("systemFontOfSize:"), fontSize)
	return rv
}

// Sets this font as the font for the current graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/set()
func (f_ Font) Set() {
	objc.Send[objc.ID](f_.ID, objc.Sel("set"))
}

// The character set containing all of the nominal characters that the font can render.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/coveredCharacterSet
func (f_ Font) CoveredCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("coveredCharacterSet"))
	return rv
}

// The scalable PostScript font corresponding to current font.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/printer
func (f_ Font) PrinterFont() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("printerFont"))
	return rv
}



