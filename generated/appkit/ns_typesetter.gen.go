// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/coreml"
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
	HyphenationFactor() float32 /* primitive/slice/pointer. */
	SetHyphenationFactor(value float32 /* primitive/slice/pointer. */)
	LayoutManager() objc.IObject /* cross-framework: LayoutManager */
	LineFragmentPadding() float64 /* primitive/slice/pointer. */
	SetLineFragmentPadding(value float64 /* primitive/slice/pointer. */)
	ParagraphSeparatorCharacterRange() objc.IObject /* cross-framework: Range */
	ParagraphSeparatorGlyphRange() objc.IObject /* cross-framework: Range */
	TextContainers() []TextContainer /* primitive/slice/pointer. */
	UsesFontLeading() bool /* primitive/slice/pointer. */
	SetUsesFontLeading(value bool /* primitive/slice/pointer. */)
	AttributedString() objc.IObject /* cross-framework: AttributedString */
	SetAttributedString(value objc.IObject /* cross-framework: AttributedString */)
	AttributesForExtraLineFragment() objc.IObject /* cross-framework: Key */
	SetAttributesForExtraLineFragment(value objc.IObject /* cross-framework: Key */)
	BidiProcessingEnabled() bool /* primitive/slice/pointer. */
	SetBidiProcessingEnabled(value bool /* primitive/slice/pointer. */)
	CurrentParagraphStyle() IParagraphStyle
	SetCurrentParagraphStyle(value IParagraphStyle)
	CurrentTextContainer() ITextContainer
	SetCurrentTextContainer(value ITextContainer)
	ParagraphCharacterRange() objc.IObject /* cross-framework: Range */
	SetParagraphCharacterRange(value objc.IObject /* cross-framework: Range */)
	ParagraphGlyphRange() objc.IObject /* cross-framework: Range */
	SetParagraphGlyphRange(value objc.IObject /* cross-framework: Range */)
	TypesetterBehavior() TypesetterBehavior
	SetTypesetterBehavior(value TypesetterBehavior)
	// methods:
	ActionForControlCharacterAtIndex(charIndex uint /* primitive/slice/pointer. */) TypesetterControlCharacterAction /* not a class type */
	BaselineOffsetInLayoutManagerGlyphIndex(layoutMgr objc.IObject /* cross-framework LayoutManager */, glyphIndex uint /* primitive/slice/pointer. */) float64 /* primitive/slice/pointer. */
	BeginLineWithGlyphAtIndex(glyphIndex uint /* primitive/slice/pointer. */)
	BeginParagraph()
	BoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(glyphIndex uint /* primitive/slice/pointer. */, textContainer ITextContainer, proposedRect objc.IObject /* cross-framework Rect */, glyphPosition objc.IObject /* cross-framework Point */, charIndex uint /* primitive/slice/pointer. */) objc.IObject /* cross-framework: Rect */
	EndParagraph()
	GetLineFragmentRectUsedRectRemainingRectForStartingGlyphAtIndexProposedRectLineSpacingParagraphSpacingBeforeParagraphSpacingAfter(lineFragmentRect RectPointer /* not a class type */, lineFragmentUsedRect RectPointer /* not a class type */, remainingRect RectPointer /* not a class type */, startingGlyphIndex uint /* primitive/slice/pointer. */, proposedRect objc.IObject /* cross-framework Rect */, lineSpacing float64 /* primitive/slice/pointer. */, paragraphSpacingBefore float64 /* primitive/slice/pointer. */, paragraphSpacingAfter float64 /* primitive/slice/pointer. */)
	GlyphRangeForCharacterRangeActualCharacterRange(charRange objc.IObject /* cross-framework Range */, actualCharRange RangePointer /* not a class type */) objc.IObject /* cross-framework: Range */
	HyphenCharacterForGlyphAtIndex(glyphIndex uint /* primitive/slice/pointer. */) unsafe.Pointer
	HyphenationFactorForGlyphAtIndex(glyphIndex uint /* primitive/slice/pointer. */) float32 /* primitive/slice/pointer. */
	LayoutGlyphsInLayoutManagerStartingAtGlyphIndexMaxNumberOfLineFragmentsNextGlyphIndex(layoutManager objc.IObject /* cross-framework LayoutManager */, startGlyphIndex uint /* primitive/slice/pointer. */, maxNumLines uint /* primitive/slice/pointer. */, nextGlyph UInteger /* not a class type */)
	SetDrawsOutsideLineFragmentForGlyphRange(flag bool /* primitive/slice/pointer. */, glyphRange objc.IObject /* cross-framework Range */)
	SetHardInvalidationForGlyphRange(flag bool /* primitive/slice/pointer. */, glyphRange objc.IObject /* cross-framework Range */)
	SetLocationWithAdvancementsForStartOfGlyphRange(location objc.IObject /* cross-framework Point */, advancements corefoundation.CGFloat, glyphRange objc.IObject /* cross-framework Range */)
	SetNotShownAttributeForGlyphRange(flag bool /* primitive/slice/pointer. */, glyphRange objc.IObject /* cross-framework Range */)
	ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	SubstituteFontForFont(originalFont IFont) IFont
	WillSetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(lineRect RectPointer /* not a class type */, glyphRange objc.IObject /* cross-framework Range */, usedRect RectPointer /* not a class type */, baselineOffset corefoundation.CGFloat)
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



// Returns the interglyph spacing in the specified range when sent to a printer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/printingAdjustment(in:forNominallySpacedGlyphRange:packedGlyphs:count:)
func (tc _TypesetterClass) PrintingAdjustmentInLayoutManagerForNominallySpacedGlyphRangePackedGlyphsCount(layoutMgr objc.IObject /* cross-framework LayoutManager */, nominallySpacedGlyphsRange objc.IObject /* cross-framework Range */, packedGlyphs unsafe.Pointer, packedGlyphsCount uint /* primitive/slice/pointer. */) objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](objc.ID(tc.class), objc.Sel("printingAdjustmentInLayoutManager:forNominallySpacedGlyphRange:packedGlyphs:count:"), layoutMgr, nominallySpacedGlyphsRange, packedGlyphs, packedGlyphsCount)
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
func (t_ Typesetter) ActionForControlCharacterAtIndex(charIndex uint /* primitive/slice/pointer. */) TypesetterControlCharacterAction /* not a class type */ {
	rv := objc.Send[TypesetterControlCharacterAction](t_.ID, objc.Sel("actionForControlCharacterAtIndex:"), charIndex)
	return rv
}


// Returns the distance from the bottom of the line fragment rectangle in which the glyph resides to the glyph baseline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/baselineOffset(in:glyphIndex:)
func (t_ Typesetter) BaselineOffsetInLayoutManagerGlyphIndex(layoutMgr objc.IObject /* cross-framework LayoutManager */, glyphIndex uint /* primitive/slice/pointer. */) float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](t_.ID, objc.Sel("baselineOffsetInLayoutManager:glyphIndex:"), layoutMgr, glyphIndex)
	return rv
}


// Sets up layout parameters at the beginning of a line during typesetting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/beginLine(withGlyphAt:)
func (t_ Typesetter) BeginLineWithGlyphAtIndex(glyphIndex uint /* primitive/slice/pointer. */) {
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
func (t_ Typesetter) BoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(glyphIndex uint /* primitive/slice/pointer. */, textContainer ITextContainer, proposedRect objc.IObject /* cross-framework Rect */, glyphPosition objc.IObject /* cross-framework Point */, charIndex uint /* primitive/slice/pointer. */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[Rect](t_.ID, objc.Sel("boundingBoxForControlGlyphAtIndex:forTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), glyphIndex, textContainer, proposedRect, glyphPosition, charIndex)
	return rv
}


// Sets up layout parameters at the end of a paragraph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/endParagraph()
func (t_ Typesetter) EndParagraph() {
	objc.Send[objc.ID](t_.ID, objc.Sel("endParagraph"))
}


// Calculates line fragment rectangle, line fragment used rectangle, and remaining rectangle for a line fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/getLineFragmentRect(_:usedRect:remaining:forStartingGlyphAt:proposedRect:lineSpacing:paragraphSpacingBefore:paragraphSpacingAfter:)
func (t_ Typesetter) GetLineFragmentRectUsedRectRemainingRectForStartingGlyphAtIndexProposedRectLineSpacingParagraphSpacingBeforeParagraphSpacingAfter(lineFragmentRect RectPointer /* not a class type */, lineFragmentUsedRect RectPointer /* not a class type */, remainingRect RectPointer /* not a class type */, startingGlyphIndex uint /* primitive/slice/pointer. */, proposedRect objc.IObject /* cross-framework Rect */, lineSpacing float64 /* primitive/slice/pointer. */, paragraphSpacingBefore float64 /* primitive/slice/pointer. */, paragraphSpacingAfter float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("getLineFragmentRect:usedRect:remainingRect:forStartingGlyphAtIndex:proposedRect:lineSpacing:paragraphSpacingBefore:paragraphSpacingAfter:"), lineFragmentRect, lineFragmentUsedRect, remainingRect, startingGlyphIndex, proposedRect, lineSpacing, paragraphSpacingBefore, paragraphSpacingAfter)
}


// Returns the range for the glyphs mapped to the characters of the text store in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/glyphRange(forCharacterRange:actualCharacterRange:)
func (t_ Typesetter) GlyphRangeForCharacterRangeActualCharacterRange(charRange objc.IObject /* cross-framework Range */, actualCharRange RangePointer /* not a class type */) objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[Range](t_.ID, objc.Sel("glyphRangeForCharacterRange:actualCharacterRange:"), charRange, actualCharRange)
	return rv
}


// Returns the hyphen character to be inserted after the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/hyphenCharacter(forGlyphAt:)
func (t_ Typesetter) HyphenCharacterForGlyphAtIndex(glyphIndex uint /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("hyphenCharacterForGlyphAtIndex:"), glyphIndex)
	return rv
}


// Returns the hyphenation factor in effect at a specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/hyphenationFactor(forGlyphAt:)
func (t_ Typesetter) HyphenationFactorForGlyphAtIndex(glyphIndex uint /* primitive/slice/pointer. */) float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](t_.ID, objc.Sel("hyphenationFactorForGlyphAtIndex:"), glyphIndex)
	return rv
}


// Lays out glyphs in the specified layout manager starting at a specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/layoutGlyphs(in:startingAtGlyphIndex:maxNumberOfLineFragments:nextGlyphIndex:)
func (t_ Typesetter) LayoutGlyphsInLayoutManagerStartingAtGlyphIndexMaxNumberOfLineFragmentsNextGlyphIndex(layoutManager objc.IObject /* cross-framework LayoutManager */, startGlyphIndex uint /* primitive/slice/pointer. */, maxNumLines uint /* primitive/slice/pointer. */, nextGlyph UInteger /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("layoutGlyphsInLayoutManager:startingAtGlyphIndex:maxNumberOfLineFragments:nextGlyphIndex:"), layoutManager, startGlyphIndex, maxNumLines, nextGlyph)
}


// Sets whether the specified glyphs exceed the bounds of the line fragment in which they are laid out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setDrawsOutsideLineFragment(_:forGlyphRange:)
func (t_ Typesetter) SetDrawsOutsideLineFragmentForGlyphRange(flag bool /* primitive/slice/pointer. */, glyphRange objc.IObject /* cross-framework Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsOutsideLineFragment:forGlyphRange:"), flag, glyphRange)
}


// Sets whether to force the layout manager to invalidate the specified portion of the glyph cache when invalidating layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setHardInvalidation(_:forGlyphRange:)
func (t_ Typesetter) SetHardInvalidationForGlyphRange(flag bool /* primitive/slice/pointer. */, glyphRange objc.IObject /* cross-framework Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHardInvalidation:forGlyphRange:"), flag, glyphRange)
}


// Sets the location where the specified glyphs are laid out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setLocation(_:withAdvancements:forStartOfGlyphRange:)
func (t_ Typesetter) SetLocationWithAdvancementsForStartOfGlyphRange(location objc.IObject /* cross-framework Point */, advancements corefoundation.CGFloat, glyphRange objc.IObject /* cross-framework Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLocation:withAdvancements:forStartOfGlyphRange:"), location, advancements, glyphRange)
}


// Sets whether the specified glyphs are not shown.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/setNotShownAttribute(_:forGlyphRange:)
func (t_ Typesetter) SetNotShownAttributeForGlyphRange(flag bool /* primitive/slice/pointer. */, glyphRange objc.IObject /* cross-framework Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNotShownAttribute:forGlyphRange:"), flag, glyphRange)
}


// Returns whether the line being laid out should be broken by a word break at the specified character.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/shouldBreakLine(byWordBeforeCharacterAt:)
func (t_ Typesetter) ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
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


// Called by the typesetter just prior to storing the actual line fragment rectangle location in the layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/willSetLineFragmentRect(_:forGlyphRange:usedRect:baselineOffset:)
func (t_ Typesetter) WillSetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(lineRect RectPointer /* not a class type */, glyphRange objc.IObject /* cross-framework Range */, usedRect RectPointer /* not a class type */, baselineOffset corefoundation.CGFloat) {
	objc.Send[objc.ID](t_.ID, objc.Sel("willSetLineFragmentRect:forGlyphRange:usedRect:baselineOffset:"), lineRect, glyphRange, usedRect, baselineOffset)
}


// Returns the current hyphenation factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/hyphenationFactor
func (t_ Typesetter) HyphenationFactor() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](t_.ID, objc.Sel("hyphenationFactor"))
	return rv
}


// Returns the current hyphenation factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/hyphenationFactor
func (t_ Typesetter) SetHyphenationFactor(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setHyphenationFactor:"), value)
}


// Returns the layout manager for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/layoutManager
func (t_ Typesetter) LayoutManager() objc.IObject /* cross-framework: LayoutManager */ {
	rv := objc.Send[LayoutManager](t_.ID, objc.Sel("layoutManager"))
	return rv
}


// Returns the current line fragment padding, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/lineFragmentPadding
func (t_ Typesetter) LineFragmentPadding() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](t_.ID, objc.Sel("lineFragmentPadding"))
	return rv
}


// Returns the current line fragment padding, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/lineFragmentPadding
func (t_ Typesetter) SetLineFragmentPadding(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLineFragmentPadding:"), value)
}


// Returns the current paragraph separator character range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphSeparatorCharacterRange
func (t_ Typesetter) ParagraphSeparatorCharacterRange() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[Range](t_.ID, objc.Sel("paragraphSeparatorCharacterRange"))
	return rv
}


// Returns the current paragraph separator range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/paragraphSeparatorGlyphRange
func (t_ Typesetter) ParagraphSeparatorGlyphRange() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[Range](t_.ID, objc.Sel("paragraphSeparatorGlyphRange"))
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
func (t_ Typesetter) TextContainers() []TextContainer /* primitive/slice/pointer. */ {
	rv := objc.Send[[]TextContainer](t_.ID, objc.Sel("textContainers"))
	return rv
}


// Returns whether the typesetter uses the leading (or line gap) value specified in the font metric information of the current font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/usesFontLeading
func (t_ Typesetter) UsesFontLeading() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontLeading"))
	return rv
}


// Returns whether the typesetter uses the leading (or line gap) value specified in the font metric information of the current font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter/usesFontLeading
func (t_ Typesetter) SetUsesFontLeading(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontLeading:"), value)
}


// Returns the text backing store, usually an instance of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/attributedstring
func (t_ Typesetter) AttributedString() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[AttributedString](t_.ID, objc.Sel("attributedString"))
	return rv
}


// Returns the text backing store, usually an instance of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/attributedstring
func (t_ Typesetter) SetAttributedString(value objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttributedString:"), value)
}


// Returns the attributes used to lay out the extra line fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/attributesforextralinefragment
func (t_ Typesetter) AttributesForExtraLineFragment() objc.IObject /* cross-framework: Key */ {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("attributesForExtraLineFragment"))
	return rv
}


// Returns the attributes used to lay out the extra line fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/attributesforextralinefragment
func (t_ Typesetter) SetAttributesForExtraLineFragment(value objc.IObject /* cross-framework: Key */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAttributesForExtraLineFragment:"), value)
}


// Returns whether bidirectional text processing is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/bidiprocessingenabled
func (t_ Typesetter) BidiProcessingEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("bidiProcessingEnabled"))
	return rv
}


// Returns whether bidirectional text processing is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/bidiprocessingenabled
func (t_ Typesetter) SetBidiProcessingEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBidiProcessingEnabled:"), value)
}


// Returns the paragraph style object for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/currentparagraphstyle
func (t_ Typesetter) CurrentParagraphStyle() IParagraphStyle {
	rv := objc.Send[ParagraphStyle](t_.ID, objc.Sel("currentParagraphStyle"))
	return rv
}


// Returns the paragraph style object for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/currentparagraphstyle
func (t_ Typesetter) SetCurrentParagraphStyle(value IParagraphStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentParagraphStyle:"), value)
}


// Returns the text container for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/currenttextcontainer
func (t_ Typesetter) CurrentTextContainer() ITextContainer {
	rv := objc.Send[TextContainer](t_.ID, objc.Sel("currentTextContainer"))
	return rv
}


// Returns the text container for the text being typeset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/currenttextcontainer
func (t_ Typesetter) SetCurrentTextContainer(value ITextContainer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentTextContainer:"), value)
}


// Returns the character range currently being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphcharacterrange
func (t_ Typesetter) ParagraphCharacterRange() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[Range](t_.ID, objc.Sel("paragraphCharacterRange"))
	return rv
}


// Returns the character range currently being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphcharacterrange
func (t_ Typesetter) SetParagraphCharacterRange(value objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParagraphCharacterRange:"), value)
}


// Returns the glyph range currently being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphglyphrange
func (t_ Typesetter) ParagraphGlyphRange() objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[Range](t_.ID, objc.Sel("paragraphGlyphRange"))
	return rv
}


// Returns the glyph range currently being processed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/paragraphglyphrange
func (t_ Typesetter) SetParagraphGlyphRange(value objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParagraphGlyphRange:"), value)
}


// Returns the current typesetter behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/typesetterbehavior
func (t_ Typesetter) TypesetterBehavior() TypesetterBehavior {
	rv := objc.Send[TypesetterBehavior](t_.ID, objc.Sel("typesetterBehavior"))
	return rv
}


// Returns the current typesetter behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/typesetterbehavior
func (t_ Typesetter) SetTypesetterBehavior(value TypesetterBehavior) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTypesetterBehavior:"), value)
}



