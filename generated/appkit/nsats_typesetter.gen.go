// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	// properties:
	AttributedString() objc.IObject /* cross-framework: AttributedString */
	SetAttributedString(value objc.IObject /* cross-framework: AttributedString */)
	BidiProcessingEnabled() bool /* primitive/slice/pointer. */
	SetBidiProcessingEnabled(value bool /* primitive/slice/pointer. */)
	CurrentTextContainer() ITextContainer
	HyphenationFactor() float32 /* primitive/slice/pointer. */
	SetHyphenationFactor(value float32 /* primitive/slice/pointer. */)
	LayoutManager() objc.IObject /* cross-framework: LayoutManager */
	LineFragmentPadding() float64 /* primitive/slice/pointer. */
	SetLineFragmentPadding(value float64 /* primitive/slice/pointer. */)
	ParagraphGlyphRange() foundation.objc.IObject /* cross-framework: Range */
	ParagraphSeparatorGlyphRange() foundation.objc.IObject /* cross-framework: Range */
	TypesetterBehavior() TypesetterBehavior
	SetTypesetterBehavior(value TypesetterBehavior)
	UsesFontLeading() bool /* primitive/slice/pointer. */
	SetUsesFontLeading(value bool /* primitive/slice/pointer. */)
	// methods:
	BoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(glyphIndex uint /* primitive/slice/pointer. */, textContainer ITextContainer, proposedRect coregraphics.CGRect, glyphPosition coregraphics.CGPoint, charIndex uint /* primitive/slice/pointer. */) coregraphics.CGRect
	GetLineFragmentRectUsedRectForParagraphSeparatorGlyphRangeAtProposedOrigin(lineFragmentRect coregraphics.CGRect, lineFragmentUsedRect coregraphics.CGRect, paragraphSeparatorGlyphRange foundation.objc.IObject /* cross-framework Range */, lineOrigin coregraphics.CGPoint)
	HyphenCharacterForGlyphAtIndex(glyphIndex uint /* primitive/slice/pointer. */) unsafe.Pointer
	HyphenationFactorForGlyphAtIndex(glyphIndex uint /* primitive/slice/pointer. */) float32 /* primitive/slice/pointer. */
	LayoutParagraphAtPoint(lineFragmentOrigin coregraphics.CGPoint) uint /* primitive/slice/pointer. */
	LineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint /* primitive/slice/pointer. */, rect coregraphics.CGRect) float64 /* primitive/slice/pointer. */
	ParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint /* primitive/slice/pointer. */, rect coregraphics.CGRect) float64 /* primitive/slice/pointer. */
	ParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint /* primitive/slice/pointer. */, rect coregraphics.CGRect) float64 /* primitive/slice/pointer. */
	SetHardInvalidationForGlyphRange(flag bool /* primitive/slice/pointer. */, glyphRange foundation.objc.IObject /* cross-framework Range */)
	SetParagraphGlyphRangeSeparatorGlyphRange(paragraphRange foundation.objc.IObject /* cross-framework Range */, paragraphSeparatorRange foundation.objc.IObject /* cross-framework Range */)
	ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(charIndex uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	SubstituteFontForFont(originalFont IFont) IFont
	TextTabForGlyphLocationWritingDirectionMaxLocation(glyphLocation float64 /* primitive/slice/pointer. */, direction WritingDirection, maxLocation float64 /* primitive/slice/pointer. */) ITextTab
	WillSetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(lineRect coregraphics.CGRect, glyphRange foundation.objc.IObject /* cross-framework Range */, usedRect coregraphics.CGRect, baselineOffset coregraphics.float64 /* primitive/slice/pointer. */)
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



// Returns a shared instance of the typesetter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/shared
func (ac _ATSTypesetterClass) SharedTypesetter() ATSTypesetter {
	rv := objc.Send[ATSTypesetter](objc.ID(ac.class), objc.Sel("sharedTypesetter"))
	return rv
}

// Returns the bounding rectangle for a control glyph, at the specified glyph position and character index in the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/boundingBox(forControlGlyphAt:for:proposedLineFragment:glyphPosition:characterIndex:)
func (a_ ATSTypesetter) BoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(glyphIndex uint /* primitive/slice/pointer. */, textContainer ITextContainer, proposedRect coregraphics.CGRect, glyphPosition coregraphics.CGPoint, charIndex uint /* primitive/slice/pointer. */) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](a_.ID, objc.Sel("boundingBoxForControlGlyphAtIndex:forTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), glyphIndex, textContainer, proposedRect, glyphPosition, charIndex)
	return rv
}


// Calculates the line fragment rectangle and the portion of the rectangle that contains marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/getLineFragmentRect(_:usedRect:forParagraphSeparatorGlyphRange:atProposedOrigin:)
func (a_ ATSTypesetter) GetLineFragmentRectUsedRectForParagraphSeparatorGlyphRangeAtProposedOrigin(lineFragmentRect coregraphics.CGRect, lineFragmentUsedRect coregraphics.CGRect, paragraphSeparatorGlyphRange foundation.objc.IObject /* cross-framework Range */, lineOrigin coregraphics.CGPoint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("getLineFragmentRect:usedRect:forParagraphSeparatorGlyphRange:atProposedOrigin:"), lineFragmentRect, lineFragmentUsedRect, paragraphSeparatorGlyphRange, lineOrigin)
}


// Returns the hyphen character to be inserted after the specified glyph when hyphenation is enabled in the layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/hyphenCharacter(forGlyphAt:)
func (a_ ATSTypesetter) HyphenCharacterForGlyphAtIndex(glyphIndex uint /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("hyphenCharacterForGlyphAtIndex:"), glyphIndex)
	return rv
}


// Returns the hyphenation factor in effect at the specified glyph index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/hyphenationFactor(forGlyphAt:)
func (a_ ATSTypesetter) HyphenationFactorForGlyphAtIndex(glyphIndex uint /* primitive/slice/pointer. */) float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](a_.ID, objc.Sel("hyphenationFactorForGlyphAtIndex:"), glyphIndex)
	return rv
}


// Lays out glyphs in the current glyph range until the next paragraph separator is reached.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/layoutParagraph(at:)
func (a_ ATSTypesetter) LayoutParagraphAtPoint(lineFragmentOrigin coregraphics.CGPoint) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](a_.ID, objc.Sel("layoutParagraphAtPoint:"), lineFragmentOrigin)
	return rv
}


// Returns the line spacing in effect following the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/lineSpacing(afterGlyphAt:withProposedLineFragmentRect:)
func (a_ ATSTypesetter) LineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint /* primitive/slice/pointer. */, rect coregraphics.CGRect) float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](a_.ID, objc.Sel("lineSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}


// Returns the number of points of space added following a paragraph, in effect after the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/paragraphSpacing(afterGlyphAt:withProposedLineFragmentRect:)
func (a_ ATSTypesetter) ParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint /* primitive/slice/pointer. */, rect coregraphics.CGRect) float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](a_.ID, objc.Sel("paragraphSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}


// Returns the number of points of space added before a paragraph, which is in effect before the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/paragraphSpacing(beforeGlyphAt:withProposedLineFragmentRect:)
func (a_ ATSTypesetter) ParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint /* primitive/slice/pointer. */, rect coregraphics.CGRect) float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](a_.ID, objc.Sel("paragraphSpacingBeforeGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}


// Sets a Boolean value that determines whether the layout manager invalidates the specified portion of the glyph cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/setHardInvalidation(_:forGlyphRange:)
func (a_ ATSTypesetter) SetHardInvalidationForGlyphRange(flag bool /* primitive/slice/pointer. */, glyphRange foundation.objc.IObject /* cross-framework Range */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHardInvalidation:forGlyphRange:"), flag, glyphRange)
}


// Sets the glyph range being processed and the paragraph separator glyph range (the range of the paragraph separator character or characters).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/setParagraphGlyphRange(_:separatorGlyphRange:)
func (a_ ATSTypesetter) SetParagraphGlyphRangeSeparatorGlyphRange(paragraphRange foundation.objc.IObject /* cross-framework Range */, paragraphSeparatorRange foundation.objc.IObject /* cross-framework Range */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setParagraphGlyphRange:separatorGlyphRange:"), paragraphRange, paragraphSeparatorRange)
}


// Breaks a line by hyphenating before the character at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/shouldBreakLine(byHyphenatingBeforeCharacterAt:)
func (a_ ATSTypesetter) ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(charIndex uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldBreakLineByHyphenatingBeforeCharacterAtIndex:"), charIndex)
	return rv
}


// Breaks a line by word-wrapping before the character at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/shouldBreakLine(byWordBeforeCharacterAt:)
func (a_ ATSTypesetter) ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldBreakLineByWordBeforeCharacterAtIndex:"), charIndex)
	return rv
}


// Returns a screen font suitable for use in place of the specified original font,.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/substituteFont(for:)
func (a_ ATSTypesetter) SubstituteFontForFont(originalFont IFont) IFont {
	rv := objc.Send[Font](a_.ID, objc.Sel("substituteFontForFont:"), originalFont)
	return rv
}


// Returns the text tab closest to the specified glyph location and not beyond a maximum position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/textTab(forGlyphLocation:writingDirection:maxLocation:)
func (a_ ATSTypesetter) TextTabForGlyphLocationWritingDirectionMaxLocation(glyphLocation float64 /* primitive/slice/pointer. */, direction WritingDirection, maxLocation float64 /* primitive/slice/pointer. */) ITextTab {
	rv := objc.Send[TextTab](a_.ID, objc.Sel("textTabForGlyphLocation:writingDirection:maxLocation:"), glyphLocation, direction, maxLocation)
	return rv
}


// Notifies subclasses that the typesetter is about to set a new line fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/willSetLineFragmentRect(_:forGlyphRange:usedRect:baselineOffset:)
func (a_ ATSTypesetter) WillSetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(lineRect coregraphics.CGRect, glyphRange foundation.objc.IObject /* cross-framework Range */, usedRect coregraphics.CGRect, baselineOffset coregraphics.float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("willSetLineFragmentRect:forGlyphRange:usedRect:baselineOffset:"), lineRect, glyphRange, usedRect, baselineOffset)
}


// The backing store that contains the text on which this typesetter operates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/attributedString
func (a_ ATSTypesetter) AttributedString() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[AttributedString](a_.ID, objc.Sel("attributedString"))
	return rv
}


// The backing store that contains the text on which this typesetter operates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/attributedString
func (a_ ATSTypesetter) SetAttributedString(value objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributedString:"), value)
}


// A Boolean value controlling whether the typesetter performs bidirectional text processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/bidiProcessingEnabled
func (a_ ATSTypesetter) BidiProcessingEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("bidiProcessingEnabled"))
	return rv
}


// A Boolean value controlling whether the typesetter performs bidirectional text processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/bidiProcessingEnabled
func (a_ ATSTypesetter) SetBidiProcessingEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBidiProcessingEnabled:"), value)
}


// The text container for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/currentTextContainer
func (a_ ATSTypesetter) CurrentTextContainer() ITextContainer {
	rv := objc.Send[TextContainer](a_.ID, objc.Sel("currentTextContainer"))
	return rv
}


// The threshold controlling when hyphenation is attempted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/hyphenationFactor
func (a_ ATSTypesetter) HyphenationFactor() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](a_.ID, objc.Sel("hyphenationFactor"))
	return rv
}


// The threshold controlling when hyphenation is attempted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/hyphenationFactor
func (a_ ATSTypesetter) SetHyphenationFactor(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHyphenationFactor:"), value)
}


// The layout manager for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/layoutManager
func (a_ ATSTypesetter) LayoutManager() objc.IObject /* cross-framework: LayoutManager */ {
	rv := objc.Send[LayoutManager](a_.ID, objc.Sel("layoutManager"))
	return rv
}


// The amount (in points) by which text is inset within line fragment rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/lineFragmentPadding
func (a_ ATSTypesetter) LineFragmentPadding() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](a_.ID, objc.Sel("lineFragmentPadding"))
	return rv
}


// The amount (in points) by which text is inset within line fragment rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/lineFragmentPadding
func (a_ ATSTypesetter) SetLineFragmentPadding(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLineFragmentPadding:"), value)
}


// The current glyph range being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/paragraphGlyphRange
func (a_ ATSTypesetter) ParagraphGlyphRange() foundation.objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[foundation.Range](a_.ID, objc.Sel("paragraphGlyphRange"))
	return rv
}


// The current paragraph separator range that contains the current glyph range and extends from one paragraph separator character to the next.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/paragraphSeparatorGlyphRange
func (a_ ATSTypesetter) ParagraphSeparatorGlyphRange() foundation.objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[foundation.Range](a_.ID, objc.Sel("paragraphSeparatorGlyphRange"))
	return rv
}


// Returns a shared instance of the typesetter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/shared
func (a_ ATSTypesetter) SharedTypesetter() IATSTypesetter {
	rv := objc.Send[ATSTypesetter](a_.ID, objc.Sel("sharedTypesetter"))
	return rv
}


// The current typesetter behavior value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/typesetterBehavior
func (a_ ATSTypesetter) TypesetterBehavior() TypesetterBehavior {
	rv := objc.Send[TypesetterBehavior](a_.ID, objc.Sel("typesetterBehavior"))
	return rv
}


// The current typesetter behavior value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/typesetterBehavior
func (a_ ATSTypesetter) SetTypesetterBehavior(value TypesetterBehavior) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTypesetterBehavior:"), value)
}


// A Boolean value controlling whether the typesetter uses the leading (or line gap) value specified in the font metric information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/usesFontLeading
func (a_ ATSTypesetter) UsesFontLeading() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("usesFontLeading"))
	return rv
}


// A Boolean value controlling whether the typesetter uses the leading (or line gap) value specified in the font metric information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter/usesFontLeading
func (a_ ATSTypesetter) SetUsesFontLeading(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUsesFontLeading:"), value)
}



