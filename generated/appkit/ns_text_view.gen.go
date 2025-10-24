// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSTextView */


/* debug [class_header]: Header for NSTextView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextView */
// An interface definition for the [TextView] class.
type ITextView interface {
	IText
	
/* debug [class_interface_properties]: Properties for TextView */
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
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
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
	TextContainerInset() Size /* not a class type */
	SetTextContainerInset(value Size /* not a class type */)
	TextContainerOrigin() vision.Point
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextView */
	// methods:
	AlignJustified(sender objc.IObject)
	BreakUndoCoalescing()
	ChangeAttributes(sender objc.IObject)
	ChangeColor(sender objc.IObject)
	ChangeDocumentBackgroundColor(sender objc.IObject)
	ChangeLayoutOrientation(sender objc.IObject)
	CharacterIndexForInsertionAtPoint(point vision.Point) uint
	CheckTextInRangeTypesOptions(range_ corefoundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary)
	CheckTextInDocument(sender objc.IObject)
	CheckTextInSelection(sender objc.IObject)
	CleanUpAfterDragOperation()
	ClickedOnLinkAtIndex(link objc.IObject, charIndex uint)
	Complete(sender objc.IObject)
	CompletionsForPartialWordRangeIndexOfSelectedItem(charRange corefoundation.Range, index int) []string
	DidChangeText()
	DragImageForSelectionWithEventOrigin(event IEvent, origin PointPointer /* not a class type */) IImage
	DragOperationForDraggingInfoType(dragInfo unsafe.Pointer, type_ PasteboardType /* typedef */) DragOperation
	DragSelectionWithEventOffsetSlideBack(event IEvent, mouseOffset Size /* not a class type */, slideBack bool) bool
	DrawViewBackgroundInRect(rect Rect /* not a class type */)
	DrawInsertionPointInRectColorTurnedOn(rect Rect /* not a class type */, color IColor, flag bool)
	DrawTextHighlightBackgroundForTextRangeOrigin(textRange ITextRange, origin vision.Point)
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
	PreferredPasteboardTypeFromArrayRestrictedToTypesFromArray(availableTypes []string, allowedTypes []string) PasteboardType /* typedef */
	QuickLookPreviewableItemsInRanges(ranges []foundation.Value) []objc.ID
	RaiseBaseline(sender objc.IObject)
	ReadSelectionFromPasteboard(pboard IPasteboard) bool
	ReadSelectionFromPasteboardType(pboard IPasteboard, type_ PasteboardType /* typedef */) bool
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
	SetConstrainedFrameSize(desiredSize Size /* not a class type */)
	SetLayoutOrientation(orientation TextLayoutOrientation)
	SetNeedsDisplayInRectAvoidAdditionalLayout(rect Rect /* not a class type */, flag bool)
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
	ValidRequestorForSendTypeReturnType(sendType PasteboardType /* typedef */, returnType PasteboardType /* typedef */) objc.ID
	WriteSelectionToPasteboardType(pboard IPasteboard, type_ PasteboardType /* typedef */) bool
	WriteSelectionToPasteboardTypes(pboard IPasteboard, types []string) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextView */
// Alloc allocates a new instance without initialization.
func (tc _TextViewClass) Alloc() TextView {
	rv := objc.Send[TextView](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextView */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextView */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/init(usingTextLayoutManager:)
func NewTextViewUsingTextLayoutManager(usingTextLayoutManager bool) TextView {
	instance := getTextViewClass().Alloc()
	rv := objc.Send[TextView](instance.ID, objc.Sel("initUsingTextLayoutManager:"), usingTextLayoutManager)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextViewUsingTextLayoutManager */


// Initializes a text view with data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/init(coder:)
func NewTextViewWithCoder(coder foundation.Coder) TextView {
	instance := getTextViewClass().Alloc()
	rv := objc.Send[TextView](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextViewWithCoder */


// Initializes a text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/init(frame:)
func NewTextViewWithFrame(frameRect Rect /* not a class type */) TextView {
	instance := getTextViewClass().Alloc()
	rv := objc.Send[TextView](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextViewWithFrame */


// Initializes a text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/init(frame:textContainer:)
func NewTextViewWithFrameTextContainer(frameRect Rect /* not a class type */, container ITextContainer) TextView {
	instance := getTextViewClass().Alloc()
	rv := objc.Send[TextView](instance.ID, objc.Sel("initWithFrame:textContainer:"), frameRect, container)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextViewWithFrameTextContainer */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextView */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/fieldEditor()
func (tc _TextViewClass) FieldEditor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("fieldEditor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FieldEditor) */


// Registers send and return types for the Services facility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/registerForServices()
func (tc _TextViewClass) RegisterForServices() {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("registerForServices"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RegisterForServices) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/scrollableDocumentContentTextView()
func (tc _TextViewClass) ScrollableDocumentContentTextView() IScrollView {
	rv := objc.Send[ScrollView](objc.ID(tc.class), objc.Sel("scrollableDocumentContentTextView"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ScrollableDocumentContentTextView) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/scrollablePlainDocumentContentTextView()
func (tc _TextViewClass) ScrollablePlainDocumentContentTextView() IScrollView {
	rv := objc.Send[ScrollView](objc.ID(tc.class), objc.Sel("scrollablePlainDocumentContentTextView"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ScrollablePlainDocumentContentTextView) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/scrollableTextView()
func (tc _TextViewClass) ScrollableTextView() IScrollView {
	rv := objc.Send[ScrollView](objc.ID(tc.class), objc.Sel("scrollableTextView"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ScrollableTextView) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textViewUsingTextLayoutManager:
func (tc _TextViewClass) TextViewUsingTextLayoutManager(usingTextLayoutManager bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("textViewUsingTextLayoutManager:"), usingTextLayoutManager)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TextViewUsingTextLayoutManager) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextView */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/stronglyReferencesTextStorage
func (tc _TextViewClass) StronglyReferencesTextStorage() bool {
	rv := objc.Send[bool](objc.ID(tc.class), objc.Sel("stronglyReferencesTextStorage"))
	return rv
}/* debug [class_properties_class/property]: stronglyReferencesTextStorage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextView */

// Applies full justification to selected paragraphs (or all text, if the receiver is a plain text object).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/alignJustified(_:)
func (t_ TextView) AlignJustified(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("alignJustified:"), sender)
}/* debug [instance_methods/method]: AlignJustified */


// Informs the receiver that it should begin coalescing successive typing operations in a new undo grouping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/breakUndoCoalescing()
func (t_ TextView) BreakUndoCoalescing() {
	objc.Send[objc.ID](t_.ID, objc.Sel("breakUndoCoalescing"))
}/* debug [instance_methods/method]: BreakUndoCoalescing */


// Changes the attributes of the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/changeAttributes(_:)
func (t_ TextView) ChangeAttributes(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("changeAttributes:"), sender)
}/* debug [instance_methods/method]: ChangeAttributes */


// Sets the color of the selected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/changeColor(_:)
func (t_ TextView) ChangeColor(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("changeColor:"), sender)
}/* debug [instance_methods/method]: ChangeColor */


// An action method used to set the background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/changeDocumentBackgroundColor(_:)
func (t_ TextView) ChangeDocumentBackgroundColor(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("changeDocumentBackgroundColor:"), sender)
}/* debug [instance_methods/method]: ChangeDocumentBackgroundColor */


// An action method that sets the layout orientation of the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/changeLayoutOrientation(_:)
func (t_ TextView) ChangeLayoutOrientation(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("changeLayoutOrientation:"), sender)
}/* debug [instance_methods/method]: ChangeLayoutOrientation */


// Returns a character index appropriate for placing a zero-length selection for an insertion point associated with the mouse at the given point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/characterIndexForInsertion(at:)
func (t_ TextView) CharacterIndexForInsertionAtPoint(point vision.Point) uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("characterIndexForInsertionAtPoint:"), point)
	return rv
}/* debug [instance_methods/method]: CharacterIndexForInsertionAtPoint */


// Check and replace the text in the range using the specified checking types and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/checkText(in:types:options:)
func (t_ TextView) CheckTextInRangeTypesOptions(range_ corefoundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("checkTextInRange:types:options:"), range_, checkingTypes, options)
}/* debug [instance_methods/method]: CheckTextInRangeTypesOptions */


// Performs the default text checking on the entire document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/checkTextInDocument(_:)
func (t_ TextView) CheckTextInDocument(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("checkTextInDocument:"), sender)
}/* debug [instance_methods/method]: CheckTextInDocument */


// Performs the default text checking on the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/checkTextInSelection(_:)
func (t_ TextView) CheckTextInSelection(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("checkTextInSelection:"), sender)
}/* debug [instance_methods/method]: CheckTextInSelection */


// Releases the drag information still existing after the dragging session has completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/cleanUpAfterDragOperation()
func (t_ TextView) CleanUpAfterDragOperation() {
	objc.Send[objc.ID](t_.ID, objc.Sel("cleanUpAfterDragOperation"))
}/* debug [instance_methods/method]: CleanUpAfterDragOperation */


// Causes the text view to act as if the user clicked on some text with the given link as the value of a link attribute associated with the text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/clicked(onLink:at:)
func (t_ TextView) ClickedOnLinkAtIndex(link objc.IObject, charIndex uint) {
	objc.Send[objc.ID](t_.ID, objc.Sel("clickedOnLink:atIndex:"), link, charIndex)
}/* debug [instance_methods/method]: ClickedOnLinkAtIndex */


// Invokes completion in a text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/complete(_:)
func (t_ TextView) Complete(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("complete:"), sender)
}/* debug [instance_methods/method]: Complete */


// Returns an array of potential completions, in the order to be presented, representing possible word completions available from a partial word.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/completions(forPartialWordRange:indexOfSelectedItem:)
func (t_ TextView) CompletionsForPartialWordRangeIndexOfSelectedItem(charRange corefoundation.Range, index int) []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("completionsForPartialWordRange:indexOfSelectedItem:"), charRange, index)
	return rv
}/* debug [instance_methods/method]: CompletionsForPartialWordRangeIndexOfSelectedItem */


// Sends out necessary notifications when a text change completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/didChangeText()
func (t_ TextView) DidChangeText() {
	objc.Send[objc.ID](t_.ID, objc.Sel("didChangeText"))
}/* debug [instance_methods/method]: DidChangeText */


// Returns an appropriate drag image for the drag initiated by the specified event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/dragImageForSelection(with:origin:)
func (t_ TextView) DragImageForSelectionWithEventOrigin(event IEvent, origin PointPointer /* not a class type */) IImage {
	rv := objc.Send[Image](t_.ID, objc.Sel("dragImageForSelectionWithEvent:origin:"), event, origin)
	return rv
}/* debug [instance_methods/method]: DragImageForSelectionWithEventOrigin */


// Returns the type of drag operation that should be performed if the image were released now.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/dragOperation(for:type:)
func (t_ TextView) DragOperationForDraggingInfoType(dragInfo unsafe.Pointer, type_ PasteboardType /* typedef */) DragOperation {
	rv := objc.Send[DragOperation](t_.ID, objc.Sel("dragOperationForDraggingInfo:type:"), dragInfo, type_)
	return rv
}/* debug [instance_methods/method]: DragOperationForDraggingInfoType */


// Begins dragging the current selected text range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/dragSelection(with:offset:slideBack:)
func (t_ TextView) DragSelectionWithEventOffsetSlideBack(event IEvent, mouseOffset Size /* not a class type */, slideBack bool) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("dragSelectionWithEvent:offset:slideBack:"), event, mouseOffset, slideBack)
	return rv
}/* debug [instance_methods/method]: DragSelectionWithEventOffsetSlideBack */


// Draws the background of the text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/drawBackground(in:)
func (t_ TextView) DrawViewBackgroundInRect(rect Rect /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawViewBackgroundInRect:"), rect)
}/* debug [instance_methods/method]: DrawViewBackgroundInRect */


// Draws or erases the insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/drawInsertionPoint(in:color:turnedOn:)
func (t_ TextView) DrawInsertionPointInRectColorTurnedOn(rect Rect /* not a class type */, color IColor, flag bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawInsertionPointInRect:color:turnedOn:"), rect, color, flag)
}/* debug [instance_methods/method]: DrawInsertionPointInRectColorTurnedOn */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/drawTextHighlightBackground(for:origin:)
func (t_ TextView) DrawTextHighlightBackgroundForTextRangeOrigin(textRange ITextRange, origin vision.Point) {
	objc.Send[objc.ID](t_.ID, objc.Sel("drawTextHighlightBackgroundForTextRange:origin:"), textRange, origin)
}/* debug [instance_methods/method]: DrawTextHighlightBackgroundForTextRangeOrigin */


// Handles the text checking results returned by the text view
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/handleTextCheckingResults(_:forRange:types:options:orthography:wordCount:)
func (t_ TextView) HandleTextCheckingResultsForRangeTypesOptionsOrthographyWordCount(results []foundation.TextCheckingResult, range_ corefoundation.Range, checkingTypes TextCheckingTypes /* not a class type */, options foundation.IDictionary, orthography foundation.Orthography, wordCount int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("handleTextCheckingResults:forRange:types:options:orthography:wordCount:"), results, range_, checkingTypes, options, orthography, wordCount)
}/* debug [instance_methods/method]: HandleTextCheckingResultsForRangeTypesOptionsOrthographyWordCount */


// An action for toggling in the receiver’s selected range. The sender should be a menu item with a of type ( ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/highlight(_:)
func (t_ TextView) Highlight(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("highlight:"), sender)
}/* debug [instance_methods/method]: Highlight */


// Inserts the selected completion into the text at the appropriate location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/insertCompletion(_:forPartialWordRange:movement:isFinal:)
func (t_ TextView) InsertCompletionForPartialWordRangeMovementIsFinal(word objc.IObject /* cross-framework: NSString */, charRange corefoundation.Range, movement int, flag bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertCompletion:forPartialWordRange:movement:isFinal:"), word, charRange, movement, flag)
}/* debug [instance_methods/method]: InsertCompletionForPartialWordRangeMovementIsFinal */


// Invalidates the calculated origin of the text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/invalidateTextContainerOrigin()
func (t_ TextView) InvalidateTextContainerOrigin() {
	objc.Send[objc.ID](t_.ID, objc.Sel("invalidateTextContainerOrigin"))
}/* debug [instance_methods/method]: InvalidateTextContainerOrigin */


// Increases the space between glyphs in the receiver’s selection, or in all text if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/loosenKerning(_:)
func (t_ TextView) LoosenKerning(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("loosenKerning:"), sender)
}/* debug [instance_methods/method]: LoosenKerning */


// Lowers the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/lowerBaseline(_:)
func (t_ TextView) LowerBaseline(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("lowerBaseline:"), sender)
}/* debug [instance_methods/method]: LowerBaseline */


// Brings forward a panel allowing the user to manipulate links in the text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontLinkPanel(_:)
func (t_ TextView) OrderFrontLinkPanel(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("orderFrontLinkPanel:"), sender)
}/* debug [instance_methods/method]: OrderFrontLinkPanel */


// Brings forward a panel allowing the user to manipulate text lists in the text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontListPanel(_:)
func (t_ TextView) OrderFrontListPanel(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("orderFrontListPanel:"), sender)
}/* debug [instance_methods/method]: OrderFrontListPanel */


// Creates and displays a new instance of the sharing service picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontSharingServicePicker(_:)
func (t_ TextView) OrderFrontSharingServicePicker(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("orderFrontSharingServicePicker:"), sender)
}/* debug [instance_methods/method]: OrderFrontSharingServicePicker */


// Brings forward a panel allowing the user to manipulate text line heights, interline spacing, and paragraph spacing, in the text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontSpacingPanel(_:)
func (t_ TextView) OrderFrontSpacingPanel(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("orderFrontSpacingPanel:"), sender)
}/* debug [instance_methods/method]: OrderFrontSpacingPanel */


// Brings forward a panel allowing the user to specify string substitutions in the text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontSubstitutionsPanel(_:)
func (t_ TextView) OrderFrontSubstitutionsPanel(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("orderFrontSubstitutionsPanel:"), sender)
}/* debug [instance_methods/method]: OrderFrontSubstitutionsPanel */


// Brings forward a panel allowing the user to manipulate text tables in the text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontTablePanel(_:)
func (t_ TextView) OrderFrontTablePanel(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("orderFrontTablePanel:"), sender)
}/* debug [instance_methods/method]: OrderFrontTablePanel */


// Adds the outline attribute to the selected text attributes if absent; removes the attribute if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/outline(_:)
func (t_ TextView) Outline(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("outline:"), sender)
}/* debug [instance_methods/method]: Outline */


// Inserts the contents of the pasteboard into the receiver’s text as plain text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/pasteAsPlainText(_:)
func (t_ TextView) PasteAsPlainText(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("pasteAsPlainText:"), sender)
}/* debug [instance_methods/method]: PasteAsPlainText */


// This action method inserts the contents of the pasteboard into the receiver’s text as rich text, maintaining its attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/pasteAsRichText(_:)
func (t_ TextView) PasteAsRichText(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("pasteAsRichText:"), sender)
}/* debug [instance_methods/method]: PasteAsRichText */


// Performs a find panel action specified by the sender’s tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/performFindPanelAction(_:)
func (t_ TextView) PerformFindPanelAction(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("performFindPanelAction:"), sender)
}/* debug [instance_methods/method]: PerformFindPanelAction */


// Replaces text in the range you specify with the attributed string you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/performValidatedReplacement(in:with:)
func (t_ TextView) PerformValidatedReplacementInRangeWithAttributedString(range_ corefoundation.Range, attributedString foundation.AttributedString) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("performValidatedReplacementInRange:withAttributedString:"), range_, attributedString)
	return rv
}/* debug [instance_methods/method]: PerformValidatedReplacementInRangeWithAttributedString */


// Returns whatever type on the pasteboard would be most preferred for copying data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/preferredPasteboardType(from:restrictedToTypesFrom:)
func (t_ TextView) PreferredPasteboardTypeFromArrayRestrictedToTypesFromArray(availableTypes []string, allowedTypes []string) PasteboardType /* typedef */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("preferredPasteboardTypeFromArray:restrictedToTypesFromArray:"), availableTypes, allowedTypes)
	return rv
}/* debug [instance_methods/method]: PreferredPasteboardTypeFromArrayRestrictedToTypesFromArray */


// Returns an array of URLs for items that can be displayed by QuickLook in the specified ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/quickLookPreviewableItems(inRanges:)
func (t_ TextView) QuickLookPreviewableItemsInRanges(ranges []foundation.Value) []objc.ID {
	rv := objc.Send[[]objc.ID](t_.ID, objc.Sel("quickLookPreviewableItemsInRanges:"), ranges)
	return rv
}/* debug [instance_methods/method]: QuickLookPreviewableItemsInRanges */


// Raises the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/raiseBaseline(_:)
func (t_ TextView) RaiseBaseline(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("raiseBaseline:"), sender)
}/* debug [instance_methods/method]: RaiseBaseline */


// Reads the text view’s preferred type of data from the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/readSelection(from:)
func (t_ TextView) ReadSelectionFromPasteboard(pboard IPasteboard) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("readSelectionFromPasteboard:"), pboard)
	return rv
}/* debug [instance_methods/method]: ReadSelectionFromPasteboard */


// Reads data of the given type from the specified pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/readSelection(from:type:)
func (t_ TextView) ReadSelectionFromPasteboardType(pboard IPasteboard, type_ PasteboardType /* typedef */) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("readSelectionFromPasteboard:type:"), pboard, type_)
	return rv
}/* debug [instance_methods/method]: ReadSelectionFromPasteboardType */


// Replaces the text container for the group of text system objects containing the receiver, keeping the association between the receiver and its layout manager intact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/replaceTextContainer(_:)
func (t_ TextView) ReplaceTextContainer(newContainer ITextContainer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("replaceTextContainer:"), newContainer)
}/* debug [instance_methods/method]: ReplaceTextContainer */


// Modifies the paragraph style of the paragraphs containing the selection to accommodate a new marker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:didAdd:)
func (t_ TextView) RulerViewDidAddMarker(ruler IRulerView, marker IRulerMarker) {
	objc.Send[objc.ID](t_.ID, objc.Sel("rulerView:didAddMarker:"), ruler, marker)
}/* debug [instance_methods/method]: RulerViewDidAddMarker */


// Modifies the paragraph style of the paragraphs containing the selection to record the new location of the marker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:didMove:)
func (t_ TextView) RulerViewDidMoveMarker(ruler IRulerView, marker IRulerMarker) {
	objc.Send[objc.ID](t_.ID, objc.Sel("rulerView:didMoveMarker:"), ruler, marker)
}/* debug [instance_methods/method]: RulerViewDidMoveMarker */


// Modifies the paragraph style of the paragraphs containing the selection—if possible—by removing the specified marker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:didRemove:)
func (t_ TextView) RulerViewDidRemoveMarker(ruler IRulerView, marker IRulerMarker) {
	objc.Send[objc.ID](t_.ID, objc.Sel("rulerView:didRemoveMarker:"), ruler, marker)
}/* debug [instance_methods/method]: RulerViewDidRemoveMarker */


// Adds a left tab marker to the ruler at the location clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:handleMouseDownWith:)
func (t_ TextView) RulerViewHandleMouseDown(ruler IRulerView, event IEvent) {
	objc.Send[objc.ID](t_.ID, objc.Sel("rulerView:handleMouseDown:"), ruler, event)
}/* debug [instance_methods/method]: RulerViewHandleMouseDown */


// Returns whether a new marker can be added.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:shouldAdd:)
func (t_ TextView) RulerViewShouldAddMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rulerView:shouldAddMarker:"), ruler, marker)
	return rv
}/* debug [instance_methods/method]: RulerViewShouldAddMarker */


// Returns whether the marker should be moved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:shouldMove:)
func (t_ TextView) RulerViewShouldMoveMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rulerView:shouldMoveMarker:"), ruler, marker)
	return rv
}/* debug [instance_methods/method]: RulerViewShouldMoveMarker */


// Returns whether the marker should be removed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:shouldRemove:)
func (t_ TextView) RulerViewShouldRemoveMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rulerView:shouldRemoveMarker:"), ruler, marker)
	return rv
}/* debug [instance_methods/method]: RulerViewShouldRemoveMarker */


// Returns a potentially modified location to which the marker should be added.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:willAdd:atLocation:)
func (t_ TextView) RulerViewWillAddMarkerAtLocation(ruler IRulerView, marker IRulerMarker, location float64) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("rulerView:willAddMarker:atLocation:"), ruler, marker, location)
	return rv
}/* debug [instance_methods/method]: RulerViewWillAddMarkerAtLocation */


// Returns a potentially modified location to which the marker should be moved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rulerView(_:willMove:toLocation:)
func (t_ TextView) RulerViewWillMoveMarkerToLocation(ruler IRulerView, marker IRulerMarker, location float64) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("rulerView:willMoveMarker:toLocation:"), ruler, marker, location)
	return rv
}/* debug [instance_methods/method]: RulerViewWillMoveMarkerToLocation */


// Returns an adjusted selected range based on the selection granularity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectionRange(forProposedRange:granularity:)
func (t_ TextView) SelectionRangeForProposedRangeGranularity(proposedCharRange corefoundation.Range, granularity SelectionGranularity) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("selectionRangeForProposedRange:granularity:"), proposedCharRange, granularity)
	return rv
}/* debug [instance_methods/method]: SelectionRangeForProposedRangeGranularity */


// Sets the alignment of the paragraphs containing characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setAlignment(_:range:)
func (t_ TextView) SetAlignmentRange(alignment TextAlignment, range_ corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlignment:range:"), alignment, range_)
}/* debug [instance_methods/method]: SetAlignmentRange */


// Sets the base writing direction of a range of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setBaseWritingDirection(_:range:)
func (t_ TextView) SetBaseWritingDirectionRange(writingDirection WritingDirection, range_ corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBaseWritingDirection:range:"), writingDirection, range_)
}/* debug [instance_methods/method]: SetBaseWritingDirectionRange */


// Attempts to set the frame size as if by user action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setConstrainedFrameSize(_:)
func (t_ TextView) SetConstrainedFrameSize(desiredSize Size /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setConstrainedFrameSize:"), desiredSize)
}/* debug [instance_methods/method]: SetConstrainedFrameSize */


// Changes the receiver’s layout orientation and invalidates the contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setLayoutOrientation(_:)
func (t_ TextView) SetLayoutOrientation(orientation TextLayoutOrientation) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLayoutOrientation:"), orientation)
}/* debug [instance_methods/method]: SetLayoutOrientation */


// Marks the receiver as requiring display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setNeedsDisplay(_:avoidAdditionalLayout:)
func (t_ TextView) SetNeedsDisplayInRectAvoidAdditionalLayout(rect Rect /* not a class type */, flag bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setNeedsDisplayInRect:avoidAdditionalLayout:"), rect, flag)
}/* debug [instance_methods/method]: SetNeedsDisplayInRectAvoidAdditionalLayout */


// Selects the specified range of characters in response to user action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setSelectedRange(_:)
func (t_ TextView) SetSelectedRange(charRange corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedRange:"), charRange)
}/* debug [instance_methods/method]: SetSelectedRange */


// Sets the selection to a range of characters in response to user action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setSelectedRange(_:affinity:stillSelecting:)
func (t_ TextView) SetSelectedRangeAffinityStillSelecting(charRange corefoundation.Range, affinity SelectionAffinity, stillSelectingFlag bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedRange:affinity:stillSelecting:"), charRange, affinity, stillSelectingFlag)
}/* debug [instance_methods/method]: SetSelectedRangeAffinityStillSelecting */


// Sets the selection to the characters in an array of ranges in response to user action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setSelectedRanges(_:affinity:stillSelecting:)
func (t_ TextView) SetSelectedRangesAffinityStillSelecting(ranges []foundation.Value, affinity SelectionAffinity, stillSelectingFlag bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedRanges:affinity:stillSelecting:"), ranges, affinity, stillSelectingFlag)
}/* debug [instance_methods/method]: SetSelectedRangesAffinityStillSelecting */


// Sets the spelling state, which controls the display of the spelling and grammar indicators on the given text range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/setSpellingState(_:range:)
func (t_ TextView) SetSpellingStateRange(value int, charRange corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSpellingState:range:"), value, charRange)
}/* debug [instance_methods/method]: SetSpellingStateRange */


// Initiates a series of delegate messages (and general notifications) to determine whether modifications can be made to the characters and attributes of the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/shouldChangeText(in:replacementString:)
func (t_ TextView) ShouldChangeTextInRangeReplacementString(affectedCharRange corefoundation.Range, replacementString objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("shouldChangeTextInRange:replacementString:"), affectedCharRange, replacementString)
	return rv
}/* debug [instance_methods/method]: ShouldChangeTextInRangeReplacementString */


// Initiates a series of delegate messages (and general notifications) to determine whether modifications can be made to the characters and attributes of the receiver’s text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/shouldChangeText(inRanges:replacementStrings:)
func (t_ TextView) ShouldChangeTextInRangesReplacementStrings(affectedRanges []foundation.Value, replacementStrings []string) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("shouldChangeTextInRanges:replacementStrings:"), affectedRanges, replacementStrings)
	return rv
}/* debug [instance_methods/method]: ShouldChangeTextInRangesReplacementStrings */


// Causes a temporary highlighting effect to appear around the visible portion (or portions) of the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/showFindIndicator(for:)
func (t_ TextView) ShowFindIndicatorForRange(charRange corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("showFindIndicatorForRange:"), charRange)
}/* debug [instance_methods/method]: ShowFindIndicatorForRange */


// Returns an extended range that includes adjacent whitespace that should be deleted along with the proposed range in order to preserve proper spacing and punctuation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/smartDeleteRange(forProposedRange:)
func (t_ TextView) SmartDeleteRangeForProposedRange(proposedCharRange corefoundation.Range) corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("smartDeleteRangeForProposedRange:"), proposedCharRange)
	return rv
}/* debug [instance_methods/method]: SmartDeleteRangeForProposedRange */


// Returns any whitespace that needs to be added after the string to preserve proper spacing and punctuation when the string replaces the characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/smartInsert(afterStringFor:replacing:)
func (t_ TextView) SmartInsertAfterStringForStringReplacingRange(pasteString objc.IObject /* cross-framework: NSString */, charRangeToReplace corefoundation.Range) foundation.String {
	rv := objc.Send[foundation.String](t_.ID, objc.Sel("smartInsertAfterStringForString:replacingRange:"), pasteString, charRangeToReplace)
	return rv
}/* debug [instance_methods/method]: SmartInsertAfterStringForStringReplacingRange */


// Returns any whitespace that needs to be added before the string to preserve proper spacing and punctuation when the string replaces the characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/smartInsert(beforeStringFor:replacing:)
func (t_ TextView) SmartInsertBeforeStringForStringReplacingRange(pasteString objc.IObject /* cross-framework: NSString */, charRangeToReplace corefoundation.Range) foundation.String {
	rv := objc.Send[foundation.String](t_.ID, objc.Sel("smartInsertBeforeStringForString:replacingRange:"), pasteString, charRangeToReplace)
	return rv
}/* debug [instance_methods/method]: SmartInsertBeforeStringForStringReplacingRange */


// Determines whether whitespace needs to be added around the string to preserve proper spacing and punctuation when it replaces the characters in the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/smartInsert(for:replacing:before:after:)
func (t_ TextView) SmartInsertForStringReplacingRangeBeforeStringAfterString(pasteString objc.IObject /* cross-framework: NSString */, charRangeToReplace corefoundation.Range, beforeString objc.IObject /* cross-framework: NSString */, afterString objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("smartInsertForString:replacingRange:beforeString:afterString:"), pasteString, charRangeToReplace, beforeString, afterString)
}/* debug [instance_methods/method]: SmartInsertForStringReplacingRangeBeforeStringAfterString */


// Speaks the selected text, or all text if no selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/startSpeaking(_:)
func (t_ TextView) StartSpeaking(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("startSpeaking:"), sender)
}/* debug [instance_methods/method]: StartSpeaking */


// Stops the speaking of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/stopSpeaking(_:)
func (t_ TextView) StopSpeaking(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("stopSpeaking:"), sender)
}/* debug [instance_methods/method]: StopSpeaking */


// Decreases the space between glyphs in the receiver’s selection, or for all glyphs if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/tightenKerning(_:)
func (t_ TextView) TightenKerning(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("tightenKerning:"), sender)
}/* debug [instance_methods/method]: TightenKerning */


// Toggles the state of the automatic dash substitution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticDashSubstitution(_:)
func (t_ TextView) ToggleAutomaticDashSubstitution(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleAutomaticDashSubstitution:"), sender)
}/* debug [instance_methods/method]: ToggleAutomaticDashSubstitution */


// Toggles the state of the automatic data detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticDataDetection(_:)
func (t_ TextView) ToggleAutomaticDataDetection(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleAutomaticDataDetection:"), sender)
}/* debug [instance_methods/method]: ToggleAutomaticDataDetection */


// Changes the state of automatic link detection from enabled to disabled and vice versa.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticLinkDetection(_:)
func (t_ TextView) ToggleAutomaticLinkDetection(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleAutomaticLinkDetection:"), sender)
}/* debug [instance_methods/method]: ToggleAutomaticLinkDetection */


// Changes the state of automatic quotation mark substitution from enabled to disabled and vice versa.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticQuoteSubstitution(_:)
func (t_ TextView) ToggleAutomaticQuoteSubstitution(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleAutomaticQuoteSubstitution:"), sender)
}/* debug [instance_methods/method]: ToggleAutomaticQuoteSubstitution */


// Toggles the state of the automatic spelling correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticSpellingCorrection(_:)
func (t_ TextView) ToggleAutomaticSpellingCorrection(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleAutomaticSpellingCorrection:"), sender)
}/* debug [instance_methods/method]: ToggleAutomaticSpellingCorrection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticTextCompletion(_:)
func (t_ TextView) ToggleAutomaticTextCompletion(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleAutomaticTextCompletion:"), sender)
}/* debug [instance_methods/method]: ToggleAutomaticTextCompletion */


// Toggles the state of the automatic text replacement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticTextReplacement(_:)
func (t_ TextView) ToggleAutomaticTextReplacement(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleAutomaticTextReplacement:"), sender)
}/* debug [instance_methods/method]: ToggleAutomaticTextReplacement */


// Toggles whether continuous spell checking is enabled for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleContinuousSpellChecking(_:)
func (t_ TextView) ToggleContinuousSpellChecking(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleContinuousSpellChecking:"), sender)
}/* debug [instance_methods/method]: ToggleContinuousSpellChecking */


// Changes the state of grammar checking from enabled to disabled and vice versa.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleGrammarChecking(_:)
func (t_ TextView) ToggleGrammarChecking(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleGrammarChecking:"), sender)
}/* debug [instance_methods/method]: ToggleGrammarChecking */


// An action message that toggles the visibility state of the Quick Look preview panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleQuickLookPreviewPanel(_:)
func (t_ TextView) ToggleQuickLookPreviewPanel(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleQuickLookPreviewPanel:"), sender)
}/* debug [instance_methods/method]: ToggleQuickLookPreviewPanel */


// Changes the state of smart insert and delete from enabled to disabled and vice versa.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/toggleSmartInsertDelete(_:)
func (t_ TextView) ToggleSmartInsertDelete(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("toggleSmartInsertDelete:"), sender)
}/* debug [instance_methods/method]: ToggleSmartInsertDelete */


// Sets the receiver to use nominal glyph spacing for the glyphs in its selection, or for all glyphs if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/turnOffKerning(_:)
func (t_ TextView) TurnOffKerning(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("turnOffKerning:"), sender)
}/* debug [instance_methods/method]: TurnOffKerning */


// Sets the receiver to use only required ligatures when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/turnOffLigatures(_:)
func (t_ TextView) TurnOffLigatures(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("turnOffLigatures:"), sender)
}/* debug [instance_methods/method]: TurnOffLigatures */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/updateCandidates()
func (t_ TextView) UpdateCandidates() {
	objc.Send[objc.ID](t_.ID, objc.Sel("updateCandidates"))
}/* debug [instance_methods/method]: UpdateCandidates */


// Updates the acceptable drag types of all text views associated with the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/updateDragTypeRegistration()
func (t_ TextView) UpdateDragTypeRegistration() {
	objc.Send[objc.ID](t_.ID, objc.Sel("updateDragTypeRegistration"))
}/* debug [instance_methods/method]: UpdateDragTypeRegistration */


// Updates the Font panel to contain the font attributes of the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/updateFontPanel()
func (t_ TextView) UpdateFontPanel() {
	objc.Send[objc.ID](t_.ID, objc.Sel("updateFontPanel"))
}/* debug [instance_methods/method]: UpdateFontPanel */


// Updates the insertion point’s location and optionally restarts the blinking cursor timer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/updateInsertionPointStateAndRestartTimer(_:)
func (t_ TextView) UpdateInsertionPointStateAndRestartTimer(restartFlag bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("updateInsertionPointStateAndRestartTimer:"), restartFlag)
}/* debug [instance_methods/method]: UpdateInsertionPointStateAndRestartTimer */


// Notifies the QuickLook panel that an update may be required.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/updateQuickLookPreviewPanel()
func (t_ TextView) UpdateQuickLookPreviewPanel() {
	objc.Send[objc.ID](t_.ID, objc.Sel("updateQuickLookPreviewPanel"))
}/* debug [instance_methods/method]: UpdateQuickLookPreviewPanel */


// Updates the ruler view in the receiver’s enclosing scroll view to reflect the selection’s paragraph and marker attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/updateRuler()
func (t_ TextView) UpdateRuler() {
	objc.Send[objc.ID](t_.ID, objc.Sel("updateRuler"))
}/* debug [instance_methods/method]: UpdateRuler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/updateTextTouchBarItems()
func (t_ TextView) UpdateTextTouchBarItems() {
	objc.Send[objc.ID](t_.ID, objc.Sel("updateTextTouchBarItems"))
}/* debug [instance_methods/method]: UpdateTextTouchBarItems */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/updateTouchBarItemIdentifiers()
func (t_ TextView) UpdateTouchBarItemIdentifiers() {
	objc.Send[objc.ID](t_.ID, objc.Sel("updateTouchBarItemIdentifiers"))
}/* debug [instance_methods/method]: UpdateTouchBarItemIdentifiers */


// Sets the receiver to use all ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/useAllLigatures(_:)
func (t_ TextView) UseAllLigatures(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("useAllLigatures:"), sender)
}/* debug [instance_methods/method]: UseAllLigatures */


// Set the receiver to use pair kerning data for the glyphs in its selection, or for all glyphs if the receiver is a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/useStandardKerning(_:)
func (t_ TextView) UseStandardKerning(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("useStandardKerning:"), sender)
}/* debug [instance_methods/method]: UseStandardKerning */


// Sets the receiver to use the standard ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/useStandardLigatures(_:)
func (t_ TextView) UseStandardLigatures(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("useStandardLigatures:"), sender)
}/* debug [instance_methods/method]: UseStandardLigatures */


// Returns if the text view can provide and accept the specified data types, or if it can’t.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/validRequestor(forSendType:returnType:)
func (t_ TextView) ValidRequestorForSendTypeReturnType(sendType PasteboardType /* typedef */, returnType PasteboardType /* typedef */) objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("validRequestorForSendType:returnType:"), sendType, returnType)
	return rv
}/* debug [instance_methods/method]: ValidRequestorForSendTypeReturnType */


// Writes the current selection to the specified pasteboard using the given type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/writeSelection(to:type:)
func (t_ TextView) WriteSelectionToPasteboardType(pboard IPasteboard, type_ PasteboardType /* typedef */) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("writeSelectionToPasteboard:type:"), pboard, type_)
	return rv
}/* debug [instance_methods/method]: WriteSelectionToPasteboardType */


// Writes the current selection to the specified pasteboard under each given type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/writeSelection(to:types:)
func (t_ TextView) WriteSelectionToPasteboardTypes(pboard IPasteboard, types []string) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("writeSelectionToPasteboard:types:"), pboard, types)
	return rv
}/* debug [instance_methods/method]: WriteSelectionToPasteboardTypes */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextView */

// The data types that the receiver accepts as the destination view of a dragging operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/acceptableDragTypes
func (t_ TextView) AcceptableDragTypes() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("acceptableDragTypes"))
	return rv
}/* debug [instance_properties/getter]: acceptableDragTypes */


// A Boolean value that indicates whether the receiver accepts the glyph info attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/acceptsGlyphInfo
func (t_ TextView) AcceptsGlyphInfo() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("acceptsGlyphInfo"))
	return rv
}/* debug [instance_properties/getter]: acceptsGlyphInfo */


// A Boolean value that indicates whether the receiver accepts the glyph info attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/acceptsGlyphInfo
func (t_ TextView) SetAcceptsGlyphInfo(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAcceptsGlyphInfo:"), value)
}/* debug [instance_properties/setter]: acceptsGlyphInfo */


// An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowedInputSourceLocales
func (t_ TextView) AllowedInputSourceLocales() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("allowedInputSourceLocales"))
	return rv
}/* debug [instance_properties/getter]: allowedInputSourceLocales */


// An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowedInputSourceLocales
func (t_ TextView) SetAllowedInputSourceLocales(value []string) {
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
}/* debug [instance_properties/setter]: allowedInputSourceLocales */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowedWritingToolsResultOptions
func (t_ TextView) AllowedWritingToolsResultOptions() WritingToolsResultOptions {
	rv := objc.Send[WritingToolsResultOptions](t_.ID, objc.Sel("allowedWritingToolsResultOptions"))
	return rv
}/* debug [instance_properties/getter]: allowedWritingToolsResultOptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowedWritingToolsResultOptions
func (t_ TextView) SetAllowedWritingToolsResultOptions(value WritingToolsResultOptions) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedWritingToolsResultOptions:"), value)
}/* debug [instance_properties/setter]: allowedWritingToolsResultOptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowsCharacterPickerTouchBarItem
func (t_ TextView) AllowsCharacterPickerTouchBarItem() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsCharacterPickerTouchBarItem"))
	return rv
}/* debug [instance_properties/getter]: allowsCharacterPickerTouchBarItem */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowsCharacterPickerTouchBarItem
func (t_ TextView) SetAllowsCharacterPickerTouchBarItem(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsCharacterPickerTouchBarItem:"), value)
}/* debug [instance_properties/setter]: allowsCharacterPickerTouchBarItem */


// A Boolean value that indicates whether the receiver allows its background color to change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowsDocumentBackgroundColorChange
func (t_ TextView) AllowsDocumentBackgroundColorChange() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsDocumentBackgroundColorChange"))
	return rv
}/* debug [instance_properties/getter]: allowsDocumentBackgroundColorChange */


// A Boolean value that indicates whether the receiver allows its background color to change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowsDocumentBackgroundColorChange
func (t_ TextView) SetAllowsDocumentBackgroundColorChange(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsDocumentBackgroundColorChange:"), value)
}/* debug [instance_properties/setter]: allowsDocumentBackgroundColorChange */


// Indicates whether image attachments should permit editing of their images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowsImageEditing
func (t_ TextView) AllowsImageEditing() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsImageEditing"))
	return rv
}/* debug [instance_properties/getter]: allowsImageEditing */


// Indicates whether image attachments should permit editing of their images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowsImageEditing
func (t_ TextView) SetAllowsImageEditing(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsImageEditing:"), value)
}/* debug [instance_properties/setter]: allowsImageEditing */


// A Boolean value that indicates whether the receiver allows undo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowsUndo
func (t_ TextView) AllowsUndo() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsUndo"))
	return rv
}/* debug [instance_properties/getter]: allowsUndo */


// A Boolean value that indicates whether the receiver allows undo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/allowsUndo
func (t_ TextView) SetAllowsUndo(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsUndo:"), value)
}/* debug [instance_properties/setter]: allowsUndo */


// The receiver’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/backgroundColor
func (t_ TextView) BackgroundColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The receiver’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/backgroundColor
func (t_ TextView) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/candidateListTouchBarItem
func (t_ TextView) CandidateListTouchBarItem() ICandidateListTouchBarItem {
	rv := objc.Send[CandidateListTouchBarItem](t_.ID, objc.Sel("candidateListTouchBarItem"))
	return rv
}/* debug [instance_properties/getter]: candidateListTouchBarItem */


// The receiver’s default paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/defaultParagraphStyle
func (t_ TextView) DefaultParagraphStyle() IParagraphStyle {
	rv := objc.Send[ParagraphStyle](t_.ID, objc.Sel("defaultParagraphStyle"))
	return rv
}/* debug [instance_properties/getter]: defaultParagraphStyle */


// The receiver’s default paragraph style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/defaultParagraphStyle
func (t_ TextView) SetDefaultParagraphStyle(value IParagraphStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDefaultParagraphStyle:"), value)
}/* debug [instance_properties/setter]: defaultParagraphStyle */


// The delegate for all text views sharing the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/delegate
func (t_ TextView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for all text views sharing the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/delegate
func (t_ TextView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/displaysLinkToolTips
func (t_ TextView) DisplaysLinkToolTips() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("displaysLinkToolTips"))
	return rv
}/* debug [instance_properties/getter]: displaysLinkToolTips */


// A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/displaysLinkToolTips
func (t_ TextView) SetDisplaysLinkToolTips(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDisplaysLinkToolTips:"), value)
}/* debug [instance_properties/setter]: displaysLinkToolTips */


// A Boolean value that indicates whether the receiver draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/drawsBackground
func (t_ TextView) DrawsBackground() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("drawsBackground"))
	return rv
}/* debug [instance_properties/getter]: drawsBackground */


// A Boolean value that indicates whether the receiver draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/drawsBackground
func (t_ TextView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDrawsBackground:"), value)
}/* debug [instance_properties/setter]: drawsBackground */


// The default text checking types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/enabledTextCheckingTypes
func (t_ TextView) EnabledTextCheckingTypes() TextCheckingTypes /* not a class type */ {
	rv := objc.Send[TextCheckingTypes](t_.ID, objc.Sel("enabledTextCheckingTypes"))
	return rv
}/* debug [instance_properties/getter]: enabledTextCheckingTypes */


// The default text checking types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/enabledTextCheckingTypes
func (t_ TextView) SetEnabledTextCheckingTypes(value TextCheckingTypes /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEnabledTextCheckingTypes:"), value)
}/* debug [instance_properties/setter]: enabledTextCheckingTypes */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to import files by dragging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/importsGraphics
func (t_ TextView) ImportsGraphics() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("importsGraphics"))
	return rv
}/* debug [instance_properties/getter]: importsGraphics */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to import files by dragging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/importsGraphics
func (t_ TextView) SetImportsGraphics(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImportsGraphics:"), value)
}/* debug [instance_properties/setter]: importsGraphics */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/inlinePredictionType
func (t_ TextView) InlinePredictionType() TextInputTraitType {
	rv := objc.Send[TextInputTraitType](t_.ID, objc.Sel("inlinePredictionType"))
	return rv
}/* debug [instance_properties/getter]: inlinePredictionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/inlinePredictionType
func (t_ TextView) SetInlinePredictionType(value TextInputTraitType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInlinePredictionType:"), value)
}/* debug [instance_properties/setter]: inlinePredictionType */


// The color of the insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/insertionPointColor
func (t_ TextView) InsertionPointColor() IColor {
	rv := objc.Send[Color](t_.ID, objc.Sel("insertionPointColor"))
	return rv
}/* debug [instance_properties/getter]: insertionPointColor */


// The color of the insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/insertionPointColor
func (t_ TextView) SetInsertionPointColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInsertionPointColor:"), value)
}/* debug [instance_properties/setter]: insertionPointColor */


// A Boolean value that indicates whether automatic dash substitution is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticDashSubstitutionEnabled
func (t_ TextView) AutomaticDashSubstitutionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticDashSubstitutionEnabled"))
	return rv
}/* debug [instance_properties/getter]: automaticDashSubstitutionEnabled */


// A Boolean value that indicates whether automatic dash substitution is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticDashSubstitutionEnabled
func (t_ TextView) SetAutomaticDashSubstitutionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticDashSubstitutionEnabled:"), value)
}/* debug [instance_properties/setter]: automaticDashSubstitutionEnabled */


// A Boolean value that indicates whether automatic data detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticDataDetectionEnabled
func (t_ TextView) AutomaticDataDetectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticDataDetectionEnabled"))
	return rv
}/* debug [instance_properties/getter]: automaticDataDetectionEnabled */


// A Boolean value that indicates whether automatic data detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticDataDetectionEnabled
func (t_ TextView) SetAutomaticDataDetectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticDataDetectionEnabled:"), value)
}/* debug [instance_properties/setter]: automaticDataDetectionEnabled */


// A Boolean value that enables or disables automatic link detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticLinkDetectionEnabled
func (t_ TextView) AutomaticLinkDetectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticLinkDetectionEnabled"))
	return rv
}/* debug [instance_properties/getter]: automaticLinkDetectionEnabled */


// A Boolean value that enables or disables automatic link detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticLinkDetectionEnabled
func (t_ TextView) SetAutomaticLinkDetectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticLinkDetectionEnabled:"), value)
}/* debug [instance_properties/setter]: automaticLinkDetectionEnabled */


// A Boolean value that enables and disables automatic quotation mark substitution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticQuoteSubstitutionEnabled
func (t_ TextView) AutomaticQuoteSubstitutionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticQuoteSubstitutionEnabled"))
	return rv
}/* debug [instance_properties/getter]: automaticQuoteSubstitutionEnabled */


// A Boolean value that enables and disables automatic quotation mark substitution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticQuoteSubstitutionEnabled
func (t_ TextView) SetAutomaticQuoteSubstitutionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticQuoteSubstitutionEnabled:"), value)
}/* debug [instance_properties/setter]: automaticQuoteSubstitutionEnabled */


// A Boolean value that indicates whether automatic spelling correction is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticSpellingCorrectionEnabled
func (t_ TextView) AutomaticSpellingCorrectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticSpellingCorrectionEnabled"))
	return rv
}/* debug [instance_properties/getter]: automaticSpellingCorrectionEnabled */


// A Boolean value that indicates whether automatic spelling correction is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticSpellingCorrectionEnabled
func (t_ TextView) SetAutomaticSpellingCorrectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticSpellingCorrectionEnabled:"), value)
}/* debug [instance_properties/setter]: automaticSpellingCorrectionEnabled */


// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticTextCompletionEnabled
func (t_ TextView) AutomaticTextCompletionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticTextCompletionEnabled"))
	return rv
}/* debug [instance_properties/getter]: automaticTextCompletionEnabled */


// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticTextCompletionEnabled
func (t_ TextView) SetAutomaticTextCompletionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticTextCompletionEnabled:"), value)
}/* debug [instance_properties/setter]: automaticTextCompletionEnabled */


// A Boolean value that indicates whether automatic text replacement is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticTextReplacementEnabled
func (t_ TextView) AutomaticTextReplacementEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticTextReplacementEnabled"))
	return rv
}/* debug [instance_properties/getter]: automaticTextReplacementEnabled */


// A Boolean value that indicates whether automatic text replacement is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticTextReplacementEnabled
func (t_ TextView) SetAutomaticTextReplacementEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticTextReplacementEnabled:"), value)
}/* debug [instance_properties/setter]: automaticTextReplacementEnabled */


// A Boolean value that indicates whether undo coalescing is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isCoalescingUndo
func (t_ TextView) CoalescingUndo() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("coalescingUndo"))
	return rv
}/* debug [instance_properties/getter]: coalescingUndo */


// A Boolean value that indicates whether the receiver has continuous spell checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isContinuousSpellCheckingEnabled
func (t_ TextView) ContinuousSpellCheckingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("continuousSpellCheckingEnabled"))
	return rv
}/* debug [instance_properties/getter]: continuousSpellCheckingEnabled */


// A Boolean value that indicates whether the receiver has continuous spell checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isContinuousSpellCheckingEnabled
func (t_ TextView) SetContinuousSpellCheckingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContinuousSpellCheckingEnabled:"), value)
}/* debug [instance_properties/setter]: continuousSpellCheckingEnabled */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isEditable
func (t_ TextView) Editable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("editable"))
	return rv
}/* debug [instance_properties/getter]: editable */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isEditable
func (t_ TextView) SetEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEditable:"), value)
}/* debug [instance_properties/setter]: editable */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isFieldEditor
func (t_ TextView) FieldEditor() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("fieldEditor"))
	return rv
}/* debug [instance_properties/getter]: fieldEditor */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isFieldEditor
func (t_ TextView) SetFieldEditor(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setFieldEditor:"), value)
}/* debug [instance_properties/setter]: fieldEditor */


// Enables and disables grammar checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isGrammarCheckingEnabled
func (t_ TextView) GrammarCheckingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("grammarCheckingEnabled"))
	return rv
}/* debug [instance_properties/getter]: grammarCheckingEnabled */


// Enables and disables grammar checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isGrammarCheckingEnabled
func (t_ TextView) SetGrammarCheckingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGrammarCheckingEnabled:"), value)
}/* debug [instance_properties/setter]: grammarCheckingEnabled */


// A Boolean value that indicates whether incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isIncrementalSearchingEnabled
func (t_ TextView) IncrementalSearchingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("incrementalSearchingEnabled"))
	return rv
}/* debug [instance_properties/getter]: incrementalSearchingEnabled */


// A Boolean value that indicates whether incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isIncrementalSearchingEnabled
func (t_ TextView) SetIncrementalSearchingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIncrementalSearchingEnabled:"), value)
}/* debug [instance_properties/setter]: incrementalSearchingEnabled */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isRichText
func (t_ TextView) RichText() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("richText"))
	return rv
}/* debug [instance_properties/getter]: richText */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isRichText
func (t_ TextView) SetRichText(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRichText:"), value)
}/* debug [instance_properties/setter]: richText */


// A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isRulerVisible
func (t_ TextView) RulerVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("rulerVisible"))
	return rv
}/* debug [instance_properties/getter]: rulerVisible */


// A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isRulerVisible
func (t_ TextView) SetRulerVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setRulerVisible:"), value)
}/* debug [instance_properties/setter]: rulerVisible */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isSelectable
func (t_ TextView) Selectable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("selectable"))
	return rv
}/* debug [instance_properties/getter]: selectable */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isSelectable
func (t_ TextView) SetSelectable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectable:"), value)
}/* debug [instance_properties/setter]: selectable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/isWritingToolsActive
func (t_ TextView) WritingToolsActive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("writingToolsActive"))
	return rv
}/* debug [instance_properties/getter]: writingToolsActive */


// The layout manager that lays out text for the receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/layoutManager
func (t_ TextView) LayoutManager() ILayoutManager {
	rv := objc.Send[LayoutManager](t_.ID, objc.Sel("layoutManager"))
	return rv
}/* debug [instance_properties/getter]: layoutManager */


// The attributes used to draw the onscreen presentation of link text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/linkTextAttributes
func (t_ TextView) LinkTextAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("linkTextAttributes"))
	return rv
}/* debug [instance_properties/getter]: linkTextAttributes */


// The attributes used to draw the onscreen presentation of link text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/linkTextAttributes
func (t_ TextView) SetLinkTextAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLinkTextAttributes:"), value)
}/* debug [instance_properties/setter]: linkTextAttributes */


// The attributes used to draw marked text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/markedTextAttributes
func (t_ TextView) MarkedTextAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("markedTextAttributes"))
	return rv
}/* debug [instance_properties/getter]: markedTextAttributes */


// The attributes used to draw marked text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/markedTextAttributes
func (t_ TextView) SetMarkedTextAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMarkedTextAttributes:"), value)
}/* debug [instance_properties/setter]: markedTextAttributes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/mathExpressionCompletionType
func (t_ TextView) MathExpressionCompletionType() TextInputTraitType {
	rv := objc.Send[TextInputTraitType](t_.ID, objc.Sel("mathExpressionCompletionType"))
	return rv
}/* debug [instance_properties/getter]: mathExpressionCompletionType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/mathExpressionCompletionType
func (t_ TextView) SetMathExpressionCompletionType(value TextInputTraitType) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMathExpressionCompletionType:"), value)
}/* debug [instance_properties/setter]: mathExpressionCompletionType */


// The range of characters affected by an action method that changes character (not paragraph) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rangeForUserCharacterAttributeChange
func (t_ TextView) RangeForUserCharacterAttributeChange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("rangeForUserCharacterAttributeChange"))
	return rv
}/* debug [instance_properties/getter]: rangeForUserCharacterAttributeChange */


// The partial range from the most recent beginning of a word up to the insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rangeForUserCompletion
func (t_ TextView) RangeForUserCompletion() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("rangeForUserCompletion"))
	return rv
}/* debug [instance_properties/getter]: rangeForUserCompletion */


// The range of characters affected by an action method that changes paragraph (not character) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rangeForUserParagraphAttributeChange
func (t_ TextView) RangeForUserParagraphAttributeChange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("rangeForUserParagraphAttributeChange"))
	return rv
}/* debug [instance_properties/getter]: rangeForUserParagraphAttributeChange */


// The range of characters affected by a method that changes characters (as opposed to attributes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rangeForUserTextChange
func (t_ TextView) RangeForUserTextChange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](t_.ID, objc.Sel("rangeForUserTextChange"))
	return rv
}/* debug [instance_properties/getter]: rangeForUserTextChange */


// An array containing the ranges of characters affected by an action method that changes character (not paragraph) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rangesForUserCharacterAttributeChange
func (t_ TextView) RangesForUserCharacterAttributeChange() []foundation.Value {
	rv := objc.Send[[]foundation.Value](t_.ID, objc.Sel("rangesForUserCharacterAttributeChange"))
	return rv
}/* debug [instance_properties/getter]: rangesForUserCharacterAttributeChange */


// An array containing the ranges of characters affected by a method that changes paragraph (not character) attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rangesForUserParagraphAttributeChange
func (t_ TextView) RangesForUserParagraphAttributeChange() []foundation.Value {
	rv := objc.Send[[]foundation.Value](t_.ID, objc.Sel("rangesForUserParagraphAttributeChange"))
	return rv
}/* debug [instance_properties/getter]: rangesForUserParagraphAttributeChange */


// An array containing the ranges of characters affected by a method that changes characters (as opposed to attributes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/rangesForUserTextChange
func (t_ TextView) RangesForUserTextChange() []foundation.Value {
	rv := objc.Send[[]foundation.Value](t_.ID, objc.Sel("rangesForUserTextChange"))
	return rv
}/* debug [instance_properties/getter]: rangesForUserTextChange */


// The types this text view can read immediately from the pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/readablePasteboardTypes
func (t_ TextView) ReadablePasteboardTypes() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("readablePasteboardTypes"))
	return rv
}/* debug [instance_properties/getter]: readablePasteboardTypes */


// An array containing the ranges of characters selected in the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectedRanges
func (t_ TextView) SelectedRanges() []foundation.Value {
	rv := objc.Send[[]foundation.Value](t_.ID, objc.Sel("selectedRanges"))
	return rv
}/* debug [instance_properties/getter]: selectedRanges */


// An array containing the ranges of characters selected in the receiver’s layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectedRanges
func (t_ TextView) SetSelectedRanges(value []foundation.Value) {
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
}/* debug [instance_properties/setter]: selectedRanges */


// The attributes used to indicate the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectedTextAttributes
func (t_ TextView) SelectedTextAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("selectedTextAttributes"))
	return rv
}/* debug [instance_properties/getter]: selectedTextAttributes */


// The attributes used to indicate the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectedTextAttributes
func (t_ TextView) SetSelectedTextAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedTextAttributes:"), value)
}/* debug [instance_properties/setter]: selectedTextAttributes */


// The preferred direction of selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectionAffinity
func (t_ TextView) SelectionAffinity() SelectionAffinity {
	rv := objc.Send[SelectionAffinity](t_.ID, objc.Sel("selectionAffinity"))
	return rv
}/* debug [instance_properties/getter]: selectionAffinity */


// The selection granularity for subsequent extension of a selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectionGranularity
func (t_ TextView) SelectionGranularity() SelectionGranularity {
	rv := objc.Send[SelectionGranularity](t_.ID, objc.Sel("selectionGranularity"))
	return rv
}/* debug [instance_properties/getter]: selectionGranularity */


// The selection granularity for subsequent extension of a selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/selectionGranularity
func (t_ TextView) SetSelectionGranularity(value SelectionGranularity) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionGranularity:"), value)
}/* debug [instance_properties/setter]: selectionGranularity */


// A Boolean value that determines whether the receiver should draw its insertion point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/shouldDrawInsertionPoint
func (t_ TextView) ShouldDrawInsertionPoint() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("shouldDrawInsertionPoint"))
	return rv
}/* debug [instance_properties/getter]: shouldDrawInsertionPoint */


// A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/smartInsertDeleteEnabled
func (t_ TextView) SmartInsertDeleteEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("smartInsertDeleteEnabled"))
	return rv
}/* debug [instance_properties/getter]: smartInsertDeleteEnabled */


// A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/smartInsertDeleteEnabled
func (t_ TextView) SetSmartInsertDeleteEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSmartInsertDeleteEnabled:"), value)
}/* debug [instance_properties/setter]: smartInsertDeleteEnabled */


// A tag identifying the text view’s text as a document for the spell checker server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/spellCheckerDocumentTag
func (t_ TextView) SpellCheckerDocumentTag() int {
	rv := objc.Send[int](t_.ID, objc.Sel("spellCheckerDocumentTag"))
	return rv
}/* debug [instance_properties/getter]: spellCheckerDocumentTag */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/stronglyReferencesTextStorage
func (t_ TextView) StronglyReferencesTextStorage() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("stronglyReferencesTextStorage"))
	return rv
}/* debug [instance_properties/getter]: stronglyReferencesTextStorage */


// The receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContainer
func (t_ TextView) TextContainer() ITextContainer {
	rv := objc.Send[TextContainer](t_.ID, objc.Sel("textContainer"))
	return rv
}/* debug [instance_properties/getter]: textContainer */


// The receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContainer
func (t_ TextView) SetTextContainer(value ITextContainer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainer:"), value)
}/* debug [instance_properties/setter]: textContainer */


// The empty space the receiver leaves around its associated text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContainerInset
func (t_ TextView) TextContainerInset() Size /* not a class type */ {
	rv := objc.Send[Size](t_.ID, objc.Sel("textContainerInset"))
	return rv
}/* debug [instance_properties/getter]: textContainerInset */


// The empty space the receiver leaves around its associated text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContainerInset
func (t_ TextView) SetTextContainerInset(value Size /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextContainerInset:"), value)
}/* debug [instance_properties/setter]: textContainerInset */


// The origin of the receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContainerOrigin
func (t_ TextView) TextContainerOrigin() vision.Point {
	rv := objc.Send[vision.Point](t_.ID, objc.Sel("textContainerOrigin"))
	return rv
}/* debug [instance_properties/getter]: textContainerOrigin */


// The receiver’s text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textContentStorage
func (t_ TextView) TextContentStorage() ITextContentStorage {
	rv := objc.Send[TextContentStorage](t_.ID, objc.Sel("textContentStorage"))
	return rv
}/* debug [instance_properties/getter]: textContentStorage */


// ************************* Text Highlight support **************************
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textHighlightAttributes
func (t_ TextView) TextHighlightAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("textHighlightAttributes"))
	return rv
}/* debug [instance_properties/getter]: textHighlightAttributes */


// ************************* Text Highlight support **************************
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textHighlightAttributes
func (t_ TextView) SetTextHighlightAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTextHighlightAttributes:"), value)
}/* debug [instance_properties/setter]: textHighlightAttributes */


// The manager that lays out text for the receiver’s text container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textLayoutManager
func (t_ TextView) TextLayoutManager() ITextLayoutManager {
	rv := objc.Send[TextLayoutManager](t_.ID, objc.Sel("textLayoutManager"))
	return rv
}/* debug [instance_properties/getter]: textLayoutManager */


// The receiver’s text storage object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/textStorage
func (t_ TextView) TextStorage() ITextStorage {
	rv := objc.Send[TextStorage](t_.ID, objc.Sel("textStorage"))
	return rv
}/* debug [instance_properties/getter]: textStorage */


// The receiver’s typing attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/typingAttributes
func (t_ TextView) TypingAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("typingAttributes"))
	return rv
}/* debug [instance_properties/getter]: typingAttributes */


// The receiver’s typing attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/typingAttributes
func (t_ TextView) SetTypingAttributes(value foundation.IDictionary) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTypingAttributes:"), value)
}/* debug [instance_properties/setter]: typingAttributes */


// A Boolean value that indicates whether the framework should use adaptive color mapping for dark appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesAdaptiveColorMappingForDarkAppearance
func (t_ TextView) UsesAdaptiveColorMappingForDarkAppearance() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesAdaptiveColorMappingForDarkAppearance"))
	return rv
}/* debug [instance_properties/getter]: usesAdaptiveColorMappingForDarkAppearance */


// A Boolean value that indicates whether the framework should use adaptive color mapping for dark appearance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesAdaptiveColorMappingForDarkAppearance
func (t_ TextView) SetUsesAdaptiveColorMappingForDarkAppearance(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesAdaptiveColorMappingForDarkAppearance:"), value)
}/* debug [instance_properties/setter]: usesAdaptiveColorMappingForDarkAppearance */


// A Boolean value that indicates whether to use the find bar for this text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesFindBar
func (t_ TextView) UsesFindBar() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFindBar"))
	return rv
}/* debug [instance_properties/getter]: usesFindBar */


// A Boolean value that indicates whether to use the find bar for this text view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesFindBar
func (t_ TextView) SetUsesFindBar(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFindBar:"), value)
}/* debug [instance_properties/setter]: usesFindBar */


// A Boolean value that indicates whether the receiver allows for a find panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesFindPanel
func (t_ TextView) UsesFindPanel() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFindPanel"))
	return rv
}/* debug [instance_properties/getter]: usesFindPanel */


// A Boolean value that indicates whether the receiver allows for a find panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesFindPanel
func (t_ TextView) SetUsesFindPanel(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFindPanel:"), value)
}/* debug [instance_properties/setter]: usesFindPanel */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager use the Font panel and Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesFontPanel
func (t_ TextView) UsesFontPanel() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesFontPanel"))
	return rv
}/* debug [instance_properties/getter]: usesFontPanel */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager use the Font panel and Font menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesFontPanel
func (t_ TextView) SetUsesFontPanel(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesFontPanel:"), value)
}/* debug [instance_properties/setter]: usesFontPanel */


// A Boolean value that indicates whether this text view uses the inspector bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesInspectorBar
func (t_ TextView) UsesInspectorBar() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesInspectorBar"))
	return rv
}/* debug [instance_properties/getter]: usesInspectorBar */


// A Boolean value that indicates whether this text view uses the inspector bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesInspectorBar
func (t_ TextView) SetUsesInspectorBar(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesInspectorBar:"), value)
}/* debug [instance_properties/setter]: usesInspectorBar */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesRolloverButtonForSelection
func (t_ TextView) UsesRolloverButtonForSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesRolloverButtonForSelection"))
	return rv
}/* debug [instance_properties/getter]: usesRolloverButtonForSelection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesRolloverButtonForSelection
func (t_ TextView) SetUsesRolloverButtonForSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesRolloverButtonForSelection:"), value)
}/* debug [instance_properties/setter]: usesRolloverButtonForSelection */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesRuler
func (t_ TextView) UsesRuler() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("usesRuler"))
	return rv
}/* debug [instance_properties/getter]: usesRuler */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/usesRuler
func (t_ TextView) SetUsesRuler(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setUsesRuler:"), value)
}/* debug [instance_properties/setter]: usesRuler */


// The pasteboard types that can be provided from the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/writablePasteboardTypes
func (t_ TextView) WritablePasteboardTypes() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("writablePasteboardTypes"))
	return rv
}/* debug [instance_properties/getter]: writablePasteboardTypes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/writingToolsBehavior
func (t_ TextView) WritingToolsBehavior() WritingToolsBehavior {
	rv := objc.Send[WritingToolsBehavior](t_.ID, objc.Sel("writingToolsBehavior"))
	return rv
}/* debug [instance_properties/getter]: writingToolsBehavior */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextView/writingToolsBehavior
func (t_ TextView) SetWritingToolsBehavior(value WritingToolsBehavior) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setWritingToolsBehavior:"), value)
}/* debug [instance_properties/setter]: writingToolsBehavior */


// A Boolean value that indicates whether automatic dash substitution is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdashsubstitutionenabled
func (t_ TextView) IsAutomaticDashSubstitutionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticDashSubstitutionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAutomaticDashSubstitutionEnabled */


// A Boolean value that indicates whether automatic dash substitution is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdashsubstitutionenabled
func (t_ TextView) SetIsAutomaticDashSubstitutionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticDashSubstitutionEnabled:"), value)
}/* debug [instance_properties/setter]: isAutomaticDashSubstitutionEnabled */


// A Boolean value that indicates whether automatic data detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdatadetectionenabled
func (t_ TextView) IsAutomaticDataDetectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticDataDetectionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAutomaticDataDetectionEnabled */


// A Boolean value that indicates whether automatic data detection is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticdatadetectionenabled
func (t_ TextView) SetIsAutomaticDataDetectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticDataDetectionEnabled:"), value)
}/* debug [instance_properties/setter]: isAutomaticDataDetectionEnabled */


// A Boolean value that enables or disables automatic link detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticlinkdetectionenabled
func (t_ TextView) IsAutomaticLinkDetectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticLinkDetectionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAutomaticLinkDetectionEnabled */


// A Boolean value that enables or disables automatic link detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticlinkdetectionenabled
func (t_ TextView) SetIsAutomaticLinkDetectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticLinkDetectionEnabled:"), value)
}/* debug [instance_properties/setter]: isAutomaticLinkDetectionEnabled */


// A Boolean value that enables and disables automatic quotation mark substitution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticquotesubstitutionenabled
func (t_ TextView) IsAutomaticQuoteSubstitutionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticQuoteSubstitutionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAutomaticQuoteSubstitutionEnabled */


// A Boolean value that enables and disables automatic quotation mark substitution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticquotesubstitutionenabled
func (t_ TextView) SetIsAutomaticQuoteSubstitutionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticQuoteSubstitutionEnabled:"), value)
}/* debug [instance_properties/setter]: isAutomaticQuoteSubstitutionEnabled */


// A Boolean value that indicates whether automatic spelling correction is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticspellingcorrectionenabled
func (t_ TextView) IsAutomaticSpellingCorrectionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticSpellingCorrectionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAutomaticSpellingCorrectionEnabled */


// A Boolean value that indicates whether automatic spelling correction is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomaticspellingcorrectionenabled
func (t_ TextView) SetIsAutomaticSpellingCorrectionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticSpellingCorrectionEnabled:"), value)
}/* debug [instance_properties/setter]: isAutomaticSpellingCorrectionEnabled */


// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextcompletionenabled
func (t_ TextView) IsAutomaticTextCompletionEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticTextCompletionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAutomaticTextCompletionEnabled */


// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextcompletionenabled
func (t_ TextView) SetIsAutomaticTextCompletionEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticTextCompletionEnabled:"), value)
}/* debug [instance_properties/setter]: isAutomaticTextCompletionEnabled */


// A Boolean value that indicates whether automatic text replacement is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextreplacementenabled
func (t_ TextView) IsAutomaticTextReplacementEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticTextReplacementEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAutomaticTextReplacementEnabled */


// A Boolean value that indicates whether automatic text replacement is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isautomatictextreplacementenabled
func (t_ TextView) SetIsAutomaticTextReplacementEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticTextReplacementEnabled:"), value)
}/* debug [instance_properties/setter]: isAutomaticTextReplacementEnabled */


// A Boolean value that indicates whether undo coalescing is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscoalescingundo
func (t_ TextView) IsCoalescingUndo() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isCoalescingUndo"))
	return rv
}/* debug [instance_properties/getter]: isCoalescingUndo */


// A Boolean value that indicates whether undo coalescing is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscoalescingundo
func (t_ TextView) SetIsCoalescingUndo(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsCoalescingUndo:"), value)
}/* debug [instance_properties/setter]: isCoalescingUndo */


// A Boolean value that indicates whether the receiver has continuous spell checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscontinuousspellcheckingenabled
func (t_ TextView) IsContinuousSpellCheckingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isContinuousSpellCheckingEnabled"))
	return rv
}/* debug [instance_properties/getter]: isContinuousSpellCheckingEnabled */


// A Boolean value that indicates whether the receiver has continuous spell checking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iscontinuousspellcheckingenabled
func (t_ TextView) SetIsContinuousSpellCheckingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsContinuousSpellCheckingEnabled:"), value)
}/* debug [instance_properties/setter]: isContinuousSpellCheckingEnabled */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iseditable
func (t_ TextView) IsEditable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isEditable"))
	return rv
}/* debug [instance_properties/getter]: isEditable */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to edit text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iseditable
func (t_ TextView) SetIsEditable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsEditable:"), value)
}/* debug [instance_properties/setter]: isEditable */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isfieldeditor
func (t_ TextView) IsFieldEditor() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isFieldEditor"))
	return rv
}/* debug [instance_properties/getter]: isFieldEditor */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager behave as field editors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isfieldeditor
func (t_ TextView) SetIsFieldEditor(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsFieldEditor:"), value)
}/* debug [instance_properties/setter]: isFieldEditor */


// Enables and disables grammar checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isgrammarcheckingenabled
func (t_ TextView) IsGrammarCheckingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isGrammarCheckingEnabled"))
	return rv
}/* debug [instance_properties/getter]: isGrammarCheckingEnabled */


// Enables and disables grammar checking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isgrammarcheckingenabled
func (t_ TextView) SetIsGrammarCheckingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsGrammarCheckingEnabled:"), value)
}/* debug [instance_properties/setter]: isGrammarCheckingEnabled */


// A Boolean value that indicates whether incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isincrementalsearchingenabled
func (t_ TextView) IsIncrementalSearchingEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isIncrementalSearchingEnabled"))
	return rv
}/* debug [instance_properties/getter]: isIncrementalSearchingEnabled */


// A Boolean value that indicates whether incremental searching is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isincrementalsearchingenabled
func (t_ TextView) SetIsIncrementalSearchingEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsIncrementalSearchingEnabled:"), value)
}/* debug [instance_properties/setter]: isIncrementalSearchingEnabled */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrichtext
func (t_ TextView) IsRichText() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRichText"))
	return rv
}/* debug [instance_properties/getter]: isRichText */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to apply attributes to specific ranges of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrichtext
func (t_ TextView) SetIsRichText(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRichText:"), value)
}/* debug [instance_properties/setter]: isRichText */


// A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrulervisible
func (t_ TextView) IsRulerVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRulerVisible"))
	return rv
}/* debug [instance_properties/getter]: isRulerVisible */


// A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isrulervisible
func (t_ TextView) SetIsRulerVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRulerVisible:"), value)
}/* debug [instance_properties/setter]: isRulerVisible */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isselectable
func (t_ TextView) IsSelectable() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSelectable"))
	return rv
}/* debug [instance_properties/getter]: isSelectable */


// A Boolean value that controls whether the text views sharing the receiver’s layout manager allow the user to select text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/isselectable
func (t_ TextView) SetIsSelectable(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsSelectable:"), value)
}/* debug [instance_properties/setter]: isSelectable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iswritingtoolsactive
func (t_ TextView) IsWritingToolsActive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isWritingToolsActive"))
	return rv
}/* debug [instance_properties/getter]: isWritingToolsActive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/iswritingtoolsactive
func (t_ TextView) SetIsWritingToolsActive(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsWritingToolsActive:"), value)
}/* debug [instance_properties/setter]: isWritingToolsActive */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextView */


