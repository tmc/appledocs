// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PLayoutManagerDelegate is the NSLayoutManagerDelegate protocol interface.
//
// A set of optional methods that delegates of layout manager objects implement.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSLayoutManagerDelegate
type PLayoutManagerDelegate interface {
	// Optional methods
	LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(layoutManager ILayoutManager, glyphIndex uint, textContainer ITextContainer, proposedRect corefoundation.CGRect, glyphPosition corefoundation.CGPoint, charIndex uint) corefoundation.CGRect
	HasLayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex() bool
	LayoutManagerDidCompleteLayoutForTextContainerAtEnd(layoutManager ILayoutManager, textContainer ITextContainer, layoutFinishedFlag bool)
	HasLayoutManagerDidCompleteLayoutForTextContainerAtEnd() bool
	LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64
	HasLayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect() bool
	LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64
	HasLayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect() bool
	LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64
	HasLayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect() bool
	LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex(layoutManager ILayoutManager, charIndex uint) bool
	HasLayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex() bool
	LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex(layoutManager ILayoutManager, charIndex uint) bool
	HasLayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex() bool
	LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange(layoutManager ILayoutManager, glyphs Glyph, props GlyphProperty, charIndexes uint, aFont IFont, glyphRange foundation.Range) uint
	HasLayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange() bool
	LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange(layoutManager ILayoutManager, lineFragmentRect corefoundation.CGRect, lineFragmentUsedRect corefoundation.CGRect, baselineOffset corefoundation.CGFloat, textContainer ITextContainer, glyphRange foundation.Range) bool
	HasLayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange() bool
	LayoutManagerShouldUseActionForControlCharacterAtIndex(layoutManager ILayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction
	HasLayoutManagerShouldUseActionForControlCharacterAtIndex() bool
	LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(layoutManager ILayoutManager, attrs foundation.IDictionary, toScreen bool, charIndex uint, effectiveCharRange RangePointer /* not a class type */) foundation.IDictionary
	HasLayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange() bool
	LayoutManagerTextContainerDidChangeGeometryFromSize(layoutManager ILayoutManager, textContainer ITextContainer, oldSize corefoundation.CGSize)
	HasLayoutManagerTextContainerDidChangeGeometryFromSize() bool
	LayoutManagerDidInvalidateLayout(sender ILayoutManager)
	HasLayoutManagerDidInvalidateLayout() bool
}

// LayoutManagerDelegate is a delegate implementation builder for the PLayoutManagerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type LayoutManagerDelegate struct {
	_LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex func(layoutManager ILayoutManager, glyphIndex uint, textContainer ITextContainer, proposedRect corefoundation.CGRect, glyphPosition corefoundation.CGPoint, charIndex uint) corefoundation.CGRect
	_LayoutManagerDidCompleteLayoutForTextContainerAtEnd func(layoutManager ILayoutManager, textContainer ITextContainer, layoutFinishedFlag bool)
	_LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect func(layoutManager ILayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64
	_LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect func(layoutManager ILayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64
	_LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect func(layoutManager ILayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64
	_LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex func(layoutManager ILayoutManager, charIndex uint) bool
	_LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex func(layoutManager ILayoutManager, charIndex uint) bool
	_LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange func(layoutManager ILayoutManager, glyphs Glyph, props GlyphProperty, charIndexes uint, aFont IFont, glyphRange foundation.Range) uint
	_LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange func(layoutManager ILayoutManager, lineFragmentRect corefoundation.CGRect, lineFragmentUsedRect corefoundation.CGRect, baselineOffset corefoundation.CGFloat, textContainer ITextContainer, glyphRange foundation.Range) bool
	_LayoutManagerShouldUseActionForControlCharacterAtIndex func(layoutManager ILayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction
	_LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange func(layoutManager ILayoutManager, attrs foundation.IDictionary, toScreen bool, charIndex uint, effectiveCharRange RangePointer /* not a class type */) foundation.IDictionary
	_LayoutManagerTextContainerDidChangeGeometryFromSize func(layoutManager ILayoutManager, textContainer ITextContainer, oldSize corefoundation.CGSize)
	_LayoutManagerDidInvalidateLayout func(sender ILayoutManager)
}

// SetLayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex sets the handler for the LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex delegate method.
//
// Returns the bounding rectangle for the specified control glyph with the specified parameters.
func (d *LayoutManagerDelegate) SetLayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(f func(layoutManager ILayoutManager, glyphIndex uint, textContainer ITextContainer, proposedRect corefoundation.CGRect, glyphPosition corefoundation.CGPoint, charIndex uint) corefoundation.CGRect) {
	d._LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex = f
}

// SetLayoutManagerDidCompleteLayoutForTextContainerAtEnd sets the handler for the LayoutManagerDidCompleteLayoutForTextContainerAtEnd delegate method.
//
// Informs the delegate when the layout manager finishes laying out text in the specified text container.
func (d *LayoutManagerDelegate) SetLayoutManagerDidCompleteLayoutForTextContainerAtEnd(f func(layoutManager ILayoutManager, textContainer ITextContainer, layoutFinishedFlag bool)) {
	d._LayoutManagerDidCompleteLayoutForTextContainerAtEnd = f
}

// SetLayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect sets the handler for the LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect delegate method.
//
// Returns the amount of space to add to the end of a line.
func (d *LayoutManagerDelegate) SetLayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(f func(layoutManager ILayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64) {
	d._LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect = f
}

// SetLayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect sets the handler for the LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect delegate method.
//
// Returns the amount of space to add at the end of a paragraph.
func (d *LayoutManagerDelegate) SetLayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(f func(layoutManager ILayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64) {
	d._LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect = f
}

// SetLayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect sets the handler for the LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect delegate method.
//
// Returns the amount of space to add at the beginning of a paragraph.
func (d *LayoutManagerDelegate) SetLayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(f func(layoutManager ILayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64) {
	d._LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect = f
}

// SetLayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex sets the handler for the LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex delegate method.
//
// Asks the delegate whether to break the line at the specified character.
func (d *LayoutManagerDelegate) SetLayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex(f func(layoutManager ILayoutManager, charIndex uint) bool) {
	d._LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex = f
}

// SetLayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex sets the handler for the LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex delegate method.
//
// Asks the delegate whether to break the line at the specified word.
func (d *LayoutManagerDelegate) SetLayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex(f func(layoutManager ILayoutManager, charIndex uint) bool) {
	d._LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex = f
}

// SetLayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange sets the handler for the LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange delegate method.
//
// Enables customization of the initial glyph generation process.
func (d *LayoutManagerDelegate) SetLayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange(f func(layoutManager ILayoutManager, glyphs Glyph, props GlyphProperty, charIndexes uint, aFont IFont, glyphRange foundation.Range) uint) {
	d._LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange = f
}

// SetLayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange sets the handler for the LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange delegate method.
//
// Customizes the line fragment geometry before committing it to the layout cache.
func (d *LayoutManagerDelegate) SetLayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange(f func(layoutManager ILayoutManager, lineFragmentRect corefoundation.CGRect, lineFragmentUsedRect corefoundation.CGRect, baselineOffset corefoundation.CGFloat, textContainer ITextContainer, glyphRange foundation.Range) bool) {
	d._LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange = f
}

// SetLayoutManagerShouldUseActionForControlCharacterAtIndex sets the handler for the LayoutManagerShouldUseActionForControlCharacterAtIndex delegate method.
//
// Returns the control character action for the control character at the specified character index.
func (d *LayoutManagerDelegate) SetLayoutManagerShouldUseActionForControlCharacterAtIndex(f func(layoutManager ILayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction) {
	d._LayoutManagerShouldUseActionForControlCharacterAtIndex = f
}

// SetLayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange sets the handler for the LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange delegate method.
//
// Asks the delegate whether to use temporary attributes when drawing the text.
func (d *LayoutManagerDelegate) SetLayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(f func(layoutManager ILayoutManager, attrs foundation.IDictionary, toScreen bool, charIndex uint, effectiveCharRange RangePointer /* not a class type */) foundation.IDictionary) {
	d._LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange = f
}

// SetLayoutManagerTextContainerDidChangeGeometryFromSize sets the handler for the LayoutManagerTextContainerDidChangeGeometryFromSize delegate method.
//
// Informs the delegate when the layout manager invalidates layout due to a change in the geometry of the specified text container.
func (d *LayoutManagerDelegate) SetLayoutManagerTextContainerDidChangeGeometryFromSize(f func(layoutManager ILayoutManager, textContainer ITextContainer, oldSize corefoundation.CGSize)) {
	d._LayoutManagerTextContainerDidChangeGeometryFromSize = f
}

// SetLayoutManagerDidInvalidateLayout sets the handler for the LayoutManagerDidInvalidateLayout delegate method.
//
// Informs the delegate when the specified layout manager invalidates layout information (not glyph information).
func (d *LayoutManagerDelegate) SetLayoutManagerDidInvalidateLayout(f func(sender ILayoutManager)) {
	d._LayoutManagerDidInvalidateLayout = f
}

// LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex implements the PLayoutManagerDelegate interface.
func (d *LayoutManagerDelegate) LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(layoutManager ILayoutManager, glyphIndex uint, textContainer ITextContainer, proposedRect corefoundation.CGRect, glyphPosition corefoundation.CGPoint, charIndex uint) corefoundation.CGRect {
	if d._LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex != nil {
		return d._LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(layoutManager, glyphIndex, textContainer, proposedRect, glyphPosition, charIndex)
	}
	var zero corefoundation.CGRect
	return zero
}

// HasLayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex returns true if a handler for LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex has been set.
func (d *LayoutManagerDelegate) HasLayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex() bool {
	return d._LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex != nil
}

// LayoutManagerDidCompleteLayoutForTextContainerAtEnd implements the PLayoutManagerDelegate interface.
func (d *LayoutManagerDelegate) LayoutManagerDidCompleteLayoutForTextContainerAtEnd(layoutManager ILayoutManager, textContainer ITextContainer, layoutFinishedFlag bool) {
	if d._LayoutManagerDidCompleteLayoutForTextContainerAtEnd != nil {
		d._LayoutManagerDidCompleteLayoutForTextContainerAtEnd(layoutManager, textContainer, layoutFinishedFlag)
	}
}

// HasLayoutManagerDidCompleteLayoutForTextContainerAtEnd returns true if a handler for LayoutManagerDidCompleteLayoutForTextContainerAtEnd has been set.
func (d *LayoutManagerDelegate) HasLayoutManagerDidCompleteLayoutForTextContainerAtEnd() bool {
	return d._LayoutManagerDidCompleteLayoutForTextContainerAtEnd != nil
}

// LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect implements the PLayoutManagerDelegate interface.
func (d *LayoutManagerDelegate) LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64 {
	if d._LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect != nil {
		return d._LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager, glyphIndex, rect)
	}
	var zero float64
	return zero
}

// HasLayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect returns true if a handler for LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect has been set.
func (d *LayoutManagerDelegate) HasLayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect() bool {
	return d._LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect != nil
}

// LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect implements the PLayoutManagerDelegate interface.
func (d *LayoutManagerDelegate) LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64 {
	if d._LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect != nil {
		return d._LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager, glyphIndex, rect)
	}
	var zero float64
	return zero
}

// HasLayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect returns true if a handler for LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect has been set.
func (d *LayoutManagerDelegate) HasLayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect() bool {
	return d._LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect != nil
}

// LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect implements the PLayoutManagerDelegate interface.
func (d *LayoutManagerDelegate) LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64 {
	if d._LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect != nil {
		return d._LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(layoutManager, glyphIndex, rect)
	}
	var zero float64
	return zero
}

// HasLayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect returns true if a handler for LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect has been set.
func (d *LayoutManagerDelegate) HasLayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect() bool {
	return d._LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect != nil
}

// LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex implements the PLayoutManagerDelegate interface.
func (d *LayoutManagerDelegate) LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex(layoutManager ILayoutManager, charIndex uint) bool {
	if d._LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex != nil {
		return d._LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex(layoutManager, charIndex)
	}
	var zero bool
	return zero
}

// HasLayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex returns true if a handler for LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex has been set.
func (d *LayoutManagerDelegate) HasLayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex() bool {
	return d._LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex != nil
}

// LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex implements the PLayoutManagerDelegate interface.
func (d *LayoutManagerDelegate) LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex(layoutManager ILayoutManager, charIndex uint) bool {
	if d._LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex != nil {
		return d._LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex(layoutManager, charIndex)
	}
	var zero bool
	return zero
}

// HasLayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex returns true if a handler for LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex has been set.
func (d *LayoutManagerDelegate) HasLayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex() bool {
	return d._LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex != nil
}

// LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange implements the PLayoutManagerDelegate interface.
func (d *LayoutManagerDelegate) LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange(layoutManager ILayoutManager, glyphs Glyph, props GlyphProperty, charIndexes uint, aFont IFont, glyphRange foundation.Range) uint {
	if d._LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange != nil {
		return d._LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange(layoutManager, glyphs, props, charIndexes, aFont, glyphRange)
	}
	var zero uint
	return zero
}

// HasLayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange returns true if a handler for LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange has been set.
func (d *LayoutManagerDelegate) HasLayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange() bool {
	return d._LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange != nil
}

// LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange implements the PLayoutManagerDelegate interface.
func (d *LayoutManagerDelegate) LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange(layoutManager ILayoutManager, lineFragmentRect corefoundation.CGRect, lineFragmentUsedRect corefoundation.CGRect, baselineOffset corefoundation.CGFloat, textContainer ITextContainer, glyphRange foundation.Range) bool {
	if d._LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange != nil {
		return d._LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange(layoutManager, lineFragmentRect, lineFragmentUsedRect, baselineOffset, textContainer, glyphRange)
	}
	var zero bool
	return zero
}

// HasLayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange returns true if a handler for LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange has been set.
func (d *LayoutManagerDelegate) HasLayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange() bool {
	return d._LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange != nil
}

// LayoutManagerShouldUseActionForControlCharacterAtIndex implements the PLayoutManagerDelegate interface.
func (d *LayoutManagerDelegate) LayoutManagerShouldUseActionForControlCharacterAtIndex(layoutManager ILayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction {
	if d._LayoutManagerShouldUseActionForControlCharacterAtIndex != nil {
		return d._LayoutManagerShouldUseActionForControlCharacterAtIndex(layoutManager, action, charIndex)
	}
	var zero ControlCharacterAction
	return zero
}

// HasLayoutManagerShouldUseActionForControlCharacterAtIndex returns true if a handler for LayoutManagerShouldUseActionForControlCharacterAtIndex has been set.
func (d *LayoutManagerDelegate) HasLayoutManagerShouldUseActionForControlCharacterAtIndex() bool {
	return d._LayoutManagerShouldUseActionForControlCharacterAtIndex != nil
}

// LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange implements the PLayoutManagerDelegate interface.
func (d *LayoutManagerDelegate) LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(layoutManager ILayoutManager, attrs foundation.IDictionary, toScreen bool, charIndex uint, effectiveCharRange RangePointer /* not a class type */) foundation.IDictionary {
	if d._LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange != nil {
		return d._LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(layoutManager, attrs, toScreen, charIndex, effectiveCharRange)
	}
	var zero foundation.IDictionary
	return zero
}

// HasLayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange returns true if a handler for LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange has been set.
func (d *LayoutManagerDelegate) HasLayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange() bool {
	return d._LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange != nil
}

// LayoutManagerTextContainerDidChangeGeometryFromSize implements the PLayoutManagerDelegate interface.
func (d *LayoutManagerDelegate) LayoutManagerTextContainerDidChangeGeometryFromSize(layoutManager ILayoutManager, textContainer ITextContainer, oldSize corefoundation.CGSize) {
	if d._LayoutManagerTextContainerDidChangeGeometryFromSize != nil {
		d._LayoutManagerTextContainerDidChangeGeometryFromSize(layoutManager, textContainer, oldSize)
	}
}

// HasLayoutManagerTextContainerDidChangeGeometryFromSize returns true if a handler for LayoutManagerTextContainerDidChangeGeometryFromSize has been set.
func (d *LayoutManagerDelegate) HasLayoutManagerTextContainerDidChangeGeometryFromSize() bool {
	return d._LayoutManagerTextContainerDidChangeGeometryFromSize != nil
}

// LayoutManagerDidInvalidateLayout implements the PLayoutManagerDelegate interface.
func (d *LayoutManagerDelegate) LayoutManagerDidInvalidateLayout(sender ILayoutManager) {
	if d._LayoutManagerDidInvalidateLayout != nil {
		d._LayoutManagerDidInvalidateLayout(sender)
	}
}

// HasLayoutManagerDidInvalidateLayout returns true if a handler for LayoutManagerDidInvalidateLayout has been set.
func (d *LayoutManagerDelegate) HasLayoutManagerDidInvalidateLayout() bool {
	return d._LayoutManagerDidInvalidateLayout != nil
}

// LayoutManagerDelegateObject wraps an existing Objective-C object that conforms to the PLayoutManagerDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type LayoutManagerDelegateObject struct {
	objectivec.Object
}

// NewLayoutManagerDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSLayoutManagerDelegate protocol.
func NewLayoutManagerDelegateObject(obj objectivec.Object) *LayoutManagerDelegateObject {
	return &LayoutManagerDelegateObject{obj}
}

// Make sure LayoutManagerDelegateObject implements PLayoutManagerDelegate.
var _ PLayoutManagerDelegate = (*LayoutManagerDelegateObject)(nil)

// LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex implements the PLayoutManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *LayoutManagerDelegateObject) LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(layoutManager ILayoutManager, glyphIndex uint, textContainer ITextContainer, proposedRect corefoundation.CGRect, glyphPosition corefoundation.CGPoint, charIndex uint) corefoundation.CGRect {
	return objc.Send[corefoundation.CGRect](o.ID, objc.Sel("layoutManager:boundingBoxForControlGlyphAtIndex:forTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), layoutManager, glyphIndex, textContainer, proposedRect, glyphPosition, charIndex)
}

// HasLayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex returns true; this is a placeholder for optional method checks.
func (o *LayoutManagerDelegateObject) HasLayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// LayoutManagerDidCompleteLayoutForTextContainerAtEnd implements the PLayoutManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *LayoutManagerDelegateObject) LayoutManagerDidCompleteLayoutForTextContainerAtEnd(layoutManager ILayoutManager, textContainer ITextContainer, layoutFinishedFlag bool) {
	objc.Send[objc.ID](o.ID, objc.Sel("layoutManager:didCompleteLayoutForTextContainer:atEnd:"), layoutManager, textContainer, layoutFinishedFlag)
}

// HasLayoutManagerDidCompleteLayoutForTextContainerAtEnd returns true; this is a placeholder for optional method checks.
func (o *LayoutManagerDelegateObject) HasLayoutManagerDidCompleteLayoutForTextContainerAtEnd() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect implements the PLayoutManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *LayoutManagerDelegateObject) LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64 {
	return objc.Send[float64](o.ID, objc.Sel("layoutManager:lineSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), layoutManager, glyphIndex, rect)
}

// HasLayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect returns true; this is a placeholder for optional method checks.
func (o *LayoutManagerDelegateObject) HasLayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect implements the PLayoutManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *LayoutManagerDelegateObject) LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64 {
	return objc.Send[float64](o.ID, objc.Sel("layoutManager:paragraphSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), layoutManager, glyphIndex, rect)
}

// HasLayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect returns true; this is a placeholder for optional method checks.
func (o *LayoutManagerDelegateObject) HasLayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect implements the PLayoutManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *LayoutManagerDelegateObject) LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect corefoundation.CGRect) float64 {
	return objc.Send[float64](o.ID, objc.Sel("layoutManager:paragraphSpacingBeforeGlyphAtIndex:withProposedLineFragmentRect:"), layoutManager, glyphIndex, rect)
}

// HasLayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect returns true; this is a placeholder for optional method checks.
func (o *LayoutManagerDelegateObject) HasLayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex implements the PLayoutManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *LayoutManagerDelegateObject) LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex(layoutManager ILayoutManager, charIndex uint) bool {
	return objc.Send[bool](o.ID, objc.Sel("layoutManager:shouldBreakLineByHyphenatingBeforeCharacterAtIndex:"), layoutManager, charIndex)
}

// HasLayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex returns true; this is a placeholder for optional method checks.
func (o *LayoutManagerDelegateObject) HasLayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex implements the PLayoutManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *LayoutManagerDelegateObject) LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex(layoutManager ILayoutManager, charIndex uint) bool {
	return objc.Send[bool](o.ID, objc.Sel("layoutManager:shouldBreakLineByWordBeforeCharacterAtIndex:"), layoutManager, charIndex)
}

// HasLayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex returns true; this is a placeholder for optional method checks.
func (o *LayoutManagerDelegateObject) HasLayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange implements the PLayoutManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *LayoutManagerDelegateObject) LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange(layoutManager ILayoutManager, glyphs Glyph, props GlyphProperty, charIndexes uint, aFont IFont, glyphRange foundation.Range) uint {
	return objc.Send[uint](o.ID, objc.Sel("layoutManager:shouldGenerateGlyphs:properties:characterIndexes:font:forGlyphRange:"), layoutManager, glyphs, props, charIndexes, aFont, glyphRange)
}

// HasLayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange returns true; this is a placeholder for optional method checks.
func (o *LayoutManagerDelegateObject) HasLayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange implements the PLayoutManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *LayoutManagerDelegateObject) LayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange(layoutManager ILayoutManager, lineFragmentRect corefoundation.CGRect, lineFragmentUsedRect corefoundation.CGRect, baselineOffset corefoundation.CGFloat, textContainer ITextContainer, glyphRange foundation.Range) bool {
	return objc.Send[bool](o.ID, objc.Sel("layoutManager:shouldSetLineFragmentRect:lineFragmentUsedRect:baselineOffset:inTextContainer:forGlyphRange:"), layoutManager, lineFragmentRect, lineFragmentUsedRect, baselineOffset, textContainer, glyphRange)
}

// HasLayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange returns true; this is a placeholder for optional method checks.
func (o *LayoutManagerDelegateObject) HasLayoutManagerShouldSetLineFragmentRectLineFragmentUsedRectBaselineOffsetInTextContainerForGlyphRange() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// LayoutManagerShouldUseActionForControlCharacterAtIndex implements the PLayoutManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *LayoutManagerDelegateObject) LayoutManagerShouldUseActionForControlCharacterAtIndex(layoutManager ILayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction {
	return objc.Send[ControlCharacterAction](o.ID, objc.Sel("layoutManager:shouldUseAction:forControlCharacterAtIndex:"), layoutManager, action, charIndex)
}

// HasLayoutManagerShouldUseActionForControlCharacterAtIndex returns true; this is a placeholder for optional method checks.
func (o *LayoutManagerDelegateObject) HasLayoutManagerShouldUseActionForControlCharacterAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange implements the PLayoutManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *LayoutManagerDelegateObject) LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(layoutManager ILayoutManager, attrs foundation.IDictionary, toScreen bool, charIndex uint, effectiveCharRange RangePointer /* not a class type */) foundation.IDictionary {
	return objc.Send[foundation.IDictionary](o.ID, objc.Sel("layoutManager:shouldUseTemporaryAttributes:forDrawingToScreen:atCharacterIndex:effectiveRange:"), layoutManager, attrs, toScreen, charIndex, effectiveCharRange)
}

// HasLayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange returns true; this is a placeholder for optional method checks.
func (o *LayoutManagerDelegateObject) HasLayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// LayoutManagerTextContainerDidChangeGeometryFromSize implements the PLayoutManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *LayoutManagerDelegateObject) LayoutManagerTextContainerDidChangeGeometryFromSize(layoutManager ILayoutManager, textContainer ITextContainer, oldSize corefoundation.CGSize) {
	objc.Send[objc.ID](o.ID, objc.Sel("layoutManager:textContainer:didChangeGeometryFromSize:"), layoutManager, textContainer, oldSize)
}

// HasLayoutManagerTextContainerDidChangeGeometryFromSize returns true; this is a placeholder for optional method checks.
func (o *LayoutManagerDelegateObject) HasLayoutManagerTextContainerDidChangeGeometryFromSize() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// LayoutManagerDidInvalidateLayout implements the PLayoutManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *LayoutManagerDelegateObject) LayoutManagerDidInvalidateLayout(sender ILayoutManager) {
	objc.Send[objc.ID](o.ID, objc.Sel("layoutManagerDidInvalidateLayout:"), sender)
}

// HasLayoutManagerDidInvalidateLayout returns true; this is a placeholder for optional method checks.
func (o *LayoutManagerDelegateObject) HasLayoutManagerDidInvalidateLayout() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
