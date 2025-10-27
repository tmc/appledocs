// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PTextViewDelegate is the NSTextViewDelegate protocol interface.
//
// A set of optional methods that text view delegates can use to manage selection, set text attributes, work with the spell checker, and more.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSTextViewDelegate
type PTextViewDelegate interface {
	// Required methods
	TextViewClickedOnCellInRect(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect)
	TextViewClickedOnLink(textView ITextView, link objectivec.IObject) bool
	TextViewDoubleClickedOnCellInRect(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect)
	TextViewDraggedCellInRectEvent(view ITextView, cell unsafe.Pointer, rect corefoundation.CGRect, event IEvent)
	// Optional methods
	TextViewShouldSetSpellingStateRange(textView ITextView, value int, affectedCharRange foundation.Range) int
	HasTextViewShouldSetSpellingStateRange() bool
	TextViewWritingToolsIgnoredRangesInEnclosingRange(textView ITextView, enclosingRange foundation.Range) []foundation.Value
	HasTextViewWritingToolsIgnoredRangesInEnclosingRange() bool
	TextViewWritingToolsDidEnd(textView ITextView)
	HasTextViewWritingToolsDidEnd() bool
	TextViewWritingToolsWillBegin(textView ITextView)
	HasTextViewWritingToolsWillBegin() bool
	TextViewCandidatesForSelectedRange(textView ITextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult
	HasTextViewCandidatesForSelectedRange() bool
	TextViewClickedOnCellInRectAtIndex(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect, charIndex uint)
	HasTextViewClickedOnCellInRectAtIndex() bool
	TextViewClickedOnLinkAtIndex(textView ITextView, link objectivec.IObject, charIndex uint) bool
	HasTextViewClickedOnLinkAtIndex() bool
	TextViewCompletionsForPartialWordRangeIndexOfSelectedItem(textView ITextView, words []string, charRange foundation.Range, index int) []string
	HasTextViewCompletionsForPartialWordRangeIndexOfSelectedItem() bool
	TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount(view ITextView, range_ foundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.TextCheckingResult
	HasTextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount() bool
	TextViewDoCommandBySelector(textView ITextView, commandSelector objc.SEL) bool
	HasTextViewDoCommandBySelector() bool
	TextViewDoubleClickedOnCellInRectAtIndex(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect, charIndex uint)
	HasTextViewDoubleClickedOnCellInRectAtIndex() bool
	TextViewDraggedCellInRectEventAtIndex(view ITextView, cell unsafe.Pointer, rect corefoundation.CGRect, event IEvent, charIndex uint)
	HasTextViewDraggedCellInRectEventAtIndex() bool
	TextViewMenuForEventAtIndex(view ITextView, menu IMenu, event IEvent, charIndex uint) IMenu
	HasTextViewMenuForEventAtIndex() bool
	TextViewShouldChangeTextInRangeReplacementString(textView ITextView, affectedCharRange foundation.Range, replacementString foundation.foundation.INSString) bool
	HasTextViewShouldChangeTextInRangeReplacementString() bool
	TextViewShouldChangeTextInRangesReplacementStrings(textView ITextView, affectedRanges []foundation.Value, replacementStrings []string) bool
	HasTextViewShouldChangeTextInRangesReplacementStrings() bool
	TextViewShouldChangeTypingAttributesToAttributes(textView ITextView, oldTypingAttributes foundation.IDictionary, newTypingAttributes foundation.IDictionary) foundation.IDictionary
	HasTextViewShouldChangeTypingAttributesToAttributes() bool
	TextViewShouldSelectCandidateAtIndex(textView ITextView, index uint) bool
	HasTextViewShouldSelectCandidateAtIndex() bool
	TextViewShouldUpdateTouchBarItemIdentifiers(textView ITextView, identifiers []string) []string
	HasTextViewShouldUpdateTouchBarItemIdentifiers() bool
	TextViewURLForContentsOfTextAttachmentAtIndex(textView ITextView, textAttachment ITextAttachment, charIndex uint) foundation.URL
	HasTextViewURLForContentsOfTextAttachmentAtIndex() bool
	TextViewWillChangeSelectionFromCharacterRangeToCharacterRange(textView ITextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range
	HasTextViewWillChangeSelectionFromCharacterRangeToCharacterRange() bool
	TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges(textView ITextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.Value
	HasTextViewWillChangeSelectionFromCharacterRangesToCharacterRanges() bool
	TextViewWillCheckTextInRangeOptionsTypes(view ITextView, range_ foundation.Range, options foundation.IDictionary, checkingTypes TextCheckingTypes /* not a class type */) foundation.IDictionary
	HasTextViewWillCheckTextInRangeOptionsTypes() bool
	TextViewWillDisplayToolTipForCharacterAtIndex(textView ITextView, tooltip foundation.foundation.INSString, characterIndex uint) foundation.String
	HasTextViewWillDisplayToolTipForCharacterAtIndex() bool
	TextViewWillShowSharingServicePickerForItems(textView ITextView, servicePicker ISharingServicePicker, items foundation.foundation.INSArray) ISharingServicePicker
	HasTextViewWillShowSharingServicePickerForItems() bool
	TextViewWritablePasteboardTypesForCellAtIndex(view ITextView, cell unsafe.Pointer, charIndex uint) []string
	HasTextViewWritablePasteboardTypesForCellAtIndex() bool
	TextViewWriteCellAtIndexToPasteboardType(view ITextView, cell unsafe.Pointer, charIndex uint, pboard IPasteboard, type_ PasteboardType) bool
	HasTextViewWriteCellAtIndexToPasteboardType() bool
	TextViewDidChangeSelection(notification foundation.foundation.INSNotification)
	HasTextViewDidChangeSelection() bool
	TextViewDidChangeTypingAttributes(notification foundation.foundation.INSNotification)
	HasTextViewDidChangeTypingAttributes() bool
	UndoManagerForTextView(view ITextView) foundation.UndoManager
	HasUndoManagerForTextView() bool
}

// TextViewDelegate is a delegate implementation builder for the PTextViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TextViewDelegate struct {
	_TextViewShouldSetSpellingStateRange func(textView ITextView, value int, affectedCharRange foundation.Range) int
	_TextViewWritingToolsIgnoredRangesInEnclosingRange func(textView ITextView, enclosingRange foundation.Range) []foundation.Value
	_TextViewWritingToolsDidEnd func(textView ITextView)
	_TextViewWritingToolsWillBegin func(textView ITextView)
	_TextViewCandidatesForSelectedRange func(textView ITextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult
	_TextViewClickedOnCellInRectAtIndex func(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect, charIndex uint)
	_TextViewClickedOnLinkAtIndex func(textView ITextView, link objectivec.IObject, charIndex uint) bool
	_TextViewCompletionsForPartialWordRangeIndexOfSelectedItem func(textView ITextView, words []string, charRange foundation.Range, index int) []string
	_TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount func(view ITextView, range_ foundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.TextCheckingResult
	_TextViewDoCommandBySelector func(textView ITextView, commandSelector objc.SEL) bool
	_TextViewDoubleClickedOnCellInRectAtIndex func(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect, charIndex uint)
	_TextViewDraggedCellInRectEventAtIndex func(view ITextView, cell unsafe.Pointer, rect corefoundation.CGRect, event IEvent, charIndex uint)
	_TextViewMenuForEventAtIndex func(view ITextView, menu IMenu, event IEvent, charIndex uint) IMenu
	_TextViewShouldChangeTextInRangeReplacementString func(textView ITextView, affectedCharRange foundation.Range, replacementString foundation.foundation.INSString) bool
	_TextViewShouldChangeTextInRangesReplacementStrings func(textView ITextView, affectedRanges []foundation.Value, replacementStrings []string) bool
	_TextViewShouldChangeTypingAttributesToAttributes func(textView ITextView, oldTypingAttributes foundation.IDictionary, newTypingAttributes foundation.IDictionary) foundation.IDictionary
	_TextViewShouldSelectCandidateAtIndex func(textView ITextView, index uint) bool
	_TextViewShouldUpdateTouchBarItemIdentifiers func(textView ITextView, identifiers []string) []string
	_TextViewURLForContentsOfTextAttachmentAtIndex func(textView ITextView, textAttachment ITextAttachment, charIndex uint) foundation.URL
	_TextViewWillChangeSelectionFromCharacterRangeToCharacterRange func(textView ITextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range
	_TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges func(textView ITextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.Value
	_TextViewWillCheckTextInRangeOptionsTypes func(view ITextView, range_ foundation.Range, options foundation.IDictionary, checkingTypes TextCheckingTypes /* not a class type */) foundation.IDictionary
	_TextViewWillDisplayToolTipForCharacterAtIndex func(textView ITextView, tooltip foundation.foundation.INSString, characterIndex uint) foundation.String
	_TextViewWillShowSharingServicePickerForItems func(textView ITextView, servicePicker ISharingServicePicker, items foundation.foundation.INSArray) ISharingServicePicker
	_TextViewWritablePasteboardTypesForCellAtIndex func(view ITextView, cell unsafe.Pointer, charIndex uint) []string
	_TextViewWriteCellAtIndexToPasteboardType func(view ITextView, cell unsafe.Pointer, charIndex uint, pboard IPasteboard, type_ PasteboardType) bool
	_TextViewDidChangeSelection func(notification foundation.foundation.INSNotification)
	_TextViewDidChangeTypingAttributes func(notification foundation.foundation.INSNotification)
	_UndoManagerForTextView func(view ITextView) foundation.UndoManager
	_TextViewClickedOnCellInRect func(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect)
	_TextViewClickedOnLink func(textView ITextView, link objectivec.IObject) bool
	_TextViewDoubleClickedOnCellInRect func(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect)
	_TextViewDraggedCellInRectEvent func(view ITextView, cell unsafe.Pointer, rect corefoundation.CGRect, event IEvent)
}

// SetTextViewShouldSetSpellingStateRange sets the handler for the TextViewShouldSetSpellingStateRange delegate method.
//
// Sent when the spelling state is changed.
func (d *TextViewDelegate) SetTextViewShouldSetSpellingStateRange(f func(textView ITextView, value int, affectedCharRange foundation.Range) int) {
	d._TextViewShouldSetSpellingStateRange = f
}

// SetTextViewWritingToolsIgnoredRangesInEnclosingRange sets the handler for the TextViewWritingToolsIgnoredRangesInEnclosingRange delegate method.
func (d *TextViewDelegate) SetTextViewWritingToolsIgnoredRangesInEnclosingRange(f func(textView ITextView, enclosingRange foundation.Range) []foundation.Value) {
	d._TextViewWritingToolsIgnoredRangesInEnclosingRange = f
}

// SetTextViewWritingToolsDidEnd sets the handler for the TextViewWritingToolsDidEnd delegate method.
func (d *TextViewDelegate) SetTextViewWritingToolsDidEnd(f func(textView ITextView)) {
	d._TextViewWritingToolsDidEnd = f
}

// SetTextViewWritingToolsWillBegin sets the handler for the TextViewWritingToolsWillBegin delegate method.
func (d *TextViewDelegate) SetTextViewWritingToolsWillBegin(f func(textView ITextView)) {
	d._TextViewWritingToolsWillBegin = f
}

// SetTextViewCandidatesForSelectedRange sets the handler for the TextViewCandidatesForSelectedRange delegate method.
//
// Returns an array of text objects to include in a text selection.
func (d *TextViewDelegate) SetTextViewCandidatesForSelectedRange(f func(textView ITextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult) {
	d._TextViewCandidatesForSelectedRange = f
}

// SetTextViewClickedOnCellInRectAtIndex sets the handler for the TextViewClickedOnCellInRectAtIndex delegate method.
//
// Sent when the user clicks a cell.
func (d *TextViewDelegate) SetTextViewClickedOnCellInRectAtIndex(f func(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect, charIndex uint)) {
	d._TextViewClickedOnCellInRectAtIndex = f
}

// SetTextViewClickedOnLinkAtIndex sets the handler for the TextViewClickedOnLinkAtIndex delegate method.
//
// Sent after the user clicks a link.
func (d *TextViewDelegate) SetTextViewClickedOnLinkAtIndex(f func(textView ITextView, link objectivec.IObject, charIndex uint) bool) {
	d._TextViewClickedOnLinkAtIndex = f
}

// SetTextViewCompletionsForPartialWordRangeIndexOfSelectedItem sets the handler for the TextViewCompletionsForPartialWordRangeIndexOfSelectedItem delegate method.
//
// Returns the actual completions for a partial word.
func (d *TextViewDelegate) SetTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(f func(textView ITextView, words []string, charRange foundation.Range, index int) []string) {
	d._TextViewCompletionsForPartialWordRangeIndexOfSelectedItem = f
}

// SetTextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount sets the handler for the TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount delegate method.
//
// Invoked to allow the delegate to modify the text checking results after checking has occurred.
func (d *TextViewDelegate) SetTextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount(f func(view ITextView, range_ foundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.TextCheckingResult) {
	d._TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount = f
}

// SetTextViewDoCommandBySelector sets the handler for the TextViewDoCommandBySelector delegate method.
//
// Sent to allow the delegate to perform the command for the text view.
func (d *TextViewDelegate) SetTextViewDoCommandBySelector(f func(textView ITextView, commandSelector objc.SEL) bool) {
	d._TextViewDoCommandBySelector = f
}

// SetTextViewDoubleClickedOnCellInRectAtIndex sets the handler for the TextViewDoubleClickedOnCellInRectAtIndex delegate method.
//
// Sent when the user double-clicks a cell.
func (d *TextViewDelegate) SetTextViewDoubleClickedOnCellInRectAtIndex(f func(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect, charIndex uint)) {
	d._TextViewDoubleClickedOnCellInRectAtIndex = f
}

// SetTextViewDraggedCellInRectEventAtIndex sets the handler for the TextViewDraggedCellInRectEventAtIndex delegate method.
//
// Sent when the user attempts to drag a cell.
func (d *TextViewDelegate) SetTextViewDraggedCellInRectEventAtIndex(f func(view ITextView, cell unsafe.Pointer, rect corefoundation.CGRect, event IEvent, charIndex uint)) {
	d._TextViewDraggedCellInRectEventAtIndex = f
}

// SetTextViewMenuForEventAtIndex sets the handler for the TextViewMenuForEventAtIndex delegate method.
//
// Allows delegate to control the context menu returned by the text view.
func (d *TextViewDelegate) SetTextViewMenuForEventAtIndex(f func(view ITextView, menu IMenu, event IEvent, charIndex uint) IMenu) {
	d._TextViewMenuForEventAtIndex = f
}

// SetTextViewShouldChangeTextInRangeReplacementString sets the handler for the TextViewShouldChangeTextInRangeReplacementString delegate method.
//
// Sent when a text view needs to determine if text in a specified range should be changed.
func (d *TextViewDelegate) SetTextViewShouldChangeTextInRangeReplacementString(f func(textView ITextView, affectedCharRange foundation.Range, replacementString foundation.foundation.INSString) bool) {
	d._TextViewShouldChangeTextInRangeReplacementString = f
}

// SetTextViewShouldChangeTextInRangesReplacementStrings sets the handler for the TextViewShouldChangeTextInRangesReplacementStrings delegate method.
//
// Sent when a text view needs to determine if text in an array of specified ranges should be changed.
func (d *TextViewDelegate) SetTextViewShouldChangeTextInRangesReplacementStrings(f func(textView ITextView, affectedRanges []foundation.Value, replacementStrings []string) bool) {
	d._TextViewShouldChangeTextInRangesReplacementStrings = f
}

// SetTextViewShouldChangeTypingAttributesToAttributes sets the handler for the TextViewShouldChangeTypingAttributesToAttributes delegate method.
//
// Sent when the typing attributes are changed.
func (d *TextViewDelegate) SetTextViewShouldChangeTypingAttributesToAttributes(f func(textView ITextView, oldTypingAttributes foundation.IDictionary, newTypingAttributes foundation.IDictionary) foundation.IDictionary) {
	d._TextViewShouldChangeTypingAttributesToAttributes = f
}

// SetTextViewShouldSelectCandidateAtIndex sets the handler for the TextViewShouldSelectCandidateAtIndex delegate method.
//
// Returns a Boolean value that indicates whether to select the text object at the index.
func (d *TextViewDelegate) SetTextViewShouldSelectCandidateAtIndex(f func(textView ITextView, index uint) bool) {
	d._TextViewShouldSelectCandidateAtIndex = f
}

// SetTextViewShouldUpdateTouchBarItemIdentifiers sets the handler for the TextViewShouldUpdateTouchBarItemIdentifiers delegate method.
//
// Returns and array of touch bar elements for the framework to update.
func (d *TextViewDelegate) SetTextViewShouldUpdateTouchBarItemIdentifiers(f func(textView ITextView, identifiers []string) []string) {
	d._TextViewShouldUpdateTouchBarItemIdentifiers = f
}

// SetTextViewURLForContentsOfTextAttachmentAtIndex sets the handler for the TextViewURLForContentsOfTextAttachmentAtIndex delegate method.
//
// Returns a URL representing the document contents for a text attachment.
func (d *TextViewDelegate) SetTextViewURLForContentsOfTextAttachmentAtIndex(f func(textView ITextView, textAttachment ITextAttachment, charIndex uint) foundation.URL) {
	d._TextViewURLForContentsOfTextAttachmentAtIndex = f
}

// SetTextViewWillChangeSelectionFromCharacterRangeToCharacterRange sets the handler for the TextViewWillChangeSelectionFromCharacterRangeToCharacterRange delegate method.
//
// Returns the actual range to select.
func (d *TextViewDelegate) SetTextViewWillChangeSelectionFromCharacterRangeToCharacterRange(f func(textView ITextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range) {
	d._TextViewWillChangeSelectionFromCharacterRangeToCharacterRange = f
}

// SetTextViewWillChangeSelectionFromCharacterRangesToCharacterRanges sets the handler for the TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges delegate method.
//
// Returns the actual character ranges to select.
func (d *TextViewDelegate) SetTextViewWillChangeSelectionFromCharacterRangesToCharacterRanges(f func(textView ITextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.Value) {
	d._TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges = f
}

// SetTextViewWillCheckTextInRangeOptionsTypes sets the handler for the TextViewWillCheckTextInRangeOptionsTypes delegate method.
//
// Invoked to allow the delegate to modify the text checking process before it occurs.
func (d *TextViewDelegate) SetTextViewWillCheckTextInRangeOptionsTypes(f func(view ITextView, range_ foundation.Range, options foundation.IDictionary, checkingTypes TextCheckingTypes /* not a class type */) foundation.IDictionary) {
	d._TextViewWillCheckTextInRangeOptionsTypes = f
}

// SetTextViewWillDisplayToolTipForCharacterAtIndex sets the handler for the TextViewWillDisplayToolTipForCharacterAtIndex delegate method.
//
// Returns the actual tooltip to display.
func (d *TextViewDelegate) SetTextViewWillDisplayToolTipForCharacterAtIndex(f func(textView ITextView, tooltip foundation.foundation.INSString, characterIndex uint) foundation.String) {
	d._TextViewWillDisplayToolTipForCharacterAtIndex = f
}

// SetTextViewWillShowSharingServicePickerForItems sets the handler for the TextViewWillShowSharingServicePickerForItems delegate method.
//
// Returns a sharing service picker for the current selection.
func (d *TextViewDelegate) SetTextViewWillShowSharingServicePickerForItems(f func(textView ITextView, servicePicker ISharingServicePicker, items foundation.foundation.INSArray) ISharingServicePicker) {
	d._TextViewWillShowSharingServicePickerForItems = f
}

// SetTextViewWritablePasteboardTypesForCellAtIndex sets the handler for the TextViewWritablePasteboardTypesForCellAtIndex delegate method.
//
// Returns the writable pasteboard types for a given cell.
func (d *TextViewDelegate) SetTextViewWritablePasteboardTypesForCellAtIndex(f func(view ITextView, cell unsafe.Pointer, charIndex uint) []string) {
	d._TextViewWritablePasteboardTypesForCellAtIndex = f
}

// SetTextViewWriteCellAtIndexToPasteboardType sets the handler for the TextViewWriteCellAtIndexToPasteboardType delegate method.
//
// Returns whether data of the specified type for the given cell could be written to the specified pasteboard.
func (d *TextViewDelegate) SetTextViewWriteCellAtIndexToPasteboardType(f func(view ITextView, cell unsafe.Pointer, charIndex uint, pboard IPasteboard, type_ PasteboardType) bool) {
	d._TextViewWriteCellAtIndexToPasteboardType = f
}

// SetTextViewDidChangeSelection sets the handler for the TextViewDidChangeSelection delegate method.
//
// Sent when the selection changes in the text view.
func (d *TextViewDelegate) SetTextViewDidChangeSelection(f func(notification foundation.foundation.INSNotification)) {
	d._TextViewDidChangeSelection = f
}

// SetTextViewDidChangeTypingAttributes sets the handler for the TextViewDidChangeTypingAttributes delegate method.
//
// Sent when a text view’s typing attributes change.
func (d *TextViewDelegate) SetTextViewDidChangeTypingAttributes(f func(notification foundation.foundation.INSNotification)) {
	d._TextViewDidChangeTypingAttributes = f
}

// SetUndoManagerForTextView sets the handler for the UndoManagerForTextView delegate method.
//
// Returns the undo manager for the specified text view.
func (d *TextViewDelegate) SetUndoManagerForTextView(f func(view ITextView) foundation.UndoManager) {
	d._UndoManagerForTextView = f
}

// SetTextViewClickedOnCellInRect sets the handler for the TextViewClickedOnCellInRect delegate method.
//
// Sent when the user clicks a cell.
func (d *TextViewDelegate) SetTextViewClickedOnCellInRect(f func(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect)) {
	d._TextViewClickedOnCellInRect = f
}

// SetTextViewClickedOnLink sets the handler for the TextViewClickedOnLink delegate method.
//
// Sent after the user clicks on a link.
func (d *TextViewDelegate) SetTextViewClickedOnLink(f func(textView ITextView, link objectivec.IObject) bool) {
	d._TextViewClickedOnLink = f
}

// SetTextViewDoubleClickedOnCellInRect sets the handler for the TextViewDoubleClickedOnCellInRect delegate method.
//
// Sent when the user double-clicks a cell.
func (d *TextViewDelegate) SetTextViewDoubleClickedOnCellInRect(f func(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect)) {
	d._TextViewDoubleClickedOnCellInRect = f
}

// SetTextViewDraggedCellInRectEvent sets the handler for the TextViewDraggedCellInRectEvent delegate method.
//
// Sent when the user attempts to drag a cell.
func (d *TextViewDelegate) SetTextViewDraggedCellInRectEvent(f func(view ITextView, cell unsafe.Pointer, rect corefoundation.CGRect, event IEvent)) {
	d._TextViewDraggedCellInRectEvent = f
}

// TextViewShouldSetSpellingStateRange implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewShouldSetSpellingStateRange(textView ITextView, value int, affectedCharRange foundation.Range) int {
	if d._TextViewShouldSetSpellingStateRange != nil {
		return d._TextViewShouldSetSpellingStateRange(textView, value, affectedCharRange)
	}
	var zero int
	return zero
}

// HasTextViewShouldSetSpellingStateRange returns true if a handler for TextViewShouldSetSpellingStateRange has been set.
func (d *TextViewDelegate) HasTextViewShouldSetSpellingStateRange() bool {
	return d._TextViewShouldSetSpellingStateRange != nil
}

// TextViewWritingToolsIgnoredRangesInEnclosingRange implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewWritingToolsIgnoredRangesInEnclosingRange(textView ITextView, enclosingRange foundation.Range) []foundation.Value {
	if d._TextViewWritingToolsIgnoredRangesInEnclosingRange != nil {
		return d._TextViewWritingToolsIgnoredRangesInEnclosingRange(textView, enclosingRange)
	}
	var zero []foundation.Value
	return zero
}

// HasTextViewWritingToolsIgnoredRangesInEnclosingRange returns true if a handler for TextViewWritingToolsIgnoredRangesInEnclosingRange has been set.
func (d *TextViewDelegate) HasTextViewWritingToolsIgnoredRangesInEnclosingRange() bool {
	return d._TextViewWritingToolsIgnoredRangesInEnclosingRange != nil
}

// TextViewWritingToolsDidEnd implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewWritingToolsDidEnd(textView ITextView) {
	if d._TextViewWritingToolsDidEnd != nil {
		d._TextViewWritingToolsDidEnd(textView)
	}
}

// HasTextViewWritingToolsDidEnd returns true if a handler for TextViewWritingToolsDidEnd has been set.
func (d *TextViewDelegate) HasTextViewWritingToolsDidEnd() bool {
	return d._TextViewWritingToolsDidEnd != nil
}

// TextViewWritingToolsWillBegin implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewWritingToolsWillBegin(textView ITextView) {
	if d._TextViewWritingToolsWillBegin != nil {
		d._TextViewWritingToolsWillBegin(textView)
	}
}

// HasTextViewWritingToolsWillBegin returns true if a handler for TextViewWritingToolsWillBegin has been set.
func (d *TextViewDelegate) HasTextViewWritingToolsWillBegin() bool {
	return d._TextViewWritingToolsWillBegin != nil
}

// TextViewCandidatesForSelectedRange implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewCandidatesForSelectedRange(textView ITextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	if d._TextViewCandidatesForSelectedRange != nil {
		return d._TextViewCandidatesForSelectedRange(textView, candidates, selectedRange)
	}
	var zero []foundation.TextCheckingResult
	return zero
}

// HasTextViewCandidatesForSelectedRange returns true if a handler for TextViewCandidatesForSelectedRange has been set.
func (d *TextViewDelegate) HasTextViewCandidatesForSelectedRange() bool {
	return d._TextViewCandidatesForSelectedRange != nil
}

// TextViewClickedOnCellInRectAtIndex implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewClickedOnCellInRectAtIndex(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect, charIndex uint) {
	if d._TextViewClickedOnCellInRectAtIndex != nil {
		d._TextViewClickedOnCellInRectAtIndex(textView, cell, cellFrame, charIndex)
	}
}

// HasTextViewClickedOnCellInRectAtIndex returns true if a handler for TextViewClickedOnCellInRectAtIndex has been set.
func (d *TextViewDelegate) HasTextViewClickedOnCellInRectAtIndex() bool {
	return d._TextViewClickedOnCellInRectAtIndex != nil
}

// TextViewClickedOnLinkAtIndex implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewClickedOnLinkAtIndex(textView ITextView, link objectivec.IObject, charIndex uint) bool {
	if d._TextViewClickedOnLinkAtIndex != nil {
		return d._TextViewClickedOnLinkAtIndex(textView, link, charIndex)
	}
	var zero bool
	return zero
}

// HasTextViewClickedOnLinkAtIndex returns true if a handler for TextViewClickedOnLinkAtIndex has been set.
func (d *TextViewDelegate) HasTextViewClickedOnLinkAtIndex() bool {
	return d._TextViewClickedOnLinkAtIndex != nil
}

// TextViewCompletionsForPartialWordRangeIndexOfSelectedItem implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewCompletionsForPartialWordRangeIndexOfSelectedItem(textView ITextView, words []string, charRange foundation.Range, index int) []string {
	if d._TextViewCompletionsForPartialWordRangeIndexOfSelectedItem != nil {
		return d._TextViewCompletionsForPartialWordRangeIndexOfSelectedItem(textView, words, charRange, index)
	}
	var zero []string
	return zero
}

// HasTextViewCompletionsForPartialWordRangeIndexOfSelectedItem returns true if a handler for TextViewCompletionsForPartialWordRangeIndexOfSelectedItem has been set.
func (d *TextViewDelegate) HasTextViewCompletionsForPartialWordRangeIndexOfSelectedItem() bool {
	return d._TextViewCompletionsForPartialWordRangeIndexOfSelectedItem != nil
}

// TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount(view ITextView, range_ foundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.TextCheckingResult {
	if d._TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount != nil {
		return d._TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount(view, range_, checkingTypes, options, results, orthography, wordCount)
	}
	var zero []foundation.TextCheckingResult
	return zero
}

// HasTextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount returns true if a handler for TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount has been set.
func (d *TextViewDelegate) HasTextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount() bool {
	return d._TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount != nil
}

// TextViewDoCommandBySelector implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewDoCommandBySelector(textView ITextView, commandSelector objc.SEL) bool {
	if d._TextViewDoCommandBySelector != nil {
		return d._TextViewDoCommandBySelector(textView, commandSelector)
	}
	var zero bool
	return zero
}

// HasTextViewDoCommandBySelector returns true if a handler for TextViewDoCommandBySelector has been set.
func (d *TextViewDelegate) HasTextViewDoCommandBySelector() bool {
	return d._TextViewDoCommandBySelector != nil
}

// TextViewDoubleClickedOnCellInRectAtIndex implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewDoubleClickedOnCellInRectAtIndex(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect, charIndex uint) {
	if d._TextViewDoubleClickedOnCellInRectAtIndex != nil {
		d._TextViewDoubleClickedOnCellInRectAtIndex(textView, cell, cellFrame, charIndex)
	}
}

// HasTextViewDoubleClickedOnCellInRectAtIndex returns true if a handler for TextViewDoubleClickedOnCellInRectAtIndex has been set.
func (d *TextViewDelegate) HasTextViewDoubleClickedOnCellInRectAtIndex() bool {
	return d._TextViewDoubleClickedOnCellInRectAtIndex != nil
}

// TextViewDraggedCellInRectEventAtIndex implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewDraggedCellInRectEventAtIndex(view ITextView, cell unsafe.Pointer, rect corefoundation.CGRect, event IEvent, charIndex uint) {
	if d._TextViewDraggedCellInRectEventAtIndex != nil {
		d._TextViewDraggedCellInRectEventAtIndex(view, cell, rect, event, charIndex)
	}
}

// HasTextViewDraggedCellInRectEventAtIndex returns true if a handler for TextViewDraggedCellInRectEventAtIndex has been set.
func (d *TextViewDelegate) HasTextViewDraggedCellInRectEventAtIndex() bool {
	return d._TextViewDraggedCellInRectEventAtIndex != nil
}

// TextViewMenuForEventAtIndex implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewMenuForEventAtIndex(view ITextView, menu IMenu, event IEvent, charIndex uint) IMenu {
	if d._TextViewMenuForEventAtIndex != nil {
		return d._TextViewMenuForEventAtIndex(view, menu, event, charIndex)
	}
	var zero IMenu
	return zero
}

// HasTextViewMenuForEventAtIndex returns true if a handler for TextViewMenuForEventAtIndex has been set.
func (d *TextViewDelegate) HasTextViewMenuForEventAtIndex() bool {
	return d._TextViewMenuForEventAtIndex != nil
}

// TextViewShouldChangeTextInRangeReplacementString implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewShouldChangeTextInRangeReplacementString(textView ITextView, affectedCharRange foundation.Range, replacementString foundation.foundation.INSString) bool {
	if d._TextViewShouldChangeTextInRangeReplacementString != nil {
		return d._TextViewShouldChangeTextInRangeReplacementString(textView, affectedCharRange, replacementString)
	}
	var zero bool
	return zero
}

// HasTextViewShouldChangeTextInRangeReplacementString returns true if a handler for TextViewShouldChangeTextInRangeReplacementString has been set.
func (d *TextViewDelegate) HasTextViewShouldChangeTextInRangeReplacementString() bool {
	return d._TextViewShouldChangeTextInRangeReplacementString != nil
}

// TextViewShouldChangeTextInRangesReplacementStrings implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewShouldChangeTextInRangesReplacementStrings(textView ITextView, affectedRanges []foundation.Value, replacementStrings []string) bool {
	if d._TextViewShouldChangeTextInRangesReplacementStrings != nil {
		return d._TextViewShouldChangeTextInRangesReplacementStrings(textView, affectedRanges, replacementStrings)
	}
	var zero bool
	return zero
}

// HasTextViewShouldChangeTextInRangesReplacementStrings returns true if a handler for TextViewShouldChangeTextInRangesReplacementStrings has been set.
func (d *TextViewDelegate) HasTextViewShouldChangeTextInRangesReplacementStrings() bool {
	return d._TextViewShouldChangeTextInRangesReplacementStrings != nil
}

// TextViewShouldChangeTypingAttributesToAttributes implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewShouldChangeTypingAttributesToAttributes(textView ITextView, oldTypingAttributes foundation.IDictionary, newTypingAttributes foundation.IDictionary) foundation.IDictionary {
	if d._TextViewShouldChangeTypingAttributesToAttributes != nil {
		return d._TextViewShouldChangeTypingAttributesToAttributes(textView, oldTypingAttributes, newTypingAttributes)
	}
	var zero foundation.IDictionary
	return zero
}

// HasTextViewShouldChangeTypingAttributesToAttributes returns true if a handler for TextViewShouldChangeTypingAttributesToAttributes has been set.
func (d *TextViewDelegate) HasTextViewShouldChangeTypingAttributesToAttributes() bool {
	return d._TextViewShouldChangeTypingAttributesToAttributes != nil
}

// TextViewShouldSelectCandidateAtIndex implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewShouldSelectCandidateAtIndex(textView ITextView, index uint) bool {
	if d._TextViewShouldSelectCandidateAtIndex != nil {
		return d._TextViewShouldSelectCandidateAtIndex(textView, index)
	}
	var zero bool
	return zero
}

// HasTextViewShouldSelectCandidateAtIndex returns true if a handler for TextViewShouldSelectCandidateAtIndex has been set.
func (d *TextViewDelegate) HasTextViewShouldSelectCandidateAtIndex() bool {
	return d._TextViewShouldSelectCandidateAtIndex != nil
}

// TextViewShouldUpdateTouchBarItemIdentifiers implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewShouldUpdateTouchBarItemIdentifiers(textView ITextView, identifiers []string) []string {
	if d._TextViewShouldUpdateTouchBarItemIdentifiers != nil {
		return d._TextViewShouldUpdateTouchBarItemIdentifiers(textView, identifiers)
	}
	var zero []string
	return zero
}

// HasTextViewShouldUpdateTouchBarItemIdentifiers returns true if a handler for TextViewShouldUpdateTouchBarItemIdentifiers has been set.
func (d *TextViewDelegate) HasTextViewShouldUpdateTouchBarItemIdentifiers() bool {
	return d._TextViewShouldUpdateTouchBarItemIdentifiers != nil
}

// TextViewURLForContentsOfTextAttachmentAtIndex implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewURLForContentsOfTextAttachmentAtIndex(textView ITextView, textAttachment ITextAttachment, charIndex uint) foundation.URL {
	if d._TextViewURLForContentsOfTextAttachmentAtIndex != nil {
		return d._TextViewURLForContentsOfTextAttachmentAtIndex(textView, textAttachment, charIndex)
	}
	var zero foundation.URL
	return zero
}

// HasTextViewURLForContentsOfTextAttachmentAtIndex returns true if a handler for TextViewURLForContentsOfTextAttachmentAtIndex has been set.
func (d *TextViewDelegate) HasTextViewURLForContentsOfTextAttachmentAtIndex() bool {
	return d._TextViewURLForContentsOfTextAttachmentAtIndex != nil
}

// TextViewWillChangeSelectionFromCharacterRangeToCharacterRange implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewWillChangeSelectionFromCharacterRangeToCharacterRange(textView ITextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range {
	if d._TextViewWillChangeSelectionFromCharacterRangeToCharacterRange != nil {
		return d._TextViewWillChangeSelectionFromCharacterRangeToCharacterRange(textView, oldSelectedCharRange, newSelectedCharRange)
	}
	var zero foundation.Range
	return zero
}

// HasTextViewWillChangeSelectionFromCharacterRangeToCharacterRange returns true if a handler for TextViewWillChangeSelectionFromCharacterRangeToCharacterRange has been set.
func (d *TextViewDelegate) HasTextViewWillChangeSelectionFromCharacterRangeToCharacterRange() bool {
	return d._TextViewWillChangeSelectionFromCharacterRangeToCharacterRange != nil
}

// TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges(textView ITextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.Value {
	if d._TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges != nil {
		return d._TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges(textView, oldSelectedCharRanges, newSelectedCharRanges)
	}
	var zero []foundation.Value
	return zero
}

// HasTextViewWillChangeSelectionFromCharacterRangesToCharacterRanges returns true if a handler for TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges has been set.
func (d *TextViewDelegate) HasTextViewWillChangeSelectionFromCharacterRangesToCharacterRanges() bool {
	return d._TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges != nil
}

// TextViewWillCheckTextInRangeOptionsTypes implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewWillCheckTextInRangeOptionsTypes(view ITextView, range_ foundation.Range, options foundation.IDictionary, checkingTypes TextCheckingTypes /* not a class type */) foundation.IDictionary {
	if d._TextViewWillCheckTextInRangeOptionsTypes != nil {
		return d._TextViewWillCheckTextInRangeOptionsTypes(view, range_, options, checkingTypes)
	}
	var zero foundation.IDictionary
	return zero
}

// HasTextViewWillCheckTextInRangeOptionsTypes returns true if a handler for TextViewWillCheckTextInRangeOptionsTypes has been set.
func (d *TextViewDelegate) HasTextViewWillCheckTextInRangeOptionsTypes() bool {
	return d._TextViewWillCheckTextInRangeOptionsTypes != nil
}

// TextViewWillDisplayToolTipForCharacterAtIndex implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewWillDisplayToolTipForCharacterAtIndex(textView ITextView, tooltip foundation.foundation.INSString, characterIndex uint) foundation.String {
	if d._TextViewWillDisplayToolTipForCharacterAtIndex != nil {
		return d._TextViewWillDisplayToolTipForCharacterAtIndex(textView, tooltip, characterIndex)
	}
	var zero foundation.String
	return zero
}

// HasTextViewWillDisplayToolTipForCharacterAtIndex returns true if a handler for TextViewWillDisplayToolTipForCharacterAtIndex has been set.
func (d *TextViewDelegate) HasTextViewWillDisplayToolTipForCharacterAtIndex() bool {
	return d._TextViewWillDisplayToolTipForCharacterAtIndex != nil
}

// TextViewWillShowSharingServicePickerForItems implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewWillShowSharingServicePickerForItems(textView ITextView, servicePicker ISharingServicePicker, items foundation.foundation.INSArray) ISharingServicePicker {
	if d._TextViewWillShowSharingServicePickerForItems != nil {
		return d._TextViewWillShowSharingServicePickerForItems(textView, servicePicker, items)
	}
	var zero ISharingServicePicker
	return zero
}

// HasTextViewWillShowSharingServicePickerForItems returns true if a handler for TextViewWillShowSharingServicePickerForItems has been set.
func (d *TextViewDelegate) HasTextViewWillShowSharingServicePickerForItems() bool {
	return d._TextViewWillShowSharingServicePickerForItems != nil
}

// TextViewWritablePasteboardTypesForCellAtIndex implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewWritablePasteboardTypesForCellAtIndex(view ITextView, cell unsafe.Pointer, charIndex uint) []string {
	if d._TextViewWritablePasteboardTypesForCellAtIndex != nil {
		return d._TextViewWritablePasteboardTypesForCellAtIndex(view, cell, charIndex)
	}
	var zero []string
	return zero
}

// HasTextViewWritablePasteboardTypesForCellAtIndex returns true if a handler for TextViewWritablePasteboardTypesForCellAtIndex has been set.
func (d *TextViewDelegate) HasTextViewWritablePasteboardTypesForCellAtIndex() bool {
	return d._TextViewWritablePasteboardTypesForCellAtIndex != nil
}

// TextViewWriteCellAtIndexToPasteboardType implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewWriteCellAtIndexToPasteboardType(view ITextView, cell unsafe.Pointer, charIndex uint, pboard IPasteboard, type_ PasteboardType) bool {
	if d._TextViewWriteCellAtIndexToPasteboardType != nil {
		return d._TextViewWriteCellAtIndexToPasteboardType(view, cell, charIndex, pboard, type_)
	}
	var zero bool
	return zero
}

// HasTextViewWriteCellAtIndexToPasteboardType returns true if a handler for TextViewWriteCellAtIndexToPasteboardType has been set.
func (d *TextViewDelegate) HasTextViewWriteCellAtIndexToPasteboardType() bool {
	return d._TextViewWriteCellAtIndexToPasteboardType != nil
}

// TextViewDidChangeSelection implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewDidChangeSelection(notification foundation.foundation.INSNotification) {
	if d._TextViewDidChangeSelection != nil {
		d._TextViewDidChangeSelection(notification)
	}
}

// HasTextViewDidChangeSelection returns true if a handler for TextViewDidChangeSelection has been set.
func (d *TextViewDelegate) HasTextViewDidChangeSelection() bool {
	return d._TextViewDidChangeSelection != nil
}

// TextViewDidChangeTypingAttributes implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewDidChangeTypingAttributes(notification foundation.foundation.INSNotification) {
	if d._TextViewDidChangeTypingAttributes != nil {
		d._TextViewDidChangeTypingAttributes(notification)
	}
}

// HasTextViewDidChangeTypingAttributes returns true if a handler for TextViewDidChangeTypingAttributes has been set.
func (d *TextViewDelegate) HasTextViewDidChangeTypingAttributes() bool {
	return d._TextViewDidChangeTypingAttributes != nil
}

// UndoManagerForTextView implements the PTextViewDelegate interface.
func (d *TextViewDelegate) UndoManagerForTextView(view ITextView) foundation.UndoManager {
	if d._UndoManagerForTextView != nil {
		return d._UndoManagerForTextView(view)
	}
	var zero foundation.UndoManager
	return zero
}

// HasUndoManagerForTextView returns true if a handler for UndoManagerForTextView has been set.
func (d *TextViewDelegate) HasUndoManagerForTextView() bool {
	return d._UndoManagerForTextView != nil
}

// TextViewClickedOnCellInRect implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewClickedOnCellInRect(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect) {
	if d._TextViewClickedOnCellInRect != nil {
		d._TextViewClickedOnCellInRect(textView, cell, cellFrame)
	}
}

// HasTextViewClickedOnCellInRect returns true if a handler for TextViewClickedOnCellInRect has been set.
func (d *TextViewDelegate) HasTextViewClickedOnCellInRect() bool {
	return d._TextViewClickedOnCellInRect != nil
}

// TextViewClickedOnLink implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewClickedOnLink(textView ITextView, link objectivec.IObject) bool {
	if d._TextViewClickedOnLink != nil {
		return d._TextViewClickedOnLink(textView, link)
	}
	var zero bool
	return zero
}

// HasTextViewClickedOnLink returns true if a handler for TextViewClickedOnLink has been set.
func (d *TextViewDelegate) HasTextViewClickedOnLink() bool {
	return d._TextViewClickedOnLink != nil
}

// TextViewDoubleClickedOnCellInRect implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewDoubleClickedOnCellInRect(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect) {
	if d._TextViewDoubleClickedOnCellInRect != nil {
		d._TextViewDoubleClickedOnCellInRect(textView, cell, cellFrame)
	}
}

// HasTextViewDoubleClickedOnCellInRect returns true if a handler for TextViewDoubleClickedOnCellInRect has been set.
func (d *TextViewDelegate) HasTextViewDoubleClickedOnCellInRect() bool {
	return d._TextViewDoubleClickedOnCellInRect != nil
}

// TextViewDraggedCellInRectEvent implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewDraggedCellInRectEvent(view ITextView, cell unsafe.Pointer, rect corefoundation.CGRect, event IEvent) {
	if d._TextViewDraggedCellInRectEvent != nil {
		d._TextViewDraggedCellInRectEvent(view, cell, rect, event)
	}
}

// HasTextViewDraggedCellInRectEvent returns true if a handler for TextViewDraggedCellInRectEvent has been set.
func (d *TextViewDelegate) HasTextViewDraggedCellInRectEvent() bool {
	return d._TextViewDraggedCellInRectEvent != nil
}

// TextViewDelegateObject wraps an existing Objective-C object that conforms to the PTextViewDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type TextViewDelegateObject struct {
	objectivec.Object
}

// NewTextViewDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSTextViewDelegate protocol.
func NewTextViewDelegateObject(obj objectivec.Object) *TextViewDelegateObject {
	return &TextViewDelegateObject{obj}
}

// Make sure TextViewDelegateObject implements PTextViewDelegate.
var _ PTextViewDelegate = (*TextViewDelegateObject)(nil)

// TextViewClickedOnCellInRect implements the PTextViewDelegate interface.
// This required method is always available on objects conforming to TextViewClickedOnCellInRect.
func (o *TextViewDelegateObject) TextViewClickedOnCellInRect(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect) {
	objc.Send[objc.ID](o.ID, objc.Sel("textView:clickedOnCell:inRect:"), textView, cell, cellFrame)
}

// TextViewClickedOnLink implements the PTextViewDelegate interface.
// This required method is always available on objects conforming to TextViewClickedOnLink.
func (o *TextViewDelegateObject) TextViewClickedOnLink(textView ITextView, link objectivec.IObject) bool {
	return objc.Send[bool](o.ID, objc.Sel("textView:clickedOnLink:"), textView, link)
}

// TextViewDoubleClickedOnCellInRect implements the PTextViewDelegate interface.
// This required method is always available on objects conforming to TextViewDoubleClickedOnCellInRect.
func (o *TextViewDelegateObject) TextViewDoubleClickedOnCellInRect(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect) {
	objc.Send[objc.ID](o.ID, objc.Sel("textView:doubleClickedOnCell:inRect:"), textView, cell, cellFrame)
}

// TextViewDraggedCellInRectEvent implements the PTextViewDelegate interface.
// This required method is always available on objects conforming to TextViewDraggedCellInRectEvent.
func (o *TextViewDelegateObject) TextViewDraggedCellInRectEvent(view ITextView, cell unsafe.Pointer, rect corefoundation.CGRect, event IEvent) {
	objc.Send[objc.ID](o.ID, objc.Sel("textView:draggedCell:inRect:event:"), view, cell, rect, event)
}

// TextViewShouldSetSpellingStateRange implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewShouldSetSpellingStateRange(textView ITextView, value int, affectedCharRange foundation.Range) int {
	return objc.Send[int](o.ID, objc.Sel("textView:shouldSetSpellingState:range:"), textView, value, affectedCharRange)
}

// HasTextViewShouldSetSpellingStateRange returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewShouldSetSpellingStateRange() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewWritingToolsIgnoredRangesInEnclosingRange implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewWritingToolsIgnoredRangesInEnclosingRange(textView ITextView, enclosingRange foundation.Range) []foundation.Value {
	return objc.Send[[]foundation.Value](o.ID, objc.Sel("textView:writingToolsIgnoredRangesInEnclosingRange:"), textView, enclosingRange)
}

// HasTextViewWritingToolsIgnoredRangesInEnclosingRange returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewWritingToolsIgnoredRangesInEnclosingRange() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewWritingToolsDidEnd implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewWritingToolsDidEnd(textView ITextView) {
	objc.Send[objc.ID](o.ID, objc.Sel("textViewWritingToolsDidEnd:"), textView)
}

// HasTextViewWritingToolsDidEnd returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewWritingToolsDidEnd() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewWritingToolsWillBegin implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewWritingToolsWillBegin(textView ITextView) {
	objc.Send[objc.ID](o.ID, objc.Sel("textViewWritingToolsWillBegin:"), textView)
}

// HasTextViewWritingToolsWillBegin returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewWritingToolsWillBegin() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewCandidatesForSelectedRange implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewCandidatesForSelectedRange(textView ITextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	return objc.Send[[]foundation.TextCheckingResult](o.ID, objc.Sel("textView:candidates:forSelectedRange:"), textView, candidates, selectedRange)
}

// HasTextViewCandidatesForSelectedRange returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewCandidatesForSelectedRange() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewClickedOnCellInRectAtIndex implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewClickedOnCellInRectAtIndex(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect, charIndex uint) {
	objc.Send[objc.ID](o.ID, objc.Sel("textView:clickedOnCell:inRect:atIndex:"), textView, cell, cellFrame, charIndex)
}

// HasTextViewClickedOnCellInRectAtIndex returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewClickedOnCellInRectAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewClickedOnLinkAtIndex implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewClickedOnLinkAtIndex(textView ITextView, link objectivec.IObject, charIndex uint) bool {
	return objc.Send[bool](o.ID, objc.Sel("textView:clickedOnLink:atIndex:"), textView, link, charIndex)
}

// HasTextViewClickedOnLinkAtIndex returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewClickedOnLinkAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewCompletionsForPartialWordRangeIndexOfSelectedItem implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewCompletionsForPartialWordRangeIndexOfSelectedItem(textView ITextView, words []string, charRange foundation.Range, index int) []string {
	return objc.Send[[]string](o.ID, objc.Sel("textView:completions:forPartialWordRange:indexOfSelectedItem:"), textView, words, charRange, index)
}

// HasTextViewCompletionsForPartialWordRangeIndexOfSelectedItem returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewCompletionsForPartialWordRangeIndexOfSelectedItem() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount(view ITextView, range_ foundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.TextCheckingResult {
	return objc.Send[[]foundation.TextCheckingResult](o.ID, objc.Sel("textView:didCheckTextInRange:types:options:results:orthography:wordCount:"), view, range_, checkingTypes, options, results, orthography, wordCount)
}

// HasTextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewDoCommandBySelector implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewDoCommandBySelector(textView ITextView, commandSelector objc.SEL) bool {
	return objc.Send[bool](o.ID, objc.Sel("textView:doCommandBySelector:"), textView, commandSelector)
}

// HasTextViewDoCommandBySelector returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewDoCommandBySelector() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewDoubleClickedOnCellInRectAtIndex implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewDoubleClickedOnCellInRectAtIndex(textView ITextView, cell unsafe.Pointer, cellFrame corefoundation.CGRect, charIndex uint) {
	objc.Send[objc.ID](o.ID, objc.Sel("textView:doubleClickedOnCell:inRect:atIndex:"), textView, cell, cellFrame, charIndex)
}

// HasTextViewDoubleClickedOnCellInRectAtIndex returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewDoubleClickedOnCellInRectAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewDraggedCellInRectEventAtIndex implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewDraggedCellInRectEventAtIndex(view ITextView, cell unsafe.Pointer, rect corefoundation.CGRect, event IEvent, charIndex uint) {
	objc.Send[objc.ID](o.ID, objc.Sel("textView:draggedCell:inRect:event:atIndex:"), view, cell, rect, event, charIndex)
}

// HasTextViewDraggedCellInRectEventAtIndex returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewDraggedCellInRectEventAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewMenuForEventAtIndex implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewMenuForEventAtIndex(view ITextView, menu IMenu, event IEvent, charIndex uint) IMenu {
	return objc.Send[IMenu](o.ID, objc.Sel("textView:menu:forEvent:atIndex:"), view, menu, event, charIndex)
}

// HasTextViewMenuForEventAtIndex returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewMenuForEventAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewShouldChangeTextInRangeReplacementString implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewShouldChangeTextInRangeReplacementString(textView ITextView, affectedCharRange foundation.Range, replacementString foundation.foundation.INSString) bool {
	return objc.Send[bool](o.ID, objc.Sel("textView:shouldChangeTextInRange:replacementString:"), textView, affectedCharRange, replacementString)
}

// HasTextViewShouldChangeTextInRangeReplacementString returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewShouldChangeTextInRangeReplacementString() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewShouldChangeTextInRangesReplacementStrings implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewShouldChangeTextInRangesReplacementStrings(textView ITextView, affectedRanges []foundation.Value, replacementStrings []string) bool {
	return objc.Send[bool](o.ID, objc.Sel("textView:shouldChangeTextInRanges:replacementStrings:"), textView, affectedRanges, replacementStrings)
}

// HasTextViewShouldChangeTextInRangesReplacementStrings returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewShouldChangeTextInRangesReplacementStrings() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewShouldChangeTypingAttributesToAttributes implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewShouldChangeTypingAttributesToAttributes(textView ITextView, oldTypingAttributes foundation.IDictionary, newTypingAttributes foundation.IDictionary) foundation.IDictionary {
	return objc.Send[foundation.IDictionary](o.ID, objc.Sel("textView:shouldChangeTypingAttributes:toAttributes:"), textView, oldTypingAttributes, newTypingAttributes)
}

// HasTextViewShouldChangeTypingAttributesToAttributes returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewShouldChangeTypingAttributesToAttributes() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewShouldSelectCandidateAtIndex implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewShouldSelectCandidateAtIndex(textView ITextView, index uint) bool {
	return objc.Send[bool](o.ID, objc.Sel("textView:shouldSelectCandidateAtIndex:"), textView, index)
}

// HasTextViewShouldSelectCandidateAtIndex returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewShouldSelectCandidateAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewShouldUpdateTouchBarItemIdentifiers implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewShouldUpdateTouchBarItemIdentifiers(textView ITextView, identifiers []string) []string {
	return objc.Send[[]string](o.ID, objc.Sel("textView:shouldUpdateTouchBarItemIdentifiers:"), textView, identifiers)
}

// HasTextViewShouldUpdateTouchBarItemIdentifiers returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewShouldUpdateTouchBarItemIdentifiers() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewURLForContentsOfTextAttachmentAtIndex implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewURLForContentsOfTextAttachmentAtIndex(textView ITextView, textAttachment ITextAttachment, charIndex uint) foundation.URL {
	return objc.Send[foundation.URL](o.ID, objc.Sel("textView:URLForContentsOfTextAttachment:atIndex:"), textView, textAttachment, charIndex)
}

// HasTextViewURLForContentsOfTextAttachmentAtIndex returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewURLForContentsOfTextAttachmentAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewWillChangeSelectionFromCharacterRangeToCharacterRange implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewWillChangeSelectionFromCharacterRangeToCharacterRange(textView ITextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range {
	return objc.Send[foundation.Range](o.ID, objc.Sel("textView:willChangeSelectionFromCharacterRange:toCharacterRange:"), textView, oldSelectedCharRange, newSelectedCharRange)
}

// HasTextViewWillChangeSelectionFromCharacterRangeToCharacterRange returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewWillChangeSelectionFromCharacterRangeToCharacterRange() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges(textView ITextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.Value {
	return objc.Send[[]foundation.Value](o.ID, objc.Sel("textView:willChangeSelectionFromCharacterRanges:toCharacterRanges:"), textView, oldSelectedCharRanges, newSelectedCharRanges)
}

// HasTextViewWillChangeSelectionFromCharacterRangesToCharacterRanges returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewWillChangeSelectionFromCharacterRangesToCharacterRanges() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewWillCheckTextInRangeOptionsTypes implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewWillCheckTextInRangeOptionsTypes(view ITextView, range_ foundation.Range, options foundation.IDictionary, checkingTypes TextCheckingTypes /* not a class type */) foundation.IDictionary {
	return objc.Send[foundation.IDictionary](o.ID, objc.Sel("textView:willCheckTextInRange:options:types:"), view, range_, options, checkingTypes)
}

// HasTextViewWillCheckTextInRangeOptionsTypes returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewWillCheckTextInRangeOptionsTypes() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewWillDisplayToolTipForCharacterAtIndex implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewWillDisplayToolTipForCharacterAtIndex(textView ITextView, tooltip foundation.foundation.INSString, characterIndex uint) foundation.String {
	return objc.Send[foundation.String](o.ID, objc.Sel("textView:willDisplayToolTip:forCharacterAtIndex:"), textView, tooltip, characterIndex)
}

// HasTextViewWillDisplayToolTipForCharacterAtIndex returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewWillDisplayToolTipForCharacterAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewWillShowSharingServicePickerForItems implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewWillShowSharingServicePickerForItems(textView ITextView, servicePicker ISharingServicePicker, items foundation.foundation.INSArray) ISharingServicePicker {
	return objc.Send[ISharingServicePicker](o.ID, objc.Sel("textView:willShowSharingServicePicker:forItems:"), textView, servicePicker, items)
}

// HasTextViewWillShowSharingServicePickerForItems returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewWillShowSharingServicePickerForItems() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewWritablePasteboardTypesForCellAtIndex implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewWritablePasteboardTypesForCellAtIndex(view ITextView, cell unsafe.Pointer, charIndex uint) []string {
	return objc.Send[[]string](o.ID, objc.Sel("textView:writablePasteboardTypesForCell:atIndex:"), view, cell, charIndex)
}

// HasTextViewWritablePasteboardTypesForCellAtIndex returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewWritablePasteboardTypesForCellAtIndex() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewWriteCellAtIndexToPasteboardType implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewWriteCellAtIndexToPasteboardType(view ITextView, cell unsafe.Pointer, charIndex uint, pboard IPasteboard, type_ PasteboardType) bool {
	return objc.Send[bool](o.ID, objc.Sel("textView:writeCell:atIndex:toPasteboard:type:"), view, cell, charIndex, pboard, type_)
}

// HasTextViewWriteCellAtIndexToPasteboardType returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewWriteCellAtIndexToPasteboardType() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewDidChangeSelection implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewDidChangeSelection(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("textViewDidChangeSelection:"), notification)
}

// HasTextViewDidChangeSelection returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewDidChangeSelection() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// TextViewDidChangeTypingAttributes implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) TextViewDidChangeTypingAttributes(notification foundation.foundation.INSNotification) {
	objc.Send[objc.ID](o.ID, objc.Sel("textViewDidChangeTypingAttributes:"), notification)
}

// HasTextViewDidChangeTypingAttributes returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasTextViewDidChangeTypingAttributes() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// UndoManagerForTextView implements the PTextViewDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *TextViewDelegateObject) UndoManagerForTextView(view ITextView) foundation.UndoManager {
	return objc.Send[foundation.UndoManager](o.ID, objc.Sel("undoManagerForTextView:"), view)
}

// HasUndoManagerForTextView returns true; this is a placeholder for optional method checks.
func (o *TextViewDelegateObject) HasUndoManagerForTextView() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
