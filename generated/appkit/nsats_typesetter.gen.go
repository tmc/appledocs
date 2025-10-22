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
	AttributedString() foundation.AttributedString
	SetAttributedString(value foundation.IAttributedString)
	BidiProcessingEnabled() bool
	SetBidiProcessingEnabled(value bool)
	CurrentTextContainer() NSTextContainer
	SetCurrentTextContainer(value ITextContainer)
	HyphenationFactor() float32
	SetHyphenationFactor(value float32)
	LayoutManager() NSLayoutManager
	SetLayoutManager(value ILayoutManager)
	LineFragmentPadding() float64
	SetLineFragmentPadding(value float64)
	ParagraphGlyphRange() foundation.Range
	SetParagraphGlyphRange(value foundation.Range)
	ParagraphSeparatorGlyphRange() foundation.Range
	SetParagraphSeparatorGlyphRange(value foundation.Range)
	TypesetterBehavior() unsafe.Pointer
	SetTypesetterBehavior(value unsafe.Pointer)
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
}

// A concrete typesetter object that places glyphs during the text layout process.
//
// An object creates line fragment rectangles, positions glyphs within the line fragments, determines line breaks by word wrapping and hyphenation, and handles tab positioning. This object encapsulates the advanced typesetting capabilities of Core Text. provides line and character spacing accuracy and supports many languages, including bidirectional languages.


// A concrete typesetter object that places glyphs during the text layout process.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/attributedstring

func (a_ ATSTypesetter) AttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](a_.ID, objc.Sel("attributedString"))
	return rv
}


// The backing store that contains the text on which this typesetter operates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/attributedstring

func (a_ ATSTypesetter) SetAttributedString(value foundation.IAttributedString) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributedString:"), value)
}


// A Boolean value controlling whether the typesetter performs bidirectional text processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/bidiprocessingenabled

func (a_ ATSTypesetter) BidiProcessingEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("bidiProcessingEnabled"))
	return rv
}


// A Boolean value controlling whether the typesetter performs bidirectional text processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/bidiprocessingenabled

func (a_ ATSTypesetter) SetBidiProcessingEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBidiProcessingEnabled:"), value)
}


// The text container for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/currenttextcontainer

func (a_ ATSTypesetter) CurrentTextContainer() NSTextContainer {
	rv := objc.Send[NSTextContainer](a_.ID, objc.Sel("currentTextContainer"))
	return rv
}


// The text container for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/currenttextcontainer

func (a_ ATSTypesetter) SetCurrentTextContainer(value ITextContainer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentTextContainer:"), value)
}


// The threshold controlling when hyphenation is attempted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/hyphenationfactor

func (a_ ATSTypesetter) HyphenationFactor() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("hyphenationFactor"))
	return rv
}


// The threshold controlling when hyphenation is attempted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/hyphenationfactor

func (a_ ATSTypesetter) SetHyphenationFactor(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHyphenationFactor:"), value)
}


// The layout manager for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/layoutmanager

func (a_ ATSTypesetter) LayoutManager() NSLayoutManager {
	rv := objc.Send[NSLayoutManager](a_.ID, objc.Sel("layoutManager"))
	return rv
}


// The layout manager for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/layoutmanager

func (a_ ATSTypesetter) SetLayoutManager(value ILayoutManager) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLayoutManager:"), value)
}


// The amount (in points) by which text is inset within line fragment rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/linefragmentpadding

func (a_ ATSTypesetter) LineFragmentPadding() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("lineFragmentPadding"))
	return rv
}


// The amount (in points) by which text is inset within line fragment rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/linefragmentpadding

func (a_ ATSTypesetter) SetLineFragmentPadding(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLineFragmentPadding:"), value)
}


// The current glyph range being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/paragraphglyphrange

func (a_ ATSTypesetter) ParagraphGlyphRange() foundation.Range {
	rv := objc.Send[foundation.Range](a_.ID, objc.Sel("paragraphGlyphRange"))
	return rv
}


// The current glyph range being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/paragraphglyphrange

func (a_ ATSTypesetter) SetParagraphGlyphRange(value foundation.Range) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setParagraphGlyphRange:"), value)
}


// The current paragraph separator range that contains the current glyph range and extends from one paragraph separator character to the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/paragraphseparatorglyphrange

func (a_ ATSTypesetter) ParagraphSeparatorGlyphRange() foundation.Range {
	rv := objc.Send[foundation.Range](a_.ID, objc.Sel("paragraphSeparatorGlyphRange"))
	return rv
}


// The current paragraph separator range that contains the current glyph range and extends from one paragraph separator character to the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/paragraphseparatorglyphrange

func (a_ ATSTypesetter) SetParagraphSeparatorGlyphRange(value foundation.Range) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setParagraphSeparatorGlyphRange:"), value)
}


// The current typesetter behavior value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/typesetterbehavior

func (a_ ATSTypesetter) TypesetterBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("typesetterBehavior"))
	return rv
}


// The current typesetter behavior value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/typesetterbehavior

func (a_ ATSTypesetter) SetTypesetterBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTypesetterBehavior:"), value)
}


// A Boolean value controlling whether the typesetter uses the leading (or line gap) value specified in the font metric information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/usesfontleading

func (a_ ATSTypesetter) UsesFontLeading() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("usesFontLeading"))
	return rv
}


// A Boolean value controlling whether the typesetter uses the leading (or line gap) value specified in the font metric information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsatstypesetter/usesfontleading

func (a_ ATSTypesetter) SetUsesFontLeading(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUsesFontLeading:"), value)
}



