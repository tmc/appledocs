// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextLineFragment] class.
var (
	TextLineFragmentClass     _TextLineFragmentClass
	TextLineFragmentClassOnce sync.Once
)

func getTextLineFragmentClass() _TextLineFragmentClass {
	TextLineFragmentClassOnce.Do(func() {
		TextLineFragmentClass = _TextLineFragmentClass{objc.GetClass("NSTextLineFragment")}
	})
	return TextLineFragmentClass
}

type _TextLineFragmentClass struct {
	class objc.Class
}

// An interface definition for the [TextLineFragment] class.
type ITextLineFragment interface {
	objectivec.IObject
}

// A class that represents a line fragment as a single textual layout and rendering unit inside a text layout fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLineFragment
type TextLineFragment struct {
	objectivec.Object
}

// TextLineFragmentFrom constructs a [TextLineFragment] from an unsafe.Pointer.
//
// A class that represents a line fragment as a single textual layout and rendering unit inside a text layout fragment.
func TextLineFragmentFrom(ptr unsafe.Pointer) TextLineFragment {
	return TextLineFragment{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TextLineFragmentClass) Alloc() TextLineFragment {
	rv := objc.Send[TextLineFragment](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextLineFragmentClass) New() TextLineFragment {
	rv := objc.Send[TextLineFragment](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextLineFragment) Init() TextLineFragment {
	rv := objc.Send[TextLineFragment](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextLineFragment) Autorelease() TextLineFragment {
	rv := objc.Send[TextLineFragment](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextLineFragment creates a new TextLineFragment instance.
func NewTextLineFragment() TextLineFragment {
	return getTextLineFragmentClass().New()
}


// The source attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlinefragment/attributedstring
func (t_ TextLineFragment) AttributedString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("attributedString"))
	return rv
}


// SetAttributedString sets the value of the attributedString property.
// The source attributed string.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlinefragment/attributedstring
func (t_ TextLineFragment) SetAttributedString(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttributedString:"), value)
}

// The string range for the source attributed string that corresponds to this line fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlinefragment/characterrange
func (t_ TextLineFragment) CharacterRange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("characterRange"))
	return rv
}


// SetCharacterRange sets the value of the characterRange property.
// The string range for the source attributed string that corresponds to this line fragment.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlinefragment/characterrange
func (t_ TextLineFragment) SetCharacterRange(value foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCharacterRange:"), value)
}

// Rendering origin for the left-most glyph in the line fragment coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlinefragment/glyphorigin
func (t_ TextLineFragment) GlyphOrigin() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](t_.ID, objc.Sel("glyphOrigin"))
	return rv
}


// SetGlyphOrigin sets the value of the glyphOrigin property.
// Rendering origin for the left-most glyph in the line fragment coordinate system.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlinefragment/glyphorigin
func (t_ TextLineFragment) SetGlyphOrigin(value coregraphics.CGPoint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGlyphOrigin:"), value)
}

// The typographic bounds that specifies the dimensions of the line fragment for laying out line fragments to each other.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlinefragment/typographicbounds
func (t_ TextLineFragment) TypographicBounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](t_.ID, objc.Sel("typographicBounds"))
	return rv
}


// SetTypographicBounds sets the value of the typographicBounds property.
// The typographic bounds that specifies the dimensions of the line fragment for laying out line fragments to each other.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextlinefragment/typographicbounds
func (t_ TextLineFragment) SetTypographicBounds(value coregraphics.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTypographicBounds:"), value)
}



