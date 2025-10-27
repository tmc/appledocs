// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	

	// properties:
	AttributedString() foundation.foundation.INSAttributedString
	SetAttributedString(value foundation.foundation.INSAttributedString)
	AttributesForExtraLineFragment() foundation.IDictionary
	BidiProcessingEnabled() bool
	SetBidiProcessingEnabled(value bool)
	CurrentParagraphStyle() IParagraphStyle
	CurrentTextContainer() ITextContainer
	HyphenationFactor() float32
	SetHyphenationFactor(value float32)
	LayoutManager() ILayoutManager
	LineFragmentPadding() float64
	SetLineFragmentPadding(value float64)
	ParagraphCharacterRange() foundation.Range
	ParagraphGlyphRange() foundation.Range
	ParagraphSeparatorCharacterRange() foundation.Range
	ParagraphSeparatorGlyphRange() foundation.Range
	TextContainers() []TextContainer
	TypesetterBehavior() TypesetterBehavior
	SetTypesetterBehavior(value TypesetterBehavior)
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)


	

	// methods:
	ActionForControlCharacterAtIndex(charIndex uint) TypesetterControlCharacterAction
	BaselineOffsetInLayoutManagerGlyphIndex(layoutMgr ILayoutManager, glyphIndex uint) float64
	BeginLineWithGlyphAtIndex(glyphIndex uint)
	BeginParagraph()
	BoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(glyphIndex uint, textContainer ITextContainer, proposedRect corefoundation.CGRect, glyphPosition corefoundation.CGPoint, charIndex uint) corefoundation.CGRect
	CharacterRangeForGlyphRangeActualGlyphRange(glyphRange foundation.Range, actualGlyphRange RangePointer /* not a class type */) foundation.Range
	EndLineWithGlyphRange(lineGlyphRange foundation.Range)
	EndParagraph()
	GetLineFragmentRectUsedRectForParagraphSeparatorGlyphRangeAtProposedOrigin(lineFragmentRect RectPointer /* not a class type */, lineFragmentUsedRect RectPointer /* not a class type */, paragraphSeparatorGlyphRange foundation.Range, lineOrigin corefoundation.CGPoint)
	GetLineFragmentRectUsedRectRemainingRectForStartingGlyphAtIndexProposedRectLineSpacingParagraphSpacingBeforeParagraphSpacingAfter(lineFragmentRect RectPointer /* not a class type */, lineFragmentUsedRect RectPointer /* not a class type */, remainingRect RectPointer /* not a class type */, startingGlyphIndex uint, proposedRect corefoundation.CGRect, lineSpacing float64, paragraphSpacingBefore float64, paragraphSpacingAfter float64)
	GlyphRangeForCharacterRangeActualCharacterRange(charRange foundation.Range, actualCharRange RangePointer /* not a class type */) foundation.Range
	HyphenCharacterForGlyphAtIndex(glyphIndex uint) objectivec.IObject
	HyphenationFactorForGlyphAtIndex(glyphIndex uint) float32
	LayoutCharactersInRangeForLayoutManagerMaximumNumberOfLineFragments(characterRange foundation.Range, layoutManager ILayoutManager, maxNumLines uint) foundation.Range
	LayoutGlyphsInLayoutManagerStartingAtGlyphIndexMaxNumberOfLineFragmentsNextGlyphIndex(layoutManager ILayoutManager, startGlyphIndex uint, maxNumLines uint, nextGlyph uint)
	LayoutParagraphAtPoint(lineFragmentOrigin PointPointer /* not a class type */) uint
	LineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect corefoundation.CGRect) float64
	ParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect corefoundation.CGRect) float64
	ParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect corefoundation.CGRect) float64
	SetAttachmentSizeForGlyphRange(attachmentSize corefoundation.CGSize, glyphRange foundation.Range)
	SetBidiLevelsForGlyphRange(levels objectivec.IObject, glyphRange foundation.Range)
	SetDrawsOutsideLineFragmentForGlyphRange(flag bool, glyphRange foundation.Range)
	SetHardInvalidationForGlyphRange(flag bool, glyphRange foundation.Range)
	SetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(fragmentRect corefoundation.CGRect, glyphRange foundation.Range, usedRect corefoundation.CGRect, baselineOffset float64)
	SetLocationWithAdvancementsForStartOfGlyphRange(location corefoundation.CGPoint, advancements corefoundation.CGFloat, glyphRange foundation.Range)
	SetNotShownAttributeForGlyphRange(flag bool, glyphRange foundation.Range)
	SetParagraphGlyphRangeSeparatorGlyphRange(paragraphRange foundation.Range, paragraphSeparatorRange foundation.Range)
	ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(charIndex uint) bool
	ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint) bool
	SubstituteFontForFont(originalFont IFont) IFont
	TextTabForGlyphLocationWritingDirectionMaxLocation(glyphLocation float64, direction WritingDirection, maxLocation float64) ITextTab
	WillSetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(lineRect RectPointer /* not a class type */, glyphRange foundation.Range, usedRect RectPointer /* not a class type */, baselineOffset corefoundation.CGFloat)


}





// Alloc allocates a new instance without initialization.
func (tc _TypesetterClass) Alloc() Typesetter {
	rv := objc.Send[Typesetter](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An abstract class that performs various type layout tasks.
//
// uses concrete subclasses of to perform line layout, which includes word wrapping, hyphenation, and line breaking in either vertical or horizontal rectangles. By default, the text system uses the concrete subclass .


// An abstract class that performs various type layout tasks.
//
// [Full Topic]
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










// Returns the interglyph spacing in the specified range when sent to a printer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/printingAdjustment(in:forNominallySpacedGlyphRange:packedGlyphs:count:)
func (tc _TypesetterClass) PrintingAdjustmentInLayoutManagerForNominallySpacedGlyphRangePackedGlyphsCount(layoutMgr ILayoutManager, nominallySpacedGlyphsRange foundation.Range, packedGlyphs objectivec.IObject, packedGlyphsCount uint) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](objc.ID(tc.class), objc.Sel("printingAdjustmentInLayoutManager:forNominallySpacedGlyphRange:packedGlyphs:count:"), layoutMgr, nominallySpacedGlyphsRange, packedGlyphs, packedGlyphsCount)
	return rv
}


// Returns a shared instance of a reentrant typesetter that implements typesetting with the specified behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/sharedSystemTypesetter(for:)
func (tc _TypesetterClass) SharedSystemTypesetterForBehavior(behavior TypesetterBehavior) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("sharedSystemTypesetterForBehavior:"), behavior)
	return rv
}







// Returns the default typesetter behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/defaultTypesetterBehavior
func (tc _TypesetterClass) DefaultTypesetterBehavior() TypesetterBehavior {
	rv := objc.Send[TypesetterBehavior](objc.ID(tc.class), objc.Sel("defaultTypesetterBehavior"))
	return rv
}

// Returns a shared instance of a reentrant typesetter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/sharedSystemTypesetter
func (tc _TypesetterClass) SharedSystemTypesetter() Typesetter {
	rv := objc.Send[Typesetter](objc.ID(tc.class), objc.Sel("sharedSystemTypesetter"))
	return rv
}






// Returns the action associated with a control character.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/actionForControlCharacter(at:)
func (t_ Typesetter) ActionForControlCharacterAtIndex(charIndex uint) TypesetterControlCharacterAction {
	rv := objc.Send[TypesetterControlCharacterAction](t_.ID, objc.Sel("actionForControlCharacterAtIndex:"), charIndex)
	return rv
}


// Returns the distance from the bottom of the line fragment rectangle in which the glyph resides to the glyph baseline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/baselineOffset(in:glyphIndex:)
func (t_ Typesetter) BaselineOffsetInLayoutManagerGlyphIndex(layoutMgr ILayoutManager, glyphIndex uint) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("baselineOffsetInLayoutManager:glyphIndex:"), layoutMgr, glyphIndex)
	return rv
}


// Sets up layout parameters at the beginning of a line during typesetting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/beginLine(withGlyphAt:)
func (t_ Typesetter) BeginLineWithGlyphAtIndex(glyphIndex uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("beginLineWithGlyphAtIndex:"), glyphIndex)
}


// Sets up layout parameters at the beginning of a paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/beginParagraph()
func (t_ Typesetter) BeginParagraph() {
	objc.Send[objc.ID](t_.ID, objc.Sel("beginParagraph"))
}


// Returns the bounding rectangle for the specified control glyph with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/boundingBox(forControlGlyphAt:for:proposedLineFragment:glyphPosition:characterIndex:)
func (t_ Typesetter) BoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(glyphIndex uint, textContainer ITextContainer, proposedRect corefoundation.CGRect, glyphPosition corefoundation.CGPoint, charIndex uint) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("boundingBoxForControlGlyphAtIndex:forTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), glyphIndex, textContainer, proposedRect, glyphPosition, charIndex)
	return rv
}


// Returns the range for the characters in the receiver’s text store that are mapped to the specified glyphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/characterRange(forGlyphRange:actualGlyphRange:)
func (t_ Typesetter) CharacterRangeForGlyphRangeActualGlyphRange(glyphRange foundation.Range, actualGlyphRange RangePointer /* not a class type */) foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("characterRangeForGlyphRange:actualGlyphRange:"), glyphRange, actualGlyphRange)
	return rv
}


// Sets up layout parameters at the end of a line during typesetting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/endLine(withGlyphRange:)
func (t_ Typesetter) EndLineWithGlyphRange(lineGlyphRange foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("endLineWithGlyphRange:"), lineGlyphRange)
}


// Sets up layout parameters at the end of a paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/endParagraph()
func (t_ Typesetter) EndParagraph() {
	objc.Send[objc.ID](t_.ID, objc.Sel("endParagraph"))
}


// Calculates the line fragment rectangle and line fragment used rectangle for blank lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/getLineFragmentRect(_:usedRect:forParagraphSeparatorGlyphRange:atProposedOrigin:)
func (t_ Typesetter) GetLineFragmentRectUsedRectForParagraphSeparatorGlyphRangeAtProposedOrigin(lineFragmentRect RectPointer /* not a class type */, lineFragmentUsedRect RectPointer /* not a class type */, paragraphSeparatorGlyphRange foundation.Range, lineOrigin corefoundation.CGPoint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("getLineFragmentRect:usedRect:forParagraphSeparatorGlyphRange:atProposedOrigin:"), lineFragmentRect, lineFragmentUsedRect, paragraphSeparatorGlyphRange, lineOrigin)
}


// Calculates line fragment rectangle, line fragment used rectangle, and remaining rectangle for a line fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/getLineFragmentRect(_:usedRect:remaining:forStartingGlyphAt:proposedRect:lineSpacing:paragraphSpacingBefore:paragraphSpacingAfter:)
func (t_ Typesetter) GetLineFragmentRectUsedRectRemainingRectForStartingGlyphAtIndexProposedRectLineSpacingParagraphSpacingBeforeParagraphSpacingAfter(lineFragmentRect RectPointer /* not a class type */, lineFragmentUsedRect RectPointer /* not a class type */, remainingRect RectPointer /* not a class type */, startingGlyphIndex uint, proposedRect corefoundation.CGRect, lineSpacing float64, paragraphSpacingBefore float64, paragraphSpacingAfter float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("getLineFragmentRect:usedRect:remainingRect:forStartingGlyphAtIndex:proposedRect:lineSpacing:paragraphSpacingBefore:paragraphSpacingAfter:"), lineFragmentRect, lineFragmentUsedRect, remainingRect, startingGlyphIndex, proposedRect, lineSpacing, paragraphSpacingBefore, paragraphSpacingAfter)
}


// Returns the range for the glyphs mapped to the characters of the text store in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/glyphRange(forCharacterRange:actualCharacterRange:)
func (t_ Typesetter) GlyphRangeForCharacterRangeActualCharacterRange(charRange foundation.Range, actualCharRange RangePointer /* not a class type */) foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("glyphRangeForCharacterRange:actualCharacterRange:"), charRange, actualCharRange)
	return rv
}


// Returns the hyphen character to be inserted after the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/hyphenCharacter(forGlyphAt:)
func (t_ Typesetter) HyphenCharacterForGlyphAtIndex(glyphIndex uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("hyphenCharacterForGlyphAtIndex:"), glyphIndex)
	return rv
}


// Returns the hyphenation factor in effect at a specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/hyphenationFactor(forGlyphAt:)
func (t_ Typesetter) HyphenationFactorForGlyphAtIndex(glyphIndex uint) float32 {
	rv := objc.Send[float32](t_.ID, objc.Sel("hyphenationFactorForGlyphAtIndex:"), glyphIndex)
	return rv
}


// Lays out characters in the given character range for the specified layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/layoutCharacters(in:for:maximumNumberOfLineFragments:)
func (t_ Typesetter) LayoutCharactersInRangeForLayoutManagerMaximumNumberOfLineFragments(characterRange foundation.Range, layoutManager ILayoutManager, maxNumLines uint) foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("layoutCharactersInRange:forLayoutManager:maximumNumberOfLineFragments:"), characterRange, layoutManager, maxNumLines)
	return rv
}


// Lays out glyphs in the specified layout manager starting at a specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/layoutGlyphs(in:startingAtGlyphIndex:maxNumberOfLineFragments:nextGlyphIndex:)
func (t_ Typesetter) LayoutGlyphsInLayoutManagerStartingAtGlyphIndexMaxNumberOfLineFragmentsNextGlyphIndex(layoutManager ILayoutManager, startGlyphIndex uint, maxNumLines uint, nextGlyph uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("layoutGlyphsInLayoutManager:startingAtGlyphIndex:maxNumberOfLineFragments:nextGlyphIndex:"), layoutManager, startGlyphIndex, maxNumLines, nextGlyph)
}


// Lays out glyphs in the current glyph range until the next paragraph separator is reached.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/layoutParagraph(at:)
func (t_ Typesetter) LayoutParagraphAtPoint(lineFragmentOrigin PointPointer /* not a class type */) uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("layoutParagraphAtPoint:"), lineFragmentOrigin)
	return rv
}


// Returns the line spacing in effect following the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/lineSpacing(afterGlyphAt:withProposedLineFragmentRect:)
func (t_ Typesetter) LineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect corefoundation.CGRect) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("lineSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}


// Returns the paragraph spacing that is in effect after the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphSpacing(afterGlyphAt:withProposedLineFragmentRect:)
func (t_ Typesetter) ParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect corefoundation.CGRect) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("paragraphSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}


// Returns the number of points of space—added before a paragraph—that is in effect before the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphSpacing(beforeGlyphAt:withProposedLineFragmentRect:)
func (t_ Typesetter) ParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect corefoundation.CGRect) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("paragraphSpacingBeforeGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}


// Sets the size the specified glyphs (assumed to be attachments) will be asked to draw themselves at.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setAttachmentSize(_:forGlyphRange:)
func (t_ Typesetter) SetAttachmentSizeForGlyphRange(attachmentSize corefoundation.CGSize, glyphRange foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttachmentSize:forGlyphRange:"), attachmentSize, glyphRange)
}


// Sets the direction of the specified glyphs for bidirectional text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setBidiLevels(_:forGlyphRange:)
func (t_ Typesetter) SetBidiLevelsForGlyphRange(levels objectivec.IObject, glyphRange foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBidiLevels:forGlyphRange:"), levels, glyphRange)
}


// Sets whether the specified glyphs exceed the bounds of the line fragment in which they are laid out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setDrawsOutsideLineFragment(_:forGlyphRange:)
func (t_ Typesetter) SetDrawsOutsideLineFragmentForGlyphRange(flag bool, glyphRange foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsOutsideLineFragment:forGlyphRange:"), flag, glyphRange)
}


// Sets whether to force the layout manager to invalidate the specified portion of the glyph cache when invalidating layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setHardInvalidation(_:forGlyphRange:)
func (t_ Typesetter) SetHardInvalidationForGlyphRange(flag bool, glyphRange foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHardInvalidation:forGlyphRange:"), flag, glyphRange)
}


// Sets the line fragment rectangle where the specified glyphs are laid out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setLineFragmentRect(_:forGlyphRange:usedRect:baselineOffset:)
func (t_ Typesetter) SetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(fragmentRect corefoundation.CGRect, glyphRange foundation.Range, usedRect corefoundation.CGRect, baselineOffset float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineFragmentRect:forGlyphRange:usedRect:baselineOffset:"), fragmentRect, glyphRange, usedRect, baselineOffset)
}


// Sets the location where the specified glyphs are laid out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setLocation(_:withAdvancements:forStartOfGlyphRange:)
func (t_ Typesetter) SetLocationWithAdvancementsForStartOfGlyphRange(location corefoundation.CGPoint, advancements corefoundation.CGFloat, glyphRange foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLocation:withAdvancements:forStartOfGlyphRange:"), location, advancements, glyphRange)
}


// Sets whether the specified glyphs are not shown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setNotShownAttribute(_:forGlyphRange:)
func (t_ Typesetter) SetNotShownAttributeForGlyphRange(flag bool, glyphRange foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNotShownAttribute:forGlyphRange:"), flag, glyphRange)
}


// Sets the current glyph range being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setParagraphGlyphRange(_:separatorGlyphRange:)
func (t_ Typesetter) SetParagraphGlyphRangeSeparatorGlyphRange(paragraphRange foundation.Range, paragraphSeparatorRange foundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParagraphGlyphRange:separatorGlyphRange:"), paragraphRange, paragraphSeparatorRange)
}


// Returns whether the line being laid out should be broken by hyphenating at the specified character.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/shouldBreakLine(byHyphenatingBeforeCharacterAt:)
func (t_ Typesetter) ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(charIndex uint) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("shouldBreakLineByHyphenatingBeforeCharacterAtIndex:"), charIndex)
	return rv
}


// Returns whether the line being laid out should be broken by a word break at the specified character.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/shouldBreakLine(byWordBeforeCharacterAt:)
func (t_ Typesetter) ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("shouldBreakLineByWordBeforeCharacterAtIndex:"), charIndex)
	return rv
}


// Returns a screen font suitable for use in place of a given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/substituteFont(for:)
func (t_ Typesetter) SubstituteFontForFont(originalFont IFont) IFont {
	rv := objc.Send[Font](t_.ID, objc.Sel("substituteFontForFont:"), originalFont)
	return rv
}


// Returns the text tab next closest to a given glyph location within the given parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/textTab(forGlyphLocation:writingDirection:maxLocation:)
func (t_ Typesetter) TextTabForGlyphLocationWritingDirectionMaxLocation(glyphLocation float64, direction WritingDirection, maxLocation float64) ITextTab {
	rv := objc.Send[TextTab](t_.ID, objc.Sel("textTabForGlyphLocation:writingDirection:maxLocation:"), glyphLocation, direction, maxLocation)
	return rv
}


// Called by the typesetter just prior to storing the actual line fragment rectangle location in the layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/willSetLineFragmentRect(_:forGlyphRange:usedRect:baselineOffset:)
func (t_ Typesetter) WillSetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(lineRect RectPointer /* not a class type */, glyphRange foundation.Range, usedRect RectPointer /* not a class type */, baselineOffset corefoundation.CGFloat) {
	objc.Send[objc.ID](t_.ID, objc.Sel("willSetLineFragmentRect:forGlyphRange:usedRect:baselineOffset:"), lineRect, glyphRange, usedRect, baselineOffset)
}







// Returns the text backing store, usually an instance of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/attributedString
func (t_ Typesetter) AttributedString() foundation.foundation.INSAttributedString {
	rv := objc.Send[foundation.NSAttributedString](t_.ID, objc.Sel("attributedString"))
	return rv
}


// Returns the text backing store, usually an instance of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/attributedString
func (t_ Typesetter) SetAttributedString(value foundation.foundation.INSAttributedString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttributedString:"), value)
}


// Returns the attributes used to lay out the extra line fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/attributesForExtraLineFragment
func (t_ Typesetter) AttributesForExtraLineFragment() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("attributesForExtraLineFragment"))
	return rv
}


// Returns whether bidirectional text processing is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/bidiProcessingEnabled
func (t_ Typesetter) BidiProcessingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("bidiProcessingEnabled"))
	return rv
}


// Returns whether bidirectional text processing is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/bidiProcessingEnabled
func (t_ Typesetter) SetBidiProcessingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBidiProcessingEnabled:"), value)
}


// Returns the paragraph style object for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/currentParagraphStyle
func (t_ Typesetter) CurrentParagraphStyle() IParagraphStyle {
	rv := objc.Send[ParagraphStyle](t_.ID, objc.Sel("currentParagraphStyle"))
	return rv
}


// Returns the text container for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/currentTextContainer
func (t_ Typesetter) CurrentTextContainer() ITextContainer {
	rv := objc.Send[TextContainer](t_.ID, objc.Sel("currentTextContainer"))
	return rv
}


// Returns the default typesetter behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/defaultTypesetterBehavior
func (t_ Typesetter) DefaultTypesetterBehavior() TypesetterBehavior {
	rv := objc.Send[TypesetterBehavior](t_.ID, objc.Sel("defaultTypesetterBehavior"))
	return rv
}


// Returns the current hyphenation factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/hyphenationFactor
func (t_ Typesetter) HyphenationFactor() float32 {
	rv := objc.Send[float32](t_.ID, objc.Sel("hyphenationFactor"))
	return rv
}


// Returns the current hyphenation factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/hyphenationFactor
func (t_ Typesetter) SetHyphenationFactor(value float32) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHyphenationFactor:"), value)
}


// Returns the layout manager for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/layoutManager
func (t_ Typesetter) LayoutManager() ILayoutManager {
	rv := objc.Send[LayoutManager](t_.ID, objc.Sel("layoutManager"))
	return rv
}


// Returns the current line fragment padding, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/lineFragmentPadding
func (t_ Typesetter) LineFragmentPadding() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("lineFragmentPadding"))
	return rv
}


// Returns the current line fragment padding, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/lineFragmentPadding
func (t_ Typesetter) SetLineFragmentPadding(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineFragmentPadding:"), value)
}


// Returns the character range currently being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphCharacterRange
func (t_ Typesetter) ParagraphCharacterRange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("paragraphCharacterRange"))
	return rv
}


// Returns the glyph range currently being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphGlyphRange
func (t_ Typesetter) ParagraphGlyphRange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("paragraphGlyphRange"))
	return rv
}


// Returns the current paragraph separator character range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphSeparatorCharacterRange
func (t_ Typesetter) ParagraphSeparatorCharacterRange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("paragraphSeparatorCharacterRange"))
	return rv
}


// Returns the current paragraph separator range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphSeparatorGlyphRange
func (t_ Typesetter) ParagraphSeparatorGlyphRange() foundation.Range {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("paragraphSeparatorGlyphRange"))
	return rv
}


// Returns a shared instance of a reentrant typesetter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/sharedSystemTypesetter
func (t_ Typesetter) SharedSystemTypesetter() ITypesetter {
	rv := objc.Send[Typesetter](t_.ID, objc.Sel("sharedSystemTypesetter"))
	return rv
}


// Returns an array containing the text containers belonging to the current layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/textContainers
func (t_ Typesetter) TextContainers() []TextContainer {
	rv := objc.Send[[]TextContainer](t_.ID, objc.Sel("textContainers"))
	return rv
}


// Returns the current typesetter behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/typesetterBehavior
func (t_ Typesetter) TypesetterBehavior() TypesetterBehavior {
	rv := objc.Send[TypesetterBehavior](t_.ID, objc.Sel("typesetterBehavior"))
	return rv
}


// Returns the current typesetter behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/typesetterBehavior
func (t_ Typesetter) SetTypesetterBehavior(value TypesetterBehavior) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTypesetterBehavior:"), value)
}


// Returns whether the typesetter uses the leading (or line gap) value specified in the font metric information of the current font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/usesFontLeading
func (t_ Typesetter) UsesFontLeading() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontLeading"))
	return rv
}


// Returns whether the typesetter uses the leading (or line gap) value specified in the font metric information of the current font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/usesFontLeading
func (t_ Typesetter) SetUsesFontLeading(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontLeading:"), value)
}








