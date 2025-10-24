// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LayoutManager] class.
var (
	LayoutManagerClass     _LayoutManagerClass
	LayoutManagerClassOnce sync.Once
)

func getLayoutManagerClass() _LayoutManagerClass {
	LayoutManagerClassOnce.Do(func() {
		LayoutManagerClass = _LayoutManagerClass{objc.GetClass("NSLayoutManager")}
	})
	return LayoutManagerClass
}

type _LayoutManagerClass struct {
	class objc.Class
}

// An interface definition for the [LayoutManager] class.
type ILayoutManager interface {
	objectivec.IObject
	// properties:
	AllowsNonContiguousLayout() bool
	SetAllowsNonContiguousLayout(value bool)
	BackgroundLayoutEnabled() bool
	SetBackgroundLayoutEnabled(value bool)
	DefaultAttachmentScaling() ImageScaling
	SetDefaultAttachmentScaling(value ImageScaling)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	ExtraLineFragmentRect() objc.IObject /* cross-framework: Rect */
	ExtraLineFragmentTextContainer() ITextContainer
	ExtraLineFragmentUsedRect() objc.IObject /* cross-framework: Rect */
	FirstTextView() ITextView
	GlyphGenerator() IGlyphGenerator
	SetGlyphGenerator(value IGlyphGenerator)
	HasNonContiguousLayout() bool
	HyphenationFactor() float32
	SetHyphenationFactor(value float32)
	LimitsLayoutForSuspiciousContents() bool
	SetLimitsLayoutForSuspiciousContents(value bool)
	NumberOfGlyphs() uint
	ShowsControlCharacters() bool
	SetShowsControlCharacters(value bool)
	ShowsInvisibleCharacters() bool
	SetShowsInvisibleCharacters(value bool)
	TextContainers() []TextContainer
	TextStorage() ITextStorage
	SetTextStorage(value ITextStorage)
	TextViewForBeginningOfSelection() ITextView
	Typesetter() ITypesetter
	SetTypesetter(value ITypesetter)
	TypesetterBehavior() TypesetterBehavior
	SetTypesetterBehavior(value TypesetterBehavior)
	UsesDefaultHyphenation() bool
	SetUsesDefaultHyphenation(value bool)
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
	UsesScreenFonts() bool
	SetUsesScreenFonts(value bool)
	// methods:
	AddTemporaryAttributeValueForCharacterRange(attrName AttributedStringKey /* not a class type */, value objc.IObject, charRange corefoundation.Range)
	AddTemporaryAttributesForCharacterRange(attrs foundation.IDictionary, charRange corefoundation.Range)
	AddTextContainer(container ITextContainer)
	AttachmentSizeForGlyphAtIndex(glyphIndex uint) objc.IObject /* cross-framework: Size */
	BoundingRectForGlyphRangeInTextContainer(glyphRange corefoundation.Range, container ITextContainer) objc.IObject /* cross-framework: Rect */
	BoundsRectForTextBlockAtIndexEffectiveRange(block ITextBlock, glyphIndex uint, effectiveGlyphRange RangePointer /* not a class type */) objc.IObject /* cross-framework: Rect */
	BoundsRectForTextBlockGlyphRange(block ITextBlock, glyphRange corefoundation.Range) objc.IObject /* cross-framework: Rect */
	CGGlyphAtIndex(glyphIndex uint) objc.IObject /* cross-framework: Glyph */
	CGGlyphAtIndexIsValidIndex(glyphIndex uint, isValidIndex unsafe.Pointer) objc.IObject /* cross-framework: Glyph */
	CharacterIndexForPointInTextContainerFractionOfDistanceBetweenInsertionPoints(point objc.IObject /* cross-framework: Point */, container ITextContainer, partialFraction corefoundation.CGFloat) uint
	CharacterIndexForGlyphAtIndex(glyphIndex uint) uint
	CharacterRangeForGlyphRangeActualGlyphRange(glyphRange corefoundation.Range, actualGlyphRange RangePointer /* not a class type */) corefoundation.Range
	DefaultBaselineOffsetForFont(theFont IFont) float64
	DefaultLineHeightForFont(theFont IFont) float64
	DrawBackgroundForGlyphRangeAtPoint(glyphsToShow corefoundation.Range, origin objc.IObject /* cross-framework: Point */)
	DrawGlyphsForGlyphRangeAtPoint(glyphsToShow corefoundation.Range, origin objc.IObject /* cross-framework: Point */)
	DrawStrikethroughForGlyphRangeStrikethroughTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange corefoundation.Range, strikethroughVal UnderlineStyle, baselineOffset float64, lineRect objc.IObject /* cross-framework: Rect */, lineGlyphRange corefoundation.Range, containerOrigin objc.IObject /* cross-framework: Point */)
	DrawUnderlineForGlyphRangeUnderlineTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange corefoundation.Range, underlineVal UnderlineStyle, baselineOffset float64, lineRect objc.IObject /* cross-framework: Rect */, lineGlyphRange corefoundation.Range, containerOrigin objc.IObject /* cross-framework: Point */)
	DrawsOutsideLineFragmentForGlyphAtIndex(glyphIndex uint) bool
	EnsureGlyphsForCharacterRange(charRange corefoundation.Range)
	EnsureGlyphsForGlyphRange(glyphRange corefoundation.Range)
	EnsureLayoutForTextContainer(container ITextContainer)
	EnsureLayoutForBoundingRectInTextContainer(bounds objc.IObject /* cross-framework: Rect */, container ITextContainer)
	EnsureLayoutForCharacterRange(charRange corefoundation.Range)
	EnsureLayoutForGlyphRange(glyphRange corefoundation.Range)
	EnumerateEnclosingRectsForGlyphRangeWithinSelectedGlyphRangeInTextContainerUsingBlock(glyphRange corefoundation.Range, selectedRange corefoundation.Range, textContainer ITextContainer, block unsafe.Pointer)
	EnumerateLineFragmentsForGlyphRangeUsingBlock(glyphRange corefoundation.Range, block unsafe.Pointer)
	FillBackgroundRectArrayCountForCharacterRangeColor(rectArray objc.IObject /* cross-framework: Rect */, rectCount uint, charRange corefoundation.Range, color IColor)
	FirstUnlaidCharacterIndex() uint
	FirstUnlaidGlyphIndex() uint
	FractionOfDistanceThroughGlyphForPointInTextContainer(point objc.IObject /* cross-framework: Point */, container ITextContainer) float64
	GetFirstUnlaidCharacterIndexGlyphIndex(charIndex uint, glyphIndex uint)
	GetGlyphsInRangeGlyphsPropertiesCharacterIndexesBidiLevels(glyphRange corefoundation.Range, glyphBuffer objc.IObject /* cross-framework: Glyph */, props GlyphProperty, charIndexBuffer uint, bidiLevelBuffer unsafe.Pointer) uint
	GetLineFragmentInsertionPointsForCharacterAtIndexAlternatePositionsInDisplayOrderPositionsCharacterIndexes(charIndex uint, aFlag bool, dFlag bool, positions corefoundation.CGFloat, charIndexes uint) uint
	GlyphAtIndex(glyphIndex uint) objc.IObject /* cross-framework: Glyph */
	GlyphAtIndexIsValidIndex(glyphIndex uint, isValidIndex unsafe.Pointer) objc.IObject /* cross-framework: Glyph */
	GlyphIndexForPointInTextContainer(point objc.IObject /* cross-framework: Point */, container ITextContainer) uint
	GlyphIndexForPointInTextContainerFractionOfDistanceThroughGlyph(point objc.IObject /* cross-framework: Point */, container ITextContainer, partialFraction corefoundation.CGFloat) uint
	GlyphIndexForCharacterAtIndex(charIndex uint) uint
	GlyphRangeForTextContainer(container ITextContainer) corefoundation.Range
	GlyphRangeForBoundingRectInTextContainer(bounds objc.IObject /* cross-framework: Rect */, container ITextContainer) corefoundation.Range
	GlyphRangeForBoundingRectWithoutAdditionalLayoutInTextContainer(bounds objc.IObject /* cross-framework: Rect */, container ITextContainer) corefoundation.Range
	GlyphRangeForCharacterRangeActualCharacterRange(charRange corefoundation.Range, actualCharRange RangePointer /* not a class type */) corefoundation.Range
	InsertTextContainerAtIndex(container ITextContainer, index uint)
	InvalidateDisplayForCharacterRange(charRange corefoundation.Range)
	InvalidateDisplayForGlyphRange(glyphRange corefoundation.Range)
	InvalidateGlyphsForCharacterRangeChangeInLengthActualCharacterRange(charRange corefoundation.Range, delta int, actualCharRange RangePointer /* not a class type */)
	InvalidateLayoutForCharacterRangeActualCharacterRange(charRange corefoundation.Range, actualCharRange RangePointer /* not a class type */)
	IsValidGlyphIndex(glyphIndex uint) bool
	LayoutManagerOwnsFirstResponderInWindow(window IWindow) bool
	LayoutRectForTextBlockAtIndexEffectiveRange(block ITextBlock, glyphIndex uint, effectiveGlyphRange RangePointer /* not a class type */) objc.IObject /* cross-framework: Rect */
	LayoutRectForTextBlockGlyphRange(block ITextBlock, glyphRange corefoundation.Range) objc.IObject /* cross-framework: Rect */
	LineFragmentRectForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange RangePointer /* not a class type */) objc.IObject /* cross-framework: Rect */
	LineFragmentRectForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange RangePointer /* not a class type */, flag bool) objc.IObject /* cross-framework: Rect */
	LineFragmentUsedRectForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange RangePointer /* not a class type */) objc.IObject /* cross-framework: Rect */
	LineFragmentUsedRectForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange RangePointer /* not a class type */, flag bool) objc.IObject /* cross-framework: Rect */
	LocationForGlyphAtIndex(glyphIndex uint) objc.IObject /* cross-framework: Point */
	NotShownAttributeForGlyphAtIndex(glyphIndex uint) bool
	ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange(textStorage ITextStorage, editMask TextStorageEditActions, newCharRange corefoundation.Range, delta int, invalidatedCharRange corefoundation.Range)
	PropertyForGlyphAtIndex(glyphIndex uint) GlyphProperty
	RangeOfNominallySpacedGlyphsContainingIndex(glyphIndex uint) corefoundation.Range
	RectArrayForCharacterRangeWithinSelectedCharacterRangeInTextContainerRectCount(charRange corefoundation.Range, selCharRange corefoundation.Range, container ITextContainer, rectCount uint) RectArray /* not a class type */
	RectArrayForGlyphRangeWithinSelectedGlyphRangeInTextContainerRectCount(glyphRange corefoundation.Range, selGlyphRange corefoundation.Range, container ITextContainer, rectCount uint) RectArray /* not a class type */
	RemoveTemporaryAttributeForCharacterRange(attrName AttributedStringKey /* not a class type */, charRange corefoundation.Range)
	RemoveTextContainerAtIndex(index uint)
	ReplaceTextStorage(newTextStorage ITextStorage)
	RulerAccessoryViewForTextViewParagraphStyleRulerEnabled(view ITextView, style IParagraphStyle, ruler IRulerView, isEnabled bool) IView
	RulerMarkersForTextViewParagraphStyleRuler(view ITextView, style IParagraphStyle, ruler IRulerView) []RulerMarker
	SetAttachmentSizeForGlyphRange(attachmentSize objc.IObject /* cross-framework: Size */, glyphRange corefoundation.Range)
	SetBoundsRectForTextBlockGlyphRange(rect objc.IObject /* cross-framework: Rect */, block ITextBlock, glyphRange corefoundation.Range)
	SetDrawsOutsideLineFragmentForGlyphAtIndex(flag bool, glyphIndex uint)
	SetExtraLineFragmentRectUsedRectTextContainer(fragmentRect objc.IObject /* cross-framework: Rect */, usedRect objc.IObject /* cross-framework: Rect */, container ITextContainer)
	SetGlyphsPropertiesCharacterIndexesFontForGlyphRange(glyphs objc.IObject /* cross-framework: Glyph */, props GlyphProperty, charIndexes uint, aFont IFont, glyphRange corefoundation.Range)
	SetLayoutRectForTextBlockGlyphRange(rect objc.IObject /* cross-framework: Rect */, block ITextBlock, glyphRange corefoundation.Range)
	SetLineFragmentRectForGlyphRangeUsedRect(fragmentRect objc.IObject /* cross-framework: Rect */, glyphRange corefoundation.Range, usedRect objc.IObject /* cross-framework: Rect */)
	SetLocationForStartOfGlyphRange(location objc.IObject /* cross-framework: Point */, glyphRange corefoundation.Range)
	SetNotShownAttributeForGlyphAtIndex(flag bool, glyphIndex uint)
	SetTemporaryAttributesForCharacterRange(attrs foundation.IDictionary, charRange corefoundation.Range)
	SetTextContainerForGlyphRange(container ITextContainer, glyphRange corefoundation.Range)
	ShowAttachmentCellInRectCharacterIndex(cell ICell, rect objc.IObject /* cross-framework: Rect */, attachmentIndex uint)
	ShowCGGlyphsPositionsCountFontTextMatrixAttributesInContext(glyphs objc.IObject /* cross-framework: Glyph */, positions corefoundation.CGPoint, glyphCount int, font IFont, textMatrix corefoundation.CGAffineTransform, attributes foundation.IDictionary, CGContext ContextRef /* not a class type */)
	StrikethroughGlyphRangeStrikethroughTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange corefoundation.Range, strikethroughVal UnderlineStyle, lineRect objc.IObject /* cross-framework: Rect */, lineGlyphRange corefoundation.Range, containerOrigin objc.IObject /* cross-framework: Point */)
	TemporaryAttributeAtCharacterIndexEffectiveRange(attrName AttributedStringKey /* not a class type */, location uint, range_ RangePointer /* not a class type */) objc.ID
	TemporaryAttributeAtCharacterIndexLongestEffectiveRangeInRange(attrName AttributedStringKey /* not a class type */, location uint, range_ RangePointer /* not a class type */, rangeLimit corefoundation.Range) objc.ID
	TemporaryAttributesAtCharacterIndexEffectiveRange(charIndex uint, effectiveCharRange RangePointer /* not a class type */) foundation.IDictionary
	TemporaryAttributesAtCharacterIndexLongestEffectiveRangeInRange(location uint, range_ RangePointer /* not a class type */, rangeLimit corefoundation.Range) foundation.IDictionary
	TextContainerForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange RangePointer /* not a class type */) ITextContainer
	TextContainerForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange RangePointer /* not a class type */, flag bool) ITextContainer
	TextContainerChangedGeometry(container ITextContainer)
	TextContainerChangedTextView(container ITextContainer)
	TruncatedGlyphRangeInLineFragmentForGlyphAtIndex(glyphIndex uint) corefoundation.Range
	UnderlineGlyphRangeUnderlineTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange corefoundation.Range, underlineVal UnderlineStyle, lineRect objc.IObject /* cross-framework: Rect */, lineGlyphRange corefoundation.Range, containerOrigin objc.IObject /* cross-framework: Point */)
	UsedRectForTextContainer(container ITextContainer) objc.IObject /* cross-framework: Rect */
}

// An object that coordinates the layout and display of text characters.
//
// maps Unicode character codes to glyphs, sets the glyphs in a series of objects, and displays them in a series of objects. In addition to its core function of laying out text, a layout manager object coordinates its text view objects, provides services to those text views to support instances for editing paragraph styles, and handles the layout and display of text attributes not inherent in glyphs (such as underline or strikethrough). You can create a subclass of to handle additional text attributes, whether inherent or not.


// An object that coordinates the layout and display of text characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager
type LayoutManager struct {
	objectivec.Object
}

// LayoutManagerFrom constructs a [LayoutManager] from an unsafe.Pointer.
//
// An object that coordinates the layout and display of text characters.
func LayoutManagerFrom(ptr unsafe.Pointer) LayoutManager {
	return LayoutManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LayoutManagerClass) Alloc() LayoutManager {
	rv := objc.Send[LayoutManager](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LayoutManagerClass) New() LayoutManager {
	rv := objc.Send[LayoutManager](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LayoutManager) Init() LayoutManager {
	rv := objc.Send[LayoutManager](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LayoutManager) Autorelease() LayoutManager {
	rv := objc.Send[LayoutManager](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayoutManager creates a new LayoutManager instance.
func NewLayoutManager() LayoutManager {
	return getLayoutManagerClass().New()
}



// Creates a layout manager from data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/init(coder:)
func NewLayoutManagerWithCoder(coder foundation.Coder) LayoutManager {
	instance := getLayoutManagerClass().Alloc()
	rv := objc.Send[LayoutManager](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Adds a temporary attribute to the characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/addTemporaryAttribute(_:value:forCharacterRange:)
func (l_ LayoutManager) AddTemporaryAttributeValueForCharacterRange(attrName AttributedStringKey /* not a class type */, value objc.IObject, charRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("addTemporaryAttribute:value:forCharacterRange:"), attrName, value, charRange)
}


// Appends one or more temporary attributes to the attributes dictionary of the specified character range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/addTemporaryAttributes(_:forCharacterRange:)
func (l_ LayoutManager) AddTemporaryAttributesForCharacterRange(attrs foundation.IDictionary, charRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("addTemporaryAttributes:forCharacterRange:"), attrs, charRange)
}


// Appends the specified text container to the series of text containers where the layout manager arranges text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/addTextContainer(_:)
func (l_ LayoutManager) AddTextContainer(container ITextContainer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("addTextContainer:"), container)
}


// Returns the size of the attachment glyph at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/attachmentSize(forGlyphAt:)
func (l_ LayoutManager) AttachmentSizeForGlyphAtIndex(glyphIndex uint) objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](l_.ID, objc.Sel("attachmentSizeForGlyphAtIndex:"), glyphIndex)
	return rv
}


// Returns the bounding rectangle for the specified glyphs in a container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/boundingRect(forGlyphRange:in:)
func (l_ LayoutManager) BoundingRectForGlyphRangeInTextContainer(glyphRange corefoundation.Range, container ITextContainer) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](l_.ID, objc.Sel("boundingRectForGlyphRange:inTextContainer:"), glyphRange, container)
	return rv
}


// Returns the bounding rectangle for the specified text block and glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/boundsRect(for:at:effectiveRange:)
func (l_ LayoutManager) BoundsRectForTextBlockAtIndexEffectiveRange(block ITextBlock, glyphIndex uint, effectiveGlyphRange RangePointer /* not a class type */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](l_.ID, objc.Sel("boundsRectForTextBlock:atIndex:effectiveRange:"), block, glyphIndex, effectiveGlyphRange)
	return rv
}


// Returns the bounding rectangle that encloses the specified text block and glyph range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/boundsRect(for:glyphRange:)
func (l_ LayoutManager) BoundsRectForTextBlockGlyphRange(block ITextBlock, glyphRange corefoundation.Range) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](l_.ID, objc.Sel("boundsRectForTextBlock:glyphRange:"), block, glyphRange)
	return rv
}


// Returns the glyph at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/cgGlyph(at:)
func (l_ LayoutManager) CGGlyphAtIndex(glyphIndex uint) objc.IObject /* cross-framework: Glyph */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("CGGlyphAtIndex:"), glyphIndex)
	return rv
}


// Returns the glyph at the specified index along with information about whether the glyph index is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/cgGlyph(at:isValidIndex:)
func (l_ LayoutManager) CGGlyphAtIndexIsValidIndex(glyphIndex uint, isValidIndex unsafe.Pointer) objc.IObject /* cross-framework: Glyph */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("CGGlyphAtIndex:isValidIndex:"), glyphIndex, isValidIndex)
	return rv
}


// Returns the index of the character that lies beneath the specified point using the specified container’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/characterIndex(for:in:fractionOfDistanceBetweenInsertionPoints:)
func (l_ LayoutManager) CharacterIndexForPointInTextContainerFractionOfDistanceBetweenInsertionPoints(point objc.IObject /* cross-framework: Point */, container ITextContainer, partialFraction corefoundation.CGFloat) uint {
	rv := objc.Send[uint](l_.ID, objc.Sel("characterIndexForPoint:inTextContainer:fractionOfDistanceBetweenInsertionPoints:"), point, container, partialFraction)
	return rv
}


// Returns the index in the text storage for the first character of the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/characterIndexForGlyph(at:)
func (l_ LayoutManager) CharacterIndexForGlyphAtIndex(glyphIndex uint) uint {
	rv := objc.Send[uint](l_.ID, objc.Sel("characterIndexForGlyphAtIndex:"), glyphIndex)
	return rv
}


// Returns the range of characters that correspond to the glyphs in the specified glyph range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/characterRange(forGlyphRange:actualGlyphRange:)
func (l_ LayoutManager) CharacterRangeForGlyphRangeActualGlyphRange(glyphRange corefoundation.Range, actualGlyphRange RangePointer /* not a class type */) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](l_.ID, objc.Sel("characterRangeForGlyphRange:actualGlyphRange:"), glyphRange, actualGlyphRange)
	return rv
}


// Returns the default baseline offset that the layout manager’s typesetter uses for the specified font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/defaultBaselineOffset(for:)
func (l_ LayoutManager) DefaultBaselineOffsetForFont(theFont IFont) float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("defaultBaselineOffsetForFont:"), theFont)
	return rv
}


// Returns the default line height for a line of text that uses a specified font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/defaultLineHeight(for:)
func (l_ LayoutManager) DefaultLineHeightForFont(theFont IFont) float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("defaultLineHeightForFont:"), theFont)
	return rv
}


// Draws background marks for the specified glyphs, which must lie completely within a single text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/drawBackground(forGlyphRange:at:)
func (l_ LayoutManager) DrawBackgroundForGlyphRangeAtPoint(glyphsToShow corefoundation.Range, origin objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("drawBackgroundForGlyphRange:atPoint:"), glyphsToShow, origin)
}


// Draws the specified glyphs, which must lie completely within a single text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/drawGlyphs(forGlyphRange:at:)
func (l_ LayoutManager) DrawGlyphsForGlyphRangeAtPoint(glyphsToShow corefoundation.Range, origin objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("drawGlyphsForGlyphRange:atPoint:"), glyphsToShow, origin)
}


// Draws a strikethrough for the specified glyphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/drawStrikethrough(forGlyphRange:strikethroughType:baselineOffset:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:)
func (l_ LayoutManager) DrawStrikethroughForGlyphRangeStrikethroughTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange corefoundation.Range, strikethroughVal UnderlineStyle, baselineOffset float64, lineRect objc.IObject /* cross-framework: Rect */, lineGlyphRange corefoundation.Range, containerOrigin objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("drawStrikethroughForGlyphRange:strikethroughType:baselineOffset:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, strikethroughVal, baselineOffset, lineRect, lineGlyphRange, containerOrigin)
}


// Draws underlining for the glyphs in a specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/drawUnderline(forGlyphRange:underlineType:baselineOffset:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:)
func (l_ LayoutManager) DrawUnderlineForGlyphRangeUnderlineTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange corefoundation.Range, underlineVal UnderlineStyle, baselineOffset float64, lineRect objc.IObject /* cross-framework: Rect */, lineGlyphRange corefoundation.Range, containerOrigin objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("drawUnderlineForGlyphRange:underlineType:baselineOffset:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, underlineVal, baselineOffset, lineRect, lineGlyphRange, containerOrigin)
}


// Indicates whether the glyph draws outside its line fragment rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/drawsOutsideLineFragment(forGlyphAt:)
func (l_ LayoutManager) DrawsOutsideLineFragmentForGlyphAtIndex(glyphIndex uint) bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("drawsOutsideLineFragmentForGlyphAtIndex:"), glyphIndex)
	return rv
}


// Forces the layout manager to generate glyphs for the specified character range if it hasn’t already.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ensureGlyphs(forCharacterRange:)
func (l_ LayoutManager) EnsureGlyphsForCharacterRange(charRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("ensureGlyphsForCharacterRange:"), charRange)
}


// Forces the layout manager to generate glyphs for the specified glyph range if it hasn’t already.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ensureGlyphs(forGlyphRange:)
func (l_ LayoutManager) EnsureGlyphsForGlyphRange(glyphRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("ensureGlyphsForGlyphRange:"), glyphRange)
}


// Forces the layout manager to perform layout for the specified text container if it hasn’t already.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ensureLayout(for:)
func (l_ LayoutManager) EnsureLayoutForTextContainer(container ITextContainer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("ensureLayoutForTextContainer:"), container)
}


// Forces the layout manager to perform layout for the specified area in the specified text container if it hasn’t already.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ensureLayout(forBoundingRect:in:)
func (l_ LayoutManager) EnsureLayoutForBoundingRectInTextContainer(bounds objc.IObject /* cross-framework: Rect */, container ITextContainer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("ensureLayoutForBoundingRect:inTextContainer:"), bounds, container)
}


// Forces the layout manager to perform layout for the specified character range if it hasn’t already.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ensureLayout(forCharacterRange:)
func (l_ LayoutManager) EnsureLayoutForCharacterRange(charRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("ensureLayoutForCharacterRange:"), charRange)
}


// Forces the layout manager to perform layout for the specified glyph range if it hasn’t already.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ensureLayout(forGlyphRange:)
func (l_ LayoutManager) EnsureLayoutForGlyphRange(glyphRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("ensureLayoutForGlyphRange:"), glyphRange)
}


// Enumerates enclosing rectangles for the specified glyph range in a text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/enumerateEnclosingRects(forGlyphRange:withinSelectedGlyphRange:in:using:)
func (l_ LayoutManager) EnumerateEnclosingRectsForGlyphRangeWithinSelectedGlyphRangeInTextContainerUsingBlock(glyphRange corefoundation.Range, selectedRange corefoundation.Range, textContainer ITextContainer, block unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("enumerateEnclosingRectsForGlyphRange:withinSelectedGlyphRange:inTextContainer:usingBlock:"), glyphRange, selectedRange, textContainer, block)
}


// Enumerates line fragments intersecting with the specified glyph range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/enumerateLineFragments(forGlyphRange:using:)
func (l_ LayoutManager) EnumerateLineFragmentsForGlyphRangeUsingBlock(glyphRange corefoundation.Range, block unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("enumerateLineFragmentsForGlyphRange:usingBlock:"), glyphRange, block)
}


// Fills background rectangles with a color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/fillBackgroundRectArray(_:count:forCharacterRange:color:)
func (l_ LayoutManager) FillBackgroundRectArrayCountForCharacterRangeColor(rectArray objc.IObject /* cross-framework: Rect */, rectCount uint, charRange corefoundation.Range, color IColor) {
	objc.Send[objc.ID](l_.ID, objc.Sel("fillBackgroundRectArray:count:forCharacterRange:color:"), rectArray, rectCount, charRange, color)
}


// Returns the index for the first character in the layout manager that isn’t in the layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/firstUnlaidCharacterIndex()
func (l_ LayoutManager) FirstUnlaidCharacterIndex() uint {
	rv := objc.Send[uint](l_.ID, objc.Sel("firstUnlaidCharacterIndex"))
	return rv
}


// Returns the index for the first glyph in the layout manager that isn’t in the layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/firstUnlaidGlyphIndex()
func (l_ LayoutManager) FirstUnlaidGlyphIndex() uint {
	rv := objc.Send[uint](l_.ID, objc.Sel("firstUnlaidGlyphIndex"))
	return rv
}


// Returns the fraction of the distance between the glyph at the specified point and the next glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/fractionOfDistanceThroughGlyph(for:in:)
func (l_ LayoutManager) FractionOfDistanceThroughGlyphForPointInTextContainer(point objc.IObject /* cross-framework: Point */, container ITextContainer) float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("fractionOfDistanceThroughGlyphForPoint:inTextContainer:"), point, container)
	return rv
}


// Returns the indexes for the first character and glyph that have invalid layout information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/getFirstUnlaidCharacterIndex(_:glyphIndex:)
func (l_ LayoutManager) GetFirstUnlaidCharacterIndexGlyphIndex(charIndex uint, glyphIndex uint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("getFirstUnlaidCharacterIndex:glyphIndex:"), charIndex, glyphIndex)
}


// Fills a passed-in buffer with a sequence of glyphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/getGlyphs(in:glyphs:properties:characterIndexes:bidiLevels:)
func (l_ LayoutManager) GetGlyphsInRangeGlyphsPropertiesCharacterIndexesBidiLevels(glyphRange corefoundation.Range, glyphBuffer objc.IObject /* cross-framework: Glyph */, props GlyphProperty, charIndexBuffer uint, bidiLevelBuffer unsafe.Pointer) uint {
	rv := objc.Send[uint](l_.ID, objc.Sel("getGlyphsInRange:glyphs:properties:characterIndexes:bidiLevels:"), glyphRange, glyphBuffer, props, charIndexBuffer, bidiLevelBuffer)
	return rv
}


// Returns insertion points in bulk for a specified line fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/getLineFragmentInsertionPoints(forCharacterAt:alternatePositions:inDisplayOrder:positions:characterIndexes:)
func (l_ LayoutManager) GetLineFragmentInsertionPointsForCharacterAtIndexAlternatePositionsInDisplayOrderPositionsCharacterIndexes(charIndex uint, aFlag bool, dFlag bool, positions corefoundation.CGFloat, charIndexes uint) uint {
	rv := objc.Send[uint](l_.ID, objc.Sel("getLineFragmentInsertionPointsForCharacterAtIndex:alternatePositions:inDisplayOrder:positions:characterIndexes:"), charIndex, aFlag, dFlag, positions, charIndexes)
	return rv
}


// Returns the glyph at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyph(at:)
func (l_ LayoutManager) GlyphAtIndex(glyphIndex uint) objc.IObject /* cross-framework: Glyph */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("glyphAtIndex:"), glyphIndex)
	return rv
}


// Returns the glyph at a specified index, and optionally returns a flag indicating whether the requested index is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyph(at:isValidIndex:)
func (l_ LayoutManager) GlyphAtIndexIsValidIndex(glyphIndex uint, isValidIndex unsafe.Pointer) objc.IObject /* cross-framework: Glyph */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("glyphAtIndex:isValidIndex:"), glyphIndex, isValidIndex)
	return rv
}


// Returns the index of the glyph at the specified location in a text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphIndex(for:in:)
func (l_ LayoutManager) GlyphIndexForPointInTextContainer(point objc.IObject /* cross-framework: Point */, container ITextContainer) uint {
	rv := objc.Send[uint](l_.ID, objc.Sel("glyphIndexForPoint:inTextContainer:"), point, container)
	return rv
}


// Returns the index of the glyph at the specified point using the container’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphIndex(for:in:fractionOfDistanceThroughGlyph:)
func (l_ LayoutManager) GlyphIndexForPointInTextContainerFractionOfDistanceThroughGlyph(point objc.IObject /* cross-framework: Point */, container ITextContainer, partialFraction corefoundation.CGFloat) uint {
	rv := objc.Send[uint](l_.ID, objc.Sel("glyphIndexForPoint:inTextContainer:fractionOfDistanceThroughGlyph:"), point, container, partialFraction)
	return rv
}


// Returns the index of the first glyph of the character at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphIndexForCharacter(at:)
func (l_ LayoutManager) GlyphIndexForCharacterAtIndex(charIndex uint) uint {
	rv := objc.Send[uint](l_.ID, objc.Sel("glyphIndexForCharacterAtIndex:"), charIndex)
	return rv
}


// Returns the range of glyphs lying within the specified text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphRange(for:)
func (l_ LayoutManager) GlyphRangeForTextContainer(container ITextContainer) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](l_.ID, objc.Sel("glyphRangeForTextContainer:"), container)
	return rv
}


// Returns the smallest contiguous range for glyphs lying wholly or partially within the specified rectangle of the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphRange(forBoundingRect:in:)
func (l_ LayoutManager) GlyphRangeForBoundingRectInTextContainer(bounds objc.IObject /* cross-framework: Rect */, container ITextContainer) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](l_.ID, objc.Sel("glyphRangeForBoundingRect:inTextContainer:"), bounds, container)
	return rv
}


// Returns the smallest contiguous range for glyphs lying wholly or partially within the specified rectangle of the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphRange(forBoundingRectWithoutAdditionalLayout:in:)
func (l_ LayoutManager) GlyphRangeForBoundingRectWithoutAdditionalLayoutInTextContainer(bounds objc.IObject /* cross-framework: Rect */, container ITextContainer) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](l_.ID, objc.Sel("glyphRangeForBoundingRectWithoutAdditionalLayout:inTextContainer:"), bounds, container)
	return rv
}


// Returns the range of glyphs that the specified range of characters generates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphRange(forCharacterRange:actualCharacterRange:)
func (l_ LayoutManager) GlyphRangeForCharacterRangeActualCharacterRange(charRange corefoundation.Range, actualCharRange RangePointer /* not a class type */) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](l_.ID, objc.Sel("glyphRangeForCharacterRange:actualCharacterRange:"), charRange, actualCharRange)
	return rv
}


// Inserts a text container at the specified index in the list of text containers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/insertTextContainer(_:at:)
func (l_ LayoutManager) InsertTextContainerAtIndex(container ITextContainer, index uint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("insertTextContainer:atIndex:"), container, index)
}


// Invalidates display for the specified character range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/invalidateDisplay(forCharacterRange:)
func (l_ LayoutManager) InvalidateDisplayForCharacterRange(charRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("invalidateDisplayForCharacterRange:"), charRange)
}


// Invalidates a range of glyphs, requiring new layout information, and updates the appropriate regions of any text views that display those glyphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/invalidateDisplay(forGlyphRange:)
func (l_ LayoutManager) InvalidateDisplayForGlyphRange(glyphRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("invalidateDisplayForGlyphRange:"), glyphRange)
}


// Invalidates and adjusts the glyphs in the specified character range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/invalidateGlyphs(forCharacterRange:changeInLength:actualCharacterRange:)
func (l_ LayoutManager) InvalidateGlyphsForCharacterRangeChangeInLengthActualCharacterRange(charRange corefoundation.Range, delta int, actualCharRange RangePointer /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("invalidateGlyphsForCharacterRange:changeInLength:actualCharacterRange:"), charRange, delta, actualCharRange)
}


// Invalidates the layout information for the glyphs that map to the specified character range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/invalidateLayout(forCharacterRange:actualCharacterRange:)
func (l_ LayoutManager) InvalidateLayoutForCharacterRangeActualCharacterRange(charRange corefoundation.Range, actualCharRange RangePointer /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("invalidateLayoutForCharacterRange:actualCharacterRange:"), charRange, actualCharRange)
}


// Indicates whether the specified index refers to a valid glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/isValidGlyphIndex(_:)
func (l_ LayoutManager) IsValidGlyphIndex(glyphIndex uint) bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isValidGlyphIndex:"), glyphIndex)
	return rv
}


// Indicates whether the first responder in the specified window is a text view for the layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/layoutManagerOwnsFirstResponder(in:)
func (l_ LayoutManager) LayoutManagerOwnsFirstResponderInWindow(window IWindow) bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("layoutManagerOwnsFirstResponderInWindow:"), window)
	return rv
}


// Returns the rectangle for the layout of the specified text block and glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/layoutRect(for:at:effectiveRange:)
func (l_ LayoutManager) LayoutRectForTextBlockAtIndexEffectiveRange(block ITextBlock, glyphIndex uint, effectiveGlyphRange RangePointer /* not a class type */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](l_.ID, objc.Sel("layoutRectForTextBlock:atIndex:effectiveRange:"), block, glyphIndex, effectiveGlyphRange)
	return rv
}


// Returns the rectangle for the layout of the specified text block and glyph range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/layoutRect(for:glyphRange:)
func (l_ LayoutManager) LayoutRectForTextBlockGlyphRange(block ITextBlock, glyphRange corefoundation.Range) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](l_.ID, objc.Sel("layoutRectForTextBlock:glyphRange:"), block, glyphRange)
	return rv
}


// Returns the rectangle for the line fragment where the glyph lies and (optionally), by reference, the entire range of glyphs in that fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/lineFragmentRect(forGlyphAt:effectiveRange:)
func (l_ LayoutManager) LineFragmentRectForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange RangePointer /* not a class type */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](l_.ID, objc.Sel("lineFragmentRectForGlyphAtIndex:effectiveRange:"), glyphIndex, effectiveGlyphRange)
	return rv
}


// Returns the line fragment rectangle that contains the glyph at the specified glyph index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/lineFragmentRect(forGlyphAt:effectiveRange:withoutAdditionalLayout:)
func (l_ LayoutManager) LineFragmentRectForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange RangePointer /* not a class type */, flag bool) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](l_.ID, objc.Sel("lineFragmentRectForGlyphAtIndex:effectiveRange:withoutAdditionalLayout:"), glyphIndex, effectiveGlyphRange, flag)
	return rv
}


// Returns the usage rectangle for the line fragment and (optionally) returns the entire range of glyphs in that fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/lineFragmentUsedRect(forGlyphAt:effectiveRange:)
func (l_ LayoutManager) LineFragmentUsedRectForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange RangePointer /* not a class type */) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](l_.ID, objc.Sel("lineFragmentUsedRectForGlyphAtIndex:effectiveRange:"), glyphIndex, effectiveGlyphRange)
	return rv
}


// Returns the usage rectangle for the line fragment and (optionally) returns the entire range of glyphs in that fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/lineFragmentUsedRect(forGlyphAt:effectiveRange:withoutAdditionalLayout:)
func (l_ LayoutManager) LineFragmentUsedRectForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange RangePointer /* not a class type */, flag bool) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](l_.ID, objc.Sel("lineFragmentUsedRectForGlyphAtIndex:effectiveRange:withoutAdditionalLayout:"), glyphIndex, effectiveGlyphRange, flag)
	return rv
}


// Returns the location for the specified glyph within its line fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/location(forGlyphAt:)
func (l_ LayoutManager) LocationForGlyphAtIndex(glyphIndex uint) objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[corefoundation.Point](l_.ID, objc.Sel("locationForGlyphAtIndex:"), glyphIndex)
	return rv
}


// Indicates whether the glyph at the specified index has a visible representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/notShownAttribute(forGlyphAt:)
func (l_ LayoutManager) NotShownAttributeForGlyphAtIndex(glyphIndex uint) bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("notShownAttributeForGlyphAtIndex:"), glyphIndex)
	return rv
}


// Notifies the layout manager when an edit action changes the contents of its text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/processEditing(for:edited:range:changeInLength:invalidatedRange:)
func (l_ LayoutManager) ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange(textStorage ITextStorage, editMask TextStorageEditActions, newCharRange corefoundation.Range, delta int, invalidatedCharRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("processEditingForTextStorage:edited:range:changeInLength:invalidatedRange:"), textStorage, editMask, newCharRange, delta, invalidatedCharRange)
}


// Returns the glyph property of the glyph at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/propertyForGlyph(at:)
func (l_ LayoutManager) PropertyForGlyphAtIndex(glyphIndex uint) GlyphProperty {
	rv := objc.Send[GlyphProperty](l_.ID, objc.Sel("propertyForGlyphAtIndex:"), glyphIndex)
	return rv
}


// Returns the range of displayable glyphs that surround the glyph at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/range(ofNominallySpacedGlyphsContaining:)
func (l_ LayoutManager) RangeOfNominallySpacedGlyphsContainingIndex(glyphIndex uint) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](l_.ID, objc.Sel("rangeOfNominallySpacedGlyphsContainingIndex:"), glyphIndex)
	return rv
}


// Returns an array of rectangles and, by reference, the number of such rectangles, that define the region in the given container enclosing the given character range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/rectArray(forCharacterRange:withinSelectedCharacterRange:in:rectCount:)
func (l_ LayoutManager) RectArrayForCharacterRangeWithinSelectedCharacterRangeInTextContainerRectCount(charRange corefoundation.Range, selCharRange corefoundation.Range, container ITextContainer, rectCount uint) RectArray /* not a class type */ {
	rv := objc.Send[RectArray](l_.ID, objc.Sel("rectArrayForCharacterRange:withinSelectedCharacterRange:inTextContainer:rectCount:"), charRange, selCharRange, container, rectCount)
	return rv
}


// Returns an array of rectangles and, by reference, the number of such rectangles, that define the region in the given container enclosing the given glyph range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/rectArray(forGlyphRange:withinSelectedGlyphRange:in:rectCount:)
func (l_ LayoutManager) RectArrayForGlyphRangeWithinSelectedGlyphRangeInTextContainerRectCount(glyphRange corefoundation.Range, selGlyphRange corefoundation.Range, container ITextContainer, rectCount uint) RectArray /* not a class type */ {
	rv := objc.Send[RectArray](l_.ID, objc.Sel("rectArrayForGlyphRange:withinSelectedGlyphRange:inTextContainer:rectCount:"), glyphRange, selGlyphRange, container, rectCount)
	return rv
}


// Removes a temporary attribute from the list of attributes for the specified character range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/removeTemporaryAttribute(_:forCharacterRange:)
func (l_ LayoutManager) RemoveTemporaryAttributeForCharacterRange(attrName AttributedStringKey /* not a class type */, charRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("removeTemporaryAttribute:forCharacterRange:"), attrName, charRange)
}


// Removes the text container at the specified index and invalidates the layout as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/removeTextContainer(at:)
func (l_ LayoutManager) RemoveTextContainerAtIndex(index uint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("removeTextContainerAtIndex:"), index)
}


// Replaces the layout manager’s current text storage object with the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/replaceTextStorage(_:)
func (l_ LayoutManager) ReplaceTextStorage(newTextStorage ITextStorage) {
	objc.Send[objc.ID](l_.ID, objc.Sel("replaceTextStorage:"), newTextStorage)
}


// Returns the accessory view that the text system uses for its ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/rulerAccessoryView(for:paragraphStyle:ruler:enabled:)
func (l_ LayoutManager) RulerAccessoryViewForTextViewParagraphStyleRulerEnabled(view ITextView, style IParagraphStyle, ruler IRulerView, isEnabled bool) IView {
	rv := objc.Send[View](l_.ID, objc.Sel("rulerAccessoryViewForTextView:paragraphStyle:ruler:enabled:"), view, style, ruler, isEnabled)
	return rv
}


// Returns an array of text ruler objects for the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/rulerMarkers(for:paragraphStyle:ruler:)
func (l_ LayoutManager) RulerMarkersForTextViewParagraphStyleRuler(view ITextView, style IParagraphStyle, ruler IRulerView) []RulerMarker {
	rv := objc.Send[[]RulerMarker](l_.ID, objc.Sel("rulerMarkersForTextView:paragraphStyle:ruler:"), view, style, ruler)
	return rv
}


// Sets the size to use when drawing a glyph that represents an attachment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setAttachmentSize(_:forGlyphRange:)
func (l_ LayoutManager) SetAttachmentSizeForGlyphRange(attachmentSize objc.IObject /* cross-framework: Size */, glyphRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAttachmentSize:forGlyphRange:"), attachmentSize, glyphRange)
}


// Sets the bounding rectangle that encloses the specified text block and glyph range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setBoundsRect(_:for:glyphRange:)
func (l_ LayoutManager) SetBoundsRectForTextBlockGlyphRange(rect objc.IObject /* cross-framework: Rect */, block ITextBlock, glyphRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBoundsRect:forTextBlock:glyphRange:"), rect, block, glyphRange)
}


// Indicates whether the specified glyph exceeds the bounds of the line fragment for its layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setDrawsOutsideLineFragment(_:forGlyphAt:)
func (l_ LayoutManager) SetDrawsOutsideLineFragmentForGlyphAtIndex(flag bool, glyphIndex uint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDrawsOutsideLineFragment:forGlyphAtIndex:"), flag, glyphIndex)
}


// Sets the bounds and container for the extra line fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setExtraLineFragmentRect(_:usedRect:textContainer:)
func (l_ LayoutManager) SetExtraLineFragmentRectUsedRectTextContainer(fragmentRect objc.IObject /* cross-framework: Rect */, usedRect objc.IObject /* cross-framework: Rect */, container ITextContainer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setExtraLineFragmentRect:usedRect:textContainer:"), fragmentRect, usedRect, container)
}


// Stores the initial glyphs and glyph properties for a character range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setGlyphs(_:properties:characterIndexes:font:forGlyphRange:)
func (l_ LayoutManager) SetGlyphsPropertiesCharacterIndexesFontForGlyphRange(glyphs objc.IObject /* cross-framework: Glyph */, props GlyphProperty, charIndexes uint, aFont IFont, glyphRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setGlyphs:properties:characterIndexes:font:forGlyphRange:"), glyphs, props, charIndexes, aFont, glyphRange)
}


// Sets the layout rectangle that encloses the specified text block and glyph range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setLayoutRect(_:for:glyphRange:)
func (l_ LayoutManager) SetLayoutRectForTextBlockGlyphRange(rect objc.IObject /* cross-framework: Rect */, block ITextBlock, glyphRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLayoutRect:forTextBlock:glyphRange:"), rect, block, glyphRange)
}


// Associates the line fragment bounds for the specified range of glyphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setLineFragmentRect(_:forGlyphRange:usedRect:)
func (l_ LayoutManager) SetLineFragmentRectForGlyphRangeUsedRect(fragmentRect objc.IObject /* cross-framework: Rect */, glyphRange corefoundation.Range, usedRect objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLineFragmentRect:forGlyphRange:usedRect:"), fragmentRect, glyphRange, usedRect)
}


// Sets the location for the first glyph in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setLocation(_:forStartOfGlyphRange:)
func (l_ LayoutManager) SetLocationForStartOfGlyphRange(location objc.IObject /* cross-framework: Point */, glyphRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLocation:forStartOfGlyphRange:"), location, glyphRange)
}


// Sets the visibility of the glyph at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setNotShownAttribute(_:forGlyphAt:)
func (l_ LayoutManager) SetNotShownAttributeForGlyphAtIndex(flag bool, glyphIndex uint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNotShownAttribute:forGlyphAtIndex:"), flag, glyphIndex)
}


// Sets one or more temporary attributes for the specified character range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setTemporaryAttributes(_:forCharacterRange:)
func (l_ LayoutManager) SetTemporaryAttributesForCharacterRange(attrs foundation.IDictionary, charRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTemporaryAttributes:forCharacterRange:"), attrs, charRange)
}


// Associates a text container with the specified range of glyphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/setTextContainer(_:forGlyphRange:)
func (l_ LayoutManager) SetTextContainerForGlyphRange(container ITextContainer, glyphRange corefoundation.Range) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTextContainer:forGlyphRange:"), container, glyphRange)
}


// Draws an attachment cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/showAttachmentCell(_:in:characterIndex:)
func (l_ LayoutManager) ShowAttachmentCellInRectCharacterIndex(cell ICell, rect objc.IObject /* cross-framework: Rect */, attachmentIndex uint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("showAttachmentCell:inRect:characterIndex:"), cell, rect, attachmentIndex)
}


// Renders the glyphs at the specified positions, using the specified attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/showCGGlyphs(_:positions:count:font:textMatrix:attributes:in:)
func (l_ LayoutManager) ShowCGGlyphsPositionsCountFontTextMatrixAttributesInContext(glyphs objc.IObject /* cross-framework: Glyph */, positions corefoundation.CGPoint, glyphCount int, font IFont, textMatrix corefoundation.CGAffineTransform, attributes foundation.IDictionary, CGContext ContextRef /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("showCGGlyphs:positions:count:font:textMatrix:attributes:inContext:"), glyphs, positions, glyphCount, font, textMatrix, attributes, CGContext)
}


// Calculates and draws strikethrough for the specified glyphs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/strikethroughGlyphRange(_:strikethroughType:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:)
func (l_ LayoutManager) StrikethroughGlyphRangeStrikethroughTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange corefoundation.Range, strikethroughVal UnderlineStyle, lineRect objc.IObject /* cross-framework: Rect */, lineGlyphRange corefoundation.Range, containerOrigin objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("strikethroughGlyphRange:strikethroughType:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, strikethroughVal, lineRect, lineGlyphRange, containerOrigin)
}


// Returns the value for the temporary attribute of a character, and the range it applies to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/temporaryAttribute(_:atCharacterIndex:effectiveRange:)
func (l_ LayoutManager) TemporaryAttributeAtCharacterIndexEffectiveRange(attrName AttributedStringKey /* not a class type */, location uint, range_ RangePointer /* not a class type */) objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("temporaryAttribute:atCharacterIndex:effectiveRange:"), attrName, location, range_)
	return rv
}


// Returns the value for the temporary attribute of a character, and the maximum range it applies to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/temporaryAttribute(_:atCharacterIndex:longestEffectiveRange:in:)
func (l_ LayoutManager) TemporaryAttributeAtCharacterIndexLongestEffectiveRangeInRange(attrName AttributedStringKey /* not a class type */, location uint, range_ RangePointer /* not a class type */, rangeLimit corefoundation.Range) objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("temporaryAttribute:atCharacterIndex:longestEffectiveRange:inRange:"), attrName, location, range_, rangeLimit)
	return rv
}


// Returns the dictionary of temporary attributes for the specified character range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/temporaryAttributes(atCharacterIndex:effectiveRange:)
func (l_ LayoutManager) TemporaryAttributesAtCharacterIndexEffectiveRange(charIndex uint, effectiveCharRange RangePointer /* not a class type */) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](l_.ID, objc.Sel("temporaryAttributesAtCharacterIndex:effectiveRange:"), charIndex, effectiveCharRange)
	return rv
}


// Returns the temporary attributes for a character, and the maximum range they apply to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/temporaryAttributes(atCharacterIndex:longestEffectiveRange:in:)
func (l_ LayoutManager) TemporaryAttributesAtCharacterIndexLongestEffectiveRangeInRange(location uint, range_ RangePointer /* not a class type */, rangeLimit corefoundation.Range) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](l_.ID, objc.Sel("temporaryAttributesAtCharacterIndex:longestEffectiveRange:inRange:"), location, range_, rangeLimit)
	return rv
}


// Returns the text container that manages the layout for the specified glyph, causing layout to happen as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/textContainer(forGlyphAt:effectiveRange:)
func (l_ LayoutManager) TextContainerForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange RangePointer /* not a class type */) ITextContainer {
	rv := objc.Send[TextContainer](l_.ID, objc.Sel("textContainerForGlyphAtIndex:effectiveRange:"), glyphIndex, effectiveGlyphRange)
	return rv
}


// Returns the text container that manages the layout for the specified glyph.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/textContainer(forGlyphAt:effectiveRange:withoutAdditionalLayout:)
func (l_ LayoutManager) TextContainerForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange RangePointer /* not a class type */, flag bool) ITextContainer {
	rv := objc.Send[TextContainer](l_.ID, objc.Sel("textContainerForGlyphAtIndex:effectiveRange:withoutAdditionalLayout:"), glyphIndex, effectiveGlyphRange, flag)
	return rv
}


// Invalidates the layout information, and possibly glyphs, for the specified text container and all subsequent text container objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/textContainerChangedGeometry(_:)
func (l_ LayoutManager) TextContainerChangedGeometry(container ITextContainer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("textContainerChangedGeometry:"), container)
}


// Updates the information necessary to manage text view objects for the specified text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/textContainerChangedTextView(_:)
func (l_ LayoutManager) TextContainerChangedTextView(container ITextContainer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("textContainerChangedTextView:"), container)
}


// Returns the range of truncated glyphs for a line fragment that contains the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/truncatedGlyphRange(inLineFragmentForGlyphAt:)
func (l_ LayoutManager) TruncatedGlyphRangeInLineFragmentForGlyphAtIndex(glyphIndex uint) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](l_.ID, objc.Sel("truncatedGlyphRangeInLineFragmentForGlyphAtIndex:"), glyphIndex)
	return rv
}


// Calculates subranges to underline for the specified glyphs and draws the underlining as appropriate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/underlineGlyphRange(_:underlineType:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:)
func (l_ LayoutManager) UnderlineGlyphRangeUnderlineTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange corefoundation.Range, underlineVal UnderlineStyle, lineRect objc.IObject /* cross-framework: Rect */, lineGlyphRange corefoundation.Range, containerOrigin objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("underlineGlyphRange:underlineType:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, underlineVal, lineRect, lineGlyphRange, containerOrigin)
}


// Returns the bounding rectangle for the glyphs in the specified text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/usedRect(for:)
func (l_ LayoutManager) UsedRectForTextContainer(container ITextContainer) objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](l_.ID, objc.Sel("usedRectForTextContainer:"), container)
	return rv
}


// A Boolean value that indicates whether the layout manager allows noncontiguous layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/allowsNonContiguousLayout
func (l_ LayoutManager) AllowsNonContiguousLayout() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("allowsNonContiguousLayout"))
	return rv
}


// A Boolean value that indicates whether the layout manager allows noncontiguous layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/allowsNonContiguousLayout
func (l_ LayoutManager) SetAllowsNonContiguousLayout(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAllowsNonContiguousLayout:"), value)
}


// A Boolean value that indicates whether the layout manager generates glyphs and lays them out when the app’s run loop is idle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/backgroundLayoutEnabled
func (l_ LayoutManager) BackgroundLayoutEnabled() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("backgroundLayoutEnabled"))
	return rv
}


// A Boolean value that indicates whether the layout manager generates glyphs and lays them out when the app’s run loop is idle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/backgroundLayoutEnabled
func (l_ LayoutManager) SetBackgroundLayoutEnabled(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBackgroundLayoutEnabled:"), value)
}


// The default amount of scaling to apply when an attachment image is too large to fit in a text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/defaultAttachmentScaling
func (l_ LayoutManager) DefaultAttachmentScaling() ImageScaling {
	rv := objc.Send[ImageScaling](l_.ID, objc.Sel("defaultAttachmentScaling"))
	return rv
}


// The default amount of scaling to apply when an attachment image is too large to fit in a text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/defaultAttachmentScaling
func (l_ LayoutManager) SetDefaultAttachmentScaling(value ImageScaling) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDefaultAttachmentScaling:"), value)
}


// The layout manager’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/delegate
func (l_ LayoutManager) Delegate() objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("delegate"))
	return rv
}


// The layout manager’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/delegate
func (l_ LayoutManager) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDelegate:"), value)
}


// The rectangle for the extra line fragment at the end of a document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/extraLineFragmentRect
func (l_ LayoutManager) ExtraLineFragmentRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](l_.ID, objc.Sel("extraLineFragmentRect"))
	return rv
}


// The text container for the extra line fragment rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/extraLineFragmentTextContainer
func (l_ LayoutManager) ExtraLineFragmentTextContainer() ITextContainer {
	rv := objc.Send[TextContainer](l_.ID, objc.Sel("extraLineFragmentTextContainer"))
	return rv
}


// The rectangle that encloses the insertion point in the extra line fragment rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/extraLineFragmentUsedRect
func (l_ LayoutManager) ExtraLineFragmentUsedRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](l_.ID, objc.Sel("extraLineFragmentUsedRect"))
	return rv
}


// The first text view in the layout manager’s series of text views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/firstTextView
func (l_ LayoutManager) FirstTextView() ITextView {
	rv := objc.Send[TextView](l_.ID, objc.Sel("firstTextView"))
	return rv
}


// The glyph generator that the layout manager uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphGenerator
func (l_ LayoutManager) GlyphGenerator() IGlyphGenerator {
	rv := objc.Send[GlyphGenerator](l_.ID, objc.Sel("glyphGenerator"))
	return rv
}


// The glyph generator that the layout manager uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/glyphGenerator
func (l_ LayoutManager) SetGlyphGenerator(value IGlyphGenerator) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setGlyphGenerator:"), value)
}


// A Boolean value that indicates whether the layout manager currently has any areas of noncontiguous layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/hasNonContiguousLayout
func (l_ LayoutManager) HasNonContiguousLayout() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("hasNonContiguousLayout"))
	return rv
}


// The threshold controlling when hyphenation is done.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/hyphenationFactor
func (l_ LayoutManager) HyphenationFactor() float32 {
	rv := objc.Send[float32](l_.ID, objc.Sel("hyphenationFactor"))
	return rv
}


// The threshold controlling when hyphenation is done.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/hyphenationFactor
func (l_ LayoutManager) SetHyphenationFactor(value float32) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setHyphenationFactor:"), value)
}


// A Boolean value that indicates whether the layout manager avoids laying out unusually long or suspicious input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/limitsLayoutForSuspiciousContents
func (l_ LayoutManager) LimitsLayoutForSuspiciousContents() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("limitsLayoutForSuspiciousContents"))
	return rv
}


// A Boolean value that indicates whether the layout manager avoids laying out unusually long or suspicious input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/limitsLayoutForSuspiciousContents
func (l_ LayoutManager) SetLimitsLayoutForSuspiciousContents(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLimitsLayoutForSuspiciousContents:"), value)
}


// The number of glyphs in the layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/numberOfGlyphs
func (l_ LayoutManager) NumberOfGlyphs() uint {
	rv := objc.Send[uint](l_.ID, objc.Sel("numberOfGlyphs"))
	return rv
}


// A Boolean value that indicates whether the layout manager substitutes visible glyphs for control characters in the layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/showsControlCharacters
func (l_ LayoutManager) ShowsControlCharacters() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("showsControlCharacters"))
	return rv
}


// A Boolean value that indicates whether the layout manager substitutes visible glyphs for control characters in the layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/showsControlCharacters
func (l_ LayoutManager) SetShowsControlCharacters(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShowsControlCharacters:"), value)
}


// A Boolean value that indicates whether to substitute visible glyphs for whitespace and other typically invisible characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/showsInvisibleCharacters
func (l_ LayoutManager) ShowsInvisibleCharacters() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("showsInvisibleCharacters"))
	return rv
}


// A Boolean value that indicates whether to substitute visible glyphs for whitespace and other typically invisible characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/showsInvisibleCharacters
func (l_ LayoutManager) SetShowsInvisibleCharacters(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShowsInvisibleCharacters:"), value)
}


// The current text containers of the layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/textContainers
func (l_ LayoutManager) TextContainers() []TextContainer {
	rv := objc.Send[[]TextContainer](l_.ID, objc.Sel("textContainers"))
	return rv
}


// The text storage object that contains the content to lay out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/textStorage
func (l_ LayoutManager) TextStorage() ITextStorage {
	rv := objc.Send[TextStorage](l_.ID, objc.Sel("textStorage"))
	return rv
}


// The text storage object that contains the content to lay out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/textStorage
func (l_ LayoutManager) SetTextStorage(value ITextStorage) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTextStorage:"), value)
}


// The text view that contains the first glyph in the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/textViewForBeginningOfSelection
func (l_ LayoutManager) TextViewForBeginningOfSelection() ITextView {
	rv := objc.Send[TextView](l_.ID, objc.Sel("textViewForBeginningOfSelection"))
	return rv
}


// The current typesetter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/typesetter
func (l_ LayoutManager) Typesetter() ITypesetter {
	rv := objc.Send[Typesetter](l_.ID, objc.Sel("typesetter"))
	return rv
}


// The current typesetter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/typesetter
func (l_ LayoutManager) SetTypesetter(value ITypesetter) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTypesetter:"), value)
}


// The default typesetter behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/typesetterBehavior-swift.property
func (l_ LayoutManager) TypesetterBehavior() TypesetterBehavior {
	rv := objc.Send[TypesetterBehavior](l_.ID, objc.Sel("typesetterBehavior"))
	return rv
}


// The default typesetter behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/typesetterBehavior-swift.property
func (l_ LayoutManager) SetTypesetterBehavior(value TypesetterBehavior) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTypesetterBehavior:"), value)
}


// A Boolean value that indicates whether the layout manager uses the default hyphenation rules to wrap lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/usesDefaultHyphenation
func (l_ LayoutManager) UsesDefaultHyphenation() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("usesDefaultHyphenation"))
	return rv
}


// A Boolean value that indicates whether the layout manager uses the default hyphenation rules to wrap lines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/usesDefaultHyphenation
func (l_ LayoutManager) SetUsesDefaultHyphenation(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setUsesDefaultHyphenation:"), value)
}


// A Boolean value that indicates whether the layout manager uses the leading of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/usesFontLeading
func (l_ LayoutManager) UsesFontLeading() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("usesFontLeading"))
	return rv
}


// A Boolean value that indicates whether the layout manager uses the leading of the font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/usesFontLeading
func (l_ LayoutManager) SetUsesFontLeading(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setUsesFontLeading:"), value)
}


// A Boolean that controls using screen fonts to calculate layout and display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/usesScreenFonts
func (l_ LayoutManager) UsesScreenFonts() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("usesScreenFonts"))
	return rv
}


// A Boolean that controls using screen fonts to calculate layout and display text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/usesScreenFonts
func (l_ LayoutManager) SetUsesScreenFonts(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setUsesScreenFonts:"), value)
}


