// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Typesetter] class.
var (
	TypesetterClass     _TypesetterClass
	TypesetterClassOnce sync.Once
)

func getTypesetterClass() _TypesetterClass {
	TypesetterClassOnce.Do(func() {
		TypesetterClass = _TypesetterClass{objc.GetClass("NSTypesetter")}
	})
	return TypesetterClass
}

type _TypesetterClass struct {
	class objc.Class
}

// An interface definition for the [Typesetter] class.
type ITypesetter interface {
	objectivec.IObject
}

// An abstract class that performs various type layout tasks.
//
// uses concrete subclasses of to perform line layout, which includes word wrapping, hyphenation, and line breaking in either vertical or horizontal rectangles. By default, the text system uses the concrete subclass .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter
type Typesetter struct {
	objectivec.Object
}

// TypesetterFrom constructs a [Typesetter] from an unsafe.Pointer.
//
// An abstract class that performs various type layout tasks.
func TypesetterFrom(ptr unsafe.Pointer) Typesetter {
	return Typesetter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TypesetterClass) Alloc() Typesetter {
	rv := objc.Send[Typesetter](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TypesetterClass) New() Typesetter {
	rv := objc.Send[Typesetter](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Typesetter) Init() Typesetter {
	rv := objc.Send[Typesetter](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Typesetter) Autorelease() Typesetter {
	rv := objc.Send[Typesetter](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTypesetter creates a new Typesetter instance.
func NewTypesetter() Typesetter {
	return getTypesetterClass().New()
}


// Returns the text backing store, usually an instance of
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/attributedstring
func (t_ Typesetter) AttributedString() AttributedString {
	rv := objc.Send[AttributedString](t_.ID, objc.Sel("attributedString"))
	return rv
}


// SetAttributedString sets the value of the attributedString property.
// Returns the text backing store, usually an instance of

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/attributedstring
func (t_ Typesetter) SetAttributedString(value IAttributedString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttributedString:"), value)
}

// Returns the attributes used to lay out the extra line fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/attributesforextralinefragment
func (t_ Typesetter) AttributesForExtraLineFragment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("attributesForExtraLineFragment"))
	return rv
}


// SetAttributesForExtraLineFragment sets the value of the attributesForExtraLineFragment property.
// Returns the attributes used to lay out the extra line fragment.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/attributesforextralinefragment
func (t_ Typesetter) SetAttributesForExtraLineFragment(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttributesForExtraLineFragment:"), value)
}

// Returns whether bidirectional text processing is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/bidiprocessingenabled
func (t_ Typesetter) BidiProcessingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("bidiProcessingEnabled"))
	return rv
}


// SetBidiProcessingEnabled sets the value of the bidiProcessingEnabled property.
// Returns whether bidirectional text processing is enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/bidiprocessingenabled
func (t_ Typesetter) SetBidiProcessingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBidiProcessingEnabled:"), value)
}

// Returns the paragraph style object for the text being typeset.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/currentparagraphstyle
func (t_ Typesetter) CurrentParagraphStyle() NSParagraphStyle {
	rv := objc.Send[NSParagraphStyle](t_.ID, objc.Sel("currentParagraphStyle"))
	return rv
}


// SetCurrentParagraphStyle sets the value of the currentParagraphStyle property.
// Returns the paragraph style object for the text being typeset.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/currentparagraphstyle
func (t_ Typesetter) SetCurrentParagraphStyle(value NSParagraphStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentParagraphStyle:"), value)
}

// Returns the text container for the text being typeset.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/currenttextcontainer
func (t_ Typesetter) CurrentTextContainer() NSTextContainer {
	rv := objc.Send[NSTextContainer](t_.ID, objc.Sel("currentTextContainer"))
	return rv
}


// SetCurrentTextContainer sets the value of the currentTextContainer property.
// Returns the text container for the text being typeset.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/currenttextcontainer
func (t_ Typesetter) SetCurrentTextContainer(value ITextContainer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentTextContainer:"), value)
}

// Returns the current hyphenation factor.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/hyphenationfactor
func (t_ Typesetter) HyphenationFactor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("hyphenationFactor"))
	return rv
}


// SetHyphenationFactor sets the value of the hyphenationFactor property.
// Returns the current hyphenation factor.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/hyphenationfactor
func (t_ Typesetter) SetHyphenationFactor(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHyphenationFactor:"), value)
}

// Returns the layout manager for the text being typeset.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/layoutmanager
func (t_ Typesetter) LayoutManager() NSLayoutManager {
	rv := objc.Send[NSLayoutManager](t_.ID, objc.Sel("layoutManager"))
	return rv
}


// SetLayoutManager sets the value of the layoutManager property.
// Returns the layout manager for the text being typeset.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/layoutmanager
func (t_ Typesetter) SetLayoutManager(value ILayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutManager:"), value)
}

// Returns the current line fragment padding, in points.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/linefragmentpadding
func (t_ Typesetter) LineFragmentPadding() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("lineFragmentPadding"))
	return rv
}


// SetLineFragmentPadding sets the value of the lineFragmentPadding property.
// Returns the current line fragment padding, in points.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/linefragmentpadding
func (t_ Typesetter) SetLineFragmentPadding(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineFragmentPadding:"), value)
}

// Returns the character range currently being processed.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphcharacterrange
func (t_ Typesetter) ParagraphCharacterRange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("paragraphCharacterRange"))
	return rv
}


// SetParagraphCharacterRange sets the value of the paragraphCharacterRange property.
// Returns the character range currently being processed.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphcharacterrange
func (t_ Typesetter) SetParagraphCharacterRange(value foundation.IRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParagraphCharacterRange:"), value)
}

// Returns the glyph range currently being processed.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphglyphrange
func (t_ Typesetter) ParagraphGlyphRange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("paragraphGlyphRange"))
	return rv
}


// SetParagraphGlyphRange sets the value of the paragraphGlyphRange property.
// Returns the glyph range currently being processed.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphglyphrange
func (t_ Typesetter) SetParagraphGlyphRange(value foundation.IRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParagraphGlyphRange:"), value)
}

// Returns the current paragraph separator character range.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphseparatorcharacterrange
func (t_ Typesetter) ParagraphSeparatorCharacterRange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("paragraphSeparatorCharacterRange"))
	return rv
}


// SetParagraphSeparatorCharacterRange sets the value of the paragraphSeparatorCharacterRange property.
// Returns the current paragraph separator character range.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphseparatorcharacterrange
func (t_ Typesetter) SetParagraphSeparatorCharacterRange(value foundation.IRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParagraphSeparatorCharacterRange:"), value)
}

// Returns the current paragraph separator range.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphseparatorglyphrange
func (t_ Typesetter) ParagraphSeparatorGlyphRange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("paragraphSeparatorGlyphRange"))
	return rv
}


// SetParagraphSeparatorGlyphRange sets the value of the paragraphSeparatorGlyphRange property.
// Returns the current paragraph separator range.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphseparatorglyphrange
func (t_ Typesetter) SetParagraphSeparatorGlyphRange(value foundation.IRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParagraphSeparatorGlyphRange:"), value)
}

// Returns an array containing the text containers belonging to the current layout manager.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/textcontainers
func (t_ Typesetter) TextContainers() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("textContainers"))
	return rv
}


// SetTextContainers sets the value of the textContainers property.
// Returns an array containing the text containers belonging to the current layout manager.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/textcontainers
func (t_ Typesetter) SetTextContainers(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainers:"), value)
}

// Returns the current typesetter behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/typesetterbehavior
func (t_ Typesetter) TypesetterBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("typesetterBehavior"))
	return rv
}


// SetTypesetterBehavior sets the value of the typesetterBehavior property.
// Returns the current typesetter behavior.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/typesetterbehavior
func (t_ Typesetter) SetTypesetterBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTypesetterBehavior:"), value)
}

// Returns whether the typesetter uses the leading (or line gap) value specified in the font metric information of the current font.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/usesfontleading
func (t_ Typesetter) UsesFontLeading() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontLeading"))
	return rv
}


// SetUsesFontLeading sets the value of the usesFontLeading property.
// Returns whether the typesetter uses the leading (or line gap) value specified in the font metric information of the current font.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/usesfontleading
func (t_ Typesetter) SetUsesFontLeading(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontLeading:"), value)
}



