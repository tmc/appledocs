// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ATSTypesetter] class.
var (
	ATSTypesetterClass     _ATSTypesetterClass
	ATSTypesetterClassOnce sync.Once
)

func getATSTypesetterClass() _ATSTypesetterClass {
	ATSTypesetterClassOnce.Do(func() {
		ATSTypesetterClass = _ATSTypesetterClass{objc.GetClass("NSATSTypesetter")}
	})
	return ATSTypesetterClass
}

type _ATSTypesetterClass struct {
	class objc.Class
}

// An interface definition for the [ATSTypesetter] class.
type IATSTypesetter interface {
	ITypesetter
}

// A concrete typesetter object that places glyphs during the text layout process.
//
// An object creates line fragment rectangles, positions glyphs within the line fragments, determines line breaks by word wrapping and hyphenation, and handles tab positioning. This object encapsulates the advanced typesetting capabilities of Core Text. provides line and character spacing accuracy and supports many languages, including bidirectional languages.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter
type ATSTypesetter struct {
	Typesetter
}

// ATSTypesetterFrom constructs a [ATSTypesetter] from an unsafe.Pointer.
//
// A concrete typesetter object that places glyphs during the text layout process.
func ATSTypesetterFrom(ptr unsafe.Pointer) ATSTypesetter {
	return ATSTypesetter{
		Typesetter: TypesetterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _ATSTypesetterClass) Alloc() ATSTypesetter {
	rv := objc.Send[ATSTypesetter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ATSTypesetterClass) New() ATSTypesetter {
	rv := objc.Send[ATSTypesetter](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ATSTypesetter) Init() ATSTypesetter {
	rv := objc.Send[ATSTypesetter](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ATSTypesetter) Autorelease() ATSTypesetter {
	rv := objc.Send[ATSTypesetter](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewATSTypesetter creates a new ATSTypesetter instance.
func NewATSTypesetter() ATSTypesetter {
	return getATSTypesetterClass().New()
}


// The backing store that contains the text on which this typesetter operates.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/attributedstring
func (a_ ATSTypesetter) AttributedString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("attributedString"))
	return rv
}


// SetAttributedString sets the value of the attributedString property.
// The backing store that contains the text on which this typesetter operates.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/attributedstring
func (a_ ATSTypesetter) SetAttributedString(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributedString:"), value)
}

// A Boolean value controlling whether the typesetter performs bidirectional text processing.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/bidiprocessingenabled
func (a_ ATSTypesetter) BidiProcessingEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("bidiProcessingEnabled"))
	return rv
}


// SetBidiProcessingEnabled sets the value of the bidiProcessingEnabled property.
// A Boolean value controlling whether the typesetter performs bidirectional text processing.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/bidiprocessingenabled
func (a_ ATSTypesetter) SetBidiProcessingEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBidiProcessingEnabled:"), value)
}

// The text container for the text being typeset.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/currenttextcontainer
func (a_ ATSTypesetter) CurrentTextContainer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("currentTextContainer"))
	return rv
}


// SetCurrentTextContainer sets the value of the currentTextContainer property.
// The text container for the text being typeset.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/currenttextcontainer
func (a_ ATSTypesetter) SetCurrentTextContainer(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentTextContainer:"), value)
}

// The threshold controlling when hyphenation is attempted.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/hyphenationfactor
func (a_ ATSTypesetter) HyphenationFactor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("hyphenationFactor"))
	return rv
}


// SetHyphenationFactor sets the value of the hyphenationFactor property.
// The threshold controlling when hyphenation is attempted.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/hyphenationfactor
func (a_ ATSTypesetter) SetHyphenationFactor(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHyphenationFactor:"), value)
}

// The layout manager for the text being typeset.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/layoutmanager
func (a_ ATSTypesetter) LayoutManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("layoutManager"))
	return rv
}


// SetLayoutManager sets the value of the layoutManager property.
// The layout manager for the text being typeset.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/layoutmanager
func (a_ ATSTypesetter) SetLayoutManager(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLayoutManager:"), value)
}

// The amount (in points) by which text is inset within line fragment rectangles.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/linefragmentpadding
func (a_ ATSTypesetter) LineFragmentPadding() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("lineFragmentPadding"))
	return rv
}


// SetLineFragmentPadding sets the value of the lineFragmentPadding property.
// The amount (in points) by which text is inset within line fragment rectangles.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/linefragmentpadding
func (a_ ATSTypesetter) SetLineFragmentPadding(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLineFragmentPadding:"), value)
}

// The current glyph range being processed.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/paragraphglyphrange
func (a_ ATSTypesetter) ParagraphGlyphRange() foundation.Range {
	rv := objc.Send[foundation.Range](a_.ID, objc.Sel("paragraphGlyphRange"))
	return rv
}


// SetParagraphGlyphRange sets the value of the paragraphGlyphRange property.
// The current glyph range being processed.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/paragraphglyphrange
func (a_ ATSTypesetter) SetParagraphGlyphRange(value foundation.Range) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setParagraphGlyphRange:"), value)
}

// The current paragraph separator range that contains the current glyph range and extends from one paragraph separator character to the next.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/paragraphseparatorglyphrange
func (a_ ATSTypesetter) ParagraphSeparatorGlyphRange() foundation.Range {
	rv := objc.Send[foundation.Range](a_.ID, objc.Sel("paragraphSeparatorGlyphRange"))
	return rv
}


// SetParagraphSeparatorGlyphRange sets the value of the paragraphSeparatorGlyphRange property.
// The current paragraph separator range that contains the current glyph range and extends from one paragraph separator character to the next.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/paragraphseparatorglyphrange
func (a_ ATSTypesetter) SetParagraphSeparatorGlyphRange(value foundation.Range) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setParagraphSeparatorGlyphRange:"), value)
}

// The current typesetter behavior value.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/typesetterbehavior
func (a_ ATSTypesetter) TypesetterBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("typesetterBehavior"))
	return rv
}


// SetTypesetterBehavior sets the value of the typesetterBehavior property.
// The current typesetter behavior value.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/typesetterbehavior
func (a_ ATSTypesetter) SetTypesetterBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTypesetterBehavior:"), value)
}

// A Boolean value controlling whether the typesetter uses the leading (or line gap) value specified in the font metric information.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/usesfontleading
func (a_ ATSTypesetter) UsesFontLeading() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("usesFontLeading"))
	return rv
}


// SetUsesFontLeading sets the value of the usesFontLeading property.
// A Boolean value controlling whether the typesetter uses the leading (or line gap) value specified in the font metric information.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/usesfontleading
func (a_ ATSTypesetter) SetUsesFontLeading(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUsesFontLeading:"), value)
}



