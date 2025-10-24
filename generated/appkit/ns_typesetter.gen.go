// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSTypesetter */


/* debug [class_header]: Header for NSTypesetter */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Typesetter */
// An interface definition for the [Typesetter] class.
type ITypesetter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Typesetter */
	// properties:
	AttributedString() foundation.AttributedString
	SetAttributedString(value foundation.AttributedString)
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
	ParagraphCharacterRange() corefoundation.Range
	ParagraphGlyphRange() corefoundation.Range
	ParagraphSeparatorCharacterRange() corefoundation.Range
	ParagraphSeparatorGlyphRange() corefoundation.Range
	TextContainers() []TextContainer
	TypesetterBehavior() TypesetterBehavior
	SetTypesetterBehavior(value TypesetterBehavior)
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Typesetter */
	// methods:
	ActionForControlCharacterAtIndex(charIndex uint) TypesetterControlCharacterAction
	BaselineOffsetInLayoutManagerGlyphIndex(layoutMgr ILayoutManager, glyphIndex uint) float64
	BeginLineWithGlyphAtIndex(glyphIndex uint)
	BeginParagraph()
	BoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(glyphIndex uint, textContainer ITextContainer, proposedRect Rect /* not a class type */, glyphPosition vision.Point, charIndex uint) Rect /* not a class type */
	CharacterRangeForGlyphRangeActualGlyphRange(glyphRange corefoundation.Range, actualGlyphRange RangePointer /* not a class type */) corefoundation.Range
	EndLineWithGlyphRange(lineGlyphRange corefoundation.Range)
	EndParagraph()
	GetLineFragmentRectUsedRectForParagraphSeparatorGlyphRangeAtProposedOrigin(lineFragmentRect RectPointer /* not a class type */, lineFragmentUsedRect RectPointer /* not a class type */, paragraphSeparatorGlyphRange corefoundation.Range, lineOrigin vision.Point)
	GetLineFragmentRectUsedRectRemainingRectForStartingGlyphAtIndexProposedRectLineSpacingParagraphSpacingBeforeParagraphSpacingAfter(lineFragmentRect RectPointer /* not a class type */, lineFragmentUsedRect RectPointer /* not a class type */, remainingRect RectPointer /* not a class type */, startingGlyphIndex uint, proposedRect Rect /* not a class type */, lineSpacing float64, paragraphSpacingBefore float64, paragraphSpacingAfter float64)
	GlyphRangeForCharacterRangeActualCharacterRange(charRange corefoundation.Range, actualCharRange RangePointer /* not a class type */) corefoundation.Range
	HyphenCharacterForGlyphAtIndex(glyphIndex uint) objectivec.IObject
	HyphenationFactorForGlyphAtIndex(glyphIndex uint) float32
	LayoutCharactersInRangeForLayoutManagerMaximumNumberOfLineFragments(characterRange corefoundation.Range, layoutManager ILayoutManager, maxNumLines uint) corefoundation.Range
	LayoutGlyphsInLayoutManagerStartingAtGlyphIndexMaxNumberOfLineFragmentsNextGlyphIndex(layoutManager ILayoutManager, startGlyphIndex uint, maxNumLines uint, nextGlyph uint)
	LayoutParagraphAtPoint(lineFragmentOrigin PointPointer /* not a class type */) uint
	LineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect Rect /* not a class type */) float64
	ParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect Rect /* not a class type */) float64
	ParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect Rect /* not a class type */) float64
	SetAttachmentSizeForGlyphRange(attachmentSize Size /* not a class type */, glyphRange corefoundation.Range)
	SetBidiLevelsForGlyphRange(levels objectivec.IObject, glyphRange corefoundation.Range)
	SetDrawsOutsideLineFragmentForGlyphRange(flag bool, glyphRange corefoundation.Range)
	SetHardInvalidationForGlyphRange(flag bool, glyphRange corefoundation.Range)
	SetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(fragmentRect Rect /* not a class type */, glyphRange corefoundation.Range, usedRect Rect /* not a class type */, baselineOffset float64)
	SetLocationWithAdvancementsForStartOfGlyphRange(location vision.Point, advancements corefoundation.CGFloat, glyphRange corefoundation.Range)
	SetNotShownAttributeForGlyphRange(flag bool, glyphRange corefoundation.Range)
	SetParagraphGlyphRangeSeparatorGlyphRange(paragraphRange corefoundation.Range, paragraphSeparatorRange corefoundation.Range)
	ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(charIndex uint) bool
	ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint) bool
	SubstituteFontForFont(originalFont IFont) IFont
	TextTabForGlyphLocationWritingDirectionMaxLocation(glyphLocation float64, direction WritingDirection, maxLocation float64) ITextTab
	WillSetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(lineRect RectPointer /* not a class type */, glyphRange corefoundation.Range, usedRect RectPointer /* not a class type */, baselineOffset corefoundation.CGFloat)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Typesetter */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Typesetter */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Typesetter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Typesetter */

// Returns the interglyph spacing in the specified range when sent to a printer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/printingAdjustment(in:forNominallySpacedGlyphRange:packedGlyphs:count:)
func (tc _TypesetterClass) PrintingAdjustmentInLayoutManagerForNominallySpacedGlyphRangePackedGlyphsCount(layoutMgr ILayoutManager, nominallySpacedGlyphsRange corefoundation.Range, packedGlyphs objectivec.IObject, packedGlyphsCount uint) Size /* not a class type */ {
	rv := objc.Send[Size](objc.ID(tc.class), objc.Sel("printingAdjustmentInLayoutManager:forNominallySpacedGlyphRange:packedGlyphs:count:"), layoutMgr, nominallySpacedGlyphsRange, packedGlyphs, packedGlyphsCount)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrintingAdjustmentInLayoutManagerForNominallySpacedGlyphRangePackedGlyphsCount) */


// Returns a shared instance of a reentrant typesetter that implements typesetting with the specified behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/sharedSystemTypesetter(for:)
func (tc _TypesetterClass) SharedSystemTypesetterForBehavior(behavior TypesetterBehavior) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("sharedSystemTypesetterForBehavior:"), behavior)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedSystemTypesetterForBehavior) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Typesetter */

// Returns the default typesetter behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/defaultTypesetterBehavior
func (tc _TypesetterClass) DefaultTypesetterBehavior() TypesetterBehavior {
	rv := objc.Send[TypesetterBehavior](objc.ID(tc.class), objc.Sel("defaultTypesetterBehavior"))
	return rv
}/* debug [class_properties_class/property]: defaultTypesetterBehavior */

// Returns a shared instance of a reentrant typesetter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/sharedSystemTypesetter
func (tc _TypesetterClass) SharedSystemTypesetter() Typesetter {
	rv := objc.Send[Typesetter](objc.ID(tc.class), objc.Sel("sharedSystemTypesetter"))
	return rv
}/* debug [class_properties_class/property]: sharedSystemTypesetter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Typesetter */

// Returns the action associated with a control character.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/actionForControlCharacter(at:)
func (t_ Typesetter) ActionForControlCharacterAtIndex(charIndex uint) TypesetterControlCharacterAction {
	rv := objc.Send[TypesetterControlCharacterAction](t_.ID, objc.Sel("actionForControlCharacterAtIndex:"), charIndex)
	return rv
}/* debug [instance_methods/method]: ActionForControlCharacterAtIndex */


// Returns the distance from the bottom of the line fragment rectangle in which the glyph resides to the glyph baseline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/baselineOffset(in:glyphIndex:)
func (t_ Typesetter) BaselineOffsetInLayoutManagerGlyphIndex(layoutMgr ILayoutManager, glyphIndex uint) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("baselineOffsetInLayoutManager:glyphIndex:"), layoutMgr, glyphIndex)
	return rv
}/* debug [instance_methods/method]: BaselineOffsetInLayoutManagerGlyphIndex */


// Sets up layout parameters at the beginning of a line during typesetting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/beginLine(withGlyphAt:)
func (t_ Typesetter) BeginLineWithGlyphAtIndex(glyphIndex uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("beginLineWithGlyphAtIndex:"), glyphIndex)
}/* debug [instance_methods/method]: BeginLineWithGlyphAtIndex */


// Sets up layout parameters at the beginning of a paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/beginParagraph()
func (t_ Typesetter) BeginParagraph() {
	objc.Send[objc.ID](t_.ID, objc.Sel("beginParagraph"))
}/* debug [instance_methods/method]: BeginParagraph */


// Returns the bounding rectangle for the specified control glyph with the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/boundingBox(forControlGlyphAt:for:proposedLineFragment:glyphPosition:characterIndex:)
func (t_ Typesetter) BoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(glyphIndex uint, textContainer ITextContainer, proposedRect Rect /* not a class type */, glyphPosition vision.Point, charIndex uint) Rect /* not a class type */ {
	rv := objc.Send[Rect](t_.ID, objc.Sel("boundingBoxForControlGlyphAtIndex:forTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), glyphIndex, textContainer, proposedRect, glyphPosition, charIndex)
	return rv
}/* debug [instance_methods/method]: BoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex */


// Returns the range for the characters in the receiver’s text store that are mapped to the specified glyphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/characterRange(forGlyphRange:actualGlyphRange:)
func (t_ Typesetter) CharacterRangeForGlyphRangeActualGlyphRange(glyphRange corefoundation.Range, actualGlyphRange RangePointer /* not a class type */) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("characterRangeForGlyphRange:actualGlyphRange:"), glyphRange, actualGlyphRange)
	return rv
}/* debug [instance_methods/method]: CharacterRangeForGlyphRangeActualGlyphRange */


// Sets up layout parameters at the end of a line during typesetting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/endLine(withGlyphRange:)
func (t_ Typesetter) EndLineWithGlyphRange(lineGlyphRange corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("endLineWithGlyphRange:"), lineGlyphRange)
}/* debug [instance_methods/method]: EndLineWithGlyphRange */


// Sets up layout parameters at the end of a paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/endParagraph()
func (t_ Typesetter) EndParagraph() {
	objc.Send[objc.ID](t_.ID, objc.Sel("endParagraph"))
}/* debug [instance_methods/method]: EndParagraph */


// Calculates the line fragment rectangle and line fragment used rectangle for blank lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/getLineFragmentRect(_:usedRect:forParagraphSeparatorGlyphRange:atProposedOrigin:)
func (t_ Typesetter) GetLineFragmentRectUsedRectForParagraphSeparatorGlyphRangeAtProposedOrigin(lineFragmentRect RectPointer /* not a class type */, lineFragmentUsedRect RectPointer /* not a class type */, paragraphSeparatorGlyphRange corefoundation.Range, lineOrigin vision.Point) {
	objc.Send[objc.ID](t_.ID, objc.Sel("getLineFragmentRect:usedRect:forParagraphSeparatorGlyphRange:atProposedOrigin:"), lineFragmentRect, lineFragmentUsedRect, paragraphSeparatorGlyphRange, lineOrigin)
}/* debug [instance_methods/method]: GetLineFragmentRectUsedRectForParagraphSeparatorGlyphRangeAtProposedOrigin */


// Calculates line fragment rectangle, line fragment used rectangle, and remaining rectangle for a line fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/getLineFragmentRect(_:usedRect:remaining:forStartingGlyphAt:proposedRect:lineSpacing:paragraphSpacingBefore:paragraphSpacingAfter:)
func (t_ Typesetter) GetLineFragmentRectUsedRectRemainingRectForStartingGlyphAtIndexProposedRectLineSpacingParagraphSpacingBeforeParagraphSpacingAfter(lineFragmentRect RectPointer /* not a class type */, lineFragmentUsedRect RectPointer /* not a class type */, remainingRect RectPointer /* not a class type */, startingGlyphIndex uint, proposedRect Rect /* not a class type */, lineSpacing float64, paragraphSpacingBefore float64, paragraphSpacingAfter float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("getLineFragmentRect:usedRect:remainingRect:forStartingGlyphAtIndex:proposedRect:lineSpacing:paragraphSpacingBefore:paragraphSpacingAfter:"), lineFragmentRect, lineFragmentUsedRect, remainingRect, startingGlyphIndex, proposedRect, lineSpacing, paragraphSpacingBefore, paragraphSpacingAfter)
}/* debug [instance_methods/method]: GetLineFragmentRectUsedRectRemainingRectForStartingGlyphAtIndexProposedRectLineSpacingParagraphSpacingBeforeParagraphSpacingAfter */


// Returns the range for the glyphs mapped to the characters of the text store in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/glyphRange(forCharacterRange:actualCharacterRange:)
func (t_ Typesetter) GlyphRangeForCharacterRangeActualCharacterRange(charRange corefoundation.Range, actualCharRange RangePointer /* not a class type */) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("glyphRangeForCharacterRange:actualCharacterRange:"), charRange, actualCharRange)
	return rv
}/* debug [instance_methods/method]: GlyphRangeForCharacterRangeActualCharacterRange */


// Returns the hyphen character to be inserted after the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/hyphenCharacter(forGlyphAt:)
func (t_ Typesetter) HyphenCharacterForGlyphAtIndex(glyphIndex uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("hyphenCharacterForGlyphAtIndex:"), glyphIndex)
	return rv
}/* debug [instance_methods/method]: HyphenCharacterForGlyphAtIndex */


// Returns the hyphenation factor in effect at a specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/hyphenationFactor(forGlyphAt:)
func (t_ Typesetter) HyphenationFactorForGlyphAtIndex(glyphIndex uint) float32 {
	rv := objc.Send[float32](t_.ID, objc.Sel("hyphenationFactorForGlyphAtIndex:"), glyphIndex)
	return rv
}/* debug [instance_methods/method]: HyphenationFactorForGlyphAtIndex */


// Lays out characters in the given character range for the specified layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/layoutCharacters(in:for:maximumNumberOfLineFragments:)
func (t_ Typesetter) LayoutCharactersInRangeForLayoutManagerMaximumNumberOfLineFragments(characterRange corefoundation.Range, layoutManager ILayoutManager, maxNumLines uint) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("layoutCharactersInRange:forLayoutManager:maximumNumberOfLineFragments:"), characterRange, layoutManager, maxNumLines)
	return rv
}/* debug [instance_methods/method]: LayoutCharactersInRangeForLayoutManagerMaximumNumberOfLineFragments */


// Lays out glyphs in the specified layout manager starting at a specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/layoutGlyphs(in:startingAtGlyphIndex:maxNumberOfLineFragments:nextGlyphIndex:)
func (t_ Typesetter) LayoutGlyphsInLayoutManagerStartingAtGlyphIndexMaxNumberOfLineFragmentsNextGlyphIndex(layoutManager ILayoutManager, startGlyphIndex uint, maxNumLines uint, nextGlyph uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("layoutGlyphsInLayoutManager:startingAtGlyphIndex:maxNumberOfLineFragments:nextGlyphIndex:"), layoutManager, startGlyphIndex, maxNumLines, nextGlyph)
}/* debug [instance_methods/method]: LayoutGlyphsInLayoutManagerStartingAtGlyphIndexMaxNumberOfLineFragmentsNextGlyphIndex */


// Lays out glyphs in the current glyph range until the next paragraph separator is reached.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/layoutParagraph(at:)
func (t_ Typesetter) LayoutParagraphAtPoint(lineFragmentOrigin PointPointer /* not a class type */) uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("layoutParagraphAtPoint:"), lineFragmentOrigin)
	return rv
}/* debug [instance_methods/method]: LayoutParagraphAtPoint */


// Returns the line spacing in effect following the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/lineSpacing(afterGlyphAt:withProposedLineFragmentRect:)
func (t_ Typesetter) LineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect Rect /* not a class type */) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("lineSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}/* debug [instance_methods/method]: LineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect */


// Returns the paragraph spacing that is in effect after the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphSpacing(afterGlyphAt:withProposedLineFragmentRect:)
func (t_ Typesetter) ParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect Rect /* not a class type */) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("paragraphSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}/* debug [instance_methods/method]: ParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect */


// Returns the number of points of space—added before a paragraph—that is in effect before the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphSpacing(beforeGlyphAt:withProposedLineFragmentRect:)
func (t_ Typesetter) ParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect Rect /* not a class type */) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("paragraphSpacingBeforeGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}/* debug [instance_methods/method]: ParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect */


// Sets the size the specified glyphs (assumed to be attachments) will be asked to draw themselves at.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setAttachmentSize(_:forGlyphRange:)
func (t_ Typesetter) SetAttachmentSizeForGlyphRange(attachmentSize Size /* not a class type */, glyphRange corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttachmentSize:forGlyphRange:"), attachmentSize, glyphRange)
}/* debug [instance_methods/method]: SetAttachmentSizeForGlyphRange */


// Sets the direction of the specified glyphs for bidirectional text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setBidiLevels(_:forGlyphRange:)
func (t_ Typesetter) SetBidiLevelsForGlyphRange(levels objectivec.IObject, glyphRange corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBidiLevels:forGlyphRange:"), levels, glyphRange)
}/* debug [instance_methods/method]: SetBidiLevelsForGlyphRange */


// Sets whether the specified glyphs exceed the bounds of the line fragment in which they are laid out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setDrawsOutsideLineFragment(_:forGlyphRange:)
func (t_ Typesetter) SetDrawsOutsideLineFragmentForGlyphRange(flag bool, glyphRange corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsOutsideLineFragment:forGlyphRange:"), flag, glyphRange)
}/* debug [instance_methods/method]: SetDrawsOutsideLineFragmentForGlyphRange */


// Sets whether to force the layout manager to invalidate the specified portion of the glyph cache when invalidating layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setHardInvalidation(_:forGlyphRange:)
func (t_ Typesetter) SetHardInvalidationForGlyphRange(flag bool, glyphRange corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHardInvalidation:forGlyphRange:"), flag, glyphRange)
}/* debug [instance_methods/method]: SetHardInvalidationForGlyphRange */


// Sets the line fragment rectangle where the specified glyphs are laid out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setLineFragmentRect(_:forGlyphRange:usedRect:baselineOffset:)
func (t_ Typesetter) SetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(fragmentRect Rect /* not a class type */, glyphRange corefoundation.Range, usedRect Rect /* not a class type */, baselineOffset float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineFragmentRect:forGlyphRange:usedRect:baselineOffset:"), fragmentRect, glyphRange, usedRect, baselineOffset)
}/* debug [instance_methods/method]: SetLineFragmentRectForGlyphRangeUsedRectBaselineOffset */


// Sets the location where the specified glyphs are laid out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setLocation(_:withAdvancements:forStartOfGlyphRange:)
func (t_ Typesetter) SetLocationWithAdvancementsForStartOfGlyphRange(location vision.Point, advancements corefoundation.CGFloat, glyphRange corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLocation:withAdvancements:forStartOfGlyphRange:"), location, advancements, glyphRange)
}/* debug [instance_methods/method]: SetLocationWithAdvancementsForStartOfGlyphRange */


// Sets whether the specified glyphs are not shown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setNotShownAttribute(_:forGlyphRange:)
func (t_ Typesetter) SetNotShownAttributeForGlyphRange(flag bool, glyphRange corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNotShownAttribute:forGlyphRange:"), flag, glyphRange)
}/* debug [instance_methods/method]: SetNotShownAttributeForGlyphRange */


// Sets the current glyph range being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setParagraphGlyphRange(_:separatorGlyphRange:)
func (t_ Typesetter) SetParagraphGlyphRangeSeparatorGlyphRange(paragraphRange corefoundation.Range, paragraphSeparatorRange corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParagraphGlyphRange:separatorGlyphRange:"), paragraphRange, paragraphSeparatorRange)
}/* debug [instance_methods/method]: SetParagraphGlyphRangeSeparatorGlyphRange */


// Returns whether the line being laid out should be broken by hyphenating at the specified character.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/shouldBreakLine(byHyphenatingBeforeCharacterAt:)
func (t_ Typesetter) ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(charIndex uint) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("shouldBreakLineByHyphenatingBeforeCharacterAtIndex:"), charIndex)
	return rv
}/* debug [instance_methods/method]: ShouldBreakLineByHyphenatingBeforeCharacterAtIndex */


// Returns whether the line being laid out should be broken by a word break at the specified character.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/shouldBreakLine(byWordBeforeCharacterAt:)
func (t_ Typesetter) ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("shouldBreakLineByWordBeforeCharacterAtIndex:"), charIndex)
	return rv
}/* debug [instance_methods/method]: ShouldBreakLineByWordBeforeCharacterAtIndex */


// Returns a screen font suitable for use in place of a given font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/substituteFont(for:)
func (t_ Typesetter) SubstituteFontForFont(originalFont IFont) IFont {
	rv := objc.Send[Font](t_.ID, objc.Sel("substituteFontForFont:"), originalFont)
	return rv
}/* debug [instance_methods/method]: SubstituteFontForFont */


// Returns the text tab next closest to a given glyph location within the given parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/textTab(forGlyphLocation:writingDirection:maxLocation:)
func (t_ Typesetter) TextTabForGlyphLocationWritingDirectionMaxLocation(glyphLocation float64, direction WritingDirection, maxLocation float64) ITextTab {
	rv := objc.Send[TextTab](t_.ID, objc.Sel("textTabForGlyphLocation:writingDirection:maxLocation:"), glyphLocation, direction, maxLocation)
	return rv
}/* debug [instance_methods/method]: TextTabForGlyphLocationWritingDirectionMaxLocation */


// Called by the typesetter just prior to storing the actual line fragment rectangle location in the layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/willSetLineFragmentRect(_:forGlyphRange:usedRect:baselineOffset:)
func (t_ Typesetter) WillSetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(lineRect RectPointer /* not a class type */, glyphRange corefoundation.Range, usedRect RectPointer /* not a class type */, baselineOffset corefoundation.CGFloat) {
	objc.Send[objc.ID](t_.ID, objc.Sel("willSetLineFragmentRect:forGlyphRange:usedRect:baselineOffset:"), lineRect, glyphRange, usedRect, baselineOffset)
}/* debug [instance_methods/method]: WillSetLineFragmentRectForGlyphRangeUsedRectBaselineOffset */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Typesetter */

// Returns the text backing store, usually an instance of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/attributedString
func (t_ Typesetter) AttributedString() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](t_.ID, objc.Sel("attributedString"))
	return rv
}/* debug [instance_properties/getter]: attributedString */


// Returns the text backing store, usually an instance of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/attributedString
func (t_ Typesetter) SetAttributedString(value foundation.AttributedString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttributedString:"), value)
}/* debug [instance_properties/setter]: attributedString */


// Returns the attributes used to lay out the extra line fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/attributesForExtraLineFragment
func (t_ Typesetter) AttributesForExtraLineFragment() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("attributesForExtraLineFragment"))
	return rv
}/* debug [instance_properties/getter]: attributesForExtraLineFragment */


// Returns whether bidirectional text processing is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/bidiProcessingEnabled
func (t_ Typesetter) BidiProcessingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("bidiProcessingEnabled"))
	return rv
}/* debug [instance_properties/getter]: bidiProcessingEnabled */


// Returns whether bidirectional text processing is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/bidiProcessingEnabled
func (t_ Typesetter) SetBidiProcessingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBidiProcessingEnabled:"), value)
}/* debug [instance_properties/setter]: bidiProcessingEnabled */


// Returns the paragraph style object for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/currentParagraphStyle
func (t_ Typesetter) CurrentParagraphStyle() IParagraphStyle {
	rv := objc.Send[ParagraphStyle](t_.ID, objc.Sel("currentParagraphStyle"))
	return rv
}/* debug [instance_properties/getter]: currentParagraphStyle */


// Returns the text container for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/currentTextContainer
func (t_ Typesetter) CurrentTextContainer() ITextContainer {
	rv := objc.Send[TextContainer](t_.ID, objc.Sel("currentTextContainer"))
	return rv
}/* debug [instance_properties/getter]: currentTextContainer */


// Returns the default typesetter behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/defaultTypesetterBehavior
func (t_ Typesetter) DefaultTypesetterBehavior() TypesetterBehavior {
	rv := objc.Send[TypesetterBehavior](t_.ID, objc.Sel("defaultTypesetterBehavior"))
	return rv
}/* debug [instance_properties/getter]: defaultTypesetterBehavior */


// Returns the current hyphenation factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/hyphenationFactor
func (t_ Typesetter) HyphenationFactor() float32 {
	rv := objc.Send[float32](t_.ID, objc.Sel("hyphenationFactor"))
	return rv
}/* debug [instance_properties/getter]: hyphenationFactor */


// Returns the current hyphenation factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/hyphenationFactor
func (t_ Typesetter) SetHyphenationFactor(value float32) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHyphenationFactor:"), value)
}/* debug [instance_properties/setter]: hyphenationFactor */


// Returns the layout manager for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/layoutManager
func (t_ Typesetter) LayoutManager() ILayoutManager {
	rv := objc.Send[LayoutManager](t_.ID, objc.Sel("layoutManager"))
	return rv
}/* debug [instance_properties/getter]: layoutManager */


// Returns the current line fragment padding, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/lineFragmentPadding
func (t_ Typesetter) LineFragmentPadding() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("lineFragmentPadding"))
	return rv
}/* debug [instance_properties/getter]: lineFragmentPadding */


// Returns the current line fragment padding, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/lineFragmentPadding
func (t_ Typesetter) SetLineFragmentPadding(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineFragmentPadding:"), value)
}/* debug [instance_properties/setter]: lineFragmentPadding */


// Returns the character range currently being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphCharacterRange
func (t_ Typesetter) ParagraphCharacterRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("paragraphCharacterRange"))
	return rv
}/* debug [instance_properties/getter]: paragraphCharacterRange */


// Returns the glyph range currently being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphGlyphRange
func (t_ Typesetter) ParagraphGlyphRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("paragraphGlyphRange"))
	return rv
}/* debug [instance_properties/getter]: paragraphGlyphRange */


// Returns the current paragraph separator character range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphSeparatorCharacterRange
func (t_ Typesetter) ParagraphSeparatorCharacterRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("paragraphSeparatorCharacterRange"))
	return rv
}/* debug [instance_properties/getter]: paragraphSeparatorCharacterRange */


// Returns the current paragraph separator range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphSeparatorGlyphRange
func (t_ Typesetter) ParagraphSeparatorGlyphRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("paragraphSeparatorGlyphRange"))
	return rv
}/* debug [instance_properties/getter]: paragraphSeparatorGlyphRange */


// Returns a shared instance of a reentrant typesetter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/sharedSystemTypesetter
func (t_ Typesetter) SharedSystemTypesetter() ITypesetter {
	rv := objc.Send[Typesetter](t_.ID, objc.Sel("sharedSystemTypesetter"))
	return rv
}/* debug [instance_properties/getter]: sharedSystemTypesetter */


// Returns an array containing the text containers belonging to the current layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/textContainers
func (t_ Typesetter) TextContainers() []TextContainer {
	rv := objc.Send[[]TextContainer](t_.ID, objc.Sel("textContainers"))
	return rv
}/* debug [instance_properties/getter]: textContainers */


// Returns the current typesetter behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/typesetterBehavior
func (t_ Typesetter) TypesetterBehavior() TypesetterBehavior {
	rv := objc.Send[TypesetterBehavior](t_.ID, objc.Sel("typesetterBehavior"))
	return rv
}/* debug [instance_properties/getter]: typesetterBehavior */


// Returns the current typesetter behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/typesetterBehavior
func (t_ Typesetter) SetTypesetterBehavior(value TypesetterBehavior) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTypesetterBehavior:"), value)
}/* debug [instance_properties/setter]: typesetterBehavior */


// Returns whether the typesetter uses the leading (or line gap) value specified in the font metric information of the current font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/usesFontLeading
func (t_ Typesetter) UsesFontLeading() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontLeading"))
	return rv
}/* debug [instance_properties/getter]: usesFontLeading */


// Returns whether the typesetter uses the leading (or line gap) value specified in the font metric information of the current font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/usesFontLeading
func (t_ Typesetter) SetUsesFontLeading(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontLeading:"), value)
}/* debug [instance_properties/setter]: usesFontLeading */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTypesetter */



