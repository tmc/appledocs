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
	AcceptableDragTypes() []string
	AcceptsGlyphInfo() bool
	SetAcceptsGlyphInfo(value bool)
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
	AllowedWritingToolsResultOptions() WritingToolsResultOptions
	SetAllowedWritingToolsResultOptions(value WritingToolsResultOptions)
	AllowsCharacterPickerTouchBarItem() bool
	SetAllowsCharacterPickerTouchBarItem(value bool)
	AllowsDocumentBackgroundColorChange() bool
	SetAllowsDocumentBackgroundColorChange(value bool)
	AllowsImageEditing() bool
	SetAllowsImageEditing(value bool)
	AllowsUndo() bool
	SetAllowsUndo(value bool)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	CandidateListTouchBarItem() ICandidateListTouchBarItem
	DefaultParagraphStyle() IParagraphStyle
	SetDefaultParagraphStyle(value IParagraphStyle)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	DisplaysLinkToolTips() bool
	SetDisplaysLinkToolTips(value bool)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	EnabledTextCheckingTypes() TextCheckingTypes /* not a class type */
	SetEnabledTextCheckingTypes(value TextCheckingTypes /* not a class type */)
	ImportsGraphics() bool
	SetImportsGraphics(value bool)
	InlinePredictionType() TextInputTraitType
	SetInlinePredictionType(value TextInputTraitType)
	InsertionPointColor() IColor
	SetInsertionPointColor(value IColor)
	AutomaticDashSubstitutionEnabled() bool
	SetAutomaticDashSubstitutionEnabled(value bool)
	AutomaticDataDetectionEnabled() bool
	SetAutomaticDataDetectionEnabled(value bool)
	AutomaticLinkDetectionEnabled() bool
	SetAutomaticLinkDetectionEnabled(value bool)
	AutomaticQuoteSubstitutionEnabled() bool
	SetAutomaticQuoteSubstitutionEnabled(value bool)
	AutomaticSpellingCorrectionEnabled() bool
	SetAutomaticSpellingCorrectionEnabled(value bool)
	AutomaticTextCompletionEnabled() bool
	SetAutomaticTextCompletionEnabled(value bool)
	AutomaticTextReplacementEnabled() bool
	SetAutomaticTextReplacementEnabled(value bool)
	CoalescingUndo() bool
	ContinuousSpellCheckingEnabled() bool
	SetContinuousSpellCheckingEnabled(value bool)
	Editable() bool
	SetEditable(value bool)
	FieldEditor() bool
	SetFieldEditor(value bool)
	GrammarCheckingEnabled() bool
	SetGrammarCheckingEnabled(value bool)
	IncrementalSearchingEnabled() bool
	SetIncrementalSearchingEnabled(value bool)
	RichText() bool
	SetRichText(value bool)
	RulerVisible() bool
	SetRulerVisible(value bool)
	Selectable() bool
	SetSelectable(value bool)
	WritingToolsActive() bool
	LayoutManager() ILayoutManager
	LinkTextAttributes() foundation.IDictionary
	SetLinkTextAttributes(value foundation.IDictionary)
	MarkedTextAttributes() foundation.IDictionary
	SetMarkedTextAttributes(value foundation.IDictionary)
	MathExpressionCompletionType() TextInputTraitType
	SetMathExpressionCompletionType(value TextInputTraitType)
	RangeForUserCharacterAttributeChange() corefoundation.Range
	RangeForUserCompletion() corefoundation.Range
	RangeForUserParagraphAttributeChange() corefoundation.Range
	RangeForUserTextChange() corefoundation.Range
	RangesForUserCharacterAttributeChange() []foundation.Value
	RangesForUserParagraphAttributeChange() []foundation.Value
	RangesForUserTextChange() []foundation.Value
	ReadablePasteboardTypes() []string
	SelectedRanges() []foundation.Value
	SetSelectedRanges(value []foundation.Value)
	SelectedTextAttributes() foundation.IDictionary
	SetSelectedTextAttributes(value foundation.IDictionary)
	SelectionAffinity() SelectionAffinity
	SelectionGranularity() SelectionGranularity
	SetSelectionGranularity(value SelectionGranularity)
	ShouldDrawInsertionPoint() bool
	SmartInsertDeleteEnabled() bool
	SetSmartInsertDeleteEnabled(value bool)
	SpellCheckerDocumentTag() int
	TextContainer() ITextContainer
	SetTextContainer(value ITextContainer)
	TextContainerInset() objc.IObject /* cross-framework: Size */
	SetTextContainerInset(value objc.IObject /* cross-framework: Size */)
	TextContainerOrigin() objc.IObject /* cross-framework: Point */
	TextContentStorage() ITextContentStorage
	TextHighlightAttributes() foundation.IDictionary
	SetTextHighlightAttributes(value foundation.IDictionary)
	TextLayoutManager() ITextLayoutManager
	TextStorage() ITextStorage
	TypingAttributes() foundation.IDictionary
	SetTypingAttributes(value foundation.IDictionary)
	UsesAdaptiveColorMappingForDarkAppearance() bool
	SetUsesAdaptiveColorMappingForDarkAppearance(value bool)
	UsesFindBar() bool
	SetUsesFindBar(value bool)
	UsesFindPanel() bool
	SetUsesFindPanel(value bool)
	UsesFontPanel() bool
	SetUsesFontPanel(value bool)
	UsesInspectorBar() bool
	SetUsesInspectorBar(value bool)
	UsesRolloverButtonForSelection() bool
	SetUsesRolloverButtonForSelection(value bool)
	UsesRuler() bool
	SetUsesRuler(value bool)
	WritablePasteboardTypes() []string
	WritingToolsBehavior() WritingToolsBehavior
	SetWritingToolsBehavior(value WritingToolsBehavior)
	IsAutomaticDashSubstitutionEnabled() bool
	SetIsAutomaticDashSubstitutionEnabled(value bool)
	IsAutomaticDataDetectionEnabled() bool
	SetIsAutomaticDataDetectionEnabled(value bool)
	IsAutomaticLinkDetectionEnabled() bool
	SetIsAutomaticLinkDetectionEnabled(value bool)
	IsAutomaticQuoteSubstitutionEnabled() bool
	SetIsAutomaticQuoteSubstitutionEnabled(value bool)
	IsAutomaticSpellingCorrectionEnabled() bool
	SetIsAutomaticSpellingCorrectionEnabled(value bool)
	IsAutomaticTextCompletionEnabled() bool
	SetIsAutomaticTextCompletionEnabled(value bool)
	IsAutomaticTextReplacementEnabled() bool
	SetIsAutomaticTextReplacementEnabled(value bool)
	IsCoalescingUndo() bool
	SetIsCoalescingUndo(value bool)
	IsContinuousSpellCheckingEnabled() bool
	SetIsContinuousSpellCheckingEnabled(value bool)
	IsEditable() bool
	SetIsEditable(value bool)
	IsFieldEditor() bool
	SetIsFieldEditor(value bool)
	IsGrammarCheckingEnabled() bool
	SetIsGrammarCheckingEnabled(value bool)
	IsIncrementalSearchingEnabled() bool
	SetIsIncrementalSearchingEnabled(value bool)
	IsRichText() bool
	SetIsRichText(value bool)
	IsRulerVisible() bool
	SetIsRulerVisible(value bool)
	IsSelectable() bool
	SetIsSelectable(value bool)
	IsWritingToolsActive() bool
	SetIsWritingToolsActive(value bool)
	// methods:
	AlignJustified(sender objc.IObject)
	BreakUndoCoalescing()
	ChangeAttributes(sender objc.IObject)
	ChangeColor(sender objc.IObject)
	ChangeDocumentBackgroundColor(sender objc.IObject)
	ChangeLayoutOrientation(sender objc.IObject)
	CharacterIndexForInsertionAtPoint(point objc.IObject /* cross-framework: Point */) uint
	CheckTextInRangeTypesOptions(range_ corefoundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary)
	CheckTextInDocument(sender objc.IObject)
	CheckTextInSelection(sender objc.IObject)
	CleanUpAfterDragOperation()
	ClickedOnLinkAtIndex(link objc.IObject, charIndex uint)
	Complete(sender objc.IObject)
	CompletionsForPartialWordRangeIndexOfSelectedItem(charRange corefoundation.Range, index int) []string
	DidChangeText()
	DragImageForSelectionWithEventOrigin(event IEvent, origin PointPointer /* not a class type */) IImage
	DragOperationForDraggingInfoType(dragInfo objc.IObject, type_ objc.IObject /* cross-framework: PasteboardType */) DragOperation
	DragSelectionWithEventOffsetSlideBack(event IEvent, mouseOffset objc.IObject /* cross-framework: Size */, slideBack bool) bool
	DrawViewBackgroundInRect(rect objc.IObject /* cross-framework: Rect */)
	DrawInsertionPointInRectColorTurnedOn(rect objc.IObject /* cross-framework: Rect */, color IColor, flag bool)
	DrawTextHighlightBackgroundForTextRangeOrigin(textRange ITextRange, origin objc.IObject /* cross-framework: Point */)
	HandleTextCheckingResultsForRangeTypesOptionsOrthographyWordCount(results []foundation.TextCheckingResult, range_ corefoundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, orthography foundation.Orthography, wordCount int)
	Highlight(sender objc.IObject)
	InsertCompletionForPartialWordRangeMovementIsFinal(word objc.IObject /* cross-framework: NSString */, charRange corefoundation.Range, movement int, flag bool)
	InvalidateTextContainerOrigin()
	LoosenKerning(sender objc.IObject)
	LowerBaseline(sender objc.IObject)
	OrderFrontLinkPanel(sender objc.IObject)
	OrderFrontListPanel(sender objc.IObject)
	OrderFrontSharingServicePicker(sender objc.IObject)
	OrderFrontSpacingPanel(sender objc.IObject)
	OrderFrontSubstitutionsPanel(sender objc.IObject)
	OrderFrontTablePanel(sender objc.IObject)
	Outline(sender objc.IObject)
	PasteAsPlainText(sender objc.IObject)
	PasteAsRichText(sender objc.IObject)
	PerformFindPanelAction(sender objc.IObject)
	PerformValidatedReplacementInRangeWithAttributedString(range_ corefoundation.Range, attributedString foundation.AttributedString) bool
	PreferredPasteboardTypeFromArrayRestrictedToTypesFromArray(availableTypes []string, allowedTypes []string) objc.IObject /* cross-framework: PasteboardType */
	QuickLookPreviewableItemsInRanges(ranges []foundation.Value) []objc.ID
	RaiseBaseline(sender objc.IObject)
	ReadSelectionFromPasteboard(pboard IPasteboard) bool
	ReadSelectionFromPasteboardType(pboard IPasteboard, type_ objc.IObject /* cross-framework: PasteboardType */) bool
	ReplaceTextContainer(newContainer ITextContainer)
	RulerViewDidAddMarker(ruler IRulerView, marker IRulerMarker)
	RulerViewDidMoveMarker(ruler IRulerView, marker IRulerMarker)
	RulerViewDidRemoveMarker(ruler IRulerView, marker IRulerMarker)
	RulerViewHandleMouseDown(ruler IRulerView, event IEvent)
	RulerViewShouldAddMarker(ruler IRulerView, marker IRulerMarker) bool
	RulerViewShouldMoveMarker(ruler IRulerView, marker IRulerMarker) bool
	RulerViewShouldRemoveMarker(ruler IRulerView, marker IRulerMarker) bool
	RulerViewWillAddMarkerAtLocation(ruler IRulerView, marker IRulerMarker, location float64) float64
	RulerViewWillMoveMarkerToLocation(ruler IRulerView, marker IRulerMarker, location float64) float64
	SelectionRangeForProposedRangeGranularity(proposedCharRange corefoundation.Range, granularity SelectionGranularity) corefoundation.Range
	SetAlignmentRange(alignment TextAlignment, range_ corefoundation.Range)
	SetBaseWritingDirectionRange(writingDirection WritingDirection, range_ corefoundation.Range)
	SetConstrainedFrameSize(desiredSize objc.IObject /* cross-framework: Size */)
	SetLayoutOrientation(orientation TextLayoutOrientation)
	SetNeedsDisplayInRectAvoidAdditionalLayout(rect objc.IObject /* cross-framework: Rect */, flag bool)
	SetSelectedRange(charRange corefoundation.Range)
	SetSelectedRangeAffinityStillSelecting(charRange corefoundation.Range, affinity SelectionAffinity, stillSelectingFlag bool)
	SetSelectedRangesAffinityStillSelecting(ranges []foundation.Value, affinity SelectionAffinity, stillSelectingFlag bool)
	SetSpellingStateRange(value int, charRange corefoundation.Range)
	ShouldChangeTextInRangeReplacementString(affectedCharRange corefoundation.Range, replacementString objc.IObject /* cross-framework: NSString */) bool
	ShouldChangeTextInRangesReplacementStrings(affectedRanges []foundation.Value, replacementStrings []string) bool
	ShowFindIndicatorForRange(charRange corefoundation.Range)
	SmartDeleteRangeForProposedRange(proposedCharRange corefoundation.Range) corefoundation.Range
	SmartInsertAfterStringForStringReplacingRange(pasteString objc.IObject /* cross-framework: NSString */, charRangeToReplace corefoundation.Range) foundation.String
	SmartInsertBeforeStringForStringReplacingRange(pasteString objc.IObject /* cross-framework: NSString */, charRangeToReplace corefoundation.Range) foundation.String
	SmartInsertForStringReplacingRangeBeforeStringAfterString(pasteString objc.IObject /* cross-framework: NSString */, charRangeToReplace corefoundation.Range, beforeString objc.IObject /* cross-framework: NSString */, afterString objc.IObject /* cross-framework: NSString */)
	StartSpeaking(sender objc.IObject)
	StopSpeaking(sender objc.IObject)
	TightenKerning(sender objc.IObject)
	ToggleAutomaticDashSubstitution(sender objc.IObject)
	ToggleAutomaticDataDetection(sender objc.IObject)
	ToggleAutomaticLinkDetection(sender objc.IObject)
	ToggleAutomaticQuoteSubstitution(sender objc.IObject)
	ToggleAutomaticSpellingCorrection(sender objc.IObject)
	ToggleAutomaticTextCompletion(sender objc.IObject)
	ToggleAutomaticTextReplacement(sender objc.IObject)
	ToggleContinuousSpellChecking(sender objc.IObject)
	ToggleGrammarChecking(sender objc.IObject)
	ToggleQuickLookPreviewPanel(sender objc.IObject)
	ToggleSmartInsertDelete(sender objc.IObject)
	TurnOffKerning(sender objc.IObject)
	TurnOffLigatures(sender objc.IObject)
	UpdateCandidates()
	UpdateDragTypeRegistration()
	UpdateFontPanel()
	UpdateInsertionPointStateAndRestartTimer(restartFlag bool)
	UpdateQuickLookPreviewPanel()
	UpdateRuler()
	UpdateTextTouchBarItems()
	UpdateTouchBarItemIdentifiers()
	UseAllLigatures(sender objc.IObject)
	UseStandardKerning(sender objc.IObject)
	UseStandardLigatures(sender objc.IObject)
	ValidRequestorForSendTypeReturnType(sendType objc.IObject /* cross-framework: PasteboardType */, returnType objc.IObject /* cross-framework: PasteboardType */) objc.ID
	WriteSelectionToPasteboardType(pboard IPasteboard, type_ objc.IObject /* cross-framework: PasteboardType */) bool
	WriteSelectionToPasteboardTypes(pboard IPasteboard, types []string) bool
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/init(usingTextLayoutManager:)
func NewTextViewUsingTextLayoutManager(usingTextLayoutManager bool) TextView {
	instance := getTextViewClass().Alloc()
	rv := objc.Send[TextView](instance.ID, objc.Sel("initUsingTextLayoutManager:"), usingTextLayoutManager)
	rv.Autorelease()
	return rv
}


// Initializes a text view with data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/init(coder:)
func NewTextViewWithCoder(coder foundation.Coder) TextView {
	instance := getTextViewClass().Alloc()
	rv := objc.Send[TextView](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Initializes a text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/init(frame:)
func NewTextViewWithFrame(frameRect objc.IObject /* cross-framework: Rect */) TextView {
	instance := getTextViewClass().Alloc()
	rv := objc.Send[TextView](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}


// Initializes a text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/init(frame:textContainer:)
func NewTextViewWithFrameTextContainer(frameRect objc.IObject /* cross-framework: Rect */, container ITextContainer) TextView {
	instance := getTextViewClass().Alloc()
	rv := objc.Send[TextView](instance.ID, objc.Sel("initWithFrame:textContainer:"), frameRect, container)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/fieldEditor()
func (tc _TextViewClass) FieldEditor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("fieldEditor"))
	return rv
}


// Registers send and return types for the Services facility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/registerForServices()
func (tc _TextViewClass) RegisterForServices() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("registerForServices"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/scrollableDocumentContentTextView()
func (tc _TextViewClass) ScrollableDocumentContentTextView() IScrollView {
	rv := objc.Send[ScrollView](objc.ID(tc.class), objc.Sel("scrollableDocumentContentTextView"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/scrollablePlainDocumentContentTextView()
func (tc _TextViewClass) ScrollablePlainDocumentContentTextView() IScrollView {
	rv := objc.Send[ScrollView](objc.ID(tc.class), objc.Sel("scrollablePlainDocumentContentTextView"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/scrollableTextView()
func (tc _TextViewClass) ScrollableTextView() IScrollView {
	rv := objc.Send[ScrollView](objc.ID(tc.class), objc.Sel("scrollableTextView"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textViewUsingTextLayoutManager:
func (tc _TextViewClass) TextViewUsingTextLayoutManager(usingTextLayoutManager bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("textViewUsingTextLayoutManager:"), usingTextLayoutManager)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/stronglyReferencesTextStorage
func (tc _TextViewClass) StronglyReferencesTextStorage() bool {
	rv := objc.Send[bool](objc.ID(tc.class), objc.Sel("stronglyReferencesTextStorage"))
	return rv
}

// Applies full justification to selected paragraphs (or all text, if the receiver is a plain text object).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/alignJustified(_:)
func (t_ TextView) AlignJustified(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("alignJustified:"), sender)
}


// Informs the receiver that it should begin coalescing successive typing operations in a new undo grouping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/breakUndoCoalescing()
func (t_ TextView) BreakUndoCoalescing() {
	objc.Send[objc.ID](t_.ID, objc.Sel("breakUndoCoalescing"))
}


// Changes the attributes of the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/changeAttributes(_:)
func (t_ TextView) ChangeAttributes(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("changeAttributes:"), sender)
}


// Sets the color of the selected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/changeColor(_:)
func (t_ TextView) ChangeColor(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("changeColor:"), sender)
}


// An action method used to set the background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/changeDocumentBackgroundColor(_:)
func (t_ TextView) ChangeDocumentBackgroundColor(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("changeDocumentBackgroundColor:"), sender)
}


// An action method that sets the layout orientation of the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/changeLayoutOrientation(_:)
func (t_ TextView) ChangeLayoutOrientation(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("changeLayoutOrientation:"), sender)
}


// Returns a character index appropriate for placing a zero-length selection for an insertion point associated with the mouse at the given point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/characterIndexForInsertion(at:)
func (t_ TextView) CharacterIndexForInsertionAtPoint(point objc.IObject /* cross-framework: Point */) uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("characterIndexForInsertionAtPoint:"), point)
	return rv
}


// Check and replace the text in the range using the specified checking types and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/checkText(in:types:options:)
func (t_ TextView) CheckTextInRangeTypesOptions(range_ corefoundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("checkTextInRange:types:options:"), range_, checkingTypes, options)
}


// Performs the default text checking on the entire document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/checkTextInDocument(_:)
func (t_ TextView) CheckTextInDocument(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("checkTextInDocument:"), sender)
}


// Performs the default text checking on the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/checkTextInSelection(_:)
func (t_ TextView) CheckTextInSelection(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("checkTextInSelection:"), sender)
}


// Releases the drag information still existing after the dragging session has completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/cleanUpAfterDragOperation()
func (t_ TextView) CleanUpAfterDragOperation() {
	objc.Send[objc.ID](t_.ID, objc.Sel("cleanUpAfterDragOperation"))
}


// Causes the text view to act as if the user clicked on some text with the given link as the value of a link attribute associated with the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/clicked(onLink:at:)
func (t_ TextView) ClickedOnLinkAtIndex(link objc.IObject, charIndex uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("clickedOnLink:atIndex:"), link, charIndex)
}


// Invokes completion in a text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/complete(_:)
func (t_ TextView) Complete(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("complete:"), sender)
}


// Returns an array of potential completions, in the order to be presented, representing possible word completions available from a partial word.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/completions(forPartialWordRange:indexOfSelectedItem:)
func (t_ TextView) CompletionsForPartialWordRangeIndexOfSelectedItem(charRange corefoundation.Range, index int) []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("completionsForPartialWordRange:indexOfSelectedItem:"), charRange, index)
	return rv
}


// Sends out necessary notifications when a text change completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/didChangeText()
func (t_ TextView) DidChangeText() {
	objc.Send[objc.ID](t_.ID, objc.Sel("didChangeText"))
}


// Returns an appropriate drag image for the drag initiated by the specified event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/dragImageForSelection(with:origin:)
func (t_ TextView) DragImageForSelectionWithEventOrigin(event IEvent, origin PointPointer /* not a class type */) IImage {
	rv := objc.Send[Image](t_.ID, objc.Sel("dragImageForSelectionWithEvent:origin:"), event, origin)
	return rv
}


// Returns the type of drag operation that should be performed if the image were released now.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/dragOperation(for:type:)
func (t_ TextView) DragOperationForDraggingInfoType(dragInfo objc.IObject, type_ objc.IObject /* cross-framework: PasteboardType */) DragOperation {
	rv := objc.Send[DragOperation](t_.ID, objc.Sel("dragOperationForDraggingInfo:type:"), dragInfo, type_)
	return rv
}


// Begins dragging the current selected text range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/dragSelection(with:offset:slideBack:)
func (t_ TextView) DragSelectionWithEventOffsetSlideBack(event IEvent, mouseOffset objc.IObject /* cross-framework: Size */, slideBack bool) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("dragSelectionWithEvent:offset:slideBack:"), event, mouseOffset, slideBack)
	return rv
}


// Draws the background of the text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/drawBackground(in:)
func (t_ TextView) DrawViewBackgroundInRect(rect objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawViewBackgroundInRect:"), rect)
}


// Draws or erases the insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/drawInsertionPoint(in:color:turnedOn:)
func (t_ TextView) DrawInsertionPointInRectColorTurnedOn(rect objc.IObject /* cross-framework: Rect */, color IColor, flag bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawInsertionPointInRect:color:turnedOn:"), rect, color, flag)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/drawTextHighlightBackground(for:origin:)
func (t_ TextView) DrawTextHighlightBackgroundForTextRangeOrigin(textRange ITextRange, origin objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawTextHighlightBackgroundForTextRange:origin:"), textRange, origin)
}


// Handles the text checking results returned by the text view
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/handleTextCheckingResults(_:forRange:types:options:orthography:wordCount:)
func (t_ TextView) HandleTextCheckingResultsForRangeTypesOptionsOrthographyWordCount(results []foundation.TextCheckingResult, range_ corefoundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, orthography foundation.Orthography, wordCount int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("handleTextCheckingResults:forRange:types:options:orthography:wordCount:"), results, range_, checkingTypes, options, orthography, wordCount)
}


// An action for toggling in the receiver’s selected range. The sender should be a menu item with a of type ( ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/highlight(_:)
func (t_ TextView) Highlight(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("highlight:"), sender)
}


// Inserts the selected completion into the text at the appropriate location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/insertCompletion(_:forPartialWordRange:movement:isFinal:)
func (t_ TextView) InsertCompletionForPartialWordRangeMovementIsFinal(word objc.IObject /* cross-framework: NSString */, charRange corefoundation.Range, movement int, flag bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertCompletion:forPartialWordRange:movement:isFinal:"), word, charRange, movement, flag)
}


// Invalidates the calculated origin of the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/invalidateTextContainerOrigin()
func (t_ TextView) InvalidateTextContainerOrigin() {
	objc.Send[objc.ID](t_.ID, objc.Sel("invalidateTextContainerOrigin"))
}


// Increases the space between glyphs in the receiver’s selection, or in all text if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/loosenKerning(_:)
func (t_ TextView) LoosenKerning(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("loosenKerning:"), sender)
}


// Lowers the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/lowerBaseline(_:)
func (t_ TextView) LowerBaseline(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("lowerBaseline:"), sender)
}


// Brings forward a panel allowing the user to manipulate links in the text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontLinkPanel(_:)
func (t_ TextView) OrderFrontLinkPanel(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("orderFrontLinkPanel:"), sender)
}


// Brings forward a panel allowing the user to manipulate text lists in the text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontListPanel(_:)
func (t_ TextView) OrderFrontListPanel(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("orderFrontListPanel:"), sender)
}


// Creates and displays a new instance of the sharing service picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontSharingServicePicker(_:)
func (t_ TextView) OrderFrontSharingServicePicker(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("orderFrontSharingServicePicker:"), sender)
}


// Brings forward a panel allowing the user to manipulate text line heights, interline spacing, and paragraph spacing, in the text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontSpacingPanel(_:)
func (t_ TextView) OrderFrontSpacingPanel(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("orderFrontSpacingPanel:"), sender)
}


// Brings forward a panel allowing the user to specify string substitutions in the text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontSubstitutionsPanel(_:)
func (t_ TextView) OrderFrontSubstitutionsPanel(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("orderFrontSubstitutionsPanel:"), sender)
}


// Brings forward a panel allowing the user to manipulate text tables in the text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontTablePanel(_:)
func (t_ TextView) OrderFrontTablePanel(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("orderFrontTablePanel:"), sender)
}


// Adds the outline attribute to the selected text attributes if absent; removes the attribute if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/outline(_:)
func (t_ TextView) Outline(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("outline:"), sender)
}


// Inserts the contents of the pasteboard into the receiver’s text as plain text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/pasteAsPlainText(_:)
func (t_ TextView) PasteAsPlainText(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("pasteAsPlainText:"), sender)
}


// This action method inserts the contents of the pasteboard into the receiver’s text as rich text, maintaining its attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/pasteAsRichText(_:)
func (t_ TextView) PasteAsRichText(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("pasteAsRichText:"), sender)
}


// Performs a find panel action specified by the sender’s tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/performFindPanelAction(_:)
func (t_ TextView) PerformFindPanelAction(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("performFindPanelAction:"), sender)
}


// Replaces text in the range you specify with the attributed string you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/performValidatedReplacement(in:with:)
func (t_ TextView) PerformValidatedReplacementInRangeWithAttributedString(range_ corefoundation.Range, attributedString foundation.AttributedString) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("performValidatedReplacementInRange:withAttributedString:"), range_, attributedString)
	return rv
}


// Returns whatever type on the pasteboard would be most preferred for copying data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/preferredPasteboardType(from:restrictedToTypesFrom:)
func (t_ TextView) PreferredPasteboardTypeFromArrayRestrictedToTypesFromArray(availableTypes []string, allowedTypes []string) objc.IObject /* cross-framework: PasteboardType */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("preferredPasteboardTypeFromArray:restrictedToTypesFromArray:"), availableTypes, allowedTypes)
	return rv
}


// Returns an array of URLs for items that can be displayed by QuickLook in the specified ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/quickLookPreviewableItems(inRanges:)
func (t_ TextView) QuickLookPreviewableItemsInRanges(ranges []foundation.Value) []objc.ID {
	rv := objc.Send[[]objc.ID](t_.ID, objc.Sel("quickLookPreviewableItemsInRanges:"), ranges)
	return rv
}


// Raises the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/raiseBaseline(_:)
func (t_ TextView) RaiseBaseline(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("raiseBaseline:"), sender)
}


// Reads the text view’s preferred type of data from the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/readSelection(from:)
func (t_ TextView) ReadSelectionFromPasteboard(pboard IPasteboard) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("readSelectionFromPasteboard:"), pboard)
	return rv
}


// Reads data of the given type from the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/readSelection(from:type:)
func (t_ TextView) ReadSelectionFromPasteboardType(pboard IPasteboard, type_ objc.IObject /* cross-framework: PasteboardType */) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("readSelectionFromPasteboard:type:"), pboard, type_)
	return rv
}


// Replaces the text container for the group of text system objects containing the receiver, keeping the association between the receiver and its layout manager intact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/replaceTextContainer(_:)
func (t_ TextView) ReplaceTextContainer(newContainer ITextContainer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceTextContainer:"), newContainer)
}


// Modifies the paragraph style of the paragraphs containing the selection to accommodate a new marker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:didAdd:)
func (t_ TextView) RulerViewDidAddMarker(ruler IRulerView, marker IRulerMarker) {
	objc.Send[objc.ID](t_.ID, objc.Sel("rulerView:didAddMarker:"), ruler, marker)
}


// Modifies the paragraph style of the paragraphs containing the selection to record the new location of the marker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:didMove:)
func (t_ TextView) RulerViewDidMoveMarker(ruler IRulerView, marker IRulerMarker) {
	objc.Send[objc.ID](t_.ID, objc.Sel("rulerView:didMoveMarker:"), ruler, marker)
}


// Modifies the paragraph style of the paragraphs containing the selection—if possible—by removing the specified marker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:didRemove:)
func (t_ TextView) RulerViewDidRemoveMarker(ruler IRulerView, marker IRulerMarker) {
	objc.Send[objc.ID](t_.ID, objc.Sel("rulerView:didRemoveMarker:"), ruler, marker)
}


// Adds a left tab marker to the ruler at the location clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:handleMouseDownWith:)
func (t_ TextView) RulerViewHandleMouseDown(ruler IRulerView, event IEvent) {
	objc.Send[objc.ID](t_.ID, objc.Sel("rulerView:handleMouseDown:"), ruler, event)
}


// Returns whether a new marker can be added.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:shouldAdd:)
func (t_ TextView) RulerViewShouldAddMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rulerView:shouldAddMarker:"), ruler, marker)
	return rv
}


// Returns whether the marker should be moved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:shouldMove:)
func (t_ TextView) RulerViewShouldMoveMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rulerView:shouldMoveMarker:"), ruler, marker)
	return rv
}


// Returns whether the marker should be removed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:shouldRemove:)
func (t_ TextView) RulerViewShouldRemoveMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rulerView:shouldRemoveMarker:"), ruler, marker)
	return rv
}


// Returns a potentially modified location to which the marker should be added.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:willAdd:atLocation:)
func (t_ TextView) RulerViewWillAddMarkerAtLocation(ruler IRulerView, marker IRulerMarker, location float64) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("rulerView:willAddMarker:atLocation:"), ruler, marker, location)
	return rv
}


// Returns a potentially modified location to which the marker should be moved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:willMove:toLocation:)
func (t_ TextView) RulerViewWillMoveMarkerToLocation(ruler IRulerView, marker IRulerMarker, location float64) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("rulerView:willMoveMarker:toLocation:"), ruler, marker, location)
	return rv
}


// Returns an adjusted selected range based on the selection granularity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectionRange(forProposedRange:granularity:)
func (t_ TextView) SelectionRangeForProposedRangeGranularity(proposedCharRange corefoundation.Range, granularity SelectionGranularity) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("selectionRangeForProposedRange:granularity:"), proposedCharRange, granularity)
	return rv
}


// Sets the alignment of the paragraphs containing characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setAlignment(_:range:)
func (t_ TextView) SetAlignmentRange(alignment TextAlignment, range_ corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlignment:range:"), alignment, range_)
}


// Sets the base writing direction of a range of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setBaseWritingDirection(_:range:)
func (t_ TextView) SetBaseWritingDirectionRange(writingDirection WritingDirection, range_ corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBaseWritingDirection:range:"), writingDirection, range_)
}


// Attempts to set the frame size as if by user action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setConstrainedFrameSize(_:)
func (t_ TextView) SetConstrainedFrameSize(desiredSize objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setConstrainedFrameSize:"), desiredSize)
}


// Changes the receiver’s layout orientation and invalidates the contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setLayoutOrientation(_:)
func (t_ TextView) SetLayoutOrientation(orientation TextLayoutOrientation) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutOrientation:"), orientation)
}


// Marks the receiver as requiring display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setNeedsDisplay(_:avoidAdditionalLayout:)
func (t_ TextView) SetNeedsDisplayInRectAvoidAdditionalLayout(rect objc.IObject /* cross-framework: Rect */, flag bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNeedsDisplayInRect:avoidAdditionalLayout:"), rect, flag)
}


// Selects the specified range of characters in response to user action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setSelectedRange(_:)
func (t_ TextView) SetSelectedRange(charRange corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedRange:"), charRange)
}


// Sets the selection to a range of characters in response to user action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setSelectedRange(_:affinity:stillSelecting:)
func (t_ TextView) SetSelectedRangeAffinityStillSelecting(charRange corefoundation.Range, affinity SelectionAffinity, stillSelectingFlag bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedRange:affinity:stillSelecting:"), charRange, affinity, stillSelectingFlag)
}


// Sets the selection to the characters in an array of ranges in response to user action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setSelectedRanges(_:affinity:stillSelecting:)
func (t_ TextView) SetSelectedRangesAffinityStillSelecting(ranges []foundation.Value, affinity SelectionAffinity, stillSelectingFlag bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedRanges:affinity:stillSelecting:"), ranges, affinity, stillSelectingFlag)
}


// Sets the spelling state, which controls the display of the spelling and grammar indicators on the given text range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setSpellingState(_:range:)
func (t_ TextView) SetSpellingStateRange(value int, charRange corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSpellingState:range:"), value, charRange)
}


// Initiates a series of delegate messages (and general notifications) to determine whether modifications can be made to the characters and attributes of the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/shouldChangeText(in:replacementString:)
func (t_ TextView) ShouldChangeTextInRangeReplacementString(affectedCharRange corefoundation.Range, replacementString objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("shouldChangeTextInRange:replacementString:"), affectedCharRange, replacementString)
	return rv
}


// Initiates a series of delegate messages (and general notifications) to determine whether modifications can be made to the characters and attributes of the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/shouldChangeText(inRanges:replacementStrings:)
func (t_ TextView) ShouldChangeTextInRangesReplacementStrings(affectedRanges []foundation.Value, replacementStrings []string) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("shouldChangeTextInRanges:replacementStrings:"), affectedRanges, replacementStrings)
	return rv
}


// Causes a temporary highlighting effect to appear around the visible portion (or portions) of the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/showFindIndicator(for:)
func (t_ TextView) ShowFindIndicatorForRange(charRange corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("showFindIndicatorForRange:"), charRange)
}


// Returns an extended range that includes adjacent whitespace that should be deleted along with the proposed range in order to preserve proper spacing and punctuation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/smartDeleteRange(forProposedRange:)
func (t_ TextView) SmartDeleteRangeForProposedRange(proposedCharRange corefoundation.Range) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("smartDeleteRangeForProposedRange:"), proposedCharRange)
	return rv
}


// Returns any whitespace that needs to be added after the string to preserve proper spacing and punctuation when the string replaces the characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/smartInsert(afterStringFor:replacing:)
func (t_ TextView) SmartInsertAfterStringForStringReplacingRange(pasteString objc.IObject /* cross-framework: NSString */, charRangeToReplace corefoundation.Range) foundation.String {
	rv := objc.Send[foundation.String](t_.ID, objc.Sel("smartInsertAfterStringForString:replacingRange:"), pasteString, charRangeToReplace)
	return rv
}


// Returns any whitespace that needs to be added before the string to preserve proper spacing and punctuation when the string replaces the characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/smartInsert(beforeStringFor:replacing:)
func (t_ TextView) SmartInsertBeforeStringForStringReplacingRange(pasteString objc.IObject /* cross-framework: NSString */, charRangeToReplace corefoundation.Range) foundation.String {
	rv := objc.Send[foundation.String](t_.ID, objc.Sel("smartInsertBeforeStringForString:replacingRange:"), pasteString, charRangeToReplace)
	return rv
}


// Determines whether whitespace needs to be added around the string to preserve proper spacing and punctuation when it replaces the characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/smartInsert(for:replacing:before:after:)
func (t_ TextView) SmartInsertForStringReplacingRangeBeforeStringAfterString(pasteString objc.IObject /* cross-framework: NSString */, charRangeToReplace corefoundation.Range, beforeString objc.IObject /* cross-framework: NSString */, afterString objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("smartInsertForString:replacingRange:beforeString:afterString:"), pasteString, charRangeToReplace, beforeString, afterString)
}


// Speaks the selected text, or all text if no selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/startSpeaking(_:)
func (t_ TextView) StartSpeaking(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("startSpeaking:"), sender)
}


// Stops the speaking of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/stopSpeaking(_:)
func (t_ TextView) StopSpeaking(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("stopSpeaking:"), sender)
}


// Decreases the space between glyphs in the receiver’s selection, or for all glyphs if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/tightenKerning(_:)
func (t_ TextView) TightenKerning(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("tightenKerning:"), sender)
}


// Toggles the state of the automatic dash substitution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticDashSubstitution(_:)
func (t_ TextView) ToggleAutomaticDashSubstitution(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleAutomaticDashSubstitution:"), sender)
}


// Toggles the state of the automatic data detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticDataDetection(_:)
func (t_ TextView) ToggleAutomaticDataDetection(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleAutomaticDataDetection:"), sender)
}


// Changes the state of automatic link detection from enabled to disabled and vice versa.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticLinkDetection(_:)
func (t_ TextView) ToggleAutomaticLinkDetection(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleAutomaticLinkDetection:"), sender)
}


// Changes the state of automatic quotation mark substitution from enabled to disabled and vice versa.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticQuoteSubstitution(_:)
func (t_ TextView) ToggleAutomaticQuoteSubstitution(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleAutomaticQuoteSubstitution:"), sender)
}


// Toggles the state of the automatic spelling correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticSpellingCorrection(_:)
func (t_ TextView) ToggleAutomaticSpellingCorrection(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleAutomaticSpellingCorrection:"), sender)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticTextCompletion(_:)
func (t_ TextView) ToggleAutomaticTextCompletion(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleAutomaticTextCompletion:"), sender)
}


// Toggles the state of the automatic text replacement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticTextReplacement(_:)
func (t_ TextView) ToggleAutomaticTextReplacement(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleAutomaticTextReplacement:"), sender)
}


// Toggles whether continuous spell checking is enabled for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleContinuousSpellChecking(_:)
func (t_ TextView) ToggleContinuousSpellChecking(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleContinuousSpellChecking:"), sender)
}


// Changes the state of grammar checking from enabled to disabled and vice versa.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleGrammarChecking(_:)
func (t_ TextView) ToggleGrammarChecking(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleGrammarChecking:"), sender)
}


// An action message that toggles the visibility state of the Quick Look preview panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleQuickLookPreviewPanel(_:)
func (t_ TextView) ToggleQuickLookPreviewPanel(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleQuickLookPreviewPanel:"), sender)
}


// Changes the state of smart insert and delete from enabled to disabled and vice versa.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleSmartInsertDelete(_:)
func (t_ TextView) ToggleSmartInsertDelete(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleSmartInsertDelete:"), sender)
}


// Sets the receiver to use nominal glyph spacing for the glyphs in its selection, or for all glyphs if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/turnOffKerning(_:)
func (t_ TextView) TurnOffKerning(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("turnOffKerning:"), sender)
}


// Sets the receiver to use only required ligatures when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/turnOffLigatures(_:)
func (t_ TextView) TurnOffLigatures(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("turnOffLigatures:"), sender)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/updateCandidates()
func (t_ TextView) UpdateCandidates() {
	objc.Send[objc.ID](t_.ID, objc.Sel("updateCandidates"))
}


// Updates the acceptable drag types of all text views associated with the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/updateDragTypeRegistration()
func (t_ TextView) UpdateDragTypeRegistration() {
	objc.Send[objc.ID](t_.ID, objc.Sel("updateDragTypeRegistration"))
}


// Updates the Font panel to contain the font attributes of the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/updateFontPanel()
func (t_ TextView) UpdateFontPanel() {
	objc.Send[objc.ID](t_.ID, objc.Sel("updateFontPanel"))
}


// Updates the insertion point’s location and optionally restarts the blinking cursor timer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/updateInsertionPointStateAndRestartTimer(_:)
func (t_ TextView) UpdateInsertionPointStateAndRestartTimer(restartFlag bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("updateInsertionPointStateAndRestartTimer:"), restartFlag)
}


// Notifies the QuickLook panel that an update may be required.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/updateQuickLookPreviewPanel()
func (t_ TextView) UpdateQuickLookPreviewPanel() {
	objc.Send[objc.ID](t_.ID, objc.Sel("updateQuickLookPreviewPanel"))
}


// Updates the ruler view in the receiver’s enclosing scroll view to reflect the selection’s paragraph and marker attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/updateRuler()
func (t_ TextView) UpdateRuler() {
	objc.Send[objc.ID](t_.ID, objc.Sel("updateRuler"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/updateTextTouchBarItems()
func (t_ TextView) UpdateTextTouchBarItems() {
	objc.Send[objc.ID](t_.ID, objc.Sel("updateTextTouchBarItems"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/updateTouchBarItemIdentifiers()
func (t_ TextView) UpdateTouchBarItemIdentifiers() {
	objc.Send[objc.ID](t_.ID, objc.Sel("updateTouchBarItemIdentifiers"))
}


// Sets the receiver to use all ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/useAllLigatures(_:)
func (t_ TextView) UseAllLigatures(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("useAllLigatures:"), sender)
}


// Set the receiver to use pair kerning data for the glyphs in its selection, or for all glyphs if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/useStandardKerning(_:)
func (t_ TextView) UseStandardKerning(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("useStandardKerning:"), sender)
}


// Sets the receiver to use the standard ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/useStandardLigatures(_:)
func (t_ TextView) UseStandardLigatures(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("useStandardLigatures:"), sender)
}


// Returns if the text view can provide and accept the specified data types, or if it can’t.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/validRequestor(forSendType:returnType:)
func (t_ TextView) ValidRequestorForSendTypeReturnType(sendType objc.IObject /* cross-framework: PasteboardType */, returnType objc.IObject /* cross-framework: PasteboardType */) objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("validRequestorForSendType:returnType:"), sendType, returnType)
	return rv
}


// Writes the current selection to the specified pasteboard using the given type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/writeSelection(to:type:)
func (t_ TextView) WriteSelectionToPasteboardType(pboard IPasteboard, type_ objc.IObject /* cross-framework: PasteboardType */) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("writeSelectionToPasteboard:type:"), pboard, type_)
	return rv
}


// Writes the current selection to the specified pasteboard under each given type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/writeSelection(to:types:)
func (t_ TextView) WriteSelectionToPasteboardTypes(pboard IPasteboard, types []string) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("writeSelectionToPasteboard:types:"), pboard, types)
	return rv
}


// The data types that the receiver accepts as the destination view of a dragging operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/acceptableDragTypes
func (t_ TextView) AcceptableDragTypes() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("acceptableDragTypes"))
	return rv
}


// A Boolean value that indicates whether the receiver accepts the glyph info attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/acceptsGlyphInfo
func (t_ TextView) AcceptsGlyphInfo() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("acceptsGlyphInfo"))
	return rv
}


// A Boolean value that indicates whether the receiver accepts the glyph info attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/acceptsGlyphInfo
func (t_ TextView) SetAcceptsGlyphInfo(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAcceptsGlyphInfo:"), value)
}


// An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowedInputSourceLocales
func (t_ TextView) AllowedInputSourceLocales() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("allowedInputSourceLocales"))
	return rv
}


// An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowedInputSourceLocales
func (t_ TextView) SetAllowedInputSourceLocales(value []string) {
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
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedInputSourceLocales:"), nsArray)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowedWritingToolsResultOptions
func (t_ TextView) AllowedWritingToolsResultOptions() WritingToolsResultOptions {
	rv := objc.Send[WritingToolsResultOptions](t_.ID, objc.Sel("allowedWritingToolsResultOptions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowedWritingToolsResultOptions
func (t_ TextView) SetAllowedWritingToolsResultOptions(value WritingToolsResultOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedWritingToolsResultOptions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowsCharacterPickerTouchBarItem
func (t_ TextView) AllowsCharacterPickerTouchBarItem() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsCharacterPickerTouchBarItem"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowsCharacterPickerTouchBarItem
func (t_ TextView) SetAllowsCharacterPickerTouchBarItem(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsCharacterPickerTouchBarItem:"), value)
}


// A Boolean value that indicates whether the receiver allows its background color to change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowsDocumentBackgroundColorChange
func (t_ TextView) AllowsDocumentBackgroundColorChange() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsDocumentBackgroundColorChange"))
	return rv
}


// A Boolean value that indicates whether the receiver allows its background color to change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowsDocumentBackgroundColorChange
func (t_ TextView) SetAllowsDocumentBackgroundColorChange(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsDocumentBackgroundColorChange:"), value)
}


// Indicates whether image attachments should permit editing of their images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowsImageEditing
func (t_ TextView) AllowsImageEditing() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsImageEditing"))
	return rv
}


// Indicates whether image attachments should permit editing of their images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowsImageEditing
func (t_ TextView) SetAllowsImageEditing(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsImageEditing:"), value)
}


// A Boolean value that indicates whether the receiver allows undo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowsUndo
func (t_ TextView) AllowsUndo() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsUndo"))
	return rv
}


// A Boolean value that indicates whether the receiver allows undo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowsUndo
func (t_ TextView) SetAllowsUndo(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsUndo:"), value)
}


// The receiver’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/backgroundColor
func (t_ TextView) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The receiver’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/backgroundColor
func (t_ TextView) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/candidateListTouchBarItem
func (t_ TextView) CandidateListTouchBarItem() ICandidateListTouchBarItem {
	rv := objc.Send[CandidateListTouchBarItem](t_.ID, objc.Sel("candidateListTouchBarItem"))
	return rv
}


// The receiver’s default paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/defaultParagraphStyle
func (t_ TextView) DefaultParagraphStyle() IParagraphStyle {
	rv := objc.Send[ParagraphStyle](t_.ID, objc.Sel("defaultParagraphStyle"))
	return rv
}


// The receiver’s default paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/defaultParagraphStyle
func (t_ TextView) SetDefaultParagraphStyle(value IParagraphStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDefaultParagraphStyle:"), value)
}


// The delegate for all text views sharing the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/delegate
func (t_ TextView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate for all text views sharing the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/delegate
func (t_ TextView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/displaysLinkToolTips
func (t_ TextView) DisplaysLinkToolTips() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("displaysLinkToolTips"))
	return rv
}


// A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/displaysLinkToolTips
func (t_ TextView) SetDisplaysLinkToolTips(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDisplaysLinkToolTips:"), value)
}


// A Boolean value that indicates whether the receiver draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/drawsBackground
func (t_ TextView) DrawsBackground() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("drawsBackground"))
	return rv
}


// A Boolean value that indicates whether the receiver draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/drawsBackground
func (t_ TextView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsBackground:"), value)
}


// The default text checking types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/enabledTextCheckingTypes
func (t_ TextView) EnabledTextCheckingTypes() TextCheckingTypes /* not a class type */ {
	rv := objc.Send[TextCheckingTypes](t_.ID, objc.Sel("enabledTextCheckingTypes"))
	return rv
}


// The default text checking types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/enabledTextCheckingTypes
func (t_ TextView) SetEnabledTextCheckingTypes(value TextCheckingTypes /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEnabledTextCheckingTypes:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to import files by dragging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/importsGraphics
func (t_ TextView) ImportsGraphics() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("importsGraphics"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to import files by dragging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/importsGraphics
func (t_ TextView) SetImportsGraphics(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImportsGraphics:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/inlinePredictionType
func (t_ TextView) InlinePredictionType() TextInputTraitType {
	rv := objc.Send[TextInputTraitType](t_.ID, objc.Sel("inlinePredictionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/inlinePredictionType
func (t_ TextView) SetInlinePredictionType(value TextInputTraitType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInlinePredictionType:"), value)
}


// The color of the insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/insertionPointColor
func (t_ TextView) InsertionPointColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("insertionPointColor"))
	return rv
}


// The color of the insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/insertionPointColor
func (t_ TextView) SetInsertionPointColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInsertionPointColor:"), value)
}


// A Boolean value that indicates whether automatic dash substitution is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticDashSubstitutionEnabled
func (t_ TextView) AutomaticDashSubstitutionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticDashSubstitutionEnabled"))
	return rv
}


// A Boolean value that indicates whether automatic dash substitution is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticDashSubstitutionEnabled
func (t_ TextView) SetAutomaticDashSubstitutionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticDashSubstitutionEnabled:"), value)
}


// A Boolean value that indicates whether automatic data detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticDataDetectionEnabled
func (t_ TextView) AutomaticDataDetectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticDataDetectionEnabled"))
	return rv
}


// A Boolean value that indicates whether automatic data detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticDataDetectionEnabled
func (t_ TextView) SetAutomaticDataDetectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticDataDetectionEnabled:"), value)
}


// A Boolean value that enables or disables automatic link detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticLinkDetectionEnabled
func (t_ TextView) AutomaticLinkDetectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticLinkDetectionEnabled"))
	return rv
}


// A Boolean value that enables or disables automatic link detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticLinkDetectionEnabled
func (t_ TextView) SetAutomaticLinkDetectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticLinkDetectionEnabled:"), value)
}


// A Boolean value that enables and disables automatic quotation mark substitution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticQuoteSubstitutionEnabled
func (t_ TextView) AutomaticQuoteSubstitutionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticQuoteSubstitutionEnabled"))
	return rv
}


// A Boolean value that enables and disables automatic quotation mark substitution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticQuoteSubstitutionEnabled
func (t_ TextView) SetAutomaticQuoteSubstitutionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticQuoteSubstitutionEnabled:"), value)
}


// A Boolean value that indicates whether automatic spelling correction is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticSpellingCorrectionEnabled
func (t_ TextView) AutomaticSpellingCorrectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticSpellingCorrectionEnabled"))
	return rv
}


// A Boolean value that indicates whether automatic spelling correction is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticSpellingCorrectionEnabled
func (t_ TextView) SetAutomaticSpellingCorrectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticSpellingCorrectionEnabled:"), value)
}


// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticTextCompletionEnabled
func (t_ TextView) AutomaticTextCompletionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticTextCompletionEnabled"))
	return rv
}


// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticTextCompletionEnabled
func (t_ TextView) SetAutomaticTextCompletionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticTextCompletionEnabled:"), value)
}


// A Boolean value that indicates whether automatic text replacement is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticTextReplacementEnabled
func (t_ TextView) AutomaticTextReplacementEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticTextReplacementEnabled"))
	return rv
}


// A Boolean value that indicates whether automatic text replacement is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticTextReplacementEnabled
func (t_ TextView) SetAutomaticTextReplacementEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticTextReplacementEnabled:"), value)
}


// A Boolean value that indicates whether undo coalescing is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isCoalescingUndo
func (t_ TextView) CoalescingUndo() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("coalescingUndo"))
	return rv
}


// A Boolean value that indicates whether the receiver has continuous spell checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isContinuousSpellCheckingEnabled
func (t_ TextView) ContinuousSpellCheckingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("continuousSpellCheckingEnabled"))
	return rv
}


// A Boolean value that indicates whether the receiver has continuous spell checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isContinuousSpellCheckingEnabled
func (t_ TextView) SetContinuousSpellCheckingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContinuousSpellCheckingEnabled:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isEditable
func (t_ TextView) Editable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("editable"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isEditable
func (t_ TextView) SetEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEditable:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isFieldEditor
func (t_ TextView) FieldEditor() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("fieldEditor"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isFieldEditor
func (t_ TextView) SetFieldEditor(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFieldEditor:"), value)
}


// Enables and disables grammar checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isGrammarCheckingEnabled
func (t_ TextView) GrammarCheckingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("grammarCheckingEnabled"))
	return rv
}


// Enables and disables grammar checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isGrammarCheckingEnabled
func (t_ TextView) SetGrammarCheckingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGrammarCheckingEnabled:"), value)
}


// A Boolean value that indicates whether incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isIncrementalSearchingEnabled
func (t_ TextView) IncrementalSearchingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("incrementalSearchingEnabled"))
	return rv
}


// A Boolean value that indicates whether incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isIncrementalSearchingEnabled
func (t_ TextView) SetIncrementalSearchingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIncrementalSearchingEnabled:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isRichText
func (t_ TextView) RichText() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("richText"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isRichText
func (t_ TextView) SetRichText(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRichText:"), value)
}


// A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isRulerVisible
func (t_ TextView) RulerVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rulerVisible"))
	return rv
}


// A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isRulerVisible
func (t_ TextView) SetRulerVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRulerVisible:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isSelectable
func (t_ TextView) Selectable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("selectable"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isSelectable
func (t_ TextView) SetSelectable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectable:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isWritingToolsActive
func (t_ TextView) WritingToolsActive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("writingToolsActive"))
	return rv
}


// The layout manager that lays out text for the receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/layoutManager
func (t_ TextView) LayoutManager() ILayoutManager {
	rv := objc.Send[LayoutManager](t_.ID, objc.Sel("layoutManager"))
	return rv
}


// The attributes used to draw the onscreen presentation of link text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/linkTextAttributes
func (t_ TextView) LinkTextAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("linkTextAttributes"))
	return rv
}


// The attributes used to draw the onscreen presentation of link text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/linkTextAttributes
func (t_ TextView) SetLinkTextAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLinkTextAttributes:"), value)
}


// The attributes used to draw marked text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/markedTextAttributes
func (t_ TextView) MarkedTextAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("markedTextAttributes"))
	return rv
}


// The attributes used to draw marked text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/markedTextAttributes
func (t_ TextView) SetMarkedTextAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMarkedTextAttributes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/mathExpressionCompletionType
func (t_ TextView) MathExpressionCompletionType() TextInputTraitType {
	rv := objc.Send[TextInputTraitType](t_.ID, objc.Sel("mathExpressionCompletionType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/mathExpressionCompletionType
func (t_ TextView) SetMathExpressionCompletionType(value TextInputTraitType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMathExpressionCompletionType:"), value)
}


// The range of characters affected by an action method that changes character (not paragraph) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rangeForUserCharacterAttributeChange
func (t_ TextView) RangeForUserCharacterAttributeChange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("rangeForUserCharacterAttributeChange"))
	return rv
}


// The partial range from the most recent beginning of a word up to the insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rangeForUserCompletion
func (t_ TextView) RangeForUserCompletion() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("rangeForUserCompletion"))
	return rv
}


// The range of characters affected by an action method that changes paragraph (not character) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rangeForUserParagraphAttributeChange
func (t_ TextView) RangeForUserParagraphAttributeChange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("rangeForUserParagraphAttributeChange"))
	return rv
}


// The range of characters affected by a method that changes characters (as opposed to attributes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rangeForUserTextChange
func (t_ TextView) RangeForUserTextChange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("rangeForUserTextChange"))
	return rv
}


// An array containing the ranges of characters affected by an action method that changes character (not paragraph) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rangesForUserCharacterAttributeChange
func (t_ TextView) RangesForUserCharacterAttributeChange() []foundation.Value {
	rv := objc.Send[[]foundation.Value](t_.ID, objc.Sel("rangesForUserCharacterAttributeChange"))
	return rv
}


// An array containing the ranges of characters affected by a method that changes paragraph (not character) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rangesForUserParagraphAttributeChange
func (t_ TextView) RangesForUserParagraphAttributeChange() []foundation.Value {
	rv := objc.Send[[]foundation.Value](t_.ID, objc.Sel("rangesForUserParagraphAttributeChange"))
	return rv
}


// An array containing the ranges of characters affected by a method that changes characters (as opposed to attributes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rangesForUserTextChange
func (t_ TextView) RangesForUserTextChange() []foundation.Value {
	rv := objc.Send[[]foundation.Value](t_.ID, objc.Sel("rangesForUserTextChange"))
	return rv
}


// The types this text view can read immediately from the pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/readablePasteboardTypes
func (t_ TextView) ReadablePasteboardTypes() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("readablePasteboardTypes"))
	return rv
}


// An array containing the ranges of characters selected in the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectedRanges
func (t_ TextView) SelectedRanges() []foundation.Value {
	rv := objc.Send[[]foundation.Value](t_.ID, objc.Sel("selectedRanges"))
	return rv
}


// An array containing the ranges of characters selected in the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectedRanges
func (t_ TextView) SetSelectedRanges(value []foundation.Value) {
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


// The attributes used to indicate the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectedTextAttributes
func (t_ TextView) SelectedTextAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("selectedTextAttributes"))
	return rv
}


// The attributes used to indicate the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectedTextAttributes
func (t_ TextView) SetSelectedTextAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedTextAttributes:"), value)
}


// The preferred direction of selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectionAffinity
func (t_ TextView) SelectionAffinity() SelectionAffinity {
	rv := objc.Send[SelectionAffinity](t_.ID, objc.Sel("selectionAffinity"))
	return rv
}


// The selection granularity for subsequent extension of a selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectionGranularity
func (t_ TextView) SelectionGranularity() SelectionGranularity {
	rv := objc.Send[SelectionGranularity](t_.ID, objc.Sel("selectionGranularity"))
	return rv
}


// The selection granularity for subsequent extension of a selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectionGranularity
func (t_ TextView) SetSelectionGranularity(value SelectionGranularity) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionGranularity:"), value)
}


// A Boolean value that determines whether the receiver should draw its insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/shouldDrawInsertionPoint
func (t_ TextView) ShouldDrawInsertionPoint() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("shouldDrawInsertionPoint"))
	return rv
}


// A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/smartInsertDeleteEnabled
func (t_ TextView) SmartInsertDeleteEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("smartInsertDeleteEnabled"))
	return rv
}


// A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/smartInsertDeleteEnabled
func (t_ TextView) SetSmartInsertDeleteEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSmartInsertDeleteEnabled:"), value)
}


// A tag identifying the text view’s text as a document for the spell checker server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/spellCheckerDocumentTag
func (t_ TextView) SpellCheckerDocumentTag() int {
	rv := objc.Send[int](t_.ID, objc.Sel("spellCheckerDocumentTag"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/stronglyReferencesTextStorage
func (t_ TextView) StronglyReferencesTextStorage() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("stronglyReferencesTextStorage"))
	return rv
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


// The empty space the receiver leaves around its associated text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContainerInset
func (t_ TextView) TextContainerInset() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](t_.ID, objc.Sel("textContainerInset"))
	return rv
}


// The empty space the receiver leaves around its associated text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContainerInset
func (t_ TextView) SetTextContainerInset(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainerInset:"), value)
}


// The origin of the receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContainerOrigin
func (t_ TextView) TextContainerOrigin() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[corefoundation.Point](t_.ID, objc.Sel("textContainerOrigin"))
	return rv
}


// The receiver’s text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContentStorage
func (t_ TextView) TextContentStorage() ITextContentStorage {
	rv := objc.Send[TextContentStorage](t_.ID, objc.Sel("textContentStorage"))
	return rv
}


// ************************* Text Highlight support **************************
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textHighlightAttributes
func (t_ TextView) TextHighlightAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("textHighlightAttributes"))
	return rv
}


// ************************* Text Highlight support **************************
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textHighlightAttributes
func (t_ TextView) SetTextHighlightAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextHighlightAttributes:"), value)
}


// The manager that lays out text for the receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textLayoutManager
func (t_ TextView) TextLayoutManager() ITextLayoutManager {
	rv := objc.Send[TextLayoutManager](t_.ID, objc.Sel("textLayoutManager"))
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
func (t_ TextView) TypingAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("typingAttributes"))
	return rv
}


// The receiver’s typing attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/typingAttributes
func (t_ TextView) SetTypingAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTypingAttributes:"), value)
}


// A Boolean value that indicates whether the framework should use adaptive color mapping for dark appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesAdaptiveColorMappingForDarkAppearance
func (t_ TextView) UsesAdaptiveColorMappingForDarkAppearance() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesAdaptiveColorMappingForDarkAppearance"))
	return rv
}


// A Boolean value that indicates whether the framework should use adaptive color mapping for dark appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesAdaptiveColorMappingForDarkAppearance
func (t_ TextView) SetUsesAdaptiveColorMappingForDarkAppearance(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesAdaptiveColorMappingForDarkAppearance:"), value)
}


// A Boolean value that indicates whether to use the find bar for this text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesFindBar
func (t_ TextView) UsesFindBar() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFindBar"))
	return rv
}


// A Boolean value that indicates whether to use the find bar for this text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesFindBar
func (t_ TextView) SetUsesFindBar(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFindBar:"), value)
}


// A Boolean value that indicates whether the receiver allows for a find panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesFindPanel
func (t_ TextView) UsesFindPanel() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFindPanel"))
	return rv
}


// A Boolean value that indicates whether the receiver allows for a find panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesFindPanel
func (t_ TextView) SetUsesFindPanel(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFindPanel:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager use the Font panel and Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesFontPanel
func (t_ TextView) UsesFontPanel() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontPanel"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager use the Font panel and Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesFontPanel
func (t_ TextView) SetUsesFontPanel(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontPanel:"), value)
}


// A Boolean value that indicates whether this text view uses the inspector bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesInspectorBar
func (t_ TextView) UsesInspectorBar() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesInspectorBar"))
	return rv
}


// A Boolean value that indicates whether this text view uses the inspector bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesInspectorBar
func (t_ TextView) SetUsesInspectorBar(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesInspectorBar:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesRolloverButtonForSelection
func (t_ TextView) UsesRolloverButtonForSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesRolloverButtonForSelection"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesRolloverButtonForSelection
func (t_ TextView) SetUsesRolloverButtonForSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesRolloverButtonForSelection:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesRuler
func (t_ TextView) UsesRuler() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesRuler"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesRuler
func (t_ TextView) SetUsesRuler(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesRuler:"), value)
}


// The pasteboard types that can be provided from the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/writablePasteboardTypes
func (t_ TextView) WritablePasteboardTypes() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("writablePasteboardTypes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/writingToolsBehavior
func (t_ TextView) WritingToolsBehavior() WritingToolsBehavior {
	rv := objc.Send[WritingToolsBehavior](t_.ID, objc.Sel("writingToolsBehavior"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/writingToolsBehavior
func (t_ TextView) SetWritingToolsBehavior(value WritingToolsBehavior) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWritingToolsBehavior:"), value)
}


// A Boolean value that indicates whether automatic dash substitution is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdashsubstitutionenabled
func (t_ TextView) IsAutomaticDashSubstitutionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticDashSubstitutionEnabled"))
	return rv
}


// A Boolean value that indicates whether automatic dash substitution is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdashsubstitutionenabled
func (t_ TextView) SetIsAutomaticDashSubstitutionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticDashSubstitutionEnabled:"), value)
}


// A Boolean value that indicates whether automatic data detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdatadetectionenabled
func (t_ TextView) IsAutomaticDataDetectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticDataDetectionEnabled"))
	return rv
}


// A Boolean value that indicates whether automatic data detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdatadetectionenabled
func (t_ TextView) SetIsAutomaticDataDetectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticDataDetectionEnabled:"), value)
}


// A Boolean value that enables or disables automatic link detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticlinkdetectionenabled
func (t_ TextView) IsAutomaticLinkDetectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticLinkDetectionEnabled"))
	return rv
}


// A Boolean value that enables or disables automatic link detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticlinkdetectionenabled
func (t_ TextView) SetIsAutomaticLinkDetectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticLinkDetectionEnabled:"), value)
}


// A Boolean value that enables and disables automatic quotation mark substitution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticquotesubstitutionenabled
func (t_ TextView) IsAutomaticQuoteSubstitutionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticQuoteSubstitutionEnabled"))
	return rv
}


// A Boolean value that enables and disables automatic quotation mark substitution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticquotesubstitutionenabled
func (t_ TextView) SetIsAutomaticQuoteSubstitutionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticQuoteSubstitutionEnabled:"), value)
}


// A Boolean value that indicates whether automatic spelling correction is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticspellingcorrectionenabled
func (t_ TextView) IsAutomaticSpellingCorrectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticSpellingCorrectionEnabled"))
	return rv
}


// A Boolean value that indicates whether automatic spelling correction is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticspellingcorrectionenabled
func (t_ TextView) SetIsAutomaticSpellingCorrectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticSpellingCorrectionEnabled:"), value)
}


// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextcompletionenabled
func (t_ TextView) IsAutomaticTextCompletionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticTextCompletionEnabled"))
	return rv
}


// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextcompletionenabled
func (t_ TextView) SetIsAutomaticTextCompletionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticTextCompletionEnabled:"), value)
}


// A Boolean value that indicates whether automatic text replacement is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextreplacementenabled
func (t_ TextView) IsAutomaticTextReplacementEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticTextReplacementEnabled"))
	return rv
}


// A Boolean value that indicates whether automatic text replacement is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextreplacementenabled
func (t_ TextView) SetIsAutomaticTextReplacementEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticTextReplacementEnabled:"), value)
}


// A Boolean value that indicates whether undo coalescing is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscoalescingundo
func (t_ TextView) IsCoalescingUndo() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isCoalescingUndo"))
	return rv
}


// A Boolean value that indicates whether undo coalescing is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscoalescingundo
func (t_ TextView) SetIsCoalescingUndo(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsCoalescingUndo:"), value)
}


// A Boolean value that indicates whether the receiver has continuous spell checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscontinuousspellcheckingenabled
func (t_ TextView) IsContinuousSpellCheckingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isContinuousSpellCheckingEnabled"))
	return rv
}


// A Boolean value that indicates whether the receiver has continuous spell checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscontinuousspellcheckingenabled
func (t_ TextView) SetIsContinuousSpellCheckingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsContinuousSpellCheckingEnabled:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iseditable
func (t_ TextView) IsEditable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEditable"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iseditable
func (t_ TextView) SetIsEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEditable:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isfieldeditor
func (t_ TextView) IsFieldEditor() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isFieldEditor"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isfieldeditor
func (t_ TextView) SetIsFieldEditor(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsFieldEditor:"), value)
}


// Enables and disables grammar checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isgrammarcheckingenabled
func (t_ TextView) IsGrammarCheckingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isGrammarCheckingEnabled"))
	return rv
}


// Enables and disables grammar checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isgrammarcheckingenabled
func (t_ TextView) SetIsGrammarCheckingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsGrammarCheckingEnabled:"), value)
}


// A Boolean value that indicates whether incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isincrementalsearchingenabled
func (t_ TextView) IsIncrementalSearchingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isIncrementalSearchingEnabled"))
	return rv
}


// A Boolean value that indicates whether incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isincrementalsearchingenabled
func (t_ TextView) SetIsIncrementalSearchingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsIncrementalSearchingEnabled:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrichtext
func (t_ TextView) IsRichText() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRichText"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrichtext
func (t_ TextView) SetIsRichText(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRichText:"), value)
}


// A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrulervisible
func (t_ TextView) IsRulerVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRulerVisible"))
	return rv
}


// A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrulervisible
func (t_ TextView) SetIsRulerVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRulerVisible:"), value)
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isselectable
func (t_ TextView) IsSelectable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSelectable"))
	return rv
}


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isselectable
func (t_ TextView) SetIsSelectable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsSelectable:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iswritingtoolsactive
func (t_ TextView) IsWritingToolsActive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isWritingToolsActive"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iswritingtoolsactive
func (t_ TextView) SetIsWritingToolsActive(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsWritingToolsActive:"), value)
}


