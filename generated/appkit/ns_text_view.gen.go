// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TextView] class.
var (
	TextViewClass     _TextViewClass
	TextViewClassOnce sync.Once
)

func getTextViewClass() _TextViewClass {
	TextViewClassOnce.Do(func() {
		TextViewClass = _TextViewClass{objc.GetClass("NSTextView")}
	})
	return TextViewClass
}

type _TextViewClass struct {
	class objc.Class
}

// An interface definition for the [TextView] class.
type ITextView interface {
	IText
	// properties:
	LayoutManager() objc.IObject /* cross-framework: LayoutManager */
	RangeForUserCompletion() foundation.objc.IObject /* cross-framework: Range */
	SelectedRanges() []Value /* primitive/slice/pointer. */
	SetSelectedRanges(value []Value /* primitive/slice/pointer. */)
	TextContainer() ITextContainer
	SetTextContainer(value ITextContainer)
	TextContainerOrigin() coregraphics.CGPoint
	TextStorage() ITextStorage
	TypingAttributes() foundation.IDictionary /* already interface */
	SetTypingAttributes(value foundation.IDictionary /* already interface */)
	AcceptableDragTypes() objc.IObject /* cross-framework: PasteboardType */
	SetAcceptableDragTypes(value objc.IObject /* cross-framework: PasteboardType */)
	AcceptsGlyphInfo() bool /* primitive/slice/pointer. */
	SetAcceptsGlyphInfo(value bool /* primitive/slice/pointer. */)
	AllowedInputSourceLocales() string /* primitive/slice/pointer. */
	SetAllowedInputSourceLocales(value string /* primitive/slice/pointer. */)
	AllowedWritingToolsResultOptions() WritingToolsResultOptions
	SetAllowedWritingToolsResultOptions(value WritingToolsResultOptions)
	AllowsCharacterPickerTouchBarItem() bool /* primitive/slice/pointer. */
	SetAllowsCharacterPickerTouchBarItem(value bool /* primitive/slice/pointer. */)
	AllowsDocumentBackgroundColorChange() bool /* primitive/slice/pointer. */
	SetAllowsDocumentBackgroundColorChange(value bool /* primitive/slice/pointer. */)
	AllowsImageEditing() bool /* primitive/slice/pointer. */
	SetAllowsImageEditing(value bool /* primitive/slice/pointer. */)
	AllowsUndo() bool /* primitive/slice/pointer. */
	SetAllowsUndo(value bool /* primitive/slice/pointer. */)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	CandidateListTouchBarItem() objc.IObject /* cross-framework: CandidateListTouchBarItem */
	SetCandidateListTouchBarItem(value objc.IObject /* cross-framework: CandidateListTouchBarItem */)
	DefaultParagraphStyle() IParagraphStyle
	SetDefaultParagraphStyle(value IParagraphStyle)
	Delegate() TextViewDelegate /* not a class type */
	SetDelegate(value TextViewDelegate /* not a class type */)
	DisplaysLinkToolTips() bool /* primitive/slice/pointer. */
	SetDisplaysLinkToolTips(value bool /* primitive/slice/pointer. */)
	DrawsBackground() bool /* primitive/slice/pointer. */
	SetDrawsBackground(value bool /* primitive/slice/pointer. */)
	EnabledTextCheckingTypes() TextCheckingTypes /* not a class type */
	SetEnabledTextCheckingTypes(value TextCheckingTypes /* not a class type */)
	ImportsGraphics() bool /* primitive/slice/pointer. */
	SetImportsGraphics(value bool /* primitive/slice/pointer. */)
	InlinePredictionType() TextInputTraitType /* not a class type */
	SetInlinePredictionType(value TextInputTraitType /* not a class type */)
	InsertionPointColor() IColor
	SetInsertionPointColor(value IColor)
	IsAutomaticDashSubstitutionEnabled() bool /* primitive/slice/pointer. */
	SetIsAutomaticDashSubstitutionEnabled(value bool /* primitive/slice/pointer. */)
	IsAutomaticDataDetectionEnabled() bool /* primitive/slice/pointer. */
	SetIsAutomaticDataDetectionEnabled(value bool /* primitive/slice/pointer. */)
	IsAutomaticLinkDetectionEnabled() bool /* primitive/slice/pointer. */
	SetIsAutomaticLinkDetectionEnabled(value bool /* primitive/slice/pointer. */)
	IsAutomaticQuoteSubstitutionEnabled() bool /* primitive/slice/pointer. */
	SetIsAutomaticQuoteSubstitutionEnabled(value bool /* primitive/slice/pointer. */)
	IsAutomaticSpellingCorrectionEnabled() bool /* primitive/slice/pointer. */
	SetIsAutomaticSpellingCorrectionEnabled(value bool /* primitive/slice/pointer. */)
	IsAutomaticTextCompletionEnabled() bool /* primitive/slice/pointer. */
	SetIsAutomaticTextCompletionEnabled(value bool /* primitive/slice/pointer. */)
	IsAutomaticTextReplacementEnabled() bool /* primitive/slice/pointer. */
	SetIsAutomaticTextReplacementEnabled(value bool /* primitive/slice/pointer. */)
	IsCoalescingUndo() bool /* primitive/slice/pointer. */
	SetIsCoalescingUndo(value bool /* primitive/slice/pointer. */)
	IsContinuousSpellCheckingEnabled() bool /* primitive/slice/pointer. */
	SetIsContinuousSpellCheckingEnabled(value bool /* primitive/slice/pointer. */)
	IsEditable() bool /* primitive/slice/pointer. */
	SetIsEditable(value bool /* primitive/slice/pointer. */)
	IsFieldEditor() bool /* primitive/slice/pointer. */
	SetIsFieldEditor(value bool /* primitive/slice/pointer. */)
	IsGrammarCheckingEnabled() bool /* primitive/slice/pointer. */
	SetIsGrammarCheckingEnabled(value bool /* primitive/slice/pointer. */)
	IsIncrementalSearchingEnabled() bool /* primitive/slice/pointer. */
	SetIsIncrementalSearchingEnabled(value bool /* primitive/slice/pointer. */)
	IsRichText() bool /* primitive/slice/pointer. */
	SetIsRichText(value bool /* primitive/slice/pointer. */)
	IsRulerVisible() bool /* primitive/slice/pointer. */
	SetIsRulerVisible(value bool /* primitive/slice/pointer. */)
	IsSelectable() bool /* primitive/slice/pointer. */
	SetIsSelectable(value bool /* primitive/slice/pointer. */)
	IsWritingToolsActive() bool /* primitive/slice/pointer. */
	SetIsWritingToolsActive(value bool /* primitive/slice/pointer. */)
	LinkTextAttributes() coreml.objc.IObject /* cross-framework: Key */
	SetLinkTextAttributes(value coreml.objc.IObject /* cross-framework: Key */)
	MarkedTextAttributes() coreml.objc.IObject /* cross-framework: Key */
	SetMarkedTextAttributes(value coreml.objc.IObject /* cross-framework: Key */)
	MathExpressionCompletionType() TextInputTraitType /* not a class type */
	SetMathExpressionCompletionType(value TextInputTraitType /* not a class type */)
	RangeForUserCharacterAttributeChange() foundation.objc.IObject /* cross-framework: Range */
	SetRangeForUserCharacterAttributeChange(value foundation.objc.IObject /* cross-framework: Range */)
	RangeForUserParagraphAttributeChange() foundation.objc.IObject /* cross-framework: Range */
	SetRangeForUserParagraphAttributeChange(value foundation.objc.IObject /* cross-framework: Range */)
	RangeForUserTextChange() foundation.objc.IObject /* cross-framework: Range */
	SetRangeForUserTextChange(value foundation.objc.IObject /* cross-framework: Range */)
	RangesForUserCharacterAttributeChange() Value /* not a class type */
	SetRangesForUserCharacterAttributeChange(value Value /* not a class type */)
	RangesForUserParagraphAttributeChange() Value /* not a class type */
	SetRangesForUserParagraphAttributeChange(value Value /* not a class type */)
	RangesForUserTextChange() Value /* not a class type */
	SetRangesForUserTextChange(value Value /* not a class type */)
	ReadablePasteboardTypes() objc.IObject /* cross-framework: PasteboardType */
	SetReadablePasteboardTypes(value objc.IObject /* cross-framework: PasteboardType */)
	SelectedTextAttributes() coreml.objc.IObject /* cross-framework: Key */
	SetSelectedTextAttributes(value coreml.objc.IObject /* cross-framework: Key */)
	SelectionAffinity() SelectionAffinity /* not a class type */
	SetSelectionAffinity(value SelectionAffinity /* not a class type */)
	SelectionGranularity() SelectionGranularity /* not a class type */
	SetSelectionGranularity(value SelectionGranularity /* not a class type */)
	ShouldDrawInsertionPoint() bool /* primitive/slice/pointer. */
	SetShouldDrawInsertionPoint(value bool /* primitive/slice/pointer. */)
	SmartInsertDeleteEnabled() bool /* primitive/slice/pointer. */
	SetSmartInsertDeleteEnabled(value bool /* primitive/slice/pointer. */)
	SpellCheckerDocumentTag() int /* primitive/slice/pointer. */
	SetSpellCheckerDocumentTag(value int /* primitive/slice/pointer. */)
	TextContainerInset() coregraphics.CGSize
	SetTextContainerInset(value coregraphics.CGSize)
	TextContentStorage() ITextContentStorage
	SetTextContentStorage(value ITextContentStorage)
	TextHighlightAttributes() coreml.objc.IObject /* cross-framework: Key */
	SetTextHighlightAttributes(value coreml.objc.IObject /* cross-framework: Key */)
	TextLayoutManager() ITextLayoutManager
	SetTextLayoutManager(value ITextLayoutManager)
	UsesAdaptiveColorMappingForDarkAppearance() bool /* primitive/slice/pointer. */
	SetUsesAdaptiveColorMappingForDarkAppearance(value bool /* primitive/slice/pointer. */)
	UsesFindBar() bool /* primitive/slice/pointer. */
	SetUsesFindBar(value bool /* primitive/slice/pointer. */)
	UsesFindPanel() bool /* primitive/slice/pointer. */
	SetUsesFindPanel(value bool /* primitive/slice/pointer. */)
	UsesFontPanel() bool /* primitive/slice/pointer. */
	SetUsesFontPanel(value bool /* primitive/slice/pointer. */)
	UsesInspectorBar() bool /* primitive/slice/pointer. */
	SetUsesInspectorBar(value bool /* primitive/slice/pointer. */)
	UsesRolloverButtonForSelection() bool /* primitive/slice/pointer. */
	SetUsesRolloverButtonForSelection(value bool /* primitive/slice/pointer. */)
	UsesRuler() bool /* primitive/slice/pointer. */
	SetUsesRuler(value bool /* primitive/slice/pointer. */)
	WritablePasteboardTypes() objc.IObject /* cross-framework: PasteboardType */
	SetWritablePasteboardTypes(value objc.IObject /* cross-framework: PasteboardType */)
	WritingToolsBehavior() WritingToolsBehavior
	SetWritingToolsBehavior(value WritingToolsBehavior)
	// methods:
	AlignJustified(sender objectivec.IObject)
	ChangeAttributes(sender objectivec.IObject)
	ChangeColor(sender objectivec.IObject)
	HandleTextCheckingResultsForRangeTypesOptionsOrthographyWordCount(results []foundation.objc.IObject /* cross-framework TextCheckingResult */, range_ foundation.objc.IObject /* cross-framework Range */, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary /* already interface */, orthography objc.IObject /* cross-framework Orthography */, wordCount int /* primitive/slice/pointer. */)
	LoosenKerning(sender objectivec.IObject)
	LowerBaseline(sender objectivec.IObject)
	QuickLookPreviewableItemsInRanges(ranges []Value /* primitive/slice/pointer. */) []objc.ID /* already interface */
	RaiseBaseline(sender objectivec.IObject)
	RulerViewShouldRemoveMarker(ruler IRulerView, marker IRulerMarker) bool /* primitive/slice/pointer. */
	SetAlignmentRange(alignment TextAlignment, range_ foundation.objc.IObject /* cross-framework Range */)
	TightenKerning(sender objectivec.IObject)
	ToggleAutomaticTextReplacement(sender objectivec.IObject)
	ToggleSmartInsertDelete(sender objectivec.IObject)
	TurnOffKerning(sender objectivec.IObject)
	TurnOffLigatures(sender objectivec.IObject)
	UseAllLigatures(sender objectivec.IObject)
	UseStandardKerning(sender objectivec.IObject)
	UseStandardLigatures(sender objectivec.IObject)
}

// A view that draws text and handles user interactions with that text.
//
// The class is the front-end class to the AppKit text system. The class draws the text managed by the back-end components and handles user events to select and modify its text, in addition to supporting rich text, attachments, input management, and key binding, and marked text attributes. is the principal means to obtain a text object that caters to almost all needs for displaying and managing text at the user interface level. While is a subclass of the class — which declares the most general Cocoa interface to the text system — adds major features beyond the capabilities of . You can also do more powerful and more creative text manipulation (such as displaying text in a circle) using , , , and related classes. You’re more likely to use the class than . It’s also important to remember that conforms to a large number of protocols, the methods of which are available to instances of the class. communicates with its delegate through methods declared both by the and by its superclass’s protocol, . All delegation messages come from the first text view. In macOS 12 and later, if you explicitly call the property on a text view or text container, the framework reverts to a compatibility mode that uses . The text view also switches to this compatibility mode when it encounters text content that’s not yet supported, such as .


// A view that draws text and handles user interactions with that text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView
type TextView struct {
	Text
}

// TextViewFrom constructs a [TextView] from an unsafe.Pointer.
//
// A view that draws text and handles user interactions with that text.
func TextViewFrom(ptr unsafe.Pointer) TextView {
	return TextView{
		Text: TextFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TextViewClass) Alloc() TextView {
	rv := objc.Send[TextView](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextViewClass) New() TextView {
	rv := objc.Send[TextView](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextView) Init() TextView {
	rv := objc.Send[TextView](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextView) Autorelease() TextView {
	rv := objc.Send[TextView](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextView creates a new TextView instance.
func NewTextView() TextView {
	return getTextViewClass().New()
}



// Applies full justification to selected paragraphs (or all text, if the receiver is a plain text object).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/alignJustified(_:)
func (t_ TextView) AlignJustified(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("alignJustified:"), sender)
}


// Changes the attributes of the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/changeAttributes(_:)
func (t_ TextView) ChangeAttributes(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("changeAttributes:"), sender)
}


// Sets the color of the selected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/changeColor(_:)
func (t_ TextView) ChangeColor(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("changeColor:"), sender)
}


// Handles the text checking results returned by the text view
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/handleTextCheckingResults(_:forRange:types:options:orthography:wordCount:)
func (t_ TextView) HandleTextCheckingResultsForRangeTypesOptionsOrthographyWordCount(results []foundation.objc.IObject /* cross-framework TextCheckingResult */, range_ foundation.objc.IObject /* cross-framework Range */, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary /* already interface */, orthography objc.IObject /* cross-framework Orthography */, wordCount int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("handleTextCheckingResults:forRange:types:options:orthography:wordCount:"), results, range_, checkingTypes, options, orthography, wordCount)
}


// Increases the space between glyphs in the receiver’s selection, or in all text if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/loosenKerning(_:)
func (t_ TextView) LoosenKerning(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("loosenKerning:"), sender)
}


// Lowers the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/lowerBaseline(_:)
func (t_ TextView) LowerBaseline(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("lowerBaseline:"), sender)
}


// Returns an array of URLs for items that can be displayed by QuickLook in the specified ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/quickLookPreviewableItems(inRanges:)
func (t_ TextView) QuickLookPreviewableItemsInRanges(ranges []Value /* primitive/slice/pointer. */) []objc.ID /* already interface */ {
	rv := objc.Send[[]objc.ID](t_.ID, objc.Sel("quickLookPreviewableItemsInRanges:"), ranges)
	return rv
}


// Raises the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/raiseBaseline(_:)
func (t_ TextView) RaiseBaseline(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("raiseBaseline:"), sender)
}


// Returns whether the marker should be removed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:shouldRemove:)
func (t_ TextView) RulerViewShouldRemoveMarker(ruler IRulerView, marker IRulerMarker) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("rulerView:shouldRemoveMarker:"), ruler, marker)
	return rv
}


// Sets the alignment of the paragraphs containing characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setAlignment(_:range:)
func (t_ TextView) SetAlignmentRange(alignment TextAlignment, range_ foundation.objc.IObject /* cross-framework Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlignment:range:"), alignment, range_)
}


// Decreases the space between glyphs in the receiver’s selection, or for all glyphs if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/tightenKerning(_:)
func (t_ TextView) TightenKerning(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("tightenKerning:"), sender)
}


// Toggles the state of the automatic text replacement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticTextReplacement(_:)
func (t_ TextView) ToggleAutomaticTextReplacement(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleAutomaticTextReplacement:"), sender)
}


// Changes the state of smart insert and delete from enabled to disabled and vice versa.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleSmartInsertDelete(_:)
func (t_ TextView) ToggleSmartInsertDelete(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleSmartInsertDelete:"), sender)
}


// Sets the receiver to use nominal glyph spacing for the glyphs in its selection, or for all glyphs if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/turnOffKerning(_:)
func (t_ TextView) TurnOffKerning(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("turnOffKerning:"), sender)
}


// Sets the receiver to use only required ligatures when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/turnOffLigatures(_:)
func (t_ TextView) TurnOffLigatures(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("turnOffLigatures:"), sender)
}


// Sets the receiver to use all ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/useAllLigatures(_:)
func (t_ TextView) UseAllLigatures(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("useAllLigatures:"), sender)
}


// Set the receiver to use pair kerning data for the glyphs in its selection, or for all glyphs if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/useStandardKerning(_:)
func (t_ TextView) UseStandardKerning(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("useStandardKerning:"), sender)
}


// Sets the receiver to use the standard ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/useStandardLigatures(_:)
func (t_ TextView) UseStandardLigatures(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("useStandardLigatures:"), sender)
}


// The layout manager that lays out text for the receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/layoutManager
func (t_ TextView) LayoutManager() objc.IObject /* cross-framework: LayoutManager */ {
	rv := objc.Send[LayoutManager](t_.ID, objc.Sel("layoutManager"))
	return rv
}


// The partial range from the most recent beginning of a word up to the insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rangeForUserCompletion
func (t_ TextView) RangeForUserCompletion() foundation.objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("rangeForUserCompletion"))
	return rv
}


// An array containing the ranges of characters selected in the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectedRanges
func (t_ TextView) SelectedRanges() []Value /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Value](t_.ID, objc.Sel("selectedRanges"))
	return rv
}


// An array containing the ranges of characters selected in the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectedRanges
func (t_ TextView) SetSelectedRanges(value []Value /* primitive/slice/pointer. */) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedRanges:"), nsArray)
}


// The receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContainer
func (t_ TextView) TextContainer() ITextContainer {
	rv := objc.Send[TextContainer](t_.ID, objc.Sel("textContainer"))
	return rv
}


// The receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContainer
func (t_ TextView) SetTextContainer(value ITextContainer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainer:"), value)
}


// The origin of the receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContainerOrigin
func (t_ TextView) TextContainerOrigin() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](t_.ID, objc.Sel("textContainerOrigin"))
	return rv
}


// The receiver’s text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textStorage
func (t_ TextView) TextStorage() ITextStorage {
	rv := objc.Send[TextStorage](t_.ID, objc.Sel("textStorage"))
	return rv
}


// The receiver’s typing attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/typingAttributes
func (t_ TextView) TypingAttributes() foundation.IDictionary /* already interface */ {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("typingAttributes"))
	return rv
}


// The receiver’s typing attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/typingAttributes
func (t_ TextView) SetTypingAttributes(value foundation.IDictionary /* already interface */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTypingAttributes:"), value)
}


// The data types that the receiver accepts as the destination view of a dragging operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/acceptabledragtypes
func (t_ TextView) AcceptableDragTypes() objc.IObject /* cross-framework: PasteboardType */ {
	rv := objc.Send[PasteboardType](t_.ID, objc.Sel("acceptableDragTypes"))
	return rv
}


// The data types that the receiver accepts as the destination view of a dragging operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/acceptabledragtypes
func (t_ TextView) SetAcceptableDragTypes(value objc.IObject /* cross-framework: PasteboardType */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAcceptableDragTypes:"), value)
}


// A Boolean value that indicates whether the receiver accepts the glyph info attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/acceptsglyphinfo
func (t_ TextView) AcceptsGlyphInfo() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("acceptsGlyphInfo"))
	return rv
}


// A Boolean value that indicates whether the receiver accepts the glyph info attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/acceptsglyphinfo
func (t_ TextView) SetAcceptsGlyphInfo(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAcceptsGlyphInfo:"), value)
}


// An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowedinputsourcelocales
func (t_ TextView) AllowedInputSourceLocales() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](t_.ID, objc.Sel("allowedInputSourceLocales"))
	return rv
}


// An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowedinputsourcelocales
func (t_ TextView) SetAllowedInputSourceLocales(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedInputSourceLocales:"), objc.String(value))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowedwritingtoolsresultoptions
func (t_ TextView) AllowedWritingToolsResultOptions() WritingToolsResultOptions {
	rv := objc.Send[WritingToolsResultOptions](t_.ID, objc.Sel("allowedWritingToolsResultOptions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowedwritingtoolsresultoptions
func (t_ TextView) SetAllowedWritingToolsResultOptions(value WritingToolsResultOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedWritingToolsResultOptions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowscharacterpickertouchbaritem
func (t_ TextView) AllowsCharacterPickerTouchBarItem() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsCharacterPickerTouchBarItem"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowscharacterpickertouchbaritem
func (t_ TextView) SetAllowsCharacterPickerTouchBarItem(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsCharacterPickerTouchBarItem:"), value)
}


// A Boolean value that indicates whether the receiver allows its background color to change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsdocumentbackgroundcolorchange
func (t_ TextView) AllowsDocumentBackgroundColorChange() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsDocumentBackgroundColorChange"))
	return rv
}


// A Boolean value that indicates whether the receiver allows its background color to change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsdocumentbackgroundcolorchange
func (t_ TextView) SetAllowsDocumentBackgroundColorChange(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsDocumentBackgroundColorChange:"), value)
}


// Indicates whether image attachments should permit editing of their images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsimageediting
func (t_ TextView) AllowsImageEditing() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsImageEditing"))
	return rv
}


// Indicates whether image attachments should permit editing of their images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsimageediting
func (t_ TextView) SetAllowsImageEditing(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsImageEditing:"), value)
}


// A Boolean value that indicates whether the receiver allows undo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsundo
func (t_ TextView) AllowsUndo() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsUndo"))
	return rv
}


// A Boolean value that indicates whether the receiver allows undo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/allowsundo
func (t_ TextView) SetAllowsUndo(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsUndo:"), value)
}


// The receiver’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/backgroundcolor
func (t_ TextView) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The receiver’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/backgroundcolor
func (t_ TextView) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/candidatelisttouchbaritem
func (t_ TextView) CandidateListTouchBarItem() objc.IObject /* cross-framework: CandidateListTouchBarItem */ {
	rv := objc.Send[CandidateListTouchBarItem](t_.ID, objc.Sel("candidateListTouchBarItem"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/candidatelisttouchbaritem
func (t_ TextView) SetCandidateListTouchBarItem(value objc.IObject /* cross-framework: CandidateListTouchBarItem */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCandidateListTouchBarItem:"), value)
}


// The receiver’s default paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/defaultparagraphstyle
func (t_ TextView) DefaultParagraphStyle() IParagraphStyle {
	rv := objc.Send[ParagraphStyle](t_.ID, objc.Sel("defaultParagraphStyle"))
	return rv
}


// The receiver’s default paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/defaultparagraphstyle
func (t_ TextView) SetDefaultParagraphStyle(value IParagraphStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDefaultParagraphStyle:"), value)
}


// The delegate for all text views sharing the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/delegate
func (t_ TextView) Delegate() TextViewDelegate /* not a class type */ {
	rv := objc.Send[TextViewDelegate](t_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate for all text views sharing the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/delegate
func (t_ TextView) SetDelegate(value TextViewDelegate /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/displayslinktooltips
func (t_ TextView) DisplaysLinkToolTips() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("displaysLinkToolTips"))
	return rv
}


// A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/displayslinktooltips
func (t_ TextView) SetDisplaysLinkToolTips(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDisplaysLinkToolTips:"), value)
}


// A Boolean value that indicates whether the receiver draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/drawsbackground
func (t_ TextView) DrawsBackground() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("drawsBackground"))
	return rv
}


// A Boolean value that indicates whether the receiver draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/drawsbackground
func (t_ TextView) SetDrawsBackground(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsBackground:"), value)
}


// The default text checking types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/enabledtextcheckingtypes
func (t_ TextView) EnabledTextCheckingTypes() TextCheckingTypes /* not a class type */ {
	rv := objc.Send[TextCheckingTypes](t_.ID, objc.Sel("enabledTextCheckingTypes"))
	return rv
}


// The default text checking types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/enabledtextcheckingtypes
func (t_ TextView) SetEnabledTextCheckingTypes(value TextCheckingTypes /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEnabledTextCheckingTypes:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to import files by dragging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/importsgraphics
func (t_ TextView) ImportsGraphics() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("importsGraphics"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to import files by dragging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/importsgraphics
func (t_ TextView) SetImportsGraphics(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImportsGraphics:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/inlinepredictiontype
func (t_ TextView) InlinePredictionType() TextInputTraitType /* not a class type */ {
	rv := objc.Send[TextInputTraitType](t_.ID, objc.Sel("inlinePredictionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/inlinepredictiontype
func (t_ TextView) SetInlinePredictionType(value TextInputTraitType /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInlinePredictionType:"), value)
}


// The color of the insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/insertionpointcolor
func (t_ TextView) InsertionPointColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("insertionPointColor"))
	return rv
}


// The color of the insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/insertionpointcolor
func (t_ TextView) SetInsertionPointColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInsertionPointColor:"), value)
}


// A Boolean value that indicates whether automatic dash substitution is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdashsubstitutionenabled
func (t_ TextView) IsAutomaticDashSubstitutionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticDashSubstitutionEnabled"))
	return rv
}


// A Boolean value that indicates whether automatic dash substitution is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdashsubstitutionenabled
func (t_ TextView) SetIsAutomaticDashSubstitutionEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticDashSubstitutionEnabled:"), value)
}


// A Boolean value that indicates whether automatic data detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdatadetectionenabled
func (t_ TextView) IsAutomaticDataDetectionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticDataDetectionEnabled"))
	return rv
}


// A Boolean value that indicates whether automatic data detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdatadetectionenabled
func (t_ TextView) SetIsAutomaticDataDetectionEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticDataDetectionEnabled:"), value)
}


// A Boolean value that enables or disables automatic link detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticlinkdetectionenabled
func (t_ TextView) IsAutomaticLinkDetectionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticLinkDetectionEnabled"))
	return rv
}


// A Boolean value that enables or disables automatic link detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticlinkdetectionenabled
func (t_ TextView) SetIsAutomaticLinkDetectionEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticLinkDetectionEnabled:"), value)
}


// A Boolean value that enables and disables automatic quotation mark substitution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticquotesubstitutionenabled
func (t_ TextView) IsAutomaticQuoteSubstitutionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticQuoteSubstitutionEnabled"))
	return rv
}


// A Boolean value that enables and disables automatic quotation mark substitution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticquotesubstitutionenabled
func (t_ TextView) SetIsAutomaticQuoteSubstitutionEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticQuoteSubstitutionEnabled:"), value)
}


// A Boolean value that indicates whether automatic spelling correction is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticspellingcorrectionenabled
func (t_ TextView) IsAutomaticSpellingCorrectionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticSpellingCorrectionEnabled"))
	return rv
}


// A Boolean value that indicates whether automatic spelling correction is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticspellingcorrectionenabled
func (t_ TextView) SetIsAutomaticSpellingCorrectionEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticSpellingCorrectionEnabled:"), value)
}


// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextcompletionenabled
func (t_ TextView) IsAutomaticTextCompletionEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticTextCompletionEnabled"))
	return rv
}


// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextcompletionenabled
func (t_ TextView) SetIsAutomaticTextCompletionEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticTextCompletionEnabled:"), value)
}


// A Boolean value that indicates whether automatic text replacement is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextreplacementenabled
func (t_ TextView) IsAutomaticTextReplacementEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticTextReplacementEnabled"))
	return rv
}


// A Boolean value that indicates whether automatic text replacement is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextreplacementenabled
func (t_ TextView) SetIsAutomaticTextReplacementEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticTextReplacementEnabled:"), value)
}


// A Boolean value that indicates whether undo coalescing is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscoalescingundo
func (t_ TextView) IsCoalescingUndo() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isCoalescingUndo"))
	return rv
}


// A Boolean value that indicates whether undo coalescing is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscoalescingundo
func (t_ TextView) SetIsCoalescingUndo(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsCoalescingUndo:"), value)
}


// A Boolean value that indicates whether the receiver has continuous spell checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscontinuousspellcheckingenabled
func (t_ TextView) IsContinuousSpellCheckingEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isContinuousSpellCheckingEnabled"))
	return rv
}


// A Boolean value that indicates whether the receiver has continuous spell checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscontinuousspellcheckingenabled
func (t_ TextView) SetIsContinuousSpellCheckingEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsContinuousSpellCheckingEnabled:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iseditable
func (t_ TextView) IsEditable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEditable"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iseditable
func (t_ TextView) SetIsEditable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEditable:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isfieldeditor
func (t_ TextView) IsFieldEditor() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isFieldEditor"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isfieldeditor
func (t_ TextView) SetIsFieldEditor(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsFieldEditor:"), value)
}


// Enables and disables grammar checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isgrammarcheckingenabled
func (t_ TextView) IsGrammarCheckingEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isGrammarCheckingEnabled"))
	return rv
}


// Enables and disables grammar checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isgrammarcheckingenabled
func (t_ TextView) SetIsGrammarCheckingEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsGrammarCheckingEnabled:"), value)
}


// A Boolean value that indicates whether incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isincrementalsearchingenabled
func (t_ TextView) IsIncrementalSearchingEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isIncrementalSearchingEnabled"))
	return rv
}


// A Boolean value that indicates whether incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isincrementalsearchingenabled
func (t_ TextView) SetIsIncrementalSearchingEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsIncrementalSearchingEnabled:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrichtext
func (t_ TextView) IsRichText() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRichText"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrichtext
func (t_ TextView) SetIsRichText(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRichText:"), value)
}


// A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrulervisible
func (t_ TextView) IsRulerVisible() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRulerVisible"))
	return rv
}


// A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrulervisible
func (t_ TextView) SetIsRulerVisible(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRulerVisible:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isselectable
func (t_ TextView) IsSelectable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSelectable"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isselectable
func (t_ TextView) SetIsSelectable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsSelectable:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iswritingtoolsactive
func (t_ TextView) IsWritingToolsActive() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("isWritingToolsActive"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iswritingtoolsactive
func (t_ TextView) SetIsWritingToolsActive(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsWritingToolsActive:"), value)
}


// The attributes used to draw the onscreen presentation of link text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/linktextattributes
func (t_ TextView) LinkTextAttributes() coreml.objc.IObject /* cross-framework: Key */ {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("linkTextAttributes"))
	return rv
}


// The attributes used to draw the onscreen presentation of link text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/linktextattributes
func (t_ TextView) SetLinkTextAttributes(value coreml.objc.IObject /* cross-framework: Key */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLinkTextAttributes:"), value)
}


// The attributes used to draw marked text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/markedtextattributes
func (t_ TextView) MarkedTextAttributes() coreml.objc.IObject /* cross-framework: Key */ {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("markedTextAttributes"))
	return rv
}


// The attributes used to draw marked text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/markedtextattributes
func (t_ TextView) SetMarkedTextAttributes(value coreml.objc.IObject /* cross-framework: Key */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMarkedTextAttributes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/mathexpressioncompletiontype
func (t_ TextView) MathExpressionCompletionType() TextInputTraitType /* not a class type */ {
	rv := objc.Send[TextInputTraitType](t_.ID, objc.Sel("mathExpressionCompletionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/mathexpressioncompletiontype
func (t_ TextView) SetMathExpressionCompletionType(value TextInputTraitType /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMathExpressionCompletionType:"), value)
}


// The range of characters affected by an action method that changes character (not paragraph) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforusercharacterattributechange
func (t_ TextView) RangeForUserCharacterAttributeChange() foundation.objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("rangeForUserCharacterAttributeChange"))
	return rv
}


// The range of characters affected by an action method that changes character (not paragraph) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforusercharacterattributechange
func (t_ TextView) SetRangeForUserCharacterAttributeChange(value foundation.objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangeForUserCharacterAttributeChange:"), value)
}


// The range of characters affected by an action method that changes paragraph (not character) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforuserparagraphattributechange
func (t_ TextView) RangeForUserParagraphAttributeChange() foundation.objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("rangeForUserParagraphAttributeChange"))
	return rv
}


// The range of characters affected by an action method that changes paragraph (not character) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforuserparagraphattributechange
func (t_ TextView) SetRangeForUserParagraphAttributeChange(value foundation.objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangeForUserParagraphAttributeChange:"), value)
}


// The range of characters affected by a method that changes characters (as opposed to attributes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforusertextchange
func (t_ TextView) RangeForUserTextChange() foundation.objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[foundation.Range](t_.ID, objc.Sel("rangeForUserTextChange"))
	return rv
}


// The range of characters affected by a method that changes characters (as opposed to attributes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangeforusertextchange
func (t_ TextView) SetRangeForUserTextChange(value foundation.objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangeForUserTextChange:"), value)
}


// An array containing the ranges of characters affected by an action method that changes character (not paragraph) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforusercharacterattributechange
func (t_ TextView) RangesForUserCharacterAttributeChange() Value /* not a class type */ {
	rv := objc.Send[Value](t_.ID, objc.Sel("rangesForUserCharacterAttributeChange"))
	return rv
}


// An array containing the ranges of characters affected by an action method that changes character (not paragraph) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforusercharacterattributechange
func (t_ TextView) SetRangesForUserCharacterAttributeChange(value Value /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangesForUserCharacterAttributeChange:"), value)
}


// An array containing the ranges of characters affected by a method that changes paragraph (not character) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforuserparagraphattributechange
func (t_ TextView) RangesForUserParagraphAttributeChange() Value /* not a class type */ {
	rv := objc.Send[Value](t_.ID, objc.Sel("rangesForUserParagraphAttributeChange"))
	return rv
}


// An array containing the ranges of characters affected by a method that changes paragraph (not character) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforuserparagraphattributechange
func (t_ TextView) SetRangesForUserParagraphAttributeChange(value Value /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangesForUserParagraphAttributeChange:"), value)
}


// An array containing the ranges of characters affected by a method that changes characters (as opposed to attributes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforusertextchange
func (t_ TextView) RangesForUserTextChange() Value /* not a class type */ {
	rv := objc.Send[Value](t_.ID, objc.Sel("rangesForUserTextChange"))
	return rv
}


// An array containing the ranges of characters affected by a method that changes characters (as opposed to attributes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/rangesforusertextchange
func (t_ TextView) SetRangesForUserTextChange(value Value /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRangesForUserTextChange:"), value)
}


// The types this text view can read immediately from the pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/readablepasteboardtypes
func (t_ TextView) ReadablePasteboardTypes() objc.IObject /* cross-framework: PasteboardType */ {
	rv := objc.Send[PasteboardType](t_.ID, objc.Sel("readablePasteboardTypes"))
	return rv
}


// The types this text view can read immediately from the pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/readablepasteboardtypes
func (t_ TextView) SetReadablePasteboardTypes(value objc.IObject /* cross-framework: PasteboardType */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setReadablePasteboardTypes:"), value)
}


// The attributes used to indicate the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectedtextattributes
func (t_ TextView) SelectedTextAttributes() coreml.objc.IObject /* cross-framework: Key */ {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("selectedTextAttributes"))
	return rv
}


// The attributes used to indicate the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectedtextattributes
func (t_ TextView) SetSelectedTextAttributes(value coreml.objc.IObject /* cross-framework: Key */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedTextAttributes:"), value)
}


// The preferred direction of selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectionaffinity
func (t_ TextView) SelectionAffinity() SelectionAffinity /* not a class type */ {
	rv := objc.Send[SelectionAffinity](t_.ID, objc.Sel("selectionAffinity"))
	return rv
}


// The preferred direction of selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectionaffinity
func (t_ TextView) SetSelectionAffinity(value SelectionAffinity /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionAffinity:"), value)
}


// The selection granularity for subsequent extension of a selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectiongranularity
func (t_ TextView) SelectionGranularity() SelectionGranularity /* not a class type */ {
	rv := objc.Send[SelectionGranularity](t_.ID, objc.Sel("selectionGranularity"))
	return rv
}


// The selection granularity for subsequent extension of a selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/selectiongranularity
func (t_ TextView) SetSelectionGranularity(value SelectionGranularity /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionGranularity:"), value)
}


// A Boolean value that determines whether the receiver should draw its insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/shoulddrawinsertionpoint
func (t_ TextView) ShouldDrawInsertionPoint() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("shouldDrawInsertionPoint"))
	return rv
}


// A Boolean value that determines whether the receiver should draw its insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/shoulddrawinsertionpoint
func (t_ TextView) SetShouldDrawInsertionPoint(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setShouldDrawInsertionPoint:"), value)
}


// A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/smartinsertdeleteenabled
func (t_ TextView) SmartInsertDeleteEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("smartInsertDeleteEnabled"))
	return rv
}


// A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/smartinsertdeleteenabled
func (t_ TextView) SetSmartInsertDeleteEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSmartInsertDeleteEnabled:"), value)
}


// A tag identifying the text view’s text as a document for the spell checker server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/spellcheckerdocumenttag
func (t_ TextView) SpellCheckerDocumentTag() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](t_.ID, objc.Sel("spellCheckerDocumentTag"))
	return rv
}


// A tag identifying the text view’s text as a document for the spell checker server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/spellcheckerdocumenttag
func (t_ TextView) SetSpellCheckerDocumentTag(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSpellCheckerDocumentTag:"), value)
}


// The empty space the receiver leaves around its associated text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontainerinset
func (t_ TextView) TextContainerInset() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](t_.ID, objc.Sel("textContainerInset"))
	return rv
}


// The empty space the receiver leaves around its associated text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontainerinset
func (t_ TextView) SetTextContainerInset(value coregraphics.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainerInset:"), value)
}


// The receiver’s text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontentstorage
func (t_ TextView) TextContentStorage() ITextContentStorage {
	rv := objc.Send[TextContentStorage](t_.ID, objc.Sel("textContentStorage"))
	return rv
}


// The receiver’s text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textcontentstorage
func (t_ TextView) SetTextContentStorage(value ITextContentStorage) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContentStorage:"), value)
}


// ************************* Text Highlight support **************************
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/texthighlightattributes
func (t_ TextView) TextHighlightAttributes() coreml.objc.IObject /* cross-framework: Key */ {
	rv := objc.Send[coreml.Key](t_.ID, objc.Sel("textHighlightAttributes"))
	return rv
}


// ************************* Text Highlight support **************************
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/texthighlightattributes
func (t_ TextView) SetTextHighlightAttributes(value coreml.objc.IObject /* cross-framework: Key */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextHighlightAttributes:"), value)
}


// The manager that lays out text for the receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textlayoutmanager
func (t_ TextView) TextLayoutManager() ITextLayoutManager {
	rv := objc.Send[TextLayoutManager](t_.ID, objc.Sel("textLayoutManager"))
	return rv
}


// The manager that lays out text for the receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/textlayoutmanager
func (t_ TextView) SetTextLayoutManager(value ITextLayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextLayoutManager:"), value)
}


// A Boolean value that indicates whether the framework should use adaptive color mapping for dark appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesadaptivecolormappingfordarkappearance
func (t_ TextView) UsesAdaptiveColorMappingForDarkAppearance() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesAdaptiveColorMappingForDarkAppearance"))
	return rv
}


// A Boolean value that indicates whether the framework should use adaptive color mapping for dark appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesadaptivecolormappingfordarkappearance
func (t_ TextView) SetUsesAdaptiveColorMappingForDarkAppearance(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesAdaptiveColorMappingForDarkAppearance:"), value)
}


// A Boolean value that indicates whether to use the find bar for this text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfindbar
func (t_ TextView) UsesFindBar() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFindBar"))
	return rv
}


// A Boolean value that indicates whether to use the find bar for this text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfindbar
func (t_ TextView) SetUsesFindBar(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFindBar:"), value)
}


// A Boolean value that indicates whether the receiver allows for a find panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfindpanel
func (t_ TextView) UsesFindPanel() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFindPanel"))
	return rv
}


// A Boolean value that indicates whether the receiver allows for a find panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfindpanel
func (t_ TextView) SetUsesFindPanel(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFindPanel:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager use the Font panel and Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfontpanel
func (t_ TextView) UsesFontPanel() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontPanel"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager use the Font panel and Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesfontpanel
func (t_ TextView) SetUsesFontPanel(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontPanel:"), value)
}


// A Boolean value that indicates whether this text view uses the inspector bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesinspectorbar
func (t_ TextView) UsesInspectorBar() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesInspectorBar"))
	return rv
}


// A Boolean value that indicates whether this text view uses the inspector bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesinspectorbar
func (t_ TextView) SetUsesInspectorBar(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesInspectorBar:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesrolloverbuttonforselection
func (t_ TextView) UsesRolloverButtonForSelection() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesRolloverButtonForSelection"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesrolloverbuttonforselection
func (t_ TextView) SetUsesRolloverButtonForSelection(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesRolloverButtonForSelection:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesruler
func (t_ TextView) UsesRuler() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesRuler"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/usesruler
func (t_ TextView) SetUsesRuler(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesRuler:"), value)
}


// The pasteboard types that can be provided from the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/writablepasteboardtypes
func (t_ TextView) WritablePasteboardTypes() objc.IObject /* cross-framework: PasteboardType */ {
	rv := objc.Send[PasteboardType](t_.ID, objc.Sel("writablePasteboardTypes"))
	return rv
}


// The pasteboard types that can be provided from the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/writablepasteboardtypes
func (t_ TextView) SetWritablePasteboardTypes(value objc.IObject /* cross-framework: PasteboardType */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWritablePasteboardTypes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/writingtoolsbehavior
func (t_ TextView) WritingToolsBehavior() WritingToolsBehavior {
	rv := objc.Send[WritingToolsBehavior](t_.ID, objc.Sel("writingToolsBehavior"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/writingtoolsbehavior
func (t_ TextView) SetWritingToolsBehavior(value WritingToolsBehavior) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWritingToolsBehavior:"), value)
}



