// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
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
	TextViewClickedOnCellInRect(textView ITextView, cell objc.IObject, cellFrame objc.IObject /* cross-framework: Rect */)
	TextViewClickedOnLink(textView ITextView, link objc.IObject) bool
	TextViewDoubleClickedOnCellInRect(textView ITextView, cell objc.IObject, cellFrame objc.IObject /* cross-framework: Rect */)
	TextViewDraggedCellInRectEvent(view ITextView, cell objc.IObject, rect objc.IObject /* cross-framework: Rect */, event IEvent)
	// Optional methods
	TextViewShouldSetSpellingStateRange(textView ITextView, value int, affectedCharRange corefoundation.Range) int
	HasTextViewShouldSetSpellingStateRange() bool
	TextViewWritingToolsIgnoredRangesInEnclosingRange(textView ITextView, enclosingRange corefoundation.Range) []foundation.Value
	HasTextViewWritingToolsIgnoredRangesInEnclosingRange() bool
	TextViewWritingToolsDidEnd(textView ITextView)
	HasTextViewWritingToolsDidEnd() bool
	TextViewWritingToolsWillBegin(textView ITextView)
	HasTextViewWritingToolsWillBegin() bool
	TextViewCandidatesForSelectedRange(textView ITextView, candidates []foundation.TextCheckingResult, selectedRange corefoundation.Range) []foundation.TextCheckingResult
	HasTextViewCandidatesForSelectedRange() bool
	TextViewClickedOnCellInRectAtIndex(textView ITextView, cell objc.IObject, cellFrame objc.IObject /* cross-framework: Rect */, charIndex uint)
	HasTextViewClickedOnCellInRectAtIndex() bool
	TextViewClickedOnLinkAtIndex(textView ITextView, link objc.IObject, charIndex uint) bool
	HasTextViewClickedOnLinkAtIndex() bool
	TextViewCompletionsForPartialWordRangeIndexOfSelectedItem(textView ITextView, words []string, charRange corefoundation.Range, index int) []string
	HasTextViewCompletionsForPartialWordRangeIndexOfSelectedItem() bool
	TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount(view ITextView, range_ corefoundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.TextCheckingResult
	HasTextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount() bool
	TextViewDoCommandBySelector(textView ITextView, commandSelector objc.SEL) bool
	HasTextViewDoCommandBySelector() bool
	TextViewDoubleClickedOnCellInRectAtIndex(textView ITextView, cell objc.IObject, cellFrame objc.IObject /* cross-framework: Rect */, charIndex uint)
	HasTextViewDoubleClickedOnCellInRectAtIndex() bool
	TextViewDraggedCellInRectEventAtIndex(view ITextView, cell objc.IObject, rect objc.IObject /* cross-framework: Rect */, event IEvent, charIndex uint)
	HasTextViewDraggedCellInRectEventAtIndex() bool
	TextViewMenuForEventAtIndex(view ITextView, menu IMenu, event IEvent, charIndex uint) Menu
	HasTextViewMenuForEventAtIndex() bool
	TextViewShouldChangeTextInRangeReplacementString(textView ITextView, affectedCharRange corefoundation.Range, replacementString objc.IObject /* cross-framework: NSString */) bool
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
	TextViewWillChangeSelectionFromCharacterRangeToCharacterRange(textView ITextView, oldSelectedCharRange corefoundation.Range, newSelectedCharRange corefoundation.Range) corefoundation.Range
	HasTextViewWillChangeSelectionFromCharacterRangeToCharacterRange() bool
	TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges(textView ITextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.Value
	HasTextViewWillChangeSelectionFromCharacterRangesToCharacterRanges() bool
	TextViewWillCheckTextInRangeOptionsTypes(view ITextView, range_ corefoundation.Range, options foundation.IDictionary, checkingTypes TextCheckingTypes /* not a class type */) foundation.IDictionary
	HasTextViewWillCheckTextInRangeOptionsTypes() bool
	TextViewWillDisplayToolTipForCharacterAtIndex(textView ITextView, tooltip objc.IObject /* cross-framework: NSString */, characterIndex uint) foundation.String
	HasTextViewWillDisplayToolTipForCharacterAtIndex() bool
	TextViewWillShowSharingServicePickerForItems(textView ITextView, servicePicker ISharingServicePicker, items objc.IObject /* cross-framework: NSArray */) SharingServicePicker
	HasTextViewWillShowSharingServicePickerForItems() bool
	TextViewWritablePasteboardTypesForCellAtIndex(view ITextView, cell objc.IObject, charIndex uint) []string
	HasTextViewWritablePasteboardTypesForCellAtIndex() bool
	TextViewWriteCellAtIndexToPasteboardType(view ITextView, cell objc.IObject, charIndex uint, pboard IPasteboard, type_ objc.IObject /* cross-framework: PasteboardType */) bool
	HasTextViewWriteCellAtIndexToPasteboardType() bool
	TextViewDidChangeSelection(notification foundation.Notification)
	HasTextViewDidChangeSelection() bool
	TextViewDidChangeTypingAttributes(notification foundation.Notification)
	HasTextViewDidChangeTypingAttributes() bool
	UndoManagerForTextView(view ITextView) foundation.UndoManager
	HasUndoManagerForTextView() bool
}

// TextViewDelegate is a delegate implementation builder for the PTextViewDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TextViewDelegate struct {
	_TextViewShouldSetSpellingStateRange func(textView ITextView, value int, affectedCharRange corefoundation.Range) int
	_TextViewWritingToolsIgnoredRangesInEnclosingRange func(textView ITextView, enclosingRange corefoundation.Range) []foundation.Value
	_TextViewWritingToolsDidEnd func(textView ITextView)
	_TextViewWritingToolsWillBegin func(textView ITextView)
	_TextViewCandidatesForSelectedRange func(textView ITextView, candidates []foundation.TextCheckingResult, selectedRange corefoundation.Range) []foundation.TextCheckingResult
	_TextViewClickedOnCellInRectAtIndex func(textView ITextView, cell objc.IObject, cellFrame objc.IObject /* cross-framework: Rect */, charIndex uint)
	_TextViewClickedOnLinkAtIndex func(textView ITextView, link objc.IObject, charIndex uint) bool
	_TextViewCompletionsForPartialWordRangeIndexOfSelectedItem func(textView ITextView, words []string, charRange corefoundation.Range, index int) []string
	_TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount func(view ITextView, range_ corefoundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.TextCheckingResult
	_TextViewDoCommandBySelector func(textView ITextView, commandSelector objc.SEL) bool
	_TextViewDoubleClickedOnCellInRectAtIndex func(textView ITextView, cell objc.IObject, cellFrame objc.IObject /* cross-framework: Rect */, charIndex uint)
	_TextViewDraggedCellInRectEventAtIndex func(view ITextView, cell objc.IObject, rect objc.IObject /* cross-framework: Rect */, event IEvent, charIndex uint)
	_TextViewMenuForEventAtIndex func(view ITextView, menu IMenu, event IEvent, charIndex uint) Menu
	_TextViewShouldChangeTextInRangeReplacementString func(textView ITextView, affectedCharRange corefoundation.Range, replacementString objc.IObject /* cross-framework: NSString */) bool
	_TextViewShouldChangeTextInRangesReplacementStrings func(textView ITextView, affectedRanges []foundation.Value, replacementStrings []string) bool
	_TextViewShouldChangeTypingAttributesToAttributes func(textView ITextView, oldTypingAttributes foundation.IDictionary, newTypingAttributes foundation.IDictionary) foundation.IDictionary
	_TextViewShouldSelectCandidateAtIndex func(textView ITextView, index uint) bool
	_TextViewShouldUpdateTouchBarItemIdentifiers func(textView ITextView, identifiers []string) []string
	_TextViewURLForContentsOfTextAttachmentAtIndex func(textView ITextView, textAttachment ITextAttachment, charIndex uint) foundation.URL
	_TextViewWillChangeSelectionFromCharacterRangeToCharacterRange func(textView ITextView, oldSelectedCharRange corefoundation.Range, newSelectedCharRange corefoundation.Range) corefoundation.Range
	_TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges func(textView ITextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.Value
	_TextViewWillCheckTextInRangeOptionsTypes func(view ITextView, range_ corefoundation.Range, options foundation.IDictionary, checkingTypes TextCheckingTypes /* not a class type */) foundation.IDictionary
	_TextViewWillDisplayToolTipForCharacterAtIndex func(textView ITextView, tooltip objc.IObject /* cross-framework: NSString */, characterIndex uint) foundation.String
	_TextViewWillShowSharingServicePickerForItems func(textView ITextView, servicePicker ISharingServicePicker, items objc.IObject /* cross-framework: NSArray */) SharingServicePicker
	_TextViewWritablePasteboardTypesForCellAtIndex func(view ITextView, cell objc.IObject, charIndex uint) []string
	_TextViewWriteCellAtIndexToPasteboardType func(view ITextView, cell objc.IObject, charIndex uint, pboard IPasteboard, type_ objc.IObject /* cross-framework: PasteboardType */) bool
	_TextViewDidChangeSelection func(notification foundation.Notification)
	_TextViewDidChangeTypingAttributes func(notification foundation.Notification)
	_UndoManagerForTextView func(view ITextView) foundation.UndoManager
	_TextViewClickedOnCellInRect func(textView ITextView, cell objc.IObject, cellFrame objc.IObject /* cross-framework: Rect */)
	_TextViewClickedOnLink func(textView ITextView, link objc.IObject) bool
	_TextViewDoubleClickedOnCellInRect func(textView ITextView, cell objc.IObject, cellFrame objc.IObject /* cross-framework: Rect */)
	_TextViewDraggedCellInRectEvent func(view ITextView, cell objc.IObject, rect objc.IObject /* cross-framework: Rect */, event IEvent)
}

// SetTextViewShouldSetSpellingStateRange sets the handler for the TextViewShouldSetSpellingStateRange delegate method.
//
// Sent when the spelling state is changed.
func (d *TextViewDelegate) SetTextViewShouldSetSpellingStateRange(f func(textView ITextView, value int, affectedCharRange corefoundation.Range) int) {
	d._TextViewShouldSetSpellingStateRange = f
}

// SetTextViewWritingToolsIgnoredRangesInEnclosingRange sets the handler for the TextViewWritingToolsIgnoredRangesInEnclosingRange delegate method.
func (d *TextViewDelegate) SetTextViewWritingToolsIgnoredRangesInEnclosingRange(f func(textView ITextView, enclosingRange corefoundation.Range) []foundation.Value) {
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
func (d *TextViewDelegate) SetTextViewCandidatesForSelectedRange(f func(textView ITextView, candidates []foundation.TextCheckingResult, selectedRange corefoundation.Range) []foundation.TextCheckingResult) {
	d._TextViewCandidatesForSelectedRange = f
}

// SetTextViewClickedOnCellInRectAtIndex sets the handler for the TextViewClickedOnCellInRectAtIndex delegate method.
//
// Sent when the user clicks a cell.
func (d *TextViewDelegate) SetTextViewClickedOnCellInRectAtIndex(f func(textView ITextView, cell objc.IObject, cellFrame objc.IObject /* cross-framework: Rect */, charIndex uint)) {
	d._TextViewClickedOnCellInRectAtIndex = f
}

// SetTextViewClickedOnLinkAtIndex sets the handler for the TextViewClickedOnLinkAtIndex delegate method.
//
// Sent after the user clicks a link.
func (d *TextViewDelegate) SetTextViewClickedOnLinkAtIndex(f func(textView ITextView, link objc.IObject, charIndex uint) bool) {
	d._TextViewClickedOnLinkAtIndex = f
}

// SetTextViewCompletionsForPartialWordRangeIndexOfSelectedItem sets the handler for the TextViewCompletionsForPartialWordRangeIndexOfSelectedItem delegate method.
//
// Returns the actual completions for a partial word.
func (d *TextViewDelegate) SetTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(f func(textView ITextView, words []string, charRange corefoundation.Range, index int) []string) {
	d._TextViewCompletionsForPartialWordRangeIndexOfSelectedItem = f
}

// SetTextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount sets the handler for the TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount delegate method.
//
// Invoked to allow the delegate to modify the text checking results after checking has occurred.
func (d *TextViewDelegate) SetTextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount(f func(view ITextView, range_ corefoundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.TextCheckingResult) {
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
func (d *TextViewDelegate) SetTextViewDoubleClickedOnCellInRectAtIndex(f func(textView ITextView, cell objc.IObject, cellFrame objc.IObject /* cross-framework: Rect */, charIndex uint)) {
	d._TextViewDoubleClickedOnCellInRectAtIndex = f
}

// SetTextViewDraggedCellInRectEventAtIndex sets the handler for the TextViewDraggedCellInRectEventAtIndex delegate method.
//
// Sent when the user attempts to drag a cell.
func (d *TextViewDelegate) SetTextViewDraggedCellInRectEventAtIndex(f func(view ITextView, cell objc.IObject, rect objc.IObject /* cross-framework: Rect */, event IEvent, charIndex uint)) {
	d._TextViewDraggedCellInRectEventAtIndex = f
}

// SetTextViewMenuForEventAtIndex sets the handler for the TextViewMenuForEventAtIndex delegate method.
//
// Allows delegate to control the context menu returned by the text view.
func (d *TextViewDelegate) SetTextViewMenuForEventAtIndex(f func(view ITextView, menu IMenu, event IEvent, charIndex uint) Menu) {
	d._TextViewMenuForEventAtIndex = f
}

// SetTextViewShouldChangeTextInRangeReplacementString sets the handler for the TextViewShouldChangeTextInRangeReplacementString delegate method.
//
// Sent when a text view needs to determine if text in a specified range should be changed.
func (d *TextViewDelegate) SetTextViewShouldChangeTextInRangeReplacementString(f func(textView ITextView, affectedCharRange corefoundation.Range, replacementString objc.IObject /* cross-framework: NSString */) bool) {
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
func (d *TextViewDelegate) SetTextViewWillChangeSelectionFromCharacterRangeToCharacterRange(f func(textView ITextView, oldSelectedCharRange corefoundation.Range, newSelectedCharRange corefoundation.Range) corefoundation.Range) {
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
func (d *TextViewDelegate) SetTextViewWillCheckTextInRangeOptionsTypes(f func(view ITextView, range_ corefoundation.Range, options foundation.IDictionary, checkingTypes TextCheckingTypes /* not a class type */) foundation.IDictionary) {
	d._TextViewWillCheckTextInRangeOptionsTypes = f
}

// SetTextViewWillDisplayToolTipForCharacterAtIndex sets the handler for the TextViewWillDisplayToolTipForCharacterAtIndex delegate method.
//
// Returns the actual tooltip to display.
func (d *TextViewDelegate) SetTextViewWillDisplayToolTipForCharacterAtIndex(f func(textView ITextView, tooltip objc.IObject /* cross-framework: NSString */, characterIndex uint) foundation.String) {
	d._TextViewWillDisplayToolTipForCharacterAtIndex = f
}

// SetTextViewWillShowSharingServicePickerForItems sets the handler for the TextViewWillShowSharingServicePickerForItems delegate method.
//
// Returns a sharing service picker for the current selection.
func (d *TextViewDelegate) SetTextViewWillShowSharingServicePickerForItems(f func(textView ITextView, servicePicker ISharingServicePicker, items objc.IObject /* cross-framework: NSArray */) SharingServicePicker) {
	d._TextViewWillShowSharingServicePickerForItems = f
}

// SetTextViewWritablePasteboardTypesForCellAtIndex sets the handler for the TextViewWritablePasteboardTypesForCellAtIndex delegate method.
//
// Returns the writable pasteboard types for a given cell.
func (d *TextViewDelegate) SetTextViewWritablePasteboardTypesForCellAtIndex(f func(view ITextView, cell objc.IObject, charIndex uint) []string) {
	d._TextViewWritablePasteboardTypesForCellAtIndex = f
}

// SetTextViewWriteCellAtIndexToPasteboardType sets the handler for the TextViewWriteCellAtIndexToPasteboardType delegate method.
//
// Returns whether data of the specified type for the given cell could be written to the specified pasteboard.
func (d *TextViewDelegate) SetTextViewWriteCellAtIndexToPasteboardType(f func(view ITextView, cell objc.IObject, charIndex uint, pboard IPasteboard, type_ objc.IObject /* cross-framework: PasteboardType */) bool) {
	d._TextViewWriteCellAtIndexToPasteboardType = f
}

// SetTextViewDidChangeSelection sets the handler for the TextViewDidChangeSelection delegate method.
//
// Sent when the selection changes in the text view.
func (d *TextViewDelegate) SetTextViewDidChangeSelection(f func(notification foundation.Notification)) {
	d._TextViewDidChangeSelection = f
}

// SetTextViewDidChangeTypingAttributes sets the handler for the TextViewDidChangeTypingAttributes delegate method.
//
// Sent when a text view’s typing attributes change.
func (d *TextViewDelegate) SetTextViewDidChangeTypingAttributes(f func(notification foundation.Notification)) {
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
func (d *TextViewDelegate) SetTextViewClickedOnCellInRect(f func(textView ITextView, cell objc.IObject, cellFrame objc.IObject /* cross-framework: Rect */)) {
	d._TextViewClickedOnCellInRect = f
}

// SetTextViewClickedOnLink sets the handler for the TextViewClickedOnLink delegate method.
//
// Sent after the user clicks on a link.
func (d *TextViewDelegate) SetTextViewClickedOnLink(f func(textView ITextView, link objc.IObject) bool) {
	d._TextViewClickedOnLink = f
}

// SetTextViewDoubleClickedOnCellInRect sets the handler for the TextViewDoubleClickedOnCellInRect delegate method.
//
// Sent when the user double-clicks a cell.
func (d *TextViewDelegate) SetTextViewDoubleClickedOnCellInRect(f func(textView ITextView, cell objc.IObject, cellFrame objc.IObject /* cross-framework: Rect */)) {
	d._TextViewDoubleClickedOnCellInRect = f
}

// SetTextViewDraggedCellInRectEvent sets the handler for the TextViewDraggedCellInRectEvent delegate method.
//
// Sent when the user attempts to drag a cell.
func (d *TextViewDelegate) SetTextViewDraggedCellInRectEvent(f func(view ITextView, cell objc.IObject, rect objc.IObject /* cross-framework: Rect */, event IEvent)) {
	d._TextViewDraggedCellInRectEvent = f
}

// TextViewShouldSetSpellingStateRange implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewShouldSetSpellingStateRange(textView ITextView, value int, affectedCharRange corefoundation.Range) int {
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
func (d *TextViewDelegate) TextViewWritingToolsIgnoredRangesInEnclosingRange(textView ITextView, enclosingRange corefoundation.Range) []foundation.Value {
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
func (d *TextViewDelegate) TextViewCandidatesForSelectedRange(textView ITextView, candidates []foundation.TextCheckingResult, selectedRange corefoundation.Range) []foundation.TextCheckingResult {
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
func (d *TextViewDelegate) TextViewClickedOnCellInRectAtIndex(textView ITextView, cell objc.IObject, cellFrame objc.IObject /* cross-framework: Rect */, charIndex uint) {
	if d._TextViewClickedOnCellInRectAtIndex != nil {
		d._TextViewClickedOnCellInRectAtIndex(textView, cell, cellFrame, charIndex)
	}
}

// HasTextViewClickedOnCellInRectAtIndex returns true if a handler for TextViewClickedOnCellInRectAtIndex has been set.
func (d *TextViewDelegate) HasTextViewClickedOnCellInRectAtIndex() bool {
	return d._TextViewClickedOnCellInRectAtIndex != nil
}

// TextViewClickedOnLinkAtIndex implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewClickedOnLinkAtIndex(textView ITextView, link objc.IObject, charIndex uint) bool {
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
func (d *TextViewDelegate) TextViewCompletionsForPartialWordRangeIndexOfSelectedItem(textView ITextView, words []string, charRange corefoundation.Range, index int) []string {
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
func (d *TextViewDelegate) TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount(view ITextView, range_ corefoundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.TextCheckingResult {
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
func (d *TextViewDelegate) TextViewDoubleClickedOnCellInRectAtIndex(textView ITextView, cell objc.IObject, cellFrame objc.IObject /* cross-framework: Rect */, charIndex uint) {
	if d._TextViewDoubleClickedOnCellInRectAtIndex != nil {
		d._TextViewDoubleClickedOnCellInRectAtIndex(textView, cell, cellFrame, charIndex)
	}
}

// HasTextViewDoubleClickedOnCellInRectAtIndex returns true if a handler for TextViewDoubleClickedOnCellInRectAtIndex has been set.
func (d *TextViewDelegate) HasTextViewDoubleClickedOnCellInRectAtIndex() bool {
	return d._TextViewDoubleClickedOnCellInRectAtIndex != nil
}

// TextViewDraggedCellInRectEventAtIndex implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewDraggedCellInRectEventAtIndex(view ITextView, cell objc.IObject, rect objc.IObject /* cross-framework: Rect */, event IEvent, charIndex uint) {
	if d._TextViewDraggedCellInRectEventAtIndex != nil {
		d._TextViewDraggedCellInRectEventAtIndex(view, cell, rect, event, charIndex)
	}
}

// HasTextViewDraggedCellInRectEventAtIndex returns true if a handler for TextViewDraggedCellInRectEventAtIndex has been set.
func (d *TextViewDelegate) HasTextViewDraggedCellInRectEventAtIndex() bool {
	return d._TextViewDraggedCellInRectEventAtIndex != nil
}

// TextViewMenuForEventAtIndex implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewMenuForEventAtIndex(view ITextView, menu IMenu, event IEvent, charIndex uint) Menu {
	if d._TextViewMenuForEventAtIndex != nil {
		return d._TextViewMenuForEventAtIndex(view, menu, event, charIndex)
	}
	var zero Menu
	return zero
}

// HasTextViewMenuForEventAtIndex returns true if a handler for TextViewMenuForEventAtIndex has been set.
func (d *TextViewDelegate) HasTextViewMenuForEventAtIndex() bool {
	return d._TextViewMenuForEventAtIndex != nil
}

// TextViewShouldChangeTextInRangeReplacementString implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewShouldChangeTextInRangeReplacementString(textView ITextView, affectedCharRange corefoundation.Range, replacementString objc.IObject /* cross-framework: NSString */) bool {
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
func (d *TextViewDelegate) TextViewWillChangeSelectionFromCharacterRangeToCharacterRange(textView ITextView, oldSelectedCharRange corefoundation.Range, newSelectedCharRange corefoundation.Range) corefoundation.Range {
	if d._TextViewWillChangeSelectionFromCharacterRangeToCharacterRange != nil {
		return d._TextViewWillChangeSelectionFromCharacterRangeToCharacterRange(textView, oldSelectedCharRange, newSelectedCharRange)
	}
	var zero corefoundation.Range
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
func (d *TextViewDelegate) TextViewWillCheckTextInRangeOptionsTypes(view ITextView, range_ corefoundation.Range, options foundation.IDictionary, checkingTypes TextCheckingTypes /* not a class type */) foundation.IDictionary {
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
func (d *TextViewDelegate) TextViewWillDisplayToolTipForCharacterAtIndex(textView ITextView, tooltip objc.IObject /* cross-framework: NSString */, characterIndex uint) foundation.String {
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
func (d *TextViewDelegate) TextViewWillShowSharingServicePickerForItems(textView ITextView, servicePicker ISharingServicePicker, items objc.IObject /* cross-framework: NSArray */) SharingServicePicker {
	if d._TextViewWillShowSharingServicePickerForItems != nil {
		return d._TextViewWillShowSharingServicePickerForItems(textView, servicePicker, items)
	}
	var zero SharingServicePicker
	return zero
}

// HasTextViewWillShowSharingServicePickerForItems returns true if a handler for TextViewWillShowSharingServicePickerForItems has been set.
func (d *TextViewDelegate) HasTextViewWillShowSharingServicePickerForItems() bool {
	return d._TextViewWillShowSharingServicePickerForItems != nil
}

// TextViewWritablePasteboardTypesForCellAtIndex implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewWritablePasteboardTypesForCellAtIndex(view ITextView, cell objc.IObject, charIndex uint) []string {
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
func (d *TextViewDelegate) TextViewWriteCellAtIndexToPasteboardType(view ITextView, cell objc.IObject, charIndex uint, pboard IPasteboard, type_ objc.IObject /* cross-framework: PasteboardType */) bool {
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
func (d *TextViewDelegate) TextViewDidChangeSelection(notification foundation.Notification) {
	if d._TextViewDidChangeSelection != nil {
		d._TextViewDidChangeSelection(notification)
	}
}

// HasTextViewDidChangeSelection returns true if a handler for TextViewDidChangeSelection has been set.
func (d *TextViewDelegate) HasTextViewDidChangeSelection() bool {
	return d._TextViewDidChangeSelection != nil
}

// TextViewDidChangeTypingAttributes implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewDidChangeTypingAttributes(notification foundation.Notification) {
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
func (d *TextViewDelegate) TextViewClickedOnCellInRect(textView ITextView, cell objc.IObject, cellFrame objc.IObject /* cross-framework: Rect */) {
	if d._TextViewClickedOnCellInRect != nil {
		d._TextViewClickedOnCellInRect(textView, cell, cellFrame)
	}
}

// HasTextViewClickedOnCellInRect returns true if a handler for TextViewClickedOnCellInRect has been set.
func (d *TextViewDelegate) HasTextViewClickedOnCellInRect() bool {
	return d._TextViewClickedOnCellInRect != nil
}

// TextViewClickedOnLink implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewClickedOnLink(textView ITextView, link objc.IObject) bool {
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
func (d *TextViewDelegate) TextViewDoubleClickedOnCellInRect(textView ITextView, cell objc.IObject, cellFrame objc.IObject /* cross-framework: Rect */) {
	if d._TextViewDoubleClickedOnCellInRect != nil {
		d._TextViewDoubleClickedOnCellInRect(textView, cell, cellFrame)
	}
}

// HasTextViewDoubleClickedOnCellInRect returns true if a handler for TextViewDoubleClickedOnCellInRect has been set.
func (d *TextViewDelegate) HasTextViewDoubleClickedOnCellInRect() bool {
	return d._TextViewDoubleClickedOnCellInRect != nil
}

// TextViewDraggedCellInRectEvent implements the PTextViewDelegate interface.
func (d *TextViewDelegate) TextViewDraggedCellInRectEvent(view ITextView, cell objc.IObject, rect objc.IObject /* cross-framework: Rect */, event IEvent) {
	if d._TextViewDraggedCellInRectEvent != nil {
		d._TextViewDraggedCellInRectEvent(view, cell, rect, event)
	}
}

// HasTextViewDraggedCellInRectEvent returns true if a handler for TextViewDraggedCellInRectEvent has been set.
func (d *TextViewDelegate) HasTextViewDraggedCellInRectEvent() bool {
	return d._TextViewDraggedCellInRectEvent != nil
}
